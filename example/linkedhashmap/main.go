package main

import (
	"log"

	"github.com/474420502/structure/map/linkedhashmap"
)

func main() {
	// []
	lmap := linkedhashmap.New()

	log.Println("Put PushFront PushBack Get String")
	// [{1:1}]
	log.Println(lmap.Put(1, 1)) // true
	// [{1:1} {2:2}]
	log.Println(lmap.Put(2, 2)) // true
	// [{1:1} {2:2}]
	log.Println(lmap.Put(1, 1)) // false
	// [{1:1} {2:2} {3:3}]
	log.Println(lmap.PushBack(3, 3)) // true
	// [{0:0} {1:1} {2:2} {3:3}]
	log.Println(lmap.PushFront(0, 0)) // true

	log.Println(lmap.Get(0))   // 1 true
	log.Println(lmap.Get(100)) // nil false

	log.Println(lmap.String()) // [{0:0} {1:1} {2:2} {3:3}]

	log.Println("Set SetFront SetBack")

	// [{0:0} {1:1} {2:2} {3:0}]
	log.Println(lmap.Set(3, 0)) // true

	// [{0:0} {1:1} {2:2} {3:0}]
	log.Println(lmap.Set(4, 4)) // false.  not insert

	// [{0:0} {1:1} {2:2} {3:0} {4:4}]
	log.Println(lmap.SetBack(4, 4)) // false.  insert

	// [{-1:-1} {0:0} {1:1} {2:2} {3:0} {4:4}]
	log.Println(lmap.SetFront(-1, -1)) // false. insert

	// [{3:3} {-1:-1} {0:0} {1:1} {2:2} {4:4}]
	log.Println(lmap.SetFront(3, 3)) // true. move the slice to head

	// [{-1:-1} {0:0} {1:1} {2:2} {4:4} {3:3}]
	log.Println(lmap.SetBack(3, 3)) // true. move the slice to tail

	log.Println("Keys Values Slices")

	log.Println(lmap.Keys())   // [-1 0 1 2 4 3]
	log.Println(lmap.Values()) // [-1 0 1 2 4 3]
	log.Println(lmap.Slices()) // [{-1:-1} {0:0} {1:1} {2:2} {4:4} {3:3}]

	log.Println("Size Empty Clear")

	log.Println(lmap.Size())  // 6
	log.Println(lmap.Empty()) // false
	lmap.Clear()
	log.Println(lmap.Empty()) // true
}
