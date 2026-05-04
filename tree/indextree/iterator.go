package indextree

type Iterator[KEY any, VALUE any] struct {
	tree *Tree[KEY, VALUE]

	cur *hNode[KEY, VALUE]

	pos   int64
}

func (iter *Iterator[KEY, VALUE]) Key() KEY {
	return iter.cur.Key
}

func (iter *Iterator[KEY, VALUE]) Value() VALUE {
	return iter.cur.Value
}

func (iter *Iterator[KEY, VALUE]) Valid() bool {
	return iter.cur != nil
}

func (iter *Iterator[KEY, VALUE]) Index() int64 {
	return iter.pos
}

func (iter *Iterator[KEY, VALUE]) SeekByIndex(index int64) {
	if index < 0 || index >= iter.tree.Size() {
		iter.cur = nil
		return
	}

	iter.cur = iter.tree.index(index)
	iter.pos = index
}

func (iter *Iterator[KEY, VALUE]) SeekToFirst() {
	iter.cur = minimumNode(iter.tree.getRoot())
	iter.pos = 0
}

func (iter *Iterator[KEY, VALUE]) SeekToLast() {
	iter.cur = maximumNode(iter.tree.getRoot())
	iter.pos = iter.tree.Size() - 1
}

func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY) bool {
	return iter.seekEqual(key, 0)
}

func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY) bool {
	return iter.seekThan(key, 0)
}

func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY) bool {
	return iter.seekEqual(key, 1)
}

func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY) bool {
	return iter.seekThan(key, 1)
}

func (iter *Iterator[KEY, VALUE]) Prev() {
	if iter.cur == nil {
		return
	}
	iter.cur = predecessorNode(iter.cur, iter.tree.root)
	iter.pos--
}

func (iter *Iterator[KEY, VALUE]) Next() {
	if iter.cur == nil {
		return
	}
	iter.cur = successorNode(iter.cur, iter.tree.root)
	iter.pos++
}

func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE] {
	return &Iterator[KEY, VALUE]{tree: iter.tree, cur: iter.cur, pos: iter.pos}
}
