package treelist

import (
	"fmt"

	"github.com/474420502/structure/compare"
)

func init() {

}

type Slice struct {
	Key   []byte
	Value interface{}
}

// func copybytes(key []byte) []byte {
// 	var buf []byte = make([]byte, len(key))
// 	copy(buf, key)
// 	return buf
// }

// String show the string of keyvalue
func (s *Slice) String() string {
	return fmt.Sprintf("{%v:%v}", string(s.Key), s.Value)
}

type treeNode struct {
	Parent   *treeNode
	Children [2]*treeNode
	Direct   [2]*treeNode

	Size int64

	Slice
}

func (n *treeNode) String() string {
	return string(n.Key)
}

// Tree the struct of treelist
type Tree struct {
	root    *treeNode
	compare compare.Compare[[]byte]

	// rcount int
}

func compareBytes(s1, s2 []byte) int {
	switch {
	case len(s1) > len(s2):
		for i := 0; i < len(s2); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 1
	case len(s1) < len(s2):
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return -1
	default:
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 0
	}
}

func compareBytesLen(s1, s2 []byte) int {
	switch {
	case len(s1) > len(s2):
		return 1
	case len(s1) < len(s2):
		return -1
	default:
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 0
	}
}

// New create a object of tree
func New() *Tree {
	return &Tree{compare: compareBytes, root: &treeNode{}}
}

func (tree *Tree) SetCompare(comp compare.Compare[[]byte]) {
	tree.compare = comp
}

// Iterator Return the Iterator of tree. similar to list or skiplist
func (tree *Tree) Iterator() *Iterator {
	return &Iterator{tree: tree}
}

// IteratorRange Return the Iterator of tree. similar to list or skiplist.
//
// the struct can set range.
func (tree *Tree) IteratorRange() *IteratorRange {
	return &IteratorRange{tree: tree}
}

// Size return the size of treelist
func (tree *Tree) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.Size
	}
	return 0
}

// Get Get Value from key.
func (tree *Tree) Get(key []byte) (interface{}, bool) {
	if cur := tree.getNode(key); cur != nil {
		return cur.Value, true
	}
	return nil, false
}

