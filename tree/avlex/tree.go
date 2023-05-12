package avlex

import (
	"github.com/474420502/structure/compare"
)

type Tree[KEY, VALUE any] struct {
	Center  *Node[KEY, VALUE]
	Compare compare.Compare[KEY]
	Size    uint
	zero    VALUE
}

func NewTree[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE] {

	tree := &Tree[KEY, VALUE]{
		Center:  &Node[KEY, VALUE]{Height: 65535},
		Compare: Compare,
		Size:    0,
	}

	tree.Center.Children[0] = tree.Center
	return tree
}

func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool {
	target, isExists, _ := tree.put(tree.Center, 1, key)
	target.Key = key
	target.Value = value
	if !isExists {
		tree.Size += 1
	}
	return !isExists
}

func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	target, isExists, _ := tree.put(tree.Center, 1, key)
	if !isExists {
		target.Value = value
		tree.Size += 1
	}
	return !isExists
}

func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	cur := tree.get(key, tree.Center.Children[1])
	if cur == nil {
		return tree.zero, false
	}

	return cur.Value, true
}

func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool) {
	target, _ := tree.remove(key, tree.Center, 0, 1)
	if target != nil {
		tree.Size -= 1
		return *target, true
	}
	return tree.zero, false
}
