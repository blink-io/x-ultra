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

type icache[V any] struct {
	c   *expirable.LRU[string, V]
	ttl time.Duration
}

func New[V any](ctx context.Context, ttl time.Duration) (cache.Cache[V], error) {
	c := expirable.NewLRU[string, V](1000, nil, ttl)
	return &icache[V]{c, ttl}, nil
}

func (l *icache[V]) Set(key string, value V) {
	l.c.Add(key, value)
}

func (l *icache[V]) Get(key string) (V, bool) {
	return l.c.Get(key)
}

func (l *icache[V]) Del(key string) {
	l.c.Remove(key)
}

func (l *icache[V]) Name() string {
	return Name
}