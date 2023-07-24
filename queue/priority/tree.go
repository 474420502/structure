package treequeue

import (
	"log"

	"github.com/474420502/structure/compare"
)

type Tree[KEY, VALUE any] struct {
	Center *Node[KEY, VALUE]

	Left  *Node[KEY, VALUE]
	Right *Node[KEY, VALUE]

	Compare compare.Compare[KEY]
	// hight       int
	zero VALUE
	// rotateCount int
}

func New[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE] {

	tree := &Tree[KEY, VALUE]{
		Center:  &Node[KEY, VALUE]{Size: 0},
		Compare: Compare,
	}

	tree.Center.Children[0] = tree.Center
	return tree
}

func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	target, isExists := tree.put(tree.Center, 1, key)
	if isExists == 1 {
		return false
	}
	target.Value = value

	return true
}

func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	cur := tree.getfirst(key, tree.Center, 1, nil)
	log.Println(tree.viewEx(cur))
	if cur == nil {
		return tree.zero, false
	}

	return cur.Value, true
}

// Gets get the key all values. contains the same key
func (tree *Tree[KEY, VALUE]) Gets(key KEY) (result []VALUE) {
	cur := tree.seekRoot(key, tree.getRoot())
	if cur == nil {
		return
	}
	tree.traverse(key, cur, &result)
	return result
}

func (tree *Tree[KEY, VALUE]) Index(idx int) VALUE {

	defer func() {
		if err := recover(); err != nil {
			log.Panicf("%v size: %d index: %d", errOutOfIndex, tree.Size(), idx)
		}
	}()

	if idx < 0 {
		return tree.index(tree.Size() + idx).Value
	}
	return tree.index(idx).Value
}

func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool) {

	target := tree.removeLeft(key, tree.Center, 0, 1)
	if target != nil {
		return *target, true
	}
	return tree.zero, false
}

// RemoveIndex remove key value by index and return value that be removed
func (tree *Tree[KEY, VALUE]) RemoveIndex(index int) VALUE {
	defer func() {
		if err := recover(); err != nil {
			log.Panicf("%v size: %d index: %d", errOutOfIndex, tree.Size()+1, index)
		}
	}()

	cur := tree.Center
	var idx int = cur.Children[1].Children[0].getSize()

	if idx < 0 {
		return *tree.removeIndex(cur, 1, idx, tree.Size()+index)
	}
	return *tree.removeIndex(cur, 1, idx, index)
}

func (tree *Tree[KEY, VALUE]) PopHead() (VALUE, bool) {
	if tree.Size() == 0 {
		return tree.zero, false
	}

	parent := tree.Center
	cur := parent.Children[1]

	for cur.Children[0] != nil {
		cur.Size--
		parent = cur
		cur = cur.Children[0]
	}

	if parent == tree.Center {
		parent.Children[1] = cur.Children[1]
	} else {
		parent.Children[0] = cur.Children[1]
	}

	return cur.Value, true
}

func (tree *Tree[KEY, VALUE]) PopTail() (VALUE, bool) {
	if tree.Size() == 0 {
		return tree.zero, false
	}

	parent := tree.Center
	cur := parent.Children[1]

	for cur.Children[1] != nil {
		cur.Size--
		parent = cur
		cur = cur.Children[1]
	}

	parent.Children[1] = cur.Children[0]
	return cur.Value, true
}

func (tree *Tree[KEY, VALUE]) Clear() {
	tree.Center.Children[1] = nil
}

func (tree *Tree[KEY, VALUE]) Size() int {
	return getSize(tree.getRoot())
}

func (tree *Tree[KEY, VALUE]) Traverse(every func(KEY, VALUE) bool) {

	var traverse func(cur *Node[KEY, VALUE]) bool
	traverse = func(cur *Node[KEY, VALUE]) bool {
		if cur == nil {
			return true
		}
		if !traverse(cur.Children[0]) {
			return false
		}
		if !every(cur.Key, cur.Value) {
			return false
		}
		if !traverse(cur.Children[1]) {
			return false
		}
		return true
	}
	traverse(tree.getRoot())
}

func (tree *Tree[KEY, VALUE]) Values() []VALUE {
	if tree.Center.Children[1] == nil {
		return nil
	}
	result := make([]VALUE, 0, tree.Size())
	tree.Traverse(func(k KEY, v VALUE) bool {
		result = append(result, v)
		return true
	})
	return result
}

func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE] {
	return newIterator(tree)
}
