// Package boltdb contains the BoltDB store implementation.
package boltdb

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/blink-io/x/kvstore"

	"go.etcd.io/bbolt"
)

var (
	// ErrMultipleEndpointsUnsupported is thrown when multiple endpoints specified for BoltDB.
	// Endpoint has to be a local file path.
	ErrMultipleEndpointsUnsupported = errors.New("boltdb supports one endpoint and should be a file path")
	// ErrBoltBucketOptionMissing is thrown when boltBucket config option is missing.
	ErrBoltBucketOptionMissing = errors.New("boltBucket config option missing")
)

// StoreName the name of the store.
const StoreName = "boltdb"

const filePerm os.FileMode = 0o644

const (
	metadataLen      = 8
	transientTimeout = time.Duration(10) * time.Second
)

// registers boltdb to kvstore.
func init() {
	kvstore.Register(StoreName, newStore)
}

// Config the BoltDB configuration.
type Config struct {
	Bucket            string
	PersistConnection bool
	ConnectionTimeout time.Duration
}

func newStore(ctx context.Context, endpoints []string, options kvstore.Config) (kvstore.Store, error) {
	cfg, ok := options.(*Config)
	if !ok && options != nil {
		return nil, &kvstore.InvalidConfigurationError{Store: StoreName, Config: options}
	}

	return New(ctx, endpoints, cfg)
}

// Store implements the store.Store interface.
type Store struct {
	client     *bbolt.DB
	boltBucket []byte
	dbIndex    uint64
	path       string
	timeout    time.Duration
	// By default, kvstore opens and closes the BoltDB connection for every get/put operation.
	// This allows multiple apps to use a BoltDB at the same time.
	// PersistConnection flag provides an option to override ths behavior.
	// ie: open the connection in New and use it till Close is called.
	PersistConnection bool
	mu                sync.Mutex
}

// New creates a new BoltDB client.
func New(_ context.Context, endpoints []string, options *Config) (*Store, error) {
	if len(endpoints) > 1 {
		return nil, ErrMultipleEndpointsUnsupported
	}

	if options == nil || options.Bucket == "" {
		return nil, ErrBoltBucketOptionMissing
	}

	dbPath := endpoints[0]

	err := os.MkdirAll(filepath.Dir(dbPath), 0o750)
	if err != nil {
		return nil, err
	}

	var db *bbolt.DB
	if options.PersistConnection {
		boltOptions := &bbolt.Options{Timeout: options.ConnectionTimeout}
		db, err = bbolt.Open(dbPath, filePerm, boltOptions)
		if err != nil {
			return nil, err
		}
	}

	timeout := transientTimeout
	if options.ConnectionTimeout != 0 {
		timeout = options.ConnectionTimeout
	}

	b := &Store{
		client:            db,
		path:              dbPath,
		boltBucket:        []byte(options.Bucket),
		timeout:           timeout,
		PersistConnection: options.PersistConnection,
	}

	return b, nil
}

// Get the value at "key".
// BoltDB doesn't provide an inbuilt last modified index with every kv pair.
// It's implemented by an atomic counter maintained by the kvstore
// and appended to the value passed by the client.
func (b *Store) Get(_ context.Context, key string, _ *kvstore.ReadOptions) (*kvstore.KVPair, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	db, err := b.getDBHandle()
	if err != nil {
		return nil, err
	}
	defer b.releaseDBHandle()

	var val []byte

	err = db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(b.boltBucket)
		if bucket == nil {
			return kvstore.ErrKeyNotFound
		}

		v := bucket.Get([]byte(key))
		val = make([]byte, len(v))
		copy(val, v)

		return nil
	})

	if len(val) == 0 {
		return nil, kvstore.ErrKeyNotFound
	}
	if err != nil {
		return nil, err
	}

	dbIndex := binary.LittleEndian.Uint64(val[:metadataLen])
	val = val[metadataLen:]

	return &kvstore.KVPair{Key: key, Value: val, LastIndex: dbIndex}, nil
}

