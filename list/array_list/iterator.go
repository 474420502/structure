package arraylist

import "log"

type Iterator[T comparable] struct {
	al  *ArrayList[T]
	cur uint
}

func (iter *Iterator[T]) Value() T {
	return iter.al.Index(iter.cur)
}

func (iter *Iterator[T]) Vaild() bool {
	return iter.cur < iter.al.size
}

// Swap  Swap iter Value
func (iter *Iterator[T]) Swap(other *Iterator[T]) {
	cidx := iter.cur + iter.al.headidx + 1
	oidx := other.cur + other.al.headidx + 1

	temp := iter.al.data[cidx]
	iter.al.data[cidx] = other.al.data[oidx]
	other.al.data[oidx] = temp
}

// SetValue Iter Value
func (iter *Iterator[T]) SetValue(value T) {
	iter.al.data[iter.cur+iter.al.headidx+1] = value
}

// Index index to the data
func (iter *Iterator[T]) Index(idx uint) {
	if idx >= iter.al.size {
		log.Panic("out of size")
	}
	iter.cur = idx
}

// RemoveToNext Remove self and to Next. must iter.Vaild() == true
func (iter *Iterator[T]) RemoveToNext() {
	iter.al.Remove(iter.cur)
	if iter.cur == iter.al.size {
		iter.cur--
	}
	return
}

// RemoveToNext Remove self and to Prev.  must iter.Vaild() == true
func (iter *Iterator[T]) RemoveToPrev() {
	iter.al.Remove(iter.cur)
	if iter.cur == 0 {
		return
	}
	iter.cur--
	return
}

// Prev to prev
func (iter *Iterator[T]) Prev() {
	iter.cur--
}

// Next to next
func (iter *Iterator[T]) Next() {
	iter.cur++
}

// ToHead to Head
func (iter *Iterator[T]) ToHead() {
	iter.cur = 0
}

// ToTail to Tail
func (iter *Iterator[T]) ToTail() {
	iter.cur = iter.al.size - 1
}
