# avl

```go
import "github.com/474420502/structure/tree/avl"
```

`tree/avl` provides a generic ordered map backed by an AVL tree.

## Features

- generic `KEY` and `VALUE`
- strict or relaxed balancing via `NewEx`
- ordered lookup, insert, update, and removal
- in-order traversal with iterators
- optional benchmark counters for rotation and shape analysis

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
- `ResetBenchmarkStats()`
- `BenchmarkStats() BenchmarkStats`

The iterator supports:

- `SeekToFirst`, `SeekToLast`
- `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`
- `Next`, `Prev`, `Clone`
- `Key`, `Value`, `Valid`

## Notes

- `Put` preserves the existing value when the key already exists and returns `false` in that case.
- `Set` overwrites existing keys and inserts missing ones.
- `NewEx` allows a larger height difference than a strict AVL tree, which can shift the balance between update cost and lookup shape.
- Example usage is available at [../../example/avl/main.go](../../example/avl/main.go).

## Validation

Behavior is covered by [avl_test.go](./avl_test.go) and [iterator_test.go](./iterator_test.go).