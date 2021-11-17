package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/set/treeset"
)

func main() {
	set := treeset.New(compare.Int)

	for _, v := range []int{5, 25, 4, 11, 0} {
		set.Add(v) // return false
	}
	log.Println(set.Add(5))  // true
	log.Println(set.Add(10)) // false

	log.Println(set.String()) // (0, 4, 5, 11, 25)

	log.Println(set.Contains(11)) // true
	log.Println(set.Contains(3))  // false

	// iter := set.Iterator()
	// iter.SeekGE()
}
