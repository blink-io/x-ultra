package tlru

import (
	"time"

	"github.com/ammario/tlru"

	"github.com/blink-io/x/cache"
)

const Name = "tlru"

const DefaultCost = 100_000

type icache[V any] struct {
	c   *tlru.Cache[string, V]
	ttl time.Duration
}

func New[V any](ttl time.Duration) cache.TTLCache[V] {
	c := tlru.New[string](tlru.ConstantCost[V], DefaultCost)
	return &icache[V]{c, ttl}
}

func (l *icache[V]) Set(key string, value V) {
	l.c.Set(key, value, l.ttl)
}

func (l *icache[V]) SetWithTTL(key string, value V, ttl time.Duration) {
	l.c.Set(key, value, ttl)
}

func (l *icache[V]) Get(key string) (V, bool) {
	var v V
	var exists bool
	v, _, exists = l.c.Get(key)
	return v, exists
}

func (l *icache[V]) Del(key string) {
	l.c.Delete(key)
}

func (l *icache[V]) Name() string {
	return Name
}
