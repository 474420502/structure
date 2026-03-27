# stack array

```go
import arraystack "github.com/474420502/structure/stack/array"
```

`stack/array` provides a generic LIFO stack backed by a contiguous slice.

## Features

- generic `T` values
- constant-time push, peek, and pop on the top of the stack
- snapshot export through `Values`
- compact implementation for simple stack workloads

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

- `Values` returns the current stack contents in storage order, from bottom to top.
- This is the simplest stack implementation in the repository when linked structure semantics are not needed.
- Example usage is available at [../../example/arraystack/main.go](../../example/arraystack/main.go).

## Validation

Behavior is covered by [stack_test.go](./stack_test.go).