package itree

import (
	"log"
)

var errOutOfIndex = "out of index"
var errLowerGtHigh = "low is behind high"

func (tree *Tree[KEY, VALUE]) getRoot() *Node[KEY, VALUE] {
	return tree.Center.Children[1]
}

func (tree *Tree[KEY, VALUE]) put(parent *Node[KEY, VALUE], child int, key KEY) (target *Node[KEY, VALUE], isExists int) {

	cur := parent.Children[child]
	if cur == nil {
		target = newNode[KEY, VALUE]()
		target.Key = key
		parent.Children[child] = target
		if parent.Children[^child+2] == nil {
			return target, 0
		}
		return target, 2
	}

	cmp := tree.Compare(cur.Key, key)

	if cmp < 0 {
		return cur, 1
	} else {
		target, isExists = tree.put(cur, cmp, key)
		if isExists == 1 {
			return target, isExists
		}
	}

	cur.Size++
	if cur.Size > 2 && isExists != 2 {
		if tree.rebalance(parent, child) {
			isExists = 2
		}
	}

	// lastsize := maysize
	// maysize = maysize << 1
	// r := cur.view()
	// if maysize-cur.Size-1 > lastsize {
	// 	tree.rebalance(parent, child)
	// 	log.Println(lastsize, maysize, cur.Size, r)
	// 	log.Println()

	// }

	return target, isExists
}

func (tree *Tree[KEY, VALUE]) get(key KEY, cur *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	if cur == nil {
		return nil
	}
	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		return cur
	}
	return tree.get(key, cur.Children[cmp])
}

func (tree *Tree[KEY, VALUE]) index(i int) *Node[KEY, VALUE] {

	cur := tree.getRoot()
	var idx int = cur.Children[0].getSize()
	for {
		if idx > i {
			cur = cur.Children[0]
			idx -= cur.Children[1].getSize() + 1
		} else if idx < i {
			cur = cur.Children[1]
			idx += cur.Children[0].getSize() + 1
		} else {
			return cur
		}
	}

}

func (tree *Tree[KEY, VALUE]) removeIndex(parent *Node[KEY, VALUE], child, idx, i int) (target *VALUE) {

	// cur := tree.getRoot()
	// var idx int = cur.Children[0].getSize()
	parent.Size--
	cur := parent.Children[child]
	if idx == i {
		if cur.Children[0] == nil {
			parent.Children[child] = cur.Children[1]
			return &cur.Value
		}

		if cur.Children[1] == nil {
			parent.Children[child] = cur.Children[0]
			return &cur.Value
		}

		var result = cur.Value
		target = &result

		replacer, _ := tree.neighboring(cur, child, ^child+2)
		cur.Key = replacer.Key
		cur.Value = replacer.Value
		cur.updateSize()
		tree.rebalance(parent, child)

		return target
	}

	if idx > i {
		cur = cur.Children[0]
		idx -= cur.Children[1].getSize() + 1
		return tree.removeIndex(cur, 0, idx, i)
	} else {
		cur = cur.Children[1]
		idx += cur.Children[0].getSize() + 1
		return tree.removeIndex(cur, 1, idx, i)
	}

}

