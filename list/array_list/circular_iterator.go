package arraylist

import "log"

type CircularIterator[T comparable] struct {
	al  *ArrayList[T]
	cur uint
}

// Swap  Swap iter Value
func (iter *CircularIterator[T]) Swap(other *CircularIterator[T]) {
	cidx := iter.cur + iter.al.headidx + 1
	oidx := other.cur + other.al.headidx + 1

	temp := iter.al.data[cidx]
	iter.al.data[cidx] = other.al.data[oidx]
	other.al.data[oidx] = temp
}

// SetValue Iter Value
func (iter *CircularIterator[T]) SetValue(value T) {
	iter.al.data[iter.cur+iter.al.headidx+1] = value
}

// Index index to the data
func (iter *CircularIterator[T]) Index(idx uint) {
	if idx >= iter.al.size {
		log.Panic("out of size")
	}
	iter.cur = idx
}

// RemoveToNext Remove self and to Next. if iter is tail. isTail = true else false
func (iter *CircularIterator[T]) RemoveToNext() {
	iter.al.Remove(iter.cur)
	return
}

// RemoveToNext Remove self and to Prev.  if iter is head.  isHead = true else false
func (iter *CircularIterator[T]) RemoveToPrev() {
	iter.al.Remove(iter.cur)
	iter.cur--
	return
}

func (iter *CircularIterator[T]) Value() T {
	return iter.al.Index(iter.cur)
}

func (iter *CircularIterator[T]) Vaild() bool {
	return iter.cur < iter.al.size
}

func (iter *CircularIterator[T]) Prev() {

	if iter.al.size == 0 {
		return
	}

	if iter.cur == 0 {
		iter.cur = iter.al.size - 1
	} else {
		iter.cur--
	}
	return
}

func (iter *CircularIterator[T]) Next() {
	if iter.al.size == 0 {
		return
	}

	if iter.cur >= iter.al.size-1 {
		iter.cur = 0
	} else {
		iter.cur++
	}
	return
}

func (iter *CircularIterator[T]) ToHead() {
	iter.cur = 0
}

func (iter *CircularIterator[T]) ToTail() {
	iter.cur = iter.al.size - 1
}
