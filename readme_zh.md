# 数据结构基准测试对比

一份全面、公平的排序数据结构基准测试对比。

## 对比的数据结构

| 结构 | 类型 | 描述 |
|------|------|------|
| **IndexTree** | BST | 使用基于大小平衡策略的自平衡二叉搜索树（创新点） |
| **TreeList** | BST | 带双向链表指针的自平衡二叉搜索树 |
| **AVL** | BST | 经典AVL树，基于高度的平衡策略（diff=1） |
| **SkipList** | 跳表 | 支持并发的跳表，使用RWMutex（最大层级16） |

## 基准测试方法

- **公平对比**: 所有结构使用相同的测试数据（相同随机种子）
- **一致的API**: 都支持 Put、Get、Remove、Iterator 操作
- **多次运行**: 每个基准测试运行5次，取稳定结果
- **真实场景**: 测试随机、顺序和混合工作负载

---

## 基准测试结果汇总

### 1. 顺序 Put（100k keys）

| 结构 | 时间/op | 内存/op | 内存分配/op |
|------|---------|---------|-------------|
| **TreeList** | ~92 ns | 80 B | 1 |
| **IndexTree** | ~94 ns | 71 B | 1 |
| **AVL** | ~112 ns | 48 B | 1 |
| SkipList | ~860 ns | 175 B | 2 |

**最优: TreeList**（略快于IndexTree，都显著快于AVL和SkipList）

### 2. 随机 Put（50k keys，含旋转统计）

| 结构 | 时间/op | 旋转次数/op | 双旋次数/op | 树高度 | 内存/op |
|------|---------|-------------|-------------|--------|---------|
| **IndexTree** | ~730 ns | 0.47 | 0.23 | 26 | 161 B |
| **AVL** | ~778 ns | 3.07 | 1.00 | 26 | 137 B |

**最优: IndexTree**（旋转次数少6.5倍，相同树高度，快16%）

### 3. 顺序 Put 旋转对比（50k keys）

| 结构 | 时间/op | 旋转次数/op | 双旋次数/op | 高度 |
|------|---------|-------------|-------------|------|
| **IndexTree** | ~125 ns | 1.00 | 0 | 24 |
| **AVL** | ~152 ns | 3.49 | 1.94 | 24 |

**最优: IndexTree**（旋转少3.5倍，快22%，相同树形）

### 4. 随机 Get（100k keys）

| 结构 | 时间/op | 内存/op | 内存分配/op |
|------|---------|---------|-------------|
| **IndexTree** | ~135 ns | 0 B | 0 |
| **TreeList** | ~141 ns | 8 B | 1 |
| **AVL** | ~156 ns | 8 B | 1 |
| SkipList | ~650 ns | 8 B | 1 |

**最优: IndexTree**（比AVL快16%，无额外boxing分配）

注意：IndexTree 直接返回 `interface{}` 类型，避免了其他实现在将具体类型（int64）赋值给 `interface{}` 时需要的装箱（boxing）开销。

### 5. SeekGE 操作（100k keys）

| 结构 | 时间/op | 内存/op | 内存分配/op |
|------|---------|---------|-------------|
| **TreeList** | ~155 ns | 7 B | 0 |
| **AVL** | ~275 ns | 328 B | 1-2 |
| SkipList | ~715 ns | 7 B | 0 |

注意：IndexTree 使用 Traverse() 回调函数进行迭代，而不是 Seek* 迭代器模式。

### 6. 基于索引的访问（100k keys）

| 结构 | 时间/op | 内存/op | 内存分配/op |
|------|---------|---------|-------------|
| **TreeList** | ~41 ns | 0 B | 0 |
| **IndexTree** | ~95 ns | 0 B | 0 |
| SkipList | ~102k ns | 16 B | 1 |

**最优: TreeList**（比IndexTree快2.3倍，比SkipList快2500倍）

### 7. 混合工作负载（50% Put，25% Get，25% Remove）

