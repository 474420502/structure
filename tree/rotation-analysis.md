# Rotation Analysis Notes

Date: 2026-03-26

Scope:
- `tree/itree`
- `tree/indextree`
- `rotation_compare_bench_test.go`

This document records the investigation, benchmark design, code changes, and conclusions from the rotation strategy analysis carried out in this repository.

## Background

The original discussion started from `itree`'s rotation type decision logic in `sizeRotateType`.

Original code path:
- `tree/itree/tree_hide.go`
- `tree/queue/priority/tree_hide.go` has a structurally similar implementation

Original decision logic:

```go
sdiff1 := ls - rs - subsize - subsize - 2
sdiff2 := ls - rs - subsubsize - subsubsize - 2
if sdiff1 > sdiff2 {
    return true  // double rotation
} else {
    return false // single rotation
}
```

The first question was whether this logic was too vague because it used subtree sizes instead of real subtree heights.

The first proposed replacement was a height-based rule using recursive `getHeight`. That version was rejected for implementation because it would turn each rotation decision into a subtree traversal, which is too expensive for a hot path.

## What Was Confirmed About the Existing Logic

After inspecting `itree`, the conclusion was:

1. `itree` is not an AVL tree that maintains exact heights.
2. Its balancing model is already size-driven, not height-driven.
3. `getMaybeHeight(size)` is only a size-derived approximation used as part of the rebalance threshold.
4. `sizeRotateType` is not purely random or arbitrary. It is an attempt to choose the rotation that gives a better size distribution after rotation.

That means the old logic is not "wrong" in principle, but it is hard to read, hard to justify locally, and hard to compare experimentally.

## Heuristic Introduced for Comparison

To test a simpler rule without switching to recursive height computation, a second heuristic was added to `itree`:

```go
func sizeRotateTypeChildBias[KEY, VALUE any](cur *Node[KEY, VALUE], child int, ls, rs int) bool {
    rchild := ^child + 2
    sub := cur.Children[child]
    if sub == nil {
        return false
    }

    outerSize := getSize(sub.Children[child])
    innerSize := getSize(sub.Children[rchild])
    return innerSize > outerSize
}
```

Interpretation:
- If the heavy child is biased toward the inner side, use double rotation.
- Otherwise, use single rotation.

This rule is cheaper to explain and similar in spirit to the classic "outer-heavy vs inner-heavy" decision, while still staying entirely size-based.

## Instrumentation Added During Research

To compare the strategies and the two tree implementations, several benchmarking hooks were added.

### itree instrumentation

Added to `tree/itree`:
- configurable `rotateType` function on `Tree`
- single rotation counter
- double rotation counter
- benchmark-only exported stats wrapper
- shape metrics:
  - `height`
  - `avg-depth`
  - `p50-depth`
  - `p95-depth`

Files involved:
- `tree/itree/tree.go`
- `tree/itree/tree_hide.go`
- `tree/itree/debug_utils.go`
- `tree/itree/rotate_bench_test.go`
- `tree/itree/bench_stats_export.go`

### indextree instrumentation

Added to `tree/indextree`:
- single rotation counter
- double rotation counter
- benchmark stats export
- shape metrics aligned with `itree`

Files involved:
- `tree/indextree/tree.go`
- `tree/indextree/utils.go`
- `tree/indextree/bench_stats.go`

### cross-tree comparison benchmark

A repository-root benchmark file was added:
- `rotation_compare_bench_test.go`

This benchmark runs `indextree` and `itree` under the same workloads and reports both throughput and structure/rotation metrics.

## Phase 1: A/B Benchmark Inside itree

The first benchmark phase compared two `itree` strategies only:

1. `formula`
   - original `sizeRotateType`
2. `child-bias`
   - `sizeRotateTypeChildBias`

Benchmarks added:
- `BenchmarkRotateDecisionPut`
- `BenchmarkRotateDecisionRemove`
- `BenchmarkRotateDecisionMixed`

Representative command:

```bash
go test ./tree/itree -run '^$' -bench 'BenchmarkRotateDecision(Put|Remove|Mixed)$' -benchmem -count=8
```

### A/B result summary

Across repeated runs:

- `child-bias` consistently reduced:
  - `rot/op`
  - `double/op`
