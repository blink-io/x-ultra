package cache

import (
	"time"
)

// Cache defines cache abstract
type Cache[V any] interface {
	Name() string
	Set(string, V)
	Get(string) (V, bool)
	Del(string)
}

type TTLCache[V any] interface {
	Cache[V]
	TTLSetter[V]
}

type TTLSetter[V any] interface {
	SetWithTTL(string, V, time.Duration)
}

func IsTTLCache(v any) bool {
	_, ok := v.(TTLSetter[any])
	return ok
}
