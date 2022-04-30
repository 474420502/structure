package treelist

import (
	"fmt"

	"github.com/474420502/structure/compare"
)

// Slice the KeyValue
type Slice[T any] struct {
	Key   T
	Value interface{}
}

// String show the string of keyvalue
func (s *Slice[T]) String() string {
	return fmt.Sprintf("{%v:%v}", s.Key, s.Value)
}

type treeNode[T any] struct {
	Parent   *treeNode[T]
	Children [2]*treeNode[T]
	Direct   [2]*treeNode[T]

	Size int64

	Slice[T]
}

func (n *treeNode[T]) String() string {
	return n.Slice.String()
}

// Tree the struct of treelist
type Tree[T any] struct {
	root    *treeNode[T]
	compare compare.Compare[T]

	// rcount int
}

// New create a object of tree
func New[T any](comp compare.Compare[T]) *Tree[T] {
	return &Tree[T]{compare: comp, root: &treeNode[T]{}}
}

// func (tree *Tree[T]) SetCompare(comp compare.Compare) {
// 	tree.compare = comp
// }

// Iterator Return the Iterator of tree. like list or skiplist
func (tree *Tree[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{tree: tree}
}

// IteratorRange Return the Iterator of tree. like list or skiplist.
//
// the struct can set range.
func (tree *Tree[T]) IteratorRange() *IteratorRange[T] {
	return &IteratorRange[T]{tree: tree}
}

// Size return the size of treelist
func (tree *Tree[T]) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.Size
	}
	return 0
}

// Get Get Value from key.
func (tree *Tree[T]) Get(key T) (interface{}, bool) {
	if cur := tree.getNode(key); cur != nil {
		return cur.Value, true
	}
	return nil, false
}

