package listqueue

type ListQueue struct {
	head *Element
	tail *Element

	size int64
}

func New() *ListQueue {
	return &ListQueue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (lq *ListQueue) Size() int64 {
	return lq.size
}

func (lq *ListQueue) Front() *Element {
	return lq.head
}

func (lq *ListQueue) Back() *Element {
	return lq.tail
}

func (lq *ListQueue) PushBack(value interface{}) {
	e := &Element{value: value}
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

func (lq *ListQueue) PushFront(value interface{}) {

	e := &Element{value: value}
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

func (lq *ListQueue) PopBack() interface{} {

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

func (lq *ListQueue) PopFront() interface{} {

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
