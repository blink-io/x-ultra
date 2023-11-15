package ristretto

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/blink-io/x/session/store"
	"github.com/outcaste-io/ristretto"
)

const Name = "ristretto"

// New returns a new store.Store instance.
// The client parameter should be a pointer to an etcd client instance.
func New() store.Store {
	return newRaw()
}

func newRaw() *istore {
	cc, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 100_000,
		MaxCost:     100_000,
		BufferItems: 100_000,
	})
	s := &istore{cc: cc, tt: make(map[string]*struct{})}
	return s
}

var _ store.Store = (*istore)(nil)

type istore struct {
	tt map[string]*struct{}
	cc *ristretto.Cache
}

func (s *istore) Name() string {
	return Name
}

func (s *istore) Delete(ctx context.Context, token string) (err error) {
	s.cc.Del(token)
	delete(s.tt, token)
	return nil
}

func (s *istore) Find(ctx context.Context, token string) ([]byte, bool, error) {
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

func (s *istore) Commit(ctx context.Context, token string, data []byte, expiry time.Time) (err error) {
	ttl := time.Until(expiry)
	ok := s.cc.SetWithTTL(token, data, 1, ttl)
	if !ok {
		return fmt.Errorf("unable to store: %s", token)
	}
	s.tt[token] = store.NilStruct
	return nil
}

func (s *istore) All(ctx context.Context) (map[string][]byte, error) {
	sessions := make(map[string][]byte)
	for token := range s.tt {
		if v, ok := s.cc.Get(token); ok {
			if data, vok := v.([]byte); vok {
				sessions[token] = data
			}
		}
	}
	return sessions, nil
}