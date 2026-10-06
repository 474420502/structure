# linked hashmap

```go
import linkedhashmap "github.com/474420502/structure/map/linkedhashmap"
```

`map/linkedhashmap` combines hash lookup with a linked order of entries.

## Features

- generic `K comparable, V any` keys and values (no interface boxing)
- stable traversal order from head to tail
- append/prepend insertion helpers
- update-and-move operations with `SetFront` and `SetBack`
- standardized `InsertIfAbsent`, `Upsert`, `Delete`, and `Len` helpers
- ordered snapshots through `Keys`, `Values`, and `Slices`

## API Snapshot

- `New[K comparable, V any]() *LinkedHashmap[K, V]`
- `Put(key K, value V) bool`
- `InsertIfAbsent(key K, value V) bool`
- `PushBack(key K, value V) bool`
- `PushFront(key K, value V) bool`
- `Set(key K, value V) bool`
- `Upsert(key K, value V) bool`
- `SetBack(key K, value V) bool`
- `SetFront(key K, value V) bool`
- `Get(key K) (V, bool)`
- `Remove(key K) (V, bool)`
- `Delete(key K) (V, bool)`
- `Keys() []K`
- `Values() []V`
- `Slices() []Slice[K, V]`
- `Clear()`
- `Empty() bool`
- `Size() uint`
- `Len() int`
- `String() string`

## Notes

- `Put` is an alias for `PushBack` and does not overwrite existing entries.
- `Set` only updates existing entries and returns `false` when the key is absent.
- `InsertIfAbsent` is the preferred explicit name for insert-only writes.
- `Upsert` updates existing entries and appends new entries to the back when the key is absent.
- `Delete` and `Len` provide the preferred cross-package removal and size entry points for new code.
- `SetFront` and `SetBack` both update insertion order as part of the write.
- Example usage is available at [../../example/linkedhashmap/main.go](../../example/linkedhashmap/main.go).

## Validation

Behavior is covered by [linked_hashmap_test.go](./linked_hashmap_test.go).
