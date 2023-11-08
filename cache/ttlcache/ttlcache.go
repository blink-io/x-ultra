package ttlcache

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"

	"github.com/jellydator/ttlcache/v3"
)

const Name = "ttlcache"

type icache[V any] struct {
	c   *ttlcache.Cache[string, V]
	ttl time.Duration
}

func New[V any](ctx context.Context, ttl time.Duration) (cache.Cache[V], error) {
	c := ttlcache.New(ttlcache.WithTTL[string, V](ttl))
	return &icache[V]{c, ttl}, nil
}

func (l *icache[V]) Set(key string, value V) {
	l.c.Set(key, value, l.ttl)
}

func (l *icache[V]) Get(key string) (V, bool) {
	i := l.c.Get(key)
	var v V
	if i != nil {
		return i.Value(), i.IsExpired()

	}
	return v, false
}

func (l *icache[V]) Del(key string) {
	l.c.Delete(key)
}

func (l *icache[V]) Name() string {
	return Name
}
