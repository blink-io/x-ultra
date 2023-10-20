package ristretto

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"
	"github.com/outcaste-io/ristretto"
)

const Name = "ristretto"

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type icache[V any] struct {
	c   *ristretto.Cache
	ttl time.Duration
}

func New[V any](ctx context.Context, ttl time.Duration) (cache.Cache[V], error) {
	c, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, err
	}
	return &icache[V]{c, ttl}, nil
}

func (l *icache[V]) Set(key string, value V) {
	l.c.SetWithTTL(key, value, 1, l.ttl)
}

func (l *icache[V]) Get(key string) (V, bool) {
	var v V
	i, ok := l.c.Get(key)
	if ok {
		v = i.(V)
	}
	return v, ok
}

func (l *icache[V]) Del(key string) {
	l.c.Del(key)
}

func (l *icache[V]) Name() string {
	return Name
}
