package indextree

import (
	"fmt"

	"github.com/474420502/structure/compare"
)

var IsDebugString = false

// Node the node of tree
type hNode[T any] struct {
	Parent   *hNode[T]
	Children [2]*hNode[T]

	Size  int64
	Key   T
	Value interface{}
}

// Tree the struct of tree
type Tree[T any] struct {
	root    *hNode[T]
	compare compare.Compare[T]
}

// New create a object of tree
func New[T any](comp compare.Compare[T]) *Tree[T] {
	return &Tree[T]{compare: comp, root: &hNode[T]{}}
}

// String show the view of tree by chars
func (tree *Tree[T]) String() string {
	return tree.debugString(IsDebugString)
}

// Size get the size of tree
func (tree *Tree[T]) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.Size
	}
	return 0
}

// Get get value by key
func (tree *Tree[T]) Get(key T) (interface{}, bool) {
	if cur := tree.getNode(key); cur != nil {
		return cur.Value, true
	}
	return nil, false
}

// Put put value into tree  by Key . if key exists,not cover the value and return false. else return true
func (tree *Tree[T]) Put(key T, value interface{}) bool {

	cur := tree.getRoot()
	if cur == nil {
		tree.root.Children[0] = &hNode[T]{Key: key, Value: value, Size: 1, Parent: tree.root}
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
				node := &hNode[T]{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[L] = node
				tree.fixPut(cur)
				return true
			}

		case c > 0:

			if cur.Children[R] != nil {
				cur = cur.Children[R]
			} else {
				node := &hNode[T]{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[R] = node
				tree.fixPut(cur)
				return true
			}
		default:
			return false
		}
	}

}

// Set set value by Key. if key exists, cover the value and return true. else return false and put value into tree
func (tree *Tree[T]) Set(key T, value interface{}) bool {

	cur := tree.getRoot()
	if cur == nil {
		tree.root.Children[0] = &hNode[T]{Key: key, Value: value, Size: 1, Parent: tree.root}
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
				node := &hNode[T]{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[L] = node
				tree.fixPut(cur)
				return false
			}

		case c > 0:

			if cur.Children[R] != nil {
				cur = cur.Children[R]
			} else {
				node := &hNode[T]{Parent: cur, Key: key, Value: value, Size: 1}
				cur.Children[R] = node
				tree.fixPut(cur)
				return false
			}
		default:
			cur.Key = key
			cur.Value = value
			return true
		}
	}

}

// Index Indexing Ordered Data. like TopN
func (tree *Tree[T]) Index(i int64) (key T, value interface{}) {
	node := tree.index(i)
	return node.Key, node.Value
}

