package avl

// Iterator tree iterator
type Iterator[KEY, VALUE any] struct {
	tree *Tree[KEY, VALUE]

	cur *Node[KEY, VALUE]

	idx   int8
	stack []NodeDir[KEY, VALUE]
}

type NodeDir[KEY any, VALUE any] struct {
	N *Node[KEY, VALUE]
	D int8
}

// Key return the key of current iterator
func (iter *Iterator[KEY, VALUE]) Key() KEY {
	return iter.cur.Key
}

// Value return the value of current iterator
func (iter *Iterator[KEY, VALUE]) Value() VALUE {
	return iter.cur.Value
}

// Valid if current value is not nil return true. else return false. for use with Seek
func (iter *Iterator[KEY, VALUE]) Valid() bool {
	return iter.cur != nil
}

// SeekToFirst seek to first item
func (iter *Iterator[KEY, VALUE]) SeekToFirst() {
	iter.cur = iter.tree.getRoot()
	iter.idx = -1
	for iter.down(0) {

	}
	// iter.up(1)
}

// SeekToFirst seek to last item
func (iter *Iterator[KEY, VALUE]) SeekToLast() {
	iter.cur = iter.tree.getRoot()
	iter.idx = -1
	for iter.down(1) {

	}
	// iter.up(0)
}

// SeekLE seek to the key that less than or equal to
func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY) bool {
	return iter.seekEqual(key, 0)
}

// SeekLT seek to the key that less than
func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY) bool {
	return iter.seekThan(key, 0)
}

// SeekGE seek to the key that greater than or equal to
func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY) bool {
	return iter.seekEqual(key, 1)
}

// SeekGT seek to the key that greater than
func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY) bool {
	return iter.seekThan(key, 1)
}

// Prev the current iterator move to the prev. before call it must call Vaild() and return true.
func (iter *Iterator[KEY, VALUE]) Prev() {
	iter.move(0)
}

// Next the current iterator move to the next. before call it must call Vaild() and return true.
func (iter *Iterator[KEY, VALUE]) Next() {
	iter.move(1)
}

// Clone Copy a current iterator
func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE] {
	other := newIterator(iter.tree)
	other.cur = iter.cur
	other.idx = iter.idx
	copy(other.stack, iter.stack)
	return other
}
