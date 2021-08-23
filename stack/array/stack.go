package arraystack

import (
	"fmt"
)

// Stack 栈
type Stack struct {
	element []interface{}
}

// New  创建一个Stack
func New() *Stack {
	st := &Stack{}
	return st
}

// Push 压栈
func (st *Stack) Push(v interface{}) {
	st.element = append(st.element, v)
}

// Peek 相当与栈顶
func (st *Stack) Peek() (interface{}, bool) {
	if len(st.element) == 0 {
		return nil, false
	}
	return st.element[len(st.element)-1], true
}

// Pop 出栈
func (st *Stack) Pop() (interface{}, bool) {

	if len(st.element) == 0 {
		return nil, false
	}

	last := len(st.element) - 1
	ele := st.element[last]
	st.element = st.element[0:last]
	return ele, true
}

// Clear 清空栈数据
func (st *Stack) Clear() {
	st.element = st.element[0:0]
}

// Empty 如果空栈返回true
func (st *Stack) Empty() bool {
	l := len(st.element)
	return l == 0
}

// Size 数据量
func (st *Stack) Size() uint {
	return uint(len(st.element))
}

// String
func (st *Stack) String() string {
	return fmt.Sprintf("%v", st.element)
}

// Values 同上
func (st *Stack) Values() []interface{} {
	result := make([]interface{}, len(st.element))
	copy(result, st.element)
	return result
}
