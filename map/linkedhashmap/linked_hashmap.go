package linkedhashmap

import (
	"fmt"
)

type Slice struct {
	Key, Value interface{}
}

type hNode struct {
	Slice
	prev, next *hNode
}

// LinkedHashmap   LinkedHashmap like list + hashmap.
type LinkedHashmap struct {
	head, tail *hNode // head  tail
	hmap       map[interface{}]*hNode
}

// New create a object of LinkedHashmap
func New() *LinkedHashmap {

	lhmap := &LinkedHashmap{hmap: make(map[interface{}]*hNode)}
	lhmap.head = &hNode{}
	lhmap.tail = &hNode{}
	lhmap.head.next = lhmap.tail
	lhmap.tail.prev = lhmap.head
	return lhmap
}

// String return string of slice
func (s Slice) String() string {
	return fmt.Sprintf("{%v:%v}", s.Key, s.Value)
}

// SetBack equal to Cover, if key exists, cover and move node to back, return true. else insert new node to back, return false
func (lhmap *LinkedHashmap) SetBack(key interface{}, value interface{}) bool {

	var ok bool
	var node *hNode

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
		node = &hNode{}
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
func (lhmap *LinkedHashmap) SetFront(key interface{}, value interface{}) bool {
	var ok bool
	var node *hNode

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
		node = &hNode{} // 创建空节点. 新的头部节点

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
func (lhmap *LinkedHashmap) Put(key interface{}, value interface{}) bool {
	return lhmap.PushBack(key, value)
}

// PushBack equal to Put, if key exists, skip value and return false. size is unchanging
func (lhmap *LinkedHashmap) PushBack(key interface{}, value interface{}) bool {
	if _, ok := lhmap.hmap[key]; !ok {

		node := &hNode{} // 创建空节点. 新的尾部节点

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
func (lhmap *LinkedHashmap) PushFront(key interface{}, value interface{}) bool {
	if _, ok := lhmap.hmap[key]; !ok {

		node := &hNode{} // 创建空节点. 新的头部节点

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
func (lhmap *LinkedHashmap) Get(key interface{}) (interface{}, bool) {
	node, ok := lhmap.hmap[key]
	if ok {
		return node.Value, ok
	}
	return nil, false
}

// Set if key exists set value and return true. else return false and do nothing.
func (lhmap *LinkedHashmap) Set(key, value interface{}) bool {
	if node, ok := lhmap.hmap[key]; ok {
		node.Key = key
		node.Value = value
		return true
	}
	return false
}

// Clear clear the LinkedHashmap
func (lhmap *LinkedHashmap) Clear() {
	lhmap.head.Key = nil
	lhmap.head.Value = nil
	lhmap.head.prev = nil

	lhmap.tail.Key = nil
	lhmap.tail.Value = nil
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
		return node.Value, true
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
		result = append(result, head.Key)
		head = head.next
	}

	return
}

// Values returns all values in-order.
func (lhmap *LinkedHashmap) Values() (result []interface{}) {
	head := lhmap.head.next
	for head != lhmap.tail {
		result = append(result, head.Value)
		head = head.next
	}
	return
}

// Slices returns all keyvalue in-order.
func (lhmap *LinkedHashmap) Slices() (result []Slice) {
	head := lhmap.head.next
	for head != lhmap.tail {
		result = append(result, head.Slice)
		head = head.next
	}
	return
}

// String returns a string
func (lhmap *LinkedHashmap) String() string {
	return fmt.Sprint(lhmap.Slices())
}
