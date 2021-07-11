package indextree

func (tree *IndexTree) fixSize(cur *Node) {
	for cur != tree.root {
		cur.Size++
		cur = cur.Parent
	}
}

func (tree *IndexTree) fixPut(cur *Node) {

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
			// 右就检测左边
			if relations == R {
				lsize := getSize(cur.Children[L])
				rsize := getSize(cur.Children[R])
				if rsize-lsize >= bottomsize {
					cur = tree.sizeRrotate(cur)
					height--
				}
			} else {
				rsize := getSize(cur.Children[R])
				lsize := getSize(cur.Children[L])
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

func (tree *IndexTree) sizeRrotate(cur *Node) *Node {
	const R = 1
	llsize, lrsize := getChildrenSize(cur.Children[R])
	if llsize > lrsize {
		tree.rrotate(cur.Children[R])
	}
	return tree.lrotate(cur)
}

func (tree *IndexTree) sizeLrotate(cur *Node) *Node {
	const L = 0
	llsize, lrsize := getChildrenSize(cur.Children[L])
	if llsize < lrsize {
		tree.lrotate(cur.Children[L])
	}
	return tree.rrotate(cur)
}

func (tree *IndexTree) lrotate(cur *Node) *Node {

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

func (tree *IndexTree) rrotate(cur *Node) *Node {

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
