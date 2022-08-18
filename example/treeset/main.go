package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/set/treeset"
)

func main() {
	set := treeset.New(compare.Any[int])

	for _, v := range []int{5, 25, 4, 11, 0} {
		set.Add(v) // return true
	}
	log.Println(set.Add(5))  // false
	log.Println(set.Add(10)) // true

	log.Println(set.String()) // (0, 4, 5, 10, 11, 25)

	log.Println(set.Contains(11)) // fasle
	log.Println(set.Contains(3))  // true

	iter := set.Iterator()
	iter.SeekGE(5)
	for iter.Vaild() {
		log.Println(iter.Value()) // 5 10 11 25
		iter.Next()
	}

	iter.SeekLT(10)
	for iter.Vaild() {
		log.Println(iter.Value()) // 5 4 0
		iter.Prev()
	}
}
