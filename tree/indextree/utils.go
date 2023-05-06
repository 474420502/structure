package indextree

import (
	"log"
)

func (tree *Tree[T]) fixPutSize(cur *hNode[T]) {
	for cur != tree.root {
		cur.Size++
		cur = cur.Parent
	}
}

func (tree *Tree[T]) fixRemoveSize(cur *hNode[T]) {
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

func (tree *Tree[T]) fixPut(cur *hNode[T]) {

	cur.Size++
	if cur.Size == 3 {
		tree.fixPutSize(cur.Parent)
		return
	}

	var height int64 = 2
	var lsize, rsize int64
	var parent *hNode[T]

	cur = cur.Parent

	for cur != tree.root {
		cur.Size++
		parent = cur.Parent

		limitSize := rootSizeTable[height]

		// (1<< height) -1 允许的最大size　超过证明高度超1, 并且有最少１size的空缺
		if cur.Size < limitSize.rootsize {

			lsize, rsize = getChildrenSize(cur)
			// 右就检测左边
			if rsize > lsize {
				if rsize-lsize >= limitSize.bottomsize {
					tree.sizeRRotate(cur)
					// height--
					tree.fixPutSize(parent)
					return
				}
			} else {
				if lsize-rsize >= limitSize.bottomsize {
					tree.sizeLRotate(cur)
					// height--
					tree.fixPutSize(parent)
					return
				}
			}
		}

		height++
		cur = parent
	}
}

func (tree *Tree[T]) sizeRRotate(cur *hNode[T]) *hNode[T] {
	const R = 1
	llsize, lrsize := getChildrenSize(cur.Children[R])
	if llsize > lrsize {
		tree.rrotate(cur.Children[R])
	}
	return tree.lrotate(cur)
}

func (tree *Tree[T]) sizeLRotate(cur *hNode[T]) *hNode[T] {
	const L = 0
	llsize, lrsize := getChildrenSize(cur.Children[L])
	if llsize < lrsize {
		tree.lrotate(cur.Children[L])
	}
	return tree.rrotate(cur)
}

func (tree *Tree[T]) lrotate(cur *hNode[T]) *hNode[T] {

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

func (tree *Tree[T]) rrotate(cur *hNode[T]) *hNode[T] {

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

func getChildrenSumSize[T any](cur *hNode[T]) int64 {
	return getSize(cur.Children[0]) + getSize(cur.Children[1])
}

func getChildrenSize[T any](cur *hNode[T]) (int64, int64) {
	return getSize(cur.Children[0]), getSize(cur.Children[1])
}

func getSize[T any](cur *hNode[T]) int64 {
	if cur == nil {
		return 0
	}
	return cur.Size
}

func getRelationship[T any](cur *hNode[T]) int {
	if cur.Parent.Children[1] == cur {
		return 1
	}
	return 0
}

func (tree *Tree[T]) getRangeRoot(low, hight T) (root *hNode[T]) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	for cur != nil {
		c1 := tree.compare(low, cur.Key)
		c2 := tree.compare(hight, cur.Key)
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

func (tree *Tree[T]) mergeGroups(root *hNode[T], group *hNode[T], childGroup *hNode[T], childSize int64, LR int) {
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

func (tree *Tree[T]) fixRemoveRange(cur *hNode[T]) {
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

func (tree *Tree[T]) index(i int64) *hNode[T] {

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

func (tree *Tree[T]) getNode(key T) *hNode[T] {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	for cur != nil {
		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:
			cur = cur.Children[L]
		case c > 0:
			cur = cur.Children[R]
		default:
			return cur
		}
	}
	return nil
}

func (tree *Tree[T]) getRoot() *hNode[T] {
	return tree.root.Children[0]
}

func (tree *Tree[T]) check() {
	const L = 0
	const R = 1

	root := tree.getRoot()
	if root != nil && root.Parent != tree.root {
		panic("")
	}

	var tcheck func(root *hNode[T])
	tcheck = func(root *hNode[T]) {

		if root == nil {
			return
		}

		size := root.Size
		if size != getSize(root.Children[L])+getSize(root.Children[R])+1 {
			log.Println(tree.debugString(true))
			log.Panic("size error")
		}

		if root.Children[L] != nil {
			if root.Children[L].Parent != root {
				log.Println(tree.debugString(true))
				log.Panicln("parent error", root.Children[L].Key, root.Key)
			}
		}

		if root.Children[R] != nil {
			if root.Children[R].Parent != root {
				log.Println(tree.debugString(true))
				log.Panicln("parent error", root.Children[R].Key, root.Key)
			}
		}

		tcheck(root.Children[L])
		tcheck(root.Children[R])
	}
	tcheck(root)

}

func (tree *Tree[T]) hight() int {

	root := tree.getRoot()

	maxHight := 0
	var getHigh func(cur *hNode[T], hight int)
	getHigh = func(cur *hNode[T], hight int) {
		if cur == nil {
			return
		}
		if cur.Size == 1 {
			if maxHight < hight {
				maxHight = hight
			}
			return
		}
		getHigh(cur.Children[0], hight+1)
		getHigh(cur.Children[1], hight+1)
	}

	getHigh(root, 0)
	return maxHight
}

func (node *hNode[T]) hight() int {
	root := node

	maxHight := 0
	var getHigh func(cur *hNode[T], hight int)
	getHigh = func(cur *hNode[T], hight int) {
		if cur == nil {
			return
		}
		if cur.Size == 1 {
			if maxHight < hight {
				maxHight = hight
			}
			return
		}
		getHigh(cur.Children[0], hight+1)
		getHigh(cur.Children[1], hight+1)
	}

	getHigh(root, 0)
	return maxHight
}
