package indextree

type Iterator[T any] struct {
	tree *Tree[T]

	cur *hNode[T]

	idx   int8
	pos   int64
	stack []nodeDir[T]
}

type nodeDir[T any] struct {
	N *hNode[T]
	D int8
}

func (iter *Iterator[T]) Key() T {
	return iter.cur.Key
}

func (iter *Iterator[T]) Value() interface{} {
	return iter.cur.Value
}

func (iter *Iterator[T]) Valid() bool {
	return iter.cur != nil
}

func (iter *Iterator[T]) Index() int64 {
	return iter.pos
}

func (iter *Iterator[T]) SeekByIndex(index int64) {
	if index < 0 || index >= iter.tree.Size() {
		iter.cur = nil
		return
	}

	iter.cur = iter.tree.index(index)
	iter.pos = index
}

func (iter *Iterator[T]) SeekToFirst() {
	iter.cur = iter.tree.getRoot()
	iter.idx = -1
	iter.pos = 0
	for iter.down(0) {

	}
}

func (iter *Iterator[T]) SeekToLast() {
	iter.cur = iter.tree.getRoot()
	iter.idx = -1
	iter.pos = iter.tree.Size() - 1
	for iter.down(1) {

	}
}

func (iter *Iterator[T]) SeekLE(key T) bool {
	return iter.seekEqual(key, 0)
}

func (iter *Iterator[T]) SeekLT(key T) bool {
	return iter.seekThan(key, 0)
}

func (iter *Iterator[T]) SeekGE(key T) bool {
	return iter.seekEqual(key, 1)
}

func (iter *Iterator[T]) SeekGT(key T) bool {
	return iter.seekThan(key, 1)
}

func (iter *Iterator[T]) Prev() {
	iter.move(0)
	iter.pos--
}

func (iter *Iterator[T]) Next() {
	iter.move(1)
	iter.pos++
}

func (iter *Iterator[T]) Clone() *Iterator[T] {
	other := newIterator(iter.tree)
	other.cur = iter.cur
	other.idx = iter.idx
	other.pos = iter.pos
	copy(other.stack, iter.stack)
	return other
}
