package rueidis

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"
	"github.com/blink-io/x/cast"
	"github.com/gogo/protobuf/codec"
	"github.com/redis/rueidis"
)

const Name = "goredis"

var _ cache.TTLCache[any] = (*Cache[any])(nil)

type Cache[V any] struct {
	cc  rueidis.Client
	ttl time.Duration
	ctx context.Context
	enc codec.Codec
}

func New[V any](cc rueidis.Client, ttl time.Duration) (*Cache[V], error) {
	return &Cache[V]{
		cc:  cc,
		ttl: ttl,
		ctx: context.Background(),
	}, nil
}

func (c *Cache[V]) Set(key string, value V) {
	c.setWithTTL(key, value, c.ttl)
}

func (c *Cache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	c.setWithTTL(key, value, ttl)
}

func (c *Cache[V]) setWithTTL(key string, value V, ttl time.Duration) {
	cmd := c.cc.B().Set().Key(key).Value(cast.ToString(value)).Ex(ttl).Build()
	c.cc.Do(c.ctx, cmd)
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
	cmd := c.cc.B().Del().Key(key).Build()
	c.cc.Do(c.ctx, cmd)
}

func (c *Cache[V]) Name() string {
	return Name
}
