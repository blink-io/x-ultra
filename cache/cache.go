package cache

import (
	"time"
)

const ErrNameSuffix = "-with-error"

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

// ErrCache defines cache abstract with error
type ErrCache[V any] interface {
	Name() string
	Set(string, V) error
	Get(string) (V, bool, error)
	Del(string) error
}

type ErrTTLSetter[V any] interface {
	SetWithTTL(string, V, time.Duration) error
}

type ErrTTLCache[V any] interface {
	ErrCache[V]
	ErrTTLSetter[V]
}
