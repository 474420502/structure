# hashset

```go
import "github.com/474420502/structure/set/hashset"
```

`set/hashset` is a generic unordered set built on top of Go's native map.

## Features

- generic `T comparable` items (no interface boxing)
- variadic `Add` and `Remove`
- membership check with `Contains`
- `Values` export for all items
- `Empty`, `Clear`, `Size`, `Len`, and `String`

## API Snapshot

- `New[T comparable]() *HashSet[T]`
- `Add(items ...T)`
- `Remove(items ...T)`
- `Contains(item T) bool`
- `Values() []T`
- `Empty() bool`
- `Clear()`
- `Size() int`
- `Len() int`
- `String() string`

## Notes

- Iteration order is undefined because it follows Go map iteration semantics.

## Validation

Behavior is covered by [hashset_test.go](./hashset_test.go) and [hashset_extra_test.go](./hashset_extra_test.go).
