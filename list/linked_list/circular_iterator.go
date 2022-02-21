package linkedlist

type CircularIterator[T comparable] struct {
	ll  *LinkedList[T]
	cur *Node[T]
}

func (iter *CircularIterator[T]) Swap(other *CircularIterator[T]) {
	iter.cur.value, other.cur.value = other.cur.value, iter.cur.value
}

func (iter *CircularIterator[T]) SetValue(v T) {
	iter.cur.value = v
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
	iter.cur = iter.ll.head.next
}

func (iter *CircularIterator[T]) ToTail() {
	iter.cur = iter.ll.tail.prev
}

// Move move next(prev[if step < 0]) by step
func (iter *CircularIterator[T]) Move(step int) {
	if iter.ll.size == 0 {
		return
	}

	if step > 0 {
		if iter.cur == iter.ll.tail {
			iter.cur = iter.ll.head.next
		}

		for i := 0; i < step; i++ {

			iter.cur = iter.cur.next
			if iter.cur == iter.ll.tail {
				iter.cur = iter.ll.head.next
			}
		}
	} else {
		if iter.cur == iter.ll.head {
			iter.cur = iter.ll.tail.prev
		}

		for i := 0; i < -step; i++ {
			iter.cur = iter.cur.prev
			if iter.cur == iter.ll.head {
				iter.cur = iter.ll.tail.prev
			}
		}
	}
	return
}

// InsertFront insert T before the iterator.
func (iter *CircularIterator[T]) InsertFront(values ...T) {
	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		panic("iterator is nil. next move or next or prev to Value")
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
func (iter *CircularIterator[T]) InsertBack(values ...T) {

	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		panic("iterator is nil. next move or next or prev to Value")
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
func (iter *CircularIterator[T]) RemoveToNext() (ok bool) {
	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		return false
	}

	var temp = iter.cur.next
	remove(iter.cur)
	if temp == iter.ll.tail {
		iter.cur = iter.ll.head.next
	} else {
		iter.cur = temp
	}
	iter.ll.size--

	return true
}

// RemoveToNext Remove self and to Prev. If iterator is removed. return true.
func (iter *CircularIterator[T]) RemoveToPrev() (ok bool) {
	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		return false
	}

	var temp = iter.cur.prev
	remove(iter.cur)
	if temp == iter.ll.head {
		iter.cur = iter.ll.tail.prev
	} else {
		iter.cur = temp
	}
	iter.ll.size--

	return true
}
