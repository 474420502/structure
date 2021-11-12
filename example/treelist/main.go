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
