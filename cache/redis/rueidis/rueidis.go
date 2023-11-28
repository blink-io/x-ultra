package rueidis

import (
	"context"
	"time"

	"github.com/blink-io/x/cache"
	"github.com/blink-io/x/cast"
	"github.com/redis/rueidis"
)

const Name = "goredis"

type icache[V any] struct {
	cc  rueidis.Client
	ttl time.Duration
	ctx context.Context
}

func New[V any](cc rueidis.Client, ttl time.Duration) (cache.Cache[V], error) {
	return &icache[V]{
		cc:  cc,
		ttl: ttl,
		ctx: context.Background(),
	}, nil
}

func (l *icache[V]) Set(key string, value V) {
	cmd := l.cc.B().Set().Key(key).Value(cast.ToString(value)).Ex(l.ttl).Build()
	l.cc.Do(l.ctx, cmd)
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
	cmd := l.cc.B().Del().Key(key).Build()
	l.cc.Do(l.ctx, cmd)
}

func (l *icache[V]) Name() string {
	return Name
}
