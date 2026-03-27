# treelist

```go
import "github.com/474420502/structure/tree/treelist"
```

`tree/treelist` provides an ordered map and indexed sequence view over the same size-balanced tree.

## Features

- generic `KEY` and `VALUE`
- ordered lookup by key and direct lookup by rank
- point removal, range removal, and trimming by key or index
- iterator and iterator-range APIs
- set-style operations for ordered collections
- duplicate-aware insertion hook through `PutDuplicate`

## API Snapshot

- `New[KEY, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]`
- `Put(key KEY, value VALUE) bool`
- `PutDuplicate(key KEY, value VALUE, do func(exists *Slice[KEY, VALUE])) bool`
- `Set(key KEY, value VALUE) bool`
- `Get(key KEY) (VALUE, bool)`
- `Index(i int64) *Slice[KEY, VALUE]`
- `IndexOf(key KEY) int64`
- `Head() *Slice[KEY, VALUE]`
- `Tail() *Slice[KEY, VALUE]`
- `Remove(key KEY) *Slice[KEY, VALUE]`
- `RemoveIndex(index int64) *Slice[KEY, VALUE]`
- `RemoveHead() *Slice[KEY, VALUE]`
- `RemoveTail() *Slice[KEY, VALUE]`
- `RemoveRange(low, hight KEY) bool`
- `RemoveRangeByIndex(low, hight int64)`
- `Trim(low, hight KEY)`
- `TrimByIndex(low, hight int64)`
- `Traverse(func(*Slice[KEY, VALUE]) bool)`
- `Slices() []Slice[KEY, VALUE]`
- `Intersection(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]`
- `UnionSets(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]`
- `DifferenceSets(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]`
- `Iterator() *Iterator[KEY, VALUE]`
- `IteratorRange() *IteratorRange[KEY, VALUE]`

The iterator supports:

- `SeekToFirst`, `SeekToLast`, `SeekByIndex`
- `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`
- `Next`, `Prev`, `Clone`
- `Index`, `Compare`, `Slice`, `Key`, `Value`, `Valid`

The range iterator supports:

- `GE2LE`, `GT2LE`, `GE2LT`, `GT2LT`
- `SetDirection`, `Direction`, `Size`, `Range`

## Notes

- This package is the main ordered structure for workloads that need both sorted-key semantics and rank/index operations.
- `Intersection`, `UnionSets`, and `DifferenceSets` treat the tree as an ordered set by key.
- The benchmark comparison for ordered structures lives in [../indextree/benchmark-comparison.md](../indextree/benchmark-comparison.md).
- Example usage is available at [../../example/tree-treelist/main.go](../../example/tree-treelist/main.go).

## Validation

Behavior is covered by [tree_test.go](./tree_test.go), [iterator_test.go](./iterator_test.go), and [iterator_range_test.go](./iterator_range_test.go).