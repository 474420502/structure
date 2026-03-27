# array list

```go
import arraylist "github.com/474420502/structure/list/array_list"
```

`list/array_list` provides a generic sequential list backed by a growable array.

## Features

- generic `T` values with a comparator used by `Contains`
- O(1) amortized push and fast indexed reads
- front and back insertion/removal helpers
- mutable iterators and circular iterators
- snapshot export with `Values`

## API Snapshot

- `New[T any](comp compare.Compare[T]) *ArrayList[T]`
- `Push(value T)`
- `PushFront(values ...T)`
- `PushBack(values ...T)`
- `Front() (T, bool)`
- `Back() (T, bool)`
- `PopFront() (T, bool)`
- `PopBack() (T, bool)`
- `Index(idx uint) T`
- `Set(idx int, value T)`
- `Remove(idx uint) T`
- `Contains(values ...T) int`
- `Traverse(func(idx uint, value T) bool)`
- `Values() []T`
- `Iterator() *Iterator[T]`
- `CircularIterator() *CircularIterator[T]`

The iterator APIs support:

- `Value`, `Vaild`, `Index`, `IndexTo`
- `Next`, `Prev`, `ToHead`, `ToTail`
- `SetValue`, `Swap`
- `RemoveToNext`, `RemoveToPrev`

## Notes

- This package is optimized for dense sequential storage and random indexed access.
- Inserts and removals near the front require element shifting, so they are more expensive than append-heavy workloads.
- Example usage is available at [../../example/array_list/main.go](../../example/array_list/main.go).

## Validation

Behavior is covered by [array_list_test.go](./array_list_test.go) and [force_test.go](./force_test.go).