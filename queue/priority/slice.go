package treequeue

import "fmt"

type Slice[T any] struct {
	key   T
	value interface{}
}

func (s *Slice[T]) String() string {
	return fmt.Sprintf("(%v,%v)", s.key, s.value)
}

func (s *Slice[T]) Key() T {
	return s.key
}

func (s *Slice[T]) Value() interface{} {
	return s.value
}

func (s *Slice[T]) SetValue(v interface{}) {
	s.value = v
}
