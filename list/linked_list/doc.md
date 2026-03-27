# linked list

```go
import linkedlist "github.com/474420502/structure/list/linked_list"
```

`list/linked_list` provides a generic doubly linked list with mutable iterator support.

## Features

- generic `T` values with comparator-based `Contains`
- efficient head and tail insertion/removal
- positional access with `Index`
- standard and circular iterators
- in-place iterator operations for insertion, movement, and removal

## API Snapshot

- `New[T any](comp compare.Compare[T]) *LinkedList[T]`
- `Push(value T)`
- `PushFront(values ...T)`
- `PushBack(values ...T)`
- `Front() (T, bool)`
- `Back() (T, bool)`
- `PopFront() (T, bool)`
- `PopBack() (T, bool)`
- `Index(idx int) (T, bool)`
- `Contains(values ...T) int`
- `Traverse(func(value T) bool)`
- `Values() []T`
- `Iterator() *Iterator[T]`
- `CircularIterator() *CircularIterator[T]`

The iterator APIs support:

- `Value`, `Vaild`, `SetValue`, `Swap`
- `Next`, `Prev`, `ToHead`, `ToTail`, `Move`
- `InsertBefore`, `InsertAfter`
- `MoveBefore`, `MoveAfter`
- `RemoveToNext`, `RemoveToPrev`

## Notes

- This implementation is better suited to frequent middle insertions and removals than the array-backed list.
- Indexed lookup is supported, but it walks the list and is therefore not constant time.
- Example usage is available at [../../example/linked_list/main.go](../../example/linked_list/main.go).

## Validation

Behavior is covered by [linked_list_test.go](./linked_list_test.go).