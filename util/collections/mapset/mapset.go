package mapset

import (
	"github.com/deckarep/golang-set/v2"
)

type Set[T comparable] interface {
	mapset.Set[T]
}

type Iterator[T comparable] interface {
	mapset.Iterator[T]
}

func NewSet[T comparable]() Set[T] {
	return mapset.NewSet[T]()
}

func NewSetFromMapKeys[T comparable, V any](val map[T]V) Set[T] {
	return mapset.NewSetFromMapKeys[T](val)
}

func NewSetWithSize[T comparable](cardinality int) Set[T] {
	return mapset.NewSetWithSize[T](cardinality)
}

func NewThreadUnsafeSet[T comparable](vals ...T) Set[T] {
	return mapset.NewThreadUnsafeSet[T](vals...)
}

func NewThreadUnsafeSetFromMapKeys[T comparable, V any](val map[T]V) Set[T] {
	return mapset.NewThreadUnsafeSetFromMapKeys[T](val)
}

func NewThreadUnsafeSetWithSize[T comparable](cardinality int) Set[T] {
	return mapset.NewThreadUnsafeSetWithSize[T](cardinality)
}
