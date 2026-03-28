package main

import (
	"log"

	"github.com/474420502/structure/compare"
	treequeue "github.com/474420502/structure/queue/priority"
)

func main() {
	queue := treequeue.New[int, int](compare.Any[int])

	for _, v := range []int{5, 1, 9, 1, 4, 7} {
		queue.InsertIfAbsent(v, v)
	}

	log.Println(queue.Values())  // [1 4 5 7 9]
	log.Println(queue.Upsert(7, 70)) // true
	log.Println(queue.Get(7))        // 70 true
	log.Println(queue.Len())         // 5
	log.Println(queue.Delete(5))     // 5 true
	log.Println(queue.Values())      // [1 4 70 9]

	duplicateQueue := treequeue.New[int, int](compare.Any[int])
	for _, v := range []int{5, 1, 9, 1, 4, 7} {
		duplicateQueue.Put(v, v)
	}
	log.Println(duplicateQueue.Values()) // [1 1 4 5 7 9]
	log.Println(duplicateQueue.Gets(1))  // [1 1]

	iter := duplicateQueue.Iterator()
	log.Println(iter.SeekGEExact(4)) // true
	for iter.Valid() {
		log.Println(iter.Value()) // 4 5 7 9
		iter.Next()
	}

	// Top 5 (Top k). Keep Top 5 Size
	// queue.Remove()
	// queue.RemoveRangeByIndex(0, (queue.Size()-5)-1) // Remove 0 - 2 [0,2]
	// log.Println(queue.String())                     // [1 2 4 4 4]

	// // Top 5 (Top k). Keep data in sliding window. (date, rank ....)
	// queue.RemoveRangeByIndex(0, (queue.Size()-5)-1)
	// log.Println(queue.String()) // [94 94 94 95 97]

	// log.Println(queue.Index(0).Key()) // 94

	// queue = treequeue.New(compare.AnyDesc[int]) //  From big to small
	// for i := 0; i < 10; i++ {
	// 	queue.Put(i, i)
	// }
	// log.Println(queue.String()) // [9 8 7 6 5 4 3 2 1 0]

	// // Keep the Index. 4-7 [4,7]
	// queue.ExtractByIndex(4, 7)
	// log.Println(queue.String()) // [5 4 3 2]
	// queue.Clear()

	// for i := 0; i < 100; i++ {
	// 	v := r.Intn(100)
	// 	queue.Put(v, v)
	// }

	// // Extract 70 - 30.[70, 30] all values
	// queue.Extract(70, 30)       // becase IntDesc. big is low
	// log.Println(queue.String()) // [70 68 68 66 66 66 64 64 63 60 59 57 57 56 55 49 48 45 44 43 42 42 41 41 39 39 38 37 36 35 34 33 33]

	// // Get Top N By Index
	// log.Println(queue.Index(0).Key(), queue.Index(1).Key()) // 70 68
}
