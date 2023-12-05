package tlru

import (
	"time"

	"github.com/ammario/tlru"

	"github.com/blink-io/x/cache"
)

const Name = "tlru"

const DefaultCost = 100_000

var _ cache.TTLCache[any] = (*Cache[any])(nil)

type Cache[V any] struct {
	cc  *tlru.Cache[string, V]
	ttl time.Duration
}

func New[V any](ttl time.Duration) cache.TTLCache[V] {
	c := tlru.New[string](tlru.ConstantCost[V], DefaultCost)
	return &Cache[V]{c, ttl}
}

func (c *Cache[V]) Set(key string, value V) {
	c.cc.Set(key, value, c.ttl)
}

func (c *Cache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	c.cc.Set(key, value, ttl)
}

func (c *Cache[V]) Get(key string) (V, bool) {
	var v V
	var exists bool
	v, _, exists = c.cc.Get(key)
	return v, exists
}

func (c *Cache[V]) Del(key string) {
	c.cc.Delete(key)
}

func (c *Cache[V]) Name() string {
	return Name
}