func (tree *Tree[KEY, VALUE]) remove(key KEY, grandpa *Node[KEY, VALUE], child2, child1 int) (target *VALUE) {
	parent := grandpa.Children[child2]
	cur := parent.Children[child1]

	if cur == nil {
		return nil
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {

		// remove 两种状态. 当前值不在底, 在底
		if cur.Children[0] == nil {
			parent.Children[child1] = cur.Children[1]
			return &cur.Value
		}

		if cur.Children[1] == nil {
			parent.Children[child1] = cur.Children[0]
			return &cur.Value
		}

		var result = cur.Value
		target = &result

		replacer, _ := tree.neighboring(cur, child1, ^child1+2)
		cur.Key = replacer.Key
		cur.Value = replacer.Value
		cur.updateSize()
		tree.rebalance(parent, child1)
		return target
	}

	target = tree.remove(key, parent, child1, cmp)
	if target != nil {
		cur.Size--
		if cur != tree.Center {
			tree.rebalance(parent, child1)
		}
	}

	return target
}

func (tree *Tree[KEY, VALUE]) neighboring(parent *Node[KEY, VALUE], child2, child1 int) (*Node[KEY, VALUE], bool) {
	cur := parent.Children[child2]
	sub := cur.Children[child1]

	if sub == nil {
		other := cur.Children[^child1+2]
		parent.Children[child2] = other
		return cur, true
	}

	result, isRebalance := tree.neighboring(cur, child1, child1)
	cur.Size--
	tree.rebalance(parent, child2)

	return result, isRebalance
}

func sizeRotateType[KEY, VALUE any](cur *Node[KEY, VALUE], child int, ls, rs int) bool {
	rchild := ^child + 2
	// ls, rs := getSize(cur.Children[child]), getSize(cur.Children[rchild])

	var subsize = 0
	var subsubsize = 0
	sub := cur.Children[child]
	if sub != nil {
		subsub := sub.Children[rchild]
		subsize = getSize(subsub)
		if subsub != nil {
			subsubsize = getSize(subsub.Children[rchild])
		}
	}

	sdiff1 := ls - rs - subsize - subsize - 2
	if sdiff1 < 0 {
		sdiff1 = -sdiff1
	}

	sdiff2 := ls - rs - subsubsize - subsubsize - 2
	if sdiff2 < 0 {
		sdiff2 = -sdiff2
	}

	if sdiff1 > sdiff2 {
		// rightRotateWithLeft(parent, child)
		return true
	} else {
		// rightRotate(parent, child)
		return false
	}
}

func (tree *Tree[KEY, VALUE]) rebalance(parent *Node[KEY, VALUE], child int) bool {

	node := parent.Children[child]

	lsize, rsize := getSize(node.Children[0]), getSize(node.Children[1])
	if lsize > rsize {
		if lsize-rsize > rsize {
			lh, rh := getMaybeHeight(lsize), getMaybeHeight(rsize)
			if lh-rh > 1 {
				// log.Println("rotate")
				// tree.rotateCount++
				if sizeRotateType(node, 0, lsize, rsize) {
					rightRotateWithLeft(parent, child)
				} else {
					rightRotate(parent, child)
				}
				return true
			}
		}
	} else {
		if rsize-lsize > lsize {
			lh, rh := getMaybeHeight(lsize), getMaybeHeight(rsize)

			if rh-lh > 1 {
				// log.Println("rotate")
				// tree.rotateCount++
				if sizeRotateType(node, 1, rsize, lsize) {
					leftRotateWithRight(parent, child)
				} else {
					leftRotate(parent, child)
				}
				return true
			}

		}
	}
	return false
}

func (tree *Tree[KEY, VALUE]) trimLow(parent *Node[KEY, VALUE], child int, key KEY) *Node[KEY, VALUE] {

	cur := parent.Children[child]

	if cur == nil {
		return nil
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		cur.Children[0] = nil
		return cur
	}

	if cmp == 1 {
		cur = tree.trimLow(cur, cmp, key)
	} else {
		cur.Children[cmp] = tree.trimLow(cur, cmp, key)
	}
	if cur != nil {
		cur.updateSize()
	}
	return cur
}

func (tree *Tree[KEY, VALUE]) trimHigh(parent *Node[KEY, VALUE], child int, key KEY) *Node[KEY, VALUE] {

	cur := parent.Children[child]

	if cur == nil {
		return nil
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		cur.Children[1] = nil
		return cur
	}

	if cmp == 0 {
		cur = tree.trimHigh(cur, cmp, key)
	} else {
		cur.Children[cmp] = tree.trimHigh(cur, cmp, key)
	}
	if cur != nil {
		cur.updateSize()
	}
	return cur
}

func (tree *Tree[KEY, VALUE]) seekTrimRoot(cur *Node[KEY, VALUE], low, high KEY) *Node[KEY, VALUE] {

	if cur == nil {
		return nil
	}

	cmplow := tree.Compare(cur.Key, low)
	if cmplow < 0 {
		return cur
	}
	cmphigh := tree.Compare(cur.Key, high)
	if cmplow < 0 {
		return cur
	}

	if cmplow != cmphigh {
		return cur
	}

	return tree.seekTrimRoot(cur.Children[cmplow], low, high)
}

func (tree *Tree[KEY, VALUE]) check() (result string) {
	if errorNode := tree.checkSizeTree(tree.getRoot()); errorNode != nil {
		log.Println(tree.view())
		log.Println(errorNode.view())
		log.Panic("size error")
	}
	return
}

func (tree *Tree[KEY, VALUE]) view() (result string) {
	result = "\n"
	if tree.getRoot() == nil {
		result += "└── nil"
		return
	}
	var nmap = make(map[*Node[KEY, VALUE]]int)
	outputfordebug(nmap, tree.getRoot(), "", true, &result)
	return
}
