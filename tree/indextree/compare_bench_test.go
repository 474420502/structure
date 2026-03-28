package indextree

import (
	"testing"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
)

var sink interface{}

func BenchmarkTreePut(b *testing.B) {
	data := newBenchDataWithSeed(50000, 12345)

	b.Run("indextree", func(b *testing.B) {
		tree := New(compare.Any[int64])
		b.ResetTimer()
		for i := 0; i < b.N && i < len(data); i++ {
			sink = tree.Put(data[i], data[i])
		}
	})

	b.Run("avl", func(b *testing.B) {
		tree := avl.New[int64, int64](compare.AnyEx[int64])
		b.ResetTimer()
		for i := 0; i < b.N && i < len(data); i++ {
			sink = tree.Put(data[i], data[i])
		}
	})
}

func BenchmarkTreePutSequential(b *testing.B) {
	b.Run("indextree", func(b *testing.B) {
		tree := New(compare.Any[int])
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink = tree.Put(i, i)
		}
	})

	b.Run("avl", func(b *testing.B) {
		tree := avl.New[int, int](compare.AnyEx[int])
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink = tree.Put(i, i)
		}
	})
}

func BenchmarkTreeGet(b *testing.B) {
	data := newBenchDataWithSeed(100000, 12345)

	it := New(compare.Any[int64])
	for i := 0; i < 100000; i++ {
		it.Put(data[i], data[i])
	}

	ta := avl.New[int64, int64](compare.AnyEx[int64])
	for i := 0; i < 100000; i++ {
		ta.Put(data[i], data[i])
	}

	b.Run("indextree", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink, _ = it.Get(data[i%100000])
		}
	})

	b.Run("avl", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink, _ = ta.Get(data[i%100000])
		}
	})
}

func BenchmarkTreeRemove(b *testing.B) {
	data := newBenchDataWithSeed(100000, 12345)

	b.Run("indextree", func(b *testing.B) {
		b.StopTimer()
		tree := New(compare.Any[int64])
		for i := 0; i < 100000; i++ {
			tree.Put(data[i], data[i])
		}
		b.StartTimer()
		for i := 0; i < b.N && i < 100000; i++ {
			tree.Remove(data[i])
		}
	})

	b.Run("avl", func(b *testing.B) {
		b.StopTimer()
		tree := avl.New[int64, int64](compare.AnyEx[int64])
		for i := 0; i < 100000; i++ {
			tree.Put(data[i], data[i])
		}
		b.StartTimer()
		for i := 0; i < b.N && i < 100000; i++ {
			tree.Remove(data[i])
		}
	})
}

func TestTreeHeightCompare(t *testing.T) {
	sizes := []int{1000, 5000, 10000, 50000}
	seeds := []int64{12345, 67890}

	for _, size := range sizes {
		for _, seed := range seeds {
			data := newBenchDataWithSeed(size, seed)

			it := New(compare.Any[int64])
			for i := 0; i < size; i++ {
				it.Put(data[i], data[i])
			}
			itStats := it.BenchmarkStats()

			ta := avl.New[int64, int64](compare.AnyEx[int64])
			for i := 0; i < size; i++ {
				ta.Put(data[i], data[i])
			}
			taH := int(ta.Height())

			t.Logf("size=%d seed=%d: indextree[h=%d avg=%.2f rot/op=%.2f] avl[h=%d]",
				size, seed, itStats.Height, itStats.AvgDepth,
				float64(itStats.SingleRotations+itStats.DoubleRotations)/float64(size), taH)
		}
	}
}

func TestTreeRotationCompare(t *testing.T) {
	sizes := []int{10000, 50000}
	seeds := []int64{12345}

	for _, size := range sizes {
		for _, seed := range seeds {
			data := newBenchDataWithSeed(size, seed)

			it := New(compare.Any[int64])
			for i := 0; i < size; i++ {
				it.Put(data[i], data[i])
			}
			itStats := it.BenchmarkStats()

			ta := avl.New[int64, int64](compare.AnyEx[int64])
			for i := 0; i < size; i++ {
				ta.Put(data[i], data[i])
			}
			taStats := ta.BenchmarkStats()

			t.Logf("size=%d: indextree[rot/op=%.2f single=%d double=%d h=%d] avl[rot/op=%.2f single=%d double=%d h=%d]",
				size,
				float64(itStats.SingleRotations+itStats.DoubleRotations)/float64(size),
				itStats.SingleRotations, itStats.DoubleRotations, itStats.Height,
				float64(taStats.SingleRotations+taStats.DoubleRotations)/float64(size),
				taStats.SingleRotations, taStats.DoubleRotations, taStats.Height)
		}
	}
}

func TestIndexTreeHeightGap(t *testing.T) {
	seeds := []int64{12345, 67890}

	for _, seed := range seeds {
		data := newBenchDataWithSeed(10000, seed)

		it := New(compare.Any[int64])
		for i := 0; i < len(data); i++ {
			it.Put(data[i], data[i])
		}
		itStats := it.BenchmarkStats()

		ta := avl.New[int64, int64](compare.AnyEx[int64])
		for i := 0; i < len(data); i++ {
			ta.Put(data[i], data[i])
		}
		taStats := ta.BenchmarkStats()

		if itStats.Height > taStats.Height+2 {
			t.Fatalf("seed=%d: indextree height %d exceeds avl height %d by more than 2", seed, itStats.Height, taStats.Height)
		}

		if itStats.SingleRotations+itStats.DoubleRotations >= taStats.SingleRotations+taStats.DoubleRotations {
			t.Fatalf("seed=%d: indextree rotations %d should stay below avl rotations %d", seed, itStats.SingleRotations+itStats.DoubleRotations, taStats.SingleRotations+taStats.DoubleRotations)
		}
	}
}
