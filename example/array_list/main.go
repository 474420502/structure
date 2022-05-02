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
	log.Println("Iterator{Index IndexTo ToHead ToTail Value}")
	iter := l.Iterator()
	iter.ToHead()                           // to head
	log.Println(iter.Value(), iter.Index()) // value:0 index:0
	iter.ToTail()                           // to tail
	log.Println(iter.Value(), iter.Index()) // value:8 index:4

	iter.IndexTo(1)           //
	log.Println(iter.Value()) // 2

	log.Println("Iterator{Prev Next Vaild}")
	iter = l.Iterator()
	iter.ToHead()
	for iter.Vaild() {
		log.Println(iter.Value()) // 0 2 4 6 8
		iter.Next()
	}

	iter.ToTail()
	for iter.Vaild() {
		log.Println(iter.Value()) // 8 6 4 2 0
		iter.Prev()
	}

	log.Println("Iterator{SetValue Swap}")
	iter1 := l.Iterator()
	iter1.ToHead()
	iter1.SetValue(8)
	iter2 := l.Iterator()
	iter2.ToTail()
	iter2.SetValue(0)
	log.Println(ShowList(l)) // [0 2 4 6 8] -> [8 2 4 6 0]
	iter1.Swap(iter2)
	log.Println(ShowList(l)) // [8 2 4 6 0] -> [0 2 4 6 8]

	log.Println("Iterator{RemoveToNext RemoveToPrev}")
	iter = l.Iterator()
	iter.IndexTo(2)           // cur: 4
	iter.RemoveToNext()       // cur:6
	log.Println(iter.Value()) // 6  [0 2 6 8]
	iter.RemoveToPrev()       // cur: 2
	log.Println(iter.Value()) // 2  [0 2 8]
	log.Println(ShowList(l))  // [0 2 8]
}

func main3() {
	l := arraylist.New(compare.Any[int])
	for i := 0; i < 10; i += 2 {
		l.Push(i)
	}
	log.Println("Iterator{Index IndexTo ToHead ToTail Value}")
	iter := l.CircularIterator()
	iter.ToHead()                           // to head
	log.Println(iter.Value(), iter.Index()) // value:0 index:0
	iter.ToTail()                           // to tail
	log.Println(iter.Value(), iter.Index()) // value:8 index:4

	iter.IndexTo(1)           //
	log.Println(iter.Value()) // 2

	log.Println("Iterator{Prev Next Vaild}")
	iter = l.CircularIterator()

	var result []int
	var count int

	count = 0
	iter.ToHead()
	for iter.Vaild() {
		result = append(result, iter.Value())
		iter.Next()
		if iter.Value() == 0 {
			count++
			if count >= 2 {
				break
			}
		}
	}
	log.Println(result) // [2 4 6 8 0 2 4 6 8]

	result = nil
	count = 0
	iter.ToTail()
	for iter.Vaild() {
		result = append(result, iter.Value())
		iter.Prev()
		if iter.Value() == 8 {
			count++
			if count >= 2 {
				break
			}
		}
	}
	log.Println(result) // [8 6 4 2 0 8 6 4 2 0]

	log.Println("Iterator{SetValue Swap}")
	iter1 := l.CircularIterator()
	iter1.ToHead()
	iter1.SetValue(8)
	iter2 := l.CircularIterator()
	iter2.ToTail()
	iter2.SetValue(0)
	log.Println(ShowList(l)) // [0 2 4 6 8] -> [8 2 4 6 0]
	iter1.Swap(iter2)
	log.Println(ShowList(l)) // [8 2 4 6 0] -> [0 2 4 6 8]

	log.Println("Iterator{RemoveToNext RemoveToPrev}")
	iter = l.CircularIterator()
	iter.IndexTo(2)           // cur: 4
	iter.RemoveToNext()       // cur:6
	log.Println(iter.Value()) // 6  [0 2 6 8]
	iter.RemoveToPrev()       // cur: 2
	log.Println(iter.Value()) // 2  [0 2 8]
	log.Println(ShowList(l))  // [0 2 8]
}

func ShowList[T any](a *arraylist.ArrayList[T]) string {
	var result []T
	a.Traverse(func(idx uint, value T) bool {
		result = append(result, value)
		return true
	})
	return fmt.Sprintf("%v", result)
}
