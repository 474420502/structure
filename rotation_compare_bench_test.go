package utils

import (
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
	"github.com/474420502/structure/tree/indextree"
)

type benchTree interface {
	Put(int)
	Get(int)
	Remove(int)
	Size() int
	ResetBenchmarkStats()
	BenchmarkStats() benchStats
}

type benchStats struct {
	singleRotations int
	doubleRotations int
	height          int
	avgDepth        float64
	p50Depth        int
	p95Depth        int
}

type indexTreeAdapter struct {
	tree *indextree.Tree[int]
}

func newIndexTreeAdapter() benchTree {
	return &indexTreeAdapter{tree: indextree.New(compare.Any[int])}
}

func (adapter *indexTreeAdapter) Put(key int) {
	adapter.tree.Put(key, key)
}

func (adapter *indexTreeAdapter) Get(key int) {
	adapter.tree.Get(key)
}

func (adapter *indexTreeAdapter) Remove(key int) {
	adapter.tree.Remove(key)
}

func (adapter *indexTreeAdapter) Size() int {
	return int(adapter.tree.Size())
}

func (adapter *indexTreeAdapter) ResetBenchmarkStats() {
	adapter.tree.ResetBenchmarkStats()
}

func (adapter *indexTreeAdapter) BenchmarkStats() benchStats {
	stats := adapter.tree.BenchmarkStats()
	return benchStats{
		singleRotations: stats.SingleRotations,
		doubleRotations: stats.DoubleRotations,
		height:          stats.Height,
		avgDepth:        stats.AvgDepth,
		p50Depth:        stats.P50Depth,
		p95Depth:        stats.P95Depth,
	}
}

type avlTreeAdapter struct {
	tree *avl.Tree[int, int]
}

func newAVLTreeAdapter() benchTree {
	return &avlTreeAdapter{tree: avl.New[int, int](compare.AnyEx[int])}
}

func (adapter *avlTreeAdapter) Put(key int) {
	adapter.tree.Put(key, key)
}

func (adapter *avlTreeAdapter) Get(key int) {
	adapter.tree.Get(key)
}

func (adapter *avlTreeAdapter) Remove(key int) {
	adapter.tree.Remove(key)
}

func (adapter *avlTreeAdapter) Size() int {
	return int(adapter.tree.Size())
}

func (adapter *avlTreeAdapter) ResetBenchmarkStats() {
	adapter.tree.ResetBenchmarkStats()
}

func (adapter *avlTreeAdapter) BenchmarkStats() benchStats {
	stats := adapter.tree.BenchmarkStats()
	return benchStats{
		singleRotations: stats.SingleRotations,
		doubleRotations: stats.DoubleRotations,
		height:          stats.Height,
		avgDepth:        stats.AvgDepth,
		p50Depth:        stats.P50Depth,
		p95Depth:        stats.P95Depth,
	}
}

func reportBenchmarkMetrics(b *testing.B, tree benchTree, operations int, singleRotations int, doubleRotations int) {
	b.StopTimer()
	stats := tree.BenchmarkStats()
	totalRotations := singleRotations + doubleRotations
	b.ReportMetric(float64(totalRotations)/float64(operations), "rot/op")
	b.ReportMetric(float64(doubleRotations)/float64(operations), "double/op")
	b.ReportMetric(float64(stats.height), "height")
	b.ReportMetric(stats.avgDepth, "avg-depth")
	b.ReportMetric(float64(stats.p50Depth), "p50-depth")
	b.ReportMetric(float64(stats.p95Depth), "p95-depth")
}

func newBenchKeys(n int, seed int64) []int {
	random := rand.New(rand.NewSource(seed))
	keys := make([]int, n)
	for i := 0; i < n; i++ {
		keys[i] = random.Int()
	}
	return keys
}

func benchmarkTreePutRandom(b *testing.B, createTree func() benchTree, prepSize int) {
	b.StopTimer()
	keys := newBenchKeys(prepSize+b.N, 12345)
	tree := createTree()
	for i := 0; i < prepSize; i++ {
		tree.Put(keys[i])
	}
	tree.ResetBenchmarkStats()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Put(keys[i+prepSize])
	}

	stats := tree.BenchmarkStats()
	reportBenchmarkMetrics(b, tree, b.N, stats.singleRotations, stats.doubleRotations)
}

func benchmarkTreePutSequential(b *testing.B, createTree func() benchTree, prepSize int) {
	tree := createTree()
	tree.ResetBenchmarkStats()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Put(prepSize + i)
	}

	stats := tree.BenchmarkStats()
	reportBenchmarkMetrics(b, tree, b.N, stats.singleRotations, stats.doubleRotations)
}

func benchmarkTreeRemoveRandom(b *testing.B, createTree func() benchTree, prepSize int) {
	random := rand.New(rand.NewSource(1683721792150515321))
	tree := createTree()
	removeList := make([]int, 0, prepSize+100)
	removeIndex := 0
	totalSingleRotations := 0
	totalDoubleRotations := 0
	epochActive := false

	for i := 0; i < b.N; i++ {
		if tree.Size() == 0 {
			if epochActive {
				stats := tree.BenchmarkStats()
				totalSingleRotations += stats.singleRotations
				totalDoubleRotations += stats.doubleRotations
				epochActive = false
			}
			b.StopTimer()
			removeList = removeList[:0]
			removeIndex = 0
			tree.ResetBenchmarkStats()
			for j := 0; j < prepSize; j++ {
				value := random.Intn(prepSize)
				before := tree.Size()
				tree.Put(value)
				if tree.Size() != before {
					removeList = append(removeList, value)
				}

				if j%25 == 0 {
					removeList = append(removeList, random.Intn(prepSize))
				}
			}
			epochActive = true
			b.StartTimer()
		}

		tree.Remove(removeList[removeIndex])
		removeIndex++
	}

	if epochActive {
		stats := tree.BenchmarkStats()
		totalSingleRotations += stats.singleRotations
		totalDoubleRotations += stats.doubleRotations
	}

	reportBenchmarkMetrics(b, tree, b.N, totalSingleRotations, totalDoubleRotations)
}

