package rueidis

import (
	"context"
	"time"

	"github.com/blink-io/x/session/store"

	"github.com/redis/rueidis"
)

// istore represents the session store.
type istore struct {
	client rueidis.Client
	prefix string
}

// New returns a new store instance. The client parameter should be a pointer
// to a go-redis connection.
func New(client rueidis.Client) interface {
	store.Store
} {
	return NewWithPrefix(client, store.DefaultPrefix)
}

// NewWithPrefix returns a new store instance. The pool parameter should be a pointer
// to a redigo connection pool. The prefix parameter controls the Redis key
// prefix, which can be used to avoid naming clashes if necessary.
func NewWithPrefix(client rueidis.Client, prefix string) interface {
	store.Store
} {
	return newRawWithPrefix(client, prefix)
}

func newRaw(client rueidis.Client) *istore {
	return newRawWithPrefix(client, store.DefaultPrefix)
}

func newRawWithPrefix(client rueidis.Client, prefix string) *istore {
	return &istore{
		client: client,
		prefix: prefix,
	}
}

// Find returns the data for a given session token from the store instance.
// If the session token is not found or is expired, the returned exists flag
// will be set to false.
func (s *istore) Find(ctx context.Context, token string) (b []byte, exists bool, err error) {
	getCmd := s.client.B().Get().Key(s.prefix + token).Build()
	b, err = s.client.Do(ctx, getCmd).AsBytes()
	if err == rueidis.Nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return b, true, nil
}

// Commit adds a session token and data to the store instance with the
// given expiry time. If the session token already exists then the data and
// expiry time are updated.
func (s *istore) Commit(ctx context.Context, token string, b []byte, expiry time.Time) error {
	setCmd := s.client.B().Set().Key(s.prefix + token).Value(string(b)).Ex(expiry.Sub(time.Now())).Build()
	err := s.client.Do(ctx, setCmd).Error()
	return err
}

// Delete removes a session token and corresponding data from the store
// instance.
func (s *istore) Delete(ctx context.Context, token string) error {
	delCmd := s.client.B().Del().Key(s.prefix + token).Build()
	return s.client.Do(ctx, delCmd).Error()
}

// All returns a map containing the token and data for all active (i.e.
// not expired) sessions in the store instance.
func (s *istore) All(ctx context.Context) (map[string][]byte, error) {
	var cursor uint64
	sessions := make(map[string][]byte)

	for {
		scanCmd := s.client.B().Scan().Cursor(cursor).Match(s.prefix + "*").Build()
		v, err := s.client.Do(ctx, scanCmd).AsScanEntry()
		cursor = v.Cursor
		keys := v.Elements
		if err != nil {
			if err == rueidis.Nil {
				return nil, nil
			} else {
				return nil, err
			}
		}
		for _, key := range keys {
			token := key[len(s.prefix):]
			data, exists, err := s.Find(ctx, token)
			if err != nil {
				return nil, err
			}
			if exists {
				sessions[token] = data
			}
		}
		if cursor == 0 {
			break
		}
	}
	return sessions, nil
}
