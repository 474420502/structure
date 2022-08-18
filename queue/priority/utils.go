package treequeue

import (
	"log"
)

var errOutOfIndex = "out of index"
var errLowerGtHigh = "low is behind high"

func (tree *Queue[T]) fixPutSize(cur *qNode[T]) {
	for cur != tree.root {
		cur.Size++
		cur = cur.Parent
	}
}

func (tree *Queue[T]) fixRemoveSize(cur *qNode[T]) {
	for cur != tree.root {
		cur.Size--
		cur = cur.Parent
	}
}

type heightLimitSize struct {
	rootsize   int64
	bottomsize int64
}

var rootSizeTable []*heightLimitSize = func() []*heightLimitSize {
	table := make([]*heightLimitSize, 64)
	for i := 2; i < 64; i++ {
		root2nsize := (int64(1) << i)
		bottomsize := root2nsize >> 1
		for x := 3; x < 64; x++ {
			rsize := root2nsize >> x
			if rsize == 0 {
				break
			}
			bottomsize -= rsize
		}
		table[i] = &heightLimitSize{
			rootsize:   root2nsize,
			bottomsize: bottomsize,
		}
	}
	return table
}()

func (tree *Queue[T]) fixPut(cur *qNode[T]) {

	cur.Size++
	if cur.Size == 3 {
		tree.fixPutSize(cur.Parent)
		return
	}

	var height int64 = 2
	var lsize, rsize int64
	var relations int = L
	var parent *qNode

	cur = cur.Parent

	for cur != tree.root {
		cur.Size++
		parent = cur.Parent

		limitsize := rootSizeTable[height]
		// (1<< height) -1 允许的最大size　超过证明高度超1, 并且有最少１size的空缺
		if cur.Size < limitsize.rootsize {

			lsize, rsize = getChildrenSize(cur)
			// 右就检测左边
			if rsize > lsize {
				if rsize-lsize >= limitsize.bottomsize {
					tree.sizeRrotate(cur)
					tree.fixPutSize(parent)
					return
				}
			} else {
				if lsize-rsize >= limitsize.bottomsize {
					tree.sizeLrotate(cur)
					tree.fixPutSize(parent)
					return
				}
			}
		}

		height++
		cur = parent
	}
}

func (tree *Queue[T]) sizeRrotate(cur *qNode[T]) *qNode[T] {

	llsize, lrsize := getChildrenSize(cur.Children[1])
	if llsize > lrsize {
		tree.rrotate(cur.Children[1])
	}
	return tree.lrotate(cur)
}

func (tree *Queue[T]) sizeLrotate(cur *qNode[T]) *qNode[T] {

	llsize, lrsize := getChildrenSize(cur.Children[0])
	if llsize < lrsize {
		tree.lrotate(cur.Children[0])
	}
	return tree.rrotate(cur)
}

func (tree *Queue[T]) lrotate(cur *qNode[T]) *qNode[T] {

	const L = 1
	const R = 0
	// 1 right 0 left
	mov := cur.Children[L]
	movright := mov.Children[R]

	if cur.Parent.Children[L] == cur {
		cur.Parent.Children[L] = mov
	} else {
		cur.Parent.Children[R] = mov
	}
	mov.Parent = cur.Parent

	if movright != nil {
		cur.Children[L] = movright
		movright.Parent = cur
	} else {
		cur.Children[L] = nil
	}

	mov.Children[R] = cur
	cur.Parent = mov

	cur.Size = getChildrenSumSize(cur) + 1
	mov.Size = getChildrenSumSize(mov) + 1

	return mov
}

func (tree *Queue[T]) rrotate(cur *qNode[T]) *qNode[T] {

	const L = 0
	const R = 1
	// 1 right 0 left
	mov := cur.Children[L]
	movright := mov.Children[R]

	if cur.Parent.Children[L] == cur {
		cur.Parent.Children[L] = mov
	} else {
		cur.Parent.Children[R] = mov
	}
	mov.Parent = cur.Parent

	if movright != nil {
		cur.Children[L] = movright
		movright.Parent = cur
	} else {
		cur.Children[L] = nil
	}

	mov.Children[R] = cur
	cur.Parent = mov

	cur.Size = getChildrenSumSize(cur) + 1
	mov.Size = getChildrenSumSize(mov) + 1

	return mov
}

func getChildrenSumSize[T any](cur *qNode[T]) int64 {
	return getSize(cur.Children[0]) + getSize(cur.Children[1])
}

