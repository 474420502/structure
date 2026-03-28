# treeset

```go
import "github.com/474420502/structure/set/treeset"
```

`set/treeset` is an ordered tree-backed set or map hybrid with AVL-style balancing.

## Features

- generic `KEY` and `VALUE`
- ordered insertion through `Add`
- overwrite-or-insert behavior through `Set`
- standardized semantic aliases for insert, upsert, delete, and size access
- `Contains`, `Get`, `Remove`
- in-order traversal and value export
- iterator with forward/backward seek operations
- configurable balancing tolerance through `NewEx`
- set operations: `Union`, `Intersection`, `Difference`

## API Snapshot

- `New[KEY, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE]`
- `NewEx[KEY, VALUE any](comp compare.Compare[KEY], differenceHeight int8) *Tree[KEY, VALUE]`
- `Add(key KEY, value VALUE) bool`
- `InsertIfAbsent(key KEY, value VALUE) bool`
- `Set(key KEY, value VALUE) bool`
- `Upsert(key KEY, value VALUE) bool`
- `Contains(item KEY) bool`
- `Get(key KEY) (VALUE, bool)`
- `Remove(key KEY) (VALUE, bool)`
- `Delete(key KEY) (VALUE, bool)`
- `Empty() bool`
- `Clear()`
- `Size() uint`
- `Len() int`
- `Height() int8`
- `Traverse(func(KEY, VALUE) bool)`
- `Values() []VALUE`
- `Iterator() *Iterator[KEY, VALUE]`
- `Union(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]`
- `Intersection(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]`
- `Difference(other *Tree[KEY, VALUE]) *Tree[KEY, VALUE]`

The iterator supports `SeekToFirst`, `SeekToLast`, `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`, `SeekGEExact`, `SeekGTExact`, `SeekLEExact`, `SeekLTExact`, `Next`, `Prev`, `Clone`, and both `Valid` and the legacy `Vaild` compatibility alias.

## Notes

- `Add` does not overwrite an existing key.
- `Set` inserts when absent and overwrites when present.
- `InsertIfAbsent` is the preferred explicit name for insert-only writes.
- `Upsert` is the preferred explicit name for overwrite-or-create writes and returns whether an existing value was replaced.
- `Delete` and `Len` provide the preferred cross-package removal and size entry points for new code.
- `Seek*Exact` methods perform the same iterator positioning as `Seek*` and additionally report whether the queried key existed exactly.
- `Union`, `Intersection`, and `Difference` mutate the receiver, return the receiver, and leave `other` unchanged.
- `Valid` is the preferred iterator validity check. `Vaild` remains available for backward compatibility.
- Example usage is available at [../../example/treeset/main.go](../../example/treeset/main.go).

## Validation

Behavior is covered by [treeset_test.go](./treeset_test.go) and [iterator_test.go](./iterator_test.go).