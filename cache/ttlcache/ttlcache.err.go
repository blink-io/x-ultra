package ttlcache

import (
	"time"

	"github.com/blink-io/x/cache"
)

const ErrName = Name + cache.ErrNameSuffix

var _ cache.ErrTTLCache[any] = (*ErrCache[any])(nil)

type ErrCache[V any] struct {
	cc  cache.TTLCache[V]
	ttl time.Duration
}

func NewErr[V any](ttl time.Duration) *ErrCache[V] {
	cc := New[V](ttl)
	return &ErrCache[V]{cc: cc, ttl: ttl}
}

func (c *ErrCache[V]) Set(key string, value V) error {
	c.cc.Set(key, value)
	return nil
}

func (c *ErrCache[V]) SetWithTTL(key string, value V, ttl time.Duration) error {
	c.cc.SetWithTTL(key, value, ttl)
	return nil
}

func (c *ErrCache[V]) Get(key string) (V, bool, error) {
	v, exists := c.cc.Get(key)
	return v, exists, nil
}

func (c *ErrCache[V]) Del(key string) error {
	c.cc.Del(key)
	return nil
}

func (c *ErrCache[V]) Name() string {
	return ErrName
}