func benchmarkTreeMixed(b *testing.B, createTree func() benchTree, prepSize int) {
	random := rand.New(rand.NewSource(1683989312052736623))
	tree := createTree()

	b.StopTimer()
	for i := 0; i < prepSize; i++ {
		tree.Put(random.Intn(prepSize * 2))
	}
	tree.ResetBenchmarkStats()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		value := random.Intn(prepSize * 2)
		switch i % 3 {
		case 0:
			tree.Put(value)
		case 1:
			tree.Get(value)
		default:
			tree.Remove(value)
		}
	}

	stats := tree.BenchmarkStats()
	reportBenchmarkMetrics(b, tree, b.N, stats.singleRotations, stats.doubleRotations)
}

func BenchmarkRotationComparePutRandom10k(b *testing.B) {
	const prepSize = 10000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreePutRandom(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreePutRandom(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationComparePutRandom20k(b *testing.B) {
	const prepSize = 20000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreePutRandom(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreePutRandom(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationComparePutRandom50k(b *testing.B) {
	const prepSize = 50000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreePutRandom(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreePutRandom(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationComparePutSequential10k(b *testing.B) {
	const prepSize = 10000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreePutSequential(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreePutSequential(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationComparePutSequential20k(b *testing.B) {
	const prepSize = 20000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreePutSequential(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreePutSequential(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationComparePutSequential50k(b *testing.B) {
	const prepSize = 50000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreePutSequential(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreePutSequential(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareRemoveRandom10k(b *testing.B) {
	const prepSize = 10000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeRemoveRandom(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeRemoveRandom(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareRemoveRandom20k(b *testing.B) {
	const prepSize = 20000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeRemoveRandom(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeRemoveRandom(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareRemoveRandom50k(b *testing.B) {
	const prepSize = 50000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeRemoveRandom(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeRemoveRandom(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareMixed10k(b *testing.B) {
	const prepSize = 10000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeMixed(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeMixed(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareMixed20k(b *testing.B) {
	const prepSize = 20000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeMixed(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeMixed(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareMixed50k(b *testing.B) {
	const prepSize = 50000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeMixed(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeMixed(b, newAVLTreeAdapter, prepSize)
	})
}

func benchmarkTreeGetRandom(b *testing.B, createTree func() benchTree, prepSize int) {
	b.StopTimer()
	random := rand.New(rand.NewSource(12345))
	keys := make([]int, prepSize)
	tree := createTree()
	for i := 0; i < prepSize; i++ {
		keys[i] = random.Int()
		tree.Put(keys[i])
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Get(keys[i%prepSize])
	}

	reportBenchmarkMetrics(b, tree, b.N, 0, 0)
}

func benchmarkTreeGetSequential(b *testing.B, createTree func() benchTree, prepSize int) {
	b.StopTimer()
	tree := createTree()
	for i := 0; i < prepSize; i++ {
		tree.Put(i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Get(i % prepSize)
	}

	reportBenchmarkMetrics(b, tree, b.N, 0, 0)
}

func benchmarkTreeGetMixed(b *testing.B, createTree func() benchTree, prepSize int) {
	b.StopTimer()
	random := rand.New(rand.NewSource(1683989312052736623))
	keys := make([]int, prepSize*2)
	tree := createTree()
	for i := 0; i < prepSize; i++ {
		keys[i] = random.Intn(prepSize * 2)
		tree.Put(keys[i])
	}
	for i := 0; i < prepSize*2; i++ {
		keys[i] = random.Intn(prepSize * 2)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Get(keys[i%len(keys)])
	}

	reportBenchmarkMetrics(b, tree, b.N, 0, 0)
}

func BenchmarkRotationCompareGetRandom10k(b *testing.B) {
	const prepSize = 10000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeGetRandom(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeGetRandom(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareGetRandom20k(b *testing.B) {
	const prepSize = 20000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeGetRandom(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeGetRandom(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareGetRandom50k(b *testing.B) {
	const prepSize = 50000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeGetRandom(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeGetRandom(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareGetSequential10k(b *testing.B) {
	const prepSize = 10000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeGetSequential(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeGetSequential(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareGetSequential20k(b *testing.B) {
	const prepSize = 20000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeGetSequential(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeGetSequential(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareGetSequential50k(b *testing.B) {
	const prepSize = 50000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeGetSequential(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeGetSequential(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareGetMixed10k(b *testing.B) {
	const prepSize = 10000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeGetMixed(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeGetMixed(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareGetMixed20k(b *testing.B) {
	const prepSize = 20000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeGetMixed(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeGetMixed(b, newAVLTreeAdapter, prepSize)
	})
}

func BenchmarkRotationCompareGetMixed50k(b *testing.B) {
	const prepSize = 50000
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeGetMixed(b, newIndexTreeAdapter, prepSize)
	})
	b.Run("avl", func(b *testing.B) {
		benchmarkTreeGetMixed(b, newAVLTreeAdapter, prepSize)
	})
}
