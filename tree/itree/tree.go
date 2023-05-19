package itree

import (
	"log"

	"github.com/474420502/structure/compare"
)

type Tree[KEY, VALUE any] struct {
	Center  *Node[KEY, VALUE]
	Compare compare.Compare[KEY]
	// hight       int
	zero VALUE
	// rotateCount int
}

func New[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE] {

	tree := &Tree[KEY, VALUE]{
		Center:  &Node[KEY, VALUE]{Size: 0},
		Compare: Compare,
	}

	tree.Center.Children[0] = tree.Center
	return tree
}

func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool {
	target, isExists := tree.put(tree.Center, 1, key)
	target.Key = key
	target.Value = value

	return isExists != 1
}

func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	target, isExists := tree.put(tree.Center, 1, key)
	if isExists == 1 {
		return false
	}
	target.Value = value
	return true
}

func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	cur := tree.get(key, tree.getRoot())
	if cur == nil {
		return tree.zero, false
	}

	return cur.Value, true
}

func (tree *Tree[KEY, VALUE]) Index(idx int) VALUE {

	defer func() {
		if err := recover(); err != nil {
			log.Panicf("%v size: %d index: %d", errOutOfIndex, tree.Size(), idx)
		}
	}()

	if idx < 0 {
		return tree.index(tree.Size() + idx).Value
	}
	return tree.index(idx).Value
}

func (tree *Tree[KEY, VALUE]) IndexOf(key KEY) int {

	cur := tree.getRoot()
	if cur == nil {
		return -1
	}

	var offset int = cur.Children[0].getSize()
	for {
		cmp := tree.Compare(cur.Key, key)
		if cmp < 0 {
			return offset
		} else {
			if cmp == 0 {
				cur = cur.Children[0]
				if cur == nil {
					return -1
				}
				offset -= getSize(cur.Children[1]) + 1
			} else {
				cur = cur.Children[1]
				if cur == nil {
					return -1
				}
				offset += getSize(cur.Children[0]) + 1
			}

		}

	}
}

func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool) {
	target := tree.remove(key, tree.Center, 0, 1)
	if target != nil {
		return *target, true
	}
	return tree.zero, false
}

// RemoveIndex remove key value by index and return value that be removed
func (tree *Tree[KEY, VALUE]) RemoveIndex(index int) VALUE {
	defer func() {
		if err := recover(); err != nil {
			log.Panicf("%v size: %d index: %d", errOutOfIndex, tree.Size(), index)
		}
	}()

	cur := tree.getRoot()
	var idx int = cur.Children[0].getSize()

	if idx < 0 {
		return *tree.removeIndex(cur, 0, idx, tree.Size()+index)
	}
	return *tree.removeIndex(cur, 0, idx, index)
}

// RemoveRange remove keys values by range. [low, high]
func (tree *Tree[KEY, VALUE]) Trim(low, high KEY) {
	cur := tree.seekTrimRoot(tree.getRoot(), low, high)
	if cur == nil {
		tree.Clear()
		return
	}
	cur.Children[0] = tree.trimLow(cur, 0, low)
	cur.Children[1] = tree.trimHigh(cur, 1, high)
	cur.updateSize()
	tree.Center.Children[1] = cur
}

// TrimByIndex retain the value of the index range . [low high]
func (tree *Tree[KEY, VALUE]) TrimByIndex(low, high int) {
	root := tree.getRoot()
	if root == nil {
		return
	}

	cur := tree.seekTrimIndexRoot(root, root.Children[0].getSize(), low, high)
	if cur == nil {
		tree.Clear()
		return
	}
	log.Println(cur.view())
}

func (tree *Tree[KEY, VALUE]) seekTrimIndexRoot(cur *Node[KEY, VALUE], idx, low, high int) *Node[KEY, VALUE] {

	if idx > low && idx > high {
		cur = cur.Children[0]
		idx -= cur.Children[1].getSize() + 1
		return tree.seekTrimIndexRoot(cur, idx, low, high)
	} else if idx < low && idx < high {
		cur = cur.Children[1]
		idx += cur.Children[0].getSize() + 1
		return tree.seekTrimIndexRoot(cur, idx, low, high)
	}
	return cur
}

