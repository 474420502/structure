# compare

```go
import "github.com/474420502/structure/compare"
```

`compare` provides the reusable key comparators used by every ordered container in this repository.

## Comparator Contract

All ordered containers share one contract:

- negative: `k1 < k2`
- positive: `k1 > k2`
- zero: `k1 == k2`

Functions: `Any`, `AnyDesc`, `ArrayAny`, `ArrayLenAny`, `RunesDesc`, `RunesLenDesc`.

Every tree, set, map, priority queue, heap, and list accepts this contract, so the same comparator can be passed to any of them.

### Deprecated aliases

`AnyEx` and `ArrayAnyEx` are kept for source compatibility and now delegate to `Any` and `ArrayAny`. They previously used a legacy index-style encoding (`k1 < k2` returned `1`, `k1 > k2` returned `0`, `k1 == k2` returned `-1`) that the containers no longer consume.

- `AnyEx` — **Deprecated: use `Any`.**
- `ArrayAnyEx` — **Deprecated: use `ArrayAny`.**

## API Snapshot

- `type Compare[T any] func(k1, k2 T) int`
- `type DefaultAny interface{ ... }`
- `type ArrayType interface{ []byte | string }`
- `Any[T DefaultAny](k1, k2 T) int`
- `AnyDesc[T DefaultAny](k1, k2 T) int`
- `AnyEx[T DefaultAny](k1, k2 T) int` (deprecated)
- `ArrayAny[T ArrayType](k1, k2 T) int`
- `ArrayLenAny[T ArrayType](k1, k2 T) int`
- `ArrayAnyEx[T ArrayType](k1, k2 T) int` (deprecated)
- `RunesDesc(k2, k1 interface{}) int`
- `RunesLenDesc(k2, k1 interface{}) int`

## Notes

- Prefer `Any` for scalar keys and `ArrayAny` for `[]byte` / `string` keys in new code.
- A custom comparator must return a negative, zero, or positive value with the same meaning; otherwise tree ordering is undefined.
