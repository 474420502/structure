package lastack

import (
	"fmt"
)

type hNode[T any] struct {
	elements []T
	cur      int
	down     *hNode[T]
}

// Stack the struct of stack
type Stack[T any] struct {
	top   *hNode[T]
	cache *hNode[T]
	size  uint

	zero T
}

func (as *Stack[T]) grow() bool {
	if as.top.cur+1 == cap(as.top.elements) {

		var grownode *hNode[T]
		if as.cache != nil {
			grownode = as.cache
			grownode.cur = -1
			as.cache = nil
		} else {
			var growsize uint
			if as.size <= 256 {
				growsize = as.size << 1
			} else {
				growsize = 256 + as.size>>2
			}
			grownode = &hNode[T]{elements: make([]T, growsize), cur: -1}
		}

		grownode.down = as.top
		as.top = grownode
		return true
	}

	return false
}

func (as *Stack[T]) cacheRemove() bool {
	if as.top.cur == 0 && as.top.down != nil {
		as.cache = as.top
		as.top = as.top.down
		as.cache.down = nil
		return true
	}

	return false
}

// New  create a object of stack
func New[T any]() *Stack[T] {
	s := &Stack[T]{}
	s.size = 0
	s.top = &hNode[T]{elements: make([]T, 8), cur: -1}
	return s
}

// New  create a object of stack with the capacity
func NewWithCap[T any](cap int) *Stack[T] {
	s := &Stack[T]{}
	s.size = 0
	s.top = &hNode[T]{elements: make([]T, cap), cur: -1}
	return s
}

// Clear Clear stack data
func (as *Stack[T]) Clear() {
	as.size = 0

	as.top.down = nil
	as.top.cur = -1
}

// Empty if stack is empty, return true. else false
func (as *Stack[T]) Empty() bool {
	return as.size == 0
}

// Size return the size of stack
func (as *Stack[T]) Size() uint {
	return as.size
}

// String return the string of stack. a(top)->b->c
func (as *Stack[T]) String() string {

	return fmt.Sprintf("%v", as.Values())
}

// Values return the values of stacks
func (as *Stack[T]) Values() []T {
	result := make([]T, as.size)

	cur := as.top
	n := as.size - 1
	for ; cur != nil; cur = cur.down {
		for i := range cur.elements {
			if cur.cur >= i {
				result[n] = cur.elements[cur.cur-i]
				n--
			}
		}
	}

	return result
}

// func (as *Stack[T]) Index(idx int) (interface{}, bool) {
// 	if idx < 0 {
// 		return nil, false
// 	}

// 	cur := as.top
// 	for cur != nil && idx-cur.cur > 0 {
// 		idx = idx - cur.cur - 1
// 		cur = cur.down
// 	}

// 	if cur == nil {
// 		return nil, false
// 	}

// 	return cur.elements[cur.cur-idx], true
// }

// Push Push value into stack
func (as *Stack[T]) Push(v T) {
	as.grow()
	as.top.cur++
	as.top.elements[as.top.cur] = v
	as.size++
}

// Pop pop the value from stack
func (as *Stack[T]) Pop() (T, bool) {
	if as.size <= 0 {
		return as.zero, false
	}

	as.size--
	if as.cacheRemove() {
		return as.cache.elements[as.cache.cur], true
	}

	epop := as.top.elements[as.top.cur]
	as.top.cur--
	return epop, true
}

// Peek the top of stack
func (as *Stack[T]) Peek() (T, bool) {
	if as.size <= 0 {
		return as.zero, false
	}
	return as.top.elements[as.top.cur], true
}
