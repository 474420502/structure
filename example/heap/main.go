package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/heap"
)

func main() {
	h := heap.New(compare.AnyDesc[int])

	// Put
	for _, v := range []int{5, 2, 3, 45, 10} {
		h.Put(v)
	}
	log.Println(h.Size()) // 5

	// Pop
	var values []int
	for v, ok := h.Pop(); ok; v, ok = h.Pop() {
		values = append(values, v)
	}
	log.Println(values)    // [45 10 5 3 2]
	log.Println(h.Empty()) // true

	h.Put(1)
	log.Println(h.Empty()) // false
	h.Clear()
	log.Println(h.Empty()) // true
}
