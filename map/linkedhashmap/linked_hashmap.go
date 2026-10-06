package linkedhashmap

import (
	"fmt"
)

type Slice[K comparable, V any] struct {
	Key   K
	Value V
}

type hNode[K comparable, V any] struct {
	Slice[K, V]
	prev, next *hNode[K, V]
}

// LinkedHashmap keeps insertion order with a linked list plus a hashmap.
type LinkedHashmap[K comparable, V any] struct {
	head, tail *hNode[K, V] // head  tail
	hmap       map[K]*hNode[K, V]
}

// LinkedHashMap is the preferred exported spelling kept in sync with LinkedHashmap.
type LinkedHashMap[K comparable, V any] = LinkedHashmap[K, V]

// New create a object of LinkedHashmap
func New[K comparable, V any]() *LinkedHashmap[K, V] {
	lhmap := &LinkedHashmap[K, V]{hmap: make(map[K]*hNode[K, V])}
	lhmap.head = &hNode[K, V]{}
	lhmap.tail = &hNode[K, V]{}
	lhmap.head.next = lhmap.tail
	lhmap.tail.prev = lhmap.head
	return lhmap
}

// String return string of slice
func (s Slice[K, V]) String() string {
	return fmt.Sprintf("{%v:%v}", s.Key, s.Value)
}

// SetBack equal to Cover, if key exists, cover and move node to back, return true. else insert new node to back, return false
func (lhmap *LinkedHashmap[K, V]) SetBack(key K, value V) bool {

	var ok bool
	var node *hNode[K, V]

	if node, ok = lhmap.hmap[key]; ok {
		node.Value = value
		if node == lhmap.tail.prev {
			return ok
		}

		// 删除node的节点
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev

		tprev := lhmap.tail.prev
		// 连接尾部
		tprev.next = node
		node.prev = tprev
		node.next = lhmap.tail
		lhmap.tail.prev = node

	} else {
		node = &hNode[K, V]{}
		// 直接在尾部赋值
		lhmap.tail.Key = key
		lhmap.tail.Value = value
		lhmap.hmap[key] = lhmap.tail

		node.prev = lhmap.tail
		lhmap.tail.next = node

		lhmap.tail = node // 重新定位尾部节点, 该节点是判断是否为尾部的关键
	}

	return ok

}

// SetFront if key exists, cover and move node to front, return true. else insert new node to front. return false
func (lhmap *LinkedHashmap[K, V]) SetFront(key K, value V) bool {
	var ok bool
	var node *hNode[K, V]

	if node, ok = lhmap.hmap[key]; ok {
		node.Value = value
		if node == lhmap.head.next {
			return ok
		}

		// 删除node的节点
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev

		hnext := lhmap.head.next
		// 连接头部
		hnext.prev = node
		node.next = hnext
		node.prev = lhmap.head
		lhmap.head.next = node

	} else {
		node = &hNode[K, V]{} // 创建空节点. 新的头部节点

		// 直接在尾部赋值
		lhmap.head.Key = key
		lhmap.head.Value = value
		lhmap.hmap[key] = lhmap.head

		node.next = lhmap.head
		lhmap.head.prev = node
		lhmap.head = node // 重新定位头部节点, 该节点是判断是否为头部的关键
	}

	return ok
}

// Put equal to PushBack
func (lhmap *LinkedHashmap[K, V]) Put(key K, value V) bool {
	return lhmap.PushBack(key, value)
}

// InsertIfAbsent inserts a value only when the key does not exist.
func (lhmap *LinkedHashmap[K, V]) InsertIfAbsent(key K, value V) bool {
	return lhmap.Put(key, value)
}

// PushBack equal to Put, if key exists, skip value and return false. size is unchanging
func (lhmap *LinkedHashmap[K, V]) PushBack(key K, value V) bool {
	if _, ok := lhmap.hmap[key]; !ok {

		node := &hNode[K, V]{} // 创建空节点. 新的尾部节点

		// 直接在尾部赋值
		lhmap.tail.Key = key
		lhmap.tail.Value = value
		lhmap.hmap[key] = lhmap.tail

		node.prev = lhmap.tail
		lhmap.tail.next = node

		lhmap.tail = node // 重新定位尾部节点, 该节点是判断是否为尾部的关键

		return true
	}

	return false

}

