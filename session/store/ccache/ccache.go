package ccache

import (
	"context"
	"time"

	"github.com/blink-io/x/session/store"
	"github.com/karlseguin/ccache/v3"
)

const Name = "ccache"

// New returns a new store.Store instance.
// The client parameter should be a pointer to an etcd client instance.
func New() store.Store {
	return newRaw()
}

func newRaw() *istore {
	cfg := ccache.Configure[[]byte]()
	cc := ccache.New(cfg)
	return &istore{cc: cc, tt: make(map[string]*struct{})}
}

var _ store.Store = (*istore)(nil)

type istore struct {
	tt map[string]*struct{}
	cc *ccache.Cache[[]byte]
}

func (s *istore) Name() string {
	return Name
}

func (s *istore) Delete(ctx context.Context, token string) (err error) {
	s.cc.Delete(token)
	delete(s.tt, token)
	return nil
}

func (s *istore) Find(ctx context.Context, token string) ([]byte, bool, error) {
	if item := s.cc.Get(token); item != nil && !item.Expired() {
		return item.Value(), true, nil
	}
	return nil, false, nil
}

func (s *istore) Commit(ctx context.Context, token string, data []byte, expiry time.Time) (err error) {
	ttl := time.Until(expiry)
	s.cc.Set(token, data, ttl)
	s.tt[token] = store.NilStruct
	return nil
}

func (s *istore) All(ctx context.Context) (map[string][]byte, error) {
	sessions := make(map[string][]byte)
	for token := range s.tt {
		if item := s.cc.Get(token); item != nil && !item.Expired() {
			sessions[token] = item.Value()
		}
	}
	return sessions, nil
}
