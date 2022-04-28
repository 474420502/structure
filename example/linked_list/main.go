package main

import (
	"log"

	linkedlist "github.com/474420502/structure/list/linked_list"
)

func main() {
	l := linkedlist.New[int]()

	for i := 0; i < 5; i++ {
		l.Push(i)
	}

	log.Println(l) // [0 1 2 3 4]

	iter := l.Iterator()
	for iter.ToTail(); iter.Vaild(); iter.Prev() {
		log.Println(iter.Value()) // 4 3 2 1 0
	}

	log.Println("CircularIterator")

	citer := l.CircularIterator()
	citer.ToHead()
	for i := 0; i < 11 && citer.Vaild(); i++ {
		citer.Next()
		log.Println(citer.Value()) // 0 1 2 3 4 0 1 2 3 4 0
	}
}
