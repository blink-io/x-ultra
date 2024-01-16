package ccache

import (
	"time"

	"github.com/blink-io/x/cache"
)

const ErrName = Name + cache.ErrNameSuffix

var _ cache.ErrTTLCache[any] = (*ErrCache[any])(nil)

type ErrCache[V any] struct {
	cc  cache.TTLCache[V]
	ttl time.Duration
}

func NewErr[V any](ttl time.Duration) (*ErrCache[V], error) {
	cc, err := New[V](ttl)
	if err != nil {
		return nil, err
	}
	return &ErrCache[V]{cc: cc, ttl: ttl}, nil
}

func (l *ErrCache[V]) Set(key string, data V) error {
	l.cc.Set(key, data)
	return nil
}

func (l *ErrCache[V]) SetWithTTL(key string, data V, ttl time.Duration) error {
	l.cc.SetWithTTL(key, data, ttl)
	return nil
}

func (l *ErrCache[V]) Get(key string) (V, bool, error) {
	v, exists := l.cc.Get(key)
	return v, exists, nil
}

func (l *ErrCache[V]) Del(key string) error {
	l.cc.Del(key)
	return nil
}

func (l *ErrCache[V]) Name() string {
	return ErrName
}
