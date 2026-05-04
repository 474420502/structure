# indextree

```go
import "github.com/474420502/structure/tree/indextree"
```

`tree/indextree` is an ordered tree that balances by subtree size and exposes rank/index-oriented operations.

## Features

- generic ordered keys and typed generic values
- direct lookup by key and by in-order index
- point removal, range removal, and trimming by key or index
- split operations for partitioning a tree around a key
- bidirectional iterator with rank-aware positioning
- standardized semantic aliases for insert, upsert, delete, and size access
- benchmark counters and related performance analysis documents

## API Snapshot

- `New[T any](comp compare.Compare[T]) *Tree[T, T]`
- `NewWithValue[KEY, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]`
- `Put(key KEY, value VALUE) bool`
- `InsertIfAbsent(key KEY, value VALUE) bool`
- `Set(key KEY, value VALUE) bool`
- `Upsert(key KEY, value VALUE) bool`
- `Get(key KEY) (VALUE, bool)`
- `Index(i int64) (key KEY, value VALUE)`
- `IndexOf(key KEY) int64`
- `Remove(key KEY) (VALUE, bool)`
- `Delete(key KEY) (VALUE, bool)`
- `RemoveIndex(index int64) (VALUE, bool)`
- `RemoveRange(low, high KEY)`
- `RemoveRangeByIndex(low, hight int64)`
- `Trim(low, high KEY)`
- `TrimByIndex(low, high int64)`
- `Split(key KEY) *Tree[KEY, VALUE]`
- `SplitContain(key KEY) *Tree[KEY, VALUE]`
- `Traverse(func(KEY, VALUE) bool)`
- `Values() []VALUE`
- `Clear()`
- `Size() int64`
- `Len() int`
- `String() string`
- `Iterator() *Iterator[KEY, VALUE]`
- `ResetBenchmarkStats()`
- `BenchmarkStats() BenchmarkStats`

The iterator supports:

- `SeekToFirst`, `SeekToLast`, `SeekByIndex`
- `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`
- `Next`, `Prev`, `Clone`
- `Index`, `Key`, `Value`, `Valid`

## Notes

- `Put` preserves the existing value when the key already exists and returns `false` in that case.
- `Set` overwrites existing keys and inserts missing ones.
- `New` is the convenience constructor for the common `key == value` case; use `NewWithValue` when the value type differs from the key type.
- `InsertIfAbsent` is the preferred explicit name for insert-only writes.
- `Upsert` is the preferred explicit name for overwrite-or-create writes and returns whether an existing value was replaced.
- `Delete` and `Len` provide the preferred cross-package removal and size entry points for new code.
- The iterator `Seek*` methods already return the exact-match status directly. This package does not need separate `Seek*Exact` aliases.
- `Split` and `SplitContain` partition the tree around a boundary key and return the detached side as a new tree.
- Related performance documents are available at [benchmark-comparison.md](./benchmark-comparison.md), [benchmark-comparison.zh.md](./benchmark-comparison.zh.md), and [../rotation-analysis.md](../rotation-analysis.md).
- Example usage is available at [../../example/indextree/main.go](../../example/indextree/main.go).

## Validation

Behavior is covered by [tree_test.go](./tree_test.go), [tree_ops_test.go](./tree_ops_test.go), and [iterator_test.go](./iterator_test.go).