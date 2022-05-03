package liststack

import "fmt"

type hNode[T any] struct {
	value T
	down  *hNode[T]
}

// Stack the struct of stack
type Stack[T any] struct {
	top  *hNode[T]
	size uint
	zero T
}

// New  create a object of Stack
func New[T any]() *Stack[T] {
	s := &Stack[T]{}
	s.size = 0
	return s
}

// Clear Clear stack data
func (ls *Stack[T]) Clear() {
	ls.size = 0
	ls.top = nil
}

// Empty if stack is empty, return true. else false
func (ls *Stack[T]) Empty() bool {
	return ls.size == 0
}

// Size return the size of stack
func (ls *Stack[T]) Size() uint {
	return ls.size
}

// String return the string of stack. a(top)->b->c
func (ls *Stack[T]) String() string {
	return fmt.Sprintf("%v", ls.Values())
}

// Values return the values of stacks
func (ls *Stack[T]) Values() []T {

	if ls.size == 0 {
		return nil
	}

	result := make([]T, ls.size)

	cur := ls.top
	n := ls.size - 1
	for ; cur != nil; cur = cur.down {
		result[n] = cur.value
		n--
	}

	return result
}

// Push Push value into stack
func (ls *Stack[T]) Push(v T) {
	nv := &hNode[T]{value: v}
	nv.down = ls.top
	ls.top = nv
	ls.size++
}

// Pop pop the value from stack
func (ls *Stack[T]) Pop() (T, bool) {
	if ls.size == 0 {
		return ls.zero, false
	}

	ls.size--

	result := ls.top
	ls.top = ls.top.down
	// result.down = nil
	return result.value, true
}

// Peek the top of stack
func (ls *Stack[T]) Peek() (T, bool) {
	if ls.size == 0 {
		return ls.zero, false
	}
	return ls.top.value, true
}
