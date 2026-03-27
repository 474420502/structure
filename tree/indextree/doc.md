# indextree

```go
import "github.com/474420502/structure/tree/indextree"
```

`tree/indextree` is an ordered tree that balances by subtree size and exposes rank/index-oriented operations.

Related performance documents:

- [benchmark-comparison.md](./benchmark-comparison.md)
- [benchmark-comparison.zh.md](./benchmark-comparison.zh.md)
- [rotation-analysis.md](./rotation-analysis.md)

## Index

- [type Iterator](<#type-iterator>)
	- [func (iter *Iterator[T]) Clone() *Iterator[T]](<#func-iteratort-clone>)
	- [func (iter *Iterator[T]) Index() int64](<#func-iteratort-index>)
	- [func (iter *Iterator[T]) Key() T](<#func-iteratort-key>)
	- [func (iter *Iterator[T]) Next()](<#func-iteratort-next>)
	- [func (iter *Iterator[T]) Prev()](<#func-iteratort-prev>)
	- [func (iter *Iterator[T]) SeekByIndex(index int64)](<#func-iteratort-seekbyindex>)
	- [func (iter *Iterator[T]) SeekGE(key T) bool](<#func-iteratort-seekge>)
	- [func (iter *Iterator[T]) SeekGT(key T) bool](<#func-iteratort-seekgt>)
	- [func (iter *Iterator[T]) SeekLE(key T) bool](<#func-iteratort-seekle>)
	- [func (iter *Iterator[T]) SeekLT(key T) bool](<#func-iteratort-seeklt>)
	- [func (iter *Iterator[T]) SeekToFirst()](<#func-iteratort-seektofirst>)
	- [func (iter *Iterator[T]) SeekToLast()](<#func-iteratort-seektolast>)
	- [func (iter *Iterator[T]) Valid() bool](<#func-iteratort-valid>)
	- [func (iter *Iterator[T]) Value() interface{}](<#func-iteratort-value>)
- [type Tree](<#type-tree>)
  - [func New[T any](comp compare.Compare[T]) *Tree[T]](<#func-new>)
  - [func (tree *Tree[T]) Clear()](<#func-treet-clear>)
  - [func (tree *Tree[T]) Get(key T) (interface{}, bool)](<#func-treet-get>)
  - [func (tree *Tree[T]) Index(i int64) (key T, value interface{})](<#func-treet-index>)
  - [func (tree *Tree[T]) IndexOf(key T) int64](<#func-treet-indexof>)
	- [func (tree *Tree[T]) Iterator() *Iterator[T]](<#func-treet-iterator>)
  - [func (tree *Tree[T]) Put(key T, value interface{}) bool](<#func-treet-put>)
  - [func (tree *Tree[T]) Remove(key T) interface{}](<#func-treet-remove>)
  - [func (tree *Tree[T]) RemoveIndex(index int64) interface{}](<#func-treet-removeindex>)
  - [func (tree *Tree[T]) RemoveRange(low, high T)](<#func-treet-removerange>)
  - [func (tree *Tree[T]) RemoveRangeByIndex(low, hight int64)](<#func-treet-removerangebyindex>)
  - [func (tree *Tree[T]) Set(key T, value interface{}) bool](<#func-treet-set>)
  - [func (tree *Tree[T]) Size() int64](<#func-treet-size>)
  - [func (tree *Tree[T]) Split(key T) *Tree[T]](<#func-treet-split>)
  - [func (tree *Tree[T]) SplitContain(key T) *Tree[T]](<#func-treet-splitcontain>)
  - [func (tree *Tree[T]) String() string](<#func-treet-string>)
  - [func (tree *Tree[T]) Traverse(every func(k T, v interface{}) bool)](<#func-treet-traverse>)
  - [func (tree *Tree[T]) Trim(low, high T)](<#func-treet-trim>)
  - [func (tree *Tree[T]) TrimByIndex(low, high int64)](<#func-treet-trimbyindex>)
  - [func (tree *Tree[T]) Values() []interface{}](<#func-treet-values>)
- [examples](<#examples>)


## type Iterator

Iterator is the bidirectional ordered iterator for `Tree`.

```go
type Iterator[T any] struct {
	// contains filtered or unexported fields
}
```

### func \(\*Iterator\[T\]\) [Clone](#examples)

```go
func (iter *Iterator[T]) Clone() *Iterator[T]
```

Clone copies the current iterator state.

### func \(\*Iterator\[T\]\) [Index](#examples)

```go
func (iter *Iterator[T]) Index() int64
```

Index returns the current in-order position.

### func \(\*Iterator\[T\]\) [Key](#examples)

```go
func (iter *Iterator[T]) Key() T
```

Key returns the current key.

### func \(\*Iterator\[T\]\) [Next](#examples)

```go
func (iter *Iterator[T]) Next()
```

Next moves to the next item in sorted order.

### func \(\*Iterator\[T\]\) [Prev](#examples)

```go
func (iter *Iterator[T]) Prev()
```

Prev moves to the previous item in sorted order.

### func \(\*Iterator\[T\]\) [SeekByIndex](#examples)

```go
func (iter *Iterator[T]) SeekByIndex(index int64)
```

SeekByIndex positions the iterator at the given rank.

### func \(\*Iterator\[T\]\) [SeekGE](#examples)

```go
func (iter *Iterator[T]) SeekGE(key T) bool
```

SeekGE positions at the first key greater than or equal to `key`.

### func \(\*Iterator\[T\]\) [SeekGT](#examples)

```go
func (iter *Iterator[T]) SeekGT(key T) bool
```

SeekGT positions at the first key strictly greater than `key`.

### func \(\*Iterator\[T\]\) [SeekLE](#examples)

```go
func (iter *Iterator[T]) SeekLE(key T) bool
```

SeekLE positions at the last key less than or equal to `key`.

### func \(\*Iterator\[T\]\) [SeekLT](#examples)

```go
func (iter *Iterator[T]) SeekLT(key T) bool
```

SeekLT positions at the last key strictly less than `key`.

### func \(\*Iterator\[T\]\) [SeekToFirst](#examples)

```go
func (iter *Iterator[T]) SeekToFirst()
```

SeekToFirst positions at the smallest key.

### func \(\*Iterator\[T\]\) [SeekToLast](#examples)

```go
func (iter *Iterator[T]) SeekToLast()
```

SeekToLast positions at the largest key.

### func \(\*Iterator\[T\]\) [Valid](#examples)

```go
func (iter *Iterator[T]) Valid() bool
```

Valid reports whether the iterator is positioned on a node.

### func \(\*Iterator\[T\]\) [Value](#examples)

```go
func (iter *Iterator[T]) Value() interface{}
```

Value returns the current value.


## type [Tree](#examples)

Tree the struct of tree

```go
type Tree[T any] struct {
    // contains filtered or unexported fields
}
```

### func [New](#examples)

```go
func New[T any](comp compare.Compare[T]) *Tree[T]
```

New create a object of tree

### func \(\*Tree\[T\]\) [Clear](#examples)

```go
func (tree *Tree[T]) Clear()
```

Clear clear all node\.

### func \(\*Tree\[T\]\) [Get](#examples)

```go
func (tree *Tree[T]) Get(key T) (interface{}, bool)
```

Get get value by key

### func \(\*Tree\[T\]\) [Index](#examples)

```go
func (tree *Tree[T]) Index(i int64) (key T, value interface{})
```

Index Indexing Ordered Data\. like TopN

### func \(\*Tree\[T\]\) [IndexOf](#examples)

```go
func (tree *Tree[T]) IndexOf(key T) int64
```

Index Indexing Ordered Data\. like TopN

### func \(\*Tree\[T\]\) [Iterator](#examples)

```go
func (tree *Tree[T]) Iterator() *Iterator[T]
```

Iterator creates a bidirectional ordered iterator.

### func \(\*Tree\[T\]\) [Put](#examples)

```go
func (tree *Tree[T]) Put(key T, value interface{}) bool
```

Put put value into tree  by Key \. if key exists\,not cover the value and return false\. else return true

### func \(\*Tree\[T\]\) [Remove](#examples)

```go
func (tree *Tree[T]) Remove(key T) interface{}
```

Remove remove key value and return value that be removed

### func \(\*Tree\[T\]\) [RemoveIndex](#examples)

```go
func (tree *Tree[T]) RemoveIndex(index int64) interface{}
```

RemoveIndex remove key value by index and return value that be removed

### func \(\*Tree\[T\]\) [RemoveRange](#examples)

```go
func (tree *Tree[T]) RemoveRange(low, high T)
```

RemoveRange remove keys values by range\. \[low\, high\]

### func \(\*Tree\[T\]\) [RemoveRangeByIndex](#examples)

```go
func (tree *Tree[T]) RemoveRangeByIndex(low, hight int64)
```

RemoveRangeByIndex 1\.remove range \[low:hight\] 2.low and hight that the range must contain a value that exists. eg: [low: hight+1] [low-1: hight].  [low-1: hight+1]. error: [low-1:low-2] or [hight+1:hight+2]

### func \(\*Tree\[T\]\) [Set](#examples)

```go
func (tree *Tree[T]) Set(key T, value interface{}) bool
```

Set set value by Key\. if key exists\, cover the value and return true\. else return false and put value into tree

### func \(\*Tree\[T\]\) [Size](#examples)

```go
func (tree *Tree[T]) Size() int64
```

Size get the size of tree

### func \(\*Tree\[T\]\) [Split](#examples)

```go
func (tree *Tree[T]) Split(key T) *Tree[T]
```

Split Contain Split  Original tree not contain Key\. return  the splited tree

### func \(\*Tree\[T\]\) [SplitContain](#examples)

```go
func (tree *Tree[T]) SplitContain(key T) *Tree[T]
```

SplitContain  Original tree contain Key\. return  the splited tree

### func \(\*Tree\[T\]\) [String](#examples)

```go
func (tree *Tree[T]) String() string
```

String show the view of tree by chars

### func \(\*Tree\[T\]\) [Traverse](#examples)

```go
func (tree *Tree[T]) Traverse(every func(k T, v interface{}) bool)
```

Traverse the traversal method defaults to LDR\. from smallest to largest\.

### func \(\*Tree\[T\]\) [Trim](#examples)

```go
func (tree *Tree[T]) Trim(low, high T)
```

Trim retain the value of the range \. \[low high\]

### func \(\*Tree\[T\]\) [TrimByIndex](#examples)

```go
func (tree *Tree[T]) TrimByIndex(low, high int64)
```

TrimByIndex retain the value of the index range \. \[low high\]

### func \(\*Tree\[T\]\) [Values](#examples)

```go
func (tree *Tree[T]) Values() []interface{}
```

Values return all values. in order

## examples

```go
package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/indextree"
)

func main() {

	tree := indextree.New(compare.Any[int]) // create a  object

	log.Println("Put Set")
	tree.Put(0, 0) // put key value into the tree
	tree.Put(4, 4)
	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Set(3, 3) // Set value

	log.Println("Values")
	log.Println(tree.Values())

	log.Println("Size")
	tree.Size() // 5

	log.Println("Index IndexOf")
	log.Println(tree.Index(0))   // 0 0
	log.Println(tree.Index(4))   // 4 4
	log.Println(tree.IndexOf(2)) // 2 like RankOf

	log.Println("String")
	log.Println(tree.String())
	//
	// │       ┌── 4
	// │   ┌── 3
	// │   │   └── 2
	// └── 1
	//     └── 0

	log.Println("Traverse")
	tree.Traverse(func(k int, v interface{}) bool {
		log.Println(k, v) // 0 0 1 1 2 2 3 3 4 4
		return true
	})

	log.Println("Get")
	log.Println(tree.Get(4)) // 4,true
	log.Println(tree.Get(5)) // nil,false

	log.Println("Remove")
	log.Println(tree.Remove(2)) // 2, true
	log.Println(tree.Remove(2)) // <nil>, false
	tree.Put(5, 5)
	tree.Put(6, 6)
	log.Println(tree.String())
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	└── 4
	// 		│   ┌── 3
	// 		└── 1
	// 			└── 0
	log.Println("RemoveIndex")
	tree.RemoveIndex(0)        // Remove head.
	log.Println(tree.String()) // 1 3 4 5 6
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	└── 4
	// 		│   ┌── 3
	// 		└── 1
	log.Println("RemoveRange")
	tree.RemoveRange(2, 5)     // remove 2-5 key
	log.Println(tree.String()) // 1 6
	// 	└── 6
	// 		└── 1

	tree.Put(15, 15)
	tree.Put(16, 16)
	log.Println(tree.String())
	//	│       ┌── 16
	//	│   ┌── 15
	//	└── 6
	//	    └── 1
	log.Println("RemoveRangeByIndex")
	tree.RemoveRangeByIndex(1, 2)
	log.Println(tree.String()) // 1 16
	// 	└── 16
	// 		└── 1

	tree.Clear() // clear all the data of tree
}
```


```go
package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/indextree"
)

func main() {
	tree := indextree.New(compare.Any[int]) // create a  object

	for i := 0; i < 7; i++ {
		tree.Put(i, i)
	}

	log.Println(tree.String())
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3
	// 		│   ┌── 2
	// 		└── 1
	// 			└── 0

	log.Println("Split")
	tree2 := tree.Split(3)
	log.Println(tree.String())
	// 	│   ┌── 2
	// 	└── 1
	// 		└── 0
	log.Println(tree2.String())
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3

	tree.Clear()
	for i := 0; i < 7; i++ {
		tree.Put(i, i)
	}

	log.Println("SplitContain")
	tree2 = tree.SplitContain(3)
	log.Println(tree.String())
	// 	└── 3
	// 		│   ┌── 2
	// 		└── 1
	// 		    └── 0
	log.Println(tree2.String())
	// 	│   ┌── 6
	// 	└── 5
	// 	    └── 4

	tree.Clear()
	for i := 0; i < 7; i++ {
		tree.Put(i, i)
	}
	log.Println(tree.String())
	// 	│       ┌── 6
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3
	// 		│   ┌── 2
	// 		└── 1
	// 			└── 0

	log.Println("Trim")
	tree.Trim(2, 5) // keep the values that range with 2-5
	log.Println(tree.String())
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3
	// 	    └── 2

	tree.Clear()
	for i := 0; i < 7; i++ {
		tree.Put(i, i)
	}
	log.Println("TrimByIndex")
	tree.TrimByIndex(2, 5) // keep the values that index range with 2-5
	log.Println(tree.String())
	// 	│   ┌── 5
	// 	│   │   └── 4
	// 	└── 3
	// 	    └── 2
}
```