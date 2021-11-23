package indextreetest

import (
	"fmt"

	"github.com/474420502/structure/compare"
)

func init() {

}

type Node struct {
	Parent   *Node
	Children [2]*Node

	Size  int64
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

func (tree *Tree) String() string {
	return tree.debugString(true)
}

func (tree *Tree) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.Size
	}
	return 0
}

func (tree *Tree) Get(key interface{}) (interface{}, bool) {
	if cur := tree.getNode(key); cur != nil {
		return cur.Value, true
	}
	return nil, false
}

// Put 插入成功,返回true. 存在不插入 返回false
func (tree *Tree) Put(key, value interface{}) bool {

	cur := tree.getRoot()
	if cur == nil {
		tree.root.Children[0] = &Node{Key: key, Value: value, Size: 1, Parent: tree.root}
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
				node := &Node{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[L] = node
				tree.fixPut(cur)
				return true
			}

		case c > 0:

			if cur.Children[R] != nil {
				cur = cur.Children[R]
			} else {
				node := &Node{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[R] = node
				tree.fixPut(cur)
				return true
			}
		default:
			return false
		}
	}

}

// Set 插入成功,返回false. 存在覆盖 返回true
func (tree *Tree) Set(key, value interface{}) bool {

	cur := tree.getRoot()
	if cur == nil {
		tree.root.Children[0] = &Node{Key: key, Value: value, Size: 1, Parent: tree.root}
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
				node := &Node{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[L] = node
				tree.fixPut(cur)
				return false
			}

		case c > 0:

			if cur.Children[R] != nil {
				cur = cur.Children[R]
			} else {
				node := &Node{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[R] = node
				tree.fixPut(cur)
				return false
			}
		default:
			cur.Key = value
			cur.Value = value
			return true
		}
	}

}

