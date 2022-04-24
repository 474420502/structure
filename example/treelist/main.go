package main

import (
	"fmt"
	"log"

	"github.com/474420502/structure/search/treelist"
)

func main() {

	// API A similar to B
	queue := treelist.New()

	log.Println("Put")
	queue.Put([]byte("zero"), 0) // true
	queue.Put([]byte("apple"), 4)
	queue.Put([]byte("word1"), 1)
	queue.Put([]byte("word2"), 2)
	queue.Put([]byte("boy"), 3)

	log.Println("Slices")
	var results []string
	for _, slice := range queue.Slices() {
		results = append(results, fmt.Sprintf("%s %d", string(slice.Key), slice.Value))
	}
	log.Println(results) // [apple 4 boy 3 word1 1 word2 2 zero 0]. values in order

	log.Println("Get")
	queue.Get([]byte("apple"))  // 4, true
	queue.Get([]byte("apple1")) // nil, false

	log.Println("Head Tail")
	log.Println(queue.Head()) // apple 4
	log.Println(queue.Tail()) // zero 0

	log.Println("Index IndexOf Size")
	log.Println(queue.Index(0))                // apple
	log.Println(queue.IndexOf([]byte("boy")))  // 1
	log.Println(queue.Index(queue.Size() - 1)) // zero

	log.Println("Iterator: {Valid Next SeekGE}")
	iter := queue.Iterator()
	if iter.SeekGE([]byte("word1")) { // similar to rocksdb pebble leveldb skiplist
		for ; iter.Valid(); iter.Next() { // Vaiid Next
			log.Println(string(iter.Key())) // log: word1 word2 zero
			// you can limit by yourself
		}
	}

	log.Println("Iterator: {Valid Prev SeekLE}")
	if iter.SeekLE([]byte("word")) { // similar to rocksdb pebble leveldb skiplist
		for ; iter.Valid(); iter.Prev() { // Vaiid Next
			log.Println(string(iter.Key())) // log: boy apple
			// you can limit by yourself
		}
	}

	log.Println("Iterator: {SeekToFirst SeekToLast}")
	iter.SeekToFirst()              // get first item
	log.Println(string(iter.Key())) // apple

	iter.SeekToLast()               // get last item
	log.Println(string(iter.Key())) // zero

	log.Println("PutDuplicate")
	queue.PutDuplicate([]byte("boy"), 10, func(exists *treelist.Slice) {
		exists.Value = 100
	})
	queue.Get([]byte("boy")) // boy 100 true
}
