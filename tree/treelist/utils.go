package treelist

import (
	"log"
)

// func (tree *Tree[T]) hashString() string {
// 	var buf = bytes.NewBuffer(nil)
// 	w := zlib.NewWriter(buf)

// 	tree.Traverse(func(s *Slice[T]) bool {
// 		w.Write(s.Key)
// 		return true
// 	})

// 	err := w.Close()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return buf.String()
// }

func (tree *Tree[T]) getRoot() *treeNode[T] {
	return tree.root.Children[0]
}

// getRangeNodes 获取范围节点的左团和又团
func (tree *Tree[T]) getRangeRoot(low, hight T) (root *treeNode[T]) {
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

func (tree *Tree[T]) fixPutSize(cur *treeNode[T]) {
	for cur != tree.root {
		cur.Size++
		cur = cur.Parent
	}
}

func (tree *Tree[T]) fixRemoveSize(cur *treeNode[T]) {
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

func (tree *Tree[T]) fixPut(cur *treeNode[T]) {
	cur.Size++
	if cur.Size == 3 {
		tree.fixPutSize(cur.Parent)
		return
	}

	const L = 0
	const R = 1

	var height int64 = 2
	var lsize, rsize int64
	var relations int = L
	var parent *treeNode[T]

	if cur.Parent.Children[R] == cur {
		relations = R
	}
	cur = cur.Parent

	for cur != tree.root {
		cur.Size++
		parent = cur.Parent

		limitsize := rootSizeTable[height]
		// (1<< height) -1 允许的最大size　超过证明高度超1, 并且有最少１size的空缺
		if cur.Size < limitsize.rootsize {

			lsize, rsize = getChildrenSize(cur)
			// 右就检测左边
			if relations == R {
				if rsize-lsize >= limitsize.bottomsize {
					cur = tree.sizeRrotate(cur)
					tree.fixPutSize(parent)
					return
				}
			} else {
				if lsize-rsize >= limitsize.bottomsize {
					cur = tree.sizeLrotate(cur)
					tree.fixPutSize(parent)
					return
				}
			}
		}

		height++
		if parent.Children[R] == cur {
			relations = R
		} else {
			relations = L
		}

		cur = parent
	}
}

func (tree *Tree[T]) fixRemoveRange(cur *treeNode[T]) {
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

func (tree *Tree[T]) sizeRrotate(cur *treeNode[T]) *treeNode[T] {
	const R = 1
	llsize, lrsize := getChildrenSize(cur.Children[R])
	if llsize > lrsize {
		tree.rrotate(cur.Children[R])
	}
	return tree.lrotate(cur)
}

func (tree *Tree[T]) sizeLrotate(cur *treeNode[T]) *treeNode[T] {
	const L = 0
	llsize, lrsize := getChildrenSize(cur.Children[L])
	if llsize < lrsize {
		tree.lrotate(cur.Children[L])
	}
	return tree.rrotate(cur)
}

func (tree *Tree[T]) lrotate(cur *treeNode[T]) *treeNode[T] {
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

func (tree *Tree[T]) rrotate(cur *treeNode[T]) *treeNode[T] {
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

func (tree *Tree[T]) mergeGroups(root *treeNode[T], group *treeNode[T], childGroup *treeNode[T], childSize int64, LR int) {
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

func (tree *Tree[T]) removeNode(cur *treeNode[T]) (s *Slice[T]) {
	const L = 0
	const R = 1
	s = &Slice[T]{Key: cur.Key, Value: cur.Value}

	if cur.Size == 1 {
		parent := cur.Parent
		if parent != tree.root {
			parent.Children[getRelationship(cur)] = nil
			tree.fixRemoveSize(parent)

			dright := cur.Direct[R]
			dleft := cur.Direct[L]

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

		} else {
			parent.Children[0] = nil
			tree.root.Direct[L] = nil
			tree.root.Direct[R] = nil
		}

		return
	}

	lsize, rsize := getChildrenSize(cur)
	if lsize > rsize {
		prev := cur.Children[L]
		for prev.Children[R] != nil {
			prev = prev.Children[R]
		}

		cur.Key = prev.Key
		cur.Value = prev.Value

		prevParent := prev.Parent
		if prevParent == cur {
			cur.Children[L] = prev.Children[L]
			cleft := cur.Children[L]
			if cleft != nil {
				cleft.Parent = cur
			}

			tree.fixRemoveSize(cur)
		} else {

			prevParent.Children[R] = prev.Children[L]
			if prevParent.Children[R] != nil {
				prevParent.Children[R].Parent = prevParent
			}
			tree.fixRemoveSize(prevParent)
		}

		dleft := cur.Direct[L].Direct[L]
		if dleft != nil {
			dleft.Direct[R] = cur
		} else {
			tree.root.Direct[L] = cur
		}

		cur.Direct[L] = dleft

	} else {

		next := cur.Children[R]
		for next.Children[L] != nil {
			next = next.Children[L]
		}

		cur.Key = next.Key
		cur.Value = next.Value

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

		dright := cur.Direct[R].Direct[R]
		if dright != nil {
			dright.Direct[L] = cur
		} else {
			tree.root.Direct[R] = cur
		}
		cur.Direct[R] = dright
	}

	return
}

func (tree *Tree[T]) head() *treeNode[T] {
	return tree.root.Direct[0]
}

func (tree *Tree[T]) tail() *treeNode[T] {
	return tree.root.Direct[1]
}

func (tree *Tree[T]) index(i int64) *treeNode[T] {

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

func (tree *Tree[T]) getNode(key T) *treeNode[T] {
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

func (tree *Tree[T]) seekNodeWithIndex(key T) (node *treeNode[T], idx int64, dir int) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	var offset int64 = getSize(cur.Children[L])
	var last *treeNode[T]
	var c int
	for {
		c = tree.compare(key, cur.Key)
		last = cur

		switch {
		case c < 0:

			cur = cur.Children[L]
			if cur != nil {
				offset -= getSize(cur.Children[R]) + 1
			} else {
				return last, offset, c
			}

		case c > 0:

			cur = cur.Children[R]
			if cur != nil {
				offset += getSize(cur.Children[L]) + 1
			} else {
				return last, offset, c
			}

		default:
			return cur, offset, c
		}
	}

}

func getChildrenSumSize[T any](cur *treeNode[T]) int64 {
	return getSize(cur.Children[0]) + getSize(cur.Children[1])
}

func getChildrenSize[T any](cur *treeNode[T]) (int64, int64) {
	return getSize(cur.Children[0]), getSize(cur.Children[1])
}

func getSize[T any](cur *treeNode[T]) int64 {
	if cur == nil {
		return 0
	}
	return cur.Size
}

func getRelationship[T any](cur *treeNode[T]) int {
	if cur.Parent.Children[1] == cur {
		return 1
	}
	return 0
}

// func (tree *Tree[T]) getHeight() int {
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

func (tree *Tree[T]) find(a1, a2 []*Slice[T]) (result []*Slice[T]) {

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

func (tree *Tree[T]) unionSetSlice(other *Tree[T]) (result []*Slice[T]) {

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
			// result = append(result, &head2.Slice[T])
			head1 = head1.Direct[R]
		case c > 0:
			// result = append(result, &head1.Slice[T])
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

func (tree *Tree[T]) intersectionSlice(other *Tree[T]) (result []*Slice[T]) {

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

func (tree *Tree[T]) check() {
	const L = 0
	const R = 1

	root := tree.getRoot()
	if root != nil && root.Parent != tree.root {
		panic("")
	}

	var tcheck func(root *treeNode[T])
	tcheck = func(root *treeNode[T]) {

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

}