func (tree *Tree) Index(i int64) (key, value interface{}) {
	node := tree.index(i)
	return node.Key, node.Value
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

func (tree *Tree) RankOf(key interface{}) (idx int64, isExists bool) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	if cur == nil {
		return 0, false
	}

	var offset int64 = getSize(cur.Children[L])
	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:
			cur = cur.Children[L]
			if cur == nil {
				return offset - 1, false
			}
			offset -= getSize(cur.Children[R]) + 1
		case c > 0:
			cur = cur.Children[R]
			if cur == nil {
				return offset + 1, false
			}
			offset += getSize(cur.Children[L]) + 1
		default:
			return offset, true
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
		mszie = root.Size
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

		if cur.Size == 1 {
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

func (tree *Tree) RemoveIndex(index int64) interface{} {
	const L = 0
	const R = 1

	if cur := tree.index(index); cur != nil {

		if cur.Size == 1 {
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

func (tree *Tree) RemoveRange(low, hight interface{}) {

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
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else {
			return ltrim(root.Children[L])
		}
	}

	var lgroup *Node
	if root.Children[L] != nil {
		lgroup = ltrim(root.Children[L])
	}

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
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else {
			return rtrim(root.Children[R])
		}
	}

	var rgroup *Node
	if root.Children[R] != nil {
		rgroup = rtrim(root.Children[R])
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

// RemoveRangeByIndex 1.range [low:hight] 2.low hight 必须包含存在的值.[low: hight+1] [low-1: hight].  [low-1: hight+1]. error: [low-1:low-2] or [hight+1:hight+2]
func (tree *Tree) RemoveRangeByIndex(low, hight int64) {
	if low > hight {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf(errOutOfIndex, low, hight))
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
	var ltrim, rtrim func(idx int64, dir int, root *Node) *Node
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
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else if idx > low {
			return ltrim(idx, L, root.Children[L])
		} else {
			return root.Children[L]
		}
	}

	var lgroup *Node
	if root.Children[L] != nil {
		lgroup = ltrim(idx, L, root.Children[L])
	}

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
			}
			root.Size = getChildrenSumSize(root) + 1
			return root
		} else if idx < hight {
			return rtrim(idx, R, root.Children[R])
		} else {
			return root.Children[R]
		}
	}

	var rgroup *Node
	if root.Children[R] != nil {
		rgroup = rtrim(idx, R, root.Children[R])
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

// Trim 保留区间
func (tree *Tree) Trim(low, high interface{}) {
	// root := tree.getRoot()

	if tree.compare(low, high) > 0 {
		panic(errLowerGtHigh)
	}

	const L = 0
	const R = 1

	root := tree.getRangeRoot(low, high)

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
		c := tree.compare(high, root.Key)
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
	}

}

// TrimByIndex 保留区间 range [low:hight]
func (tree *Tree) TrimByIndex(low, high int64) {
	if low > high {
		panic(errLowerGtHigh)
	}

	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf(errOutOfIndex, low, high))
		}
	}()

	const L = 0
	const R = 1

	// log.Println(tree.debugString(true), string(low), string(hight))
	root := tree.getRoot()
	var idx int64 = getSize(root.Children[L])
	for {
		if idx > low && idx > high {
			root = root.Children[L]
			idx -= getSize(root.Children[R]) + 1
		} else if idx < high && idx < low {
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

		if idx > high {
			return rtrim(idx-getSize(root.Children[L].Children[R])-1, root.Children[L])
		} else if idx < high {
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

		rhand := root
		for rhand.Children[R] != nil {
			rhand = rhand.Children[R]
		}

	}
}

// SplitContain Split 原树不包含Key. 拆分到返回的树
func (tree *Tree) Split(key interface{}) *Tree {
	root := tree.getRoot()
	if root == nil {
		return nil
	}

	const L = 0
	const R = 1

	cur := root
	// 寻找左右根
	var lroot, rroot *Node

	for cur != nil {
		c := tree.compare(cur.Key, key)
		if c < 0 {
			lroot = cur
			cur = cur.Children[R]
			if rroot != nil {
				break
			}

		} else {
			rroot = cur
			cur = cur.Children[L]
			if lroot != nil {
				break
			}
		}
	}

	var traverse func(cur *Node, lroot, rroot *Node)
	traverse = func(cur, lroot, rroot *Node) {
		if cur == nil { // 就算是nil也要赋值拼接
			lroot.Children[R] = nil
			lroot.Size = getSize(lroot.Children[L]) + 1

			rroot.Children[L] = nil
			rroot.Size = getSize(rroot.Children[R]) + 1
			return
		}

		c := tree.compare(cur.Key, key)
		switch {
		case c > 0:
			rroot.Children[L] = cur
			cur.Parent = rroot
			traverse(cur.Children[L], lroot, cur)
			rroot.Size = getChildrenSumSize(rroot) + 1
		case c < 0:
			lroot.Children[R] = cur
			cur.Parent = lroot
			traverse(cur.Children[R], cur, rroot)
			lroot.Size = getChildrenSumSize(lroot) + 1
		default:

			// 拼接右则
			rroot.Children[L] = cur
			cur.Parent = rroot

			left := cur.Children[L]
			cur.Children[L] = nil
			cur.Size = getSize(cur.Children[R]) + 1
			rroot.Size = getChildrenSumSize(rroot) + 1

			// 拼接左则
			lroot.Children[R] = left
			if left != nil {
				left.Parent = lroot
			}
			lroot.Size = getChildrenSumSize(lroot) + 1
		}

	}

	if lroot == nil {
		tree.root.Children[0] = nil

		rtree := New(tree.compare)
		rtree.root.Children[0] = root
		root.Parent = rtree.root
		return rtree
	}

	if rroot == nil {
		return New(tree.compare)
	}

	rtree := New(tree.compare)
	if lroot.Parent != rroot {
		rtree.root.Children[0] = rroot
		rroot.Parent = rtree.root

	} else {
		rtree.root.Children[0] = root
		root.Parent = rtree.root

		tree.root.Children[0] = lroot
		lroot.Parent = tree.root
	}

	// log.Println(lroot.Key, rroot.Key)

	lroot.Children[R] = nil
	rroot.Children[L] = nil
	traverse(cur, lroot, rroot)

	// 根链接错误

	for lroot != tree.root {
		lroot.Size = getChildrenSumSize(lroot) + 1
		lroot = lroot.Parent
	}

	for rroot != rtree.root {
		rroot.Size = getChildrenSumSize(rroot) + 1
		rroot = rroot.Parent
	}

	return rtree
}

// SplitContain Split 原树包含Key
func (tree *Tree) SplitContain(key interface{}) *Tree {
	root := tree.getRoot()
	if root == nil {
		return nil
	}

	const L = 0
	const R = 1

	cur := root
	// 寻找左右根
	var lroot, rroot *Node

	for cur != nil {
		c := tree.compare(cur.Key, key)
		if c > 0 {
			rroot = cur
			cur = cur.Children[L]
			if lroot != nil {
				break
			}
		} else {
			lroot = cur
			cur = cur.Children[R]
			if rroot != nil {
				break
			}
		}
	}

	var traverse func(cur *Node, lroot, rroot *Node)
	traverse = func(cur, lroot, rroot *Node) {
		if cur == nil { // 就算是nil也要赋值拼接
			lroot.Children[R] = nil
			lroot.Size = getSize(lroot.Children[L]) + 1

			rroot.Children[L] = nil
			rroot.Size = getSize(rroot.Children[R]) + 1
			return
		}

		c := tree.compare(cur.Key, key)
		switch {
		case c > 0:
			rroot.Children[L] = cur
			cur.Parent = rroot
			traverse(cur.Children[L], lroot, cur)
			rroot.Size = getChildrenSumSize(rroot) + 1
		case c < 0:
			lroot.Children[R] = cur
			cur.Parent = lroot
			traverse(cur.Children[R], cur, rroot)
			lroot.Size = getChildrenSumSize(lroot) + 1
		default:

			// 拼接左则
			lroot.Children[R] = cur
			cur.Parent = lroot

			right := cur.Children[R]
			cur.Children[R] = nil
			cur.Size = getSize(cur.Children[L]) + 1
			lroot.Size = getChildrenSumSize(lroot) + 1

			// 拼接右则
			rroot.Children[L] = right
			if right != nil {
				right.Parent = rroot
			}
			rroot.Size = getChildrenSumSize(rroot) + 1
		}

	}

	if lroot == nil {
		tree.root.Children[0] = nil

		rtree := New(tree.compare)
		rtree.root.Children[0] = root
		root.Parent = rtree.root
		return rtree
	}

	if rroot == nil {
		return New(tree.compare)
	}

	rtree := New(tree.compare)
	if lroot.Parent != rroot {
		rtree.root.Children[0] = rroot
		rroot.Parent = rtree.root

	} else {
		rtree.root.Children[0] = root
		root.Parent = rtree.root

		tree.root.Children[0] = lroot
		lroot.Parent = tree.root
	}

	// log.Println(lroot.Key, rroot.Key)

	lroot.Children[R] = nil
	rroot.Children[L] = nil
	traverse(cur, lroot, rroot)

	// 根链接错误

	for lroot != tree.root {
		lroot.Size = getChildrenSumSize(lroot) + 1
		lroot = lroot.Parent
	}

	for rroot != rtree.root {
		rroot.Size = getChildrenSumSize(rroot) + 1
		rroot = rroot.Parent
	}

	return rtree
}

func (tree *Tree) Clear() {
	tree.root.Children[0] = nil
}
