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
	return &CircularIterator[T]{ll: l, cur: l.head}
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

func (l *LinkedList[T]) PopFront() (result T, found bool) {
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
	found = false
	return
}

func (l *LinkedList[T]) PopBack() (result T, found bool) {
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
	found = false
	return
}

func (l *LinkedList[T]) Front() (result T, found bool) {
	if l.size != 0 {
		return l.head.next.value, true
	}
	found = false
	return
}

func (l *LinkedList[T]) Back() (result T, found bool) {
	if l.size != 0 {
		return l.tail.prev.value, true
	}
	found = false
	return
}

func (l *LinkedList[T]) Index(idx int) (result T, ok bool) {

	if idx < 0 {
		ok = false
		return
	}
	var uidx = (uint)(idx)

	if uidx >= l.size || idx < 0 {
		ok = false
		return
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

	ok = false
	return
}

func remove[T comparable](cur *Node[T]) {
	curPrev := cur.prev
	curNext := cur.next
	curPrev.next = curNext
	curNext.prev = curPrev
	cur.prev = nil
	cur.next = nil
}

func (l *LinkedList[T]) Remove(idx int) (result T, ok bool) {

	if idx < 0 {
		ok = false
		return
	}

	var uidx uint = (uint)(idx)
	if l.size <= uidx {
		// log.Printf("out of list range, size is %d, idx is %d\n", l.size, idx)
		ok = false
		return
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
func (l *LinkedList[T]) RemoveIf(every func(idx uint, value T) RemoveState) (result []T, isRemoved bool) {
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

// Contains is the T  in list?
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

// Values get the values of list
func (l *LinkedList[T]) Values() (result []T) {
	l.Traverse(func(value T) bool {
		result = append(result, value)
		return true
	})
	return
}

// String fmt.Sprintf("%v", l.Values())
func (l *LinkedList[T]) String() string {
	return fmt.Sprintf("%v", l.Values())
}

// Traverse from the list of head to the tail. iterator can do it also.
func (l *LinkedList[T]) Traverse(every func(value T) bool) {
	for cur := l.head.next; cur != l.tail; cur = cur.next {
		if !every(cur.value) {
			break
		}
	}
}
