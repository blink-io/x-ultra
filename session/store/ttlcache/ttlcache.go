package ttlcache

import (
	"context"
	"time"

	"github.com/blink-io/x/session/store"
	"github.com/jellydator/ttlcache/v3"
)

const Name = "ttlcache"

var _ store.Store = (*Store)(nil)

// New returns a new store.Store instance.
// The client parameter should be a pointer to an etcd client instance.
func New() *Store {
	cc := ttlcache.New[string, []byte]()
	return &Store{cc: cc, tt: store.NewTokenMap()}
}

var _ store.Store = (*Store)(nil)

type Store struct {
	tt store.TokenMap
	cc *ttlcache.Cache[string, []byte]
}

func (s *Store) Name() string {
	return Name
}

func (s *Store) Delete(ctx context.Context, token string) (err error) {
	s.cc.Delete(token)
	delete(s.tt, token)
	return nil
}

func (s *Store) Find(ctx context.Context, token string) ([]byte, bool, error) {
	if item := s.cc.Get(token); item != nil && !item.IsExpired() {
		return item.Value(), true, nil
	}
	return nil, false, nil
}

func (s *Store) Commit(ctx context.Context, token string, data []byte, expiry time.Time) (err error) {
	ttl := time.Until(expiry)
	s.cc.Set(token, data, ttl)
	s.tt[token] = store.NilStruct
	return nil
}

func (s *Store) All(ctx context.Context) (map[string][]byte, error) {
	sessions := make(map[string][]byte)
	for token := range s.tt {
		if item := s.cc.Get(token); item != nil && !item.IsExpired() {
			sessions[token] = item.Value()
		}
	}
	return sessions, nil
}
