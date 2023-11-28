package goredis

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"
	"github.com/redis/go-redis/v9"
)

const Name = "goredis"

type icache[V any] struct {
	cc  redis.UniversalClient
	ttl time.Duration
	ctx context.Context
}

func New[V any](cc redis.UniversalClient, ttl time.Duration) (cache.Cache[V], error) {
	return &icache[V]{
		cc:  cc,
		ttl: ttl,
		ctx: context.Background(),
	}, nil
}

func (l *icache[V]) Set(key string, data V) {
	l.cc.Set(l.ctx, key, data, l.ttl)
}

func (l *icache[V]) Get(key string) (V, bool) {
	//i := l.cc.Get(l.ctx, key)
	var v V
	//if i != nil {
	//	return i.Value(), i.Expired()
	//}
	return v, false
}

func (l *icache[V]) Del(key string) {
	l.cc.Del(l.ctx, key)
}

func (l *icache[V]) Name() string {
	return Name
}
