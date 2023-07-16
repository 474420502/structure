package treequeue

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

func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	target, isExists := tree.put(tree.Center, 1, key)
	if isExists == 1 {
		return false
	}
	target.Value = value
	return true
}

func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool) {

	// cur := tree.get(key, tree.getRoot())
	cur := tree.getfirst(key, tree.Center, 1, nil)
	log.Println(tree.viewEx(cur))
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
			log.Panicf("%v size: %d index: %d", errOutOfIndex, tree.Size()+1, index)
		}
	}()

	cur := tree.Center
	var idx int = cur.Children[1].Children[0].getSize()

	if idx < 0 {
		return *tree.removeIndex(cur, 1, idx, tree.Size()+index)
	}
	return *tree.removeIndex(cur, 1, idx, index)
}

// RemoveRange remove keys values by range. [low, high]
func (tree *Tree[KEY, VALUE]) Trim(low, high KEY) {
	cur := tree.seekRangeRoot(tree.getRoot(), low, high)
	if cur == nil {
		tree.Clear()
		return
	}
	cur.Children[0] = tree.trimLow(cur.Children[0], low)
	cur.Children[1] = tree.trimHigh(cur.Children[1], high)
	cur.updateSize()
	tree.Center.Children[1] = cur
}

// TrimByIndex retain the value of the index range . [low high]
func (tree *Tree[KEY, VALUE]) TrimByIndex(low, high int) {

	root := tree.getRoot()
	if root == nil {
		return
	}
	size := root.Size
	if low < 0 {
		low = size + low
	}
	if high < 0 {
		high = size + high
	}
	if low > high {
		low, high = high, low
	}
	if high >= size {
		log.Panicf("high(index) %v >= size(%v)", high, size)
	}

	cur, idx := tree.seekTrimIndexRoot(root, root.Children[0].getSize(), low, high)
	// log.Println(cur.view(), idx)
	cur.Children[0] = tree.trimIndexLow(cur, 0, idx, low)
	cur.Children[1] = tree.trimIndexHigh(cur, 1, idx, high)
	cur.updateSize()
	tree.Center.Children[1] = cur
	// log.Println(cur.view())
}

// RemoveRange remove keys values by range. [low, high]
func (tree *Tree[KEY, VALUE]) RemoveRange(low, high KEY) {

	c := tree.Compare(low, high)
	if c < 0 {
		tree.Remove(low)
	} else if c == 0 {
		low, high = high, low
	}

	root := tree.getRoot()
	if root == nil {
		return
	}

	var lefts, rights []*Node[KEY, VALUE]

	tree.removeCollectLows(&lefts, root, low)
	tree.removeCollectHighs(&rights, root, high)

	left := tree.megreThreshold(lefts, 0, 1)
	right := tree.megreThreshold(rights, 0, 0)

	if left == nil {
		tree.Center.Children[1] = right
		return
	}

	tree.Center.Children[1] = left
	if right == nil {
		return
	}
	tree.leftMegreRight(tree.Center, right)
	// log.Println(tree.view())
}

// RemoveRangeByIndex 1.remove range [low:high]
// 2.low and hight that the range must contain a value that exists. eg: [low: high+1] [low-1: high].  [low-1: hight+1]. error: [min-1:min-2] or [max+1:max+2]
func (tree *Tree[KEY, VALUE]) RemoveRangeByIndex(low, high int) {
	var lefts, rights []*Node[KEY, VALUE]
	root := tree.getRoot()
	if root == nil {
		return
	}
	// log.Println(tree.view())
	tree.removeCollectIndexLows(&lefts, root, root.Children[0].getSize(), low)
	tree.removeCollectIndexHighs(&rights, root, root.Children[0].getSize(), high)
	// log.Println(lefts, rights)

	left := tree.megreThreshold(lefts, 0, 1)
	right := tree.megreThreshold(rights, 0, 0)

	if left == nil {
		tree.Center.Children[1] = right
		return
	}

	tree.Center.Children[1] = left
	if right == nil {
		return
	}
	tree.leftMegreRight(tree.Center, right)
}

// Split  original tree contain Key. return  the splited tree.
// eg: 1.[1 4 5 7] -> Split(5) [1 4 5] [7]; 2.[1 4 5 7] -> Split(3) [1] [4 5 7]
func (tree *Tree[KEY, VALUE]) Split(key KEY) *Tree[KEY, VALUE] {
	root := tree.getRoot()
	var lefts, rights []*Node[KEY, VALUE]
	tree.split(&lefts, &rights, root, key)

	left := tree.megreThreshold(lefts, 0, 1)
	right := tree.megreThreshold(rights, 0, 0)

	other := &Tree[KEY, VALUE]{
		Center:  &Node[KEY, VALUE]{Size: 0},
		Compare: tree.Compare,
	}

	other.Center.Children[0] = other.Center
	other.Center.Children[1] = right
	if right != nil {
		other.rebalance(other.Center, 1)
	}

	tree.Center.Children[1] = left
	if left != nil {
		tree.rebalance(tree.Center, 1)
	}

	// log.Println(tree.view(), other.view())

	return other
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
