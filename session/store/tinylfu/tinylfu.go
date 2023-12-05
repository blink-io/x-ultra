package tinylfu

import (
	"context"
	"errors"
	"time"

	"github.com/blink-io/x/session/store"
	"github.com/vmihailenco/go-tinylfu"
)

const Name = "tinylfu"

var _ store.Store = (*Store)(nil)

// New returns a new store.Store instance.
// The client parameter should be a pointer to an etcd client instance.
func New() *Store {
	cc := tinylfu.New(100_000_000, 100_000_000)
	return &Store{cc: cc, tm: store.NewTokenMap()}
}

var _ store.Store = (*Store)(nil)

type Store struct {
	tm store.TokenMap
	cc *tinylfu.T
}

func (s *Store) Name() string {
	return Name
}

func (s *Store) Delete(ctx context.Context, token string) (err error) {
	s.cc.Del(token)
	delete(s.tm, token)
	return nil
}

func (s *Store) Find(ctx context.Context, token string) ([]byte, bool, error) {
	if v, ok := s.cc.Get(token); ok {
		if data, vok := v.([]byte); vok {
			return data, true, nil
		} else {
			return nil, false, errors.New("invalid value type")
		}
	} else {
		return nil, false, nil
	}
}

func (s *Store) Commit(ctx context.Context, token string, data []byte, expiry time.Time) (err error) {
	s.cc.Set(&tinylfu.Item{
		Key:      token,
		Value:    data,
		ExpireAt: expiry,
	})
	s.tm[token] = store.NilStruct
	return nil
}

func (s *Store) All(ctx context.Context) (map[string][]byte, error) {
	sessions := make(map[string][]byte)
	for token := range s.tm {
		if v, ok := s.cc.Get(token); ok {
			if data, vok := v.([]byte); vok {
				sessions[token] = data
			}
		}
	}
	return sessions, nil
}