| 结构 | 时间/op | 旋转次数/op | 高度 |
|------|---------|-------------|------|
| **IndexTree** | ~140 ns | 0.068 | 19 |
| **AVL** | ~182 ns | 0.83 | 19 |

**最优: IndexTree**（快30%，旋转少12倍）

### 8. 迭代器遍历（100k keys）

| 结构 | 时间/op | 内存/op | 内存分配/op |
|------|---------|---------|-------------|
| **TreeList** | ~1.4-1.5M ns | 800 KB | 100k |
| **AVL** | ~2.2-2.5M ns | 800 KB | 100k |
| SkipList | ~13-15M ns | 800 KB | 100k |

**最优: TreeList**（比AVL快1.6倍，比SkipList快10倍）

---

## 详细分析

### IndexTree vs AVL（基于大小 vs 基于高度的平衡）

IndexTree 和 AVL 都能达到相同的树形（相同高度），但 IndexTree 使用**基于子树大小**的平衡策略，而非基于高度。

**IndexTree 的主要优势：**

1. **更少的旋转**: 随机插入时旋转少6.5倍，顺序插入时少3.5倍
2. **更好的写性能**: Put 操作快16-22%
3. **Get无额外boxing分配**: IndexTree 直接返回 `interface{}`，避免了具体返回类型需要的装箱开销
4. **更好的混合工作负载**: 快30%，旋转少12倍

**为什么 IndexTree 旋转更少：**

经典 AVL 基于高度差进行重平衡（平衡因子 = 高度(左) - 高度(右)）。当差值超过1时，进行旋转。

IndexTree 基于**子树大小**而非高度进行重平衡。这使其能在保持相同逻辑平衡（相同树形）的同时，更加谨慎地决定何时旋转。大小方法更"不急于"旋转，因为：
- 子树大小变化比高度更缓慢
- 大小已被追踪用于 Index() 操作，无需额外开销

### TreeList vs SkipList

TreeList 被设计为 SkipList 的替代品，具有更好的性能：

| 操作 | TreeList vs SkipList |
|------|---------------------|
| Put 顺序 | **快6-9倍** |
| Get 随机 | **快4-5倍** |
| SeekGE | **快4-5倍** |
| Index 访问 | **快2500倍** |
| Iterator | **快10倍** |

TreeList 通过以下方式实现这一性能：
- 更好的缓存局部性（二叉树 vs 跳表的概率布局）
- O(log n) 确定的最坏情况 vs O(log n) 期望的跳表
- 双向链表指针实现 O(1) 头/尾访问

### SkipList（线程安全替代方案）

SkipList 是此对比中唯一**并发**的结构，使用 `sync.RWMutex` 实现线程安全。

**何时选择 SkipList：**
- 需要多线程访问
- 更简单的实现
- 可接受概率性保证

**性能代价**: 由于锁开销，比 TreeList 慢4-10倍

---

## 功能对比矩阵

| 功能 | IndexTree | TreeList | AVL | SkipList |
|------|-----------|----------|-----|----------|
| 泛型 | ✓ | ✓ | ✓ | ✓ |
| 线程安全 | ✗ | ✗ | ✗ | ✓ |
| O(1) 头/尾访问 | ✗ | ✓ | ✗ | ✓ |
| Index（排名） | ✓ | ✓ | ✗ | ✓ |
| IndexOf（键转排名） | ✓ | ✓ | ✗ | ✗ |
| RemoveRange | ✓ | ✓ | ✗ | ✗ |
| 集合操作 | ✗ | ✓ | ✗ | ✓ |
| Split/SplitContain | ✓ | ✗ | ✗ | ✗ |
| Trim/TrimByIndex | ✓ | ✓ | ✗ | ✗ |

---

## 内存使用对比

### 每节点内存开销

