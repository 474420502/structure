# hashmap

```go
import "github.com/474420502/structure/map/hashmap"
```

`map/hashmap` is a generic wrapper around Go's native hash map.

## Features

- generic `K comparable, V any` keys and values (no interface boxing)
- optional initial capacity with `NewWithCap`
- legacy `Put` and `Set` plus standardized semantic aliases
- snapshot access to keys, values, and key/value slices

## API Snapshot

- `New[K comparable, V any]() *HashMap[K, V]`
- `NewWithCap[K comparable, V any](cap int) *HashMap[K, V]`
- `Put(key K, value V) bool`
- `InsertIfAbsent(key K, value V) bool`
- `Set(key K, value V)`
- `Upsert(key K, value V) bool`
- `Get(key K) (V, bool)`
- `Remove(key K)`
- `Delete(key K) (V, bool)`
- `Keys() []K`
- `Values() []V`
- `Slices() []Slice[K, V]`
- `Clear()`
- `Empty() bool`
- `Size() int`
- `Len() int`
- `String() string`

## Notes

- `Put` returns `false` when the key already exists and leaves the stored value unchanged.
- `Set` always overwrites or creates the entry.
- `InsertIfAbsent` is the preferred explicit name for insert-only writes.
- `Upsert` is the preferred explicit name for overwrite-or-create writes and returns whether an existing value was replaced.
- `Delete` is the preferred explicit removal helper when the previous value is needed.
- `Len` is the preferred cross-package size accessor for new code.
- Iteration order for `Keys`, `Values`, and `Slices` is not stable.
- Example usage is available at [../../example/hashmap/main.go](../../example/hashmap/main.go).

## Validation

Behavior is covered by [hashmap_test.go](./hashmap_test.go).
