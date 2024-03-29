# avl

```go
import "github.com/474420502/structure/tree/avl"
```

## Index

 
- [type Tree](<#type-tree>)
  - [func New[T any](Compare compare.Compare[T]) *Tree[T]](<#func-new>)
  - [func (tree *Tree[T]) Clear()](<#func-treet-clear>)
  - [func (tree *Tree[T]) Get(key T) (interface{}, bool)](<#func-treet-get>)
  - [func (tree *Tree[T]) Height() int](<#func-treet-height>)
  - [func (tree *Tree[T]) Iterator() *Iterator[T]](<#func-treet-iterator>)
  - [func (tree *Tree[T]) Put(key T, value interface{}) bool](<#func-treet-put>)
  - [func (tree *Tree[T]) Remove(key T) (interface{}, bool)](<#func-treet-remove>)
  - [func (tree *Tree[T]) Set(key T, value interface{}) bool](<#func-treet-set>)
  - [func (tree *Tree[T]) Size() int64](<#func-treet-size>)
  - [func (tree *Tree[T]) String() string](<#func-treet-string>)
  - [func (tree *Tree[T]) Traverse(every func(k T, v interface{}) bool)](<#func-treet-traverse>)
  - [func (tree *Tree[T]) Values() []interface{}](<#func-treet-values>)
- [type Iterator](<#type-iterator>)
  - [func (iter *Iterator[T]) Clone() *Iterator[T]](<#func-iteratort-clone>)
  - [func (iter *Iterator[T]) Key() T](<#func-iteratort-key>)
  - [func (iter *Iterator[T]) Next()](<#func-iteratort-next>)
  - [func (iter *Iterator[T]) Prev()](<#func-iteratort-prev>)
  - [func (iter *Iterator[T]) SeekGE(key T) bool](<#func-iteratort-seekge>)
  - [func (iter *Iterator[T]) SeekGT(key T) bool](<#func-iteratort-seekgt>)
  - [func (iter *Iterator[T]) SeekLE(key T) bool](<#func-iteratort-seekle>)
  - [func (iter *Iterator[T]) SeekLT(key T) bool](<#func-iteratort-seeklt>)
  - [func (iter *Iterator[T]) SeekToFirst()](<#func-iteratort-seektofirst>)
  - [func (iter *Iterator[T]) SeekToLast()](<#func-iteratort-seektolast>)
  - [func (iter *Iterator[T]) Vaild() bool](<#func-iteratort-vaild>)
  - [func (iter *Iterator[T]) Value() interface{}](<#func-iteratort-value>)

- [examples](#examples)

## type Tree

```go
type Tree[T any] struct {
    Root       *Node[T] // Tree Root
    HeightDiff int      // Allowable height difference is 1(default). if heightdiff == 2, will faster than rbtree.

    Compare compare.Compare[T] // The compare function of the key of node
    // contains filtered or unexported fields
}
```
---



### func [New](#examples)

```go
func New[T any](Compare compare.Compare[T]) *Tree[T]
```

New create a object of tree

### func \(\*Tree\[T\]\) [Clear](#examples)

```go
func (tree *Tree[T]) Clear()
```

Clear clear all nodes

### func \(\*Tree\[T\]\) [Get](#examples)

```go
func (tree *Tree[T]) Get(key T) (interface{}, bool)
```

Get get value by key

### func \(\*Tree\[T\]\) [Height](#examples)

```go
func (tree *Tree[T]) Height() int
```

Height get the height  of tree

### func \(\*Tree\[T\]\) [Iterator](#examples)

```go
func (tree *Tree[T]) Iterator() *Iterator[T]
```

Iterator must call Seek\*\.

### func \(\*Tree\[T\]\) [Put](#examples)

```go
func (tree *Tree[T]) Put(key T, value interface{}) bool
```

Put put value into tree  by Key \. if key exists\, not cover the value and return false\.

else return true

### func \(\*Tree\[T\]\) [Remove](#examples)

```go
func (tree *Tree[T]) Remove(key T) (interface{}, bool)
```

Remove remove key and return value that be removed

### func \(\*Tree\[T\]\) [Set](#examples)

```go
func (tree *Tree[T]) Set(key T, value interface{}) bool
```

Set set value by Key. if key exists, cover the value and return true.

else return false and put value into tree

### func \(\*Tree\[T\]\) [Size](#examples)

```go
func (tree *Tree[T]) Size() int64
```

Size get the size of tree

### func \(\*Tree\[T\]\) [String](#examples)

```go
func (tree *Tree[T]) String() string
```

String show the view of tree by chars

### func \(\*Tree\[T\]\) [Traverse](#examples)

```go
func (tree *Tree[T]) Traverse(every func(k T, v interface{}) bool)
```

Traverse the traversal method defaults to LDR. from smallest to largest.

### func \(\*Tree\[T\]\) [Values](#examples)

```go
func (tree *Tree[T]) Values() []interface{}
```

Values return all nodes\.


## type Iterator

Iterator tree iterator

```go
type Iterator[T any] struct {
    // contains filtered or unexported fields
}
```

### func \(\*Iterator\[T\]\) [Clone](#examples)

```go
func (iter *Iterator[T]) Clone() *Iterator[T]
```

Clone Copy a current iterator

### func \(\*Iterator\[T\]\) [Compare](#examples)

```go
func (iter *Iterator[T]) Compare(key T) int
```

Compare iterator the  current value comare to key. <br>
if cur.key > key. return 1. <br>
if cur.key == key return 0. <br>
if cur.key < key return - 1.

### func \(\*Iterator\[T\]\) [Key](#examples)

```go
func (iter *Iterator[T]) Key() T
```

Key return the key of current iterator

### func \(\*Iterator\[T\]\) [Next](#examples)

```go
func (iter *Iterator[T]) Next()
```

Next the current iterator move to the next\. before call it must call Vaild\(\) and return true\.

### func \(\*Iterator\[T\]\) [Prev](#examples)

```go
func (iter *Iterator[T]) Prev()
```

Prev the current iterator move to the prev\. before call it must call Vaild\(\) and return true\.

### func \(\*Iterator\[T\]\) [SeekGE](#examples)

```go
func (iter *Iterator[T]) SeekGE(key T) bool
```

SeekGE seek to the key that greater than or equal to

### func \(\*Iterator\[T\]\) [SeekGT](#examples)

```go
func (iter *Iterator[T]) SeekGT(key T) bool
```

SeekGE seek to the key that greater than

### func \(\*Iterator\[T\]\) [SeekLE](#examples)

```go
func (iter *Iterator[T]) SeekLE(key T) bool
```

SeekLE seek to the key that less than or equal to

### func \(\*Iterator\[T\]\) [SeekLT](#examples)

```go
func (iter *Iterator[T]) SeekLT(key T) bool
```

SeekLT seek to the key that less than

### func \(\*Iterator\[T\]\) [SeekToFirst](#examples)

```go
func (iter *Iterator[T]) SeekToFirst()
```

SeekToFirst seek to first item

### func \(\*Iterator\[T\]\) [SeekToLast](#examples)

```go
func (iter *Iterator[T]) SeekToLast()
```

SeekToFirst seek to last item

### func \(\*Iterator\[T\]\) [Vaild](#examples)

```go
func (iter *Iterator[T]) Vaild() bool
```

Vaild if current value is not nil return true\. else return false. for use with Seek

### func \(\*Iterator\[T\]\) [Value](#examples)

```go
func (iter *Iterator[T]) Value() interface{}
```

Value return the value of current iterator
 
## Examples

```go
package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
)

func main() {

	tree := avl.New(compare.Any[int]) // create a  object

	log.Println("Put Set")
	tree.Put(0, 0) // put key value into the tree
	tree.Put(4, 4)
	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Set(3, 3)

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

	iter := tree.Iterator() // create a the object of iterator

	log.Println("SeekToFirst")
	iter.SeekToFirst() // seek to the first item
	// range first to last
	for iter.Vaild() {
		log.Println(iter.Value()) // 0 1 2 3 4
		iter.Next()
	}

	log.Println("SeekToLast")
	iter.SeekToLast() // seek to the last item
	// range last to first
	for iter.Vaild() {
		log.Println(iter.Value()) // 4 3 2 1 0
		iter.Prev()
	}

	log.Println("Remove")
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

	log.Println("SeekGE")
	iter.SeekGE(3) // seek to the iterator value >= 3
	// range 3 to 4
	for iter.Vaild() {
		log.Println(iter.Value()) // 3 4
		iter.Next()
	}

	log.Println("SeekLE")
	iter.SeekLE(2) // seek to the iterator value >= 3
	// range 1 to 0
	for iter.Vaild() {
		log.Println(iter.Value()) // 1 0
		iter.Prev()
	}

	log.Println(tree.Get(2)) // nil,false
	tree.Set(0, 1)
	log.Println(tree.Get(0)) // 1, true
	
	tree.Clear()             // clear all the data of tree
}


```