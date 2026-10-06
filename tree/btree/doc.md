# btree

```go
import "github.com/474420502/structure/tree/btree"
```

`tree/btree` is a generic in-memory B-tree ordered map.

## Features

- generic `KEY` and `VALUE` with the standard comparator contract
- high fan-out nodes keep the tree shallow and cache friendly
- point lookup, insert-only, overwrite, and delete
- in-order `Traverse` and `Values`
- `Min` and `Max`
- node capacity is preallocated, so inserts amortize to `0 allocs/op`

## API Snapshot

- `New[K, V any](comp compare.Compare[K]) *Tree[K, V]`
- `NewWithDegree[K, V any](comp compare.Compare[K], degree int) *Tree[K, V]`
- `Put(key K, value V) bool`
- `InsertIfAbsent(key K, value V) bool`
- `Set(key K, value V) bool`
- `Upsert(key K, value V) bool`
- `Get(key K) (V, bool)`
- `Delete(key K) (V, bool)`
- `Remove(key K) (V, bool)`
- `Traverse(func(K, V) bool)`
- `Values() []V`
- `Min() (K, V, bool)`
- `Max() (K, V, bool)`
- `Len() int`, `Size() int`, `Empty() bool`, `Clear()`, `Height() int`

## Notes

- The default minimum degree is `16`, so a node holds up to `31` items.
- `Put` and `InsertIfAbsent` preserve an existing value and report whether a new key was inserted.
- `Set` inserts or overwrites and reports whether a new key was inserted.
- `Upsert` reports whether an existing value was replaced.
- The comparator follows the standard repository contract: negative is less-than, positive is greater-than, zero is equal.
- Nodes use binary search, so wide nodes stay cheap to scan.

## Validation

Behavior and B-tree invariants are covered by [btree_test.go](./btree_test.go) and [api_alias_test.go](./api_alias_test.go).
