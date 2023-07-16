package treeset

// TreeSet
// type TreeSet[KEY, VALUE any] struct {
// 	tree *avl.Tree[KEY, VALUE]
// }

// // New
// func New[KEY, VALUE any](Compare compare.Compare[KEY]) *TreeSet[KEY, VALUE] {
// 	return &TreeSet[KEY, VALUE]{tree: avl.New[KEY, VALUE](Compare)}
// }

// Add Not Cover the key of node. if item exists. return false
// func (set *TreeSet[KEY, VALUE]) Add(item KEY, value VALUE) bool {
// 	return set.tree.Put(item, value)

// }

// Set   Set the key of node.
//
// like add(). if compare is special, will cover key.
// eg. k1 = {a:1,b:2}. k2 = {a:1, b:3}.  k1.a == k2.a. the key will be covered by k2
// func (set *TreeSet[KEY, VALUE]) Set(item KEY, value VALUE) bool {
// 	return set.tree.Set(item, value)
// }

// Remove
// func (set *TreeSet[KEY, VALUE]) Remove(item KEY) {
// 	set.tree.Remove(item)
// }

// Values
// func (set *TreeSet[KEY, VALUE]) Values() []VALUE {
// 	return set.tree.Values()
// }

// Contains
// func (set *TreeSet[KEY, VALUE]) Contains(item KEY) bool {
// 	if _, ok := set.tree.Get(item); ok {
// 		return true
// 	}
// 	return false
// }

// Empty
// func (set *TreeSet[KEY, VALUE]) Empty() bool {
// 	return set.Size() == 0
// }

// Clear
// func (set *TreeSet[KEY, VALUE]) Clear() {
// 	set.tree.Clear()
// }

// Size
// func (set *TreeSet[KEY, VALUE]) Size() uint {
// 	return set.tree.Size()
// }

// Iterator avl Iterator
// func (set *TreeSet[KEY, VALUE]) Iterator() *Iterator[KEY] {
// 	return set.tree.Iterator()
// }

// Traverse 从左到右遍历. left -> right
// func (set *TreeSet[KEY, VALUE]) Traverse(tr func(v interface{}) bool) {
// 	set.tree.Traverse(tr)
// }
