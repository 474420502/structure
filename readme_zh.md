# structure 数据结构总览

这是一个面向 Go 的数据结构与搜索工具仓库，包含有序树、列表、栈、队列、映射、集合、布隆过滤器以及 A* 搜索等实现。

根文档现在只负责项目总览和导航。`indextree` 的性能对比与旋转分析已经下沉到 `tree/indextree/` 目录下，避免主文档继续只服务单一实现。

## 文档导航

- 英文总览: [readme.md](./readme.md)
- IndexTree 基准对比: [tree/indextree/benchmark-comparison.md](./tree/indextree/benchmark-comparison.md)
- IndexTree 基准对比中文: [tree/indextree/benchmark-comparison.zh.md](./tree/indextree/benchmark-comparison.zh.md)
- IndexTree 旋转分析: [tree/indextree/rotation-analysis.md](./tree/indextree/rotation-analysis.md)
- 仓库级旋转分析记录: [tree/rotation-analysis.md](./tree/rotation-analysis.md)

## 模块路径

```go
module github.com/474420502/structure
```

## 实现目录

### List

| 包 | 说明 | 文档 | 示例 |
|----|------|------|------|
| `list/array_list` | 基于动态数组的线性表，支持随机访问、迭代器和循环迭代器 | [doc](./list/array_list/doc.md) | [example](./example/array_list/main.go) |
| `list/linked_list` | 双向链表，支持迭代器和循环迭代器 | [doc](./list/linked_list/doc.md) | [example](./example/linked_list/main.go) |

### Map

| 包 | 说明 | 文档 | 示例 |
|----|------|------|------|
| `map/hashmap` | 基础哈希表封装，提供 `Put`、`Set`、`Get`、`Keys`、`Values`、`Slices` | [doc](./map/hashmap/doc.md) | [example](./example/hashmap/main.go) |
| `map/linkedhashmap` | 保留插入顺序的哈希表，支持头尾移动与覆盖更新 | [doc](./map/linkedhashmap/doc.md) | [example](./example/linkedhashmap/main.go) |
| `map/orderedmap.go` | 有序映射预留目录，目前只有空的 `OrderedMap` 类型，占位中 | [doc](./map/orderedmap.go/doc.md) | - |

### Queue

| 包 | 说明 | 文档 | 示例 |
|----|------|------|------|
| `queue/linkedarray` | 环形数组双端队列，支持头尾进出、按序号访问、遍历 | [doc](./queue/linkedarray/doc.md) | - |
| `queue/list` | 双向链表双端队列，`Front`/`Back` 返回节点句柄 | [doc](./queue/list/doc.md) | - |
| `queue/priority` | 基于大小平衡树的有序优先队列，支持迭代器与按索引删除 | [doc](./queue/priority/doc.md) | [example](./example/priority_queue/main.go) |

### Set

| 包 | 说明 | 文档 | 示例 |
|----|------|------|------|
| `set/hashset` | 无序哈希集合，适合快速成员判断 | [doc](./set/hashset/doc.md) | - |
| `set/treeset` | 基于 AVL 风格平衡树的有序集合，支持双向迭代 | [doc](./set/treeset/doc.md) | [example](./example/treeset/main.go) |

### Stack

| 包 | 说明 | 文档 | 示例 |
|----|------|------|------|
| `stack/array` | 基于切片的栈 | [doc](./stack/array/doc.md) | [example](./example/arraystack/main.go) |
| `stack/list` | 基于链表的栈 | [doc](./stack/list/doc.md) | [example](./example/liststack/main.go) |
| `stack/listarray` | 由链式数组块组成的分段栈 | [doc](./stack/listarray/doc.md) | [example](./example/lastack/main.go) |

### Tree 与有序结构

| 包 | 说明 | 文档 | 示例 |
|----|------|------|------|
| `tree/avl` | 经典 AVL 平衡树，支持迭代器 | [doc](./tree/avl/doc.md) | [example](./example/avl/main.go) |
| `tree/avls` | 支持重复键的 AVL 变体 | [doc](./tree/avls/doc.md) | - |
| `tree/heap` | 基于比较器的二叉堆 | [doc](./tree/heap/doc.md) | [example](./example/heap/main.go) |
| `tree/indextree` | 以子树大小维持平衡的有序树，支持排名、索引、切分、裁剪 | [doc](./tree/indextree/doc.md) | [example](./example/indextree/main.go) |
| `tree/skiplist` | 带 `RWMutex` 的并发跳表，支持迭代器、索引、裁剪、集合运算 | [doc](./tree/skiplist/doc.md) | - |
| `tree/treelist` | 带链式顺序指针的有序树，支持范围迭代与集合运算 | [doc](./tree/treelist/doc.md) | [example](./example/tree-treelist/main.go) |

### Filter

| 包 | 说明 | 文档 | 示例 |
|----|------|------|------|
| `filter/bloom` | 支持序列化的布隆过滤器 | [doc](./filter/bloom/doc.md) | [example](./example/bloom/main.go) |

### Graph Search

| 包 | 说明 | 文档 | 示例 |
|----|------|------|------|
| `graph/astar` | 网格版 A* 搜索，支持自定义邻接、代价、权重策略 | [doc](./graph/astar/doc.md) | - |

### Search

| 包 | 说明 | 文档 | 示例 |
|----|------|------|------|
| `search/treelist` | 面向搜索索引场景的字节串有序树 | [doc](./search/treelist/doc.md) | [example](./example/search-treelist/main.go) |
| `search/searchtree` | 搜索索引抽象预留目录，目前没有可直接使用的公开 API | [doc](./search/searchtree/doc.md) | - |

## 选型建议

- 需要综合读写性能与排名能力: `tree/indextree`
- 需要顺序访问、头尾访问与范围迭代: `tree/treelist`
- 需要并发安全的有序结构: `tree/skiplist`
- 需要经典平衡二叉树: `tree/avl`
- 需要简单的无序映射/集合: `map/hashmap`、`set/hashset`

## 基准与分析文档

原来放在根目录 README 的性能对比内容已经移动到 `indextree` 专项文档：

- [tree/indextree/benchmark-comparison.md](./tree/indextree/benchmark-comparison.md)
- [tree/indextree/benchmark-comparison.zh.md](./tree/indextree/benchmark-comparison.zh.md)
- [tree/indextree/rotation-analysis.md](./tree/indextree/rotation-analysis.md)
- [tree/rotation-analysis.md](./tree/rotation-analysis.md)

## 运行测试

```bash
go test ./...
```

## 运行树结构基准

```bash
go test -bench=. -benchmem ./tree/...
go test -bench=BenchmarkRotation -benchmem -count=3 ./...
```