// PutDuplicate put, when key duplicate with call do. don,t change the key of `exists`, will break the tree of blance
// 				if duplicate, will return true.
func (tree *Tree[T]) PutDuplicate(key T, value interface{}, do func(exists *Slice[T])) bool {
	const L = 0
	const R = 1

	// if key == nil {
	// 	panic(fmt.Errorf("key must not be nil"))
	// }

	cur := tree.getRoot()
	if cur == nil {
		node := &treeNode[T]{Slice: Slice[T]{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return false
	}

	var left *treeNode[T] = nil
	var right *treeNode[T] = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &treeNode[T]{Parent: cur, Slice: Slice[T]{Key: key, Value: value}, Size: 1}
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
				node := &treeNode[T]{Parent: cur, Slice: Slice[T]{Key: key, Value: value}, Size: 1}
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
func (tree *Tree[T]) Set(key T, value interface{}) bool {
	const L = 0
	const R = 1

	// if key == nil {
	// 	panic(fmt.Errorf("key must not be nil"))
	// }

	cur := tree.getRoot()
	if cur == nil {

		node := &treeNode[T]{Slice: Slice[T]{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return false
	}

	var left *treeNode[T] = nil
	var right *treeNode[T] = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &treeNode[T]{Parent: cur, Slice: Slice[T]{Key: key, Value: value}, Size: 1}
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
				node := &treeNode[T]{Parent: cur, Slice: Slice[T]{Key: key, Value: value}, Size: 1}
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
func (tree *Tree[T]) Put(key T, value interface{}) bool {
	const L = 0
	const R = 1

	// if key == nil {
	// 	panic(fmt.Errorf("key must not be nil"))
	// }

	cur := tree.getRoot()
	if cur == nil {
		node := &treeNode[T]{Slice: Slice[T]{Key: key, Value: value}, Size: 1, Parent: tree.root}
		tree.root.Children[0] = node
		tree.root.Direct[L] = node
		tree.root.Direct[R] = node
		return true
	}

	var left *treeNode[T] = nil
	var right *treeNode[T] = nil

	for {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:

			right = cur
			if cur.Children[L] != nil {
				cur = cur.Children[L]
			} else {

				node := &treeNode[T]{Parent: cur, Slice: Slice[T]{Key: key, Value: value}, Size: 1}
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
				node := &treeNode[T]{Parent: cur, Slice: Slice[T]{Key: key, Value: value}, Size: 1}
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
func (tree *Tree[T]) Index(i int64) *Slice[T] {
	node := tree.index(i)
	return &node.Slice
}

// IndexOf Get the Index of key in the Treelist(Order)
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

// Traverse the traversal method defaults to LDR. from smallest to largest.
func (tree *Tree[T]) Traverse(every func(s *Slice[T]) bool) {
	root := tree.getRoot()
	if root == nil {
		return
	}

	var traverasl func(cur *treeNode[T]) bool
	traverasl = func(cur *treeNode[T]) bool {
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
func (tree *Tree[T]) Slices() []Slice[T] {
	var mszie int64
	root := tree.getRoot()
	if root != nil {
		mszie = root.Size
	}
	result := make([]Slice[T], 0, mszie)
	tree.Traverse(func(s *Slice[T]) bool {
		result = append(result, *s)
		return true
	})
	return result
}

// Remove remove key and return value that be removed. if not exists, return nil
func (tree *Tree[T]) Remove(key T) *Slice[T] {
	if cur := tree.getNode(key); cur != nil {
		return tree.removeNode(cur)
	}
	return nil
}

// RemoveIndex remove key value by index and return value that be removed
func (tree *Tree[T]) RemoveIndex(index int64) *Slice[T] {
	if cur := tree.index(index); cur != nil {
		return tree.removeNode(cur)
	}
	return nil
}

// Head returns the head of the ordered data of tree
func (tree *Tree[T]) Head() *Slice[T] {
	h := tree.root.Direct[0]
	if h != nil {
		return &h.Slice
	}
	return nil
}

// RemoveHead remove the head of the ordered data of tree. similar to the pop function of heap
func (tree *Tree[T]) RemoveHead() *Slice[T] {
	if tree.getRoot() != nil {
		return tree.removeNode(tree.root.Direct[0])
	}
	return nil
}

// Tail returns the tail of the ordered data of tree
func (tree *Tree[T]) Tail() *Slice[T] {
	t := tree.root.Direct[1]
	if t != nil {
		return &t.Slice
	}
	return nil
}

// RemoveTail remove the tail of the ordered data of tree.
func (tree *Tree[T]) RemoveTail() *Slice[T] {
	if tree.getRoot() != nil {
		return tree.removeNode(tree.root.Direct[1])
	}
	return nil
}

// RemoveRange remove keys values by range. [low, high]
func (tree *Tree[T]) RemoveRange(low, hight T) bool {

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

	var ltrim, rtrim func(*treeNode[T]) *treeNode[T]
	var dleft *treeNode[T]
	ltrim = func(root *treeNode[T]) *treeNode[T] {
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

	var lgroup *treeNode[T]
	if root.Children[L] != nil {
		lgroup = ltrim(root.Children[L])
	} else {
		dleft = root.Direct[L]
	}

	var dright *treeNode[T]
	rtrim = func(root *treeNode[T]) *treeNode[T] {
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

	var rgroup *treeNode[T]
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
	// log.Println(low, hight, "low:", tree.index(low), "hight:", tree.index(hight), "root:", root)
	var ltrim, rtrim func(idx int64, dir int, root *treeNode[T]) *treeNode[T]
	var dleft *treeNode[T]
	ltrim = func(idx int64, dir int, root *treeNode[T]) *treeNode[T] {
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

	var lgroup *treeNode[T]
	if root.Children[L] != nil {
		lgroup = ltrim(idx, L, root.Children[L])
	} else {
		dleft = root.Direct[L]
	}

	var dright *treeNode[T]
	rtrim = func(idx int64, dir int, root *treeNode[T]) *treeNode[T] {
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

	var rgroup *treeNode[T]
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
func (tree *Tree[T]) Clear() {
	tree.root.Children[0] = nil
	tree.root.Direct[0] = nil
	tree.root.Direct[1] = nil
}

// Trim retain the value of the range . [low high]
func (tree *Tree[T]) Trim(low, hight T) {

	if tree.compare(low, hight) > 0 {
		panic(errLowerGtHigh)
	}

	const L = 0
	const R = 1

	root := tree.getRangeRoot(low, hight)

	var ltrim func(root *treeNode[T]) *treeNode[T]
	ltrim = func(root *treeNode[T]) *treeNode[T] {
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

	var rtrim func(root *treeNode[T]) *treeNode[T]
	rtrim = func(root *treeNode[T]) *treeNode[T] {
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
func (tree *Tree[T]) TrimByIndex(low, hight int64) {

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

	var ltrim func(idx int64, root *treeNode[T]) *treeNode[T]
	ltrim = func(idx int64, root *treeNode[T]) *treeNode[T] {
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

	var rtrim func(idx int64, root *treeNode[T]) *treeNode[T]
	rtrim = func(idx int64, root *treeNode[T]) *treeNode[T] {
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
func (tree *Tree[T]) Intersection(other *Tree[T]) *Tree[T] {
	// This method is a bit stupid.  There is time to prepare for refactoring
	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	result := New(tree.compare)
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
func (tree *Tree[T]) UnionSets(other *Tree[T]) *Tree[T] {

	// There is time to prepare for refactoring

	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	result := New(tree.compare)
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
func (tree *Tree[T]) DifferenceSets(other *Tree[T]) *Tree[T] {
	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	result := New(tree.compare)
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
