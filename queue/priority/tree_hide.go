package treequeue

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
		cmp = 1
	}

	target, isExists = tree.put(cur, cmp, key)
	if isExists == 1 {
		return target, isExists
	}

	cur.Size++
	if cur.Size > 2 && isExists != 2 {
		tree.rebalance(parent, child)
	}

	return target, isExists
}

// 获取到根节点后遍历
func (tree *Tree[KEY, VALUE]) get(key KEY, cur *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	if cur == nil {
		return nil
	}
	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		// cmp = 0
		return tree.get(key, cur)
	}

	return tree.get(key, cur.Children[cmp])
}

func (tree *Tree[KEY, VALUE]) getfirst(key KEY, parent *Node[KEY, VALUE], child int, sel *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	cur := parent.Children[child]
	if cur == nil {
		return sel
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		cmp = 0
		sel = cur
	}
	return tree.getfirst(key, cur, cmp, sel)
}

func (tree *Tree[KEY, VALUE]) index(i int) *Node[KEY, VALUE] {

	cur := tree.getRoot()
	var idx int = cur.Children[0].getSize()
	for {
		if idx > i {
			if cur.Children[0] == nil {
				log.Println(tree.view(), cur.view())
			}
			cur = cur.Children[0]
			idx -= cur.Children[1].getSize() + 1
		} else if idx < i {
			if cur.Children[1] == nil {
				log.Println(tree.view(), cur.view())
			}
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
		// cur = cur.Children[0]
		idx -= cur.Children[0].Children[1].getSize() + 1
		return tree.removeIndex(cur, 0, idx, i)
	} else {
		// cur = cur.Children[1]
		idx += cur.Children[1].Children[0].getSize() + 1
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

func (tree *Tree[KEY, VALUE]) trimLow(cur *Node[KEY, VALUE], key KEY) *Node[KEY, VALUE] {

	if cur == nil {
		return nil
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		cur.Children[0] = nil
		cur.updateSize()
		return cur
	}

	if cmp == 1 {
		cur = tree.trimLow(cur.Children[1], key)
	} else {
		cur.Children[0] = tree.trimLow(cur.Children[0], key)
	}
	if cur != nil {
		cur.updateSize()
	}
	return cur
}

func (tree *Tree[KEY, VALUE]) trimHigh(cur *Node[KEY, VALUE], key KEY) *Node[KEY, VALUE] {

	if cur == nil {
		return nil
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		cur.Children[1] = nil
		cur.updateSize()
		return cur
	}

	if cmp == 0 {
		cur = tree.trimHigh(cur.Children[0], key)
	} else {
		cur.Children[1] = tree.trimHigh(cur.Children[1], key)
	}
	if cur != nil {
		cur.updateSize()
	}
	return cur
}

func (tree *Tree[KEY, VALUE]) seekRangeRoot(cur *Node[KEY, VALUE], low, high KEY) *Node[KEY, VALUE] {

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

	return tree.seekRangeRoot(cur.Children[cmplow], low, high)
}

func (tree *Tree[KEY, VALUE]) seekTrimIndexRoot(cur *Node[KEY, VALUE], idx, low, high int) (*Node[KEY, VALUE], int) {

	if idx > low && idx > high {
		cur = cur.Children[0]
		idx -= cur.Children[1].getSize() + 1
		return tree.seekTrimIndexRoot(cur, idx, low, high)
	} else if idx < low && idx < high {
		cur = cur.Children[1]
		idx += cur.Children[0].getSize() + 1
		return tree.seekTrimIndexRoot(cur, idx, low, high)
	}
	return cur, idx
}

func (tree *Tree[KEY, VALUE]) trimIndexLow(parent *Node[KEY, VALUE], child, idx int, low int) *Node[KEY, VALUE] {

	cur := parent.Children[child]
	if cur == nil {
		return nil
	}

	if child == 0 {
		idx -= cur.Children[1].getSize() + 1
	} else {
		idx += cur.Children[0].getSize() + 1
	}

	if idx == low {
		cur.Children[0] = nil
		cur.updateSize()
		return cur
	}

	if idx > low {
		cur.Children[0] = tree.trimIndexLow(cur, 0, idx, low)
	} else {
		cur = tree.trimIndexLow(cur, 1, idx, low)
	}
	if cur != nil {
		cur.updateSize()
	}
	return cur
}

func (tree *Tree[KEY, VALUE]) trimIndexHigh(parent *Node[KEY, VALUE], child, idx int, high int) *Node[KEY, VALUE] {

	cur := parent.Children[child]
	if cur == nil {
		return nil
	}

	if child == 0 {
		idx -= cur.Children[1].getSize() + 1
	} else {
		idx += cur.Children[0].getSize() + 1
	}

	if idx == high {
		cur.Children[1] = nil
		cur.updateSize()
		return cur
	}

	if idx < high {
		cur.Children[1] = tree.trimIndexHigh(cur, 1, idx, high)
	} else {
		cur = tree.trimIndexHigh(cur, 0, idx, high)
	}
	if cur != nil {
		cur.updateSize()
	}
	return cur
}

// 收集低维度节点数据
func (tree *Tree[KEY, VALUE]) removeCollectLows(collect *[]*Node[KEY, VALUE], cur *Node[KEY, VALUE], low KEY) {
	if cur == nil {
		// *collect = append(*collect, cur)
		return
	}

	cmp := tree.Compare(cur.Key, low)
	if cmp < 0 {
		tree.removeCollectLows(collect, cur.Children[0], low)
		return
	}

	if cmp == 1 {
		*collect = append(*collect, cur)
	}
	tree.removeCollectLows(collect, cur.Children[cmp], low)
}

func (tree *Tree[KEY, VALUE]) removeCollectHighs(collect *[]*Node[KEY, VALUE], cur *Node[KEY, VALUE], high KEY) {
	if cur == nil {
		// *collect = append(*collect, cur)
		return
	}

	cmp := tree.Compare(cur.Key, high)
	if cmp < 0 {
		tree.removeCollectHighs(collect, cur.Children[1], high)
		return
	}

	if cmp == 0 {
		*collect = append(*collect, cur)
	}
	tree.removeCollectHighs(collect, cur.Children[cmp], high)
}

// megreThreshold 合并收集的范围节点
func (tree *Tree[KEY, VALUE]) megreThreshold(l []*Node[KEY, VALUE], idx int, t int) *Node[KEY, VALUE] {
	if idx >= len(l) {
		return nil
	}
	cur := l[idx]
	cur.Children[t] = tree.megreThreshold(l, idx+1, t)
	cur.updateSize()
	return cur
}

// leftMegreRight 左边值集合 合并 右值集合
func (tree *Tree[KEY, VALUE]) leftMegreRight(parent *Node[KEY, VALUE], right *Node[KEY, VALUE]) {
	left := parent.Children[1]
	if left.Children[1] == nil {
		left.Children[1] = right
		left.Size += right.Size
		return
	}

	tree.leftMegreRight(left, right)
	left.updateSize()
	tree.rebalance(parent, 1)
}

func (tree *Tree[KEY, VALUE]) removeCollectIndexLows(collect *[]*Node[KEY, VALUE], cur *Node[KEY, VALUE], idx int, low int) {

	if idx < low {
		*collect = append(*collect, cur)
		cur = cur.Children[1]
		if cur == nil {
			return
		}
		tree.removeCollectIndexLows(collect, cur, idx+cur.Children[0].getSize()+1, low)
	} else {
		cur = cur.Children[0]
		if cur == nil {
			return
		}
		tree.removeCollectIndexLows(collect, cur, idx-cur.Children[1].getSize()-1, low)
	}

}

func (tree *Tree[KEY, VALUE]) removeCollectIndexHighs(collect *[]*Node[KEY, VALUE], cur *Node[KEY, VALUE], idx int, high int) {

	if idx > high {
		*collect = append(*collect, cur)
		cur = cur.Children[0]
		if cur == nil {
			return
		}
		tree.removeCollectIndexHighs(collect, cur, idx-cur.Children[1].getSize()-1, high)
	} else {

		cur = cur.Children[1]
		if cur == nil {
			return
		}
		tree.removeCollectIndexHighs(collect, cur, idx+cur.Children[0].getSize()+1, high)
	}

}

func (tree *Tree[KEY, VALUE]) split(left, right *[]*Node[KEY, VALUE], cur *Node[KEY, VALUE], key KEY) {
	if cur == nil {
		return
	}
	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		*left = append(*left, cur)
		tree.split(left, right, cur.Children[1], key)
		return
	}
	if cmp == 0 {
		*right = append(*right, cur)
		tree.split(left, right, cur.Children[0], key)
	} else {
		*left = append(*left, cur)
		tree.split(left, right, cur.Children[1], key)
	}
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
	outputfordebug(nmap, tree.getRoot(), "", true, &result, nil)
	return
}

func (tree *Tree[KEY, VALUE]) viewEx(taregt *Node[KEY, VALUE]) (result string) {
	result = "\n"
	if tree.getRoot() == nil {
		result += "└── nil"
		return
	}
	var nmap = make(map[*Node[KEY, VALUE]]int)
	outputfordebug(nmap, tree.getRoot(), "", true, &result, taregt)
	return
}
