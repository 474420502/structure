package hashmap

import (
	"fmt"
)

// Slice the KeyValue
type Slice[K comparable, V any] struct {
	Key   K
	Value V
}

// HashMap is a generic hash map on top of Go's native map.
type HashMap[K comparable, V any] struct {
	hm map[K]V
}

// New instantiates a hash map.
func New[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{hm: make(map[K]V)}
}

// NewWithCap instantiates a hash map with capacity.
func NewWithCap[K comparable, V any](cap int) *HashMap[K, V] {
	return &HashMap[K, V]{hm: make(map[K]V, cap)}
}

// Put inserts element into the map With Not Cover. if key exists return false. else return true
func (hm *HashMap[K, V]) Put(key K, value V) bool {
	if _, ok := hm.hm[key]; !ok {
		hm.hm[key] = value
		return true
	}
	return false
}

// InsertIfAbsent inserts a value only when the key does not exist.
func (hm *HashMap[K, V]) InsertIfAbsent(key K, value V) bool {
	return hm.Put(key, value)
}

// Set inserts element into the map With Set.
func (hm *HashMap[K, V]) Set(key K, value V) {
	hm.hm[key] = value
}

// Upsert sets the value and reports whether an existing value was replaced.
func (hm *HashMap[K, V]) Upsert(key K, value V) bool {
	_, replaced := hm.hm[key]
	hm.hm[key] = value
	return replaced
}

// Get get the element by key
func (hm *HashMap[K, V]) Get(key K) (value V, isfound bool) {
	value, isfound = hm.hm[key]
	return
}

// Remove remove the element by key
func (hm *HashMap[K, V]) Remove(key K) {
	delete(hm.hm, key)
}

// Delete removes a key and returns the previous value when present.
func (hm *HashMap[K, V]) Delete(key K) (value V, ok bool) {
	value, ok = hm.hm[key]
	if ok {
		delete(hm.hm, key)
	}
	return
}

// Empty if the hashmap is empty, return true
func (hm *HashMap[K, V]) Empty() bool {
	return len(hm.hm) == 0
}

// Size return the size of hashmap
func (hm *HashMap[K, V]) Size() int {
	return len(hm.hm)
}

// Len returns the number of elements.
func (hm *HashMap[K, V]) Len() int {
	return len(hm.hm)
}

// Keys return the all keys of hashmap. non order
func (hm *HashMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(hm.hm))
	for key := range hm.hm {
		keys = append(keys, key)
	}
	return keys
}

// Values return the all values of hashmap. non order
func (hm *HashMap[K, V]) Values() []V {
	values := make([]V, 0, len(hm.hm))
	for _, value := range hm.hm {
		values = append(values, value)
	}
	return values
}

// Slices return the all keyvalue of hashmap. non order
func (hm *HashMap[K, V]) Slices() []Slice[K, V] {
	slices := make([]Slice[K, V], 0, len(hm.hm))
	for key, value := range hm.hm {
		slices = append(slices, Slice[K, V]{Key: key, Value: value})
	}
	return slices
}

// Clear clear the hashmap
func (hm *HashMap[K, V]) Clear() {
	hm.hm = make(map[K]V)
}

// String print the hashmap
func (hm *HashMap[K, V]) String() string {
	return fmt.Sprintf("%v", hm.hm)
}
