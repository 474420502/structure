package utils

import (
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/indextree"
	"github.com/474420502/structure/tree/itree"
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

type iTreeAdapter struct {
	tree *itree.Tree[int, int]
}

func newITreeAdapter() benchTree {
	return &iTreeAdapter{tree: itree.New[int, int](compare.AnyEx[int])}
}

func (adapter *iTreeAdapter) Put(key int) {
	adapter.tree.Put(key, key)
}

func (adapter *iTreeAdapter) Get(key int) {
	adapter.tree.Get(key)
}

func (adapter *iTreeAdapter) Remove(key int) {
	adapter.tree.Remove(key)
}

func (adapter *iTreeAdapter) Size() int {
	return adapter.tree.Size()
}

func (adapter *iTreeAdapter) ResetBenchmarkStats() {
	adapter.tree.ResetBenchmarkStats()
}

func (adapter *iTreeAdapter) BenchmarkStats() benchStats {
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

func benchmarkTreePutRandom(b *testing.B, createTree func() benchTree) {
	b.StopTimer()
	keys := newBenchKeys(b.N+10000, 12345)
	tree := createTree()
	for i := 0; i < 10000; i++ {
		tree.Put(keys[i])
	}
	tree.ResetBenchmarkStats()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Put(keys[i+10000])
	}

	stats := tree.BenchmarkStats()
	reportBenchmarkMetrics(b, tree, b.N, stats.singleRotations, stats.doubleRotations)
}

func benchmarkTreePutSequential(b *testing.B, createTree func() benchTree) {
	tree := createTree()
	tree.ResetBenchmarkStats()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(i)
	}

	stats := tree.BenchmarkStats()
	reportBenchmarkMetrics(b, tree, b.N, stats.singleRotations, stats.doubleRotations)
}

func benchmarkTreeRemoveRandom(b *testing.B, createTree func() benchTree) {
	random := rand.New(rand.NewSource(1683721792150515321))
	tree := createTree()
	removeList := make([]int, 0, 1100)
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
			for j := 0; j < 1000; j++ {
				value := random.Intn(1000)
				before := tree.Size()
				tree.Put(value)
				if tree.Size() != before {
					removeList = append(removeList, value)
				}

				if j%25 == 0 {
					removeList = append(removeList, random.Intn(1000))
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

func benchmarkTreeMixed(b *testing.B, createTree func() benchTree) {
	random := rand.New(rand.NewSource(1683989312052736623))
	tree := createTree()

	b.StopTimer()
	for i := 0; i < 5000; i++ {
		tree.Put(random.Intn(20000))
	}
	tree.ResetBenchmarkStats()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		value := random.Intn(20000)
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

func BenchmarkRotationComparePutRandom(b *testing.B) {
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreePutRandom(b, newIndexTreeAdapter)
	})
	b.Run("itree", func(b *testing.B) {
		benchmarkTreePutRandom(b, newITreeAdapter)
	})
}

func BenchmarkRotationComparePutSequential(b *testing.B) {
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreePutSequential(b, newIndexTreeAdapter)
	})
	b.Run("itree", func(b *testing.B) {
		benchmarkTreePutSequential(b, newITreeAdapter)
	})
}

func BenchmarkRotationCompareRemoveRandom(b *testing.B) {
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeRemoveRandom(b, newIndexTreeAdapter)
	})
	b.Run("itree", func(b *testing.B) {
		benchmarkTreeRemoveRandom(b, newITreeAdapter)
	})
}

func BenchmarkRotationCompareMixed(b *testing.B) {
	b.Run("indextree", func(b *testing.B) {
		benchmarkTreeMixed(b, newIndexTreeAdapter)
	})
	b.Run("itree", func(b *testing.B) {
		benchmarkTreeMixed(b, newITreeAdapter)
	})
}