<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# liststack

```go
import "github.com/474420502/structure/stack/list"
```

## Index

- [type Stack](<#type-stack>)
  - [func New\[T any\]() *Stack[T]](<#func-new>)
  - [func (ls *Stack[T]) Clear()](<#func-stackt-clear>)
  - [func (ls *Stack[T]) Empty() bool](<#func-stackt-empty>)
  - [func (ls *Stack[T]) Peek() (T, bool)](<#func-stackt-peek>)
  - [func (ls *Stack[T]) Pop() (T, bool)](<#func-stackt-pop>)
  - [func (ls *Stack[T]) Push(v T)](<#func-stackt-push>)
  - [func (ls *Stack[T]) Size() uint](<#func-stackt-size>)
  - [func (ls *Stack[T]) String() string](<#func-stackt-string>)
  - [func (ls *Stack[T]) Values() []T](<#func-stackt-values>)

- [examples](<#examples>)

## type [Stack](<#examples>)

Stack the struct of stack

```go
type Stack[T any] struct {
    // contains filtered or unexported fields
}
```

### func [New](<#examples>)

```go
func New[T any]() *Stack[T]
```

New  create a object of Stack

### func \(\*Stack\[T\]\) [Clear](<#examples>)

```go
func (ls *Stack[T]) Clear()
```

Clear Clear stack data

### func \(\*Stack\[T\]\) [Empty](<#examples>)

```go
func (ls *Stack[T]) Empty() bool
```

Empty if stack is empty\, return true\. else false

### func \(\*Stack\[T\]\) [Peek](<#examples>)

```go
func (ls *Stack[T]) Peek() (T, bool)
```

Peek the top of stack

### func \(\*Stack\[T\]\) [Pop](<#examples>)

```go
func (ls *Stack[T]) Pop() (T, bool)
```

Pop pop the value from stack

### func \(\*Stack\[T\]\) [Push](<#examples>)

```go
func (ls *Stack[T]) Push(v T)
```

Push Push value into stack

### func \(\*Stack\[T\]\) [Size](<#examples>)

```go
func (ls *Stack[T]) Size() uint
```

Size return the size of stack

### func \(\*Stack\[T\]\) [String](<#examples>)

```go
func (ls *Stack[T]) String() string
```

String return the string of stack\. a\(top\)\-\>b\-\>c

### func \(\*Stack\[T\]\) [Values](<#examples>)

```go
func (ls *Stack[T]) Values() []T
```

Values return the values of stacks

## examples

```go
package main

import (
	"log"

	liststack "github.com/474420502/structure/stack/list"
)

func main() {

	st := liststack.New[int]()

	log.Println("Push String Size")
	for i := 0; i < 10; i += 2 {
		st.Push(i)
	}
	log.Println(st.String()) // [0 2 4 6 8]
	log.Println(st.Size())   // 5

	log.Println("Peek Pop Empty Clear")
	log.Println(st.Peek()) // 8 true
	log.Println(st.Pop())  // 8 true
	st.Clear()
	log.Println(st.Empty()) // true
	log.Println(st.Peek())  // 0 false
	log.Println(st.Pop())   // 0 false

	log.Println("Values")
	for i := 0; i < 10; i += 2 {
		st.Push(i)
	}
	log.Println(st.Values()) // [0 2 4 6 8]
}
```
 