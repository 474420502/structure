package arraylist

import "log"

type CircularIterator[T comparable] struct {
	al     *ArrayList[T]
	cur    uint
	isInit bool
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
	if !iter.isInit {
		log.Panic("the index of iterator is unknownd, call Index(uint) or Prev() Next()  ToHead() ToTail()")
	}
	iter.al.data[iter.cur+iter.al.headidx+1] = value
}

// Index index to the data
func (iter *CircularIterator[T]) Index(idx uint) {
	if idx >= iter.al.size {
		log.Panic("out of size")
	}
	iter.cur = idx
	iter.isInit = true
}

// RemoveToNext Remove self and to Next. if iter is tail. isTail = true else false
func (iter *CircularIterator[T]) RemoveToNext() (isTail bool) {
	if !iter.isInit {
		log.Panic("the index of iterator is unknownd, call Index(uint) or Prev() Next()  ToHead() ToTail()")
	}
	iter.al.Remove(iter.cur)
	if iter.cur == iter.al.size {
		iter.cur--
		isTail = true
	}
	return
}

// RemoveToNext Remove self and to Prev.  if iter is head.  isHead = true else false
func (iter *CircularIterator[T]) RemoveToPrev() (isHead bool) {
	if !iter.isInit {
		log.Panic("the index of iterator is unknownd, call Index(uint) or Prev() Next() ToHead() ToTail()")
	}
	iter.al.Remove(iter.cur)
	if iter.cur == 0 {
		isHead = true
		return
	}
	iter.cur--
	return
}

func (iter *CircularIterator[T]) Value() T {
	return iter.al.Index(iter.cur)
}

func (iter *CircularIterator[T]) Prev() bool {

	if !iter.isInit {
		if iter.al.size != 0 {
			iter.isInit = true
			iter.cur = iter.al.size - 1
			return true
		}
		return false
	}

	if iter.al.size == 0 {
		return false
	}

	if iter.cur <= 0 {
		iter.cur = iter.al.size - 1
	} else {
		iter.cur--
	}
	return true
}

func (iter *CircularIterator[T]) Next() bool {

	if !iter.isInit {
		if iter.al.size != 0 {
			iter.isInit = true
			iter.cur = 0
			return true
		}
		return false
	}

	if iter.al.size == 0 {
		return false
	}

	if iter.cur >= iter.al.size-1 {
		iter.cur = 0
	} else {
		iter.cur++
	}
	return true
}

func (iter *CircularIterator[T]) ToHead() {
	iter.isInit = true
	iter.cur = 0
}

func (iter *CircularIterator[T]) ToTail() {
	iter.isInit = true
	iter.cur = iter.al.size - 1
}
