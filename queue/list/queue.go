package listqueue

type ListQueue[T any] struct {
	head *Element[T]
	tail *Element[T]

	size int64
}

func New[T any]() *ListQueue[T] {
	return &ListQueue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (lq *ListQueue[T]) Size() int64 {
	return lq.size
}

func (lq *ListQueue[T]) Front() *Element[T] {
	return lq.head
}

func (lq *ListQueue[T]) Back() *Element[T] {
	return lq.tail
}

func (lq *ListQueue[T]) PushBack(value T) {
	e := &Element[T]{value: value}
	lq.size++
	if lq.size == 1 {
		lq.head = e
		lq.tail = e
		return
	}

	e.prev = lq.tail
	lq.tail.next = e
	lq.tail = e
}

func (lq *ListQueue[T]) PushFront(value T) {

	e := &Element[T]{value: value}
	lq.size++
	if lq.size == 1 {
		lq.head = e
		lq.tail = e
		return
	}

	e.next = lq.head
	lq.head.prev = e
	lq.head = e

}

func (lq *ListQueue[T]) PopBack() interface{} {

	if lq.size == 0 {
		return nil
	}

	lq.size--
	if lq.size == 0 {
		p := lq.tail
		lq.head = nil
		lq.tail = nil
		return p.value
	}

	p := lq.tail

	prev := p.prev
	prev.next = nil
	p.prev = nil

	lq.tail = prev

	return p.value
}

func (lq *ListQueue[T]) PopFront() interface{} {

	if lq.size == 0 {
		return nil
	}

	lq.size--
	if lq.size == 0 {
		p := lq.head
		lq.head = nil
		lq.tail = nil
		return p.value
	}

	p := lq.head

	next := p.next
	next.prev = nil
	p.next = nil

	lq.head = next

	return p.value
}
