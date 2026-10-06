package hashset

import (
	"fmt"
	"strings"
)

var nullItem = struct{}{}

// HashSet is a generic unordered set on top of Go's native map.
type HashSet[T comparable] struct {
	hm map[T]struct{}
}

// New instantiates an empty set.
func New[T comparable]() *HashSet[T] {
	return &HashSet[T]{hm: make(map[T]struct{})}
}

// Add inserts items into the set.
func (set *HashSet[T]) Add(items ...T) {
	for _, item := range items {
		if _, ok := set.hm[item]; !ok {
			set.hm[item] = nullItem
		}
	}
}

// Remove deletes items from the set.
func (set *HashSet[T]) Remove(items ...T) {
	for _, item := range items {
		delete(set.hm, item)
	}
}

// Values returns an unordered snapshot of all items.
func (set *HashSet[T]) Values() []T {
	values := make([]T, 0, len(set.hm))
	for item := range set.hm {
		values = append(values, item)
	}
	return values
}

// Contains reports whether item is a member of the set.
func (set *HashSet[T]) Contains(item T) bool {
	_, contains := set.hm[item]
	return contains
}

// Empty reports whether the set has no items.
func (set *HashSet[T]) Empty() bool {
	return len(set.hm) == 0
}

// Clear removes every item.
func (set *HashSet[T]) Clear() {
	set.hm = make(map[T]struct{})
}

// Size returns the number of items.
func (set *HashSet[T]) Size() int {
	return len(set.hm)
}

// Len returns the number of items.
func (set *HashSet[T]) Len() int {
	return len(set.hm)
}

// String renders the set for debugging.
func (set *HashSet[T]) String() string {
	content := "["
	items := make([]string, 0, len(set.hm))
	for k := range set.hm {
		items = append(items, fmt.Sprintf("%v", k))
	}
	content += strings.Join(items, ",")
	content += "]"
	return content
}
