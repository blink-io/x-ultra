package tinylfu

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"
	"github.com/vmihailenco/go-tinylfu"
)

const Name = "tinylfu"

type icache[V any] struct {
	c   *tinylfu.T
	ttl time.Duration
}

func New[V any](ctx context.Context, ttl time.Duration) (cache.Cache[V], error) {
	c := tinylfu.New(100_000_000, 100_000_000)
	return &icache[V]{c, ttl}, nil
}

func (l *icache[V]) Set(key string, value V) {
	l.SetWithTTL(key, value, l.ttl)
}

func (l *icache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	expiresAt := time.Now().Add(ttl)
	i := &tinylfu.Item{
		Key:      key,
		Value:    value,
		ExpireAt: expiresAt,
	}
	l.c.Set(i)
}

func (l *icache[V]) Get(key string) (V, bool) {
	rv, ok := l.c.Get(key)
	return rv.(V), ok
}

func (l *icache[V]) Del(key string) {
	l.c.Del(key)
}

func (l *icache[V]) Name() string {
	return Name
}
