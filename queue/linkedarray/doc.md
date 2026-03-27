# linkedarray queue

```go
import "github.com/474420502/structure/queue/linkedarray"
```

`queue/linkedarray` implements a generic double-ended queue on top of a circular array buffer.

## Features

- `PushFront` and `PushBack` for deque-style insertion
- `PopFront` and `PopBack` for removal from both ends
- `Front`, `Back`, and `Index` for direct access
- automatic grow and shrink of the backing buffer
- `Traverse` and `Values` for ordered iteration

## API Snapshot

- `New[T any]() *ArrayQueue[T]`
- `PushFront(value T)`
- `PushBack(value T)`
- `PopFront() interface{}`
- `PopBack() interface{}`
- `Front() interface{}`
- `Back() interface{}`
- `Index(idx int64) interface{}`
- `Size() int64`
- `Traverse(func(idx int64, value T) bool)`
- `Values() []T`

## Notes

- This package behaves as a deque rather than a strict FIFO queue.
- The public read/pop accessors currently return `interface{}` even though the container itself is generic.
- The implementation is intended for single-threaded use.

## Validation

Behavior is covered by [queue_test.go](./queue_test.go).