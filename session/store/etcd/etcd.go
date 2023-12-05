package etcd

import (
	"context"
	"time"

	"github.com/blink-io/x/session/store"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const Name = "etcd"

var _ store.Store = (*Store)(nil)

// Store represents the session store.
type Store struct {
	client *clientv3.Client
	prefix string
}

// New returns a new store.Store instance.
// The client parameter should be a pointer to an etcd client instance.
func New(client *clientv3.Client) *Store {
	return NewWithPrefix(client, store.DefaultPrefix)
}

// NewWithPrefix returns a new store.Store instance. The client parameter should be a pointer
// to a etcd client instance. The prefix parameter controls the etcd key
// prefix, which can be used to avoid naming clashes if necessary.
func NewWithPrefix(client *clientv3.Client, prefix string) *Store {
	return &Store{
		client: client,
		prefix: prefix,
	}
}

func (e *Store) Name() string {
	return Name
}

// Find returns the data for a given session token from the store.Store instance.
// If the session token is not found or is expired, the returned exists flag will
// be set to false.
func (e *Store) Find(ctx context.Context, token string) (b []byte, exists bool, err error) {
	res, err := e.client.Get(ctx, e.prefix+token)
	if err != nil {
		return nil, false, err
	}

	if len(res.Kvs) == 0 {
		return nil, false, nil
	}

	return res.Kvs[0].Value, true, nil
}

// Commit adds a session token and data to the store.Store instance with the
// given expiry time. If the session token already exists then the data and expiry
// time are updated.
func (e *Store) Commit(ctx context.Context, token string, b []byte, expiry time.Time) error {
	lease, _ := e.client.Grant(ctx, int64(time.Until(expiry).Seconds()))
	_, err := e.client.Put(ctx, e.prefix+token, string(b), clientv3.WithLease(lease.ID))
	return err
}

// Delete removes a session token and corresponding data from the store.Store instance.
func (e *Store) Delete(ctx context.Context, token string) error {
	_, err := e.client.Delete(ctx, e.prefix+token)
	return err
}

// All returns a map containing the token and data for all active (i.e.
// not expired) sessions in the Store instance.
func (e *Store) All(ctx context.Context) (map[string][]byte, error) {
	sessions := make(map[string][]byte)

	opts := []clientv3.OpOption{
		clientv3.WithPrefix(),
	}

	res, err := e.client.Get(ctx, e.prefix, opts...)
	if err != nil {
		return nil, err
	}

	if len(res.Kvs) == 0 {
		return sessions, nil
	}

	for _, kv := range res.Kvs {
		sessions[string(kv.Key)[len(e.prefix):]] = kv.Value
	}

	return sessions, nil
}
