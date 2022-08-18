package main

import (
	"log"

	lastack "github.com/474420502/structure/stack/listarray"
)

func main() {

	st := lastack.NewWithCap[int](16)

	log.Println("Push String Size")
	for i := 0; i < 10; i += 2 {
		st.Push(i)
	}
	log.Println(st.String()) // [0 2 4 6 8]
	log.Println(st.Size())   // 5

	log.Println("Peek Pop Empty Clear")
	log.Println(st.Peek()) // 8 true
	log.Println(st.Pop())  // 8 true
	st.Clear()
	log.Println(st.Empty()) // true
	log.Println(st.Peek())  // 0 false
	log.Println(st.Pop())   // 0 false

	log.Println("Values")
	for i := 0; i < 10; i += 2 {
		st.Push(i)
	}
	log.Println(st.Values()) // [0 2 4 6 8]
}
