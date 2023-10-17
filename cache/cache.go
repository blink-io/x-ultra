package cache

type Cache[V any] interface {
	Set(string, V)
	Get(string) (V, bool)
	Del(string)
}
