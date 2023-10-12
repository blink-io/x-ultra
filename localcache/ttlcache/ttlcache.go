package ttlcache

import (
	"context"
	"time"

	"github.com/blink-io/x/localcache"

	"github.com/jellydator/ttlcache/v3"
)

const Name = "ttlcache"

type cache[K comparable, V any] struct {
	c   *ttlcache.Cache[K, V]
	ttl time.Duration
}

func New[K comparable, V any](ctx context.Context, ttl time.Duration) (localcache.Cache[K, V], error) {
	c := ttlcache.New(ttlcache.WithTTL[K, V](ttl))
	return &cache[K, V]{c, ttl}, nil
}

func (l *cache[K, V]) Set(key K, data V) {
	l.c.Set(key, data, 0)
}

func (l *cache[K, V]) Get(key K) (V, bool) {
	i := l.c.Get(key)
	var v V
	if i != nil {
		return i.Value(), !i.IsExpired()

	}
	return v, false
}

func (l *cache[K, V]) Del(key K) {
	l.c.Delete(key)
}

func (l *cache[K, V]) Name() string {
	return Name
}
