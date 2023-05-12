package avlex

import (
	"log"

	"github.com/474420502/structure/compare"
)

type Tree[KEY, VALUE any] struct {
	Center  *Node[KEY, VALUE]
	Compare compare.Compare[KEY]
	zero    VALUE
}

func NewTree[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE] {

	tree := &Tree[KEY, VALUE]{
		Center:  &Node[KEY, VALUE]{Height: 0},
		Compare: Compare,
	}

	tree.Center.Children[0] = tree.Center
	return tree
}

func (tree *Tree[KEY, VALUE]) View() (result string) {
	result = "\n"
	if tree.Center.Children[1] == nil {
		result += "└── nil"
		return
	}

	var nmap = make(map[*Node[KEY, VALUE]]int)
	outputfordebug(nmap, tree.Center.Children[1], "", true, &result)
	return
}

func (tree *Tree[KEY, VALUE]) check() (result string) {
	if !checkHeightTree(tree.Center.Children[1]) {
		log.Panic("height error")
	}
	return
}

func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	target, isRepeat, _ := tree.put(tree.Center, 1, key)
	target.Value = value
	return isRepeat
}

func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	cur := tree.get(key, tree.Center.Children[1])
	if cur == nil {
		return tree.zero, false
	}

	return cur.Value, true
}

func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool) {
	target := tree.remove(key, tree.Center, 0, 1)
	if target != nil {
		return target.Value, true
	}
	return tree.zero, false
}