- tree shape remained very close to the original formula
- throughput did not regress in any catastrophic way
- `Put` and `Mixed` were competitive enough to justify a rollout test

Important observation:
- fewer rotations did not always directly produce a better depth profile
- the shape metrics stayed very close between the two heuristics

### Decision taken after A/B

`itree`'s default strategy was switched from the original formula to `child-bias`.

Default path changed in:
- `tree/itree/tree.go`
- `tree/itree/tree_hide.go`

Validation run after the switch:
- full `itree` package tests passed
- repeated stress tests passed
- repeated whole-package runs passed

Representative commands:

```bash
go test ./tree/itree
go test ./tree/itree -count=10
go test ./tree/itree -run '^(TestForce|TestSimpleForce|TestIteratorForce|TestIteratorForce2|TestRemoveRangeIndexForce|TestRankForce|TestTrim|TestRemoveRange|TestSplit)$' -count=20
```

Conclusion of phase 1:
- the `child-bias` heuristic was acceptable as `itree`'s default
- it reduced rotations and passed the available correctness checks

## Phase 2: Direct indextree vs itree Comparison

The next phase compared the two implementations directly.

Benchmarks added at repository root:
- `BenchmarkRotationComparePutRandom`
- `BenchmarkRotationComparePutSequential`
- `BenchmarkRotationCompareRemoveRandom`
- `BenchmarkRotationCompareMixed`

Representative command:

```bash
go test . -run '^$' -bench 'BenchmarkRotationCompare(PutRandom|PutSequential|RemoveRandom|Mixed)$' -benchmem -count=5
```

## Cross-Tree Findings

### 1. PutRandom

Representative metric profile:

- `indextree`
  - about `820` to `836 ns/op`
  - `rot/op ~= 0.4736`
  - `double/op ~= 0.2263`
  - `avg-depth ~= 20.1`
  - `height ~= 25`
- `itree`
  - about `874` to `917 ns/op`
  - `rot/op ~= 0.4345`
  - `double/op ~= 0.2171`
  - `avg-depth ~= 20.0`
  - `height ~= 26`

Interpretation:
- `itree` rotates slightly less than `indextree` in this workload.
- Despite that, `itree` is still slower.
- Therefore the performance gap here is not caused by extra rotations alone.
- This strongly suggests additional fixed overhead in `itree`'s insert and rebalance path.

### 2. PutSequential

Representative metric profile:

- `indextree`
  - about `124` to `142 ns/op`
  - `rot/op = 1.0`
  - `double/op = 0`
  - `height ~= 24`
- `itree`
  - about `233` to `245 ns/op`
  - `rot/op = 1.0`
  - `double/op = 0`
  - `height ~= 23`

Interpretation:
- Both implementations perform essentially one single rotation per operation.
- `itree` still costs roughly twice as much per operation.
- This means the dominant issue in this scenario is not the rotation decision rule.
- The dominant issue is the constant cost around rotation and rebalancing.

This was one of the strongest findings of the investigation.

### 3. RemoveRandom

Representative metric profile:

- `indextree`
  - about `53` to `55 ns/op`
  - `rot/op ~= 0.4353`
  - `double/op ~= 0.2110`
  - `avg-depth ~= 6.6` to `8.3`
  - `height ~= 9` to `11`
- `itree`
  - about `100` to `106 ns/op`
  - `rot/op ~= 0.6693`
  - `double/op ~= 0.3256`
  - `avg-depth ~= 7.1` to `8.5`
  - `height ~= 9` to `12`

Interpretation:
- Here the gap is largely explained by rotation volume.
- `itree` performs materially more rotations and more double rotations than `indextree`.
- Tree shape is not dramatically worse, so the extra work is mostly from maintenance activity, not from a much taller tree.

### 4. Mixed workload

Representative metric profile:

- `indextree`
  - about `109` to `111 ns/op`
  - `rot/op ~= 0.0677`
  - `double/op ~= 0.0316`
  - `avg-depth ~= 12.5`
  - `height ~= 16`
- `itree`
  - about `140` to `142 ns/op`
  - `rot/op ~= 0.1044`
  - `double/op ~= 0.0482`
  - `avg-depth ~= 12.6`
  - `height ~= 17`

