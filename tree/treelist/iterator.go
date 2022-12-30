package treelist

type hNode[KEY any, VALUE any] struct {
	cur *treeNode[KEY, VALUE]
	idx int64
}

// Iterator the iterator of treelist
type Iterator[KEY any, VALUE any] struct {
	tree *Tree[KEY, VALUE]
	// cur  *treeNode
	// idx  int64
	hNode[KEY, VALUE]
}

// SeekGE
// seek to Greater Than or Equal the key.
// if equal is not exists, take the great key
func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY) bool {
	const R = 1
	cur, idx, dir := iter.tree.seekNodeWithIndex(key)
	iter.idx = idx
	if dir > 0 {
		cur = cur.Direct[R]
		iter.idx++
	}
	iter.cur = cur
	return dir == 0
}

// SeekGT
// seek to Greater Than the key.
// take the great key
func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY) bool {
	const R = 1
	cur, idx, dir := iter.tree.seekNodeWithIndex(key)
	iter.idx = idx
	if dir >= 0 {
		cur = cur.Direct[R]
		iter.idx++
	}
	iter.cur = cur
	return dir == 0
}

// SeekByIndex
// seek to  the key by index. like index of array. index is ordered
func (iter *Iterator[KEY, VALUE]) SeekByIndex(index int64) {
	cur := iter.tree.index(index)
	iter.idx = index
	iter.cur = cur
}

// SeekLE
// seek to  less than or equal the key.
// if equal is not exists, take the less key
func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY) bool {
	const L = 0
	cur, idx, dir := iter.tree.seekNodeWithIndex(key)
	iter.idx = idx
	if dir < 0 {
		cur = cur.Direct[L]
		iter.idx--
	}
	iter.cur = cur
	return dir == 0
}

// SeekLT
// seek to  less than  the key.
func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY) bool {
	const L = 0
	cur, idx, dir := iter.tree.seekNodeWithIndex(key)
	iter.idx = idx
	if dir <= 0 {
		cur = cur.Direct[L]
		iter.idx--
	}
	iter.cur = cur
	return dir == 0
}

// SeekToFirst to the first item of the ordered sequence
func (iter *Iterator[KEY, VALUE]) SeekToFirst() {
	const L = 0
	// root := iter.tree.getRoot()

	iter.cur = iter.tree.root.Direct[L]
	iter.idx = 0

}

// SeekToLast to the last item of the ordered sequence
func (iter *Iterator[KEY, VALUE]) SeekToLast() {
	const R = 1

	iter.cur = iter.tree.root.Direct[R]
	iter.idx = iter.tree.Size() - 1

}

// Valid  if current value is not nil return true. else return false. for use with Seek
func (iter *Iterator[KEY, VALUE]) Valid() bool {
	return iter.cur != nil
}

// Prev  the current iterator move to the prev. before call it must call Vaild() and return true.
func (iter *Iterator[KEY, VALUE]) Prev() {
	const L = 0

	iter.idx--
	iter.cur = iter.cur.Direct[L]
}

// Compare iterator the  current value comare to key.
// if cur.key > key. return 1.
// if cur.key == key return 0.
// if cur.key < key return - 1.
func (iter *Iterator[KEY, VALUE]) Compare(key KEY) int {
	return iter.tree.compare(iter.cur.Key, key)
}

// Next Next the current iterator move to the next. before call it must call Vaild() and return true.
func (iter *Iterator[KEY, VALUE]) Next() {
	const R = 1
	iter.cur = iter.cur.Direct[R]
	iter.idx++
}

// Slice return the KeyValue of current
func (iter *Iterator[KEY, VALUE]) Slice() *Slice[KEY, VALUE] {
	return &iter.cur.Slice
}

// Index return the Index of the current iterator. Ordered position equivalent to the Index of an Priority Queue(Array)
func (iter *Iterator[KEY, VALUE]) Index() int64 {
	return iter.idx
}

// Key return the key of current
func (iter *Iterator[KEY, VALUE]) Key() KEY {
	return iter.cur.Key
}

// Value return the value of current
func (iter *Iterator[KEY, VALUE]) Value() VALUE {
	return iter.cur.Value
}

// Clone
// copy a iterator. eg: record iterator position
func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE] {
	return &Iterator[KEY, VALUE]{tree: iter.tree, hNode: hNode[KEY, VALUE]{cur: iter.cur, idx: iter.idx}}
}
