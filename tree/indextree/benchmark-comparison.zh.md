# IndexTree 性能对比文档

本文档只保留同口径、可直接横向比较的 benchmark 结果。仓库根目录 README 只负责项目总览与导航。

## 对比结构

| 结构 | 类型 | 说明 |
|------|------|------|
| **IndexTree** | BST | 基于子树大小平衡的自平衡二叉搜索树 |
| **TreeList** | BST | 带双向链表顺序指针的自平衡树 |
| **AVL** | BST | 经典 AVL，高度平衡 |
| **RBTree** | BST | 经典红黑树，弱平衡、低旋转开销 |
| **SkipList** | 跳表 | 基于 `RWMutex` 的并发跳表 |

## 口径说明

- 只保留满足以下条件的结果: 同一个 benchmark 函数、相同数据集和随机种子、相同预加载规模、相同公开 API 语义
- 主表数字取 `-count=3` 的中位结果，当前机器为 `AMD Ryzen 7 7700`，Go 默认 `-cpu=16`
- 共享操作对比来自 `tree/skiplist/compare_bench_test.go`
- 旋转统计来自仓库根目录 `BenchmarkRotationCompare*`，这里只保留最有代表性的 `50k` 组
- 在根目录旋转 benchmark 中，`10k/20k/50k` 是 `prepSize`。对 `PutSequential` 而言，它只影响 key 起始偏移，不代表不同的最终树规模
- `RBTree` 当前不提供 rank/index API，因此不会出现在 `Index` 这类需要顺序排名接口的对比行里
- 旧版文档里那些混用了不同预加载规模、缺少同口径实现、或曾出现近似 `0 ns/op` 的历史行，已经从摘要中移除

---

## 结果摘要

### 1. 共享 CRUD 对比

| 操作 | IndexTree | TreeList | AVL | RBTree | SkipList | 最优 |
|------|-----------|----------|-----|--------|----------|------|
| 随机 Put | `576.8 ns` / `64 B` / `1 alloc` | `636.7 ns` / `80 B` / `1 alloc` | `644.8 ns` / `48 B` / `1 alloc` | `540.6 ns` / `48 B` / `1 alloc` | `1.075 us` / `175 B` / `2 allocs` | **RBTree** |
| 顺序 Put | `80.02 ns` / `64 B` / `1 alloc` | `88.02 ns` / `80 B` / `1 alloc` | `109.6 ns` / `48 B` / `1 alloc` | `93.44 ns` / `48 B` / `1 alloc` | `817.4 ns` / `175 B` / `2 allocs` | **IndexTree** |
| 随机 Get | `115.4 ns` / `8 B` / `1 alloc` | `124.9 ns` / `8 B` / `1 alloc` | `137.0 ns` / `8 B` / `1 alloc` | `114.7 ns` / `8 B` / `1 alloc` | `411.8 ns` / `8 B` / `1 alloc` | **RBTree ~= IndexTree** |
| 随机 Remove | `521.8 ns` / `0 B` / `0 alloc` | `586.9 ns` / `16 B` / `1 alloc` | `696.9 ns` / `3 B` / `0 alloc` | `560.1 ns` / `0 B` / `0 alloc` | `1.011 us` / `16 B` / `1 alloc` | **IndexTree** |

### 2. 有序访问对比

| 操作 | IndexTree | TreeList | AVL | RBTree | SkipList | 最优 |
|------|-----------|----------|-----|--------|----------|------|
| 迭代遍历 | `1.823 ms` / `800002 B` / `100000 allocs` | `1.229 ms` / `800000 B` / `100000 allocs` | `1.990 ms` / `800323 B` / `100001 allocs` | `1.828 ms` / `800003 B` / `100000 allocs` | `4.765 ms` / `800000 B` / `100000 allocs` | **TreeList** |
| SeekGE | `130.9 ns` / `7 B` / `0 alloc` | `124.5 ns` / `7 B` / `0 alloc` | `239.7 ns` / `328 B` / `1 alloc` | `130.1 ns` / `7 B` / `0 alloc` | `411.8 ns` / `7 B` / `0 alloc` | **TreeList ~= IndexTree ~= RBTree** |
| Index（100k 预加载） | `34.99 ns` / `0 B` / `0 alloc` | `36.37 ns` / `0 B` / `0 alloc` | 不适用 | 不适用 | `103.8 us` / `16 B` / `1 alloc` | **IndexTree ~= TreeList** |

