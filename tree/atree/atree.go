package arraytree

import "github.com/474420502/structure/compare"

type Tree struct {
	compare compare.Compare
	size    int64
	datas   []*Node
}

type Slice struct {
	Key   interface{}
	Value interface{}
}

type Node struct {
	Slice
}

func New() *Tree {
	tree := &Tree{compare: compare.Int}
	tree.datas = make([]*Node, 1<<3-1)
	return tree
}

func (tree *Tree) Put(key, value interface{}) {
	dlen := len(tree.datas)
	var start = 0
	var end = dlen
	var idx int
	var cur *Node
	for {
		idx = start + (end-start)>>1
		if idx >= dlen {

		}
		cur = tree.datas[idx]
		if cur == nil {
			tree.datas[idx] = &Node{Slice{Key: key, Value: value}}
			tree.size++
			return
		}

		c := tree.compare(key, cur.Slice.Key)
		switch {
		case c < 0:
			end = idx
		case c > 0:
			start = idx
		default:
			tree.datas[idx].Value = value
			return
		}
	}
}
