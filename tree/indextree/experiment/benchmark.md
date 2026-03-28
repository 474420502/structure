# IndexTree Experiment Notes

## Scope

`tree/indextree/experiment` 保留为实验目录，但当前不再维护多阶段容差实现。

目录中的实验树现在统一采用与主包 `tree/indextree` 一致的单阶段尺寸容差算法，实验实现名为 `ShiftToleranceTree`：

- 单一参数 `shift`
- 阈值公式 `bottomsize = (1 << height >> shift) + 1`
- 不做第二阶段确认

其中 `shift=2` 代表当前主包的默认策略，实验代码中以 `indextree-default` 表示。

## Purpose

当前实验的目标是：

1. 验证主包等价算法在不同 `shift` 下的形态和旋转开销
2. 与 AVL 基线做高度、旋转次数、读写性能对比
3. 为后续新的实验方向保留独立目录和基准框架

## Current Configurations

| Config | Meaning |
|--------|---------|
| `indextree-default` | 与 `tree/indextree` 当前默认实现一致 |
| `shift=3` | 更宽松的单阶段容差 |
| `shift=4` | 更宽松的单阶段容差 |
| `shift=5` | 更宽松的单阶段容差 |
| `AVL` | 高度平衡基线 |

## Run

```bash
go test ./tree/indextree/experiment -v
go test ./tree/indextree/experiment -run TestComprehensiveReport -v
go test ./tree/indextree/experiment -bench=. -benchmem -run=^$
```

## Notes

- 如果需要做新的平衡策略实验，建议新增独立实现文件，不要直接改写 `indextree-default` 行为。
- `indextree-default` 的语义应始终与主包 `tree/indextree` 保持一致，便于回归和对照。