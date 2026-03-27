# search treelist

```go
import "github.com/474420502/structure/search/treelist"
```

`search/treelist` is the byte-keyed ordered tree variant used for search-oriented workloads.

## Features

- `[]byte` keys with configurable comparison via `SetCompare`
- ordered lookup by key and direct lookup by rank
- point removal, range removal, and trimming by key or index
- iterator and iterator-range APIs
- set-style operations over ordered byte-key collections
- duplicate-aware insertion hook through `PutDuplicate`

## API Snapshot

- `New() *Tree`
- `SetCompare(comp compare.Compare[[]byte])`
- `Put(key []byte, value interface{}) bool`
- `PutDuplicate(key []byte, value interface{}, do func(exists *Slice)) bool`
- `Set(key []byte, value interface{}) bool`
- `Get(key []byte) (interface{}, bool)`
- `Index(i int64) *Slice`
- `IndexOf(key []byte) int64`
- `Head() *Slice`
- `Tail() *Slice`
- `Remove(key []byte) *Slice`
- `RemoveIndex(index int64) *Slice`
- `RemoveHead() *Slice`
- `RemoveTail() *Slice`
- `RemoveRange(low, hight []byte) bool`
- `RemoveRangeByIndex(low, hight int64)`
- `Trim(low, hight []byte)`
- `TrimByIndex(low, hight int64)`
- `Traverse(func(*Slice) bool)`
- `Slices() []Slice`
- `Intersection(other *Tree) *Tree`
- `UnionSets(other *Tree) *Tree`
- `DifferenceSets(other *Tree) *Tree`
- `Iterator() *Iterator`
- `IteratorRange() *IteratorRange`

The iterator supports:

- `SeekToFirst`, `SeekToLast`, `SeekByIndex`
- `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`
- `Next`, `Prev`, `Clone`
- `Index`, `Compare`, `Slice`, `Key`, `Value`, `Valid`

The range iterator supports:

- `GE2LE`, `GT2LE`, `GE2LT`, `GT2LT`
- `SetDirection`, `Direction`, `Size`, `Range`

## Notes

- The default comparator is byte-oriented, but `SetCompare` allows callers to swap in custom ordering logic.
- This package mirrors the generic `tree/treelist` API with `[]byte` keys and `interface{}` values.
- Example usage is available at [../../example/search-treelist/main.go](../../example/search-treelist/main.go).

## Validation

Behavior is covered by [tree_test.go](./tree_test.go), [iterator_test.go](./iterator_test.go), and [iterator_range_test.go](./iterator_range_test.go).