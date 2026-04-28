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
- Sections 2, 3, and 7 come from the root `BenchmarkRotationCompare*` suite and compare `IndexTree` vs `AVL` with rotation statistics
- Sections 1, 4, 5, 6, and 8 come from the package-level comparison benchmarks under `tree/skiplist` and `tree/indextree`
- In the root rotation suite, `10k/20k/50k` labels are `prepSize` inputs. For `PutSequential`, they only shift the key range and do not represent different final tree sizes

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

### 2. Random Put (prepSize 50k, with rotation stats)

| Structure | Time/op | Rotations/op | Double Rotations/op | Tree Height | AvgDepth |
|-----------|---------|--------------|---------------------|-------------|----------|
| **IndexTree** | ~690 ns | 0.47 | 0.23 | 26 | 20.53 |
| **AVL** | ~710 ns | 3.07 | 1.00 | 26 | 20.90 |

Winner: IndexTree

### 3. Sequential Put Rotation Comparison (monotone insert stream)

| Structure | Time/op | Rotations/op | Double Rotations/op | Height |
|-----------|---------|--------------|---------------------|--------|
| **IndexTree** | ~123 ns | 1.00 | 0 | 24 |
| **AVL** | ~156 ns | 3.49 | 1.94 | 24 |

Winner: IndexTree

### 4. Random Get (100k fixed keys)

| Structure | Time/op | Memory/op | Allocs/op |
|-----------|---------|-----------|-----------|
| **IndexTree** | ~107 ns | 0 B | 0 |
| **TreeList** | ~128 ns | 8 B | 1 |
| **AVL** | ~140 ns | 8 B | 1 |
| SkipList | ~430 ns | 8 B | 1 |

Winner: IndexTree

Note: `IndexTree` returns `interface{}` directly in `Get`, which avoids one boxing allocation seen in the other tree implementations under this benchmark setup.

### 5. SeekGE Operation (100k keys)

| Structure | Time/op | Memory/op | Allocs/op |
|-----------|---------|-----------|-----------|
| **TreeList** | ~155 ns | 7 B | 0 |
| **AVL** | ~275 ns | 328 B | 1-2 |
| SkipList | ~715 ns | 7 B | 0 |

Note: `IndexTree` historically used `Traverse()` callbacks for iteration. Iterator APIs were added later, but the original benchmark section remains centered on the comparable `Seek*` interfaces.

### 6. Index-based Access (fixed tree benchmarks)

| Structure | Time/op | Memory/op | Allocs/op |
|-----------|---------|-----------|-----------|
| **TreeList** | ~36 ns | 0 B | 0 |
| **IndexTree** | ~88 ns | 0 B | 0 |
| SkipList | ~100k ns | 16 B | 1 |

Winner: TreeList

Note: `TreeList` and `SkipList` use `BenchmarkTreeIndex` with 100k preloaded keys. `IndexTree` uses `BenchmarkIndexOnly` with 1M preloaded keys, so this row is still useful directionally but not a perfectly unified apples-to-apples benchmark.

### 7. Mixed Workload (round-robin Put/Get/Remove, prepSize 50k)

| Structure | Time/op | Rotations/op | Height |
|-----------|---------|--------------|--------|
| **IndexTree** | ~145 ns | 0.068 | 19 |
| **AVL** | ~177 ns | 0.83 | 19 |

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

# Shift tolerance experiment benchmarks
go test -v -run TestComprehensiveReport -count=1 ./tree/indextree/experiment/...
go test -bench=. -benchmem ./tree/indextree/experiment/...
```

## Related Documents

- [rotation-analysis.md](./rotation-analysis.md)
- [benchmark-comparison.zh.md](./benchmark-comparison.zh.md)
- [../rotation-analysis.md](../rotation-analysis.md)