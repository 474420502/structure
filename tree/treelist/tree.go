package treelist

import (
	"fmt"

	"github.com/474420502/structure/compare"
)

// Slice the KeyValue
type Slice[KEY any, VALUE any] struct {
	Key   KEY
	Value VALUE
}

// String show the string of keyvalue
func (s *Slice[KEY, VALUE]) String() string {
	return fmt.Sprintf("{%v:%v}", s.Key, s.Value)
}

type treeNode[KEY any, VALUE any] struct {
	Parent   *treeNode[KEY, VALUE]
	Children [2]*treeNode[KEY, VALUE]
	Direct   [2]*treeNode[KEY, VALUE]

	Size int64

	Slice[KEY, VALUE]
}

func (n *treeNode[KEY, VALUE]) String() string {
	return n.Slice.String()
}

// Tree the struct of treelist
type Tree[KEY any, VALUE any] struct {
	root    *treeNode[KEY, VALUE]
	compare compare.Compare[KEY]

	// rcount int
}

// New create a object of tree
func New[KEY any, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE] {
	return &Tree[KEY, VALUE]{compare: comp, root: &treeNode[KEY, VALUE]{}}
}

// func (tree *Tree[KEY,VALUE]) SetCompare(comp compare.Compare) {
// 	tree.compare = comp
// }

// Iterator Return the Iterator of tree. like list or skiplist
func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE] {
	return &Iterator[KEY, VALUE]{tree: tree}
}

// IteratorRange Return the Iterator of tree. like list or skiplist.
//
// the struct can set range.
func (tree *Tree[KEY, VALUE]) IteratorRange() *IteratorRange[KEY, VALUE] {
	return &IteratorRange[KEY, VALUE]{tree: tree}
}

// Size return the size of treelist
func (tree *Tree[KEY, VALUE]) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.Size
	}
	return 0
}

// Get Get Value from key.
func (tree *Tree[KEY, VALUE]) Get(key KEY) (interface{}, bool) {
	if cur := tree.getNode(key); cur != nil {
		return cur.Value, true
	}
	return nil, false
}

