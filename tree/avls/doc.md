# avls

```go
import "github.com/474420502/structure/tree/avls"
```

`tree/avls` is the duplicate-key variant of the AVL tree implementation.

## Features

- generic `KEY` and `VALUE`
- duplicate keys are stored as independent nodes
- ordered traversal over the full multiset
- boundary-aware iterators that can land on the first or last node in an equal-key run
- balancing configuration via `NewEx`

## API Snapshot

- `New[KEY, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]`
- `NewEx[KEY, VALUE any](comp compare.Compare[KEY], differenceHeight int8) *Tree[KEY, VALUE]`
- `Put(key KEY, value VALUE) bool`
- `Set(key KEY, value VALUE) bool`
- `Get(key KEY) (VALUE, bool)`
- `Remove(key KEY) (VALUE, bool)`
- `Traverse(func(KEY, VALUE) bool)`
- `Values() []VALUE`
- `Clear()`
- `Size() uint`
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
- `Get` returns the value at the stable front of the duplicate-key run.
- `SeekGE` and `SeekLE` position to the first or last node of an equal-key run respectively.

## Validation

Behavior is covered by [avl_test.go](./avl_test.go) and [iterator_test.go](./iterator_test.go).