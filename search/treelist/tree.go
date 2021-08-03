package treelist

import (
	"bytes"
	"compress/zlib"
	"log"

	"github.com/474420502/structure/compare"
)

func init() {

}

type Slice struct {
	Key   []byte
	Value interface{}
}

type Node struct {
	Parent   *Node
	Children [2]*Node
	Direct   [2]*Node

	Size int64

	Slice
}

func (n *Node) String() string {
	return string(n.Key)
}

type Tree struct {
	root    *Node
	compare compare.Compare

	// rcount int
}

func New() *Tree {
	return &Tree{compare: compare.Bytes, root: &Node{}}
}

func (tree *Tree) SetCompare(comp compare.Compare) {
	tree.compare = comp
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

// PutDuplicate put, when key duplicate with call do. don,t change the key of `exists`, will break the tree of blance
// 				if duplicate, will return true.
func (tree *Tree) PutDuplicate(key []byte, value interface{}, do func(exists *Slice)) bool {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	if cur == nil {
		node := &Node{Slice: Slice{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return false
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
				return false
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
				return false
			}
		default:
			do(&cur.Slice)
			return true
		}
	}

}

func (tree *Tree) Put(key []byte, value interface{}) bool {
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
			log.Panicln(ErrOutOfIndex, i)
		}
	}()

	const L = 0
	const R = 1

	cur := tree.getRoot()
	var idx int64 = getSize(cur.Children[L])
	for {
		if idx > i {
			cur = cur.Children[L]
			idx -= getSize(cur.Children[R]) + 1
		} else if idx < i {
			cur = cur.Children[R]
			idx += getSize(cur.Children[L]) + 1
		} else {
			return cur
		}
	}

}

