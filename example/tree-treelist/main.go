package main

import (
	"fmt"
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/treelist"
)

func main() {

	// New a object of tree
	tree1 := treelist.New(compare.Any[int])

	log.Println("Put Set")
	tree1.Put(0, 0) // true
	tree1.Put(4, 4)
	tree1.Put(1, 1)
	// tree1.Put(2, 2)
	tree1.Set(4, 4) //   4
	tree1.Set(3, 3) //   3 insert
	tree1.Set(7, 7) //   7

	log.Println("Slices")
	var results []string
	for _, slice := range tree1.Slices() {
		results = append(results, Slice2String(&slice))
	}
	log.Println(results) // [{0:0} {1:1} {3:3} {4:4} {7:7}]. values in order

	log.Println("Get")
	tree1.Get(1)   // 1, true
	tree1.Get(100) // nil, false

	log.Println("Head Tail")
	log.Println(tree1.Head()) // {0:0}
	log.Println(tree1.Tail()) // {7:7}

	log.Println("Index IndexOf Size")
	log.Println(tree1.Index(0))   // {0:0}
	log.Println(tree1.IndexOf(1)) // 1
	log.Println(tree1.Index(4))   // {7:7}

	log.Println("Intersection UnionSets") //
	tree2 := treelist.New(compare.Any[int])
	// [1 2 5]
	tree2.Set(1, 1)
	tree2.Set(3, 3)
	tree2.Set(5, 5)

	tree3 := tree1.Intersection(tree2)            // Intersection
	log.Println(Tree2String(tree3), tree3.Size()) // [{1:1} {3:3}] 2

	tree3 = tree1.UnionSets(tree2)                // UnionSets
	log.Println(Tree2String(tree3), tree3.Size()) // [{0:0} {1:1} {3:3} {4:4} {5:5} {7:7}] 6

	// [{0:0} {1:1} {3:3} {4:4} {7:7}]
	log.Println("Iterator: {Valid Next SeekGE}")
	iter := tree1.Iterator()
	log.Println(iter.SeekGE(2))       // return false. key >= 2 similar to rocksdb pebble leveldb skiplist
	for ; iter.Valid(); iter.Next() { // Vaiid Next
		log.Println(iter.Key()) // log: 3 4 7
		// you can limit by yourself
	}

	// [{0:0} {1:1} {3:3} {4:4} {7:7}]
	log.Println("Iterator: {Valid Next SeekGT}")
	log.Println(iter.SeekGT(2))       // return false.  key > 2
	for ; iter.Valid(); iter.Next() { // Vaiid Next
		log.Println(iter.Key()) // log: 3 4 7
		// you can limit by yourself
	}

	// [{0:0} {1:1} {3:3} {4:4} {7:7}]
	log.Println("Iterator: {Valid Prev SeekLE}")
	log.Println(iter.SeekLE(3))       // return true . key  <= 3
	for ; iter.Valid(); iter.Prev() { // Vaiid Next
		log.Println(iter.Key()) // log: 3 1 0
		// you can limit by yourself
	}

	// [{0:0} {1:1} {3:3} {4:4} {7:7}]
	log.Println("Iterator: {Valid Prev SeekLT}")
	if iter.SeekLT(3) { // return true. key < 3
		for ; iter.Valid(); iter.Prev() { // Vaiid Next
			log.Println(iter.Key()) // log: 1 0
			// you can limit by yourself
		}
	}

	log.Println("Iterator: {SeekToFirst SeekToLast Index}")
	iter.SeekToFirst()      // get first item
	log.Println(iter.Key()) // 0

	iter.SeekToLast()       // get last item
	log.Println(iter.Key()) // 7

	log.Println("Iterator: {Index}") // get index, the value is `size - 1`
	log.Println(iter.Index())        // 4

	log.Println("PutDuplicate")
	tree1.PutDuplicate(10, 10, func(exists *treelist.Slice[int]) {
		exists.Value = 100 // if key is exists, set the value
	})
	// [{0:0} {1:1} {3:3} {4:4} {7:7} {10:10}]
	log.Println(Tree2String(tree1))
	tree1.Remove(10)

	// [{0:0} {1:1} {3:3} {4:4} {7:7}]. values in order
	log.Println("Trim TrimByIndex") //
	tree1.Trim(1, 4)                //
	log.Println(Tree2String(tree1)) // [{1:1} {3:3} {4:4}]

	Resotre(tree1) // Resotre Tree1

	tree1.TrimByIndex(1, 3)         //
	log.Println(Tree2String(tree1)) // [{1:1} {3:3} {4:4}]

	log.Println("Traverse")
	tree1.Traverse(func(s *treelist.Slice[int]) bool {
		log.Println(Slice2String(s))
		return true
	}) // {1:1} {3:3} {4:4}

	Resotre(tree1) // Resotre Tree1

	// [{0:0} {1:1} {3:3} {4:4} {7:7}]
	log.Println("Remove RemoveHead RemoveTail RemoveIndex")
	log.Println(tree1.Remove(3))                  // if not exists, return nil. {3:3}
	log.Println(Slice2String(tree1.RemoveHead())) // be removed. return slice -> {0:0}
	log.Println(Slice2String(tree1.RemoveTail())) // be removed. return slice -> {7:7}
	log.Println(Tree2String(tree1))               // [{0:0} {1:1} {3:3} {4:4} {7:7}] ->  [{1:1} {4:4}]
	Resotre(tree1)                                // Resotre Tree1

	//  [{0:0} {1:1} {3:3} {4:4} {7:7}] ->   [{0:0} {3:3} {4:4} {7:7}]
	log.Println(Slice2String(tree1.RemoveIndex(1))) //  be removed. return slice -> {1:1}
	//  [{0:0} {3:3} {4:4} {7:7}] -> [{0:0} {4:4} {7:7}]
	log.Println(Slice2String(tree1.RemoveIndex(1))) // be removed. return slice -> {3:3}

	log.Println("RemoveRange RemoveRangeByIndex")
	Resotre(tree1) // Resotre Tree1
	tree1.RemoveRange(2, 4)
	log.Println(Tree2String(tree1)) // [{0:0} {1:1} {3:3} {4:4} {7:7}] -> [{0:0} {1:1} {7:7}]
	Resotre(tree1)                  // Resotre Tree1
	tree1.RemoveRangeByIndex(1, 3)  // Remove by index from 1 - 3
	log.Println(Tree2String(tree1)) // [{0:0} {1:1} {3:3} {4:4} {7:7}] ->  [{0:0} {7:7}]
}

