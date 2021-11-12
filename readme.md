# Golang Data Structure

## Provide a powerful and easy-to-use data structure

* PriorityQueue

```go
package main

import (
	"log"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
	treequeue "github.com/474420502/structure/queue/priority"
)

func main() {
	// all api is fast
	r := random.New(1636706158629652669)
	queue := treequeue.New(compare.Int)

	for i := 0; i < 10; i++ {
		v := r.Intn(10)
		queue.Put(v, v)
	}

	log.Println(queue.String())     // [0 0 0 1 1 2 4 4 4 8]
	log.Println(queue.RemoveHead()) // (0,0) PopHead treequeue.Slice{} Key = 0 Value = 0
	log.Println(queue.String())     // [0 0 1 1 2 4 4 4 8]
	log.Println(queue.RemoveTail()) // (8,8) PopTail treequeue.Slice{} Key = 8 Value = 8
	log.Println(queue.String())     // [0 0 1 1 2 4 4 4]

	// Top 5 (Top k). Keep Top 5 Size
	queue.RemoveRangeByIndex(0, (queue.Size()-5)-1) // Remove 0 - 2 [0,2]
	log.Println(queue.String())                     // [1 2 4 4 4]

	for i := 0; i < 100; i++ {
		v := r.Intn(100)
		queue.Put(v, v)
	}

	// Top 5 (Top k). Keep data in sliding window. (date, rank ....)
	queue.RemoveRangeByIndex(0, (queue.Size()-5)-1)
	log.Println(queue.String()) // [94 94 94 95 97]

	log.Println(queue.Index(0).Key()) // 94

	queue = treequeue.New(compare.IntDesc) //  From big to small
	for i := 0; i < 10; i++ {
		queue.Put(i, i)
	}
	log.Println(queue.String()) // [9 8 7 6 5 4 3 2 1 0]

	// Keep the Index. 4-7 [4,7]
	queue.ExtractByIndex(4, 7)
	log.Println(queue.String()) // [5 4 3 2]
	queue.Clear()

	for i := 0; i < 100; i++ {
		v := r.Intn(100)
		queue.Put(v, v)
	}

	// Extract 70 - 30.[70, 30] all values
	queue.Extract(70, 30)       // becase IntDesc. big is low
	log.Println(queue.String()) // [70 68 68 66 66 66 64 64 63 60 59 57 57 56 55 49 48 45 44 43 42 42 41 41 39 39 38 37 36 35 34 33 33]

	// Get Top N By Index
	log.Println(queue.Index(0).Key(), queue.Index(1).Key()) // 70 68
}
```

* Treelist is like skiplist. better than skiplist

```go
package main

import (
	"log"

	"github.com/474420502/structure/search/treelist"
)

func main() {
	// all api is fast

	// API like treequeue
	queue := treelist.New()

	queue.Put([]byte("zero"), 0)
	queue.Put([]byte("apple"), 4)
	queue.Put([]byte("word1"), 1)
	queue.Put([]byte("word2"), 2)
	queue.Put([]byte("boy"), 3)

	iter := queue.Iterator()
	if iter.SeekGE([]byte("word1")) { // like rocksdb pebble leveldb skiplist
		for ; iter.Valid(); iter.Next() {
			log.Println(string(iter.Key())) // word1 word2 zero. you can limit by yourself
		}
	}

	iter.SeekToFirst()              // get first item
	log.Println(string(iter.Key())) // apple
}

```