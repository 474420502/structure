# hashmap

```go
import "github.com/474420502/structure/map/hashmap"
```

`map/hashmap` wraps Go's hash map with a small container-style API.

## Features

- arbitrary key and value types via `interface{}`
- optional initial capacity with `NewWithCap`
- insert-only `Put` and overwrite `Set`
- snapshot access to keys, values, and key/value slices

## API Snapshot

- `New() *HashMap`
- `NewWithCap(cap int) *HashMap`
- `Put(key interface{}, value interface{}) bool`
- `Set(key interface{}, value interface{})`
- `Get(key interface{}) (interface{}, bool)`
- `Remove(key interface{})`
- `Keys() []interface{}`
- `Values() []interface{}`
- `Slices() []Slice`
- `Clear()`
- `Empty() bool`
- `Size() int`
- `String() string`

## Notes

- `Put` returns `false` when the key already exists and leaves the stored value unchanged.
- `Set` always overwrites or creates the entry.
- Iteration order for `Keys`, `Values`, and `Slices` is not stable.
- Example usage is available at [../../example/hashmap/main.go](../../example/hashmap/main.go).

## Validation

Behavior is covered by [hashmap_test.go](./hashmap_test.go).