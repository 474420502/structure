# rbtree

```go
import "github.com/474420502/structure/tree/rbtree"
```

`tree/rbtree` is a conventional ordered red-black tree for workloads that want classic logarithmic balancing with a relatively small rotation budget.

## Features

- generic ordered `KEY` and `VALUE`
- classic red-black balancing with parent-linked nodes
- insert-only writes through `Put`
- overwrite-or-insert behavior through `Set`
- standardized semantic aliases for insert, upsert, delete, and size access
- bidirectional iterator with `Seek*` navigation
- benchmark counters and shape statistics for cross-tree comparison

## API Snapshot

- `New[KEY, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]`
- `Put(key KEY, value VALUE) bool`
- `InsertIfAbsent(key KEY, value VALUE) bool`
- `Set(key KEY, value VALUE) bool`
- `Upsert(key KEY, value VALUE) bool`
- `Get(key KEY) (VALUE, bool)`
- `Remove(key KEY) (VALUE, bool)`
- `Delete(key KEY) (VALUE, bool)`
- `Clear()`
- `Size() int64`
- `Len() int`
- `Height() int`
- `Traverse(func(KEY, VALUE) bool)`
- `Values() []VALUE`
- `Iterator() *Iterator[KEY, VALUE]`
- `ResetBenchmarkStats()`
- `BenchmarkStats() BenchmarkStats`

The iterator supports `SeekToFirst`, `SeekToLast`, `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`, `Next`, `Prev`, `Clone`, `Key`, `Value`, and `Valid`.

## Notes

- This package expects a standard comparator contract such as `compare.Any`, where negative means less-than, positive means greater-than, and zero means equality.
- `Put` preserves the existing value when a key already exists.
- `Set` inserts when absent and overwrites when present.
- `Upsert` returns whether an existing value was replaced.
- This tree does not currently expose rank/index operations, so it is compared only on the shared ordered-map operations.
- Example usage is available at [../../example/rbtree/main.go](../../example/rbtree/main.go).

## Validation

Behavior is covered by [rbtree_test.go](./rbtree_test.go).