package rbtree

type Iterator[KEY any, VALUE any] struct {
	tree *Tree[KEY, VALUE]
	cur  *Node[KEY, VALUE]
}

func (iter *Iterator[KEY, VALUE]) Valid() bool {
	return iter.cur != nil
}

func (iter *Iterator[KEY, VALUE]) Key() KEY {
	return iter.cur.Key
}

func (iter *Iterator[KEY, VALUE]) Value() VALUE {
	return iter.cur.Value
}

func (iter *Iterator[KEY, VALUE]) SeekToFirst() {
	iter.cur = minimum(iter.tree.root)
}

func (iter *Iterator[KEY, VALUE]) SeekToLast() {
	iter.cur = maximum(iter.tree.root)
}

func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY) bool {
	current := iter.tree.root
	var candidate *Node[KEY, VALUE]
	exact := false

	for current != nil {
		cmp := iter.tree.compare(current.Key, key)
		if cmp < 0 {
			current = current.Right
			continue
		}

		candidate = current
		if cmp == 0 {
			exact = true
			break
		}
		current = current.Left
	}

	iter.cur = candidate
	return exact
}

func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY) bool {
	current := iter.tree.root
	var candidate *Node[KEY, VALUE]
	exact := false

	for current != nil {
		cmp := iter.tree.compare(current.Key, key)
		if cmp <= 0 {
			if cmp == 0 {
				exact = true
			}
			current = current.Right
			continue
		}

		candidate = current
		current = current.Left
	}

	iter.cur = candidate
	return exact
}

func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY) bool {
	current := iter.tree.root
	var candidate *Node[KEY, VALUE]
	exact := false

	for current != nil {
		cmp := iter.tree.compare(current.Key, key)
		if cmp > 0 {
			current = current.Left
			continue
		}

		candidate = current
		if cmp == 0 {
			exact = true
			break
		}
		current = current.Right
	}

	iter.cur = candidate
	return exact
}

func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY) bool {
	current := iter.tree.root
	var candidate *Node[KEY, VALUE]
	exact := false

	for current != nil {
		cmp := iter.tree.compare(current.Key, key)
		if cmp >= 0 {
			if cmp == 0 {
				exact = true
			}
			current = current.Left
			continue
		}

		candidate = current
		current = current.Right
	}

	iter.cur = candidate
	return exact
}

func (iter *Iterator[KEY, VALUE]) Next() {
	iter.cur = successor(iter.cur)
}

func (iter *Iterator[KEY, VALUE]) Prev() {
	iter.cur = predecessor(iter.cur)
}

func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE] {
	return &Iterator[KEY, VALUE]{tree: iter.tree, cur: iter.cur}
}