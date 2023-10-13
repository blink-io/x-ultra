package lru

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

const Name = "icache"

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type icache[K comparable, V any] struct {
	c   *expirable.LRU[K, V]
	ttl time.Duration
}

func New[K comparable, V any](ctx context.Context, ttl time.Duration) (cache.Cache[K, V], error) {
	c := expirable.NewLRU[K, V](1000, nil, ttl)
	return &icache[K, V]{c, ttl}, nil
}

func (l *icache[K, V]) Set(key K, data V) {
	l.c.Add(key, data)
}

func (l *icache[K, V]) Get(key K) (V, bool) {
	return l.c.Get(key)
}

func (l *icache[K, V]) Del(key K) {
	l.c.Remove(key)
}

func (l *icache[K, V]) Name() string {
	return Name
}