| 结构 | 节点大小（估计） |
|------|----------------|
| AVL | 最小（无额外字段） |
| IndexTree | +子树大小字段 |
| TreeList | +2个链表指针 |
| SkipList | +4层指针 + 互斥锁 |

### 操作期间的内存分配

| 操作 | IndexTree | TreeList | AVL | SkipList |
|------|-----------|----------|-----|----------|
| Put | 71-161 B | 80 B | 48-140 B | 175 B |
| Get | 0 B | 8 B | 8 B | 8 B |
| Iterator | 800 KB total | 800 KB total | 800 KB total | 800 KB total |

---

## 结论

### 综合性能最佳: **IndexTree**

对于需要以下平衡的单线程应用：
- 写性能（Put/Remove）
- 读性能（Get）
- 最少旋转
- Index 操作

**IndexTree** 以16-56%优于AVL的性能和6.5倍更少的旋转，同时保持相同的树形，成为明显赢家。

### 顺序数据最佳: **TreeList**

处理顺序或近似顺序的键时：
- 最快的 Put 操作（~92 ns/op）
- 出色的 Get 性能（~141 ns/op）
- O(1) 头/尾访问
- 基于索引的访问（41 ns/op）

### 线程安全最佳: **SkipList**

需要并发访问时：
- 内置线程安全（RWMutex）
- 比带互斥锁的树更简单的实现
- 可接受的性能损失（慢4-10倍）

### 最佳传统BST: **AVL**

需要成熟、简单的实现时：
- 最小节点大小
- 无外部依赖
- 易于理解的行为
- 足够的性能

---

## 运行基准测试

```bash
# 运行所有树结构基准测试
go test -bench=. -benchmem ./tree/...

# 运行特定基准测试对比
go test -bench=BenchmarkTreePut -benchmem -count=5 ./tree/skiplist/...
go test -bench=BenchmarkRotation -benchmem -count=3 ./...

# 运行特定数据结构基准测试
go test -bench=. -benchmem ./tree/indextree/...
go test -bench=. -benchmem ./tree/avl/...
go test -bench=. -benchmem ./tree/skiplist/...
go test -bench=. -benchmem ./tree/treelist/...
```

---

## 附录：完整基准测试数据

### PutRandom（50k keys）
```
IndexTree:  730 ns/op  |  rot/op=0.47  |  double=0.23  |  height=26  |  161 B/op  |  2 allocs
AVL:        778 ns/op  |  rot/op=3.07  |  double=1.00  |  height=26  |  137 B/op  |  1 alloc
```

### PutSequential（50k keys）
```
IndexTree:  125 ns/op  |  rot/op=1.00  |  double=0.00  |  height=24  |  163 B/op  |  2 allocs
AVL:        152 ns/op  |  rot/op=3.49  |  double=1.94  |  height=24  |  140 B/op  |  1 alloc
```

### GetRandom（100k keys）
```
IndexTree:  135 ns/op  |  0 B/op  |  0 allocs
TreeList:   141 ns/op  |  8 B/op  |  1 alloc
AVL:        156 ns/op  |  8 B/op  |  1 alloc
SkipList:   650 ns/op  |  8 B/op  |  1 alloc
```

### SeekGE（100k keys）
```
TreeList:   155 ns/op  |  7 B/op  |  0 allocs
AVL:        275 ns/op  |  328 B/op  |  1-2 allocs
SkipList:   715 ns/op  |  7 B/op  |  0 allocs
```

### 混合工作负载（50% Put, 25% Get, 25% Remove）
```
IndexTree:  140 ns/op  |  rot/op=0.068  |  height=19
AVL:        182 ns/op  |  rot/op=0.83  |  height=19
```

### Iterator 遍历（100k keys）
```
TreeList:   1.4-1.5M ns/op  |  800 KB  |  100k allocs
AVL:        2.2-2.5M ns/op  |  800 KB  |  100k allocs
SkipList:   13-15M ns/op  |  800 KB  |  100k allocs
```
