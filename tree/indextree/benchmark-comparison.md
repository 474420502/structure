# IndexTree Benchmark Comparison

This document keeps only same-harness, apples-to-apples benchmark results for `tree/indextree`. The repository root README now stays focused on project-wide navigation.

## Structures Compared

| Structure | Type | Description |
|-----------|------|-------------|
| **IndexTree** | BST | Self-balancing BST using size-based balancing |
| **TreeList** | BST | Self-balancing BST with bidirectional linked-list pointers |
| **AVL** | BST | Classic AVL tree with height-based balancing |
| **RBTree** | BST | Classic red-black tree with weaker balancing and lower rotation pressure |
| **SkipList** | Skip List | Concurrent skip list with `RWMutex` |

## Methodology

- Results are kept only when they share the same benchmark function, dataset/seed, preload size, and public API semantics
- Summary numbers below use the median of `-count=3` on `AMD Ryzen 7 7700` with Go's default `-cpu=16`
- Shared operation comparisons come from `tree/skiplist/compare_bench_test.go`
- Rotation statistics come from the repository-root `BenchmarkRotationCompare*` suite; this summary keeps the representative `50k` case only
- In the root rotation suite, `10k/20k/50k` labels are `prepSize`. For `PutSequential`, they shift the starting key range and do not represent different final tree sizes
- `RBTree` does not currently expose rank/index APIs, so it is excluded from `Index`-style comparisons
- Older rows that mixed preload sizes, lacked a same-harness `IndexTree` implementation, or relied on previously invalid near-zero `ns/op` output have been removed from the summary

---

## Results Summary

### 1. Shared CRUD Benchmarks

| Operation | IndexTree | TreeList | AVL | RBTree | SkipList | Winner |
|-----------|-----------|----------|-----|--------|----------|--------|
| Random Put | `576.8 ns` / `64 B` / `1 alloc` | `636.7 ns` / `80 B` / `1 alloc` | `644.8 ns` / `48 B` / `1 alloc` | `540.6 ns` / `48 B` / `1 alloc` | `1.075 us` / `175 B` / `2 allocs` | **RBTree** |
| Sequential Put | `80.02 ns` / `64 B` / `1 alloc` | `88.02 ns` / `80 B` / `1 alloc` | `109.6 ns` / `48 B` / `1 alloc` | `93.44 ns` / `48 B` / `1 alloc` | `817.4 ns` / `175 B` / `2 allocs` | **IndexTree** |
| Random Get | `115.4 ns` / `8 B` / `1 alloc` | `124.9 ns` / `8 B` / `1 alloc` | `137.0 ns` / `8 B` / `1 alloc` | `114.7 ns` / `8 B` / `1 alloc` | `411.8 ns` / `8 B` / `1 alloc` | **RBTree ~= IndexTree** |
| Random Remove | `521.8 ns` / `0 B` / `0 alloc` | `586.9 ns` / `16 B` / `1 alloc` | `696.9 ns` / `3 B` / `0 alloc` | `560.1 ns` / `0 B` / `0 alloc` | `1.011 us` / `16 B` / `1 alloc` | **IndexTree** |

### 2. Ordered Access Benchmarks

| Operation | IndexTree | TreeList | AVL | RBTree | SkipList | Winner |
|-----------|-----------|----------|-----|--------|----------|--------|
| Iterator Traversal | `1.823 ms` / `800002 B` / `100000 allocs` | `1.229 ms` / `800000 B` / `100000 allocs` | `1.990 ms` / `800323 B` / `100001 allocs` | `1.828 ms` / `800003 B` / `100000 allocs` | `4.765 ms` / `800000 B` / `100000 allocs` | **TreeList** |
| SeekGE | `130.9 ns` / `7 B` / `0 alloc` | `124.5 ns` / `7 B` / `0 alloc` | `239.7 ns` / `328 B` / `1 alloc` | `130.1 ns` / `7 B` / `0 alloc` | `411.8 ns` / `7 B` / `0 alloc` | **TreeList ~= IndexTree ~= RBTree** |
| Index (100k preloaded keys) | `34.99 ns` / `0 B` / `0 alloc` | `36.37 ns` / `0 B` / `0 alloc` | N/A | N/A | `103.8 us` / `16 B` / `1 alloc` | **IndexTree ~= TreeList** |

