# heap

```go
import "github.com/474420502/structure/tree/heap"
```

`tree/heap` provides a generic binary heap stored in an array.

## Features

- generic `T` values with caller-provided ordering
- `Put`, `Top`, and `Pop` for priority-queue style access
- `Clear` to reset length while keeping capacity
- `Reset` to release backing storage

## API Snapshot

- `New[T any](comp compare.Compare[T]) *Tree[T]`
- `Put(v T)`
- `Top() (T, bool)`
- `Pop() (T, bool)`
- `Clear()`
- `Reset()`
- `Empty() bool`
- `Size() int`

## Notes

- Heap order depends entirely on the comparator passed to `New`.
- `Clear` keeps the current allocation for reuse, while `Reset` drops cached storage.
- Example usage is available at [../../example/heap/main.go](../../example/heap/main.go).

## Validation

Behavior is covered by [heap_test.go](./heap_test.go).