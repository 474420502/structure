# hashset

```go
import "github.com/474420502/structure/set/hashset"
```

`set/hashset` is a simple unordered set built on top of Go's native map.

## Features

- variadic `Add` and `Remove`
- membership check with `Contains`
- `Values` export for all items
- `Empty`, `Clear`, `Size`, and `String`

## API Snapshot

- `New() *HashSet`
- `Add(items ...interface{})`
- `Remove(items ...interface{})`
- `Contains(item interface{}) bool`
- `Values() []interface{}`
- `Empty() bool`
- `Clear()`
- `Size() int`
- `String() string`

## Notes

- The set is not generic yet; all values are stored as `interface{}`.
- Iteration order is undefined because it follows Go map iteration semantics.

## Validation

Behavior is covered by [hashset_test.go](./hashset_test.go).