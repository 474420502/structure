# list queue

```go
import "github.com/474420502/structure/queue/list"
```

`queue/list` implements a generic deque backed by a doubly linked list.

## Features

- `PushFront` and `PushBack`
- `PopFront` and `PopBack`
- `Front` and `Back` return explicit node handles
- node handles expose `Prev`, `Next`, and `Value`

## API Snapshot

- `New[T any]() *ListQueue[T]`
- `Size() int64`
- `Front() *Element[T]`
- `Back() *Element[T]`
- `PushFront(value T)`
- `PushBack(value T)`
- `PopFront() interface{}`
- `PopBack() interface{}`

`Element[T]` provides:

- `Prev() *Element[T]`
- `Next() *Element[T]`
- `Value() interface{}`

## Notes

- This package is useful when you need stable node links at the ends of the queue.
- Like `queue/linkedarray`, it works as a deque, not only as a FIFO queue.
- The current API returns `interface{}` from `Value` and `Pop*`.

## Validation

Behavior is covered by [queue_test.go](./queue_test.go).