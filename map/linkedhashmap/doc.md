# linked hashmap

```go
import linkedhashmap "github.com/474420502/structure/map/linkedhashmap"
```

`map/linkedhashmap` combines hash lookup with a linked order of entries.

## Features

- arbitrary key and value types via `interface{}`
- stable traversal order from head to tail
- append/prepend insertion helpers
- update-and-move operations with `SetFront` and `SetBack`
- ordered snapshots through `Keys`, `Values`, and `Slices`

## API Snapshot

- `New() *LinkedHashmap`
- `Put(key interface{}, value interface{}) bool`
- `PushBack(key interface{}, value interface{}) bool`
- `PushFront(key interface{}, value interface{}) bool`
- `Set(key, value interface{}) bool`
- `SetBack(key interface{}, value interface{}) bool`
- `SetFront(key interface{}, value interface{}) bool`
- `Get(key interface{}) (interface{}, bool)`
- `Remove(key interface{}) (interface{}, bool)`
- `Keys() []interface{}`
- `Values() []interface{}`
- `Slices() []Slice`
- `Clear()`
- `Empty() bool`
- `Size() uint`
- `String() string`

## Notes

- `Put` is an alias for `PushBack` and does not overwrite existing entries.
- `Set` only updates existing entries and returns `false` when the key is absent.
- `SetFront` and `SetBack` both update insertion order as part of the write.
- Example usage is available at [../../example/linkedhashmap/main.go](../../example/linkedhashmap/main.go).

## Validation

Behavior is covered by [linked_hashmap_test.go](./linked_hashmap_test.go).