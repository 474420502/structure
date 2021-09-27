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

func copybytes(key []byte) []byte {
	var buf []byte = make([]byte, len(key))
	copy(buf, key)
	return buf
}

func (s *Slice) String() string {
	return string(s.Key)
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

func (tree *Tree) Set(key []byte, value interface{}) bool {
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
			cur.Slice.Value = value
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

func (tree *Tree) RemoveIndex(index int64) *Slice {
	if cur := tree.index(index); cur != nil {
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

// RemoveRange
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
func (tree *Tree) RemoveRangeByIndex(low, hight int64) {

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

// Intersection 交集
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

// UnionSets 并集
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

// DifferenceSets 差集
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
