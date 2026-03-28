package orderedmap

import (
	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/indextree"
)

type OrderedMap[K any, V any] struct {
	tree *indextree.Tree[K]
}

func New[K any, V any](comp compare.Compare[K]) *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		tree: indextree.New(comp),
	}
}

func (m *OrderedMap[K, V]) Size() int64 {
	return m.tree.Size()
}

func (m *OrderedMap[K, V]) IsEmpty() bool {
	return m.tree.Size() == 0
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	v, ok := m.tree.Get(key)
	if !ok {
		return *new(V), false
	}
	return v.(V), true
}

func (m *OrderedMap[K, V]) Put(key K, value V) bool {
	return m.tree.Put(key, value)
}

func (m *OrderedMap[K, V]) Set(key K, value V) bool {
	return m.tree.Set(key, value)
}

func (m *OrderedMap[K, V]) Remove(key K) (V, bool) {
	v := m.tree.Remove(key)
	if v == nil {
		return *new(V), false
	}
	return v.(V), true
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	_, ok := m.tree.Get(key)
	return ok
}

func (m *OrderedMap[K, V]) IndexOf(key K) int64 {
	return m.tree.IndexOf(key)
}

func (m *OrderedMap[K, V]) Index(index int64) (K, V) {
	k, v := m.tree.Index(index)
	return k, v.(V)
}

func (m *OrderedMap[K, V]) RemoveIndex(index int64) (K, V, bool) {
	if index < 0 || index >= m.tree.Size() {
		return *new(K), *new(V), false
	}
	k, _ := m.tree.Index(index)
	vv := m.tree.RemoveIndex(index)
	return k, vv.(V), true
}

func (m *OrderedMap[K, V]) Keys() []K {
	var size int64
	if m.tree.Size() > 0 {
		size = m.tree.Size()
	}
	result := make([]K, 0, size)
	m.tree.Traverse(func(k K, v interface{}) bool {
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
	m.tree.Traverse(func(k K, v interface{}) bool {
		result = append(result, v.(V))
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
	iter *indextree.Iterator[K]
}

func (iter *Iterator[K, V]) Valid() bool {
	return iter.iter.Valid()
}

func (iter *Iterator[K, V]) Key() K {
	return iter.iter.Key()
}

func (iter *Iterator[K, V]) Value() V {
	return iter.iter.Value().(V)
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
