<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# itree

```go
import "github.com/474420502/structure/tree/itree"
```

## Index

- [type Iterator](<#type-iterator>)
  - [func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE]](<#func-iteratorkey-value-clone>)
  - [func (iter *Iterator[KEY, VALUE]) Key() KEY](<#func-iteratorkey-value-key>)
  - [func (iter *Iterator[KEY, VALUE]) Next()](<#func-iteratorkey-value-next>)
  - [func (iter *Iterator[KEY, VALUE]) Prev()](<#func-iteratorkey-value-prev>)
  - [func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY)](<#func-iteratorkey-value-seekge>)
  - [func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY)](<#func-iteratorkey-value-seekgt>)
  - [func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY)](<#func-iteratorkey-value-seekle>)
  - [func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY)](<#func-iteratorkey-value-seeklt>)
  - [func (iter *Iterator[KEY, VALUE]) SeekToFirst()](<#func-iteratorkey-value-seektofirst>)
  - [func (iter *Iterator[KEY, VALUE]) SeekToLast()](<#func-iteratorkey-value-seektolast>)
  - [func (iter *Iterator[KEY, VALUE]) Vaild() bool](<#func-iteratorkey-value-vaild>)
  - [func (iter *Iterator[KEY, VALUE]) Value() VALUE](<#func-iteratorkey-value-value>)
- [type Tree](<#type-tree>)
  - [func New[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE]](<#func-new>)
  - [func (tree *Tree[KEY, VALUE]) Clear()](<#func-treekey-value-clear>)
  - [func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool)](<#func-treekey-value-get>)
  - [func (tree *Tree[KEY, VALUE]) Index(idx int) VALUE](<#func-treekey-value-index>)
  - [func (tree *Tree[KEY, VALUE]) IndexOf(key KEY) int](<#func-treekey-value-indexof>)
  - [func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE]](<#func-treekey-value-iterator>)
  - [func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool](<#func-treekey-value-put>)
  - [func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool)](<#func-treekey-value-remove>)
  - [func (tree *Tree[KEY, VALUE]) RemoveIndex(index int) VALUE](<#func-treekey-value-removeindex>)
  - [func (tree *Tree[KEY, VALUE]) RemoveRange(low, high KEY)](<#func-treekey-value-removerange>)
  - [func (tree *Tree[KEY, VALUE]) RemoveRangeByIndex(low, high int)](<#func-treekey-value-removerangebyindex>)
  - [func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool](<#func-treekey-value-set>)
  - [func (tree *Tree[KEY, VALUE]) Size() int](<#func-treekey-value-size>)
  - [func (tree *Tree[KEY, VALUE]) Split(key KEY) *Tree[KEY, VALUE]](<#func-treekey-value-split>)
  - [func (tree *Tree[KEY, VALUE]) Traverse(every func(KEY, VALUE) bool)](<#func-treekey-value-traverse>)
  - [func (tree *Tree[KEY, VALUE]) Trim(low, high KEY)](<#func-treekey-value-trim>)
  - [func (tree *Tree[KEY, VALUE]) TrimByIndex(low, high int)](<#func-treekey-value-trimbyindex>)
  - [func (tree *Tree[KEY, VALUE]) Values() []VALUE](<#func-treekey-value-values>)


## type [Iterator](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L4-L11>)

Iterator tree iterator

```go
type Iterator[KEY, VALUE any] struct {
    // contains filtered or unexported fields
}
```

### func \(\*Iterator\[KEY, VALUE\]\) [Clone](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L84>)

```go
func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE]
```

Clone Copy a current iterator

### func \(\*Iterator\[KEY, VALUE\]\) [Key](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L19>)

```go
func (iter *Iterator[KEY, VALUE]) Key() KEY
```

Key return the key of current iterator

### func \(\*Iterator\[KEY, VALUE\]\) [Next](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L79>)

```go
func (iter *Iterator[KEY, VALUE]) Next()
```

Next the current iterator move to the next. before call it must call Vaild\(\) and return true.

### func \(\*Iterator\[KEY, VALUE\]\) [Prev](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L74>)

```go
func (iter *Iterator[KEY, VALUE]) Prev()
```

Prev the current iterator move to the prev. before call it must call Vaild\(\) and return true.

### func \(\*Iterator\[KEY, VALUE\]\) [SeekGE](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L64>)

```go
func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY)
```

SeekGE seek to the key that greater than or equal to

### func \(\*Iterator\[KEY, VALUE\]\) [SeekGT](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L69>)

```go
func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY)
```

SeekGT seek to the key that greater than

### func \(\*Iterator\[KEY, VALUE\]\) [SeekLE](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L54>)

```go
func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY)
```

SeekLE seek to the key that less than or equal to

### func \(\*Iterator\[KEY, VALUE\]\) [SeekLT](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L59>)

```go
func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY)
```

SeekLT seek to the key that less than

### func \(\*Iterator\[KEY, VALUE\]\) [SeekToFirst](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L34>)

```go
func (iter *Iterator[KEY, VALUE]) SeekToFirst()
```

SeekToFirst seek to first item

### func \(\*Iterator\[KEY, VALUE\]\) [SeekToLast](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L44>)

```go
func (iter *Iterator[KEY, VALUE]) SeekToLast()
```

SeekToFirst seek to last item

### func \(\*Iterator\[KEY, VALUE\]\) [Vaild](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L29>)

```go
func (iter *Iterator[KEY, VALUE]) Vaild() bool
```

Vaild if current value is not nil return true. else return false. for use with Seek

### func \(\*Iterator\[KEY, VALUE\]\) [Value](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L24>)

```go
func (iter *Iterator[KEY, VALUE]) Value() VALUE
```

Value return the value of current iterator

## type [Node](<https://github.com/474420502/structure/blob/master/tree/itree/node.go#L7-L12>)

```go
type Node[KEY any, VALUE any] struct {
    Key      KEY
    Value    VALUE
    Size     int
    Children [2]*Node[KEY, VALUE]
}
```

### func \(\*Node\[KEY, VALUE\]\) [String](<https://github.com/474420502/structure/blob/master/tree/itree/node.go#L14>)

```go
func (node *Node[KEY, VALUE]) String() string
```

## type [NodeDir](<https://github.com/474420502/structure/blob/master/tree/itree/iterator.go#L13-L16>)

```go
type NodeDir[KEY any, VALUE any] struct {
    N   *Node[KEY, VALUE]
    D   int8
}
```

## type [Tree](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L9-L15>)

```go
type Tree[KEY, VALUE any] struct {
    Center  *Node[KEY, VALUE]
    Compare compare.Compare[KEY]
    // contains filtered or unexported fields
}
```

### func [New](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L17>)

```go
func New[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE]
```

### func \(\*Tree\[KEY, VALUE\]\) [Clear](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L263>)

```go
func (tree *Tree[KEY, VALUE]) Clear()
```

### func \(\*Tree\[KEY, VALUE\]\) [Get](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L45>)

```go
func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool)
```

### func \(\*Tree\[KEY, VALUE\]\) [Index](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L54>)

```go
func (tree *Tree[KEY, VALUE]) Index(idx int) VALUE
```

### func \(\*Tree\[KEY, VALUE\]\) [IndexOf](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L68>)

```go
func (tree *Tree[KEY, VALUE]) IndexOf(key KEY) int
```

### func \(\*Tree\[KEY, VALUE\]\) [Iterator](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L304>)

```go
func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE]
```

### func \(\*Tree\[KEY, VALUE\]\) [Put](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L36>)

```go
func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool
```

### func \(\*Tree\[KEY, VALUE\]\) [Remove](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L100>)

```go
func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool)
```

### func \(\*Tree\[KEY, VALUE\]\) [RemoveIndex](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L109>)

```go
func (tree *Tree[KEY, VALUE]) RemoveIndex(index int) VALUE
```

RemoveIndex remove key value by index and return value that be removed

### func \(\*Tree\[KEY, VALUE\]\) [RemoveRange](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L169>)

```go
func (tree *Tree[KEY, VALUE]) RemoveRange(low, high KEY)
```

RemoveRange remove keys values by range. \[low, high\]

### func \(\*Tree\[KEY, VALUE\]\) [RemoveRangeByIndex](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L206>)

```go
func (tree *Tree[KEY, VALUE]) RemoveRangeByIndex(low, high int)
```

RemoveRangeByIndex 1.remove range \[low:high\] 2.low and hight that the range must contain a value that exists. eg: \[low: high\+1\] \[low\-1: high\].  \[low\-1: hight\+1\]. error: \[min\-1:min\-2\] or \[max\+1:max\+2\]

### func \(\*Tree\[KEY, VALUE\]\) [Set](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L28>)

```go
func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool
```

### func \(\*Tree\[KEY, VALUE\]\) [Size](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L267>)

```go
func (tree *Tree[KEY, VALUE]) Size() int
```

### func \(\*Tree\[KEY, VALUE\]\) [Split](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L234>)

```go
func (tree *Tree[KEY, VALUE]) Split(key KEY) *Tree[KEY, VALUE]
```

Split  original tree contain Key. return  the splited tree. eg: 1.\[1 4 5 7\] \-\> Split\(5\) \[1 4 5\] \[7\]; 2.\[1 4 5 7\] \-\> Split\(3\) \[1\] \[4 5 7\]

### func \(\*Tree\[KEY, VALUE\]\) [Traverse](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L271>)

```go
func (tree *Tree[KEY, VALUE]) Traverse(every func(KEY, VALUE) bool)
```

### func \(\*Tree\[KEY, VALUE\]\) [Trim](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L126>)

```go
func (tree *Tree[KEY, VALUE]) Trim(low, high KEY)
```

RemoveRange remove keys values by range. \[low, high\]

### func \(\*Tree\[KEY, VALUE\]\) [TrimByIndex](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L139>)

```go
func (tree *Tree[KEY, VALUE]) TrimByIndex(low, high int)
```

TrimByIndex retain the value of the index range . \[low high\]

### func \(\*Tree\[KEY, VALUE\]\) [Values](<https://github.com/474420502/structure/blob/master/tree/itree/tree.go#L292>)

```go
func (tree *Tree[KEY, VALUE]) Values() []VALUE
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
