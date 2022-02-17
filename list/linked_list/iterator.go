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
	iter.ll.size--

	return true
}

// RemoveToNext Remove self and to Prev. If iterator is removed. return true.
func (iter *Iterator[T]) RemoveToPrev() (ok bool) {
	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		return false
	}

	temp := iter.cur.prev
	remove(iter.cur)
	iter.cur = temp
	iter.ll.size--

	return true
}

func (iter *Iterator[T]) Value() T {
	return iter.cur.value
}

// Move move next(prev[if step < 0]) by step
func (iter *Iterator[T]) Move(step int) (isEnd bool) {
	if step > 0 {
		for i := 0; i < step; i++ {
			iter.cur = iter.cur.next
			if iter.cur == iter.ll.tail {
				return true
			}
		}
	} else {
		for i := 0; i < -step; i++ {
			iter.cur = iter.cur.prev
			if iter.cur == iter.ll.head {
				return true
			}
		}
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

// ToHead
func (iter *Iterator[T]) ToHead() {
	iter.cur = iter.ll.head
}

// ToTail
func (iter *Iterator[T]) ToTail() {
	iter.cur = iter.ll.tail
}

type CircularIterator[T comparable] struct {
	ll  *LinkedList[T]
	cur *Node[T]
}

func (iter *CircularIterator[T]) Value() interface{} {
	return iter.cur.value
}

func (iter *CircularIterator[T]) Prev() bool {
	if iter.ll.size == 0 {
		return false
	}

	if iter.cur == iter.ll.head {
		iter.cur = iter.ll.tail.prev
		return true
	}

	iter.cur = iter.cur.prev
	if iter.cur == iter.ll.head {
		iter.cur = iter.ll.tail.prev
	}

	return true
}

func (iter *CircularIterator[T]) Next() bool {
	if iter.ll.size == 0 {
		return false
	}

	if iter.cur == iter.ll.tail {
		iter.cur = iter.ll.head.next
		return true
	}

	iter.cur = iter.cur.next
	if iter.cur == iter.ll.tail {
		iter.cur = iter.ll.head.next
	}

	return true
}

func (iter *CircularIterator[T]) ToHead() {
	iter.cur = iter.ll.head
}

func (iter *CircularIterator[T]) ToTail() {
	iter.cur = iter.ll.tail
}

// Move move next(prev[if step < 0]) by step
func (iter *CircularIterator[T]) Move(step int) {
	if step > 0 {
		for i := 0; i < step; i++ {

			if iter.cur == iter.ll.tail {
				iter.cur = iter.ll.head.next

			}

			iter.cur = iter.cur.next
			if iter.cur == iter.ll.tail {
				iter.cur = iter.ll.head.next
			}

		}
	} else {
		for i := 0; i < -step; i++ {

			if iter.cur == iter.ll.head {
				iter.cur = iter.ll.tail.prev
			}

			iter.cur = iter.cur.prev
			if iter.cur == iter.ll.head {
				iter.cur = iter.ll.tail.prev
			}

		}
	}

	return
}
