package linkedlist

import (
	"fmt"
)

type Node[T comparable] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

func (node *Node[T]) Value() T {
	return node.value
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size uint
}

func New[T comparable]() *LinkedList[T] {
	l := &LinkedList[T]{}
	l.head = &Node[T]{}
	l.head.prev = nil

	l.tail = &Node[T]{}
	l.tail.next = nil

	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LinkedList[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{ll: l, cur: l.head}
}

func (l *LinkedList[T]) CircularIterator() *CircularIterator[T] {
	return &CircularIterator[T]{pl: l, cur: l.head}
}

func (l *LinkedList[T]) Clear() {

	l.head.next = l.tail
	l.tail.prev = l.head

	l.size = 0
}

func (l *LinkedList[T]) Empty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Size() uint {
	return l.size
}

func (l *LinkedList[T]) Push(value T) {
	var node *Node[T]
	l.size++

	node = &Node[T]{}
	node.value = value

	tprev := l.tail.prev
	tprev.next = node

	node.prev = tprev
	node.next = l.tail
	l.tail.prev = node
}

func (l *LinkedList[T]) PushFront(values ...T) {

	var node *Node[T]
	l.size += uint(len(values))
	for _, v := range values {
		node = &Node[T]{}
		node.value = v

		hnext := l.head.next
		hnext.prev = node

		node.next = hnext
		node.prev = l.head
		l.head.next = node
	}
}

func (l *LinkedList[T]) PushBack(values ...T) {

	var node *Node[T]
	l.size += uint(len(values))
	for _, v := range values {
		node = &Node[T]{}
		node.value = v

		tprev := l.tail.prev
		tprev.next = node

		node.prev = tprev
		node.next = l.tail
		l.tail.prev = node
	}
}

func (l *LinkedList[T]) PopFront() (result interface{}, found bool) {
	if l.size != 0 {
		l.size--

		temp := l.head.next
		hnext := temp.next
		hnext.prev = l.head
		l.head.next = hnext

		result = temp.value
		found = true
		return
	}
	return nil, false
}

func (l *LinkedList[T]) PopBack() (result interface{}, found bool) {
	if l.size != 0 {
		l.size--

		temp := l.tail.prev
		tprev := temp.prev
		tprev.next = l.tail
		l.tail.prev = tprev

		result = temp.value
		found = true
		return
	}
	return nil, false
}

func (l *LinkedList[T]) Front() (result interface{}, found bool) {
	if l.size != 0 {
		return l.head.next.value, true
	}
	return nil, false
}

func (l *LinkedList[T]) Back() (result interface{}, found bool) {
	if l.size != 0 {
		return l.tail.prev.value, true
	}
	return nil, false
}

func (l *LinkedList[T]) Index(idx int) (interface{}, bool) {

	if idx < 0 {
		return nil, false
	}
	var uidx = (uint)(idx)

	if uidx >= l.size || idx < 0 {
		return nil, false
	}

	if uidx > l.size/2 {
		uidx = l.size - 1 - uidx
		// 尾部
		for cur := l.tail.prev; cur != l.head; cur = cur.prev {
			if uidx == 0 {
				return cur.value, true
			}
			uidx--
		}

	} else {
		// 头部
		for cur := l.head.next; cur != l.tail; cur = cur.next {
			if uidx == 0 {
				return cur.value, true
			}
			uidx--
		}
	}

	return nil, false
}

func (l *LinkedList[T]) Insert(idx uint, values ...T) bool {
	if idx > l.size { // 插入的方式 可能导致size的范围判断不一样
		return false
	}

	if idx > l.size/2 {
		idx = l.size - idx
		// 尾部
		for cur := l.tail.prev; cur != nil; cur = cur.prev {

			if idx == 0 {

				var start *Node[T]
				var end *Node[T]

				start = &Node[T]{value: values[0]}
				end = start

				for _, value := range values[1:] {
					node := &Node[T]{value: value}
					end.next = node
					node.prev = end
					end = node
				}

				cnext := cur.next

				cur.next = start
				start.prev = cur

				end.next = cnext
				cnext.prev = end

				break
			}

			idx--
		}

	} else {
		// 头部
		for cur := l.head.next; cur != nil; cur = cur.next {
			if idx == 0 {

				var start *Node[T]
				var end *Node[T]

				start = &Node[T]{value: values[0]}
				end = start

				for _, value := range values[1:] {
					node := &Node[T]{value: value}
					end.next = node
					node.prev = end
					end = node
				}

				cprev := cur.prev

				cprev.next = start
				start.prev = cprev

				end.next = cur
				cur.prev = end

				break
			}

			idx--
		}
	}

	l.size += uint(len(values))
	return true
}

// InsertState InsertIf的every函数的枚举  从左到右 1 为前 2 为后 insert here(2) ->cur-> insert here(1)
// UninsertAndContinue 不插入并且继续
// UninsertAndBreak 不插入并且停止
// InsertBack cur后插入并且停止
// InsertFront cur前插入并且停止
type InsertState int

const (
	// UninsertAndContinue 不插入并且继续
	UninsertAndContinue InsertState = 0
	// UninsertAndBreak 不插入并且停止
	UninsertAndBreak InsertState = -1
	// InsertBack cur后插入并且停止
	InsertBack InsertState = 2
	// InsertFront cur前插入并且停止
	InsertFront InsertState = 1
)

// InsertIf  every函数的枚举  从左到右遍历 1 为前 2 为后 insert here(2) ->cur-> insert here(1)
func (l *LinkedList[T]) InsertIf(every func(idx uint, value T) InsertState, values ...T) {

	idx := uint(0)
	// 头部
	for cur := l.head.next; cur != nil; cur = cur.next {
		insertState := every(idx, cur.value)

		if insertState == UninsertAndContinue {
			continue
		}

		if insertState > 0 { // 1 为前 2 为后 insert here(2) ->cur-> insert here(1)
			var start *Node[T]
			var end *Node[T]

			start = &Node[T]{value: values[0]}
			end = start

			for _, value := range values[1:] {
				node := &Node[T]{value: value}
				end.next = node
				node.prev = end
				end = node
			}

			if insertState == InsertBack {
				cprev := cur.prev
				cprev.next = start
				start.prev = cprev
				end.next = cur
				cur.prev = end
			} else { // InsertFront
				cnext := cur.next
				cnext.prev = end
				start.prev = cur
				cur.next = start
				end.next = cnext
			}

			l.size += uint(len(values))
			return
		}

		// 必然 等于 UninsertAndBreak
		return
	}
}

func remove[T comparable](cur *Node[T]) {
	curPrev := cur.prev
	curNext := cur.next
	curPrev.next = curNext
	curNext.prev = curPrev
	cur.prev = nil
	cur.next = nil
}

func (l *LinkedList[T]) Remove(idx int) (interface{}, bool) {

	if idx < 0 {
		return nil, false
	}

	var uidx uint = (uint)(idx)
	if l.size <= uidx {
		// log.Printf("out of list range, size is %d, idx is %d\n", l.size, idx)
		return nil, false
	}

	l.size--
	if uidx > l.size/2 {
		uidx = l.size - uidx // l.size - 1 - idx,  先减size
		// 尾部
		for cur := l.tail.prev; cur != l.head; cur = cur.prev {
			if uidx == 0 {
				remove(cur)
				return cur.value, true
			}
			uidx--
		}

	} else {
		// 头部
		for cur := l.head.next; cur != l.tail; cur = cur.next {
			if uidx == 0 {
				remove(cur)
				return cur.value, true

			}
			uidx--
		}
	}

	panic(fmt.Errorf("unknown error"))
}

// RemoveState RemoveIf的every函数的枚举
// RemoveAndContinue 删除并且继续
// RemoveAndBreak 删除并且停止
// UnremoveAndBreak 不删除并且停止遍历
// UnremoveAndContinue 不删除并且继续遍历
type RemoveState int

const (
	// RemoveAndContinue 删除并且继续
	RemoveAndContinue RemoveState = iota
	// RemoveAndBreak 删除并且停止
	RemoveAndBreak
	// UnremoveAndBreak 不删除并且停止遍历
	UnremoveAndBreak
	// UnremoveAndContinue 不删除并且继续遍历
	UnremoveAndContinue
)

// RemoveIf every的遍历函数操作remove过程 如果没删除result 返回nil, isRemoved = false
func (l *LinkedList[T]) RemoveIf(every func(idx uint, value T) RemoveState) (result []interface{}, isRemoved bool) {
	// 头部
	idx := uint(0)
TOPFOR:
	for cur := l.head.next; cur != l.tail; idx++ {
		removeState := every(idx, cur.value)
		switch removeState {
		case RemoveAndContinue:
			result = append(result, cur.value)
			isRemoved = true
			temp := cur.next
			remove(cur)
			cur = temp
			l.size--
			continue TOPFOR
		case RemoveAndBreak:
			result = append(result, cur.value)
			isRemoved = true
			temp := cur.next
			remove(cur)
			cur = temp
			l.size--
			return
		case UnremoveAndContinue:
		case UnremoveAndBreak:
			return
		}

		cur = cur.next
	}
	return
}

func (l *LinkedList[T]) Contains(values ...T) bool {

	for _, searchValue := range values {
		found := false
		for cur := l.head.next; cur != l.tail; cur = cur.next {
			if cur.value == searchValue {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

func (l *LinkedList[T]) Values() (result []interface{}) {
	l.Traverse(func(value T) bool {
		result = append(result, value)
		return true
	})
	return
}

func (l *LinkedList[T]) String() string {
	return fmt.Sprintf("%v", l.Values())
}

func (l *LinkedList[T]) Traverse(every func(value T) bool) {
	for cur := l.head.next; cur != l.tail; cur = cur.next {
		if !every(cur.value) {
			break
		}
	}
}
