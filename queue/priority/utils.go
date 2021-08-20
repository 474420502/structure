package treequeue

import (
	"log"
)

var errOutOfIndex = "out of index"

func (tree *Queue) fixPutSize(cur *qNode) {
	for cur != tree.root {
		cur.Size++
		cur = cur.Parent
	}
}

func (tree *Queue) fixRemoveSize(cur *qNode) {
	for cur != tree.root {
		cur.Size--
		cur = cur.Parent
	}
}

func (tree *Queue) fixPut(cur *qNode) {

	tree.fixPutSize(cur)
	if cur.Size == 3 {
		return
	}

	const L = 0
	const R = 1

	var height int64 = 2

	var relations int = L
	if cur.Parent.Children[R] == cur {
		relations = R
	}
	cur = cur.Parent

	for cur != tree.root {

		root2nsize := (int64(1) << height)
		// (1<< height) -1 允许的最大size　超过证明高度超1, 并且有最少１size的空缺
		if cur.Size < root2nsize {

			child2nsize := root2nsize >> 2
			bottomsize := child2nsize + child2nsize>>(height>>1)
			lsize, rsize := getChildrenSize(cur)
			// 右就检测左边
			if relations == R {
				if rsize-lsize >= bottomsize {
					cur = tree.sizeRrotate(cur)
					height--
				}
			} else {
				if lsize-rsize >= bottomsize {
					cur = tree.sizeLrotate(cur)
					height--
				}
			}
		}

		height++
		if cur.Parent.Children[R] == cur {
			relations = R
		} else {
			relations = L
		}

		cur = cur.Parent
	}
}

func (tree *Queue) sizeRrotate(cur *qNode) *qNode {
	const R = 1
	llsize, lrsize := getChildrenSize(cur.Children[R])
	if llsize > lrsize {
		tree.rrotate(cur.Children[R])
	}
	return tree.lrotate(cur)
}

func (tree *Queue) sizeLrotate(cur *qNode) *qNode {
	const L = 0
	llsize, lrsize := getChildrenSize(cur.Children[L])
	if llsize < lrsize {
		tree.lrotate(cur.Children[L])
	}
	return tree.rrotate(cur)
}

func (tree *Queue) lrotate(cur *qNode) *qNode {

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

func (tree *Queue) rrotate(cur *qNode) *qNode {

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

func getChildrenSumSize(cur *qNode) int64 {
	return getSize(cur.Children[0]) + getSize(cur.Children[1])
}

func getChildrenSize(cur *qNode) (int64, int64) {
	return getSize(cur.Children[0]), getSize(cur.Children[1])
}

func getSize(cur *qNode) int64 {
	if cur == nil {
		return 0
	}
	return cur.Size
}

func getRelationship(cur *qNode) int {
	if cur.Parent.Children[1] == cur {
		return 1
	}
	return 0
}

func (tree *Queue) getRangeRoot(low, hight interface{}) (root *qNode) {
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

func (tree *Queue) mergeGroups(root *qNode, group *qNode, childGroup *qNode, childSize int64, LR int) {
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

func (tree *Queue) fixRemoveRange(cur *qNode) {
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

func (tree *Queue) index(i int64) *qNode {

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

func (tree *Queue) getNode(key interface{}) (result *qNode) {
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

func (tree *Queue) getNodes(key interface{}) (result []*qNode) {
	const L = 0
	const R = 1

	var traverse func(cur *qNode)
	traverse = func(cur *qNode) {
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

func (tree *Queue) getRoot() *qNode {
	return tree.root.Children[0]
}

func (tree *Queue) remove(cur *qNode) *Slice {

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

func (tree *Queue) check() {
	const L = 0
	const R = 1

	root := tree.getRoot()
	if root != nil && root.Parent != tree.root {
		panic("")
	}

	var tcheck func(root *qNode)
	tcheck = func(root *qNode) {

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
