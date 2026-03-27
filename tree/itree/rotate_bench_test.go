package itree

import (
	"testing"

	random "github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

type rotateTypeFunc func(cur *Node[int, int], child int, ls, rs int) bool

func reportRotateMetricsWithCounts(b *testing.B, tree *Tree[int, int], operations int, single int, double int) {
	b.StopTimer()
	height, avgDepth, p50Depth, p95Depth := tree.shapeStats()
	totalRotations := single + double

	b.ReportMetric(float64(totalRotations)/float64(operations), "rot/op")
	b.ReportMetric(float64(double)/float64(operations), "double/op")
	b.ReportMetric(float64(height), "height")
	b.ReportMetric(avgDepth, "avg-depth")
	b.ReportMetric(float64(p50Depth), "p50-depth")
	b.ReportMetric(float64(p95Depth), "p95-depth")
}

func reportRotateMetrics(b *testing.B, tree *Tree[int, int], operations int) {
	single, double := tree.rotationStats()
	reportRotateMetricsWithCounts(b, tree, operations, single, double)
}

func benchmarkPutRotateType(b *testing.B, rotateType rotateTypeFunc) {
	rand := random.New(1683721792150515321)
	tree := New[int, int](compare.AnyEx[int])
	tree.rotateType = rotateType

	b.StopTimer()
	for i := 0; i < 10000; i++ {
		v := rand.Int()
		tree.Put(v, v)
	}
	tree.resetRotationStats()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v := rand.Int()
		tree.Put(v, v)
	}

	reportRotateMetrics(b, tree, b.N)
}

func benchmarkRemoveRotateType(b *testing.B, rotateType rotateTypeFunc) {
	rand := random.New(1683721792150515321)
	tree := New[int, int](compare.AnyEx[int])
	tree.rotateType = rotateType

	var removeList []int
	removeIndex := 0
	totalSingle := 0
	totalDouble := 0
	epochActive := false

	for i := 0; i < b.N; i++ {
		if tree.Size() == 0 {
			if epochActive {
				single, double := tree.rotationStats()
				totalSingle += single
				totalDouble += double
				epochActive = false
			}
			b.StopTimer()
			removeList = removeList[:0]
			removeIndex = 0
			tree.resetRotationStats()
			for j := 0; j < 1000; j++ {
				v := rand.Intn(1000)
				if tree.Put(v, v) {
					removeList = append(removeList, v)
				}

				if j%25 == 0 {
					removeList = append(removeList, rand.Intn(1000))
				}
			}
			epochActive = true
			b.StartTimer()
		}

		v := removeList[removeIndex]
		tree.Remove(v)
		removeIndex++
	}

	if epochActive {
		single, double := tree.rotationStats()
		totalSingle += single
		totalDouble += double
	}

	reportRotateMetricsWithCounts(b, tree, b.N, totalSingle, totalDouble)
}

func benchmarkMixedRotateType(b *testing.B, rotateType rotateTypeFunc) {
	rand := random.New(1683989312052736623)
	tree := New[int, int](compare.AnyEx[int])
	tree.rotateType = rotateType

	b.StopTimer()
	for i := 0; i < 5000; i++ {
		v := rand.Intn(20000)
		tree.Put(v, v)
	}
	tree.resetRotationStats()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v := rand.Intn(20000)
		switch i % 3 {
		case 0:
			tree.Put(v, v)
		case 1:
			tree.Get(v)
		default:
			tree.Remove(v)
		}
	}

	reportRotateMetrics(b, tree, b.N)
}

func BenchmarkRotateDecisionPut(b *testing.B) {
	b.Run("formula", func(b *testing.B) {
		benchmarkPutRotateType(b, sizeRotateType[int, int])
	})
	b.Run("child-bias", func(b *testing.B) {
		benchmarkPutRotateType(b, sizeRotateTypeChildBias[int, int])
	})
}

func BenchmarkRotateDecisionRemove(b *testing.B) {
	b.Run("formula", func(b *testing.B) {
		benchmarkRemoveRotateType(b, sizeRotateType[int, int])
	})
	b.Run("child-bias", func(b *testing.B) {
		benchmarkRemoveRotateType(b, sizeRotateTypeChildBias[int, int])
	})
}

func BenchmarkRotateDecisionMixed(b *testing.B) {
	b.Run("formula", func(b *testing.B) {
		benchmarkMixedRotateType(b, sizeRotateType[int, int])
	})
	b.Run("child-bias", func(b *testing.B) {
		benchmarkMixedRotateType(b, sizeRotateTypeChildBias[int, int])
	})
}