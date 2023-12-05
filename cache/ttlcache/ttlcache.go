package ttlcache

import (
	"time"

	"github.com/blink-io/x/cache"

	"github.com/jellydator/ttlcache/v3"
)

const Name = "ttlcache"

var _ cache.TTLCache[any] = (*Cache[any])(nil)

type Cache[V any] struct {
	cc  *ttlcache.Cache[string, V]
	ttl time.Duration
}

func New[V any](ttl time.Duration) *Cache[V] {
	c := ttlcache.New(ttlcache.WithTTL[string, V](ttl))
	return &Cache[V]{c, ttl}
}

func (c *Cache[V]) Set(key string, value V) {
	c.cc.Set(key, value, c.ttl)
}

func (c *Cache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	c.cc.Set(key, value, ttl)
}

func (c *Cache[V]) Get(key string) (V, bool) {
	i := c.cc.Get(key)
	var v V
	if i != nil {
		return i.Value(), !i.IsExpired()
	}
	return v, false
}

func (c *Cache[V]) Del(key string) {
	c.cc.Delete(key)
}

func (c *Cache[V]) Name() string {
	return Name
}
