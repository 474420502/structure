package liststack

import "fmt"

type Node[T any] struct {
	value T
	down  *Node[T]
}

type Stack[T any] struct {
	top  *Node[T]
	size uint
}

func New[T any]() *Stack[T] {
	s := &Stack[T]{}
	s.size = 0
	return s
}

func (ls *Stack[T]) Clear() {
	ls.size = 0
	ls.top = nil
}

func (ls *Stack[T]) Empty() bool {
	return ls.size == 0
}

func (ls *Stack[T]) Size() uint {
	return ls.size
}

// String 从左到右 左边第一个表示Top 如链表 a(top)->b->c
func (ls *Stack[T]) String() string {
	return fmt.Sprintf("%v", ls.Values())
}

func (ls *Stack[T]) Values() []interface{} {

	if ls.size == 0 {
		return nil
	}

	result := make([]interface{}, ls.size, ls.size)

	cur := ls.top
	n := ls.size - 1
	for ; cur != nil; cur = cur.down {
		result[n] = cur.value
		n--
	}

	return result
}

func (ls *Stack[T]) Push(v T) {
	nv := &Node[T]{value: v}
	nv.down = ls.top
	ls.top = nv
	ls.size++
}

func (ls *Stack[T]) Pop() (interface{}, bool) {
	if ls.size == 0 {
		return nil, false
	}

	ls.size--

	result := ls.top
	ls.top = ls.top.down
	// result.down = nil
	return result.value, true
}

func (ls *Stack[T]) Peek() (interface{}, bool) {
	if ls.size == 0 {
		return nil, false
	}
	return ls.top.value, true
}
