package treeset

import (
	"fmt"
	"strings"

	"github.com/474420502/structure/compare"
)

// TreeSet
type TreeSet struct {
	tree *Tree
}

// New
func New(Compare compare.Compare) *TreeSet {
	return &TreeSet{tree: newAVL(Compare)}
}

// Add Not Cover the key of node
func (set *TreeSet) Add(item interface{}) bool {
	return set.tree.Put(item)

}

// Set   Set the key of node.
//
// like add(). if compare is special, will cover key.
// eg. k1 = {a:1,b:2}. k2 = {a:1, b:3}.  k1.a == k2.a. the key will be covered by k2
func (set *TreeSet) Set(item interface{}) bool {
	return set.tree.Set(item)
}

// Sets   Cover the key of nodes
func (set *TreeSet) Sets(items ...interface{}) {
	for _, item := range items {
		set.tree.Set(item)
	}
}

// Adds
func (set *TreeSet) Adds(items ...interface{}) {
	for _, item := range items {
		set.tree.Put(item)
	}
}

// Remove
func (set *TreeSet) Remove(items ...interface{}) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Values
func (set *TreeSet) Values() []interface{} {
	return set.tree.Values()
}

// Contains
func (set *TreeSet) Contains(item interface{}) bool {
	if _, ok := set.tree.Get(item); ok {
		return true
	}
	return false
}

// Empty
func (set *TreeSet) Empty() bool {
	return set.Size() == 0
}

// Clear
func (set *TreeSet) Clear() {
	set.tree.Clear()
}

// Size
func (set *TreeSet) Size() int {
	return set.tree.Size()
}

// Iterator avl Iterator
func (set *TreeSet) Iterator() *Iterator {
	return newIterator(set.tree)
}

// Traverse 从左到右遍历. left -> right
func (set *TreeSet) Traverse(tr func(v interface{}) bool) {
	set.tree.Traverse(tr)
}

// String
func (set *TreeSet) String() string {
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
