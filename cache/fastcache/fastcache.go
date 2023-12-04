package fastcache

import (
	"context"
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/blink-io/x/cache"
)

const Name = "fastcache"

const MaxCost = 100_000

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

type icache[V any] struct {
	c   *fastcache.Cache
	ttl time.Duration
	enc cache.Codec
}

func New[V any](ctx context.Context, ttl time.Duration) (cache.Cache[V], error) {
	c := fastcache.New(MaxCost)
	return &icache[V]{c: c, ttl: ttl}, nil
}

func (l *icache[V]) Set(key string, value V) {
	data, err := l.enc.Encode(value)
	if err == nil {
		l.c.Set(keyToBytes(key), data)
	}
}

func (l *icache[V]) Get(key string) (V, bool) {
	data := l.c.Get(nil, keyToBytes(key))
	var v V
	if vv, verr := l.enc.Decode(data); verr != nil {
		return v, false
	} else {
		v = vv.(V)
		return v, true
	}
}

func (l *icache[V]) Del(key string) {
	l.c.Del(keyToBytes(key))
}

func (l *icache[V]) Name() string {
	return Name
}

func keyToBytes(key string) []byte {
	return []byte(key)
}
