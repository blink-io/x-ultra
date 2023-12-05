package goredis

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"
	"github.com/redis/go-redis/v9"
)

const Name = "goredis"

var _ cache.TTLCache[any] = (*Cache[any])(nil)

type Cache[V any] struct {
	cc  redis.UniversalClient
	ttl time.Duration
	ctx context.Context
}

func New[V any](cc redis.UniversalClient, ttl time.Duration) (*Cache[V], error) {
	return &Cache[V]{
		cc:  cc,
		ttl: ttl,
		ctx: context.Background(),
	}, nil
}

func (c *Cache[V]) Set(key string, value V) {
	c.cc.Set(c.ctx, key, value, c.ttl)
}

func (c *Cache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	c.cc.Set(c.ctx, key, value, ttl)
}

func (c *Cache[V]) Get(key string) (V, bool) {
	//i := c.cc.Get(c.ctx, key)
	var v V
	//if i != nil {
	//	return i.Value(), i.Expired()
	//}
	return v, false
}

func (c *Cache[V]) Del(key string) {
	c.cc.Del(c.ctx, key)
}

func (c *Cache[V]) Name() string {
	return Name
}
