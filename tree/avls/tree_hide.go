package avls

import "log"

func (tree *Tree[KEY, VALUE]) getRoot() *Node[KEY, VALUE] {
	return tree.Center.Children[1]
}

func (tree *Tree[KEY, VALUE]) put(parent *Node[KEY, VALUE], child int, key KEY) (target *Node[KEY, VALUE], isRebalance bool) {

	cur := parent.Children[child]
	if cur == nil {
		if tree.free != nil {
			target = tree.free
			tree.free = target.Children[0]
		} else {
			target = &Node[KEY, VALUE]{}
		}
		target.Key = key
		target.Height = 1
		target.Children[0] = nil
		target.Children[1] = nil
		parent.Children[child] = target
		if parent.Children[^child+2] == nil {
			return target, true
		}
		return target, false
	}

	cmp := tree.Compare(cur.Key, key)
	dir := 0
	if cmp <= 0 {
		dir = 1
	}

	target, isRebalance = tree.put(cur, dir, key)

	if isRebalance {
		isRebalance = tree.rebalance(parent, child)
	}

	return target, isRebalance
}

func (tree *Tree[KEY, VALUE]) get(key KEY, cur *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	// Iterative version that returns the left-most equal key.
	var found *Node[KEY, VALUE]
	for cur != nil {
		cmp := tree.Compare(cur.Key, key)
		if cmp == 0 {
			found = cur
			cur = cur.Children[0]
			continue
		}
		if cmp < 0 {
			cur = cur.Children[1]
		} else {
			cur = cur.Children[0]
		}
	}
	return found
}

// recycle resets a detached node and pushes it onto the free list.
func (tree *Tree[KEY, VALUE]) recycle(node *Node[KEY, VALUE]) {
	var zeroKey KEY
	var zeroValue VALUE
	node.Key = zeroKey
	node.Value = zeroValue
	node.Height = 0
	node.Children[1] = nil
	node.Children[0] = tree.free
	tree.free = node
}

func (tree *Tree[KEY, VALUE]) removeLeft(key KEY, grandpa *Node[KEY, VALUE], child2, child1 int) (target VALUE, found, isRebalance bool) {
	parent := grandpa.Children[child2]
	cur := parent.Children[child1]

	if cur == nil {
		return tree.zero, false, false
	}

	cmp := tree.Compare(cur.Key, key)
	if cmp == 0 {
		target, found, isRebalance = tree.removeLeft(key, parent, child1, 0)
		if found {
			if cur != tree.Center && isRebalance {
				isRebalance = tree.rebalance(parent, child1)
			}
			return target, true, isRebalance
		}

		if cur.Children[0] == nil {
			parent.Children[child1] = cur.Children[1]
			target = cur.Value
			tree.recycle(cur)
			return target, true, true
		}

		if cur.Children[1] == nil {
			parent.Children[child1] = cur.Children[0]
			target = cur.Value
			tree.recycle(cur)
			return target, true, true
		}

		target = cur.Value
		found = true

		replacer, _ := tree.neighboring(cur, child1, ^child1+2)
		cur.Key = replacer.Key
		cur.Value = replacer.Value
		isRebalance = tree.rebalance(parent, child1)
		tree.recycle(replacer)

		return target, true, isRebalance
	}

	dir := 0
	if cmp < 0 {
		dir = 1
	}

	target, found, isRebalance = tree.removeLeft(key, parent, child1, dir)
	if !found {
		return tree.zero, false, false
	}
	if cur != tree.Center && isRebalance {
		isRebalance = tree.rebalance(parent, child1)
	}

	return target, true, isRebalance
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
