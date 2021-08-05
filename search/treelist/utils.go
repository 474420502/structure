package treelist

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

func (tree *Tree) fixRemoveRange(cur *Node) {
	const L = 0
	const R = 1

	if cur.Size <= 2 {
		return
	}

	// log.Println(tree.debugString(true))

	ls, rs := getChildrenSize(cur)
	if ls > rs && ls >= rs<<1 {
		cls, crs := getChildrenSize(cur.Children[L])
		if cls < crs {
			tree.lrotate(cur.Children[L])
		}
		tree.rrotate(cur)
		// tree.fixRemoveRange(cur)
		// tree.fixRemoveRange(root, level+1)
	} else if ls < rs && rs >= ls<<1 {
		cls, crs := getChildrenSize(cur.Children[R])
		if cls > crs {
			tree.rrotate(cur.Children[R])
		}
		tree.lrotate(cur)
		// tree.fixRemoveRange(cur)
		// tree.fixRemoveRange(root, level+1)
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
	// tree.rcount++

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
	// tree.rcount++

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

// func (tree *Tree) getHeight() int {
// 	root := tree.getRoot()
// 	if root == nil {
// 		return 0
// 	}

// 	var height = 1

// 	var traverse func(cur *Node, h int)
// 	traverse = func(cur *Node, h int) {

// 		if cur == nil {
// 			return
// 		}

// 		if h > height {
// 			height = h
// 		}

// 		traverse(cur.Children[0], h+1)
// 		traverse(cur.Children[1], h+1)
// 	}

// 	traverse(root, 1)

// 	return height
// }

func (tree *Tree) find(a1, a2 []*Slice) (result []*Slice) {

	h1 := 0
	h2 := 0

	// var count = 0

	for h1 < len(a1) && h2 < len(a2) {
		c := tree.compare(a1[h1].Key, a2[h2].Key)
		// count++
		switch {
		case c < 0:
			h1++
		case c > 0:
			h2++
		default:
			result = append(result, a1[h1])
			h1++
			h2++
		}
	}

	// log.Println("count:", count)
	return
}

func (tree *Tree) unionSetSlice(other *Tree) (result []*Slice) {

	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	for head1 != nil && head2 != nil {
		c := tree.compare(head1.Key, head2.Key)
		// count++
		switch {
		case c < 0:
			result = append(result, &head1.Slice)
			// result = append(result, &head2.Slice)
			head1 = head1.Direct[R]
		case c > 0:
			// result = append(result, &head1.Slice)
			result = append(result, &head2.Slice)
			head2 = head2.Direct[R]

		default:
			result = append(result, &head1.Slice)
			head1 = head1.Direct[R]
			head2 = head2.Direct[R]
		}
	}

	for head1 != nil {
		result = append(result, &head1.Slice)
		head1 = head1.Direct[R]
	}

	for head2 != nil {
		result = append(result, &head2.Slice)
		head2 = head2.Direct[R]
	}

	return
}

func (tree *Tree) intersectionSlice(other *Tree) (result []*Slice) {

	const L = 0
	const R = 1

	// count := 0

	head1 := tree.head()
	head2 := other.head()

	for head1 != nil && head2 != nil {
		c := tree.compare(head1.Key, head2.Key)
		// count++

		switch {
		case c < 0:
			head1 = head1.Direct[R]
		case c > 0:
			head2 = head2.Direct[R]
		default:
			result = append(result, &head1.Slice)
			head1 = head1.Direct[R]
			head2 = head2.Direct[R]
		}
	}

	// log.Println("count:", count, tree.Size(), other.Size())
	return
}
