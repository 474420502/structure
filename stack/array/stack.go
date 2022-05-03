package arraystack

import (
	"fmt"
)

// Stack the struct of stack
type Stack[T any] struct {
	element []T
	zero    T
}

// New  create a object of Stack
func New[T any]() *Stack[T] {
	st := &Stack[T]{}
	return st
}

// Push Push value into stack
func (st *Stack[T]) Push(v T) {
	st.element = append(st.element, v)
}

// Peek the top of stack
func (st *Stack[T]) Peek() (T, bool) {
	if len(st.element) == 0 {
		return st.zero, false
	}
	return st.element[len(st.element)-1], true
}

// Pop pop the value from stack
func (st *Stack[T]) Pop() (T, bool) {

	if len(st.element) == 0 {
		return st.zero, false
	}

	last := len(st.element) - 1
	ele := st.element[last]
	st.element = st.element[0:last]
	return ele, true
}

// Clear Clear stack data
func (st *Stack[T]) Clear() {
	st.element = st.element[0:0]
}

// Empty if stack is empty, return true. else false
func (st *Stack[T]) Empty() bool {
	l := len(st.element)
	return l == 0
}

// Size return the size of stack
func (st *Stack[T]) Size() uint {
	return uint(len(st.element))
}

// String return the string of stack
func (st *Stack[T]) String() string {
	return fmt.Sprintf("%v", st.element)
}

// Values return the values of stacks
func (st *Stack[T]) Values() []T {
	result := make([]T, len(st.element))
	copy(result, st.element)
	return result
}