Interpretation:
- `itree` again performs more rotations.
- `itree` is also slightly taller in this workload.
- The throughput gap here is explained by both more rebalancing work and slightly worse final shape.

## Consolidated Root Cause Model

The final model from the benchmark evidence is:

### Cause A: `itree` rotates too often in delete-heavy and mixed workloads

Evidence:
- `RemoveRandom`: significantly higher `rot/op` and `double/op`
- `Mixed`: significantly higher `rot/op` and `double/op`

Likely affected code:
- `tree/itree/tree_hide.go`
- especially `rebalance`
- and the threshold logic around `getMaybeHeight` and size comparisons

### Cause B: `itree` has higher fixed cost per rebalance even when rotation counts match

Evidence:
- `PutSequential` has the same `rot/op` and `double/op` as `indextree`
- but `itree` is still much slower

Likely contributors:
- recursive insertion path in `tree/itree/tree_hide.go`
- repeated rebalance checks along the path
- generic function dispatch and extra field updates
- node/parent navigation differences compared with `indextree`
- the overall amount of bookkeeping around each rotation

### Cause C: lowering rotations alone is not enough to make `itree` catch `indextree`

Evidence:
- `PutRandom` shows `itree` rotating less but still running slower

This is an important guardrail for future optimization work.

## Why child-bias Still Made Sense for itree

Even though `indextree` remains faster overall, switching `itree` to `child-bias` still made sense because:

1. the new heuristic is easier to reason about
2. it reduced rotation counts inside `itree`
3. it passed correctness tests and repeated stress runs
4. it did not make the benchmark picture worse in a meaningful way

So the change improved the local behavior of `itree`, even though it did not close the broader gap with `indextree`.

## Commands Used During Investigation

Representative commands used in this investigation:

```bash
go test ./tree/itree -run '^$' -bench 'BenchmarkRotateDecision(Put|Remove|Mixed)$' -benchmem -count=1
go test ./tree/itree -run '^$' -bench 'BenchmarkRotateDecision(Put|Remove|Mixed)$' -benchmem -count=8
go test ./tree/itree -count=10
go test ./tree/itree -run '^(TestForce|TestSimpleForce|TestIteratorForce|TestIteratorForce2|TestRemoveRangeIndexForce|TestRankForce|TestTrim|TestRemoveRange|TestSplit)$' -count=20

go test . -run '^$' -bench 'BenchmarkRotationCompare(PutRandom|PutSequential|RemoveRandom|Mixed)$' -benchmem -count=1
go test . -run '^$' -bench 'BenchmarkRotationCompare(PutRandom|PutSequential|RemoveRandom|Mixed)$' -benchmem -count=3
go test . -run '^$' -bench 'BenchmarkRotationCompare(PutRandom|PutSequential|RemoveRandom|Mixed)$' -benchmem -count=5
```

## Current Status at the End of This Analysis

At the end of this investigation:

- `itree` default rotate strategy is `sizeRotateTypeChildBias`
- `itree` has internal benchmark instrumentation for rotation and shape stats
- `indextree` has matching benchmark instrumentation
- root-level cross-tree benchmarks exist for direct comparison

## Recommended Next Steps

If more work is planned, the evidence suggests this priority order:

1. optimize `itree` fixed rebalance overhead
   - best target for `PutSequential`
   - likely more impactful than only tuning rotate selection again

2. revisit `itree` rebalance thresholds
   - best target for `RemoveRandom` and `Mixed`
   - aim to reduce `rot/op` closer to `indextree`

3. keep the cross-tree benchmark file as the primary regression harness
   - it now captures both throughput and structural behavior

4. only after that, consider whether similar heuristic changes should be applied to other size-balanced trees in the repository
   - for example `queue/priority` if its implementation remains structurally aligned with the older `itree` logic

## Short Conclusion

The most important final conclusion is:

- `itree` was improved locally by replacing the original rotation selector with `child-bias`
- but `itree` is still slower than `indextree`
- the remaining gap is explained by two different issues:
  - extra rotations in delete-heavy and mixed workloads
  - higher constant overhead even when rotation counts are the same

That means the next optimization step should focus less on rotation type selection alone and more on the full rebalance cost model.