package avlex

import "log"

func (tree *Tree[KEY, VALUE]) put(parent *Node[KEY, VALUE], child int, key KEY) (target *Node[KEY, VALUE], isExists bool, isRebalance bool) {

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
		return cur, true, false
	} else {
		target, isExists, isRebalance = tree.put(cur, cmp, key)
		if isExists || !isRebalance {
			return target, isExists, isRebalance
		}
	}

	if isRebalance {
		isRebalance = cur.rebalance(parent, child)
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

		target = &cur.Value
		replacer, _ := tree.neighboring(cur, child1, ^child1+2)
		cur.Key = replacer.Key
		cur.Value = replacer.Value

		return target, cur.rebalance(parent, child1)
	}

	target, isRebalance = tree.remove(key, parent, child1, cmp)
	if cur != tree.Center && isRebalance {
		isRebalance = cur.rebalance(parent, child1)
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
		isRebalance = cur.rebalance(parent, child2)
	}

	return result, isRebalance
}

func (tree *Tree[KEY, VALUE]) check() (result string) {
	if !tree.checkHeightTree(tree.Center.Children[1]) {
		log.Panic("height error")
	}
	return
}

func (tree *Tree[KEY, VALUE]) view() (result string) {
	result = "\n"
	if tree.Center.Children[1] == nil {
		result += "└── nil"
		return
	}
	var nmap = make(map[*Node[KEY, VALUE]]int)
	outputfordebug(nmap, tree.Center.Children[1], "", true, &result)
	return
}
