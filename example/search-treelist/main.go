package main

import (
	"fmt"
	"log"

	"github.com/474420502/structure/search/treelist"
)

func main() {

	// API A similar to B
	tree1 := treelist.New()

	log.Println("Put Set")
	tree1.Put([]byte("zero"), 0) // true
	tree1.Put([]byte("apple"), 4)
	tree1.Put([]byte("word1"), 1)
	tree1.Put([]byte("word2"), 2)
	tree1.Set([]byte("boy"), 4) // boy 4
	tree1.Set([]byte("boy"), 3) // boy 3

	log.Println("Slices")
	var results []string
	for _, slice := range tree1.Slices() {
		results = append(results, Slice2String(&slice))
	}
	log.Println(results) //[{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}]. values in order

	log.Println("Get")
	tree1.Get([]byte("apple"))  // 4, true
	tree1.Get([]byte("apple1")) // nil, false

	log.Println("Head Tail")
	log.Println(tree1.Head()) // apple 4
	log.Println(tree1.Tail()) // zero 0

	log.Println("Index IndexOf Size")
	log.Println(tree1.Index(0))                // apple
	log.Println(tree1.IndexOf([]byte("boy")))  // 1
	log.Println(tree1.Index(tree1.Size() - 1)) // 4

	log.Println("Intersection UnionSets") //
	tree2 := treelist.New()
	// word1 1 word2 2 zero 0
	tree2.Set([]byte("word1"), 1)
	tree2.Set([]byte("word2"), 2)
	tree2.Set([]byte("zero3"), 3)

	tree3 := tree1.Intersection(tree2)            // Intersection
	log.Println(Tree2String(tree3), tree3.Size()) // [{word1:1} {word2:2}] 2

	tree3 = tree1.UnionSets(tree2)                // UnionSets
	log.Println(Tree2String(tree3), tree3.Size()) // [{apple:4} {boy:3} {word1:1} {word2:2} {zero:0} {zero3:3}] 6

	//[{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}]
	log.Println("Iterator: {Valid Next SeekGE}")
	iter := tree1.Iterator()
	if iter.SeekGE([]byte("word1")) { // key >= "word1"  similar to rocksdb pebble leveldb skiplist
		for ; iter.Valid(); iter.Next() { // Vaiid Next
			log.Println(string(iter.Key())) // log: word1 word2 zero
			// you can limit by yourself
		}
	}

	//[{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}]
	log.Println("Iterator: {Valid Next SeekGT}")
	if iter.SeekGT([]byte("word1")) { //  key > "word1". to word2
		for ; iter.Valid(); iter.Next() { // Vaiid Next
			log.Println(string(iter.Key())) // log: word2 zero
			// you can limit by yourself
		}
	}

	//[{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}]
	log.Println("Iterator: {Valid Prev SeekLE}")
	if !iter.SeekLE([]byte("word")) { // key <= "word". return false . key is "boy"
		for ; iter.Valid(); iter.Prev() { // Vaiid Next
			log.Println(string(iter.Key())) // log: boy apple
			// you can limit by yourself
		}
	}

	//[{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}]
	log.Println("Iterator: {Valid Prev SeekLT}")
	if iter.SeekLT([]byte("word1")) { // key < "word1". return tree key is "boy"
		for ; iter.Valid(); iter.Prev() { // Vaiid Next
			log.Println(string(iter.Key())) // log: boy apple
			// you can limit by yourself
		}
	}

	log.Println("Iterator: {SeekToFirst SeekToLast Index}")
	iter.SeekToFirst()              // get first item
	log.Println(string(iter.Key())) // apple

	iter.SeekToLast()               // get last item
	log.Println(string(iter.Key())) // zero

	log.Println("Iterator: {Index}") // get index, the value is `size - 1`
	log.Println(iter.Index())        // 4

	log.Println("PutDuplicate")
	tree1.PutDuplicate([]byte("boy"), 10, func(exists *treelist.Slice) {
		exists.Value = 100 // if key is exists, set the value
	})

	//[{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}]. values in order
	log.Println("Trim TrimByIndex")             //
	tree1.Trim([]byte("word1"), []byte("zero")) //
	log.Println(Tree2String(tree1))             // [{word1:1} {word2:2} {zero:0}]

	Resotre(tree1) // Resotre Tree1

	tree1.TrimByIndex(1, 3)         //
	log.Println(Tree2String(tree1)) // [{boy:3} {word1:1} {word2:2}]

	log.Println("Traverse")
	tree1.Traverse(func(s *treelist.Slice) bool {
		log.Println(Slice2String(s))
		return true
	}) // {boy:3} {word1:1} {word2:2}

	Resotre(tree1) // Resotre Tree1

	log.Println("Remove RemoveHead RemoveTail RemoveIndex")
	tree1.Remove([]byte("word1"))
	log.Println(Slice2String(tree1.RemoveHead())) // be removed. return slice -> {apple:4}
	log.Println(Slice2String(tree1.RemoveTail())) // be removed. return slice -> {zero:0}
	log.Println(Tree2String(tree1))               //[{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}] ->  [{boy:3} {word2:2}]
	Resotre(tree1)                                // Resotre Tree1

	// [{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}] ->  [{apple:4} {word1:1} {word2:2} {zero:0}]
	log.Println(Slice2String(tree1.RemoveIndex(1))) //  be removed. return slice -> {boy:3}
	//  [{apple:4} {word1:1} {word2:2} {zero:0}] -> [{apple:4} {word2:2} {zero:0}]
	log.Println(Slice2String(tree1.RemoveIndex(1))) // be removed. return slice -> {word1:1}

	log.Println("RemoveRange RemoveRangeByIndex")
	Resotre(tree1) // Resotre Tree1
	tree1.RemoveRange([]byte("boy"), []byte("word2"))
	log.Println(Tree2String(tree1)) //[{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}] ->  [{boy:3} {word2:2}]
	Resotre(tree1)                  // Resotre Tree1
	tree1.RemoveRangeByIndex(1, 3)  // Remove by index from 1 - 3
	log.Println(Tree2String(tree1)) //[{apple:4} {boy:3} {word1:1} {word2:2} {zero:0}] ->  [{boy:3} {word2:2}]
}