// RemoveRange remove keys values by range. [low, high]
func (tree *Tree[KEY, VALUE]) RemoveRange(low, high KEY) {

	const L = 0
	const R = 1

	c := tree.Compare(low, high)
	if c > 0 {
		panic("key2 must greater than key1 or equal to")
	} else if c == 0 {
		tree.Remove(low)
		return
	}

	// root := tree.getRangeRoot(low, high)
	// if root == nil {
	// 	return
	// }

	// var ltrim, rtrim func(*hNode[T]) *hNode[T]
	// ltrim = func(root *hNode[T]) *hNode[T] {
	// 	if root == nil {
	// 		return nil
	// 	}
	// 	c = tree.compare(low, root.Key)
	// 	if c > 0 {
	// 		child := ltrim(root.Children[R])
	// 		root.Children[R] = child
	// 		if child != nil {
	// 			child.Parent = root
	// 		}
	// 		root.Size = getChildrenSumSize(root) + 1
	// 		return root
	// 	} else {
	// 		return ltrim(root.Children[L])
	// 	}
	// }

	// var lgroup *hNode[T]
	// if root.Children[L] != nil {
	// 	lgroup = ltrim(root.Children[L])
	// }

	// rtrim = func(root *hNode[T]) *hNode[T] {
	// 	if root == nil {
	// 		return nil
	// 	}
	// 	c = tree.compare(high, root.Key)
	// 	if c < 0 {
	// 		child := rtrim(root.Children[L])
	// 		root.Children[L] = child
	// 		if child != nil {
	// 			child.Parent = root
	// 		}
	// 		root.Size = getChildrenSumSize(root) + 1
	// 		return root
	// 	} else {
	// 		return rtrim(root.Children[R])
	// 	}
	// }

	// var rgroup *hNode[T]
	// if root.Children[R] != nil {
	// 	rgroup = rtrim(root.Children[R])
	// }

	// if lgroup == nil && rgroup == nil {
	// 	rparent := root.Parent
	// 	size := root.Size
	// 	root.Parent.Children[getRelationship(root)] = nil
	// 	for rparent != tree.root {
	// 		rparent.Size -= size
	// 		rparent = rparent.Parent
	// 	}

	// 	return
	// }

	// // Left and right tree concat
	// rsize := getSize(rgroup)
	// lsize := getSize(lgroup)
	// if lsize > rsize {
	// 	tree.mergeGroups(root, lgroup, rgroup, rsize, R)
	// } else {
	// 	tree.mergeGroups(root, rgroup, lgroup, lsize, L)
	// }
}

func (tree *Tree[KEY, VALUE]) Clear() {
	tree.Center.Children[1] = nil
}

func (tree *Tree[KEY, VALUE]) Size() int {
	return getSize(tree.getRoot())
}

func (tree *Tree[KEY, VALUE]) Traverse(every func(KEY, VALUE) bool) {

	var traverse func(cur *Node[KEY, VALUE]) bool
	traverse = func(cur *Node[KEY, VALUE]) bool {
		if cur == nil {
			return true
		}
		if !traverse(cur.Children[0]) {
			return false
		}
		if !every(cur.Key, cur.Value) {
			return false
		}
		if !traverse(cur.Children[1]) {
			return false
		}
		return true
	}
	traverse(tree.getRoot())
}

func (tree *Tree[KEY, VALUE]) Values() []VALUE {
	if tree.Center.Children[1] == nil {
		return nil
	}
	result := make([]VALUE, 0, tree.Size())
	tree.Traverse(func(k KEY, v VALUE) bool {
		result = append(result, v)
		return true
	})
	return result
}

func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE] {
	return newIterator(tree)
}
