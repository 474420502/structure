package arraylist

type Iterator[T comparable] struct {
	al     *ArrayList[T]
	cur    uint
	isInit bool
}

func (iter *Iterator[T]) Value() interface{} {
	v, _ := iter.al.Index((int)(iter.cur))
	return v
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

type CircularIterator[T comparable] struct {
	al     *ArrayList[T]
	cur    uint
	isInit bool
}

func (iter *CircularIterator[T]) Value() interface{} {
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
