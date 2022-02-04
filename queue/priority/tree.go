package treequeue

import (
	"fmt"

	"github.com/474420502/structure/compare"
)

type qNode[T any] struct {
	Parent   *qNode[T]
	Children [2]*qNode[T]

	Size int64

	Slice[T]
}

type Queue[T any] struct {
	root    *qNode[T]
	compare compare.Compare[T]
}

func New[T any](comp compare.Compare[T]) *Queue[T] {
	return &Queue[T]{compare: comp, root: &qNode[T]{}}
}

func (tree *Queue[T]) String() string {
	return fmt.Sprintf("%v", tree.Values())
}

func (tree *Queue[T]) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.Size
	}
	return 0
}

func (tree *Queue[T]) Head() *Slice[T] {
	const L = 0

	root := tree.getRoot()
	if root == nil {
		return nil
	}
	for root.Children[L] != nil {
		root = root.Children[L]
	}
	var s Slice[T] = root.Slice
	return &s
}

func (tree *Queue[T]) RemoveHead() *Slice[T] {
	const L = 0

	root := tree.getRoot()
	if root == nil {
		return nil
	}

	for root.Children[L] != nil {
		root = root.Children[L]
	}

	return tree.remove(root)
}

func (tree *Queue[T]) Tail() *Slice[T] {
	const R = 1

	root := tree.getRoot()
	if root == nil {
		return nil
	}
	for root.Children[R] != nil {
		root = root.Children[R]
	}
	var s Slice[T] = root.Slice
	return &s
}

func (tree *Queue[T]) RemoveTail() *Slice[T] {
	const R = 1

	root := tree.getRoot()
	if root == nil {
		return nil
	}
	for root.Children[R] != nil {
		root = root.Children[R]
	}

	return tree.remove(root)
}

// Get 按Key获取Value, 如果Key值相等, 返回最先入队的值
func (tree *Queue[T]) Get(key T) *Slice[T] {
	if cur := tree.getNode(key); cur != nil {
		return &cur.Slice
	}
	return nil
}

// Gets 按Key获取Value, 如果Key值相等, 返回所有相当值的Value. 顺序也按先后
func (tree *Queue[T]) Gets(key T) (result []*Slice[T]) {

	for _, node := range tree.getNodes(key) {
		result = append(result, &node.Slice)
	}

	return
}

// Put 插入数据. (队列不存在去重)
//
// put the data to queue.
func (tree *Queue[T]) Put(key T, value interface{}) {

	cur := tree.getRoot()
	if cur == nil {
		node := &qNode[T]{Size: 1, Parent: tree.root}
		node.key = key
		node.value = value
		tree.root.Children[0] = node
		return
	}

	const L = 0
	const R = 1

	for {
		c := tree.compare(key, cur.key)

		if c < 0 {

			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {
				node := &qNode[T]{Parent: cur, Size: 1}
				node.key = key
				node.value = value
				cur.Children[L] = node
				tree.fixPut(cur)
				return
			}

		} else {

			if cur.Children[R] != nil {
				cur = cur.Children[R]
			} else {
				node := &qNode[T]{Parent: cur, Size: 1}
				node.key = key
				node.value = value
				cur.Children[R] = node
				tree.fixPut(cur)
				return
			}

		}

	}

}

func (tree *Queue[T]) Index(i int64) *Slice[T] {
	node := tree.index(i)
	return &node.Slice
}

func (tree *Queue[T]) IndexOf(key T) int64 {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	if cur == nil {
		return -1
	}

	var offset int64 = getSize(cur.Children[L])
	for {
		c := tree.compare(key, cur.key)
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
			cur = cur.Children[L]
			if cur == nil {
				return offset
			}
			offset -= getSize(cur.Children[R]) + 1

			for {
				c = tree.compare(key, cur.key)
				if c == 0 {
					if cur.Children[L] == nil {
						return offset
					}
					cur = cur.Children[L]
					offset -= getSize(cur.Children[R]) + 1
				} else {
					if cur.Children[R] == nil {
						return offset + 1
					}
					cur = cur.Children[R]
					offset += getSize(cur.Children[L]) + 1
				}
			}

		}
	}

}

