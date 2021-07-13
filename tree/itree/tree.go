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

	size  int64
	Key   interface{}
	Value interface{}
}

type Tree struct {
	root    *Node
	compare compare.Compare
}

func New(comp compare.Compare) *Tree {
	return &Tree{compare: comp, root: &Node{}}
}

func (tree *Tree) getRoot() *Node {
	return tree.root.Children[0]
}

func (tree *Tree) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.size
	}
	return 0
}

func (tree *Tree) Get(key interface{}) (interface{}, bool) {
	if cur := tree.getNode(key); cur != nil {
		return cur.Value, true
	}
	return nil, false
}

func (tree *Tree) getNode(key interface{}) *Node {
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
			return cur
		}
	}
	return nil
}

func (tree *Tree) Put(key, value interface{}) bool {

	cur := tree.getRoot()
	if cur == nil {
		tree.root.Children[0] = &Node{Key: key, Value: value, size: 1, Parent: tree.root}
		return true
	}

	const L = 0
	const R = 1

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {
				node := &Node{Parent: cur, Key: key, Value: value, size: 1}
				cur.Children[L] = node
				tree.fixPut(cur)
				return true
			}

		case c > 0:

			if cur.Children[R] != nil {
				cur = cur.Children[R]
			} else {
				node := &Node{Parent: cur, Key: key, Value: value, size: 1}
				cur.Children[R] = node
				tree.fixPut(cur)
				return true
			}
		default:
			return false
		}
	}

}

func (tree *Tree) Index(i int64) (key, value interface{}) {
	node := tree.index(i)
	return node.Key, node.Value
}

func (tree *Tree) index(i int64) *Node {

	defer func() {
		if err := recover(); err != nil {
			panic(ErrOutOfIndex)
		}
	}()

	const L = 0
	const R = 1

	cur := tree.getRoot()

	var offset int64 = getSize(cur.Children[L])
	for {
		if i < offset {
			cur = cur.Children[L]
			offset -= getSize(cur.Children[L]) + 1
		} else if i > offset {
			cur = cur.Children[R]
			offset += getSize(cur.Children[L]) + 1
		} else {
			return cur
		}
	}

}

func (tree *Tree) IndexOf(key interface{}) int64 {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	if cur == nil {
		return -1
	}

	var offset int64 = getSize(cur.Children[L])
	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:
			cur = cur.Children[L]
			if cur == nil {
				return -1
			}
			offset -= getSize(cur.Children[R]) + 1
		case c > 0:
			cur = cur.Children[R]
			if cur == nil {
				return -1
			}
			offset += getSize(cur.Children[L]) + 1
		default:
			return offset
		}
	}

}

// Traverse 遍历的方法 默认是LDR 从小到大 Compare 为 l < r
func (tree *Tree) Traverse(every func(k, v interface{}) bool) {
	root := tree.getRoot()
	if root == nil {
		return
	}

	var traverasl func(cur *Node) bool
	traverasl = func(cur *Node) bool {
		if cur == nil {
			return true
		}
		if !traverasl(cur.Children[0]) {
			return false
		}
		if !every(cur.Key, cur.Value) {
			return false
		}
		if !traverasl(cur.Children[1]) {
			return false
		}
		return true
	}
	traverasl(root)
}

func (tree *Tree) Values() []interface{} {
	var mszie int64
	root := tree.getRoot()
	if root != nil {
		mszie = root.size
	}
	result := make([]interface{}, 0, mszie)
	tree.Traverse(func(k, v interface{}) bool {
		result = append(result, v)
		return true
	})
	return result
}

func (tree *Tree) Remove(key interface{}) interface{} {
	const L = 0
	const R = 1

	if cur := tree.getNode(key); cur != nil {

		if cur.size == 1 {
			parent := cur.Parent
			parent.Children[getRelationship(cur)] = nil
			tree.fixRemoveSize(parent)
			return cur.Value
		}

		lsize, rsize := getChildrenSize(cur)
		if lsize > rsize {
			prev := cur.Children[L]
			for prev.Children[R] != nil {
				prev = prev.Children[R]
			}

			value := cur.Value
			cur.Key = prev.Key
			cur.Value = prev.Value

			prevParent := prev.Parent
			if prevParent == cur {
				cur.Children[L] = prev.Children[L]
				if cur.Children[L] != nil {
					cur.Children[L].Parent = cur
				}
				tree.fixRemoveSize(cur)
			} else {
				prevParent.Children[R] = prev.Children[L]
				if prevParent.Children[R] != nil {
					prevParent.Children[R].Parent = prevParent
				}
				tree.fixRemoveSize(prevParent)
			}

			return value
		} else {

			next := cur.Children[R]
			for next.Children[L] != nil {
				next = next.Children[L]
			}

			value := cur.Value
			cur.Key = next.Key
			cur.Value = next.Value

			nextParent := next.Parent
			if nextParent == cur {
				cur.Children[R] = next.Children[R]
				if cur.Children[R] != nil {
					cur.Children[R].Parent = cur
				}
				tree.fixRemoveSize(cur)
			} else {
				nextParent.Children[L] = next.Children[R]
				if nextParent.Children[L] != nil {
					nextParent.Children[L].Parent = nextParent
				}
				tree.fixRemoveSize(nextParent)
			}

			return value

		}
	}

	return nil
}

func (tree *Tree) Clear() {
	tree.root.Children[0] = nil
}

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
