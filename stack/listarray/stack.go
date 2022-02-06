package lastack

import (
	"fmt"
)

type Node[T any] struct {
	elements []T
	cur      int
	down     *Node[T]
}

type Stack[T any] struct {
	top   *Node[T]
	cache *Node[T]
	size  uint
}

func (as *Stack[T]) grow() bool {
	if as.top.cur+1 == cap(as.top.elements) {

		var grownode *Node[T]
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
			grownode = &Node[T]{elements: make([]T, growsize, growsize), cur: -1}
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

func New[T any]() *Stack[T] {
	s := &Stack[T]{}
	s.size = 0
	s.top = &Node[T]{elements: make([]T, 8, 8), cur: -1}
	return s
}

func NewWithCap[T any](cap int) *Stack[T] {
	s := &Stack[T]{}
	s.size = 0
	s.top = &Node[T]{elements: make([]T, cap, cap), cur: -1}
	return s
}

func (as *Stack[T]) Clear() {
	as.size = 0

	as.top.down = nil
	as.top.cur = -1
}

func (as *Stack[T]) Empty() bool {
	return as.size == 0
}

func (as *Stack[T]) Size() uint {
	return as.size
}

// String 左为Top
func (as *Stack[T]) String() string {

	return fmt.Sprintf("%v", as.Values())
}

func (as *Stack[T]) Values() []T {
	result := make([]T, as.size, as.size)

	cur := as.top
	n := as.size - 1
	for ; cur != nil; cur = cur.down {
		for i, _ := range cur.elements {
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

func (as *Stack[T]) Push(v T) {
	as.grow()
	as.top.cur++
	as.top.elements[as.top.cur] = v
	as.size++
}

func (as *Stack[T]) Pop() (interface{}, bool) {
	if as.size <= 0 {
		return nil, false
	}

	as.size--
	if as.cacheRemove() {
		return as.cache.elements[as.cache.cur], true
	}

	epop := as.top.elements[as.top.cur]
	as.top.cur--
	return epop, true
}

func (as *Stack[T]) Peek() (interface{}, bool) {
	if as.size <= 0 {
		return nil, false
	}
	return as.top.elements[as.top.cur], true
}
