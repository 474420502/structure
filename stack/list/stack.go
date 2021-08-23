package liststack

import "fmt"

type Node struct {
	value interface{}
	down  *Node
}

type Stack struct {
	top  *Node
	size uint
}

func New() *Stack {
	s := &Stack{}
	s.size = 0
	return s
}

func (ls *Stack) Clear() {
	ls.size = 0
	ls.top = nil
}

func (ls *Stack) Empty() bool {
	return ls.size == 0
}

func (ls *Stack) Size() uint {
	return ls.size
}

// String 从左到右 左边第一个表示Top 如链表 a(top)->b->c
func (ls *Stack) String() string {
	return fmt.Sprintf("%v", ls.Values())
}

func (ls *Stack) Values() []interface{} {

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

func (ls *Stack) Push(v interface{}) {
	nv := &Node{value: v}
	nv.down = ls.top
	ls.top = nv
	ls.size++
}

func (ls *Stack) Pop() (interface{}, bool) {
	if ls.size == 0 {
		return nil, false
	}

	ls.size--

	result := ls.top
	ls.top = ls.top.down
	// result.down = nil
	return result.value, true
}

func (ls *Stack) Peek() (interface{}, bool) {
	if ls.size == 0 {
		return nil, false
	}
	return ls.top.value, true
}
