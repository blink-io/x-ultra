package bigcache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/blink-io/x/cache"
)

const Name = "bigcache"

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type icache[V any] struct {
	c   *bigcache.BigCache
	ttl time.Duration
	enc cache.Codec
}

func New[V any](ctx context.Context, ttl time.Duration) (cache.Cache[V], error) {
	c, err := bigcache.New(ctx, bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		return nil, err
	}
	return &icache[V]{c: c, ttl: ttl}, nil
}

func (l *icache[V]) Set(key string, value V) {
	data, err := l.enc.Encode(value)
	if err == nil {
		l.c.Set(key, data)
	}
}

func (l *icache[V]) Get(key string) (V, bool) {
	data, err := l.c.Get(key)
	var v V
	if err == nil {
		if vv, verr := l.enc.Decode(data); verr != nil {
			return v, false
		} else {
			v = vv.(V)
			return v, true
		}
	}
	return v, false
}

func (l *icache[V]) Del(key string) {
	l.c.Delete(key)
}

func (l *icache[V]) Name() string {
	return Name
}
