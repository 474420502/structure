package treeset

import (
	"fmt"
	"strings"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
)

// TreeSet
type TreeSet struct {
	tree *Tree
}

// New
func New(Compare compare.Compare) *TreeSet {
	return &TreeSet{tree: newAVL(Compare)}
}

// Add
func (set *TreeSet) Add(items ...interface{}) {
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
func (set *TreeSet) Iterator() *avl.Iterator {
	return set.tree.Iterator()
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
