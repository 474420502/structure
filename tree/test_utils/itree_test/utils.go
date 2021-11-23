package indextreetest

import (
	"log"
)

func (tree *Tree) fixPutSize(cur *Node) {
	for cur != tree.root {
		cur.Size++
		cur = cur.Parent
	}
}

func (tree *Tree) fixRemoveSize(cur *Node) {
	for cur != tree.root {
		cur.Size--
		cur = cur.Parent
	}
}

func (tree *Tree) fixPut(cur *Node) {

	cur.Size++

	if cur.Size == 3 {
		tree.fixPutSize(cur.Parent)
		return
	}

	const L = 0
	const R = 1

	var height int64 = 2
	var root2nsize, child2nsize, bottomsize, lsize, rsize int64
	var relations int = L
	var parent *Node

	if cur.Parent.Children[R] == cur {
		relations = R
	}
	cur = cur.Parent // 第三点起
	// var rcount = 0
	// var debugshow []string
	for cur != tree.root {
		cur.Size++
		parent = cur.Parent
		root2nsize = (int64(1) << height) - (height - 2)
		// (1<< height) -1 允许的最大size　超过证明高度超1, 并且有最少１size的空缺
		if cur.Size < root2nsize {

			child2nsize = root2nsize >> 2
			bottomsize = child2nsize + (child2nsize >> (height >> 1))
			lsize, rsize = getChildrenSize(cur)
			// 右就检测左边
			if relations == R {
				if rsize-lsize >= bottomsize {
					// rcount++
					// ckey := setColor(cur, color_5)
					// // log.Println(tree.debugString(false), root2nsize)
					// debugshow = append(debugshow, tree.debugString(true))
					cur = tree.sizeRrotate(cur)
					// log.Println(tree.debugString(false))
					// debugshow = append(debugshow, tree.debugString(true))
					// delColor(ckey)
					height--
					// if rcount >= 2 {
					// 	for _, s := range debugshow {
					// 		log.Println(s)
					// 	}
					// 	log.Println()
					// }
				}
			} else {
				if lsize-rsize >= bottomsize {
					// rcount++
					// ckey := setColor(cur, color_5)
					// // log.Println(tree.debugString(false), root2nsize)
					// debugshow = append(debugshow, tree.debugString(true))
					cur = tree.sizeLrotate(cur)
					// log.Println(tree.debugString(false))
					// debugshow = append(debugshow, tree.debugString(true))
					// delColor(ckey)
					height--
					// if rcount >= 2 {
					// 	for _, s := range debugshow {
					// 		log.Println(s)
					// 	}
					// 	log.Println()
					// }
				}
			}
		}

		height++
		if cur.Parent.Children[R] == cur {
			relations = R
		} else {
			relations = L
		}

		cur = parent
	}
}

func (tree *Tree) sizeRrotate(cur *Node) *Node {
	const R = 1
	llsize, lrsize := getChildrenSize(cur.Children[R])
	if llsize > lrsize {
		tree.rrotate(cur.Children[R])
	}
	return tree.lrotate(cur)
}

func (tree *Tree) sizeLrotate(cur *Node) *Node {
	const L = 0
	llsize, lrsize := getChildrenSize(cur.Children[L])
	if llsize < lrsize {
		tree.lrotate(cur.Children[L])
	}
	return tree.rrotate(cur)
}

func (tree *Tree) lrotate(cur *Node) *Node {

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

func (tree *Tree) rrotate(cur *Node) *Node {

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

func getChildrenSumSize(cur *Node) int64 {
	return getSize(cur.Children[0]) + getSize(cur.Children[1])
}

func getChildrenSize(cur *Node) (int64, int64) {
	return getSize(cur.Children[0]), getSize(cur.Children[1])
}

func getSize(cur *Node) int64 {
	if cur == nil {
		return 0
	}
	return cur.Size
}

func getRelationship(cur *Node) int {
	if cur.Parent.Children[1] == cur {
		return 1
	}
	return 0
}

func (tree *Tree) getRangeRoot(low, hight interface{}) (root *Node) {
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

func (tree *Tree) mergeGroups(root *Node, group *Node, childGroup *Node, childSize int64, LR int) {
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

func (tree *Tree) fixRemoveRange(cur *Node) {
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

func (tree *Tree) index(i int64) *Node {

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

func (tree *Tree) getNode(key interface{}) *Node {
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

func (tree *Tree) getRoot() *Node {
	return tree.root.Children[0]
}

func (tree *Tree) check() {
	const L = 0
	const R = 1

	root := tree.getRoot()
	if root != nil && root.Parent != tree.root {
		panic("")
	}

	var tcheck func(root *Node)
	tcheck = func(root *Node) {

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

func (tree *Tree) hight() int {

	root := tree.getRoot()

	maxHight := 0
	var getHigh func(cur *Node, hight int)
	getHigh = func(cur *Node, hight int) {
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
