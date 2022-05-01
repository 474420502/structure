package main

import (
	"fmt"
	"log"

	"github.com/474420502/structure/compare"
	arraylist "github.com/474420502/structure/list/array_list"
)

func main() {
	l := arraylist.New(compare.Any[int])
	log.Println("Push PushFront PushBack Set")
	l.Push(3)
	l.Push(5)
	l.Push(7)
	l.PushBack(7)
	l.PushFront(2)
	l.PushFront(2)
	l.Set(1, 10)

	log.Println("Size Empty Values String")
	log.Println(l.String()) // [2 10 3 5 7 7]
	log.Println(l.Size())   // 6
	log.Println(l.Empty())  // false
	log.Println(l.Values()) // [2 10 3 5 7 7]

	log.Println("Back Front Index Contains")
	log.Println(l.Back())         // 7 true
	log.Println(l.Front())        // 2 true
	log.Println(l.Index(1))       // 10
	log.Println(l.Contains(7))    // 2
	log.Println(l.Contains(1))    // 0
	log.Println(l.Contains(2, 7)) // 3

	log.Println("Traverse")
	l.Traverse(func(idx uint, value int) bool {
		log.Print("(", idx, value, ")") // (0 2) (1 10) (2 3) (3 5) (4 7) (5 7)
		return true
	})

	log.Println("Remove PopBack PopFront Clear")
	// [2 10 3 5 7 7]
	l.Remove(0)              // [10 3 5 7 7]
	l.PopBack()              // [10 3 5 7]
	l.PopFront()             // [3 5 7]
	l.Clear()                // []
	log.Println(ShowList(l)) // []

}

func main2() {
	l := arraylist.New(compare.Any[int])
	for i := 0; i < 10; i += 2 {
		l.Push(i)
	}
	log.Println("Iterator")
}

func ShowList[T any](a *arraylist.ArrayList[T]) string {
	var result []T
	a.Traverse(func(idx uint, value T) bool {
		result = append(result, value)
		return true
	})
	return fmt.Sprintf("%v", result)
}