// Traverse 遍历的方法 默认是LDR 从小到大 Compare 为 l < r
func (tree *Queue[T]) Traverse(every func(s *Slice[T]) bool) {
	root := tree.getRoot()
	if root == nil {
		return
	}

	var traverasl func(cur *qNode[T]) bool
	traverasl = func(cur *qNode[T]) bool {
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

func (tree *Queue[T]) Slices() []*Slice[T] {
	var mszie int64
	root := tree.getRoot()
	if root != nil {
		mszie = root.Size
	}
	result := make([]*Slice[T], 0, mszie)
	tree.Traverse(func(s *Slice[T]) bool {
		result = append(result, s)
		return true
	})
	return result
}

func (tree *Queue[T]) Values() []interface{} {
	var mszie int64
	root := tree.getRoot()
	if root != nil {
		mszie = root.Size
	}
	result := make([]interface{}, 0, mszie)
	tree.Traverse(func(s *Slice[T]) bool {
		result = append(result, s.value)
		return true
	})
	return result
}

// Remove 按Key删除一个数据, 如果存在Key相同的情况下
func (tree *Queue[T]) Remove(key T) *Slice[T] {
	if cur := tree.getNode(key); cur != nil {
		return tree.remove(cur)
	}
	return nil
}

func (tree *Queue[T]) RemoveIndex(index int64) *Slice[T] {
	const L = 0
	const R = 1

	if cur := tree.index(index); cur != nil {

		if cur.Size == 1 {
			parent := cur.Parent
			parent.Children[getRelationship(cur)] = nil
			tree.fixRemoveSize(parent)
			return &cur.Slice
		}

		lsize, rsize := getChildrenSize(cur)
		if lsize > rsize {
			prev := cur.Children[L]
			for prev.Children[R] != nil {
				prev = prev.Children[R]
			}

			s := cur.Slice
			cur.Slice = prev.Slice

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

			return &s

		} else {

			next := cur.Children[R]
			for next.Children[L] != nil {
				next = next.Children[L]
			}

			s := cur.Slice
			cur.Slice = next.Slice

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
			return &s
		}
	}

	return nil
}

// RemoveRange 删除区间. [low:hight]
func (tree *Queue[T]) RemoveRange(low, hight T) {

	const L = 0
	const R = 1

	c := tree.compare(low, hight)
	if c > 0 {
		panic("key2 must greater than key1 or equal to")
	}
	// else if c == 0 {
	// 	tree.Remove(low)
	// 	return
	// }

	root := tree.getRangeRoot(low, hight)
	if root == nil {
		return
	}

	var ltrim, rtrim func(*qNode[T]) *qNode[T]
	ltrim = func(root *qNode[T]) *qNode[T] {
		if root == nil {
			return nil
		}
		c = tree.compare(low, root.Key())
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

	var lgroup *qNode[T]
	if root.Children[L] != nil {
		lgroup = ltrim(root.Children[L])
	}

	rtrim = func(root *qNode[T]) *qNode[T] {
		if root == nil {
			return nil
		}
		c = tree.compare(hight, root.Key())
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

	var rgroup *qNode[T]
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
func (tree *Queue[T]) RemoveRangeByIndex(low, hight int64) {

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
	var ltrim, rtrim func(idx int64, dir int, root *qNode[T]) *qNode[T]
	ltrim = func(idx int64, dir int, root *qNode[T]) *qNode[T] {
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

	var lgroup *qNode[T]
	if root.Children[L] != nil {
		lgroup = ltrim(idx, L, root.Children[L])
	}

	rtrim = func(idx int64, dir int, root *qNode[T]) *qNode[T] {
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

	var rgroup *qNode[T]
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

// Extract 提取区间的数据. 区间外数据删除(与RemoveRange相反) range [low:hight]
func (tree *Queue[T]) Extract(low, hight T) {
	// root := tree.getRoot()

	if tree.compare(low, hight) > 0 {
		panic(errLowerGtHigh)
	}

	const L = 0
	const R = 1

	root := tree.getRangeRoot(low, hight)

	var ltrim func(root *qNode[T]) *qNode[T]
	ltrim = func(root *qNode[T]) *qNode[T] {
		if root == nil {
			return nil
		}
		c := tree.compare(low, root.Key())
		if c > 0 {
			return ltrim(root.Children[R])
		} else { //
			child := ltrim(root.Children[L])
			root.Children[L] = child
			if child != nil {
				child.Parent = root
			}
			root.Size = getChildrenSumSize(root) + 1
			return root

		}
	}

	ltrim(root)

	var rtrim func(root *qNode[T]) *qNode[T]
	rtrim = func(root *qNode[T]) *qNode[T] {
		if root == nil {
			return nil
		}
		c := tree.compare(hight, root.Key())
		if c < 0 {
			return rtrim(root.Children[L])
		} else { //  c >= 0
			child := rtrim(root.Children[R])
			root.Children[R] = child
			if child != nil {
				child.Parent = root
			}
			root.Size = getChildrenSumSize(root) + 1
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

// ExtractByIndex 保留区间(Extract类似, 范围用顺序索引) range [low:hight]
func (tree *Queue[T]) ExtractByIndex(low, hight int64) {

	if low > hight {
		panic(errLowerGtHigh)
	}

	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf(errOutOfIndex, low, hight))
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

	var ltrim func(idx int64, root *qNode[T]) *qNode[T]
	ltrim = func(idx int64, root *qNode[T]) *qNode[T] {
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

	var rtrim func(idx int64, root *qNode[T]) *qNode[T]
	rtrim = func(idx int64, root *qNode[T]) *qNode[T] {
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

		rhand := root
		for rhand.Children[R] != nil {
			rhand = rhand.Children[R]
		}

	}
}

func (tree *Queue[T]) Clear() {
	tree.root.Children[0] = nil
}
