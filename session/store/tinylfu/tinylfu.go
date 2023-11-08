package tinylfu

import (
	"context"
	"errors"
	"time"

	"github.com/blink-io/x/session/store"
	"github.com/vmihailenco/go-tinylfu"
)

// New returns a new store.Store instance.
// The client parameter should be a pointer to an etcd client instance.
func New() store.Store {
	return newRaw()
}

func newRaw() *istore {
	cc := tinylfu.New(100_000_000, 100_000_000)
	return &istore{cc: cc, tt: make(map[string]*struct{})}
}

var _ store.Store = (*istore)(nil)

type istore struct {
	tt map[string]*struct{}
	cc *tinylfu.T
}

func (i *istore) Delete(ctx context.Context, token string) (err error) {
	i.cc.Del(token)
	delete(i.tt, token)
	return nil
}

func (i *istore) Find(ctx context.Context, token string) ([]byte, bool, error) {
	if v, ok := i.cc.Get(token); ok {
		if data, vok := v.([]byte); vok {
			return data, true, nil
		} else {
			return nil, false, errors.New("invalid value type")
		}
	} else {
		return nil, false, nil
	}
}

func (i *istore) Commit(ctx context.Context, token string, data []byte, expiry time.Time) (err error) {
	i.cc.Set(&tinylfu.Item{
		Key:      token,
		Value:    data,
		ExpireAt: expiry,
	})
	i.tt[token] = store.NilStruct
	return nil
}

func (i *istore) All(ctx context.Context) (map[string][]byte, error) {
	sessions := make(map[string][]byte)
	for token := range i.tt {
		if v, ok := i.cc.Get(token); ok {
			if data, vok := v.([]byte); vok {
				sessions[token] = data
			}
		}
	}
	return sessions, nil
}