// Index Indexing Ordered Data. like TopN
func (tree *Tree[T]) IndexOf(key T) int64 {
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

// func (tree *Tree[T]) RankOf(key T) (idx int64, isExists bool) {
// 	const L = 0
// 	const R = 1

// 	cur := tree.getRoot()
// 	if cur == nil {
// 		return 0, false
// 	}

// 	var offset int64 = getSize(cur.Children[L])
// 	for {
// 		c := tree.compare(key, cur.Key)
// 		switch {
// 		case c < 0:
// 			cur = cur.Children[L]
// 			if cur == nil {
// 				return offset - 1, false
// 			}
// 			offset -= getSize(cur.Children[R]) + 1
// 		case c > 0:
// 			cur = cur.Children[R]
// 			if cur == nil {
// 				return offset + 1, false
// 			}
// 			offset += getSize(cur.Children[L]) + 1
// 		default:
// 			return offset, true
// 		}
// 	}

// }

// Traverse the traversal method defaults to LDR. from smallest to largest.
func (tree *Tree[T]) Traverse(every func(k T, v interface{}) bool) {
	root := tree.getRoot()
	if root == nil {
		return
	}

	var traverasl func(cur *hNode[T]) bool
	traverasl = func(cur *hNode[T]) bool {
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

// Values return all values. in order
func (tree *Tree[T]) Values() []interface{} {
	var mszie int64
	root := tree.getRoot()
	if root != nil {
		mszie = root.Size
	}
	result := make([]interface{}, 0, mszie)
	tree.Traverse(func(k T, v interface{}) bool {
		result = append(result, v)
		return true
	})
	return result
}

// Remove remove key value and return value that be removed
func (tree *Tree[T]) Remove(key T) interface{} {
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

// RemoveIndex remove key value by index and return value that be removed
func (tree *Tree[T]) RemoveIndex(index int64) interface{} {
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

// RemoveRange remove keys values by range. [low, high]
func (tree *Tree[T]) RemoveRange(low, high T) {

	const L = 0
	const R = 1

	c := tree.compare(low, high)
	if c > 0 {
		panic("key2 must greater than key1 or equal to")
	} else if c == 0 {
		tree.Remove(low)
		return
	}

	root := tree.getRangeRoot(low, high)
	if root == nil {
		return
	}

	var ltrim, rtrim func(*hNode[T]) *hNode[T]
	ltrim = func(root *hNode[T]) *hNode[T] {
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

	var lgroup *hNode[T]
	if root.Children[L] != nil {
		lgroup = ltrim(root.Children[L])
	}

	rtrim = func(root *hNode[T]) *hNode[T] {
		if root == nil {
			return nil
		}
		c = tree.compare(high, root.Key)
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

	var rgroup *hNode[T]
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

	// Left and right tree concat
	rsize := getSize(rgroup)
	lsize := getSize(lgroup)
	if lsize > rsize {
		tree.mergeGroups(root, lgroup, rgroup, rsize, R)
	} else {
		tree.mergeGroups(root, rgroup, lgroup, lsize, L)
	}
}

// RemoveRangeByIndex 1.remove range [low:hight]
// 2.low and hight that the range must contain a value that exists. eg: [low: hight+1] [low-1: hight].  [low-1: hight+1]. error: [low-1:low-2] or [hight+1:hight+2]
func (tree *Tree[T]) RemoveRangeByIndex(low, hight int64) {
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
	var ltrim, rtrim func(idx int64, dir int, root *hNode[T]) *hNode[T]
	ltrim = func(idx int64, dir int, root *hNode[T]) *hNode[T] {
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

	var lgroup *hNode[T]
	if root.Children[L] != nil {
		lgroup = ltrim(idx, L, root.Children[L])
	}

	rtrim = func(idx int64, dir int, root *hNode[T]) *hNode[T] {
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

	var rgroup *hNode[T]
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

	// Left and right tree concat
	rsize := getSize(rgroup)
	lsize := getSize(lgroup)
	if lsize > rsize {
		tree.mergeGroups(root, lgroup, rgroup, rsize, R)
	} else {
		tree.mergeGroups(root, rgroup, lgroup, lsize, L)
	}
}

// Trim retain the value of the range . [low high]
func (tree *Tree[T]) Trim(low, high T) {
	// root := tree.getRoot()

	if tree.compare(low, high) > 0 {
		panic(errLowerGtHigh)
	}

	const L = 0
	const R = 1

	root := tree.getRangeRoot(low, high)

	var ltrim func(root *hNode[T]) *hNode[T]
	ltrim = func(root *hNode[T]) *hNode[T] {
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

	var rtrim func(root *hNode[T]) *hNode[T]
	rtrim = func(root *hNode[T]) *hNode[T] {
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

// TrimByIndex retain the value of the index range . [low high]
func (tree *Tree[T]) TrimByIndex(low, high int64) {
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

	var ltrim func(idx int64, root *hNode[T]) *hNode[T]
	ltrim = func(idx int64, root *hNode[T]) *hNode[T] {
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

	var rtrim func(idx int64, root *hNode[T]) *hNode[T]
	rtrim = func(idx int64, root *hNode[T]) *hNode[T] {
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

// Split Contain Split  Original tree not contain Key. return  the splited tree
func (tree *Tree[T]) Split(key T) *Tree[T] {
	root := tree.getRoot()
	if root == nil {
		return nil
	}

	const L = 0
	const R = 1

	cur := root
	// 寻找左右根
	var lroot, rroot *hNode[T]

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

	var traverse func(cur *hNode[T], lroot, rroot *hNode[T])
	traverse = func(cur, lroot, rroot *hNode[T]) {
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

// SplitContain  Original tree contain Key. return  the splited tree
func (tree *Tree[T]) SplitContain(key T) *Tree[T] {
	root := tree.getRoot()
	if root == nil {
		return nil
	}

	const L = 0
	const R = 1

	cur := root
	// 寻找左右根
	var lroot, rroot *hNode[T]

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

	var traverse func(cur *hNode[T], lroot, rroot *hNode[T])
	traverse = func(cur, lroot, rroot *hNode[T]) {
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

// Clear clear all node.
func (tree *Tree[T]) Clear() {
	tree.root.Children[0] = nil
}