// PutDuplicate put, when key duplicate with call do. don,t change the key of `exists`, will break the tree of blance
// 				if duplicate, will return true.
func (tree *Tree[KEY, VALUE]) PutDuplicate(key KEY, value VALUE, do func(exists *Slice[KEY, VALUE])) bool {
	const L = 0
	const R = 1

	// if key == nil {
	// 	panic(fmt.Errorf("key must not be nil"))
	// }

	cur := tree.getRoot()
	if cur == nil {
		node := &treeNode[KEY, VALUE]{Slice: Slice[KEY, VALUE]{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return false
	}

	var left *treeNode[KEY, VALUE] = nil
	var right *treeNode[KEY, VALUE] = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &treeNode[KEY, VALUE]{Parent: cur, Slice: Slice[KEY, VALUE]{Key: key, Value: value}, Size: 1}
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
				node := &treeNode[KEY, VALUE]{Parent: cur, Slice: Slice[KEY, VALUE]{Key: key, Value: value}, Size: 1}
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
func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool {
	const L = 0
	const R = 1

	// if key == nil {
	// 	panic(fmt.Errorf("key must not be nil"))
	// }

	cur := tree.getRoot()
	if cur == nil {

		node := &treeNode[KEY, VALUE]{Slice: Slice[KEY, VALUE]{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return false
	}

	var left *treeNode[KEY, VALUE] = nil
	var right *treeNode[KEY, VALUE] = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &treeNode[KEY, VALUE]{Parent: cur, Slice: Slice[KEY, VALUE]{Key: key, Value: value}, Size: 1}
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
				node := &treeNode[KEY, VALUE]{Parent: cur, Slice: Slice[KEY, VALUE]{Key: key, Value: value}, Size: 1}
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
func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	const L = 0
	const R = 1

	// if key == nil {
	// 	panic(fmt.Errorf("key must not be nil"))
	// }

	cur := tree.getRoot()
	if cur == nil {
		node := &treeNode[KEY, VALUE]{Slice: Slice[KEY, VALUE]{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return true
	}

	var left *treeNode[KEY, VALUE] = nil
	var right *treeNode[KEY, VALUE] = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &treeNode[KEY, VALUE]{Parent: cur, Slice: Slice[KEY, VALUE]{Key: key, Value: value}, Size: 1}
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
				node := &treeNode[KEY, VALUE]{Parent: cur, Slice: Slice[KEY, VALUE]{Key: key, Value: value}, Size: 1}
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
func (tree *Tree[KEY, VALUE]) Index(i int64) *Slice[KEY, VALUE] {
	node := tree.index(i)
	return &node.Slice
}

// IndexOf Get the Index of key in the Treelist(Order)
func (tree *Tree[KEY, VALUE]) IndexOf(key KEY) int64 {
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
func (tree *Tree[KEY, VALUE]) Traverse(every func(s *Slice[KEY, VALUE]) bool) {
	root := tree.getRoot()
	if root == nil {
		return
	}

	var traverasl func(cur *treeNode[KEY, VALUE]) bool
	traverasl = func(cur *treeNode[KEY, VALUE]) bool {
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

// Slices  return all slice. from smallest to largest.
func (tree *Tree[KEY, VALUE]) Slices() []Slice[KEY, VALUE] {
	var mszie int64
	root := tree.getRoot()
	if root != nil {
		mszie = root.Size
	}
	result := make([]Slice[KEY, VALUE], 0, mszie)
	tree.Traverse(func(s *Slice[KEY, VALUE]) bool {
		result = append(result, *s)
		return true
	})
	return result
}

// Remove remove key and return value that be removed. if not exists, return nil
func (tree *Tree[KEY, VALUE]) Remove(key KEY) *Slice[KEY, VALUE] {
	if cur := tree.getNode(key); cur != nil {
		return tree.removeNode(cur)
	}
	return nil
}

// RemoveIndex remove key value by index and return value that be removed
func (tree *Tree[KEY, VALUE]) RemoveIndex(index int64) *Slice[KEY, VALUE] {
	if cur := tree.index(index); cur != nil {
		return tree.removeNode(cur)
	}
	return nil
}

// Head returns the head of the ordered data of tree
func (tree *Tree[KEY, VALUE]) Head() *Slice[KEY, VALUE] {
	h := tree.root.Direct[0]
	if h != nil {
		return &h.Slice
	}
	return nil
}

// RemoveHead remove the head of the ordered data of tree. similar to the pop function of heap
func (tree *Tree[KEY, VALUE]) RemoveHead() *Slice[KEY, VALUE] {
	if tree.getRoot() != nil {
		return tree.removeNode(tree.root.Direct[0])
	}
	return nil
}

// Tail returns the tail of the ordered data of tree
func (tree *Tree[KEY, VALUE]) Tail() *Slice[KEY, VALUE] {
	t := tree.root.Direct[1]
	if t != nil {
		return &t.Slice
	}
	return nil
}

// RemoveTail remove the tail of the ordered data of tree.
func (tree *Tree[KEY, VALUE]) RemoveTail() *Slice[KEY, VALUE] {
	if tree.getRoot() != nil {
		return tree.removeNode(tree.root.Direct[1])
	}
	return nil
}

// RemoveRange remove keys values by range. [low, high]
func (tree *Tree[KEY, VALUE]) RemoveRange(low, hight KEY) bool {

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

	var ltrim, rtrim func(*treeNode[KEY, VALUE]) *treeNode[KEY, VALUE]
	var dleft *treeNode[KEY, VALUE]
	ltrim = func(root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE] {
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

	var lgroup *treeNode[KEY, VALUE]
	if root.Children[L] != nil {
		lgroup = ltrim(root.Children[L])
	} else {
		dleft = root.Direct[L]
	}

	var dright *treeNode[KEY, VALUE]
	rtrim = func(root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE] {
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

	var rgroup *treeNode[KEY, VALUE]
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

// RemoveRangeByIndex 1.range [low:hight] 2.low hight 必须包含存在的值.[low: hight+1] [low-1: hight].  [low-1: hight+1]. error: [low-1:low-2] or [hight+1:hight+2]
func (tree *Tree[KEY, VALUE]) RemoveRangeByIndex(low, hight int64) {

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
	var ltrim, rtrim func(idx int64, dir int, root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE]
	var dleft *treeNode[KEY, VALUE]
	ltrim = func(idx int64, dir int, root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE] {
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

	var lgroup *treeNode[KEY, VALUE]
	if root.Children[L] != nil {
		lgroup = ltrim(idx, L, root.Children[L])
	} else {
		dleft = root.Direct[L]
	}

	var dright *treeNode[KEY, VALUE]
	rtrim = func(idx int64, dir int, root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE] {
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

	var rgroup *treeNode[KEY, VALUE]
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
func (tree *Tree[KEY, VALUE]) Clear() {
	tree.root.Children[0] = nil
	tree.root.Direct[0] = nil
	tree.root.Direct[1] = nil
}

// Trim retain the value of the range . [low high]
func (tree *Tree[KEY, VALUE]) Trim(low, hight KEY) {

	if tree.compare(low, hight) > 0 {
		panic(errLowerGtHigh)
	}

	const L = 0
	const R = 1

	root := tree.getRangeRoot(low, hight)

	var ltrim func(root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE]
	ltrim = func(root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE] {
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

	var rtrim func(root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE]
	rtrim = func(root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE] {
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
func (tree *Tree[KEY, VALUE]) TrimByIndex(low, hight int64) {

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

	var ltrim func(idx int64, root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE]
	ltrim = func(idx int64, root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE] {
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

	var rtrim func(idx int64, root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE]
	rtrim = func(idx int64, root *treeNode[KEY, VALUE]) *treeNode[KEY, VALUE] {
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
func (tree *Tree[KEY, VALUE]) Intersection(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE] {
	// This method is a bit stupid.  There is time to prepare for refactoring
	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	result := New[KEY, VALUE](tree.compare)
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
func (tree *Tree[KEY, VALUE]) UnionSets(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE] {

	// There is time to prepare for refactoring

	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	result := New[KEY, VALUE](tree.compare)
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
func (tree *Tree[KEY, VALUE]) DifferenceSets(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE] {
	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	result := New[KEY, VALUE](tree.compare)
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
