package cache

type Cache[K comparable, V any] interface {
	Set(K, V)
	Get(K) (V, bool)
	Del(K)
}
