package arraylist

import (
	"fmt"
	"log"

	"github.com/474420502/structure/compare"
)

// func assertImplementation() {
// 	var _ ilist.IList[any] = (*ArrayList[any])(nil)
// 	var _ ilist.IIterator[any] = (*Iterator[any])(nil)
// }

// ArrayList a list base on array
type ArrayList[T any] struct {
	data    []T
	headidx uint // [ nil(hdix) 1 nil(tidx) ]
	tailidx uint
	size    uint

	growthSize uint
	shrinkSize uint
	comp       compare.Compare[T]
}

const (
	listMaxLimit = uint(1) << 63 // the max size of list
	listMinLimit = uint(8)       // the min size of list
	initCap      = uint(8)
	//growthFactor = float32(2.0)  // growth by 100%
	//shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New create a object of list
func New[T any](comp compare.Compare[T]) *ArrayList[T] {
	l := &ArrayList[T]{}
	l.data = make([]T, initCap)
	l.tailidx = initCap / 2
	l.headidx = l.tailidx - 1
	// l.shrinkSize = listMinLimit
	l.comp = comp
	return l
}

// Iterator  an iterator is an object that enables a programmer to traverse a container, particularly lists
func (l *ArrayList[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{al: l, cur: 0}
}

// Iterator
// an iterator is an object that enables a programmer to traverse a container
// and can circulate
func (l *ArrayList[T]) CircularIterator() *CircularIterator[T] {
	return &CircularIterator[T]{al: l, cur: 0}
}

// Clear clear the list
func (l *ArrayList[T]) Clear() {
	l.data = make([]T, 8)
	l.tailidx = initCap / 2
	l.headidx = l.tailidx - 1
	l.size = 0
}

// Empty   if the list is empty, return true. else return false
func (l *ArrayList[T]) Empty() bool {
	return l.size == 0
}

// Size return the size of list
func (l *ArrayList[T]) Size() uint {
	return l.size
}

func (l *ArrayList[T]) shrink() {

	if l.size <= listMinLimit {
		return
	}

	if l.size <= l.shrinkSize {
		lcap := uint(len(l.data))
		nSize := lcap - lcap>>2
		temp := make([]T, nSize)

		ghidx := l.size >> 2
		gtidx := ghidx + l.size + 1
		copy(temp[ghidx+1:], l.data[l.headidx+1:l.tailidx])
		l.data = temp
		l.headidx = ghidx
		l.tailidx = gtidx

		// l.shrinkSize = l.shrinkSize - lcap>>2
		l.shrinkSize = l.size - l.size>>2

	}
}

// 后续需要优化 growth 策略
func (l *ArrayList[T]) growth() {

	if l.size >= listMaxLimit {
		log.Panic("list size is over listMaxLimit", listMaxLimit)
	}

	lcap := uint(len(l.data))
	nSize := lcap << 1
	temp := make([]T, nSize)

	ghidx := lcap / 2
	gtidx := ghidx + l.size + 1
	copy(temp[ghidx+1:], l.data[l.headidx+1:l.tailidx])
	l.data = temp
	l.headidx = ghidx
	l.tailidx = gtidx

	l.shrinkSize = l.size - l.size>>2
}

// Push Push a value to the tail of the list
func (l *ArrayList[T]) Push(value T) {
	for l.tailidx+1 > uint(len(l.data)) {
		l.growth()
	}
	l.data[l.tailidx] = value
	l.tailidx++
	l.size += 1
}

// PushFront Push values to the head of the list
func (l *ArrayList[T]) PushFront(values ...T) {
	psize := uint(len(values))
	for l.headidx+1-psize > listMaxLimit {
		l.growth()
		// panic("growth -1")
	}

	for _, v := range values {
		l.data[l.headidx] = v
		l.headidx--
	}
	l.size += psize
}

// PushBack Push  values to the tail of the list
func (l *ArrayList[T]) PushBack(values ...T) {
	psize := uint(len(values))
	for l.tailidx+psize > uint(len(l.data)) { // [0 1 2 3 4 5 6]
		l.growth()
	}

	for _, v := range values {
		l.data[l.tailidx] = v
		l.tailidx++
	}
	l.size += psize
}

// Front return the head of list
func (l *ArrayList[T]) Front() (result T, ok bool) {
	if l.size != 0 {
		return l.data[l.headidx+1], true
	}
	ok = false
	return
}

// Back return the head of list
func (l *ArrayList[T]) Back() (result T, ok bool) {
	if l.size != 0 {
		return l.data[l.tailidx-1], true
	}
	ok = false
	return
}

// PopFront pop the head of the list
func (l *ArrayList[T]) PopFront() (result T, ok bool) {
	if l.size != 0 {
		l.size--
		l.headidx++
		result = l.data[l.headidx]
		l.shrink()
		return result, true
	}
	ok = false
	return
}

// PopBack pop the back of the list
func (l *ArrayList[T]) PopBack() (result T, ok bool) {
	if l.size != 0 {
		l.size--
		l.tailidx--
		result = l.data[l.tailidx]
		l.shrink()
		return result, true
	}
	ok = false
	return
}

// Index fast to index. the feature of array list
func (l *ArrayList[T]) Index(idx uint) T {
	if idx >= l.size {
		log.Panic("out of size.", l.size)
	}
	return l.data[idx+l.headidx+1]
}

// Set similar to slice[idx] = value. the feature of array list
func (l *ArrayList[T]) Set(idx int, value T) {

	l.data[uint(idx)+l.headidx+1] = value
}

// Remove remove the value by index
func (l *ArrayList[T]) Remove(idx uint) (result T) {

	if idx >= l.size {
		log.Panic("out of size:", idx)
		return
	}

	offset := l.headidx + 1 + idx
	result = l.data[offset]

	if uint(len(l.data))-l.tailidx > l.headidx {
		copy(l.data[offset:], l.data[offset+1:l.tailidx]) // shift to the left by one (slow operation, need ways to optimize this)
		l.tailidx--
	} else {
		copy(l.data[l.headidx+2:], l.data[l.headidx+1:offset])
		l.headidx++
	}

	l.size--
	l.shrink()
	return
}

// Contains determining whether a list contains values. return the count of values.
func (l *ArrayList[T]) Contains(values ...T) (count int) {

	for _, element := range l.data[l.headidx+1 : l.tailidx] {
		for _, searchValue := range values {
			if l.comp(element, searchValue) == -1 {
				count++
			}
		}
	}

	return
}

// Values return all values of list
func (l *ArrayList[T]) Values() []T {
	newElements := make([]T, l.size)
	copy(newElements, l.data[l.headidx+1:l.tailidx])
	return newElements
}

// String print the string of list
func (l *ArrayList[T]) String() string {
	return fmt.Sprintf("%v", l.Values())
}

// Traverse Traversing containers
func (l *ArrayList[T]) Traverse(every func(idx uint, value T) bool) {
	for i := uint(0); i < l.size; i++ {
		if !every(i, l.data[i+l.headidx+1]) {
			break
		}
	}
}
