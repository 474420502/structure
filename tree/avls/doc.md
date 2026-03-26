# avls

```go
import "github.com/474420502/structure/tree/avls"
```

`avls` is the duplicate-key variant of AVL.

- `Put` always inserts a new node, even when the key already exists.
- `Set` updates one existing node for the key when present, otherwise inserts a new node.
- `Get` returns the value at the stable front of that key's duplicate run.
- `Remove` removes one node for the key at a time.
- `SeekGE` and `SeekLE` land on the first/last node of an equal-key run.
- The boolean result of `SeekGE`, `SeekGT`, `SeekLE`, and `SeekLT` reports whether the queried key exists in the tree, not whether the returned iterator position is an exact match.

## Index

 
- [type Tree](<#type-tree>)
	- [func New[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE]](<#func-new>)
	- [func (tree *Tree[KEY, VALUE]) Clear()](<#func-treekey-value-clear>)
	- [func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool)](<#func-treekey-value-get>)
	- [func (tree *Tree[KEY, VALUE]) Height() int8](<#func-treekey-value-height>)
	- [func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE]](<#func-treekey-value-iterator>)
	- [func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool](<#func-treekey-value-put>)
	- [func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool)](<#func-treekey-value-remove>)
	- [func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool](<#func-treekey-value-set>)
	- [func (tree *Tree[KEY, VALUE]) Size() uint](<#func-treekey-value-size>)
	- [func (tree *Tree[KEY, VALUE]) Traverse(every func(k KEY, v VALUE) bool)](<#func-treekey-value-traverse>)
	- [func (tree *Tree[KEY, VALUE]) Values() []VALUE](<#func-treekey-value-values>)
- [type Iterator](<#type-iterator>)
	- [func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE]](<#func-iteratorkey-value-clone>)
	- [func (iter *Iterator[KEY, VALUE]) Key() KEY](<#func-iteratorkey-value-key>)
	- [func (iter *Iterator[KEY, VALUE]) Next()](<#func-iteratorkey-value-next>)
	- [func (iter *Iterator[KEY, VALUE]) Prev()](<#func-iteratorkey-value-prev>)
	- [func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY) bool](<#func-iteratorkey-value-seekge>)
	- [func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY) bool](<#func-iteratorkey-value-seekgt>)
	- [func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY) bool](<#func-iteratorkey-value-seekle>)
	- [func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY) bool](<#func-iteratorkey-value-seeklt>)
	- [func (iter *Iterator[KEY, VALUE]) SeekToFirst()](<#func-iteratorkey-value-seektofirst>)
	- [func (iter *Iterator[KEY, VALUE]) SeekToLast()](<#func-iteratorkey-value-seektolast>)
	- [func (iter *Iterator[KEY, VALUE]) Valid() bool](<#func-iteratorkey-value-valid>)
	- [func (iter *Iterator[KEY, VALUE]) Value() VALUE](<#func-iteratorkey-value-value>)

- [examples](#examples)

## type Tree

```go
type Tree[KEY, VALUE any] struct {
	Compare compare.Compare[KEY]
	// contains filtered or unexported fields
}
```
---



### func [New](#examples)

```go
func New[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE]
```

New create a object of tree

### func \(\*Tree\[T\]\) [Clear](#examples)

```go
func (tree *Tree[KEY, VALUE]) Clear()
```

Clear clear all nodes

### func \(\*Tree\[T\]\) [Get](#examples)

```go
func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool)
```

Get returns one value for the key. When duplicates exist, it returns the value at the stable front of that duplicate run.

### func \(\*Tree\[T\]\) [Height](#examples)

```go
func (tree *Tree[KEY, VALUE]) Height() int8
```

Height get the height  of tree

### func \(\*Tree\[T\]\) [Iterator](#examples)

```go
func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE]
```

Iterator should be positioned with `Seek*`, `SeekToFirst`, or `SeekToLast` before reading.

### func \(\*Tree\[T\]\) [Put](#examples)

```go
func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool
```

Put always inserts a new node. Existing keys are not deduplicated.

### func \(\*Tree\[T\]\) [Remove](#examples)

```go
func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool)
```

Remove deletes one node for the key and returns that value.

### func \(\*Tree\[T\]\) [Set](#examples)

```go
func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool
```

Set updates one existing node for the key and returns `false`. If the key does not exist, it inserts a new node and returns `true`.

### func \(\*Tree\[T\]\) [Size](#examples)

```go
func (tree *Tree[KEY, VALUE]) Size() uint
```

Size get the size of tree

### func \(\*Tree\[T\]\) [String](#examples)

### func \(\*Tree\[KEY, VALUE\]\) [Traverse](#examples)

```go
func (tree *Tree[KEY, VALUE]) Traverse(every func(k KEY, v VALUE) bool)
```

Traverse the traversal method defaults to LDR. from smallest to largest.

### func \(\*Tree\[KEY, VALUE\]\) [Values](#examples)

```go
func (tree *Tree[KEY, VALUE]) Values() []VALUE
```

Values returns all values in in-order traversal. Duplicate keys appear multiple times.


## type Iterator

Iterator tree iterator

```go
type Iterator[KEY, VALUE any] struct {
    // contains filtered or unexported fields
}
```

### func \(\*Iterator\[KEY, VALUE\]\) [Clone](#examples)

```go
func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE]
```

Clone Copy a current iterator

### func \(\*Iterator\[KEY, VALUE\]\) [Key](#examples)

```go
func (iter *Iterator[KEY, VALUE]) Key() KEY
```

Key return the key of current iterator

### func \(\*Iterator\[KEY, VALUE\]\) [Next](#examples)

```go
func (iter *Iterator[KEY, VALUE]) Next()
```

Next moves to the next in-order node. Call it only when `Valid()` is true.

### func \(\*Iterator\[KEY, VALUE\]\) [Prev](#examples)

```go
func (iter *Iterator[KEY, VALUE]) Prev()
```

Prev moves to the previous in-order node. Call it only when `Valid()` is true.

### func \(\*Iterator\[KEY, VALUE\]\) [SeekGE](#examples)

```go
func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY) bool
```

SeekGE positions to the first node with key `>= key`. If the queried key exists, it lands on the first duplicate of that key and returns `true`.

### func \(\*Iterator\[KEY, VALUE\]\) [SeekGT](#examples)

```go
func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY) bool
```

SeekGT positions to the first node with key `> key`. The return value still reports whether the queried key exists anywhere in the tree.

### func \(\*Iterator\[KEY, VALUE\]\) [SeekLE](#examples)

```go
func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY) bool
```

SeekLE positions to the last node with key `<= key`. If the queried key exists, it lands on the last duplicate of that key and returns `true`.

### func \(\*Iterator\[KEY, VALUE\]\) [SeekLT](#examples)

```go
func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY) bool
```

SeekLT positions to the last node with key `< key`. The return value still reports whether the queried key exists anywhere in the tree.

### func \(\*Iterator\[KEY, VALUE\]\) [SeekToFirst](#examples)

```go
func (iter *Iterator[KEY, VALUE]) SeekToFirst()
```

SeekToFirst seek to first item

### func \(\*Iterator\[KEY, VALUE\]\) [SeekToLast](#examples)

```go
func (iter *Iterator[KEY, VALUE]) SeekToLast()
```

SeekToFirst seek to last item

### func \(\*Iterator\[KEY, VALUE\]\) [Valid](#examples)

```go
func (iter *Iterator[KEY, VALUE]) Valid() bool
```

Valid returns true when the iterator is positioned on a node.

### func \(\*Iterator\[KEY, VALUE\]\) [Value](#examples)

```go
func (iter *Iterator[KEY, VALUE]) Value() VALUE
```

Value return the value of current iterator
 
## Examples

```go
package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avls"
)

func main() {

	tree := avls.New[int, int](compare.AnyEx[int])

	log.Println("Put Set")
	tree.Put(0, 0)
	tree.Put(4, 40)
	tree.Put(4, 41)
	tree.Put(1, 10)
	tree.Set(4, 99) // update one existing key=4 node, does not grow size
	tree.Set(3, 30) // insert because key=3 does not exist

	log.Println("String")
	log.Println(tree.Values())
	// [0 10 30 99 41]

	log.Println("Traverse")
	tree.Traverse(func(k int, v int) bool {
		log.Println(k, v)
		return true
	})

	log.Println("Get")
	log.Println(tree.Get(4)) // 99,true
	log.Println(tree.Get(5)) // 0,false

	iter := tree.Iterator() // create a the object of iterator

	log.Println("SeekToFirst")
	iter.SeekToFirst() // seek to the first item
	// range first to last
	for iter.Valid() {
		log.Println(iter.Key(), iter.Value())
		iter.Next()
	}

	log.Println("SeekToLast")
	iter.SeekToLast() // seek to the last item
	// range last to first
	for iter.Valid() {
		log.Println(iter.Key(), iter.Value())
		iter.Prev()
	}

	log.Println("Remove")
	log.Println(tree.Remove(4)) // 99, true
	log.Println(tree.Remove(4)) // 41, true
	log.Println(tree.Remove(4)) // 0, false

	log.Println(tree.Values()) // [0 10 30]

	iter.Next() // the pos of iter is before 0. so need call Next()
	for iter.Valid() {
		log.Println(iter.Value())
		iter.Next()
	}

	log.Println("SeekGE")
	log.Println(iter.SeekGE(3), iter.Key(), iter.Value()) // true 3 30
	for iter.Valid() {
		log.Println(iter.Value())
		iter.Next()
	}

	log.Println("SeekLE")
	log.Println(iter.SeekLE(2), iter.Key(), iter.Value()) // false 1 10
	for iter.Valid() {
		log.Println(iter.Value())
		iter.Prev()
	}

	log.Println(iter.SeekGT(0), iter.Key(), iter.Value()) // true 1 10
	log.Println(iter.SeekLT(0), iter.Valid())             // true false
	
	tree.Clear()             // clear all the data of tree
}


```