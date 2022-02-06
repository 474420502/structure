package linkedlist

type Iterator[T comparable] struct {
	ll  *LinkedList[T]
	cur *Node[T]
}

func (iter *Iterator[T]) Value() interface{} {
	return iter.cur.value
}

func (iter *Iterator[T]) Prev() bool {
	if iter.cur == iter.ll.head {
		return false
	}
	iter.cur = iter.cur.prev
	return iter.cur != iter.ll.head
}

func (iter *Iterator[T]) Next() bool {
	if iter.cur == iter.ll.tail {
		return false
	}
	iter.cur = iter.cur.next
	return iter.cur != iter.ll.tail
}

func (iter *Iterator[T]) ToHead() {
	iter.cur = iter.ll.head
}

func (iter *Iterator[T]) ToTail() {
	iter.cur = iter.ll.tail
}

type CircularIterator[T comparable] struct {
	pl  *LinkedList[T]
	cur *Node[T]
}

func (iter *CircularIterator[T]) Value() interface{} {
	return iter.cur.value
}

func (iter *CircularIterator[T]) Prev() bool {
	if iter.pl.size == 0 {
		return false
	}

	if iter.cur == iter.pl.head {
		iter.cur = iter.pl.tail.prev
		return true
	}

	iter.cur = iter.cur.prev
	if iter.cur == iter.pl.head {
		iter.cur = iter.pl.tail.prev
	}

	return true
}

func (iter *CircularIterator[T]) Next() bool {
	if iter.pl.size == 0 {
		return false
	}

	if iter.cur == iter.pl.tail {
		iter.cur = iter.pl.head.next
		return true
	}

	iter.cur = iter.cur.next
	if iter.cur == iter.pl.tail {
		iter.cur = iter.pl.head.next
	}

	return true
}

func (iter *CircularIterator[T]) ToHead() {
	iter.cur = iter.pl.head
}

func (iter *CircularIterator[T]) ToTail() {
	iter.cur = iter.pl.tail
}
