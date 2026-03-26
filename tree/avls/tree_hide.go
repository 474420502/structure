package avls

import "log"

func (tree *Tree[KEY, VALUE]) getRoot() *Node[KEY, VALUE] {
	return tree.Center.Children[1]
}

func (tree *Tree[KEY, VALUE]) put(parent *Node[KEY, VALUE], child int, key KEY) (target *Node[KEY, VALUE], isRebalance bool) {

	cur := parent.Children[child]
	if cur == nil {
		target = newNode[KEY, VALUE]()
		target.Key = key
		parent.Children[child] = target
		if parent.Children[^child+2] == nil {
			return target, true
		}
		return target, false
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		cmp = 1
	}

	target, isRebalance = tree.put(cur, cmp, key)

	if isRebalance {
		isRebalance = tree.rebalance(parent, child)
	}

	return target, isRebalance
}

func (tree *Tree[KEY, VALUE]) get(key KEY, cur *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	if cur == nil {
		return nil
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		if left := tree.get(key, cur.Children[0]); left != nil {
			return left
		}
		return cur
	}

	return tree.get(key, cur.Children[cmp])
}

func (tree *Tree[KEY, VALUE]) removeLeft(key KEY, grandpa *Node[KEY, VALUE], child2, child1 int) (target *VALUE, isRebalance bool) {
	parent := grandpa.Children[child2]
	cur := parent.Children[child1]

	if cur == nil {
		return nil, false
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		target, isRebalance = tree.removeLeft(key, parent, child1, 0)
		if target != nil {
			if cur != tree.Center && isRebalance {
				isRebalance = tree.rebalance(parent, child1)
			}
			return target, isRebalance
		}

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

	target, isRebalance = tree.removeLeft(key, parent, child1, cmp)
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
		// tree.rotateCount++
		if getHeight(sub.Children[1]) > getHeight(sub.Children[0]) {
			rightRotateWithLeft(parent, child)
		} else {
			rightRotate(parent, child)
		}
		return true
	} else if diff <= -tree.differenceHeight {
		sub := node.Children[1]
		// tree.rotateCount++
		if getHeight(sub.Children[0]) > getHeight(sub.Children[1]) {
			leftRotateWithRight(parent, child)
		} else {
			leftRotate(parent, child)
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
