package fastcache

import (
	"context"
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/blink-io/x/cache"
)

const Name = "fastcache"

const MaxCost = 100_000

var _ cache.Cache[any] = (*Cache[any])(nil)

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type Cache[V any] struct {
	cc  *fastcache.Cache
	ttl time.Duration
	enc cache.Codec
}

func New[V any](ctx context.Context, ttl time.Duration) (*Cache[V], error) {
	c := fastcache.New(MaxCost)
	return &Cache[V]{cc: c, ttl: ttl}, nil
}

func (c *Cache[V]) Set(key string, value V) {
	data, err := c.enc.Encode(value)
	if err == nil {
		c.cc.Set(keyToBytes(key), data)
	}
}

func (c *Cache[V]) Get(key string) (V, bool) {
	data := c.cc.Get(nil, keyToBytes(key))
	var v V
	if vv, verr := c.enc.Decode(data); verr != nil {
		return v, false
	} else {
		v = vv.(V)
		return v, true
	}
}

func (c *Cache[V]) Del(key string) {
	c.cc.Del(keyToBytes(key))
}

func (c *Cache[V]) Name() string {
	return Name
}

func keyToBytes(key string) []byte {
	return []byte(key)
}
