package ttlcache

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"

	"github.com/jellydator/ttlcache/v3"
)

const Name = "ttlcache"

type icache[K comparable, V any] struct {
	c   *ttlcache.Cache[K, V]
	ttl time.Duration
}

func New[K comparable, V any](ctx context.Context, ttl time.Duration) (cache.Cache[K, V], error) {
	c := ttlcache.New(ttlcache.WithTTL[K, V](ttl))
	return &icache[K, V]{c, ttl}, nil
}

func (l *icache[K, V]) Set(key K, data V) {
	l.c.Set(key, data, 0)
}

func (l *icache[K, V]) Get(key K) (V, bool) {
	i := l.c.Get(key)
	var v V
	if i != nil {
		return i.Value(), !i.IsExpired()

	}
	return v, false
}

func (l *icache[K, V]) Del(key K) {
	l.c.Delete(key)
}

func (l *icache[K, V]) Name() string {
	return Name
}
