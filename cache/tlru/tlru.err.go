package tlru

import (
	"time"

	"github.com/ammario/tlru"

	"github.com/blink-io/x/cache"
)

const (
	ErrName = Name + cache.ErrNameSuffix
)

var _ cache.ErrTTLCache[any] = (*ErrCache[any])(nil)

type ErrCache[V any] struct {
	cc  *tlru.Cache[string, V]
	ttl time.Duration
}

func NewErr[V any](ttl time.Duration) cache.TTLCache[V] {
	c := tlru.New[string](tlru.ConstantCost[V], DefaultCost)
	return &Cache[V]{c, ttl}
}

func (c *ErrCache[V]) Set(key string, value V) error {
	c.cc.Set(key, value, c.ttl)
	return nil
}

func (c *ErrCache[V]) SetWithTTL(key string, value V, ttl time.Duration) error {
	c.cc.Set(key, value, ttl)
	return nil
}

func (c *ErrCache[V]) Get(key string) (V, bool, error) {
	var v V
	var exists bool
	v, _, exists = c.cc.Get(key)
	return v, exists, nil
}

func (c *ErrCache[V]) Del(key string) error {
	c.cc.Delete(key)
	return nil
}

func (c *ErrCache[V]) Name() string {
	return ErrName
}
