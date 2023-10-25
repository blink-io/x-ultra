package cache

// Cache defines cache abstract
type Cache[V any] interface {
	Set(string, V)
	Get(string) (V, bool)
	Del(string)
}
