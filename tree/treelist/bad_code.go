package treelist

// func (tree *Tree) getRangeNodeStart(cur *Node, key []byte) (groups []*Node, left *Node) {
// 	const L = 0
// 	const R = 1

// 	for {
// 		c := tree.compare(key, cur.Key)
// 		switch {
// 		case c < 0:
// 			if cur.Children[L] == nil {
// 				left = cur.Direct[L]
// 				groups = append(groups, cur.Children[L])
// 				return
// 			}
// 			cur = cur.Children[L]

// 		case c > 0:
// 			groups = append(groups, cur)
// 			if cur.Children[R] == nil {
// 				left = cur
// 				return
// 			}
// 			cur = cur.Children[R]

// 		default:
// 			left = cur.Direct[L]
// 			groups = append(groups, cur.Children[L])
// 			return
// 		}
// 	}
// }

// func (tree *Tree) getRangeNodeEnd(root *Node, key []byte) (groups []*Node, right *Node) {
// 	const L = 0
// 	const R = 1

// 	cur := root
// 	for cur != nil {
// 		c := tree.compare(key, cur.Key)
// 		switch {
// 		case c < 0:
// 			groups = append(groups, cur)
// 			if cur.Children[L] == nil {
// 				right = cur
// 				return
// 			}
// 			cur = cur.Children[L]
// 		case c > 0:
// 			if cur.Children[R] == nil {
// 				right = cur.Direct[R]
// 				groups = append(groups, cur.Children[R])
// 				return
// 			}
// 			cur = cur.Children[R]

// 		default:
// 			right = cur.Direct[R]
// 			groups = append(groups, cur.Children[R])
// 			return
// 		}
// 	}
// 	return
// }

// // getRangeNodes 获取范围节点的左团和又团
// func (tree *Tree) getRangeNodes(low, hight []byte) (root *Node, start []*Node, left *Node, end []*Node, right *Node) {
// 	const L = 0
// 	const R = 1

// 	cur := tree.getRoot()
// 	for cur != nil {
// 		c1 := tree.compare(low, cur.Key)
// 		c2 := tree.compare(hight, cur.Key)

// 		if c1 != c2 {
// 			starts, dleft := tree.getRangeNodeStart(cur, low)
// 			ends, dright := tree.getRangeNodeEnd(cur, hight)
// 			return cur, starts, dleft, ends, dright
// 		}

// 		if c1 < 0 {
// 			cur = cur.Children[L]
// 		} else {
// 			cur = cur.Children[R]
// 		}

// 	}
// 	return
// }

func (tree *Tree) trimBad(low, hight interface{}) {
	const L = 0
	const R = 1
	root := tree.getRoot()
	var trim func(root *treeNode) *treeNode
	trim = func(root *treeNode) *treeNode {
		if root == nil {
			return nil
		}

		if tree.compare(root.Key, hight) > 0 {
			return trim(root.Children[L])
		}

		if tree.compare(root.Key, low) < 0 {
			return trim(root.Children[R])
		}

		root.Children[L] = trim(root.Children[L])
		root.Children[R] = trim(root.Children[R])
		root.Size = getChildrenSumSize(root) + 1
		return root
	}

	croot := trim(root)
	if root != croot {
		tree.root.Children[L] = croot
	}
	// list
	if croot != nil {
		croot.Parent = tree.root

		lhand := croot
		for lhand.Children[L] != nil {
			lhand = lhand.Children[L]
		}
		lhand.Direct[L] = nil

		rhand := croot
		for rhand.Children[R] != nil {
			rhand = rhand.Children[R]
		}
		rhand.Direct[R] = nil
	}
}

// func (tree *Tree) trimVeryBad(root *Node, low, hight []byte) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	if tree.compare(root.Key, hight) > 0 {
// 		return tree.trimVeryBad(root.Children[0], low, hight)
// 	}

// 	if tree.compare(root.Key, low) < 0 {
// 		return tree.trimVeryBad(root.Children[1], low, hight)
// 	}

// 	root.Children[0] = tree.trimVeryBad(root.Children[0], low, hight)
// 	root.Children[1] = tree.trimVeryBad(root.Children[1], low, hight)
// 	root.Size = getChildrenSumSize(root) + 1
// 	return root
// }

// func combineGroups(starts []*Node, LR int) *Node {

// 	var group *Node
// 	var child *Node

// 	nlen := len(starts)
// 	if nlen == 0 {
// 		return nil
// 	}
// 	group = starts[nlen-1]
// 	child = group
// 	for i := nlen - 2; i >= 0; i-- {
// 		group = starts[i]
// 		combine(group, child, LR)
// 		child = group
// 	}
// 	return group
// }

// func combine(group *Node, child *Node, LR int) {
// 	if group != nil {
// 		hand := group

// 		hand.Children[LR] = child
// 		if child != nil {
// 			child.Parent = hand
// 		}

// 		for hand != group.Parent {
// 			hand.Size = getChildrenSumSize(hand) + 1
// 			hand = hand.Parent
// 		}
// 	}
// }

// func (tree *Tree) removeRangeBad(low, hight []byte) {
// 	const L = 0
// 	const R = 1

// 	c := tree.compare(low, hight)
// 	if c > 0 {
// 		panic("key2 must greater than key1 or equal to")
// 	} else if c == 0 {
// 		tree.Remove(low)
// 		return
// 	}

// 	root, starts, dleft, ends, dright := tree.getRangeNodes(low, hight)
// 	if root == nil {
// 		return
// 	}

// 	if dleft != nil {
// 		dleft.Direct[R] = dright
// 	}

// 	if dright != nil {
// 		dright.Direct[L] = dleft
// 	}

// 	// 合并左树
// 	lgroup := combineGroups(starts, R)
// 	// 合并又树
// 	rgroup := combineGroups(ends, L)

// 	if lgroup == nil && rgroup == nil {
// 		rparent := root.Parent
// 		size := root.Size
// 		root.Parent.Children[getRelationship(root)] = nil
// 		for rparent != tree.root {
// 			rparent.Size -= size
// 			rparent = rparent.Parent
// 		}
// 		return
// 	}
// 	// log.Println(debugNode(lgroup), "\n", debugNode(rgroup))
// 	// log.Println(tree.debugString(true))
// 	// log.Println(root, starts, ends)

// 	// 左右树　拼接
// 	rsize := getSize(rgroup)
// 	lsize := getSize(lgroup)
// 	if lsize > rsize {
// 		tree.mergeGroups(root, lgroup, rgroup, rsize, R)
// 	} else {
// 		tree.mergeGroups(root, rgroup, lgroup, lsize, L)
// 	}
// }