// IteratorRange
func main2() {
	var TestedBytesSimlpe = [][]byte{[]byte("c1"), []byte("c4"), []byte("c6"), []byte("a1"), []byte("a3"), []byte("a5")}

	tree := treelist.New()
	for _, v := range TestedBytesSimlpe {
		tree.Put(v, v)
	}

	//	│       ┌── c6
	//	│   ┌── c4
	//	└── c1
	//		│   ┌── a5
	//		└── a3
	//			└── a1
	// [a1 a3 a5 c1 c4 c6]

	log.Println("IteratorRange: {GE2LT}")
	func() {
		var result []string
		iter := tree.IteratorRange()
		iter.GE2LT([]byte("a4"), []byte("c4")) // a4 <= key < c4
		iter.Range(func(cur *treelist.SliceIndex) bool {
			result = append(result, string(cur.Key))
			return true
		})
		log.Println(result) // "[a5 c1]"
	}()

	// [a1 a3 a5 c1 c4 c6]
	log.Println("IteratorRange: {GT2LT}")
	func() {
		var result []string
		iter := tree.IteratorRange()
		iter.GT2LT([]byte("a4"), []byte("c9")) // a4 < key < c9
		iter.Range(func(cur *treelist.SliceIndex) bool {
			result = append(result, string(cur.Key))
			return true
		})
		log.Println(result) // "[a5 c1 c4 c6]"
	}()

	// [a1 a3 a5 c1 c4 c6]
	log.Println("IteratorRange: {GE2LE}")
	func() {
		var result []string
		iter := tree.IteratorRange()
		iter.GE2LE([]byte("a0"), []byte("c9")) // a0 <= key <= c9
		iter.Range(func(cur *treelist.SliceIndex) bool {
			result = append(result, string(cur.Key))
			return true
		})
		log.Println(result) // "[a1 a3 a5 c1 c4 c6]""
	}()

	log.Println("IteratorRange: {GT2LE SetDirection}")
	func() {
		var result []string
		iter := tree.IteratorRange()
		iter.SetDirection(treelist.Reverse)    // Reverse
		iter.GT2LE([]byte("a0"), []byte("c9")) // a0 < Key <= c9
		iter.Range(func(cur *treelist.SliceIndex) bool {
			result = append(result, string(cur.Key))
			return true
		})
		log.Println(result) // [c6 c4 c1 a5 a3 a1]
	}()
}

func Resotre(tree1 *treelist.Tree) {
	tree1.Clear()
	tree1.Put([]byte("zero"), 0) // true
	tree1.Put([]byte("apple"), 4)
	tree1.Put([]byte("word1"), 1)
	tree1.Put([]byte("word2"), 2)
	tree1.Set([]byte("boy"), 4) // boy 4
	tree1.Set([]byte("boy"), 3) // boy 3
}

func Slice2String(s *treelist.Slice) string {
	return fmt.Sprintf("{%v:%v}", string(s.Key), s.Value)
}

func Tree2String(tree *treelist.Tree) []string {
	var results []string
	for _, s := range tree.Slices() {
		results = append(results, Slice2String(&s))
	}
	return results
}
