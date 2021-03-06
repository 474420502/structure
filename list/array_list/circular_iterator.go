package arraylist

import "log"

// CircularIterator an iterator is an object that enables a programmer to traverse a container
type CircularIterator[T any] struct {
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

// IndexTo index to the data
func (iter *CircularIterator[T]) IndexTo(idx uint) {
	if idx >= iter.al.size {
		log.Panic("out of size")
	}
	iter.cur = idx
}

// IndexTo index to the data
func (iter *CircularIterator[T]) Index() uint {
	return iter.cur
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
}

// Value return current value iterator
func (iter *CircularIterator[T]) Value() T {
	return iter.al.Index(iter.cur)
}

// Vaild if current value is not nil return true. else return false. for use with Seek
func (iter *CircularIterator[T]) Vaild() bool {
	return iter.cur < iter.al.size
}

// Prev to prev
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

// Next to next
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

// ToHead to Head
func (iter *CircularIterator[T]) ToHead() {
	iter.cur = 0
}

// ToTail to Tail
func (iter *CircularIterator[T]) ToTail() {
	iter.cur = iter.al.size - 1
}
