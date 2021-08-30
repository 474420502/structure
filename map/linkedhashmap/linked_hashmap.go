package linkedhashmap

import (
	"fmt"
)

type Node struct {
	prev, next *Node
	key, value interface{}
}

// LinkedHashmap
type LinkedHashmap struct {
	head, tail *Node // 头节点 head->next 尾节点 tail->prev
	hmap       map[interface{}]*Node
}

// New
func New() *LinkedHashmap {
	lhmap := &LinkedHashmap{hmap: make(map[interface{}]*Node)}
	lhmap.head = &Node{}
	lhmap.tail = &Node{}
	lhmap.head.next = lhmap.tail
	lhmap.tail.prev = lhmap.head
	return lhmap
}

// Put equal to PushBack
func (lhmap *LinkedHashmap) Put(key interface{}, value interface{}) bool {
	return lhmap.PushBack(key, value)
}

// PushBack equal to Put, if key exists, push value replace the value is exists. size is unchanging
func (lhmap *LinkedHashmap) PushBack(key interface{}, value interface{}) bool {
	if _, ok := lhmap.hmap[key]; !ok {

		node := &Node{} // 创建空节点. 新的尾部节点

		// 直接在尾部赋值
		lhmap.tail.key = key
		lhmap.tail.value = value
		lhmap.hmap[key] = lhmap.tail

		node.prev = lhmap.tail
		lhmap.tail.next = node

		lhmap.tail = node // 重新定位尾部节点, 该节点是判断是否为尾部的关键

		return true
	}

	return false

}

// PushFront if key exists, push value replace the value is exists. size is unchanging
func (lhmap *LinkedHashmap) PushFront(key interface{}, value interface{}) bool {
	if _, ok := lhmap.hmap[key]; !ok {

		node := &Node{} // 创建空节点. 新的头部节点

		// 直接在尾部赋值
		lhmap.head.key = key
		lhmap.head.value = value
		lhmap.hmap[key] = lhmap.head

		node.next = lhmap.head

		lhmap.head.prev = node

		lhmap.head = node // 重新定位头部节点, 该节点是判断是否为头部的关键

		return true
	}

	return false
}

// Get
func (lhmap *LinkedHashmap) Get(key interface{}) (interface{}, bool) {
	node, ok := lhmap.hmap[key]
	return node.value, ok
}

// Clear
func (lhmap *LinkedHashmap) Clear() {
	lhmap.head.key = nil
	lhmap.head.value = nil
	lhmap.head.prev = nil

	lhmap.tail.key = nil
	lhmap.tail.value = nil
	lhmap.tail.next = nil

	lhmap.head.next = lhmap.tail
	lhmap.tail.prev = lhmap.head
	lhmap.hmap = make(map[interface{}]*Node)
}

// Remove if key not exists reture nil, false.
func (lhmap *LinkedHashmap) Remove(key interface{}) (interface{}, bool) {
	if node, ok := lhmap.hmap[key]; ok {
		delete(lhmap.hmap, key)
		lhmap.remove(node)
		return node.value, true
	}
	return nil, false
}

// Remove if key not exists reture nil, false.
func (lhmap *LinkedHashmap) remove(node *Node) {
	nprev := node.prev
	nnext := node.next
	nprev.next = nnext
	nnext.prev = nprev
}

// Empty returns true if map does not contain any elements
func (lhmap *LinkedHashmap) Empty() bool {
	return len(lhmap.hmap) == 0
}

// Size returns number of elements in the map.
func (lhmap *LinkedHashmap) Size() int {
	return len(lhmap.hmap)
}

// Keys returns all keys left to right (head to tail)
func (lhmap *LinkedHashmap) Keys() (result []interface{}) {

	head := lhmap.head.next
	for head != lhmap.tail {
		result = append(result, head.key)
		head = head.next
	}

	return
}

// Values returns all values in-order based on the key.
func (lhmap *LinkedHashmap) Values() (result []interface{}) {
	head := lhmap.head.next
	for head != lhmap.tail {
		result = append(result, head.value)
		head = head.next
	}
	return
}

// String returns a string
func (lhmap *LinkedHashmap) String() string {
	return fmt.Sprint(lhmap.Values())
}
