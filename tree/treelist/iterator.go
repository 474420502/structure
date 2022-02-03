package treelist

type nodePoint[T any] struct {
	cur *treeNode[T]
	idx int64
}

type Iterator[T any] struct {
	tree *Tree[T]
	// cur  *treeNode
	// idx  int64
	nodePoint[T]
}

// SeekGE 搜索到 大于等于key 的前缀. 如果存在等值的key 返回true. 否则false
//
// seek to Greater Than or Equal the key.
// [less equal greater] --> if equal is not exists, take the great
func (iter *Iterator[T]) SeekGE(key T) bool {
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
//
// seek to Greater Than the key.
// [less equal greater] -->  take the great
func (iter *Iterator[T]) SeekGT(key T) bool {
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
//
// seek to  the key by index. like index of array. index is ordered
func (iter *Iterator[T]) SeekByIndex(index int64) {
	cur := iter.tree.index(index)
	iter.idx = index
	iter.cur = cur
}

// SeekLE 搜索到 小于等于key 的前缀. 如果存在等值的key 返回true. 否则false
//
// seek to  less than or equal the key.
// [less equal greater] -->  if equal is not exists, take the less
func (iter *Iterator[T]) SeekLE(key T) bool {
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
//
// seek to  less than  the key.
// [less equal greater] --> take the less
func (iter *Iterator[T]) SeekLT(key T) bool {
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
func (iter *Iterator[T]) SeekToFirst() {
	const L = 0
	// root := iter.tree.getRoot()

	iter.cur = iter.tree.root.Direct[L]
	iter.idx = 0

}

// SeekToLast to the last item of the ordered sequence
func (iter *Iterator[T]) SeekToLast() {
	const R = 1

	iter.cur = iter.tree.root.Direct[R]
	iter.idx = iter.tree.Size() - 1

}

// Valid 校验数据是否存在 配合Seek使用
func (iter *Iterator[T]) Valid() bool {
	return iter.cur != nil
}

// Prev 位移至前一个
func (iter *Iterator[T]) Prev() {
	const L = 0

	iter.idx--
	iter.cur = iter.cur.Direct[L]
}

// Compare iterator the  current value comare to key.
//
// if cur.key > key. return 1.
//
// if cur.key == key return 0.
//
// if cur.key < key return - 1.
func (iter *Iterator[T]) Compare(key T) int {
	return iter.tree.compare(iter.cur.Key, key)
}

// Next 位移至下一个
func (iter *Iterator[T]) Next() {
	const R = 1
	iter.cur = iter.cur.Direct[R]
	iter.idx++
}

// Slice 当前item的key value
func (iter *Iterator[T]) Slice() *Slice[T] {
	return &iter.cur.Slice
}

// Index 当前item的index. 有序的位置 与数组等义
func (iter *Iterator[T]) Index() int64 {
	return iter.idx
}

// Key 当前item的key
func (iter *Iterator[T]) Key() T {
	return iter.cur.Key
}

// Value 当前item的value
func (iter *Iterator[T]) Value() interface{} {
	return iter.cur.Value
}

// Clone 复制一个当前迭代的iterator. 用于复位
//
// copy a iterator. eg: record iterator position
func (iter *Iterator[T]) Clone() *Iterator[T] {
	return &Iterator[T]{tree: iter.tree, nodePoint: nodePoint[T]{cur: iter.cur, idx: iter.idx}}
}
