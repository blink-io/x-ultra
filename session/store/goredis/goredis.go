package goredis

import (
	"context"
	"time"

	"github.com/blink-io/x/session/store"

	"github.com/redis/go-redis/v9"
)

// istore represents the session store.
type istore struct {
	client redis.UniversalClient
	prefix string
}

// New returns a new store instance. The client parameter should be a pointer
// to a go-redis connection.
func New(client redis.UniversalClient) interface {
	store.Store
} {
	return NewWithPrefix(client, store.DefaultPrefix)
}

// NewWithPrefix returns a new store instance. The pool parameter should be a pointer
// to a redigo connection pool. The prefix parameter controls the Redis key
// prefix, which can be used to avoid naming clashes if necessary.
func NewWithPrefix(client redis.UniversalClient, prefix string) interface {
	store.Store
} {
	return newRawWithPrefix(client, prefix)
}

func newRaw(client redis.UniversalClient) *istore {
	return newRawWithPrefix(client, store.DefaultPrefix)
}

func newRawWithPrefix(client redis.UniversalClient, prefix string) *istore {
	return &istore{
		client: client,
		prefix: prefix,
	}
}

// Find returns the data for a given session token from the store instance.
// If the session token is not found or is expired, the returned exists flag
// will be set to false.
func (s *istore) Find(ctx context.Context, token string) (b []byte, exists bool, err error) {
	b, err = s.client.Get(ctx, s.prefix+token).Bytes()
	if err == redis.Nil {
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
	err := s.client.Set(ctx, s.prefix+token, string(b), expiry.Sub(time.Now())).Err()
	return err
}

// Delete removes a session token and corresponding data from the store
// instance.
func (s *istore) Delete(ctx context.Context, token string) error {
	return s.client.Del(ctx, s.prefix+token).Err()
}

// All returns a map containing the token and data for all active (i.e.
// not expired) sessions in the store instance.
func (s *istore) All(ctx context.Context) (map[string][]byte, error) {
	var cursor uint64
	sessions := make(map[string][]byte)

	for {
		var keys []string
		var err error
		keys, cursor, err = s.client.Scan(ctx, cursor, s.prefix+"*", 0).Result()
		if err != nil {
			if err == redis.Nil {
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
