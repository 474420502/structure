package treelist

import (
	"github.com/474420502/structure/compare"
)

func init() {

}

type Slice struct {
	Key   []byte
	Value []byte
}

type Node struct {
	Parent   *Node
	Children [2]*Node
	Direct   [2]*Node

	Size int64

	Slice
}

type Tree struct {
	root    *Node
	compare compare.Compare
}

func New() *Tree {
	return &Tree{compare: compare.Bytes, root: &Node{}}
}

func (tree *Tree) Iterator() *Iterator {
	return &Iterator{tree: tree}
}

func (tree *Tree) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.Size
	}
	return 0
}

func (tree *Tree) Get(key []byte) (interface{}, bool) {
	if cur := tree.getNode(key); cur != nil {
		return cur.Value, true
	}
	return nil, false
}

func (tree *Tree) getNode(key []byte) *Node {
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

func (tree *Tree) getNodeWithIndex(key []byte) (node *Node, idx int64) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	var offset int64 = getSize(cur.Children[L])
	for cur != nil {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:
			cur = cur.Children[L]
			if cur != nil {
				offset -= getSize(cur.Children[L]) + 1
			} else {
				return nil, -1
			}
		case c > 0:
			cur = cur.Children[R]
			if cur != nil {
				offset += getSize(cur.Children[L]) + 1
			} else {
				return nil, -1
			}
		default:
			return cur, offset
		}
	}
	return nil, -1
}

func (tree *Tree) seekNodeWithIndex(key []byte) (node *Node, idx int64, dir int) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	var offset int64 = getSize(cur.Children[L])
	var last *Node
	var c int
	for {
		c = tree.compare(key, cur.Key)
		last = cur

		switch {
		case c < 0:

			cur = cur.Children[L]
			if cur != nil {
				offset -= getSize(cur.Children[L]) + 1
			} else {
				return last, offset, c
			}

		case c > 0:

			cur = cur.Children[R]
			if cur != nil {
				offset += getSize(cur.Children[L]) + 1
			} else {
				return last, offset, c
			}

		default:
			return cur, offset, c
		}
	}

}

func (tree *Tree) Put(key, value []byte) bool {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	if cur == nil {
		node := &Node{Slice: Slice{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return true
	}

	var left *Node = nil
	var right *Node = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &Node{Parent: cur, Slice: Slice{Key: key, Value: value}, Size: 1}
				cur.Children[L] = node

				if left != nil {
					left.Direct[R] = node
				} else {
					tree.root.Direct[L] = node
				}

				if right != nil {
					right.Direct[L] = node
				} else {
					tree.root.Direct[R] = node
				}

				node.Direct[L] = left
				node.Direct[R] = right

				tree.fixPut(cur)
				return true
			}

		case c > 0:

			left = cur
			if cur.Children[R] != nil {
				cur = cur.Children[R]
			} else {
				node := &Node{Parent: cur, Slice: Slice{Key: key, Value: value}, Size: 1}
				cur.Children[R] = node

				if left != nil {
					left.Direct[R] = node
				} else {
					tree.root.Direct[L] = node
				}
				if right != nil {
					right.Direct[L] = node
				} else {
					tree.root.Direct[R] = node
				}

				node.Direct[L] = left
				node.Direct[R] = right

				tree.fixPut(cur)
				return true
			}
		default:
			return false
		}
	}

}

func (tree *Tree) Index(i int64) *Slice {
	node := tree.index(i)
	return &node.Slice
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

func (tree *Tree) IndexOf(key []byte) int64 {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	var offset int64 = getSize(cur.Children[L])
	for cur != nil {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:
			cur = cur.Children[L]
			offset -= getSize(cur.Children[L]) + 1
		case c > 0:
			cur = cur.Children[R]
			offset += getSize(cur.Children[L]) + 1
		default:
			return offset
		}
	}
	return -1
}

func (tree *Tree) rankNode(key []byte) (*Node, int64) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	var offset int64 = getSize(cur.Children[L])
	for cur != nil {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:
			cur = cur.Children[L]
			offset -= getSize(cur.Children[R]) + 1
		case c > 0:
			cur = cur.Children[R]
			offset += getSize(cur.Children[L]) + 1
		default:
			return cur, offset
		}
	}
	return nil, -1
}

// Traverse 遍历的方法 默认是LDR 从小到大 Compare 为 l < r
func (tree *Tree) Traverse(every func(s *Slice) bool) {
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
		if !every(&cur.Slice) {
			return false
		}
		if !traverasl(cur.Children[1]) {
			return false
		}
		return true
	}
	traverasl(root)
}

func (tree *Tree) Slices() []Slice {
	var mszie int64
	root := tree.getRoot()
	if root != nil {
		mszie = root.Size
	}
	result := make([]Slice, 0, mszie)
	tree.Traverse(func(s *Slice) bool {
		result = append(result, *s)
		return true
	})
	return result
}

func (tree *Tree) Remove(key []byte) (s *Slice) {
	const L = 0
	const R = 1

	if cur := tree.getNode(key); cur != nil {
		s = &Slice{Key: cur.Key, Value: cur.Value}

		if cur.Size == 1 {
			parent := cur.Parent
			if parent != tree.root {
				parent.Children[getRelationship(cur)] = nil
				tree.fixRemoveSize(parent)

				dright := cur.Direct[R]
				dleft := cur.Direct[L]
				if dright != nil {
					dright.Direct[L] = dleft
				}
				if dleft != nil {
					dleft.Direct[R] = dright
				}

			} else {
				parent.Children[0] = nil
				tree.root.Direct[L] = nil
				tree.root.Direct[R] = nil
			}

			return
		}

		lsize, rsize := getChildrenSize(cur)
		if lsize > rsize {
			prev := cur.Children[L]
			for prev.Children[R] != nil {
				prev = prev.Children[R]
			}

			cur.Key = prev.Key
			cur.Value = prev.Value

			prevParent := prev.Parent
			if prevParent == cur {
				cur.Children[L] = prev.Children[L]
				cleft := cur.Children[L]
				if cleft != nil {
					cleft.Parent = cur
				}
				// direct
				tree.fixRemoveSize(cur)
			} else {

				prevParent.Children[R] = prev.Children[L]
				if prevParent.Children[R] != nil {
					prevParent.Children[R].Parent = prevParent
				}
				tree.fixRemoveSize(prevParent)
			}

			dleft := cur.Direct[L].Direct[L]
			if dleft != nil {
				dleft.Direct[R] = cur
			}
			cur.Direct[L] = dleft

		} else {

			next := cur.Children[R]
			for next.Children[L] != nil {
				next = next.Children[L]
			}

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

			dright := cur.Direct[R].Direct[R]
			if dright != nil {
				dright.Direct[L] = cur
			}
			cur.Direct[R] = dright
		}

		return
	}

	return nil
}

func (tree *Tree) Clear() {
	tree.root.Children[0] = nil
}

func (tree *Tree) getRoot() *Node {
	return tree.root.Children[0]
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
