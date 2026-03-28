package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/set/treeset"
)

func main() {
	set := treeset.New[int, int](compare.AnyEx[int])

	for _, v := range []int{5, 25, 4, 11, 0} {
		set.InsertIfAbsent(v, v)
	}
	log.Println(set.InsertIfAbsent(5, 5))   // false
	log.Println(set.InsertIfAbsent(10, 10)) // true
	log.Println(set.Upsert(10, 100))        // true

	log.Println(set.String()) // (0, 4, 5, 10, 11, 25)

	log.Println(set.Contains(11)) // true
	log.Println(set.Contains(3))  // false

	iter := set.Iterator()
	log.Println(iter.SeekGEExact(5)) // true
	for iter.Valid() {
		log.Println(iter.Value()) // 5 100 11 25
		iter.Next()
	}

	log.Println(iter.SeekLTExact(10)) // true
	for iter.Valid() {
		log.Println(iter.Value()) // 5 4 0
		iter.Prev()
	}
}
