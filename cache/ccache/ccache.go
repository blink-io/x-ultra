package ccache

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"

	"github.com/karlseguin/ccache/v3"
)

const Name = "ccache"

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type icache[V any] struct {
	c   *ccache.Cache[V]
	ttl time.Duration
}

func New[V any](ctx context.Context, ttl time.Duration) (cache.Cache[V], error) {
	cfg := ccache.Configure[V]()
	c := ccache.New(cfg)
	return &icache[V]{c, ttl}, nil
}

func (l *icache[V]) Set(key string, data V) {
	l.c.Set(key, data, l.ttl)
}

func (l *icache[V]) Get(key string) (V, bool) {
	i := l.c.Get(key)
	var v V
	if i != nil {
		return i.Value(), i.Expired()
	}
	return v, false
}

func (l *icache[V]) Del(key string) {
	l.c.Delete(key)
}

func (l *icache[V]) Name() string {
	return Name
}
