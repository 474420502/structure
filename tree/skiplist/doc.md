# skiplist

```go
import "github.com/474420502/structure/tree/skiplist"
```

`tree/skiplist` is the repository's concurrent ordered map implementation.

## Features

- generic `KEY` and `VALUE`
- internal `sync.RWMutex` protection
- ordered insertion, update, removal, and lookup
- head/tail access
- iterator with `Seek*` helpers and index tracking
- rank-style helpers such as `Index` and `IndexOf`
- range trimming and removal by index
- set-style operations: intersection, union, difference
- configurable maximum level via `NewWithMaxLevel`

## API Snapshot

- `New[KEY, VALUE any](comp compare.Compare[KEY]) *SkipList[KEY, VALUE]`
- `NewWithMaxLevel[KEY, VALUE any](level int, comp compare.Compare[KEY]) *SkipList[KEY, VALUE]`
- `Put`, `Set`, `PutDuplicate`
- `Get`, `Remove`, `RemoveHead`, `RemoveTail`, `RemoveIndex`, `RemoveRangeByIndex`
- `Head`, `Tail`, `Index`, `IndexOf`, `Size`, `Height`, `Clear`
- `Traverse`, `Slice`, `Slices`, `String`
- `Trim`, `TrimByIndex`, `Intersection`, `UnionSets`, `DifferenceSets`
- `Iterator` with `SeekToFirst`, `SeekToLast`, `SeekGE`, `SeekGT`, `SeekLE`, `SeekLT`, `SeekByIndex`, `Next`, `Prev`, `Clone`, `Index`

## Notes

- This is the only ordered structure in the repository explicitly designed for concurrent use.
- It offers a broader API surface than a minimal skip list, including rank and set-operation helpers.
- Benchmark comparisons with tree-based structures are documented under `tree/indextree/`.

## Validation

Behavior is covered by [skiplist_test.go](./skiplist_test.go) and [compare_bench_test.go](./compare_bench_test.go).