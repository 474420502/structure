# structure

Generic data structures and search utilities for Go.

This repository contains multiple implementations of ordered trees, lists, stacks, queues, maps, sets, probabilistic filters, and graph search helpers. The root document is the project index. Performance comparison notes for `indextree` now live under `tree/indextree/`.

## Documentation Map

- Chinese overview: [readme_zh.md](./readme_zh.md)
- IndexTree benchmark summary: [tree/indextree/benchmark-comparison.md](./tree/indextree/benchmark-comparison.md)
- IndexTree benchmark summary in Chinese: [tree/indextree/benchmark-comparison.zh.md](./tree/indextree/benchmark-comparison.zh.md)
- Rotation analysis deep dive: [tree/indextree/rotation-analysis.md](./tree/indextree/rotation-analysis.md)
- Repository-level rotation notes: [tree/rotation-analysis.md](./tree/rotation-analysis.md)

## Module

```go
module github.com/474420502/structure
```

## Package Index

### Lists

| Package | Summary | Docs | Example |
|---------|---------|------|---------|
| `list/array_list` | Dynamic array list with random access, iterators, and circular iterators | [doc](./list/array_list/doc.md) | [example](./example/array_list/main.go) |
| `list/linked_list` | Doubly linked list with iterator and circular iterator support | [doc](./list/linked_list/doc.md) | [example](./example/linked_list/main.go) |

### Maps

| Package | Summary | Docs | Example |
|---------|---------|------|---------|
| `map/hashmap` | Basic hash map wrapper with `Put`, `Set`, `Get`, `Keys`, `Values`, `Slices` | [doc](./map/hashmap/doc.md) | [example](./example/hashmap/main.go) |
| `map/linkedhashmap` | Hash map that preserves insertion order and supports front/back relocation | [doc](./map/linkedhashmap/doc.md) | [example](./example/linkedhashmap/main.go) |
| `map/orderedmap.go` | Placeholder for a future ordered map implementation. The package currently exposes only an empty `OrderedMap` type. | [doc](./map/orderedmap.go/doc.md) | - |

### Queues

| Package | Summary | Docs | Example |
|---------|---------|------|---------|
| `queue/linkedarray` | Circular-array deque with front/back push-pop, indexed access, and traversal | [doc](./queue/linkedarray/doc.md) | - |
| `queue/list` | Doubly linked deque with explicit node handles for front/back iteration | [doc](./queue/list/doc.md) | - |
| `queue/priority` | Ordered priority queue backed by a size-balanced tree with iterators and indexed removal | [doc](./queue/priority/doc.md) | [example](./example/priority_queue/main.go) |

### Sets

| Package | Summary | Docs | Example |
|---------|---------|------|---------|
| `set/hashset` | Simple hash set for unordered membership tests | [doc](./set/hashset/doc.md) | - |
| `set/treeset` | Ordered tree set with AVL-style balancing and bidirectional iterators | [doc](./set/treeset/doc.md) | [example](./example/treeset/main.go) |

### Stacks

| Package | Summary | Docs | Example |
|---------|---------|------|---------|
| `stack/array` | Slice-backed LIFO stack | [doc](./stack/array/doc.md) | [example](./example/arraystack/main.go) |
| `stack/list` | Linked-list-backed LIFO stack | [doc](./stack/list/doc.md) | [example](./example/liststack/main.go) |
| `stack/listarray` | Segmented stack using linked array blocks | [doc](./stack/listarray/doc.md) | [example](./example/lastack/main.go) |

### Trees And Ordered Structures

| Package | Summary | Docs | Example |
|---------|---------|------|---------|
| `tree/avl` | Classic AVL tree with iterator support | [doc](./tree/avl/doc.md) | [example](./example/avl/main.go) |
| `tree/avls` | Duplicate-key AVL variant | [doc](./tree/avls/doc.md) | - |
| `tree/heap` | Binary heap based on comparator ordering | [doc](./tree/heap/doc.md) | [example](./example/heap/main.go) |
| `tree/indextree` | Size-balanced ordered tree with rank/index operations and split/trim support | [doc](./tree/indextree/doc.md) | [example](./example/indextree/main.go) |
| `tree/skiplist` | Concurrent skip list with iterator, index, trim, and set-operation helpers | [doc](./tree/skiplist/doc.md) | - |
| `tree/treelist` | Ordered map/tree list with linked ordering, range iterator, and set algebra | [doc](./tree/treelist/doc.md) | [example](./example/tree-treelist/main.go) |

### Filters

| Package | Summary | Docs | Example |
|---------|---------|------|---------|
| `filter/bloom` | Bloom filter with binary encode/decode helpers | [doc](./filter/bloom/doc.md) | [example](./example/bloom/main.go) |

### Graph Search

| Package | Summary | Docs | Example |
|---------|---------|------|---------|
| `graph/astar` | Grid-based A* search with pluggable neighbor, cost, and weight strategies | [doc](./graph/astar/doc.md) | - |

### Search Utilities

| Package | Summary | Docs | Example |
|---------|---------|------|---------|
| `search/treelist` | Byte-key ordered tree for search/index use cases | [doc](./search/treelist/doc.md) | [example](./example/search-treelist/main.go) |
| `search/searchtree` | Placeholder for higher-level search index abstractions. No usable public API yet. | [doc](./search/searchtree/doc.md) | - |

## Recommended Starting Points

- For a general ordered map with strong read/write performance: `tree/indextree`
- For ordered iteration plus head/tail access: `tree/treelist`
- For concurrent ordered access: `tree/skiplist`
- For a conventional self-balancing BST: `tree/avl`
- For a simple unordered dictionary or set: `map/hashmap`, `set/hashset`

## Benchmarks And Analysis

The benchmark comparison that used to live in the root README is now scoped to `indextree` and related ordered trees:

- [tree/indextree/benchmark-comparison.md](./tree/indextree/benchmark-comparison.md)
- [tree/indextree/benchmark-comparison.zh.md](./tree/indextree/benchmark-comparison.zh.md)
- [tree/indextree/rotation-analysis.md](./tree/indextree/rotation-analysis.md)
- [tree/rotation-analysis.md](./tree/rotation-analysis.md)

## Running Tests

```bash
go test ./...
```

## Running Tree Benchmarks

```bash
go test -bench=. -benchmem ./tree/...
go test -bench=BenchmarkRotation -benchmem -count=3 ./...
```
