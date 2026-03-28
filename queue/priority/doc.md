# priority queue tree

```go
import treequeue "github.com/474420502/structure/queue/priority"
```

`queue/priority` provides an ordered priority queue backed by a size-balanced binary search tree.

## Features

- generic `KEY` and `VALUE`
- ordered insertion with `Put`
- standardized semantic aliases for insert, upsert, delete, and size access
- point lookup with `Get`
- duplicate-key collection with `Gets`
- ordered iteration through an iterator API
- indexed access and indexed removal
- head/tail popping for minimum and maximum elements

## API Snapshot

- `New[KEY, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]`
- `Put(key KEY, value VALUE) bool`
- `InsertIfAbsent(key KEY, value VALUE) bool`
- `Upsert(key KEY, value VALUE) bool`
- `Get(key KEY) (VALUE, bool)`
- `Gets(key KEY) []VALUE`
- `Index(idx int) VALUE`
- `Remove(key KEY) (VALUE, bool)`
- `Delete(key KEY) (VALUE, bool)`
- `RemoveIndex(index int) VALUE`
- `PopHead() (VALUE, bool)`
- `PopTail() (VALUE, bool)`
- `Traverse(func(KEY, VALUE) bool)`
- `Values() []VALUE`
- `Iterator() *Iterator[KEY, VALUE]`
- `Len() int`

The iterator supports:

- `SeekToFirst`, `SeekToLast`
- `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`
- `SeekGEExact`, `SeekGTExact`, `SeekLEExact`, `SeekLTExact`
- `Next`, `Prev`, `Clone`
- `Key`, `Value`, `Vaild`

## Notes

- Despite the package path, this behaves like an ordered tree container rather than a binary heap.
- The package name used in code is `treequeue`.
- `Put` keeps its historical duplicate-key behavior.
- `InsertIfAbsent` is the preferred explicit name when duplicate keys should be rejected.
- `Upsert` updates the first matching key when present and inserts when absent.
- `Delete` and `Len` provide the preferred cross-package removal and size entry points for new code.
- `Seek*Exact` methods perform the same iterator positioning as `Seek*` and additionally report whether the queried key existed exactly.
- Example usage is available at [../../example/priority_queue/main.go](../../example/priority_queue/main.go).

## Validation

Behavior is covered by [tree_test.go](./tree_test.go).