// Put the key, value pair.
// Index number metadata is prepended to the value.
func (b *Store) Put(_ context.Context, key string, value []byte, _ *kvstore.WriteOptions) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	dbval := make([]byte, metadataLen)

	db, err := b.getDBHandle()
	if err != nil {
		return err
	}
	defer b.releaseDBHandle()

	return db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(b.boltBucket)
		if err != nil {
			return err
		}

		dbIndex := atomic.AddUint64(&b.dbIndex, 1)
		binary.LittleEndian.PutUint64(dbval, dbIndex)
		dbval = append(dbval, value...)

		err = bucket.Put([]byte(key), dbval)
		if err != nil {
			return err
		}
		return nil
	})
}

// Delete the value for the given key.
func (b *Store) Delete(_ context.Context, key string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	db, err := b.getDBHandle()
	if err != nil {
		return err
	}
	defer b.releaseDBHandle()

	return db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(b.boltBucket)
		if bucket == nil {
			return kvstore.ErrKeyNotFound
		}
		err := bucket.Delete([]byte(key))
		return err
	})
}

// Exists checks if the key exists inside the kvstore.
func (b *Store) Exists(_ context.Context, key string, _ *kvstore.ReadOptions) (bool, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	db, err := b.getDBHandle()
	if err != nil {
		return false, err
	}
	defer b.releaseDBHandle()

	var val []byte

	err = db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(b.boltBucket)
		if bucket == nil {
			return kvstore.ErrKeyNotFound
		}

		val = bucket.Get([]byte(key))

		return nil
	})

	if len(val) == 0 {
		return false, err
	}
	return true, err
}

// List returns the range of keys starting with the passed in prefix.
func (b *Store) List(_ context.Context, keyPrefix string, _ *kvstore.ReadOptions) ([]*kvstore.KVPair, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	var kv []*kvstore.KVPair

	db, err := b.getDBHandle()
	if err != nil {
		return nil, err
	}
	defer b.releaseDBHandle()

	hasResult := false

	err = db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(b.boltBucket)
		if bucket == nil {
			return kvstore.ErrKeyNotFound
		}

		cursor := bucket.Cursor()
		prefix := []byte(keyPrefix)

		for key, v := cursor.Seek(prefix); key != nil && bytes.HasPrefix(key, prefix); key, v = cursor.Next() {
			hasResult = true

			dbIndex := binary.LittleEndian.Uint64(v[:metadataLen])
			v = v[metadataLen:]

			val := make([]byte, len(v))
			copy(val, v)

			if string(key) != keyPrefix {
				kv = append(kv, &kvstore.KVPair{
					Key:       string(key),
					Value:     val,
					LastIndex: dbIndex,
				})
			}
		}
		return nil
	})

	if !hasResult {
		return nil, kvstore.ErrKeyNotFound
	}

	return kv, err
}

// AtomicDelete deletes a value at "key" if the key has not been modified in the meantime,
// throws an error if this is the case.
func (b *Store) AtomicDelete(_ context.Context, key string, previous *kvstore.KVPair) (bool, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if previous == nil {
		return false, kvstore.ErrPreviousNotSpecified
	}

	db, err := b.getDBHandle()
	if err != nil {
		return false, err
	}
	defer b.releaseDBHandle()

	var val []byte

	err = db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(b.boltBucket)
		if bucket == nil {
			return kvstore.ErrKeyNotFound
		}

		val = bucket.Get([]byte(key))
		if val == nil {
			return kvstore.ErrKeyNotFound
		}
		dbIndex := binary.LittleEndian.Uint64(val[:metadataLen])
		if dbIndex != previous.LastIndex {
			return kvstore.ErrKeyModified
		}

		return bucket.Delete([]byte(key))
	})
	if err != nil {
		return false, err
	}
	return true, err
}