// IteratorRange
func main2() {
	var TestedBytesSimlpe = []int{15, 4, 11, 6, 13, 1}

	tree := treelist.New(compare.Any[int])
	for _, v := range TestedBytesSimlpe {
		tree.Put(v, v)
	}
	log.Println(Tree2String(tree))
	// [{1:1} {4:4} {6:6} {11:11} {13:13} {15:15}]

	log.Println("IteratorRange: {GE2LT}")
	func() {
		var result []int
		iter := tree.IteratorRange()
		iter.GE2LT(6, 13) // 6 <= key < 13
		iter.Range(func(cur *treelist.SliceIndex[int]) bool {
			result = append(result, cur.Key)
			return true
		})
		log.Println(result) // "[6 11]"
	}()

	// [{1:1} {4:4} {6:6} {11:11} {13:13} {15:15}]
	log.Println("IteratorRange: {GT2LT}")
	func() {
		var result []int
		iter := tree.IteratorRange()
		iter.GT2LT(6, 13) // 6 < key < 13
		iter.Range(func(cur *treelist.SliceIndex[int]) bool {
			result = append(result, cur.Key)
			return true
		})
		log.Println(result) // "[11]"
	}()

	// [{1:1} {4:4} {6:6} {11:11} {13:13} {15:15}]
	log.Println("IteratorRange: {GE2LE}")
	func() {
		var result []int
		iter := tree.IteratorRange()
		iter.GE2LE(6, 13) // 6 <= key <= 13
		iter.Range(func(cur *treelist.SliceIndex[int]) bool {
			result = append(result, cur.Key)
			return true
		})
		log.Println(result) // "[6 11 13]"
	}()

	// [{1:1} {4:4} {6:6} {11:11} {13:13} {15:15}]
	log.Println("IteratorRange: {GT2LE SetDirection}")
	func() {
		var result []int
		iter := tree.IteratorRange()
		iter.SetDirection(treelist.Reverse) // Reverse
		iter.GT2LE(6, 13)                   // 6 < Key <= 13
		iter.Range(func(cur *treelist.SliceIndex[int]) bool {
			result = append(result, cur.Key)
			return true
		})
		log.Println(result) // [13 11]
	}()

}

func Resotre[T int](tree1 *treelist.Tree[T]) {
	tree1.Clear()
	tree1.Put(0, 0) // true
	tree1.Put(4, 4)
	tree1.Put(1, 1)
	// tree1.Put(2, 2)
	tree1.Set(3, 3) //   3 insert
	tree1.Set(7, 7) //   7
}

func Slice2String[T any](s *treelist.Slice[T]) string {
	return fmt.Sprintf("{%v:%v}", s.Key, s.Value)
}

func Tree2String[T any](tree *treelist.Tree[T]) []string {
	var results []string
	for _, s := range tree.Slices() {
		results = append(results, Slice2String(&s))
	}
	return results
}
