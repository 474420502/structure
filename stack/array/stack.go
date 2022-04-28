package arraystack

import (
	"fmt"
)

// Stack 栈
type Stack[T any] struct {
	element []T
}

// New  创建一个Stack
func New[T any]() *Stack[T] {
	st := &Stack[T]{}
	return st
}

// Push 压栈
func (st *Stack[T]) Push(v T) {
	st.element = append(st.element, v)
}

// Peek 相当与栈顶
func (st *Stack[T]) Peek() (interface{}, bool) {
	if len(st.element) == 0 {
		return nil, false
	}
	return st.element[len(st.element)-1], true
}

// Pop 出栈
func (st *Stack[T]) Pop() (interface{}, bool) {

	if len(st.element) == 0 {
		return nil, false
	}

	last := len(st.element) - 1
	ele := st.element[last]
	st.element = st.element[0:last]
	return ele, true
}

// Clear 清空栈数据
func (st *Stack[T]) Clear() {
	st.element = st.element[0:0]
}

// Empty 如果空栈返回true
func (st *Stack[T]) Empty() bool {
	l := len(st.element)
	return l == 0
}

// Size 数据量
func (st *Stack[T]) Size() uint {
	return uint(len(st.element))
}

// String
func (st *Stack[T]) String() string {
	return fmt.Sprintf("%v", st.element)
}

// Values 同上
func (st *Stack[T]) Values() []T {
	result := make([]T, len(st.element))
	copy(result, st.element)
	return result
}
