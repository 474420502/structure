package listqueue

type Element struct {
	prev  *Element
	next  *Element
	value interface{}
}

func (e *Element) Prev() *Element {
	return e.prev
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Value() interface{} {
	return e.value
}
