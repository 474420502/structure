package experiment

import "log"

func (tree *Tree[KEY, VALUE]) getRoot() *Node[KEY, VALUE] {
	return tree.Center.Children[1]
}

func (tree *Tree[KEY, VALUE]) fixPutSize(cur *Node[KEY, VALUE]) {
	for cur != tree.Center {
		cur.updateSize()
		cur = cur.Parent
	}
}

func (tree *Tree[KEY, VALUE]) put(parent *Node[KEY, VALUE], child int, key KEY) (target *Node[KEY, VALUE], isExists bool, isRebalance bool) {

	cur := parent.Children[child]
	if cur == nil {
		target = newNode[KEY, VALUE]()
		target.Key = key
		target.Parent = parent
		parent.Children[child] = target
		tree.fixPutSize(parent)
		if parent.Children[^child+2] == nil {
			return target, false, true
		}
		return target, false, false
	}

	cmp := tree.Compare(cur.Key, key)

	if cmp < 0 {
		target, isExists, isRebalance = tree.put(cur, 0, key)
		if isExists || !isRebalance {
			return target, isExists, isRebalance
		}
	} else if cmp > 0 {
		target, isExists, isRebalance = tree.put(cur, 1, key)
		if isExists || !isRebalance {
			return target, isExists, isRebalance
		}
	} else {
		return cur, true, false
	}

	if isRebalance {
		isRebalance = tree.rebalance(parent, child)
		if isRebalance {
			tree.fixPutSize(parent)
		}
	}

	return target, isExists, isRebalance
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

func (tree *Tree[KEY, VALUE]) remove(key KEY, grandpa *Node[KEY, VALUE], child2, child1 int) (target *VALUE, isRebalance bool) {
	parent := grandpa.Children[child2]
	cur := parent.Children[child1]

	if cur == nil {
		return nil, false
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {

		// remove 两种状态. 当前值不在底, 在底
		if cur.Children[0] == nil {
			parent.Children[child1] = cur.Children[1]
			return &cur.Value, true
		}

		if cur.Children[1] == nil {
			parent.Children[child1] = cur.Children[0]
			return &cur.Value, true
		}

		var result = cur.Value
		target = &result

		replacer, _ := tree.neighboring(cur, child1, ^child1+2)
		cur.Key = replacer.Key
		cur.Value = replacer.Value

		return target, tree.rebalance(parent, child1)
	}

	target, isRebalance = tree.remove(key, parent, child1, cmp)
	if cur != tree.Center && isRebalance {
		isRebalance = tree.rebalance(parent, child1)
	}

	return target, isRebalance
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
	if isRebalance {
		isRebalance = tree.rebalance(parent, child2)
	}

	return result, isRebalance
}

func (tree *Tree[KEY, VALUE]) rebalance(parent *Node[KEY, VALUE], child int) bool {

	node := parent.Children[child]
	lh, rh := getHeight(node.Children[0]), getHeight(node.Children[1])

	diff := lh - rh

	if diff >= tree.differenceHeight {
		sub := node.Children[0]
		if getHeight(sub.Children[1]) > getHeight(sub.Children[0]) {
			rightRotateWithLeft(tree, parent, child)
		} else {
			rightRotate(tree, parent, child)
		}
		return true
	} else if diff <= -tree.differenceHeight {
		sub := node.Children[1]
		if getHeight(sub.Children[0]) > getHeight(sub.Children[1]) {
			leftRotateWithRight(tree, parent, child)
		} else {
			leftRotate(tree, parent, child)
		}
		return true
	} else {
		if lh > rh {
			if node.Height != lh+1 {
				node.Height = lh + 1
				return true
			}
		} else {
			if node.Height != rh+1 {
				node.Height = rh + 1
				return true
			}
		}

	}

	return false
}

func (tree *Tree[KEY, VALUE]) check() (result string) {
	if !tree.checkHeightTree(tree.getRoot()) {
		log.Panic("height error")
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
