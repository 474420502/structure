package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
)

func main() {

	tree := avl.New(compare.Any[int]) // create a  object

	log.Println("Put Set")
	tree.Put(0, 0) // put key value into the tree
	tree.Put(4, 4)
	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Set(3, 3)

	log.Println("String")
	log.Println(tree.String())
	//
	// │       ┌── 4
	// │   ┌── 3
	// │   │   └── 2
	// └── 1
	//     └── 0

	log.Println("Traverse")
	tree.Traverse(func(k int, v interface{}) bool {
		log.Println(k, v) // 0 0 1 1 2 2 3 3 4 4
		return true
	})

	log.Println("Get")
	log.Println(tree.Get(4)) // 4,true
	log.Println(tree.Get(5)) // nil,false

	iter := tree.Iterator() // create a the object of iterator

	log.Println("SeekToFirst")
	iter.SeekToFirst() // seek to the first item
	// range first to last
	for iter.Vaild() {
		log.Println(iter.Value()) // 0 1 2 3 4
		iter.Next()
	}

	log.Println("SeekToLast")
	iter.SeekToLast() // seek to the last item
	// range last to first
	for iter.Vaild() {
		log.Println(iter.Value()) // 4 3 2 1 0
		iter.Prev()
	}

	log.Println("Remove")
	log.Println(tree.Remove(2)) // 2, true
	log.Println(tree.Remove(2)) // <nil>, false

	log.Println(tree.String())
	// 	│       ┌── 4
	// 	│   ┌── 3
	// 	└── 1
	// 		└── 0

	iter.Next() // the pos of iter is before 0. so need call Next()
	for iter.Vaild() {
		log.Println(iter.Value()) // 0 1 3 4
		iter.Next()
	}

	log.Println("SeekGE")
	iter.SeekGE(3) // seek to the iterator value >= 3
	// range 3 to 4
	for iter.Vaild() {
		log.Println(iter.Value()) // 3 4
		iter.Next()
	}

	log.Println("SeekLE")
	iter.SeekLE(2) // seek to the iterator value >= 3
	// range 1 to 0
	for iter.Vaild() {
		log.Println(iter.Value()) // 1 0
		iter.Prev()
	}

	log.Println(tree.Get(2)) // nil,false
	tree.Set(0, 1)
	log.Println(tree.Get(0)) // 1, true

	tree.Clear() // clear all the data of tree
}
