package indextree

func newIterator[KEY any, VALUE any](tree *Tree[KEY, VALUE]) *Iterator[KEY, VALUE] {
	return &Iterator[KEY, VALUE]{tree: tree}
}

func minimumNode[KEY any, VALUE any](node *hNode[KEY, VALUE]) *hNode[KEY, VALUE] {
	for node != nil && node.Children[0] != nil {
		node = node.Children[0]
	}
	return node
}

func maximumNode[KEY any, VALUE any](node *hNode[KEY, VALUE]) *hNode[KEY, VALUE] {
	for node != nil && node.Children[1] != nil {
		node = node.Children[1]
	}
	return node
}

func successorNode[KEY any, VALUE any](node *hNode[KEY, VALUE], sentinel *hNode[KEY, VALUE]) *hNode[KEY, VALUE] {
	if node == nil {
		return nil
	}

	if node.Children[1] != nil {
		return minimumNode(node.Children[1])
	}

	parent := node.Parent
	for parent != nil && parent != sentinel && parent.Children[1] == node {
		node = parent
		parent = parent.Parent
	}
	if parent == sentinel {
		return nil
	}
	return parent
}

func predecessorNode[KEY any, VALUE any](node *hNode[KEY, VALUE], sentinel *hNode[KEY, VALUE]) *hNode[KEY, VALUE] {
	if node == nil {
		return nil
	}

	if node.Children[0] != nil {
		return maximumNode(node.Children[0])
	}

	parent := node.Parent
	for parent != nil && parent != sentinel && parent.Children[0] == node {
		node = parent
		parent = parent.Parent
	}
	if parent == sentinel {
		return nil
	}
	return parent
}

func (iter *Iterator[KEY, VALUE]) seekEqual(key KEY, lessAndGreater int8) bool {
	current := iter.tree.getRoot()
	if current == nil {
		iter.cur = nil
		return false
	}

	pos := getSize(current.Children[0])
	var candidate *hNode[KEY, VALUE]
	var candidatePos int64
	exact := false

	for current != nil {
		cmp := iter.tree.compare(current.Key, key)
		switch {
		case cmp == 0:
			candidate = current
			candidatePos = pos
			exact = true
			current = nil
		case cmp < 0:
			if lessAndGreater == 0 {
				candidate = current
				candidatePos = pos
			}
			current = current.Children[1]
			if current != nil {
				pos += getSize(current.Children[0]) + 1
			}
		default:
			if lessAndGreater == 1 {
				candidate = current
				candidatePos = pos
			}
			current = current.Children[0]
			if current != nil {
				pos -= getSize(current.Children[1]) + 1
			}
		}
	}

	iter.cur = candidate
	if candidate != nil {
		iter.pos = candidatePos
	} else if lessAndGreater == 0 {
		iter.pos = -1
	} else {
		iter.pos = iter.tree.Size()
	}
	return exact
}

func (iter *Iterator[KEY, VALUE]) seekThan(key KEY, lessAndGreater int8) bool {
	current := iter.tree.getRoot()
	if current == nil {
		iter.cur = nil
		return false
	}

	pos := getSize(current.Children[0])
	var candidate *hNode[KEY, VALUE]
	var candidatePos int64
	exact := false

	for current != nil {
		cmp := iter.tree.compare(current.Key, key)
		switch {
		case cmp == 0:
			exact = true
			current = current.Children[lessAndGreater]
			if current != nil {
				if lessAndGreater == 0 {
					pos -= getSize(current.Children[1]) + 1
				} else {
					pos += getSize(current.Children[0]) + 1
				}
			}
		case cmp < 0:
			if lessAndGreater == 0 {
				candidate = current
				candidatePos = pos
			}
			current = current.Children[1]
			if current != nil {
				pos += getSize(current.Children[0]) + 1
			}
		default:
			if lessAndGreater == 1 {
				candidate = current
				candidatePos = pos
			}
			current = current.Children[0]
			if current != nil {
				pos -= getSize(current.Children[1]) + 1
			}
		}
	}

	iter.cur = candidate
	if candidate != nil {
		iter.pos = candidatePos
	} else if lessAndGreater == 0 {
		iter.pos = -1
	} else {
		iter.pos = iter.tree.Size()
	}
	return exact
}
