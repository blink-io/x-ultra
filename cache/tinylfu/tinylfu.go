package tinylfu

import (
	"time"

	"github.com/blink-io/x/cache"
	"github.com/vmihailenco/go-tinylfu"
)

const Name = "tinylfu"

var _ cache.TTLCache[any] = (*Cache[any])(nil)

type Cache[V any] struct {
	cc  *tinylfu.T
	ttl time.Duration
}

func New[V any](ttl time.Duration) (*Cache[V], error) {
	c := tinylfu.New(100_000_000, 100_000_000)
	return &Cache[V]{c, ttl}, nil
}

func (c *Cache[V]) Set(key string, value V) {
	c.SetWithTTL(key, value, c.ttl)
}

func (c *Cache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	expiresAt := time.Now().Add(ttl)
	i := &tinylfu.Item{
		Key:      key,
		Value:    value,
		ExpireAt: expiresAt,
	}
	c.cc.Set(i)
}

func (c *Cache[V]) Get(key string) (V, bool) {
	rv, ok := c.cc.Get(key)
	return rv.(V), ok
}

func (c *Cache[V]) Del(key string) {
	c.cc.Del(key)
}

func (c *Cache[V]) Name() string {
	return Name
}
