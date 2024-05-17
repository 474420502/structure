package avls

import (
	"github.com/474420502/structure/compare"
)

type Tree[KEY, VALUE any] struct {
	Center           *Node[KEY, VALUE]
	Compare          compare.Compare[KEY]
	size             uint
	zero             VALUE
	differenceHeight int8
	// rotateCount      int
}

func New[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE] {

	tree := &Tree[KEY, VALUE]{
		Center:           &Node[KEY, VALUE]{Height: 0},
		Compare:          Compare,
		size:             0,
		differenceHeight: 2,
	}

	tree.Center.Children[0] = tree.Center
	return tree
}

func NewEx[KEY, VALUE any](Compare compare.Compare[KEY], differenceHeight int8) *Tree[KEY, VALUE] {

	tree := &Tree[KEY, VALUE]{
		Center:           &Node[KEY, VALUE]{Height: 0},
		Compare:          Compare,
		size:             0,
		differenceHeight: differenceHeight,
	}

	tree.Center.Children[0] = tree.Center
	return tree
}

func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool {
	target, isExists, _ := tree.put(tree.Center, 1, key)
	target.Key = key
	target.Value = value
	if !isExists {
		tree.size += 1
	}
	return !isExists
}

func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	target, isExists, _ := tree.put(tree.Center, 1, key)
	if !isExists {
		target.Value = value
		tree.size += 1
	}
	return !isExists
}

func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	cur := tree.get(key, tree.getRoot())
	if cur == nil {
		return tree.zero, false
	}

	return cur.Value, true
}

func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool) {
	target, _ := tree.remove(key, tree.Center, 0, 1)
	if target != nil {
		tree.size -= 1
		return *target, true
	}
	return tree.zero, false
}

func (tree *Tree[KEY, VALUE]) Clear() {
	tree.Center.Children[1] = nil
	tree.size = 0
}

func (tree *Tree[KEY, VALUE]) Size() uint {
	return tree.size
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
	if tree.size == 0 {
		return nil
	}
	result := make([]VALUE, 0, tree.size)
	tree.Traverse(func(k KEY, v VALUE) bool {
		result = append(result, v)
		return true
	})
	return result
}

func (tree *Tree[KEY, VALUE]) Height() int8 {
	if tree.size == 0 {
		return 0
	}
	return tree.getRoot().Height
}

func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE] {
	return newIterator(tree)
}
