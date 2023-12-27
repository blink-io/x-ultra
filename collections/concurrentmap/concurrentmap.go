package concurrentmap

import (
	"github.com/orcaman/concurrent-map/v2"
)

type ConcurrentMap[K comparable, V any] struct {
	cmap.ConcurrentMap[K, V]
}
