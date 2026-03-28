# hashmap

```go
import "github.com/474420502/structure/map/hashmap"
```

`map/hashmap` wraps Go's hash map with a small container-style API.

## Features

- arbitrary key and value types via `interface{}`
- optional initial capacity with `NewWithCap`
- legacy `Put` and `Set` plus standardized semantic aliases
- snapshot access to keys, values, and key/value slices

## API Snapshot

- `New() *HashMap`
- `NewWithCap(cap int) *HashMap`
- `Put(key interface{}, value interface{}) bool`
- `InsertIfAbsent(key interface{}, value interface{}) bool`
- `Set(key interface{}, value interface{})`
- `Upsert(key interface{}, value interface{}) bool`
- `Get(key interface{}) (interface{}, bool)`
- `Remove(key interface{})`
- `Delete(key interface{}) (interface{}, bool)`
- `Keys() []interface{}`
- `Values() []interface{}`
- `Slices() []Slice`
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