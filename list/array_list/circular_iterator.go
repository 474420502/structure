package arraylist

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
	iter.al.data[iter.cur+iter.al.headidx+1] = value
}

func (iter *CircularIterator[T]) Value() T {
	v, _ := iter.al.Index((int)(iter.cur))
	return v
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
