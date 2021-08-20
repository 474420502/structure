package treequeue

import "fmt"

type Slice struct {
	key   interface{}
	value interface{}
}

func (s *Slice) String() string {
	return fmt.Sprintf("(%v,%v)", s.key, s.value)
}

func (s *Slice) Key() interface{} {
	return s.key
}

func (s *Slice) Value() interface{} {
	return s.value
}

func (s *Slice) SetValue(v interface{}) {
	s.value = v
}
