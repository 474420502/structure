package main

import (
	"log"

	"github.com/474420502/structure/map/hashmap"
)

func main() {

	hm := hashmap.New()

	log.Println("Put Set Get String")
	log.Println(hm.Put(1, 1))       // true
	log.Println(hm.Put(1, 1))       // false
	log.Println(hm.Put('a', 2))     // true
	log.Println(hm.Put("apple", 4)) // true
	hm.Set("apple", 3)
	log.Println(hm.String()) // map[1:1 97:2 apple:3]

	log.Println(hm.Get(1))       // 1 true
	log.Println(hm.Get(2))       // <nil> false
	log.Println(hm.Get('a'))     // 2 true
	log.Println(hm.Get("apple")) // 3 true

	log.Println("Keys Values Slices")
	log.Println(hm.Keys())   // [1 'a'(97) apple]
	log.Println(hm.Values()) // [1 2 3]
	log.Println(hm.Slices()) // [{1 1} {97 2} {apple 3}]

	log.Println("Size Remove")
	log.Println(hm.Size())
	hm.Remove(1)
	log.Println(hm.String()) // map[97:2 apple:3]
	hm.Remove('a')
	log.Println(hm.String()) // map[apple:3]
	hm.Remove(2)
	log.Println(hm.String()) // map[apple:3]

	log.Println("Clear Empty")
	hm.Clear()
	log.Println(hm.Empty())  // true
	log.Println(hm.String()) // map[]
}
