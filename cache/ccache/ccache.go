package ccache

import (
	"time"

	"github.com/blink-io/x/cache"

	"github.com/karlseguin/ccache/v3"
)

const Name = "ccache"

func init() {
	//local.SetProviderFn(ProviderLRU, NewLRULocal)
}

var _ cache.TTLCache[any] = (*Cache[any])(nil)

type Cache[V any] struct {
	c   *ccache.Cache[V]
	ttl time.Duration
}

func New[V any](ttl time.Duration) (*Cache[V], error) {
	cfg := ccache.Configure[V]()
	c := ccache.New(cfg)
	return &Cache[V]{c, ttl}, nil
}

func (l *Cache[V]) Set(key string, data V) {
	l.c.Set(key, data, l.ttl)
}

func (l *Cache[V]) SetWithTTL(key string, data V, ttl time.Duration) {
	l.c.Set(key, data, ttl)
}

func (l *Cache[V]) Get(key string) (V, bool) {
	i := l.c.Get(key)
	var v V
	if i != nil {
		return i.Value(), !i.Expired()
	}
	return v, false
}

func (l *Cache[V]) Del(key string) {
	l.c.Delete(key)
}

func (l *Cache[V]) Name() string {
	return Name
}
