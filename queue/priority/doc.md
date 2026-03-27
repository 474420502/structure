# priority queue tree

```go
import treequeue "github.com/474420502/structure/queue/priority"
```

`queue/priority` provides an ordered priority queue backed by a size-balanced binary search tree.

## Features

- generic `KEY` and `VALUE`
- ordered insertion with `Put`
- point lookup with `Get`
- duplicate-key collection with `Gets`
- ordered iteration through an iterator API
- indexed access and indexed removal
- head/tail popping for minimum and maximum elements

## API Snapshot

- `New[KEY, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]`
- `Put(key KEY, value VALUE) bool`
- `Get(key KEY) (VALUE, bool)`
- `Gets(key KEY) []VALUE`
- `Index(idx int) VALUE`
- `Remove(key KEY) (VALUE, bool)`
- `RemoveIndex(index int) VALUE`
- `PopHead() (VALUE, bool)`
- `PopTail() (VALUE, bool)`
- `Traverse(func(KEY, VALUE) bool)`
- `Values() []VALUE`
- `Iterator() *Iterator[KEY, VALUE]`

The iterator supports:

- `SeekToFirst`, `SeekToLast`
- `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`
- `Next`, `Prev`, `Clone`
- `Key`, `Value`, `Vaild`

## Notes

- Despite the package path, this behaves like an ordered tree container rather than a binary heap.
- The package name used in code is `treequeue`.
- Example usage is available at [../../example/priority_queue/main.go](../../example/priority_queue/main.go).

## Validation

Behavior is covered by [tree_test.go](./tree_test.go).