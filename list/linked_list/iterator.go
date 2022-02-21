package linkedlist

type Iterator[T comparable] struct {
	ll  *LinkedList[T]
	cur *Node[T]
}

// InsertFront insert T before the iterator. must iter.Vaild() == true
func (iter *Iterator[T]) InsertFront(values ...T) {

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

// InsertBack insert T after the iterator.  must iter.Vaild() == true
func (iter *Iterator[T]) InsertBack(values ...T) {

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

// RemoveToNext Remove self and to Next. If iterator is removed. return true.  must iter.Vaild() == true
func (iter *Iterator[T]) RemoveToNext() {
	temp := iter.cur.next
	remove(iter.cur)
	iter.cur = temp
	iter.ll.size--
}

// RemoveToNext Remove self and to Prev. If iterator is removed. return true.  must iter.Vaild() == true
func (iter *Iterator[T]) RemoveToPrev() {
	temp := iter.cur.prev
	remove(iter.cur)
	iter.cur = temp
	iter.ll.size--
}

// Swap  must iter.Vaild() == true
func (iter *Iterator[T]) Swap(other *Iterator[T]) {
	iter.cur.value, other.cur.value = other.cur.value, iter.cur.value
}

//SetValue  must iter.Vaild() == true
func (iter *Iterator[T]) SetValue(v T) {
	iter.cur.value = v
}

// Value must iter.Vaild() == true
func (iter *Iterator[T]) Value() T {
	return iter.cur.value
}

// Vaild current is Vaild ?
func (iter *Iterator[T]) Vaild() bool {
	if iter.cur == iter.ll.head || iter.cur == iter.ll.tail {
		return false
	}
	return true
}

// Move move next(prev[if step < 0]) by step. must iter.Vaild() == true
func (iter *Iterator[T]) Move(step int) {

	if step > 0 {
		if iter.cur == iter.ll.tail {
			return
		}

		for i := 0; i < step; i++ {
			iter.cur = iter.cur.next
			if iter.cur == iter.ll.tail {
				return
			}
		}
	} else {
		if iter.cur == iter.ll.head {
			return
		}

		for i := 0; i < -step; i++ {
			iter.cur = iter.cur.prev
			if iter.cur == iter.ll.head {
				return
			}
		}
	}

	return
}

//Prev must iter.Vaild() == true
func (iter *Iterator[T]) Prev() {
	iter.cur = iter.cur.prev
}

//Next must iter.Vaild() == true
func (iter *Iterator[T]) Next() {
	iter.cur = iter.cur.next
}

// ToHead. to head and must iter.Vaild() == true
func (iter *Iterator[T]) ToHead() {
	iter.cur = iter.ll.head.next
}

// ToTail. to tail and must iter.Vaild() == true
func (iter *Iterator[T]) ToTail() {
	iter.cur = iter.ll.tail.prev
}
