# IndexTree 性能对比文档

本文档专门用于 `tree/indextree` 的基准测试与性能对比。仓库根目录 README 现在只负责整个项目的总览与导航。

## 对比结构

| 结构 | 类型 | 说明 |
|------|------|------|
| **IndexTree** | BST | 基于子树大小平衡的自平衡二叉搜索树 |
| **TreeList** | BST | 带双向链表顺序指针的自平衡树 |
| **AVL** | BST | 经典 AVL，高度平衡 |
| **SkipList** | 跳表 | 基于 `RWMutex` 的并发跳表 |

## 基准方法

- 公平对比: 使用相同测试数据和随机种子
- 接口可比: 覆盖 `Put`、`Get`、`Remove`、迭代遍历
- 多次运行: 通过 `-count` 提高稳定性
- 真实负载: 随机、顺序、按索引访问、混合读写

---

## 结果摘要

### 1. 顺序 Put（100k keys）

| 结构 | 时间/op | 内存/op | 分配/op |
|------|---------|---------|---------|
| **TreeList** | ~92 ns | 80 B | 1 |
| **IndexTree** | ~94 ns | 71 B | 1 |
| **AVL** | ~112 ns | 48 B | 1 |
| SkipList | ~860 ns | 175 B | 2 |

最优: TreeList

### 2. 随机 Put（50k keys，含旋转统计）

| 结构 | 时间/op | 旋转次数/op | 双旋次数/op | 树高 | 内存/op |
|------|---------|-------------|-------------|------|---------|
| **IndexTree** | ~730 ns | 0.47 | 0.23 | 26 | 161 B |
| **AVL** | ~778 ns | 3.07 | 1.00 | 26 | 137 B |

最优: IndexTree

### 3. 顺序 Put 旋转对比（50k keys）

| 结构 | 时间/op | 旋转次数/op | 双旋次数/op | 高度 |
|------|---------|-------------|-------------|------|
| **IndexTree** | ~125 ns | 1.00 | 0 | 24 |
| **AVL** | ~152 ns | 3.49 | 1.94 | 24 |

最优: IndexTree

### 4. 随机 Get（100k keys）

| 结构 | 时间/op | 内存/op | 分配/op |
|------|---------|---------|---------|
| **IndexTree** | ~135 ns | 0 B | 0 |
| **TreeList** | ~141 ns | 8 B | 1 |
| **AVL** | ~156 ns | 8 B | 1 |
| SkipList | ~650 ns | 8 B | 1 |

最优: IndexTree

说明: 当前基准下，`IndexTree` 的 `Get` 直接返回 `interface{}`，避免了其他实现把具体值装箱到接口时产生的一次额外分配。

### 5. SeekGE（100k keys）

| 结构 | 时间/op | 内存/op | 分配/op |
|------|---------|---------|---------|
| **TreeList** | ~155 ns | 7 B | 0 |
| **AVL** | ~275 ns | 328 B | 1-2 |
| SkipList | ~715 ns | 7 B | 0 |

说明: 这组历史对比主要围绕 `Seek*` 迭代接口。`IndexTree` 后续新增了迭代器 API，但原始对比仍以可直接横向比较的接口为主。

### 6. 基于索引访问（100k keys）

| 结构 | 时间/op | 内存/op | 分配/op |
|------|---------|---------|---------|
| **TreeList** | ~41 ns | 0 B | 0 |
| **IndexTree** | ~95 ns | 0 B | 0 |
| SkipList | ~102k ns | 16 B | 1 |

最优: TreeList

### 7. 混合负载（50% Put，25% Get，25% Remove）

| 结构 | 时间/op | 旋转次数/op | 高度 |
|------|---------|-------------|------|
| **IndexTree** | ~140 ns | 0.068 | 19 |
| **AVL** | ~182 ns | 0.83 | 19 |

最优: IndexTree

### 8. 迭代遍历（100k keys）

| 结构 | 时间/op | 内存/op | 分配/op |
|------|---------|---------|---------|
| **TreeList** | ~1.4-1.5M ns | 800 KB | 100k |
| **AVL** | ~2.2-2.5M ns | 800 KB | 100k |
| SkipList | ~13-15M ns | 800 KB | 100k |

最优: TreeList

---

## 结果解读

### IndexTree vs AVL

`IndexTree` 和 `AVL` 能达到相近的树高，但 `IndexTree` 以更少的结构调整获得同样的平衡效果。

- 插入密集场景下旋转次数少 3.5x 到 6.5x
- 汇总结果中写入相关吞吐高 16% 到 30%
- 混合负载下重平衡次数明显更少

结论不是 “AVL 不平衡”，而是 “更激进的高度平衡策略在这些用例中没有换来更好的最终树形”。

### TreeList vs SkipList

在单线程有序负载中，`TreeList` 基本全面领先 `SkipList`。

- 顺序插入快 6x 到 9x
- 随机读取和 `SeekGE` 快 4x 到 5x
- 按索引访问差距极大

但 `SkipList` 仍然是仓库里直接支持并发访问的有序结构。

### 选型建议

- 单线程、需要写入性能和排名访问: `IndexTree`
- 更看重顺序访问、头尾访问、范围迭代: `TreeList`
- 需要并发安全: `SkipList`
- 需要经典、容易解释的平衡树: `AVL`

---

## 运行基准

```bash
go test -bench=. -benchmem ./tree/...
go test -bench=BenchmarkTreePut -benchmem -count=5 ./tree/skiplist/...
go test -bench=BenchmarkRotation -benchmem -count=3 ./...
go test -bench=. -benchmem ./tree/indextree/...
go test -bench=. -benchmem ./tree/avl/...
go test -bench=. -benchmem ./tree/skiplist/...
go test -bench=. -benchmem ./tree/treelist/...
```

## 相关文档

- [rotation-analysis.md](./rotation-analysis.md)
- [benchmark-comparison.md](./benchmark-comparison.md)
- [../rotation-analysis.md](../rotation-analysis.md)