package lru

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

const Name = "Cache"

var _ cache.Cache[any] = (*Cache[any])(nil)

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type Cache[V any] struct {
	cc  *expirable.LRU[string, V]
	ttl time.Duration
}

func New[V any](ctx context.Context, ttl time.Duration) (*Cache[V], error) {
	c := expirable.NewLRU[string, V](1000, nil, ttl)
	return &Cache[V]{c, ttl}, nil
}

func (c *Cache[V]) Set(key string, value V) {
	c.cc.Add(key, value)
}

func (c *Cache[V]) Get(key string) (V, bool) {
	return c.cc.Get(key)
}

func (c *Cache[V]) Del(key string) {
	c.cc.Remove(key)
}

func (c *Cache[V]) Name() string {
	return Name
}
