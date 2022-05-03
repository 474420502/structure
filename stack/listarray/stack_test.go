package lastack

import (
	"container/list"
	"log"
	"testing"

	"github.com/474420502/random"
)

func TestForce(t *testing.T) {

	rand := random.New(t.Name())
	s1 := New[int]()
	s2 := list.New()

	for i := 0; i < 2000; i++ {
		v := rand.Intn(100)
		s1.Push(v)
		s2.PushBack(v)
	}

	for !s1.Empty() {

		v1, _ := s1.Peek()
		v2 := s2.Back().Value
		if v1 != v2 {
			panic("")
		}

		v1, _ = s1.Pop()
		if v1 != s2.Remove(s2.Back()) {
			panic("")
		}

		if s1.Size() != uint(s2.Len()) {
			panic("")
		}

		if rand.Intn(2) == 0 {
			v := rand.Intn(100)
			s1.Push(v)
			s2.PushBack(v)
		}
	}

	if _, ok := s1.Peek(); ok != false {
		panic("")
	}

	if _, ok := s1.Pop(); ok != false {
		panic("")
	}

	s1.Clear()
	s2.Init()

	if s1.Size() != uint(s2.Len()) {
		panic("")
	}

	for i := 0; i < 20; i++ {
		v := rand.Intn(1000)
		s1.Push(v)
		s2.PushBack(v)
	}

	for _, v := range s1.Values() {
		if v != s2.Remove(s2.Front()) {
			panic("")
		}
	}

}

func TestString(t *testing.T) {
	rand := random.New(t.Name())
	s1 := New[int]()

	for i := 0; i < 10; i++ {
		v := rand.Intn(100)
		s1.Push(v)
	}

	log.Println(s1.String())
}
