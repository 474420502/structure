package treelist

type Iterator struct {
	tree *Tree
	cur  *Node
	idx  int64
}

func (iter *Iterator) Seek(key []byte) {
	const R = 1
	cur, idx, dir := iter.tree.seekNodeWithIndex(key)
	iter.idx = idx
	if dir > 0 {
		cur = cur.Direct[R]
		iter.idx++
	}
	iter.cur = cur
}

func (iter *Iterator) SeekToPrev(key []byte) {
	const L = 1
	cur, idx, dir := iter.tree.seekNodeWithIndex(key)
	iter.idx = idx
	if dir < 0 {
		cur = cur.Direct[L]
		iter.idx--
	}
	iter.cur = cur
}

func (iter *Iterator) SeekToFirst() {
	const L = 0
	iter.cur = iter.tree.root.Direct[L]
	iter.idx = 0
}

func (iter *Iterator) SeekToLast() {
	const R = 1
	iter.cur = iter.tree.root.Direct[R]
	iter.idx = iter.tree.Size() - 1
}

func (iter *Iterator) Valid() bool {
	return iter.cur != nil
}

func (iter *Iterator) Prev() {
	const L = 0

	iter.cur = iter.cur.Direct[L]
	iter.idx--

}

func (iter *Iterator) Next() {
	const R = 1
	iter.cur = iter.cur.Direct[R]
	iter.idx++
}

func (iter *Iterator) Slice() *Slice {
	return &iter.cur.Slice
}

func (iter *Iterator) Index() int64 {
	return iter.idx
}

func (iter *Iterator) Key() []byte {
	return iter.cur.Key
}

func (iter *Iterator) Value() []byte {
	return iter.cur.Value
}
