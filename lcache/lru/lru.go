package lru

import (
	"context"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

const Name = "lru"

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type cache[K comparable, V any] struct {
	c   *expirable.LRU[K, V]
	ttl time.Duration
}

func New[K comparable, V any](ctx context.Context, ttl time.Duration) (lcache.Cache[K, V], error) {
	c := expirable.NewLRU[K, V](1000, nil, ttl)
	return &cache[K, V]{c, ttl}, nil
}

func (l *cache[K, V]) Set(key K, data V) {
	l.c.Add(key, data)
}

func (l *cache[K, V]) Get(key K) (V, bool) {
	return l.c.Get(key)
}

func (l *cache[K, V]) Del(key K) {
	l.c.Remove(key)
}

func (l *cache[K, V]) Name() string {
	return Name
}
