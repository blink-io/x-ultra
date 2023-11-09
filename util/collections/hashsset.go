package collections

type HashSet[K comparable] struct {
	base HashMap[K, struct{}]
}

// Insert an item
func (tr *HashSet[K]) Insert(key K) {
	tr.base.Set(key, struct{}{})
}

// Get a value for key
func (tr *HashSet[K]) Contains(key K) bool {
	_, ok := tr.base.Get(key)
	return ok
}

// Len returns the number of items in the tree
func (tr *HashSet[K]) Len() int {
	return tr.base.Len()
}

// Delete an item
func (tr *HashSet[K]) Delete(key K) {
	tr.base.Delete(key)
}

func (tr *HashSet[K]) Scan(iter func(key K) bool) {
	tr.base.Scan(func(key K, value struct{}) bool {
		return iter(key)
	})
}

// Keys returns all keys as a slice
func (tr *HashSet[K]) Keys() []K {
	return tr.base.Keys()
}

// Copy the set. This is a copy-on-write operation and is very fast because
// it only performs a shadow copy.
func (tr *HashSet[K]) Copy() *HashSet[K] {
	tr2 := new(HashSet[K])
	tr2.base = *tr.base.Copy()
	return tr2
}

// GetPos gets a single keys/value nearby a position.
// The pos param can be any valid uint64. Useful for grabbing a random item
// from the HashSet.
func (tr *HashSet[K]) GetPos(pos uint64) (key K, ok bool) {
	key, _, ok = tr.base.GetPos(pos)
	return key, ok
}
