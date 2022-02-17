package linkedlist

type Iterator[T comparable] struct {
	ll  *LinkedList[T]
	cur *Node[T]
}

// InsertFront insert T before the iterator.
func (iter *Iterator[T]) InsertFront(values ...T) {

	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		panic("iterator is nil")
	}

	var start *Node[T]
	var end *Node[T]

	iter.ll.size += uint(len(values))

	start = &Node[T]{value: values[0]}
	end = start

	for _, value := range values[1:] {
		node := &Node[T]{value: value}
		end.next = node
		node.prev = end
		end = node
	}

	cprev := iter.cur.prev

	cprev.next = start
	start.prev = cprev

	end.next = iter.cur
	iter.cur.prev = end
}

// InsertBack insert T after the iterator.
func (iter *Iterator[T]) InsertBack(values ...T) {

	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		panic("iterator is nil")
	}

	var start *Node[T]
	var end *Node[T]

	iter.ll.size += uint(len(values))

	start = &Node[T]{value: values[0]}
	end = start

	for _, value := range values[1:] {
		node := &Node[T]{value: value}
		end.next = node
		node.prev = end
		end = node
	}

	cnext := iter.cur.next

	iter.cur.next = start
	start.prev = iter.cur

	end.next = cnext
	cnext.prev = end
}

// RemoveToNext Remove self and to Next. If iterator is removed. return true.
func (iter *Iterator[T]) RemoveToNext() (ok bool) {
	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		return false
	}

	temp := iter.cur.next
	remove(iter.cur)
	iter.cur = temp

	return true
}

func (iter *Iterator[T]) Value() T {
	return iter.cur.value
}

// Move move next(prev[if step < 0]) by step
func (iter *Iterator[T]) Move(step int) (isEnd bool) {
	for i := 0; i < step; i++ {
		iter.cur = iter.cur.prev
	}
	return false
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