func getChildrenSize[T any](cur *qNode[T]) (int64, int64) {
	return getSize(cur.Children[0]), getSize(cur.Children[1])
}

func getSize[T any](cur *qNode[T]) int64 {
	if cur == nil {
		return 0
	}
	return cur.Size
}

func getRelationship[T any](cur *qNode[T]) int {
	if cur.Parent.Children[1] == cur {
		return 1
	}
	return 0
}

func (tree *Queue[T]) getRangeRoot(low, hight T) (root *qNode[T]) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	for cur != nil {
		c1 := tree.compare(low, cur.Key())
		c2 := tree.compare(hight, cur.Key())
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

func (tree *Queue[T]) mergeGroups(root *qNode[T], group *qNode[T], childGroup *qNode[T], childSize int64, LR int) {
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

func (tree *Queue[T]) fixRemoveRange(cur *qNode[T]) {
	const L = 0
	const R = 1

	if cur.Size <= 2 {
		return
	}

	ls, rs := getChildrenSize(cur)
	if ls > rs && ls >= rs<<1 {
		cls, crs := getChildrenSize(cur.Children[L])
		if cls < crs {
			tree.lrotate(cur.Children[L])
		}
		tree.rrotate(cur)
	} else if ls < rs && rs >= ls<<1 {
		cls, crs := getChildrenSize(cur.Children[R])
		if cls > crs {
			tree.rrotate(cur.Children[R])
		}
		tree.lrotate(cur)
	}
}

func (tree *Queue[T]) index(i int64) *qNode[T] {

	defer func() {
		if err := recover(); err != nil {
			log.Panicln(errOutOfIndex, i)
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

func (tree *Queue[T]) getNode(key T) (result *qNode[T]) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	for cur != nil {
		c := tree.compare(key, cur.Key())
		switch {
		case c < 0:
			cur = cur.Children[L]
		case c > 0:
			cur = cur.Children[R]
		default:

			// 插入最右边. 获取最左 与Put相反
			result = cur

			cur = cur.Children[L]
		}
	}
	return
}

func (tree *Queue[T]) getNodes(key T) (result []*qNode[T]) {
	const L = 0
	const R = 1

	var traverse func(cur *qNode[T])
	traverse = func(cur *qNode[T]) {
		if cur == nil {
			return
		}
		c := tree.compare(key, cur.Key())
		switch {
		case c < 0:
			traverse(cur.Children[L])
		case c > 0:
			traverse(cur.Children[R])
		default:
			// 插入最右边. 获取最左 与Put相反
			traverse(cur.Children[L])
			result = append(result, cur)
			traverse(cur.Children[R])
		}
	}

	traverse(tree.getRoot())
	return
}

func (tree *Queue[T]) getRoot() *qNode[T] {
	return tree.root.Children[0]
}

func (tree *Queue[T]) remove(cur *qNode[T]) *Slice[T] {

	const L = 0
	const R = 1

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

		var s = cur.Slice
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

		var s = cur.Slice
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

func (tree *Queue[T]) check() {
	const L = 0
	const R = 1

	root := tree.getRoot()
	if root != nil && root.Parent != tree.root {
		panic("")
	}

	var tcheck func(root *qNode[T])
	tcheck = func(root *qNode[T]) {

		if root == nil {
			return
		}

		size := root.Size
		if size != getSize(root.Children[L])+getSize(root.Children[R])+1 {
			log.Panic("size error")
		}

		if root.Children[L] != nil {
			if root.Children[L].Parent != root {
				log.Panic("parent error")
			}
		}

		if root.Children[R] != nil {
			if root.Children[R].Parent != root {
				log.Panic("parent error")
			}
		}

		tcheck(root.Children[L])
		tcheck(root.Children[R])
	}
	tcheck(root)

	// cur := tree.getRoot()
	// if cur != nil {
	// 	for cur.Children[L] != nil {
	// 		cur = cur.Children[L]
	// 	}

	// 	if cur != tree.head {
	// 		log.Println(tree.debugStringWithValue())
	// 		log.Panic("head error", tree.head, cur)
	// 	}

	// 	cur = tree.getRoot()
	// 	for cur.Children[R] != nil {
	// 		cur = cur.Children[R]
	// 	}

	// 	if cur != tree.tail {
	// 		log.Println(tree.debugStringWithValue())
	// 		log.Panic("tail error", tree.tail, cur)
	// 	}
	// }

}