// PutDuplicate put, when key duplicate with call do. don,t change the key of `exists`, will break the tree of blance
// 				if duplicate, will return true.
func (tree *Tree) PutDuplicate(key []byte, value interface{}, do func(exists *Slice)) bool {
	const L = 0
	const R = 1

	if len(key) == 0 {
		panic(fmt.Errorf("key must not be nil"))
	}

	cur := tree.getRoot()
	if cur == nil {
		node := &treeNode{Slice: Slice{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return false
	}

	var left *treeNode = nil
	var right *treeNode = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &treeNode{Parent: cur, Slice: Slice{Key: key, Value: value}, Size: 1}
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
				node := &treeNode{Parent: cur, Slice: Slice{Key: key, Value: value}, Size: 1}
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

// Set Insert the key In treelist, if key exists, cover
func (tree *Tree) Set(key []byte, value interface{}) bool {
	const L = 0
	const R = 1

	if len(key) == 0 {
		panic(fmt.Errorf("key must not be nil"))
	}

	cur := tree.getRoot()
	if cur == nil {

		node := &treeNode{Slice: Slice{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return false
	}

	var left *treeNode = nil
	var right *treeNode = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &treeNode{Parent: cur, Slice: Slice{Key: key, Value: value}, Size: 1}
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
				node := &treeNode{Parent: cur, Slice: Slice{Key: key, Value: value}, Size: 1}
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
			cur.Slice.Key = key
			cur.Slice.Value = value
			return true
		}
	}

}

// Put Insert the key In treelist, if key exists, ignore
func (tree *Tree) Put(key []byte, value interface{}) bool {
	const L = 0
	const R = 1

	if len(key) == 0 {
		panic(fmt.Errorf("key must not be nil"))
	}

	cur := tree.getRoot()
	if cur == nil {
		node := &treeNode{Slice: Slice{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return true
	}

	var left *treeNode = nil
	var right *treeNode = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &treeNode{Parent: cur, Slice: Slice{Key: key, Value: value}, Size: 1}
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
				node := &treeNode{Parent: cur, Slice: Slice{Key: key, Value: value}, Size: 1}
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

// Index return the slice by index.
//
// like the index of array(order)
func (tree *Tree) Index(i int64) *Slice {
	node := tree.index(i)
	return &node.Slice
}

// IndexOf Get the Index of key in the Treelist(Order)
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

// Traverse the traversal method defaults to LDR. from smallest to largest.
func (tree *Tree) Traverse(every func(s *Slice) bool) {
	root := tree.getRoot()
	if root == nil {
		return
	}

	var traverasl func(cur *treeNode) bool
	traverasl = func(cur *treeNode) bool {
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

//  Slices  return all slice. from smallest to largest.
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

// Remove remove key and return value that be removed
func (tree *Tree) Remove(key []byte) *Slice {
	if cur := tree.getNode(key); cur != nil {
		return tree.removeNode(cur)
	}
	return nil
}

// RemoveIndex remove key value by index and return value that be removed
func (tree *Tree) RemoveIndex(index int64) *Slice {
	if cur := tree.index(index); cur != nil {
		return tree.removeNode(cur)
	}
	return nil
}

// Head returns the head of the ordered data of tree
func (tree *Tree) Head() *Slice {
	h := tree.root.Direct[0]
	if h != nil {
		return &h.Slice
	}
	return nil
}

// RemoveHead remove the head of the ordered data of tree. similar to the pop function of heap
func (tree *Tree) RemoveHead() *Slice {
	if tree.getRoot() != nil {
		return tree.removeNode(tree.root.Direct[0])
	}
	return nil
}

// Tail returns the tail of the ordered data of tree
func (tree *Tree) Tail() *Slice {
	t := tree.root.Direct[1]
	if t != nil {
		return &t.Slice
	}
	return nil
}

// RemoveTail remove the tail of the ordered data of tree.
func (tree *Tree) RemoveTail() *Slice {
	if tree.getRoot() != nil {
		return tree.removeNode(tree.root.Direct[1])
	}
	return nil
}

// RemoveRange remove keys values by range. [low, high]
func (tree *Tree) RemoveRange(low, hight []byte) bool {

	const L = 0
	const R = 1

	c := tree.compare(low, hight)
	if c > 0 {
		panic("key2 must greater than key1 or equal to")
	} else if c == 0 {
		return tree.Remove(low) != nil
	}

	root := tree.getRangeRoot(low, hight)
	if root == nil {
		return false
	}

	var ltrim, rtrim func(*treeNode) *treeNode
	var dleft *treeNode
	ltrim = func(root *treeNode) *treeNode {
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

	var lgroup *treeNode
	if root.Children[L] != nil {
		lgroup = ltrim(root.Children[L])
	} else {
		dleft = root.Direct[L]
	}

	var dright *treeNode
	rtrim = func(root *treeNode) *treeNode {
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

	var rgroup *treeNode
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
		return true
	}

	// 左右树　拼接
	rsize := getSize(rgroup)
	lsize := getSize(lgroup)
	if lsize > rsize {
		tree.mergeGroups(root, lgroup, rgroup, rsize, R)
	} else {
		tree.mergeGroups(root, rgroup, lgroup, lsize, L)
	}

	return true
}

// RemoveRangeByIndex 1.remove range [low:hight]
// 2.low and hight that the range must contain a value that exists. eg: [low: hight+1] [low-1: hight].  [low-1: hight+1]. error: [low-1:low-2] or [hight+1:hight+2]
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
	// log.Println(low, hight, "low:", tree.index(low), "hight:", tree.index(hight), "root:", root)
	var ltrim, rtrim func(idx int64, dir int, root *treeNode) *treeNode
	var dleft *treeNode
	ltrim = func(idx int64, dir int, root *treeNode) *treeNode {
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

	var lgroup *treeNode
	if root.Children[L] != nil {
		lgroup = ltrim(idx, L, root.Children[L])
	} else {
		dleft = root.Direct[L]
	}

	var dright *treeNode
	rtrim = func(idx int64, dir int, root *treeNode) *treeNode {
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

	var rgroup *treeNode
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

// Clear. Reset the treelist.
func (tree *Tree) Clear() {
	tree.root.Children[0] = nil
	tree.root.Direct[0] = nil
	tree.root.Direct[1] = nil
}

// Trim retain the value of the range . [low high]
func (tree *Tree) Trim(low, hight []byte) {

	if tree.compare(low, hight) > 0 {
		panic(errLowerGtHigh)
	}

	const L = 0
	const R = 1

	root := tree.getRangeRoot(low, hight)

	var ltrim func(root *treeNode) *treeNode
	ltrim = func(root *treeNode) *treeNode {
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

	var rtrim func(root *treeNode) *treeNode
	rtrim = func(root *treeNode) *treeNode {
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

// TrimByIndex retain the value of the index range . [low high]
func (tree *Tree) TrimByIndex(low, hight int64) {

	if low > hight {
		panic(errLowerGtHigh)
	}

	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf(errLowerGtHigh, low, hight))
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

	var ltrim func(idx int64, root *treeNode) *treeNode
	ltrim = func(idx int64, root *treeNode) *treeNode {
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

	var rtrim func(idx int64, root *treeNode) *treeNode
	rtrim = func(idx int64, root *treeNode) *treeNode {
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

// Intersection  tree intersection with other. [1 2 3] [2 3 4] -> [2 3].
func (tree *Tree) Intersection(other *Tree) *Tree {

	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	result := New()
	result.compare = tree.compare

	for head1 != nil && head2 != nil {
		c := tree.compare(head1.Key, head2.Key)
		// count++

		switch {
		case c < 0:
			head1 = head1.Direct[R]
		case c > 0:
			head2 = head2.Direct[R]
		default:
			result.Put(head1.Key, head1.Value)
			head1 = head1.Direct[R]
			head2 = head2.Direct[R]
		}
	}

	// log.Println("count:", count, tree.Size(), other.Size())
	return result
}

// UnionSets tree unionsets with other. [1 2 3] [2 3 4] -> [1 2 3 4].
func (tree *Tree) UnionSets(other *Tree) *Tree {
	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	result := New()
	result.compare = tree.compare

	for head1 != nil && head2 != nil {
		c := tree.compare(head1.Key, head2.Key)
		// count++
		switch {
		case c < 0:
			result.Put(head1.Key, head1.Value)
			head1 = head1.Direct[R]

		case c > 0:
			result.Put(head2.Key, head2.Value)
			head2 = head2.Direct[R]

		default:
			result.Put(head1.Key, head1.Value)
			head1 = head1.Direct[R]
			head2 = head2.Direct[R]
		}
	}

	for head1 != nil {
		result.Put(head1.Key, head1.Value)
		head1 = head1.Direct[R]
	}

	for head2 != nil {
		result.Put(head2.Key, head2.Value)
		head2 = head2.Direct[R]
	}

	return result
}

// DifferenceSets The set of elements after subtracting B from A
func (tree *Tree) DifferenceSets(other *Tree) *Tree {
	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	result := New()
	result.compare = tree.compare

	for head1 != nil && head2 != nil {
		c := tree.compare(head1.Key, head2.Key)
		// count++
		switch {
		case c < 0:
			result.Put(head1.Key, head1.Value)
			head1 = head1.Direct[R]

		case c > 0:
			head2 = head2.Direct[R]
		default:
			head1 = head1.Direct[R]
			head2 = head2.Direct[R]
		}
	}

	for head1 != nil {
		result.Put(head1.Key, head1.Value)
		head1 = head1.Direct[R]
	}

	return result
}
