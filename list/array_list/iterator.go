package arraylist

type Iterator[T comparable] struct {
	al     *ArrayList[T]
	cur    uint
	isInit bool
}

func (iter *Iterator[T]) Value() T {
	v, _ := iter.al.Index((int)(iter.cur))
	return v
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

func (iter *Iterator[T]) Prev() bool {

	if !iter.isInit {
		if iter.al.size != 0 {
			iter.isInit = true
			iter.cur = iter.al.size - 1
			return true
		}
		return false
	}

	if iter.cur <= 0 {
		return false
	}
	iter.cur--
	return true
}

func (iter *Iterator[T]) Next() bool {

	if !iter.isInit {
		if iter.al.size != 0 {
			iter.isInit = true
			iter.cur = 0
			return true
		}
		return false
	}

	if iter.cur >= iter.al.size-1 {
		return false
	}
	iter.cur++
	return true
}

func (iter *Iterator[T]) ToHead() {
	iter.isInit = true
	iter.cur = 0
}

func (iter *Iterator[T]) ToTail() {
	iter.isInit = true
	iter.cur = iter.al.size - 1
}