// AtomicPut puts a value at "key"
// if the key has not been modified since the last Put,
// throws an error if this is the case.
func (b *Store) AtomicPut(_ context.Context, key string, value []byte, previous *kvstore.KVPair, _ *kvstore.WriteOptions) (bool, *kvstore.KVPair, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	dbval := make([]byte, metadataLen)

	db, err := b.getDBHandle()
	if err != nil {
		return false, nil, err
	}
	defer b.releaseDBHandle()

	var dbIndex uint64

	errUpdate := db.Update(func(tx *bbolt.Tx) error {
		var err error
		bucket := tx.Bucket(b.boltBucket)
		if bucket == nil {
			if previous != nil {
				return kvstore.ErrKeyNotFound
			}
			bucket, err = tx.CreateBucket(b.boltBucket)
			if err != nil {
				return err
			}
		}

		// AtomicPut is equivalent to Put if previous is nil and the key doesn't exist in the DB.
		val := bucket.Get([]byte(key))
		if previous == nil && len(val) != 0 {
			return kvstore.ErrKeyExists
		}

		if previous != nil {
			if len(val) == 0 {
				return kvstore.ErrKeyNotFound
			}
			dbIndex = binary.LittleEndian.Uint64(val[:metadataLen])
			if dbIndex != previous.LastIndex {
				return kvstore.ErrKeyModified
			}
		}

		dbIndex = atomic.AddUint64(&b.dbIndex, 1)
		binary.LittleEndian.PutUint64(dbval, b.dbIndex)
		dbval = append(dbval, value...)

		return bucket.Put([]byte(key), dbval)
	})
	if errUpdate != nil {
		return false, nil, errUpdate
	}

	updated := &kvstore.KVPair{
		Key:       key,
		Value:     value,
		LastIndex: dbIndex,
	}

	return true, updated, nil
}

// Close the db connection to the BoltDB.
func (b *Store) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.PersistConnection {
		return b.client.Close()
	}

	b.reset()
	return nil
}

// DeleteTree deletes a range of keys with a given prefix.
func (b *Store) DeleteTree(_ context.Context, keyPrefix string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	db, err := b.getDBHandle()
	if err != nil {
		return err
	}
	defer b.releaseDBHandle()

	return db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(b.boltBucket)
		if bucket == nil {
			return kvstore.ErrKeyNotFound
		}

		cursor := bucket.Cursor()
		prefix := []byte(keyPrefix)

		for key, _ := cursor.Seek(prefix); bytes.HasPrefix(key, prefix); key, _ = cursor.Next() {
			_ = bucket.Delete(key)
		}
		return nil
	})
}

// NewLock has to implemented at the library level since it's not supported by BoltDB.
func (b *Store) NewLock(_ context.Context, _ string, _ *kvstore.LockOptions) (kvstore.Locker, error) {
	return nil, kvstore.ErrCallNotSupported
}

// Watch has to implemented at the library level since it's not supported by BoltDB.
func (b *Store) Watch(_ context.Context, _ string, _ *kvstore.ReadOptions) (<-chan *kvstore.KVPair, error) {
	return nil, kvstore.ErrCallNotSupported
}

// WatchTree has to implemented at the library level since it's not supported by BoltDB.
func (b *Store) WatchTree(_ context.Context, _ string, _ *kvstore.ReadOptions) (<-chan []*kvstore.KVPair, error) {
	return nil, kvstore.ErrCallNotSupported
}

func (b *Store) reset() {
	b.path = ""
	b.boltBucket = []byte{}
}

func (b *Store) getDBHandle() (*bbolt.DB, error) {
	if b.PersistConnection {
		return b.client, nil
	}

	boltOptions := &bbolt.Options{Timeout: b.timeout}
	db, err := bbolt.Open(b.path, filePerm, boltOptions)
	if err != nil {
		return nil, err
	}

	b.client = db
	return b.client, nil
}

func (b *Store) releaseDBHandle() {
	if !b.PersistConnection {
		_ = b.client.Close()
	}
}
