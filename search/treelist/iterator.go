package treelist

type Iterator struct {
	tree *Tree
	cur  *Node
	idx  int64
}

// SeekGE 搜索到 大于等于key 的前缀. 如果存在等值的key 返回true. 否则false
func (iter *Iterator) SeekGE(key []byte) bool {
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

// SeekGT 搜索到 大于key 的前缀. 如果存在等值的key 返回true. 否则false
func (iter *Iterator) SeekGT(key []byte) bool {
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

// SeekByIndex 位移到 有序序列的第index个
func (iter *Iterator) SeekByIndex(index int64) {
	cur := iter.tree.index(index)
	iter.idx = index
	iter.cur = cur
}

// SeekLE 搜索到 小于等于key 的前缀. 如果存在等值的key 返回true. 否则false
func (iter *Iterator) SeekLE(key []byte) bool {
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

// SeekLE 搜索到 小于key 的前缀. 如果存在等值的key 返回true. 否则false
func (iter *Iterator) SeekLT(key []byte) bool {
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
func (iter *Iterator) SeekToFirst() {
	const L = 0
	// root := iter.tree.getRoot()

	iter.cur = iter.tree.root.Direct[L]
	iter.idx = 0

}

// SeekToLast to the last item of the ordered sequence
func (iter *Iterator) SeekToLast() {
	const R = 1

	iter.cur = iter.tree.root.Direct[R]
	iter.idx = iter.tree.Size() - 1

}

// Valid 校验数据是否存在 配合Seek使用
func (iter *Iterator) Valid() bool {
	return iter.cur != nil
}

// Prev 位移至前一个
func (iter *Iterator) Prev() {
	const L = 0
	iter.cur = iter.cur.Direct[L]
	iter.idx--
}

// Next 位移至下一个
func (iter *Iterator) Next() {
	const R = 1
	iter.cur = iter.cur.Direct[R]
	iter.idx++
}

// Slice 当前item的key value
func (iter *Iterator) Slice() *Slice {
	return &iter.cur.Slice
}

// Index 当前item的index. 有序的位置 与数组等义
func (iter *Iterator) Index() int64 {
	return iter.idx
}

// Key 当前item的key
func (iter *Iterator) Key() []byte {
	return iter.cur.Key
}

// Value 当前item的value
func (iter *Iterator) Value() interface{} {
	return iter.cur.Value
}

// Clone 复制一个当前迭代的iterator. 用于复位
func (iter *Iterator) Clone() *Iterator {
	return &Iterator{tree: iter.tree, cur: iter.cur, idx: iter.idx}
}
