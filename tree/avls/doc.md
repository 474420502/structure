# avls

```go
import "github.com/474420502/structure/tree/avls"
```

`tree/avls` is the duplicate-key variant of the AVL tree implementation.

## Features

- generic `KEY` and `VALUE`
- duplicate keys are stored as independent nodes
- standardized semantic aliases for insert, upsert, delete, and size access
- ordered traversal over the full multiset
- boundary-aware iterators that can land on the first or last node in an equal-key run
- balancing configuration via `NewEx`

## API Snapshot

- `New[KEY, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]`
- `NewEx[KEY, VALUE any](comp compare.Compare[KEY], differenceHeight int8) *Tree[KEY, VALUE]`
- `Put(key KEY, value VALUE) bool`
- `InsertIfAbsent(key KEY, value VALUE) bool`
- `Set(key KEY, value VALUE) bool`
- `Upsert(key KEY, value VALUE) bool`
- `Get(key KEY) (VALUE, bool)`
- `Remove(key KEY) (VALUE, bool)`
- `Delete(key KEY) (VALUE, bool)`
- `Traverse(func(KEY, VALUE) bool)`
- `Values() []VALUE`
- `Clear()`
- `Size() uint`
- `Len() int`
- `Height() int8`
- `Iterator() *Iterator[KEY, VALUE]`

The iterator supports:

- `SeekToFirst`, `SeekToLast`
- `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`
- `Next`, `Prev`, `Clone`
- `Key`, `Value`, `Valid`

## Notes

- `Put` always inserts a new node, even when the key already exists.
- `Set` updates one existing node for the key when present, otherwise it inserts a new node.
- `InsertIfAbsent` is the preferred explicit name when duplicate keys should be rejected.
- `Upsert` is the preferred explicit name for overwrite-or-create writes against the first matching key.
- `Delete` and `Len` provide the preferred cross-package removal and size entry points for new code.
- The iterator `Seek*` methods already return the exact-match status directly. This package does not need separate `Seek*Exact` aliases.
- `Get` returns the value at the stable front of the duplicate-key run.
- `SeekGE` and `SeekLE` position to the first or last node of an equal-key run respectively.

## Validation

Behavior is covered by [avl_test.go](./avl_test.go) and [iterator_test.go](./iterator_test.go).