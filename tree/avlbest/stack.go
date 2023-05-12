package avlbest

import (
	"fmt"
)

// Stack the struct of stack
type Stack[NTYPE any] struct {
	element []NodeDir[NTYPE]
	index   int
}

type NodeDir[NTYPE any] struct {
	Node  *NTYPE
	Child int
}

// NewStack  create a object of Stack
func NewStack[NTYPE any]() *Stack[NTYPE] {
	st := &Stack[NTYPE]{index: -1, element: make([]NodeDir[NTYPE], 16)}
	return st
}

// Push Push value into stack
func (st *Stack[NTYPE]) Push(v *NTYPE, child int) {
	st.index += 1
	if st.index >= len(st.element) {
		cap := len(st.element) << 1
		temp := make([]NodeDir[NTYPE], cap)
		copy(temp, st.element)
		st.element = temp
	}
	e := &st.element[st.index]
	e.Node = v
	e.Child = child
}

// Peek the top of stack
func (st *Stack[NTYPE]) Peek() *NodeDir[NTYPE] {
	if st.index < 0 {
		return nil
	}
	return &st.element[st.index]
}

// Pop pop the count of the value without return the values from stack
func (st *Stack[NTYPE]) PopFast(count int) {
	if st.index-count < -1 {
		// panic("no more element to pop")
		return
	}
	st.index -= count
}

// Pop pop the value from stack
func (st *Stack[NTYPE]) Pop() *NodeDir[NTYPE] {
	if st.index < 0 {
		return nil
	}

	ele := &st.element[st.index]
	st.index -= 1
	return ele
}

// Clear Clear stack data
func (st *Stack[NTYPE]) Clear() {
	st.index = -1
}

// Empty if stack is empty, return true. else false
func (st *Stack[NTYPE]) Empty() bool {
	return st.index < 0
}

// Size return the size of stack
func (st *Stack[NTYPE]) Size() int {
	return st.index + 1
}

// String return the string of stack
func (st *Stack[NTYPE]) String() string {
	return fmt.Sprintf("%v", st.element)
}
