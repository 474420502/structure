package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
)

func main() {

	// all api is fast

	// API like treequeue
	tree := avl.New(compare.Any[int])

	tree.Put(0, 0)
	tree.Put(4, 4)
	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Put(3, 3)

	log.Println(tree.String())
	//
	// │       ┌── 4
	// │   ┌── 3
	// │   │   └── 2
	// └── 1
	//     └── 0

	log.Println(tree.Get(4)) // 4,true

	iter := tree.Iterator()
	iter.SeekToFirst()

	for iter.Vaild() {
		log.Println(iter.Value()) // 0 1 2 3 4
		iter.Next()
	}

	iter.SeekToLast()
	for iter.Vaild() {
		log.Println(iter.Value()) // 4 3 2 1 0
		iter.Prev()
	}

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

	log.Println(tree.Get(2)) // nil,false
}