说明: 当前 `SeekGE` 基准测量的是“创建迭代器 + `SeekGE` + `Next`”这条完整公开 API 路径，而不是只测裸搜索步骤。改成基于 `Parent` 指针的前驱/后继迭代后，`IndexTree` 已经不再为这条路径支付额外的建迭代器开销。

说明: 当前 `Index` 基准会把 key/value 同时写入 typed sink，避免泛型化后的 `IndexTree` 因为结果收集阶段的 `interface{}` 装箱而被额外惩罚。

### 3. 旋转与树形统计（50k 代表组）

| 负载 | IndexTree | AVL | RBTree | 结论 |
|------|-----------|-----|--------|------|
| 随机 Put | `645.2 ns`，`0.4746 rot/op`，`0.2262 double/op`，`height=26`，`avgDepth=20.63` | `702.6 ns`，`3.065 rot/op`，`0.9963 double/op`，`height=26`，`avgDepth=20.94` | `582.9 ns`，`0.5820 rot/op`，`0 double/op`，`height=27`，`avgDepth=20.98` | `RBTree` 仍是随机写入基线；`IndexTree` 在泛型化后把分配压到了同阶，剩下差距主要来自平衡维护成本 |
| 顺序 Put | `105.6 ns`，`1.000 rot/op`，`0 double/op`，`height=24` | `148.8 ns`，`3.487 rot/op`，`1.944 double/op`，`height=24` | `116.9 ns`，`1.000 rot/op`，`0 double/op`，`height=45` | 泛型化后 `IndexTree` 顺序写入已反超 `RBTree`；这行里的 `45` 仍然来自 root suite 的更大最终 `b.N` |
| 随机 Get | `89.96 ns`，`0 rot/op`，`height=20`，`avgDepth=14.92` | `105.6 ns`，`0 rot/op`，`height=19`，`avgDepth=15.18` | `87.03 ns`，`0 rot/op`，`height=19`，`avgDepth=14.96` | `RBTree` 点查略快；`IndexTree` 非常接近，同时仍保留 rank 元数据 |
| 随机 Remove | `101.5 ns`，`0.4445 rot/op`，`0.2130 double/op`，`height=18`，`avgDepth=13.19` | `161.1 ns`，`5.076 rot/op`，`1.570 double/op`，`height=18`，`avgDepth=13.72` | `95.01 ns`，`0.9443 rot/op`，`0 double/op`，`height=14`，`avgDepth=10.04` | `RBTree` 删除最快，但 `IndexTree` 明显比 AVL 更接近它，同时写时重平衡仍远少于 AVL |
| 混合负载 | `132.9 ns`，`0.04997 rot/op`，`0.02218 double/op`，`height=19`，`avgDepth=14.91` | `175.1 ns`，`0.8275 rot/op`，`0.2499 double/op`，`height=19`，`avgDepth=15.10` | `133.2 ns`，`0.1317 rot/op`，`0 double/op`，`height=19`，`avgDepth=14.98` | `IndexTree` 和 `RBTree` 在吞吐上几乎持平；`IndexTree` 的重平衡次数仍更少 |

---

## 结果解读

- `RBTree` 仍然是最强的随机写入基线：随机 Put 最快，点查也略快，而 `SeekGE` 现在已经和 `IndexTree`、`TreeList` 处在同一梯队
- `IndexTree` 现在有四类更明确的优势：顺序 `Put` 最快、随机 `Remove` 最快、公平口径下的 `Index` 访问最强，而且邻近 `seek` 已经基本追平 `RBTree`
- `TreeList` 仍然是顺序访问王者，但它在邻近 `seek` 上的领先现在已经变成小幅优势，而不是结构性差距
- `RBTree` 不需要为了“顺序插入高度”额外偏离标准实现；直接顺序插入 `50k` 的回归测试仍满足红黑树高度上界
- `AVL` 依旧是最保守、最激进重平衡的一棵树；最终树高不差，但写入时付出的旋转成本最高
- `SkipList` 的核心价值仍然是并发安全，而不是单线程延迟

## 运行基准

```bash
go test -run '^$' -bench 'BenchmarkTree(Put|PutSequential|Get|Remove|Iterator|SeekGE|Index)$' -benchmem -count=3 ./tree/skiplist
go test . -run '^$' -bench 'BenchmarkRotationCompare(PutRandom50k|PutSequential50k|GetRandom50k|RemoveRandom50k|Mixed50k)$' -benchmem -count=3
```

## 相关文档

- [rotation-analysis.md](./rotation-analysis.md)
- [benchmark-comparison.md](./benchmark-comparison.md)
- [../rotation-analysis.md](../rotation-analysis.md)