Note: the current `SeekGE` benchmark measures the full public path of `Iterator()` + `SeekGE()` + `Next()`, not only the raw search step. After replacing the old stack-height setup with parent-pointer traversal, `IndexTree` no longer pays an iterator-construction penalty on this path.

Note: the current `Index` benchmark now drains both key and value into typed sinks, so the genericized `IndexTree` path is no longer penalized by interface boxing during result collection.

### 3. Rotation and Shape Profile (Representative 50k Case)

| Workload | IndexTree | AVL | RBTree | Takeaway |
|----------|-----------|-----|--------|----------|
| Random Put | `645.2 ns`, `0.4746 rot/op`, `0.2262 double/op`, `height=26`, `avgDepth=20.63` | `702.6 ns`, `3.065 rot/op`, `0.9963 double/op`, `height=26`, `avgDepth=20.94` | `582.9 ns`, `0.5820 rot/op`, `0 double/op`, `height=27`, `avgDepth=20.98` | `RBTree` remains the random-write baseline; after genericizing values, `IndexTree` now matches the one-allocation profile and the remaining gap is mostly balance-maintenance cost |
| Sequential Put | `105.6 ns`, `1.000 rot/op`, `0 double/op`, `height=24` | `148.8 ns`, `3.487 rot/op`, `1.944 double/op`, `height=24` | `116.9 ns`, `1.000 rot/op`, `0 double/op`, `height=45` | After genericizing values, `IndexTree` now beats `RBTree` on sequential writes; the `45` here still comes from the root suite's larger final `b.N` |
| Random Get | `89.96 ns`, `0 rot/op`, `height=20`, `avgDepth=14.92` | `105.6 ns`, `0 rot/op`, `height=19`, `avgDepth=15.18` | `87.03 ns`, `0 rot/op`, `height=19`, `avgDepth=14.96` | `RBTree` is marginally fastest on point lookup; `IndexTree` stays very close while still carrying rank metadata |
| Random Remove | `101.5 ns`, `0.4445 rot/op`, `0.2130 double/op`, `height=18`, `avgDepth=13.19` | `161.1 ns`, `5.076 rot/op`, `1.570 double/op`, `height=18`, `avgDepth=13.72` | `95.01 ns`, `0.9443 rot/op`, `0 double/op`, `height=14`, `avgDepth=10.04` | `RBTree` removes fastest, but `IndexTree` stays much closer to it than AVL does while doing far fewer rotations than AVL |
| Mixed Workload | `132.9 ns`, `0.04997 rot/op`, `0.02218 double/op`, `height=19`, `avgDepth=14.91` | `175.1 ns`, `0.8275 rot/op`, `0.2499 double/op`, `height=19`, `avgDepth=15.10` | `133.2 ns`, `0.1317 rot/op`, `0 double/op`, `height=19`, `avgDepth=14.98` | `IndexTree` and `RBTree` are effectively tied on throughput here; `IndexTree` still rebalances less |

---

## Interpretation

- `RBTree` remains the strongest random-write baseline in this repository: it wins random puts, is marginally ahead on point gets, and now sits in the same `SeekGE` tier as `IndexTree` and `TreeList`
- `IndexTree` now has four clearer strengths: best sequential `Put`, best random `Remove`, best fair `Index` access, and a neighbor-seek path that is now effectively tied with `RBTree`
- `TreeList` is still the best choice for ordered traversal; its neighbor-seek lead is now small rather than structural
- `RBTree` does not need non-standard handling for monotone inserts; a direct sequential-insert regression test still keeps it within the theoretical red-black height bound
- `AVL` still gives a classic aggressively balanced shape, but it remains the most rotation-heavy tree in write-intensive workloads
- `SkipList` still matters as the concurrency-safe option, not as the lowest-latency single-threaded structure

## Running The Benchmarks

```bash
go test -run '^$' -bench 'BenchmarkTree(Put|PutSequential|Get|Remove|Iterator|SeekGE|Index)$' -benchmem -count=3 ./tree/skiplist
go test . -run '^$' -bench 'BenchmarkRotationCompare(PutRandom50k|PutSequential50k|GetRandom50k|RemoveRandom50k|Mixed50k)$' -benchmem -count=3
```

## Related Documents

- [rotation-analysis.md](./rotation-analysis.md)
- [benchmark-comparison.zh.md](./benchmark-comparison.zh.md)
- [../rotation-analysis.md](../rotation-analysis.md)