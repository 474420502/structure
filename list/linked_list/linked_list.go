package linkedlist

import (
	"fmt"

	"github.com/474420502/structure/compare"
)

// func assertImplementation() {
// 	var _ ilist.IList[any] = (*LinkedList[any])(nil)
// 	var _ ilist.IIterator[any] = (*Iterator[any])(nil)
// }

type hNode[T any] struct {
	prev  *hNode[T]
	next  *hNode[T]
	value T
}

// func (node *Node[T]) Value() T {
// 	return node.value
// }

// LinkedList struct of LinkedList
type LinkedList[T any] struct {
	head *hNode[T]
	tail *hNode[T]
	size uint
	comp compare.Compare[T]
}

// New create a object of LinkedList
func New[T any](comp compare.Compare[T]) *LinkedList[T] {
	l := &LinkedList[T]{comp: comp}
	l.head = &hNode[T]{}
	l.head.prev = nil

	l.tail = &hNode[T]{}
	l.tail.next = nil

	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// Iterator  an iterator is an object that enables a programmer to traverse a container, particularly lists
func (l *LinkedList[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{ll: l, cur: l.head.next}
}

// Iterator
// an iterator is an object that enables a programmer to traverse a container
// and can circulate
func (l *LinkedList[T]) CircularIterator() *CircularIterator[T] {
	return &CircularIterator[T]{ll: l, cur: l.head.next}
}

// Clear clear the list
func (l *LinkedList[T]) Clear() {

	l.head.next = l.tail
	l.tail.prev = l.head

	l.size = 0
}

// Empty   if the list is empty, return true. else return false
func (l *LinkedList[T]) Empty() bool {
	return l.size == 0
}

// Size return the size of list
func (l *LinkedList[T]) Size() uint {
	return l.size
}

// Push Push a value to the tail of the list
func (l *LinkedList[T]) Push(value T) {
	var node *hNode[T]
	l.size++

	node = &hNode[T]{}
	node.value = value

	tprev := l.tail.prev
	tprev.next = node

	node.prev = tprev
	node.next = l.tail
	l.tail.prev = node
}

// PushFront Push values to the head of the list
func (l *LinkedList[T]) PushFront(values ...T) {

	var node *hNode[T]
	l.size += uint(len(values))
	for _, v := range values {
		node = &hNode[T]{}
		node.value = v

		hnext := l.head.next
		hnext.prev = node

		node.next = hnext
		node.prev = l.head
		l.head.next = node
	}
}

// PushBack Push  values to the tail of the list
func (l *LinkedList[T]) PushBack(values ...T) {

	var node *hNode[T]
	l.size += uint(len(values))
	for _, v := range values {
		node = &hNode[T]{}
		node.value = v

		tprev := l.tail.prev
		tprev.next = node

		node.prev = tprev
		node.next = l.tail
		l.tail.prev = node
	}
}

// PopFront pop the head of the list
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

// PopBack pop the back of the list
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

// Front return the head of list
func (l *LinkedList[T]) Front() (result T, found bool) {
	if l.size != 0 {
		return l.head.next.value, true
	}
	found = false
	return
}

// Back return the head of list
func (l *LinkedList[T]) Back() (result T, found bool) {
	if l.size != 0 {
		return l.tail.prev.value, true
	}
	found = false
	return
}

// Index slowly. is a list. need to move with idx step
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

func remove[T any](cur *hNode[T]) {
	curPrev := cur.prev
	curNext := cur.next
	curPrev.next = curNext
	curNext.prev = curPrev
	cur.prev = nil
	cur.next = nil
}

// Contains is the []T  in list?
func (l *LinkedList[T]) Contains(values ...T) (count int) {

	for cur := l.head.next; cur != l.tail; cur = cur.next {
		for _, searchValue := range values {
			if l.comp(cur.value, searchValue) == 0 {
				count++
			}
		}
	}
	return
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
