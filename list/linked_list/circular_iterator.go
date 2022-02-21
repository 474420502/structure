package linkedlist

type CircularIterator[T comparable] struct {
	ll  *LinkedList[T]
	cur *Node[T]
}

func (iter *CircularIterator[T]) Swap(other *CircularIterator[T]) {
	iter.cur.value, other.cur.value = other.cur.value, iter.cur.value
}

// SetValue set the value of current iter
func (iter *CircularIterator[T]) SetValue(v T) {
	iter.cur.value = v
}

// Value get the value of element. must iter.Vaild() == true
func (iter *CircularIterator[T]) Value() interface{} {
	return iter.cur.value
}

// Vaild current is Vaild ?
func (iter *CircularIterator[T]) Vaild() bool {
	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		return false
	}
	return true
}

// Prev the prev element
func (iter *CircularIterator[T]) Prev() {
	if iter.ll.size == 0 {
		return
	}

	if iter.cur == iter.ll.head {
		iter.cur = iter.ll.tail.prev
		return
	}

	iter.cur = iter.cur.prev
	if iter.cur == iter.ll.head {
		iter.cur = iter.ll.tail.prev
	}

	return
}

// Next the next element
func (iter *CircularIterator[T]) Next() {
	if iter.ll.size == 0 {
		return
	}

	if iter.cur == iter.ll.tail {
		iter.cur = iter.ll.head.next
		return
	}

	iter.cur = iter.cur.next
	if iter.cur == iter.ll.tail {
		iter.cur = iter.ll.head.next
	}

	return
}

// ToHead to list head element
func (iter *CircularIterator[T]) ToHead() {
	iter.cur = iter.ll.head.next
}

// ToTail to list tail element
func (iter *CircularIterator[T]) ToTail() {
	iter.cur = iter.ll.tail.prev
}

// Move move next(prev[if step < 0]) by step must iter.Vaild() == true
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

// InsertFront insert T before the iterator. must iter.Vaild() == true
func (iter *CircularIterator[T]) InsertFront(values ...T) {

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

// InsertBack insert T after the iterator. must iter.Vaild() == true
func (iter *CircularIterator[T]) InsertBack(values ...T) {

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

// RemoveToNext Remove self and to Next.
func (iter *CircularIterator[T]) RemoveToNext() {

	var temp = iter.cur.next
	remove(iter.cur)
	if temp == iter.ll.tail {
		iter.cur = iter.ll.head.next
	} else {
		iter.cur = temp
	}
	iter.ll.size--

}

// RemoveToNext Remove self and to Prev.
func (iter *CircularIterator[T]) RemoveToPrev() {

	var temp = iter.cur.prev
	remove(iter.cur)
	if temp == iter.ll.head {
		iter.cur = iter.ll.tail.prev
	} else {
		iter.cur = temp
	}
	iter.ll.size--

}
