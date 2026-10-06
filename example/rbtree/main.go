package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/rbtree"
)

func main() {
	// rbtree expects a standard comparator such as compare.Any:
	// negative means less-than, positive means greater-than, zero means equal.
	tree := rbtree.New[int, int](compare.Any[int])

	log.Println("Put InsertIfAbsent Set Upsert")
	tree.Put(0, 0)
	tree.Put(4, 4)
	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Set(3, 3)
	log.Println(tree.InsertIfAbsent(3, 99)) // false, key 3 already exists
	log.Println(tree.Upsert(3, 33))         // true, existing value replaced
	log.Println(tree.Len())                 // 5

	log.Println("Traverse")
	tree.Traverse(func(k, v int) bool {
		log.Println(k, v) // 0 0, 1 1, 2 2, 3 33, 4 4
		return true
	})

	log.Println("Get")
	log.Println(tree.Get(4)) // 4 true
	log.Println(tree.Get(5)) // 0 false

	iter := tree.Iterator()

	log.Println("SeekToFirst")
	iter.SeekToFirst()
	for iter.Valid() {
		log.Println(iter.Value()) // 0 1 2 33 4
		iter.Next()
	}

	log.Println("SeekToLast")
	iter.SeekToLast()
	for iter.Valid() {
		log.Println(iter.Value()) // 4 33 2 1 0
		iter.Prev()
	}

	log.Println("SeekGE")
	log.Println(iter.SeekGE(3)) // true, exact key exists
	for iter.Valid() {
		log.Println(iter.Value()) // 33 4
		iter.Next()
	}

	log.Println("SeekLT")
	iter.SeekToLast()
	log.Println(iter.SeekLT(3)) // true, key 3 existed; iterator lands on 2
	log.Println(iter.Key())     // 2

	log.Println("Delete")
	log.Println(tree.Delete(3)) // 33 true
	log.Println(tree.Delete(3)) // 0 false
	log.Println(tree.Len())     // 4

	tree.Clear()
}
