package orderedmap

import (
	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/indextree"
)

type OrderedMap[K any, V any] struct {
	tree *indextree.Tree[K, V]
}

// NewWith builds a container from a comparator and a prototype value. The
// prototype is not stored; it only lets the compiler infer VALUE so callers
// name the key type once:
//
//	t := orderedmap.NewWith(compare.Any[int], "")
func NewWith[K any, V any](cmp compare.Compare[K], _ V) *OrderedMap[K, V] {
	return New[K, V](cmp)
}

func New[K any, V any](comp compare.Compare[K]) *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		tree: indextree.NewWithValue[K, V](comp),
	}
}

func (m *OrderedMap[K, V]) Size() int64 {
	return m.tree.Size()
}

// Len returns the number of elements.
func (m *OrderedMap[K, V]) Len() int {
	return int(m.tree.Size())
}

func (m *OrderedMap[K, V]) IsEmpty() bool {
	return m.tree.Size() == 0
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	return m.tree.Get(key)
}

func (m *OrderedMap[K, V]) Put(key K, value V) bool {
	return m.tree.Put(key, value)
}

// InsertIfAbsent inserts a value only when the key does not exist.
func (m *OrderedMap[K, V]) InsertIfAbsent(key K, value V) bool {
	return m.tree.InsertIfAbsent(key, value)
}

func (m *OrderedMap[K, V]) Set(key K, value V) bool {
	return m.tree.Set(key, value)
}

// Upsert sets the value and reports whether an existing value was replaced.
func (m *OrderedMap[K, V]) Upsert(key K, value V) bool {
	return m.tree.Upsert(key, value)
}

func (m *OrderedMap[K, V]) Remove(key K) (V, bool) {
	return m.tree.Remove(key)
}

// Delete removes a key and returns the previous value when present.
func (m *OrderedMap[K, V]) Delete(key K) (V, bool) {
	return m.tree.Delete(key)
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	_, ok := m.tree.Get(key)
	return ok
}

func (m *OrderedMap[K, V]) IndexOf(key K) int64 {
	return m.tree.IndexOf(key)
}

func (m *OrderedMap[K, V]) Index(index int64) (K, V) {
	return m.tree.Index(index)
}

func (m *OrderedMap[K, V]) RemoveIndex(index int64) (K, V, bool) {
	if index < 0 || index >= m.tree.Size() {
		return *new(K), *new(V), false
	}
	k, _ := m.tree.Index(index)
	v, ok := m.tree.RemoveIndex(index)
	if !ok {
		return *new(K), *new(V), false
	}
	return k, v, true
}

func (m *OrderedMap[K, V]) Keys() []K {
	var size int64
	if m.tree.Size() > 0 {
		size = m.tree.Size()
	}
	result := make([]K, 0, size)
	m.tree.Traverse(func(k K, v V) bool {
		result = append(result, k)
		return true
	})
	return result
}

func (m *OrderedMap[K, V]) Values() []V {
	var size int64
	if m.tree.Size() > 0 {
		size = m.tree.Size()
	}
	result := make([]V, 0, size)
	m.tree.Traverse(func(k K, v V) bool {
		result = append(result, v)
		return true
	})
	return result
}

func (m *OrderedMap[K, V]) Clear() {
	m.tree.Clear()
}

func (m *OrderedMap[K, V]) Iterator() *Iterator[K, V] {
	return &Iterator[K, V]{iter: m.tree.Iterator()}
}

type Iterator[K any, V any] struct {
	iter *indextree.Iterator[K, V]
}

func (iter *Iterator[K, V]) Valid() bool {
	return iter.iter.Valid()
}

func (iter *Iterator[K, V]) Key() K {
	return iter.iter.Key()
}

func (iter *Iterator[K, V]) Value() V {
	return iter.iter.Value()
}

func (iter *Iterator[K, V]) Index() int64 {
	return iter.iter.Index()
}

func (iter *Iterator[K, V]) SeekToFirst() {
	iter.iter.SeekToFirst()
}

func (iter *Iterator[K, V]) SeekToLast() {
	iter.iter.SeekToLast()
}

func (iter *Iterator[K, V]) SeekGE(key K) bool {
	return iter.iter.SeekGE(key)
}

func (iter *Iterator[K, V]) SeekGT(key K) bool {
	return iter.iter.SeekGT(key)
}

func (iter *Iterator[K, V]) SeekLE(key K) bool {
	return iter.iter.SeekLE(key)
}

func (iter *Iterator[K, V]) SeekLT(key K) bool {
	return iter.iter.SeekLT(key)
}

func (iter *Iterator[K, V]) SeekByIndex(index int64) {
	iter.iter.SeekByIndex(index)
}

func (iter *Iterator[K, V]) Next() {
	iter.iter.Next()
}

func (iter *Iterator[K, V]) Prev() {
	iter.iter.Prev()
}

func (iter *Iterator[K, V]) Clone() *Iterator[K, V] {
	return &Iterator[K, V]{iter: iter.iter.Clone()}
}
