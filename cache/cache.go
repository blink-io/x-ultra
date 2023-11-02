package cache

// Cache defines cache abstract
type Cache[V any] interface {
	Name() string
	Set(string, V)
	Get(string) (V, bool)
	Del(string)
}

// CacheWithError defines cache abstract with error return
type CacheWithError[V any] interface {
	Name() string
	Set(string, V) error
	Get(string) (V, bool, error)
	Del(string) error
}