func (tree *Tree) IndexOf(key []byte) int64 {
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

func (tree *Tree) Remove(key []byte) *Slice {
	if cur := tree.getNode(key); cur != nil {
		return tree.removeNode(cur)
	}
	return nil
}

func (tree *Tree) Head() *Slice {
	h := tree.root.Direct[0]
	if h != nil {
		return &h.Slice
	}
	return nil
}

func (tree *Tree) RemoveHead() *Slice {
	if tree.getRoot() != nil {
		return tree.removeNode(tree.root.Direct[0])
	}
	return nil
}

func (tree *Tree) Tail() *Slice {
	t := tree.root.Direct[1]
	if t != nil {
		return &t.Slice
	}
	return nil
}

func (tree *Tree) RemoveTail() *Slice {
	if tree.getRoot() != nil {
		return tree.removeNode(tree.root.Direct[1])
	}
	return nil
}

func (tree *Tree) removeNode(cur *Node) (s *Slice) {
	const L = 0
	const R = 1
	s = &Slice{Key: cur.Key, Value: cur.Value}

	if cur.Size == 1 {
		parent := cur.Parent
		if parent != tree.root {
			parent.Children[getRelationship(cur)] = nil
			tree.fixRemoveSize(parent)

			dright := cur.Direct[R]
			dleft := cur.Direct[L]

			if dleft != nil {
				dleft.Direct[R] = dright
			} else {
				tree.root.Direct[L] = dright
			}

			if dright != nil {
				dright.Direct[L] = dleft
			} else {
				tree.root.Direct[R] = dleft
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
		} else {
			tree.root.Direct[L] = cur
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
		} else {
			tree.root.Direct[R] = cur
		}
		cur.Direct[R] = dright
	}

	return
}

// RemoveRange
func (tree *Tree) RemoveRange(low, hight []byte) {

	const L = 0
	const R = 1

	c := tree.compare(low, hight)
	if c > 0 {
		panic("key2 must greater than key1 or equal to")
	} else if c == 0 {
		tree.Remove(low)
		return
	}

	root := tree.getRangeRoot(low, hight)
	if root == nil {
		return
	}

	var ltrim, rtrim func(*Node) *Node
	var dleft *Node
	ltrim = func(root *Node) *Node {
		if root == nil {
			return nil
		}
		c = tree.compare(low, root.Key)
		if c > 0 {
			child := ltrim(root.Children[R])
			root.Children[R] = child
			if child != nil {
				child.Parent = root
			} else {
				dleft = root
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else if c < 0 {
			if root.Children[L] == nil {
				dleft = root.Direct[L]
			}
			return ltrim(root.Children[L])
		} else {
			dleft = root.Direct[L]
			return root.Children[L]
		}
	}

	var lgroup *Node
	if root.Children[L] != nil {
		lgroup = ltrim(root.Children[L])
	} else {
		dleft = root.Direct[L]
	}

	var dright *Node
	rtrim = func(root *Node) *Node {
		if root == nil {
			return nil
		}
		c = tree.compare(hight, root.Key)
		if c < 0 {
			child := rtrim(root.Children[L])
			root.Children[L] = child
			if child != nil {
				child.Parent = root
			} else {
				dright = root
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else if c > 0 {
			if root.Children[R] == nil {
				dright = root.Direct[R]
			}
			return rtrim(root.Children[R])
		} else {
			dright = root.Direct[R]
			return root.Children[R]
		}
	}

	var rgroup *Node
	if root.Children[R] != nil {
		rgroup = rtrim(root.Children[R])
	} else {
		dright = root.Direct[R]
	}

	if dleft != nil {
		dleft.Direct[R] = dright
	} else {
		tree.root.Direct[L] = dright
	}

	if dright != nil {
		dright.Direct[L] = dleft
	} else {
		tree.root.Direct[R] = dleft
	}

	if lgroup == nil && rgroup == nil {
		rparent := root.Parent
		size := root.Size
		root.Parent.Children[getRelationship(root)] = nil
		for rparent != tree.root {
			rparent.Size -= size
			rparent = rparent.Parent
		}
		return
	}

	// 左右树　拼接
	rsize := getSize(rgroup)
	lsize := getSize(lgroup)
	if lsize > rsize {
		tree.mergeGroups(root, lgroup, rgroup, rsize, R)
	} else {
		tree.mergeGroups(root, rgroup, lgroup, lsize, L)
	}
}

func (tree *Tree) mergeGroups(root *Node, group *Node, childGroup *Node, childSize int64, LR int) {
	rparent := root.Parent
	hand := group
	for hand.Children[LR] != nil {
		hand = hand.Children[LR]
	}
	hand.Children[LR] = childGroup
	if childGroup != nil {
		childGroup.Parent = hand
	}
	rparent.Children[getRelationship(root)] = group
	if group != nil {
		group.Parent = rparent
	}

	if childGroup != nil {
		parent := childGroup.Parent
		for parent != rparent {
			parent.Size += childSize
			temp := parent.Parent
			tree.fixRemoveRange(parent)
			parent = temp
		}
	}

	parent := rparent
	for parent != tree.root {
		parent.Size = getChildrenSumSize(parent) + 1
		parent = parent.Parent
	}
}

// RemoveRangeByIndex 1.range [low:hight] 2.low hight 必须包含存在的值.[low: hight+1] [low-1: hight].  [low-1: hight+1]. error: [low-1:low-2] or [hight+1:hight+2]
func (tree *Tree) RemoveRangeByIndex(low, hight int64) {

	defer func() {
		if err := recover(); err != nil {
			log.Panicln(ErrOutOfIndex, low, hight)
		}
	}()

	const L = 0
	const R = 1

	cur := tree.getRoot()
	var idx int64 = getSize(cur.Children[L])
	for {
		if idx > low && idx > hight {
			cur = cur.Children[L]
			idx -= getSize(cur.Children[R]) + 1
		} else if idx < hight && idx < low {
			cur = cur.Children[R]
			idx += getSize(cur.Children[L]) + 1
		} else {
			break
		}
	}

	root := cur
	// log.Println(low, hight, "low:", tree.index(low), "hight:", tree.index(hight), "root:", root)
	var ltrim, rtrim func(idx int64, dir int, root *Node) *Node
	var dleft *Node
	ltrim = func(idx int64, dir int, root *Node) *Node {
		if root == nil {
			return nil
		}

		if dir == R {
			idx += getSize(root.Children[L]) + 1
		} else {
			idx -= getSize(root.Children[R]) + 1
		}

		if idx < low {
			child := ltrim(idx, R, root.Children[R])
			root.Children[R] = child
			if child != nil {
				child.Parent = root
			} else {
				dleft = root
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else if idx > low {
			if root.Children[L] == nil {
				dleft = root.Direct[L]
			}
			return ltrim(idx, L, root.Children[L])
		} else {
			dleft = root.Direct[L]
			return root.Children[L]
		}
	}

	var lgroup *Node
	if root.Children[L] != nil {
		lgroup = ltrim(idx, L, root.Children[L])
	} else {
		dleft = root.Direct[L]
	}

	var dright *Node
	rtrim = func(idx int64, dir int, root *Node) *Node {
		if root == nil {
			return nil
		}

		if dir == R {
			idx += getSize(root.Children[L]) + 1
		} else {
			idx -= getSize(root.Children[R]) + 1
		}

		if idx > hight {
			child := rtrim(idx, L, root.Children[L])
			root.Children[L] = child
			if child != nil {
				child.Parent = root
			} else {
				dright = root
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else if idx < hight {
			if root.Children[R] == nil {
				dright = root.Direct[R]
			}
			return rtrim(idx, R, root.Children[R])
		} else {
			dright = root.Direct[R]
			return root.Children[R]
		}
	}

	var rgroup *Node
	if root.Children[R] != nil {
		rgroup = rtrim(idx, R, root.Children[R])
	} else {
		dright = root.Direct[R]
	}

	if dleft != nil {
		dleft.Direct[R] = dright
	} else {
		tree.root.Direct[L] = dright
	}

	if dright != nil {
		dright.Direct[L] = dleft
	} else {
		tree.root.Direct[R] = dleft
	}

	if lgroup == nil && rgroup == nil {
		rparent := root.Parent
		size := root.Size
		root.Parent.Children[getRelationship(root)] = nil
		for rparent != tree.root {
			rparent.Size -= size
			rparent = rparent.Parent
		}
		return
	}

	// 左右树　拼接
	rsize := getSize(rgroup)
	lsize := getSize(lgroup)
	if lsize > rsize {
		tree.mergeGroups(root, lgroup, rgroup, rsize, R)
	} else {
		tree.mergeGroups(root, rgroup, lgroup, lsize, L)
	}
}

func (tree *Tree) Clear() {
	tree.root.Children[0] = nil
}

func (tree *Tree) getRoot() *Node {
	return tree.root.Children[0]
}

// getRangeNodes 获取范围节点的左团和又团
func (tree *Tree) getRangeRoot(low, hight []byte) (root *Node) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	for cur != nil {
		c1 := tree.compare(low, cur.Key)
		c2 := tree.compare(hight, cur.Key)

		if c1 != c2 {
			return cur
		}

		if c1 < 0 {
			cur = cur.Children[L]
		} else if c1 > 0 {
			cur = cur.Children[R]
		} else {
			return cur
		}
	}
	return
}

// Trim range [low:hight]
func (tree *Tree) Trim(low, hight []byte) {

	const L = 0
	const R = 1

	root := tree.getRangeRoot(low, hight)

	var ltrim func(root *Node) *Node
	ltrim = func(root *Node) *Node {
		if root == nil {
			return nil
		}
		c := tree.compare(low, root.Key)
		if c > 0 {
			return ltrim(root.Children[R])
		} else if c < 0 {
			child := ltrim(root.Children[L])
			root.Children[L] = child
			if child != nil {
				child.Parent = root
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else {
			root.Children[L] = nil
			root.Size = getSize(root.Children[R]) + 1
			return root
		}
	}

	ltrim(root)

	var rtrim func(root *Node) *Node
	rtrim = func(root *Node) *Node {
		if root == nil {
			return nil
		}
		c := tree.compare(hight, root.Key)
		if c < 0 {
			return rtrim(root.Children[L])
		} else if c > 0 {
			child := rtrim(root.Children[R])
			root.Children[R] = child
			if child != nil {
				child.Parent = root
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else {
			root.Children[R] = nil
			root.Size = getSize(root.Children[L]) + 1
			return root
		}
	}

	rtrim(root)

	if root != tree.root {
		tree.root.Children[0] = root
	}

	if root != nil {
		root.Parent = tree.root

		lhand := root
		for lhand.Children[L] != nil {
			lhand = lhand.Children[L]
		}
		lhand.Direct[L] = nil
		tree.root.Direct[L] = lhand

		rhand := root
		for rhand.Children[R] != nil {
			rhand = rhand.Children[R]
		}
		rhand.Direct[R] = nil
		tree.root.Direct[R] = rhand
	}

}

// TrimByIndex range [low:hight]
func (tree *Tree) TrimByIndex(low, hight int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Panicln(ErrOutOfIndex, low, hight)
		}
	}()

	const L = 0
	const R = 1

	// log.Println(tree.debugString(true), string(low), string(hight))
	root := tree.getRoot()
	var idx int64 = getSize(root.Children[L])
	for {
		if idx > low && idx > hight {
			root = root.Children[L]
			idx -= getSize(root.Children[R]) + 1
		} else if idx < hight && idx < low {
			root = root.Children[R]
			idx += getSize(root.Children[L]) + 1
		} else {
			break
		}
	}

	var ltrim func(idx int64, root *Node) *Node
	ltrim = func(idx int64, root *Node) *Node {
		if root == nil {
			return nil
		}

		if idx < low {
			return ltrim(idx+getSize(root.Children[R].Children[L])+1, root.Children[R])
		} else if idx > low {
			child := ltrim(idx-getSize(root.Children[L].Children[R])-1, root.Children[L])
			root.Children[L] = child
			if child != nil {
				child.Parent = root
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else {
			root.Children[L] = nil
			root.Size = getSize(root.Children[R]) + 1
			return root
		}
	}

	ltrim(idx, root)

	var rtrim func(idx int64, root *Node) *Node
	rtrim = func(idx int64, root *Node) *Node {
		if root == nil {
			return nil
		}

		if idx > hight {
			return rtrim(idx-getSize(root.Children[L].Children[R])-1, root.Children[L])
		} else if idx < hight {
			child := rtrim(idx+getSize(root.Children[R].Children[L])+1, root.Children[R])
			root.Children[R] = child
			if child != nil {
				child.Parent = root
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else {
			root.Children[R] = nil
			root.Size = getSize(root.Children[L]) + 1
			return root
		}
	}

	rtrim(idx, root)
	// log.Println(debugNode(root))

	if root != tree.root {
		tree.root.Children[0] = root
	}

	if root != nil {
		root.Parent = tree.root

		lhand := root
		for lhand.Children[L] != nil {
			lhand = lhand.Children[L]
		}
		lhand.Direct[L] = nil
		tree.root.Direct[L] = lhand

		rhand := root
		for rhand.Children[R] != nil {
			rhand = rhand.Children[R]
		}
		rhand.Direct[R] = nil
		tree.root.Direct[R] = rhand
	}
}

func (tree *Tree) hashString() string {
	var buf = bytes.NewBuffer(nil)
	w := zlib.NewWriter(buf)

	tree.Traverse(func(s *Slice) bool {
		w.Write(s.Key)
		return true
	})

	err := w.Close()
	if err != nil {
		panic(err)
	}

	return string(buf.Bytes())
}
