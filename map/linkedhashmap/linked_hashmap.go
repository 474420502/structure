package linkedhashmap

import (
	"fmt"
)

type hNode struct {
	prev, next *hNode
	key, value interface{}
}

// LinkedHashmap
type LinkedHashmap struct {
	head, tail *hNode // 头节点 head->next 尾节点 tail->prev
	hmap       map[interface{}]*hNode
}

// New
func New() *LinkedHashmap {

	lhmap := &LinkedHashmap{hmap: make(map[interface{}]*hNode)}
	lhmap.head = &hNode{}
	lhmap.tail = &hNode{}
	lhmap.head.next = lhmap.tail
	lhmap.tail.prev = lhmap.head
	return lhmap
}

// SetBack equal to Cover, if key exists, cover and move node to back, return true. else insert new node to back, return false
func (lhmap *LinkedHashmap) SetBack(key interface{}, value interface{}) bool {

	var ok bool
	var node *hNode

	if node, ok = lhmap.hmap[key]; ok {
		node.value = value
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
		node = &hNode{}
		// 直接在尾部赋值
		lhmap.tail.key = key
		lhmap.tail.value = value
		lhmap.hmap[key] = lhmap.tail

		node.prev = lhmap.tail
		lhmap.tail.next = node

		lhmap.tail = node // 重新定位尾部节点, 该节点是判断是否为尾部的关键
	}

	return ok

}

// SetFront if key exists, cover and move node to front, return true. else insert new node to front. return false
func (lhmap *LinkedHashmap) SetFront(key interface{}, value interface{}) bool {
	var ok bool
	var node *hNode

	if node, ok = lhmap.hmap[key]; ok {
		node.value = value
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
		node = &hNode{} // 创建空节点. 新的头部节点

		// 直接在尾部赋值
		lhmap.head.key = key
		lhmap.head.value = value
		lhmap.hmap[key] = lhmap.head

		node.next = lhmap.head
		lhmap.head.prev = node
		lhmap.head = node // 重新定位头部节点, 该节点是判断是否为头部的关键
	}

	return ok
}

// Put equal to PushBack
func (lhmap *LinkedHashmap) Put(key interface{}, value interface{}) bool {
	return lhmap.PushBack(key, value)
}

// PushBack equal to Put, if key exists, skip value and return false. size is unchanging
func (lhmap *LinkedHashmap) PushBack(key interface{}, value interface{}) bool {
	if _, ok := lhmap.hmap[key]; !ok {

		node := &hNode{} // 创建空节点. 新的尾部节点

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

// PushFront if key exists, skip value and return false. size is unchanging
func (lhmap *LinkedHashmap) PushFront(key interface{}, value interface{}) bool {
	if _, ok := lhmap.hmap[key]; !ok {

		node := &hNode{} // 创建空节点. 新的头部节点

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

// Get get the value
func (lhmap *LinkedHashmap) Get(key interface{}) (interface{}, bool) {
	node, ok := lhmap.hmap[key]
	return node.value, ok
}

// Set if key exists set value and return true. else return false.
func (lhmap *LinkedHashmap) Set(key, value interface{}) bool {
	if node, ok := lhmap.hmap[key]; ok {
		node.key = key
		node.value = value
		return true
	}
	return false
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
	lhmap.hmap = make(map[interface{}]*hNode)
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
func (lhmap *LinkedHashmap) remove(node *hNode) {
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
