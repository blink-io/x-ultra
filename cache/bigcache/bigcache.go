package bigcache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/blink-io/x/cache"
)

const Name = "bigcache"

var _ cache.Cache[any] = (*Cache[any])(nil)

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type Cache[V any] struct {
	cc  *bigcache.BigCache
	ttl time.Duration
	enc cache.Codec
}

func New[V any](ctx context.Context, ttl time.Duration) (*Cache[V], error) {
	c, err := bigcache.New(ctx, bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		return nil, err
	}
	return &Cache[V]{cc: c, ttl: ttl}, nil
}

func (c *Cache[V]) Set(key string, value V) {
	data, err := c.enc.Encode(value)
	if err == nil {
		c.cc.Set(key, data)
	}
}

func (c *Cache[V]) Get(key string) (V, bool) {
	data, err := c.cc.Get(key)
	var v V
	if err == nil {
		if vv, verr := c.enc.Decode(data); verr != nil {
			return v, false
		} else {
			v = vv.(V)
			return v, true
		}
	}
	return v, false
}

func (c *Cache[V]) Del(key string) {
	_ = c.cc.Delete(key)
}

func (c *Cache[V]) Name() string {
	return Name
}
