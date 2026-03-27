# IndexTree Benchmark Comparison

This document is the dedicated benchmark comparison for `tree/indextree`. The repository root README now focuses on the full project index.

## Structures Compared

| Structure | Type | Description |
|-----------|------|-------------|
| **IndexTree** | BST | Self-balancing BST using size-based balancing |
| **TreeList** | BST | Self-balancing BST with bidirectional linked list pointers |
| **AVL** | BST | Classic AVL tree with height-based balancing |
| **SkipList** | Skip List | Concurrent skip list with `RWMutex` |

## Benchmark Methodology

- Fair comparison: identical test data and random seeds
- Comparable APIs: `Put`, `Get`, `Remove`, iterator-style traversal
- Stability: repeated runs with `-count`
- Workloads: random, sequential, index-based, and mixed read/write patterns

---

## Benchmark Results Summary

### 1. Sequential Put (100k keys)

| Structure | Time/op | Memory/op | Allocs/op |
|-----------|---------|-----------|-----------|
| **TreeList** | ~92 ns | 80 B | 1 |
| **IndexTree** | ~94 ns | 71 B | 1 |
| **AVL** | ~112 ns | 48 B | 1 |
| SkipList | ~860 ns | 175 B | 2 |

Winner: TreeList

### 2. Random Put (50k keys, with rotation stats)

| Structure | Time/op | Rotations/op | Double Rotations/op | Tree Height | Memory/op |
|-----------|---------|--------------|---------------------|-------------|-----------|
| **IndexTree** | ~730 ns | 0.47 | 0.23 | 26 | 161 B |
| **AVL** | ~778 ns | 3.07 | 1.00 | 26 | 137 B |

Winner: IndexTree

### 3. Sequential Put Rotation Comparison (50k keys)

| Structure | Time/op | Rotations/op | Double Rotations/op | Height |
|-----------|---------|--------------|---------------------|--------|
| **IndexTree** | ~125 ns | 1.00 | 0 | 24 |
| **AVL** | ~152 ns | 3.49 | 1.94 | 24 |

Winner: IndexTree

### 4. Random Get (100k keys)

| Structure | Time/op | Memory/op | Allocs/op |
|-----------|---------|-----------|-----------|
| **IndexTree** | ~135 ns | 0 B | 0 |
| **TreeList** | ~141 ns | 8 B | 1 |
| **AVL** | ~156 ns | 8 B | 1 |
| SkipList | ~650 ns | 8 B | 1 |

Winner: IndexTree

Note: `IndexTree` returns `interface{}` directly in `Get`, which avoids one boxing allocation seen in the other tree implementations under this benchmark setup.

### 5. SeekGE Operation (100k keys)

| Structure | Time/op | Memory/op | Allocs/op |
|-----------|---------|-----------|-----------|
| **TreeList** | ~155 ns | 7 B | 0 |
| **AVL** | ~275 ns | 328 B | 1-2 |
| SkipList | ~715 ns | 7 B | 0 |

Note: `IndexTree` historically used `Traverse()` callbacks for iteration. Iterator APIs were added later, but the original benchmark section remains centered on the comparable `Seek*` interfaces.

### 6. Index-based Access (100k keys)

| Structure | Time/op | Memory/op | Allocs/op |
|-----------|---------|-----------|-----------|
| **TreeList** | ~41 ns | 0 B | 0 |
| **IndexTree** | ~95 ns | 0 B | 0 |
| SkipList | ~102k ns | 16 B | 1 |

Winner: TreeList

### 7. Mixed Workload (50% Put, 25% Get, 25% Remove)

| Structure | Time/op | Rotations/op | Height |
|-----------|---------|--------------|--------|
| **IndexTree** | ~140 ns | 0.068 | 19 |
| **AVL** | ~182 ns | 0.83 | 19 |

Winner: IndexTree

### 8. Iterator Traversal (100k keys)

| Structure | Time/op | Memory/op | Allocs/op |
|-----------|---------|-----------|-----------|
| **TreeList** | ~1.4-1.5M ns | 800 KB | 100k |
| **AVL** | ~2.2-2.5M ns | 800 KB | 100k |
| SkipList | ~13-15M ns | 800 KB | 100k |

Winner: TreeList

---

## Interpretation

### IndexTree vs AVL

`IndexTree` and `AVL` reach similar tree heights, but `IndexTree` does less structural work to get there.

- 3.5x to 6.5x fewer rotations in insertion-heavy workloads
- 16% to 30% better write-heavy throughput in the summarized runs
- better mixed-workload behavior with fewer rebalance operations

The important observation is that more aggressive AVL rotations do not translate into a meaningfully better final tree shape in these benchmarks.

### TreeList vs SkipList

`TreeList` is consistently faster than `SkipList` in single-threaded ordered workloads.

- 6x to 9x faster on sequential puts
- 4x to 5x faster on random gets and `SeekGE`
- dramatically faster for index-based access

`SkipList` remains the concurrent option in this repository.

### When To Choose Which

- Choose `IndexTree` for single-threaded ordered workloads that need fast writes and rank/index operations.
- Choose `TreeList` when ordered iteration, head/tail access, and index access matter most.
- Choose `SkipList` when thread safety is required.
- Choose `AVL` when you want a conventional, easier-to-explain balanced BST.

---

## Running The Benchmarks

```bash
go test -bench=. -benchmem ./tree/...
go test -bench=BenchmarkTreePut -benchmem -count=5 ./tree/skiplist/...
go test -bench=BenchmarkRotation -benchmem -count=3 ./...
go test -bench=. -benchmem ./tree/indextree/...
go test -bench=. -benchmem ./tree/avl/...
go test -bench=. -benchmem ./tree/skiplist/...
go test -bench=. -benchmem ./tree/treelist/...
```

## Related Documents

- [rotation-analysis.md](./rotation-analysis.md)
- [benchmark-comparison.zh.md](./benchmark-comparison.zh.md)
- [../rotation-analysis.md](../rotation-analysis.md)