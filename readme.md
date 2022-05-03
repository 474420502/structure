# Golang Data Structure

## Provide a powerful and easy-to-use data structure

- [Tree](./tree)

	- [AVL](./tree/avl/doc.md) is a self-balancing binary search tree (BST)

	- [Heap](./tree/heap/doc.md)  a heap is a specialized tree-based data structure which is essentially an almost complete tree that satisfies the heap property

	- [IndexTree](./tree/indextree/doc.md)  is a self-balancing binary search tree (BST). and have some special features

	- [TreeList](./tree/treelist/doc.md)  is a self-balancing binary search tree (BST). and have some special features 

- [List](./list)

	- [ArrayList](./list/array_list/doc.md)  list are based on arrays 
	- [LinkedList](./list/linked_list/doc.md)  linkedlist are based on list 

- [Stack](./stack)

	- [ArrayStack](./stack/array/doc.md)  stack are based on arrays 
	- [ListStack](./stack/list/doc.md)  stack are based on list 
	- [LAStack](./stack/listarray/doc.md)  stack are based on both list and array


- [Search](./search)
	- [TreeList](./search/treelist/doc.md) is a self-balancing binary search tree (BST). and have some special features. key only support the type of bytes 
	
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

* bloom this a filter. is easy to use

```go 
package main

import (
	"log"
	"math/rand"

	"github.com/474420502/random"
	"github.com/474420502/structure/filter/bloom"
)

var basechars []byte = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789")

func main() {
	var keynum = uint64(100000)

	bloom := bloom.New(keynum * 10) // create bloom
	r := random.New()
	var collect [][]byte // collect the bytes

	var truecount int
	for i := uint64(0); i < keynum; i++ {
		var chars []byte
		// random create bytes
		for n := 0; n < rand.Intn(32)+5; n++ {
			s := r.Intn(len(basechars))
			chars = append(chars, basechars[s])
		}
		collect = append(collect, chars)
		if bloom.AddBytes(chars) { // add the bloom
			truecount++
		}
	}

	for _, chars := range collect {
		if !bloom.ContainsBytes(chars) { // is in bloom?
			log.Panic("bloom.ContainsBytes error")
		}
	}

	hitsize := bloom.HitSize()

	buf := bloom.Encode() // encode to buffer
	bloom.Reset()
	bloom.Decode(buf) // decode from buffer

	ratio := float64(keynum-bloom.HitSize()) / float64(keynum)
	if ratio == 0 || bloom.HitSize() == 0 || bloom.HitSize() != hitsize {
		log.Println("Encode and Decode error")
	}

	log.Println(bloom.HitSize(), ratio) // 95192(bit used by bloom) 0.04843(percentage of duplicates)
}
```

## The version of generics.

* AVL 

```go
package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
)

func main() {

	// all api is fast

	// API like treequeue
	tree := avl.New(compare.Any[int])

	tree.Put(0, 0)
	tree.Put(4, 4)
	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Put(3, 3)

	log.Println(tree.String())
	//
	// │       ┌── 4
	// │   ┌── 3
	// │   │   └── 2
	// └── 1
	//     └── 0

	log.Println(tree.Get(4)) // 4,true

	iter := tree.Iterator()
	iter.SeekToFirst()

	for iter.Vaild() {
		log.Println(iter.Value()) // 0 1 2 3 4
		iter.Next()
	}

	iter.SeekToLast()
	for iter.Vaild() {
		log.Println(iter.Value()) // 4 3 2 1 0
		iter.Prev()
	}

	log.Println(tree.Remove(2)) // 2, true
	log.Println(tree.Remove(2)) // <nil>, false

	log.Println(tree.String())
	// 	│       ┌── 4
	// 	│   ┌── 3
	// 	└── 1
	// 		└── 0

	iter.Next() // the pos of iter is before 0. so need call Next()
	for iter.Vaild() {
		log.Println(iter.Value()) // 0 1 3 4
		iter.Next()
	}

	log.Println(tree.Get(2)) // nil,false
}

```