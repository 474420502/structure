# stack listarray

```go
import lastack "github.com/474420502/structure/stack/listarray"
```

`stack/listarray` provides a generic LIFO stack built from linked array blocks.

## Features

- generic `T` values
- amortized constant-time push and pop
- optional initial capacity with `NewWithCap`
- lower reallocation pressure than a single contiguous slice under some workloads
- snapshot export through `Values`

## API Snapshot

- `New[T any]() *Stack[T]`
- `NewWithCap[T any](cap int) *Stack[T]`
- `Push(v T)`
- `Peek() (T, bool)`
- `Pop() (T, bool)`
- `Clear()`
- `Empty() bool`
- `Size() uint`
- `Values() []T`
- `String() string`

## Notes

- This implementation is a hybrid between linked storage and array storage.
- `NewWithCap` controls the initial block size and can reduce early allocations.
- Example usage is available at [../../example/lastack/main.go](../../example/lastack/main.go).

## Validation

Behavior is covered by [stack_test.go](./stack_test.go).