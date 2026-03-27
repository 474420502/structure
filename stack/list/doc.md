# stack list

```go
import liststack "github.com/474420502/structure/stack/list"
```

`stack/list` provides a generic LIFO stack backed by linked nodes.

## Features

- generic `T` values
- constant-time push, peek, and pop
- linked representation with no slice growth copies
- snapshot export through `Values`

## API Snapshot

- `New[T any]() *Stack[T]`
- `Push(v T)`
- `Peek() (T, bool)`
- `Pop() (T, bool)`
- `Clear()`
- `Empty() bool`
- `Size() uint`
- `Values() []T`
- `String() string`

## Notes

- This variant trades slice locality for node-based storage and steady push/pop behavior.
- `String` renders the stack from top toward bottom.
- Example usage is available at [../../example/liststack/main.go](../../example/liststack/main.go).

## Validation

Behavior is covered by [stack_test.go](./stack_test.go).