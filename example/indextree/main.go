package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/indextree"
)

func main1() {

	tree := indextree.New(compare.Any[int]) // create a  object

	log.Println("Put Set")
	tree.Put(0, 0) // put key value into the tree
	tree.Put(4, 4)
	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Set(3, 3) // Set value

	log.Println("Values")
	log.Println(tree.Values())

	log.Println("Size")
	tree.Size() // 5

	log.Println("Index IndexOf")
	log.Println(tree.Index(0))   // 0 0
	log.Println(tree.Index(4))   // 4 4
	log.Println(tree.IndexOf(2)) // 2 like RankOf

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

	log.Println("Remove")
	log.Println(tree.Remove(2)) // 2, true
	log.Println(tree.Remove(2)) // <nil>, false
	tree.Put(5, 5)
	tree.Put(6, 6)
	log.Println(tree.String())
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	└── 4
	// 		│   ┌── 3
	// 		└── 1
	// 			└── 0
	log.Println("RemoveIndex")
	tree.RemoveIndex(0)        // Remove head.
	log.Println(tree.String()) // 1 3 4 5 6
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	└── 4
	// 		│   ┌── 3
	// 		└── 1
	log.Println("RemoveIndex")
	tree.RemoveRange(2, 5)     // remove 2-5 key
	log.Println(tree.String()) // 1 6
	// 	└── 6
	// 		└── 1

	tree.Put(15, 15)
	tree.Put(16, 16)
	log.Println(tree.String())
	//	│       ┌── 16
	//	│   ┌── 15
	//	└── 6
	//	    └── 1
	log.Println("RemoveRangeByIndex")
	tree.RemoveRangeByIndex(1, 2)
	log.Println(tree.String()) // 1 16
	// 	└── 16
	// 		└── 1

	tree.Clear() // clear all the data of tree
}

func main() {
	tree := indextree.New(compare.Any[int]) // create a  object

	for i := 0; i < 7; i++ {
		tree.Put(i, i)
	}

	log.Println(tree.String())
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3
	// 		│   ┌── 2
	// 		└── 1
	// 			└── 0

	log.Println("Split")
	tree2 := tree.Split(3)
	log.Println(tree.String())
	// 	│   ┌── 2
	// 	└── 1
	// 		└── 0
	log.Println(tree2.String())
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3

	tree.Clear()
	for i := 0; i < 7; i++ {
		tree.Put(i, i)
	}

	log.Println("SplitContain")
	tree2 = tree.SplitContain(3)
	log.Println(tree.String())
	// 	└── 3
	// 		│   ┌── 2
	// 		└── 1
	// 		    └── 0
	log.Println(tree2.String())
	// 	│   ┌── 6
	// 	└── 5
	// 	    └── 4

	tree.Clear()
	for i := 0; i < 7; i++ {
		tree.Put(i, i)
	}
	log.Println(tree.String())
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3
	// 		│   ┌── 2
	// 		└── 1
	// 			└── 0

	log.Println("Trim")
	tree.Trim(2, 5) // keep the values that range with 2-5
	log.Println(tree.String())
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3
	// 	    └── 2

	tree.Clear()
	for i := 0; i < 7; i++ {
		tree.Put(i, i)
	}
	log.Println("TrimByIndex")
	tree.TrimByIndex(2, 5) // keep the values that index range with 2-5
	log.Println(tree.String())
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3
	// 	    └── 2
}
