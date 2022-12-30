package treeset

import (
	"fmt"
	"strings"

	"github.com/474420502/structure/compare"
)

// TreeSet
type TreeSet[T any] struct {
	tree *Tree[T]
}

// New
func New[T any](Compare compare.Compare[T]) *TreeSet[T] {
	return &TreeSet[T]{tree: newAVL(Compare)}
}

// Add Not Cover the key of node. if item exists. return false
func (set *TreeSet[T]) Add(item T) bool {
	return set.tree.Put(item)

}

// Set   Set the key of node.
//
// like add(). if compare is special, will cover key.
// eg. k1 = {a:1,b:2}. k2 = {a:1, b:3}.  k1.a == k2.a. the key will be covered by k2
func (set *TreeSet[T]) Set(item T) bool {
	return set.tree.Set(item)
}

// Sets   Cover the key of nodes
func (set *TreeSet[T]) Sets(items ...T) {
	for _, item := range items {
		set.tree.Set(item)
	}
}

// Adds
func (set *TreeSet[T]) Adds(items ...T) {
	for _, item := range items {
		set.tree.Put(item)
	}
}

// Remove
func (set *TreeSet[T]) Remove(items ...T) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Values
func (set *TreeSet[T]) Values() []interface{} {
	return set.tree.Values()
}

// Contains
func (set *TreeSet[T]) Contains(item T) bool {
	if _, ok := set.tree.Get(item); ok {
		return true
	}
	return false
}

// Empty
func (set *TreeSet[T]) Empty() bool {
	return set.Size() == 0
}

// Clear
func (set *TreeSet[T]) Clear() {
	set.tree.Clear()
}

// Size
func (set *TreeSet[T]) Size() int {
	return set.tree.Size()
}

// Iterator avl Iterator
func (set *TreeSet[T]) Iterator() *Iterator[T] {
	return newIterator(set.tree)
}

// Traverse 从左到右遍历. left -> right
func (set *TreeSet[T]) Traverse(tr func(v interface{}) bool) {
	set.tree.Traverse(tr)
}

// String
func (set *TreeSet[T]) String() string {
	// content := "HashSet\n"
	var content = ""
	items := []string{}

	set.tree.Traverse(func(v interface{}) bool {
		items = append(items, fmt.Sprintf("%v", v))
		return true
	})

	content += "(" + strings.Join(items, ", ") + ")"
	return content
}
