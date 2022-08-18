package listqueue

type Element[T any] struct {
	prev  *Element[T]
	next  *Element[T]
	value T
}

func (e *Element[T]) Prev() *Element[T] {
	return e.prev
}

func (e *Element[T]) Next() *Element[T] {
	return e.next
}

func (e *Element[T]) Value() interface{} {
	return e.value
}
