package treelist

type Iterator struct {
	tree *Tree
	cur  *Node
	idx  int64
}

func (iter *Iterator) Seek(key []byte) {
	iter.cur, iter.idx = iter.tree.getNodeWithIndex(key)
}

func (iter *Iterator) Prev() bool {
	const L = 0
	prev := iter.cur.Direct[L]
	if prev != nil {
		iter.cur = prev
		iter.idx--
		return true
	}
	return false
}

func (iter *Iterator) Next() bool {
	const R = 1
	next := iter.cur.Direct[R]
	if next != nil {
		iter.cur = next
		iter.idx++
		return true
	}
	return false
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
