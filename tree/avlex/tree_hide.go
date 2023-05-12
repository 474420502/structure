package avlex

func (tree *Tree[KEY, VALUE]) put(parent *Node[KEY, VALUE], child int, key KEY) (target *Node[KEY, VALUE], isRepeat bool, isRebalance bool) {

	cur := parent.Children[child]
	if cur == nil {
		target = newNode[KEY, VALUE]()
		target.Key = key
		parent.Children[child] = target
		if parent.Children[^child+2] == nil {
			return target, false, true
		}
		return target, false, false
	}

	cmp := tree.Compare(cur.Key, key)

	if cmp < 0 {
		cur.Key = key
		return cur, true, false
	} else {
		target, isRepeat, isRebalance = tree.put(cur, cmp, key)

		if isRepeat || !isRebalance {
			return target, isRepeat, isRebalance
		}
	}

	if isRebalance {
		isRebalance = cur.rebalance(parent, child)
	}

	return target, false, isRebalance
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

func (tree *Tree[KEY, VALUE]) remove(key KEY, grandpa *Node[KEY, VALUE], child2, child1 int) *Node[KEY, VALUE] {
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
			return cur
		}

		if cur.Children[1] == nil {
			parent.Children[child1] = cur.Children[0]
			return cur
		}

		replace := tree.moveNode(cur, child1, ^child1+2)

		parent.Children[child1] = replace
		replace.Children = cur.Children
		replace.rebalance(parent, child1)

		return cur
	}

	result := tree.remove(key, parent, child1, cmp)
	if cur != tree.Center {
		cur.rebalance(parent, child1)
	}

	return result
}

func (tree *Tree[KEY, VALUE]) moveNode(parent *Node[KEY, VALUE], child2, child1 int) *Node[KEY, VALUE] {
	cur := parent.Children[child2]
	sub := cur.Children[child1]

	if sub == nil {
		other := cur.Children[^child1+2]
		parent.Children[child2] = other
		return cur
	}

	result := tree.moveNode(cur, child1, child1)
	cur.rebalance(parent, child2)
	return result
}
