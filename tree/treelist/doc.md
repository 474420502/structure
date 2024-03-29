<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# treelist

```go
import "github.com/474420502/structure/tree/treelist"
```

## Index

- [type Iterator](<#type-iterator>)
  - [func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE]](<#func-iteratorkey-value-clone>)
  - [func (iter *Iterator[KEY, VALUE]) Compare(key KEY) int](<#func-iteratorkey-value-compare>)
  - [func (iter *Iterator[KEY, VALUE]) Index() int64](<#func-iteratorkey-value-index>)
  - [func (iter *Iterator[KEY, VALUE]) Key() KEY](<#func-iteratorkey-value-key>)
  - [func (iter *Iterator[KEY, VALUE]) Next()](<#func-iteratorkey-value-next>)
  - [func (iter *Iterator[KEY, VALUE]) Prev()](<#func-iteratorkey-value-prev>)
  - [func (iter *Iterator[KEY, VALUE]) SeekByIndex(index int64)](<#func-iteratorkey-value-seekbyindex>)
  - [func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY) bool](<#func-iteratorkey-value-seekge>)
  - [func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY) bool](<#func-iteratorkey-value-seekgt>)
  - [func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY) bool](<#func-iteratorkey-value-seekle>)
  - [func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY) bool](<#func-iteratorkey-value-seeklt>)
  - [func (iter *Iterator[KEY, VALUE]) SeekToFirst()](<#func-iteratorkey-value-seektofirst>)
  - [func (iter *Iterator[KEY, VALUE]) SeekToLast()](<#func-iteratorkey-value-seektolast>)
  - [func (iter *Iterator[KEY, VALUE]) Slice() *Slice[KEY, VALUE]](<#func-iteratorkey-value-slice>)
  - [func (iter *Iterator[KEY, VALUE]) Valid() bool](<#func-iteratorkey-value-valid>)
  - [func (iter *Iterator[KEY, VALUE]) Value() interface{}](<#func-iteratorkey-value-value>)
- [type IteratorRange](<#type-iteratorrange>)
  - [func (ir *IteratorRange[KEY, VALUE]) Direction() RangeDirection](<#func-iteratorrangekey-value-direction>)
  - [func (ir *IteratorRange[KEY, VALUE]) GE2LE(start, end KEY)](<#func-iteratorrangekey-value-ge2le>)
  - [func (ir *IteratorRange[KEY, VALUE]) GE2LT(start, end KEY)](<#func-iteratorrangekey-value-ge2lt>)
  - [func (ir *IteratorRange[KEY, VALUE]) GT2LE(start, end KEY)](<#func-iteratorrangekey-value-gt2le>)
  - [func (ir *IteratorRange[KEY, VALUE]) GT2LT(start, end KEY)](<#func-iteratorrangekey-value-gt2lt>)
  - [func (ir *IteratorRange[KEY, VALUE]) Range(do func(cur *SliceIndex[KEY, VALUE]) bool)](<#func-iteratorrangekey-value-range>)
  - [func (ir *IteratorRange[KEY, VALUE]) SetDirection(dir RangeDirection)](<#func-iteratorrangekey-value-setdirection>)
  - [func (ir *IteratorRange[KEY, VALUE]) Size() int64](<#func-iteratorrangekey-value-size>)
- [type RangeDirection](<#type-rangedirection>)
- [type Slice](<#type-slice>)
  - [func (s *Slice[KEY, VALUE]) String() string](<#func-slicekey-value-string>)
- [type SliceIndex](<#type-sliceindex>)
- [type Tree](<#type-tree>)
  - [func New[KEY any, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]](<#func-new>)
  - [func (tree *Tree[KEY, VALUE]) Clear()](<#func-treekey-value-clear>)
  - [func (tree *Tree[KEY, VALUE]) DifferenceSets(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]](<#func-treekey-value-differencesets>)
  - [func (tree *Tree[KEY, VALUE]) Get(key KEY) (interface{}, bool)](<#func-treekey-value-get>)
  - [func (tree *Tree[KEY, VALUE]) Head() *Slice[KEY, VALUE]](<#func-treekey-value-head>)
  - [func (tree *Tree[KEY, VALUE]) Index(i int64) *Slice[KEY, VALUE]](<#func-treekey-value-index>)
  - [func (tree *Tree[KEY, VALUE]) IndexOf(key KEY) int64](<#func-treekey-value-indexof>)
  - [func (tree *Tree[KEY, VALUE]) Intersection(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]](<#func-treekey-value-intersection>)
  - [func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE]](<#func-treekey-value-iterator>)
  - [func (tree *Tree[KEY, VALUE]) IteratorRange() *IteratorRange[KEY, VALUE]](<#func-treekey-value-iteratorrange>)
  - [func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool](<#func-treekey-value-put>)
  - [func (tree *Tree[KEY, VALUE]) PutDuplicate(key KEY, value VALUE, do func(exists *Slice[KEY, VALUE])) bool](<#func-treekey-value-putduplicate>)
  - [func (tree *Tree[KEY, VALUE]) Remove(key KEY) *Slice[KEY, VALUE]](<#func-treekey-value-remove>)
  - [func (tree *Tree[KEY, VALUE]) RemoveHead() *Slice[KEY, VALUE]](<#func-treekey-value-removehead>)
  - [func (tree *Tree[KEY, VALUE]) RemoveIndex(index int64) *Slice[KEY, VALUE]](<#func-treekey-value-removeindex>)
  - [func (tree *Tree[KEY, VALUE]) RemoveRange(low, hight KEY) bool](<#func-treekey-value-removerange>)
  - [func (tree *Tree[KEY, VALUE]) RemoveRangeByIndex(low, hight int64)](<#func-treekey-value-removerangebyindex>)
  - [func (tree *Tree[KEY, VALUE]) RemoveTail() *Slice[KEY, VALUE]](<#func-treekey-value-removetail>)
  - [func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool](<#func-treekey-value-set>)
  - [func (tree *Tree[KEY, VALUE]) Size() int64](<#func-treekey-value-size>)
  - [func (tree *Tree[KEY, VALUE]) Slices() []Slice[KEY, VALUE]](<#func-treekey-value-slices>)
  - [func (tree *Tree[KEY, VALUE]) Tail() *Slice[KEY, VALUE]](<#func-treekey-value-tail>)
  - [func (tree *Tree[KEY, VALUE]) Traverse(every func(s *Slice[KEY, VALUE]) bool)](<#func-treekey-value-traverse>)
  - [func (tree *Tree[KEY, VALUE]) Trim(low, hight KEY)](<#func-treekey-value-trim>)
  - [func (tree *Tree[KEY, VALUE]) TrimByIndex(low, hight int64)](<#func-treekey-value-trimbyindex>)
  - [func (tree *Tree[KEY, VALUE]) UnionSets(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]](<#func-treekey-value-unionsets>)


## type [Iterator](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L9-L14>)

Iterator the iterator of treelist

```go
type Iterator[KEY any, VALUE any] struct {
    // contains filtered or unexported fields
}
```

### func \(\*Iterator\[KEY, VALUE\]\) [Clone](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L152>)

```go
func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE]
```

Clone copy a iterator. eg: record iterator position

### func \(\*Iterator\[KEY, VALUE\]\) [Compare](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L119>)

```go
func (iter *Iterator[KEY, VALUE]) Compare(key KEY) int
```

Compare iterator the  current value comare to key. if cur.key \> key. return 1. if cur.key == key return 0. if cur.key \< key return \- 1.

### func \(\*Iterator\[KEY, VALUE\]\) [Index](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L136>)

```go
func (iter *Iterator[KEY, VALUE]) Index() int64
```

Index return the Index of the current iterator. Ordered position equivalent to the Index of an Priority Queue\(Array\)

### func \(\*Iterator\[KEY, VALUE\]\) [Key](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L141>)

```go
func (iter *Iterator[KEY, VALUE]) Key() KEY
```

Key return the key of current

### func \(\*Iterator\[KEY, VALUE\]\) [Next](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L124>)

```go
func (iter *Iterator[KEY, VALUE]) Next()
```

Next Next the current iterator move to the next. before call it must call Vaild\(\) and return true.

### func \(\*Iterator\[KEY, VALUE\]\) [Prev](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L108>)

```go
func (iter *Iterator[KEY, VALUE]) Prev()
```

Prev  the current iterator move to the prev. before call it must call Vaild\(\) and return true.

### func \(\*Iterator\[KEY, VALUE\]\) [SeekByIndex](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L48>)

```go
func (iter *Iterator[KEY, VALUE]) SeekByIndex(index int64)
```

SeekByIndex seek to  the key by index. like index of array. index is ordered

### func \(\*Iterator\[KEY, VALUE\]\) [SeekGE](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L19>)

```go
func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY) bool
```

SeekGE seek to Greater Than or Equal the key. if equal is not exists, take the great key

### func \(\*Iterator\[KEY, VALUE\]\) [SeekGT](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L34>)

```go
func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY) bool
```

SeekGT seek to Greater Than the key. take the great key

### func \(\*Iterator\[KEY, VALUE\]\) [SeekLE](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L57>)

```go
func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY) bool
```

SeekLE seek to  less than or equal the key. if equal is not exists, take the less key

### func \(\*Iterator\[KEY, VALUE\]\) [SeekLT](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L71>)

```go
func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY) bool
```

SeekLT seek to  less than  the key.

### func \(\*Iterator\[KEY, VALUE\]\) [SeekToFirst](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L84>)

```go
func (iter *Iterator[KEY, VALUE]) SeekToFirst()
```

SeekToFirst to the first item of the ordered sequence

### func \(\*Iterator\[KEY, VALUE\]\) [SeekToLast](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L94>)

```go
func (iter *Iterator[KEY, VALUE]) SeekToLast()
```

SeekToLast to the last item of the ordered sequence

### func \(\*Iterator\[KEY, VALUE\]\) [Slice](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L131>)

```go
func (iter *Iterator[KEY, VALUE]) Slice() *Slice[KEY, VALUE]
```

Slice return the KeyValue of current

### func \(\*Iterator\[KEY, VALUE\]\) [Valid](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L103>)

```go
func (iter *Iterator[KEY, VALUE]) Valid() bool
```

Valid  if current value is not nil return true. else return false. for use with Seek

### func \(\*Iterator\[KEY, VALUE\]\) [Value](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator.go#L146>)

```go
func (iter *Iterator[KEY, VALUE]) Value() interface{}
```

Value return the value of current

## type [IteratorRange](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L13-L17>)

IteratorRange the iterator for easy to range the data

```go
type IteratorRange[KEY any, VALUE any] struct {
    // contains filtered or unexported fields
}
```

### func \(\*IteratorRange\[KEY, VALUE\]\) [Direction](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L69>)

```go
func (ir *IteratorRange[KEY, VALUE]) Direction() RangeDirection
```

SetDirection set iterator range direction

### func \(\*IteratorRange\[KEY, VALUE\]\) [GE2LE](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L85>)

```go
func (ir *IteratorRange[KEY, VALUE]) GE2LE(start, end KEY)
```

GE2LE \[s,e\] start with GE, end with LE. \(like Seek\*\*\)

### func \(\*IteratorRange\[KEY, VALUE\]\) [GE2LT](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L137>)

```go
func (ir *IteratorRange[KEY, VALUE]) GE2LT(start, end KEY)
```

GE2LT \[s,e\) start with GE, end with LT. \(like Seek\*\*\)

### func \(\*IteratorRange\[KEY, VALUE\]\) [GT2LE](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L111>)

```go
func (ir *IteratorRange[KEY, VALUE]) GT2LE(start, end KEY)
```

GE2LE \(s,e\] start with GT, end with LE. \(like Seek\*\*\)

### func \(\*IteratorRange\[KEY, VALUE\]\) [GT2LT](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L163>)

```go
func (ir *IteratorRange[KEY, VALUE]) GT2LT(start, end KEY)
```

GE2LT \(s,e\) start with GT, end with LT. \(like Seek\*\*\)

### func \(\*IteratorRange\[KEY, VALUE\]\) [Range](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L25>)

```go
func (ir *IteratorRange[KEY, VALUE]) Range(do func(cur *SliceIndex[KEY, VALUE]) bool)
```

SetDirection set iterator range direction. default Forward\(start to end\)

### func \(\*IteratorRange\[KEY, VALUE\]\) [SetDirection](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L64>)

```go
func (ir *IteratorRange[KEY, VALUE]) SetDirection(dir RangeDirection)
```

SetDirection set iterator range direction. default Forward\(start to end\)

### func \(\*IteratorRange\[KEY, VALUE\]\) [Size](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L74>)

```go
func (ir *IteratorRange[KEY, VALUE]) Size() int64
```

Size get range size

## type [RangeDirection](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L3>)

```go
type RangeDirection int
```

```go
const (
    // Forward start to end
    Forward RangeDirection = 0
    // Reverse end KEYo start
    Reverse RangeDirection = 1
)
```

## type [Slice](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L10-L13>)

Slice the KeyValue

```go
type Slice[KEY any, VALUE any] struct {
    Key   KEY
    Value VALUE
}
```

### func \(\*Slice\[KEY, VALUE\]\) [String](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L16>)

```go
func (s *Slice[KEY, VALUE]) String() string
```

String show the string of keyvalue

## type [SliceIndex](<https://github.com/474420502/structure/blob/master/tree/treelist/iterator_range.go#L19-L22>)

```go
type SliceIndex[KEY any, VALUE any] struct {
    Index int64
    // contains filtered or unexported fields
}
```

## type [Tree](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L35-L40>)

Tree the struct of treelist

```go
type Tree[KEY any, VALUE any] struct {
    // contains filtered or unexported fields
}
```

### func [New](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L43>)

```go
func New[KEY any, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]
```

New create a object of tree

### func \(\*Tree\[KEY, VALUE\]\) [Clear](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L742>)

```go
func (tree *Tree[KEY, VALUE]) Clear()
```

Clear. Reset the treelist.

### func \(\*Tree\[KEY, VALUE\]\) [DifferenceSets](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L1022>)

```go
func (tree *Tree[KEY, VALUE]) DifferenceSets(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]
```

DifferenceSets The set of elements after subtracting B from A

### func \(\*Tree\[KEY, VALUE\]\) [Get](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L72>)

```go
func (tree *Tree[KEY, VALUE]) Get(key KEY) (interface{}, bool)
```

Get Get Value from key.

### func \(\*Tree\[KEY, VALUE\]\) [Head](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L440>)

```go
func (tree *Tree[KEY, VALUE]) Head() *Slice[KEY, VALUE]
```

Head returns the head of the ordered data of tree

### func \(\*Tree\[KEY, VALUE\]\) [Index](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L345>)

```go
func (tree *Tree[KEY, VALUE]) Index(i int64) *Slice[KEY, VALUE]
```

Index return the slice by index.

like the index of array\(order\)

### func \(\*Tree\[KEY, VALUE\]\) [IndexOf](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L351>)

```go
func (tree *Tree[KEY, VALUE]) IndexOf(key KEY) int64
```

IndexOf Get the Index of key in the Treelist\(Order\)

### func \(\*Tree\[KEY, VALUE\]\) [Intersection](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L940>)

```go
func (tree *Tree[KEY, VALUE]) Intersection(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]
```

Intersection  tree intersection with other. \[1 2 3\] \[2 3 4\] \-\> \[2 3\].

### func \(\*Tree\[KEY, VALUE\]\) [Iterator](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L52>)

```go
func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE]
```

Iterator Return the Iterator of tree. like list or skiplist

### func \(\*Tree\[KEY, VALUE\]\) [IteratorRange](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L59>)

```go
func (tree *Tree[KEY, VALUE]) IteratorRange() *IteratorRange[KEY, VALUE]
```

IteratorRange Return the Iterator of tree. like list or skiplist.

the struct can set range.

### func \(\*Tree\[KEY, VALUE\]\) [Put](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L257>)

```go
func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool
```

Put Insert the key In treelist, if key exists, ignore

### func \(\*Tree\[KEY, VALUE\]\) [PutDuplicate](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L81>)

```go
func (tree *Tree[KEY, VALUE]) PutDuplicate(key KEY, value VALUE, do func(exists *Slice[KEY, VALUE])) bool
```

PutDuplicate put, when key duplicate with call do. don,t change the key of \`exists\`, will break the tree of blance if duplicate, will return true.

### func \(\*Tree\[KEY, VALUE\]\) [Remove](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L424>)

```go
func (tree *Tree[KEY, VALUE]) Remove(key KEY) *Slice[KEY, VALUE]
```

Remove remove key and return value that be removed. if not exists, return nil

### func \(\*Tree\[KEY, VALUE\]\) [RemoveHead](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L449>)

```go
func (tree *Tree[KEY, VALUE]) RemoveHead() *Slice[KEY, VALUE]
```

RemoveHead remove the head of the ordered data of tree. similar to the pop function of heap

### func \(\*Tree\[KEY, VALUE\]\) [RemoveIndex](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L432>)

```go
func (tree *Tree[KEY, VALUE]) RemoveIndex(index int64) *Slice[KEY, VALUE]
```

RemoveIndex remove key value by index and return value that be removed

### func \(\*Tree\[KEY, VALUE\]\) [RemoveRange](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L474>)

```go
func (tree *Tree[KEY, VALUE]) RemoveRange(low, hight KEY) bool
```

RemoveRange remove keys values by range. \[low, high\]

### func \(\*Tree\[KEY, VALUE\]\) [RemoveRangeByIndex](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L596>)

```go
func (tree *Tree[KEY, VALUE]) RemoveRangeByIndex(low, hight int64)
```

RemoveRangeByIndex 1.range \[low:hight\] 2.low hight 必须包含存在的值.\[low: hight\+1\] \[low\-1: hight\].  \[low\-1: hight\+1\]. error: \[low\-1:low\-2\] or \[hight\+1:hight\+2\]

### func \(\*Tree\[KEY, VALUE\]\) [RemoveTail](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L466>)

```go
func (tree *Tree[KEY, VALUE]) RemoveTail() *Slice[KEY, VALUE]
```

RemoveTail remove the tail of the ordered data of tree.

### func \(\*Tree\[KEY, VALUE\]\) [Set](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L168>)

```go
func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool
```

Set Insert the key In treelist, if key exists, cover

### func \(\*Tree\[KEY, VALUE\]\) [Size](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L64>)

```go
func (tree *Tree[KEY, VALUE]) Size() int64
```

Size return the size of treelist

### func \(\*Tree\[KEY, VALUE\]\) [Slices](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L409>)

```go
func (tree *Tree[KEY, VALUE]) Slices() []Slice[KEY, VALUE]
```

Slices  return all slice. from smallest to largest.

### func \(\*Tree\[KEY, VALUE\]\) [Tail](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L457>)

```go
func (tree *Tree[KEY, VALUE]) Tail() *Slice[KEY, VALUE]
```

Tail returns the tail of the ordered data of tree

### func \(\*Tree\[KEY, VALUE\]\) [Traverse](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L383>)

```go
func (tree *Tree[KEY, VALUE]) Traverse(every func(s *Slice[KEY, VALUE]) bool)
```

Traverse the traversal method defaults to LDR. from smallest to largest.

### func \(\*Tree\[KEY, VALUE\]\) [Trim](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L749>)

```go
func (tree *Tree[KEY, VALUE]) Trim(low, hight KEY)
```

Trim retain the value of the range . \[low high\]

### func \(\*Tree\[KEY, VALUE\]\) [TrimByIndex](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L835>)

```go
func (tree *Tree[KEY, VALUE]) TrimByIndex(low, hight int64)
```

TrimByIndex retain the value of the index range . \[low high\]

### func \(\*Tree\[KEY, VALUE\]\) [UnionSets](<https://github.com/474420502/structure/blob/master/tree/treelist/tree.go#L974>)

```go
func (tree *Tree[KEY, VALUE]) UnionSets(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]
```

UnionSets tree unionsets with other. \[1 2 3\] \[2 3 4\] \-\> \[1 2 3 4\].



## examples

```go
package main

import (
	"fmt"
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/treelist"
)

func main() {

	// New a object of tree
	tree1 := treelist.New[int, int](compare.Any[int])

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
	tree2 := treelist.New[int, int](compare.Any[int])
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
	tree1.PutDuplicate(10, 10, func(exists *treelist.Slice[int, int]) {
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
	tree1.Traverse(func(s *treelist.Slice[int, int]) bool {
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

	tree := treelist.New[int, int](compare.Any[int])
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
		iter.Range(func(cur *treelist.SliceIndex[int, int]) bool {
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
		iter.Range(func(cur *treelist.SliceIndex[int, int]) bool {
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
		iter.Range(func(cur *treelist.SliceIndex[int, int]) bool {
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
		iter.Range(func(cur *treelist.SliceIndex[int, int]) bool {
			result = append(result, cur.Key)
			return true
		})
		log.Println(result) // [13 11]
	}()

}

func Resotre[KEY, VALUE int](tree1 *treelist.Tree[KEY, VALUE]) {
	tree1.Clear()
	tree1.Put(0, 0) // true
	tree1.Put(4, 4)
	tree1.Put(1, 1)
	// tree1.Put(2, 2)
	tree1.Set(3, 3) //   3 insert
	tree1.Set(7, 7) //   7
}

func Slice2String[KEY any, VALUE any](s *treelist.Slice[KEY, VALUE]) string {
	return fmt.Sprintf("{%v:%v}", s.Key, s.Value)
}

func Tree2String[KEY, VALUE any](tree *treelist.Tree[KEY, VALUE]) []string {
	var results []string
	for _, s := range tree.Slices() {
		results = append(results, Slice2String(&s))
	}
	return results
}
```
