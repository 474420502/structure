# heap

```go
import "github.com/474420502/structure/tree/heap"
```

## Index

- [type Tree](<#type-tree>)
  - [func New[T any](Compare compare.Compare[T]) *Tree[T]](<#func-new>)
  - [func (h *Tree[T]) Clear()](<#func-treet-clear>)
  - [func (h *Tree[T]) Empty() bool](<#func-treet-empty>)
  - [func (h *Tree[T]) Pop() (interface{}, bool)](<#func-treet-pop>)
  - [func (h *Tree[T]) Put(v T)](<#func-treet-put>)
  - [func (h *Tree[T]) Reset()](<#func-treet-reset>)
  - [func (h *Tree[T]) Size() int](<#func-treet-size>)
  - [func (h *Tree[T]) Top() (result T, ok bool)](<#func-treet-top>)
- [examples](#examples)
 

## type [Tree](#examples)

Tree the struct of heap with array

```go
type Tree[T any] struct {
    // contains filtered or unexported fields
}
```

### func [New](#examples)

```go
func New[T any](Compare compare.Compare[T]) *Tree[T]
```

New create a  object of heap

### func \(\*Tree\[T\]\) [Clear](#examples)

```go
func (h *Tree[T]) Clear()
```

Clear clear all node\, but not release memory

### func \(\*Tree\[T\]\) [Empty](#examples)

```go
func (h *Tree[T]) Empty() bool
```

Empty if heap size is zero\, return true\. else false

### func \(\*Tree\[T\]\) [Pop](#examples)

```go
func (h *Tree[T]) Pop() (result T, ok bool)
```

Pop pop value from heap

### func \(\*Tree\[T\]\) [Put](#examples)

```go
func (h *Tree[T]) Put(v T)
```

Put put value to heap

### func \(\*Tree\[T\]\) [Reset](#examples)

```go
func (h *Tree[T]) Reset()
```

Reset clear all node and release memory

### func \(\*Tree\[T\]\) [Size](#examples)

```go
func (h *Tree[T]) Size() int
```

Size return the size of heap

### func \(\*Tree\[T\]\) [Top](#examples)

```go
func (h *Tree[T]) Top() (result T, ok bool)
```

Top return the top of heap
 
## Examples

```go
package main

import (
	"log"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/heap"
)

func main() {
	h := heap.New(compare.AnyDesc[int])

	// Put
	for _, v := range []int{5, 2, 3, 45, 10} {
		h.Put(v)
	}
	log.Println(h.Size()) // 5

	// Pop
	var values []int
	for v, ok := h.Pop(); ok; v, ok = h.Pop() {
		values = append(values, v)
	}
	log.Println(values)    // [45 10 5 3 2]
	log.Println(h.Empty()) // true

	h.Put(1)
	log.Println(h.Empty()) // false
	h.Clear()
	log.Println(h.Empty()) // true
}

```
 