// PushFront if key exists, skip value and return false. size is unchanging
func (lhmap *LinkedHashmap[K, V]) PushFront(key K, value V) bool {
	if _, ok := lhmap.hmap[key]; !ok {

		node := &hNode[K, V]{} // 创建空节点. 新的头部节点

		// 直接在尾部赋值
		lhmap.head.Key = key
		lhmap.head.Value = value
		lhmap.hmap[key] = lhmap.head

		node.next = lhmap.head

		lhmap.head.prev = node

		lhmap.head = node // 重新定位头部节点, 该节点是判断是否为头部的关键

		return true
	}

	return false
}

// Get get the value
func (lhmap *LinkedHashmap[K, V]) Get(key K) (V, bool) {
	if node, ok := lhmap.hmap[key]; ok {
		return node.Value, true
	}
	var zero V
	return zero, false
}

// Set if key exists set value and return true. else return false and do nothing.
func (lhmap *LinkedHashmap[K, V]) Set(key K, value V) bool {
	if node, ok := lhmap.hmap[key]; ok {
		node.Key = key
		node.Value = value
		return true
	}
	return false
}

// Upsert sets the value and reports whether an existing value was replaced.
// New keys are appended to the back to match Put semantics.
func (lhmap *LinkedHashmap[K, V]) Upsert(key K, value V) bool {
	if node, ok := lhmap.hmap[key]; ok {
		node.Key = key
		node.Value = value
		return true
	}
	lhmap.Put(key, value)
	return false
}

// Clear clear the LinkedHashmap
func (lhmap *LinkedHashmap[K, V]) Clear() {
	var zeroK K
	var zeroV V

	lhmap.head.Key = zeroK
	lhmap.head.Value = zeroV
	lhmap.head.prev = nil

	lhmap.tail.Key = zeroK
	lhmap.tail.Value = zeroV
	lhmap.tail.next = nil

	lhmap.head.next = lhmap.tail
	lhmap.tail.prev = lhmap.head
	lhmap.hmap = make(map[K]*hNode[K, V])
}

// Remove if key not exists reture nil, false.
func (lhmap *LinkedHashmap[K, V]) Remove(key K) (V, bool) {
	if node, ok := lhmap.hmap[key]; ok {
		delete(lhmap.hmap, key)
		lhmap.remove(node)
		return node.Value, true
	}
	var zero V
	return zero, false
}

// Delete removes a key and returns the previous value when present.
func (lhmap *LinkedHashmap[K, V]) Delete(key K) (V, bool) {
	return lhmap.Remove(key)
}

// remove unlink a node from the order list.
func (lhmap *LinkedHashmap[K, V]) remove(node *hNode[K, V]) {
	nprev := node.prev
	nnext := node.next
	nprev.next = nnext
	nnext.prev = nprev
}

// Empty returns true if map does not contain any elements
func (lhmap *LinkedHashmap[K, V]) Empty() bool {
	return len(lhmap.hmap) == 0
}

// Size returns number of elements in the map.
func (lhmap *LinkedHashmap[K, V]) Size() uint {
	return uint(len(lhmap.hmap))
}

// Len returns the number of elements.
func (lhmap *LinkedHashmap[K, V]) Len() int {
	return len(lhmap.hmap)
}

// Keys returns all keys left to right (head to tail)
func (lhmap *LinkedHashmap[K, V]) Keys() []K {
	result := make([]K, 0, len(lhmap.hmap))
	head := lhmap.head.next
	for head != lhmap.tail {
		result = append(result, head.Key)
		head = head.next
	}
	return result
}

// Values returns all values in-order.
func (lhmap *LinkedHashmap[K, V]) Values() []V {
	result := make([]V, 0, len(lhmap.hmap))
	head := lhmap.head.next
	for head != lhmap.tail {
		result = append(result, head.Value)
		head = head.next
	}
	return result
}

// Slices returns all keyvalue in-order.
func (lhmap *LinkedHashmap[K, V]) Slices() []Slice[K, V] {
	result := make([]Slice[K, V], 0, len(lhmap.hmap))
	head := lhmap.head.next
	for head != lhmap.tail {
		result = append(result, head.Slice)
		head = head.next
	}
	return result
}

// String returns a string
func (lhmap *LinkedHashmap[K, V]) String() string {
	return fmt.Sprint(lhmap.Slices())
}
