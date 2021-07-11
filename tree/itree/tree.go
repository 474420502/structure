package indextree

import (
	"github.com/474420502/structure/compare"
)

func init() {

}

type Node struct {
	Parent   *Node
	Children [2]*Node
	// direct   [2]*Node

	Size  int64
	Key   interface{}
	Value interface{}
}

type IndexTree struct {
	root    *Node
	compare compare.Compare
}

func New(comp compare.Compare) *IndexTree {
	return &IndexTree{compare: comp, root: &Node{}}
}

func (tree *IndexTree) getRoot() *Node {
	return tree.root.Children[0]
}

func (tree *IndexTree) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.Size
	}
	return 0
}

func (tree *IndexTree) Get(key interface{}) (interface{}, bool) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	for cur != nil {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:
			cur = cur.Children[L]
		case c > 0:
			cur = cur.Children[R]
		default:
			return cur.Value, true
		}
	}
	return nil, false
}

func (tree *IndexTree) Put(key, value interface{}) bool {

	cur := tree.getRoot()
	if cur == nil {
		tree.root.Children[0] = &Node{Key: key, Value: value, Size: 1, Parent: tree.root}
		return true
	}

	// var left *Node = nil
	// var right *Node = nil

	const L = 0
	const R = 1

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:
			// right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &Node{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[L] = node

				tree.fixSize(cur)
				tree.fixPut(cur)
				return true
			}

		case c > 0:

			// left = cur
			if cur.Children[R] != nil {
				cur = cur.Children[R]
			} else {

				node := &Node{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[R] = node

				tree.fixSize(cur)
				tree.fixPut(cur)
				return true
			}
		default:
			return false
		}
	}
}

func (tree *IndexTree) Index(i int64) (key, value interface{}) {
	node := tree.index(i)
	return node.Key, node.Value
}

func (tree *IndexTree) index(i int64) *Node {

	defer func() {
		if err := recover(); err != nil {
			panic(ErrOutOfIndex)
		}
	}()

	const L = 0
	const R = 1

	cur := tree.getRoot()

	var offset int64 = 0
	for {
		lsize := getSize(cur.Children[L])
		idx := lsize + offset
		if idx > i {
			cur = cur.Children[L]
		} else if idx < i {
			cur = cur.Children[R]
			offset += lsize + 1
		} else {
			return cur
		}
	}

}

func (tree *IndexTree) Seek()

// func (tree *IndexTree) Range(start, end int64) (result [][2]interface{}) {
// 	snode := tree.index(start)
// 	result = append(result, [2]interface{}{snode.Key, snode.Value})
// 	if
// 	tree.traversal(snode, func(cur *Node) bool {
// 		result = append(result, [2]interface{}{cur.Key, cur.Value})
// 		if start == end {
// 			return false
// 		}
// 		return true
// 	})
// }

// func (tree *IndexTree) traversal(cur *Node, do func(cur *Node) bool) {
// 	if cur == nil || cur == tree.root {
// 		return
// 	}

// 	tree.traversal(cur.Children[0], do)
// 	if do(cur) {
// 		tree.traversal(cur.Children[1], do)
// 	}
// 	tree.traversal(cur.Parent, do)
// }
