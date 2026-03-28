package main

import (
	"log"

	"github.com/474420502/structure/map/linkedhashmap"
)

func main() {
	// []
	lmap := linkedhashmap.New()

	log.Println("InsertIfAbsent PushFront PushBack Get String")
	// [{1:1}]
	log.Println(lmap.InsertIfAbsent(1, 1)) // true
	// [{1:1} {2:2}]
	log.Println(lmap.InsertIfAbsent(2, 2)) // true
	// [{1:1} {2:2}]
	log.Println(lmap.InsertIfAbsent(1, 1)) // false
	// [{1:1} {2:2} {3:3}]
	log.Println(lmap.PushBack(3, 3)) // true
	// [{0:0} {1:1} {2:2} {3:3}]
	log.Println(lmap.PushFront(0, 0)) // true

	log.Println(lmap.Get(0))   // 0 true
	log.Println(lmap.Get(100)) // nil false

	log.Println(lmap.String()) // [{0:0} {1:1} {2:2} {3:3}]

	log.Println("Upsert Set SetFront SetBack")

	// [{0:0} {1:1} {2:2} {3:0}]
	log.Println(lmap.Upsert(3, 0)) // true

	// [{0:0} {1:1} {2:2} {3:0} {4:4}]
	log.Println(lmap.Upsert(4, 4)) // false. insert missing key at back

	// [{0:0} {1:1} {2:2} {3:0} {4:40}]
	log.Println(lmap.SetBack(4, 40)) // true. update existing key and move to back

	// [{-1:-1} {0:0} {1:1} {2:2} {3:0} {4:40}]
	log.Println(lmap.SetFront(-1, -1)) // false. insert

	// [{3:3} {-1:-1} {0:0} {1:1} {2:2} {4:40}]
	log.Println(lmap.SetFront(3, 3)) // true. move the slice to head

	// [{-1:-1} {0:0} {1:1} {2:2} {4:40} {3:3}]
	log.Println(lmap.SetBack(3, 3)) // true. move the slice to tail

	log.Println("Keys Values Slices")

	log.Println(lmap.Keys())   // [-1 0 1 2 4 3]
	log.Println(lmap.Values()) // [-1 0 1 2 40 3]
	log.Println(lmap.Slices()) // [{-1:-1} {0:0} {1:1} {2:2} {4:40} {3:3}]

	log.Println("Len Delete Empty Clear")

	log.Println(lmap.Len())   // 6
	log.Println(lmap.Delete(2))
	log.Println(lmap.Empty()) // false
	lmap.Clear()
	log.Println(lmap.Empty()) // true
}
