package skiplist

import (
	"math/rand"
	"testing"
	"time"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
	"github.com/474420502/structure/tree/indextree"
	"github.com/474420502/structure/tree/treelist"
)

var sink interface{}

func newBenchDataWithSeed(size int, seed int64) []int64 {
	r := rand.New(rand.NewSource(seed))
	data := make([]int64, size)
	for i := 0; i < size; i++ {
		data[i] = r.Int63()
	}
	return data
}

func BenchmarkTreePut(b *testing.B) {
	data := newBenchDataWithSeed(b.N, 12345)

	b.Run("skiplist", func(b *testing.B) {
		tree := New[int64, int64](compare.Any[int64])
		b.ResetTimer()
		for i := 0; i < b.N && i < len(data); i++ {
			sink = tree.Put(data[i], data[i])
		}
	})

	b.Run("treelist", func(b *testing.B) {
		tree := treelist.New[int64, int64](compare.Any[int64])
		b.ResetTimer()
		for i := 0; i < b.N && i < len(data); i++ {
			sink = tree.Put(data[i], data[i])
		}
	})

	b.Run("indextree", func(b *testing.B) {
		tree := indextree.New(compare.Any[int64])
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
	b.Run("skiplist", func(b *testing.B) {
		tree := New[int, int](compare.Any[int])
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink = tree.Put(i, i)
		}
	})

	b.Run("treelist", func(b *testing.B) {
		tree := treelist.New[int, int](compare.Any[int])
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink = tree.Put(i, i)
		}
	})

	b.Run("indextree", func(b *testing.B) {
		tree := indextree.New(compare.Any[int])
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
	size := 100000
	data := newBenchDataWithSeed(size, 12345)

	sk := New[int64, int64](compare.Any[int64])
	for i := 0; i < size; i++ {
		sk.Put(data[i], data[i])
	}

	tl := treelist.New[int64, int64](compare.Any[int64])
	for i := 0; i < size; i++ {
		tl.Put(data[i], data[i])
	}

	it := indextree.New(compare.Any[int64])
	for i := 0; i < size; i++ {
		it.Put(data[i], data[i])
	}

	ta := avl.New[int64, int64](compare.AnyEx[int64])
	for i := 0; i < size; i++ {
		ta.Put(data[i], data[i])
	}

	b.Run("skiplist", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink, _ = sk.Get(data[i%size])
		}
	})

	b.Run("treelist", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink, _ = tl.Get(data[i%size])
		}
	})

	b.Run("indextree", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink, _ = it.Get(data[i%size])
		}
	})

	b.Run("avl", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink, _ = ta.Get(data[i%size])
		}
	})
}

func BenchmarkTreeRemove(b *testing.B) {
	size := 100000
	data := newBenchDataWithSeed(size, 12345)

	b.Run("skiplist", func(b *testing.B) {
		b.StopTimer()
		tree := New[int64, int64](compare.Any[int64])
		for i := 0; i < size; i++ {
			tree.Put(data[i], data[i])
		}
		b.StartTimer()
		for i := 0; i < b.N && i < size; i++ {
			tree.Remove(data[i])
		}
	})

	b.Run("treelist", func(b *testing.B) {
		b.StopTimer()
		tree := treelist.New[int64, int64](compare.Any[int64])
		for i := 0; i < size; i++ {
			tree.Put(data[i], data[i])
		}
		b.StartTimer()
		for i := 0; i < b.N && i < size; i++ {
			tree.Remove(data[i])
		}
	})

	b.Run("indextree", func(b *testing.B) {
		b.StopTimer()
		tree := indextree.New(compare.Any[int64])
		for i := 0; i < size; i++ {
			tree.Put(data[i], data[i])
		}
		b.StartTimer()
		for i := 0; i < b.N && i < size; i++ {
			tree.Remove(data[i])
		}
	})

	b.Run("avl", func(b *testing.B) {
		b.StopTimer()
		tree := avl.New[int64, int64](compare.AnyEx[int64])
		for i := 0; i < size; i++ {
			tree.Put(data[i], data[i])
		}
		b.StartTimer()
		for i := 0; i < b.N && i < size; i++ {
			tree.Remove(data[i])
		}
	})
}

func BenchmarkTreeIterator(b *testing.B) {
	size := 100000
	data := newBenchDataWithSeed(size, 12345)

	sk := New[int64, int64](compare.Any[int64])
	for i := 0; i < size; i++ {
		sk.Put(data[i], data[i])
	}

	tl := treelist.New[int64, int64](compare.Any[int64])
	for i := 0; i < size; i++ {
		tl.Put(data[i], data[i])
	}

	ta := avl.New[int64, int64](compare.AnyEx[int64])
	for i := 0; i < size; i++ {
		ta.Put(data[i], data[i])
	}

	b.Run("skiplist", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			iter := sk.Iterator()
			iter.SeekToFirst()
			for iter.Valid() {
				sink = iter.Value()
				iter.Next()
			}
		}
	})

	b.Run("treelist", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			iter := tl.Iterator()
			iter.SeekToFirst()
			for iter.Valid() {
				sink = iter.Value()
				iter.Next()
			}
		}
	})

	b.Run("avl", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			iter := ta.Iterator()
			iter.SeekToFirst()
			for iter.Valid() {
				sink = iter.Value()
				iter.Next()
			}
		}
	})
}

func BenchmarkTreeSeekGE(b *testing.B) {
	size := 100000
	data := newBenchDataWithSeed(size, 12345)

	sk := New[int64, int64](compare.Any[int64])
	for i := 0; i < size; i++ {
		sk.Put(data[i], data[i])
	}

	tl := treelist.New[int64, int64](compare.Any[int64])
	for i := 0; i < size; i++ {
		tl.Put(data[i], data[i])
	}

	ta := avl.New[int64, int64](compare.AnyEx[int64])
	for i := 0; i < size; i++ {
		ta.Put(data[i], data[i])
	}

	b.Run("skiplist", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			iter := sk.Iterator()
			iter.SeekGE(data[i%size])
			if iter.Valid() {
				iter.Next()
				if iter.Valid() {
					sink = iter.Value()
				}
			}
		}
	})

	b.Run("treelist", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			iter := tl.Iterator()
			iter.SeekGE(data[i%size])
			if iter.Valid() {
				iter.Next()
				if iter.Valid() {
					sink = iter.Value()
				}
			}
		}
	})

	b.Run("avl", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			iter := ta.Iterator()
			iter.SeekGE(data[i%size])
			if iter.Valid() {
				iter.Next()
				if iter.Valid() {
					sink = iter.Value()
				}
			}
		}
	})
}

func BenchmarkTreeIndex(b *testing.B) {
	size := 100000
	data := newBenchDataWithSeed(size, 12345)

	sk := New[int64, int64](compare.Any[int64])
	for i := 0; i < size; i++ {
		sk.Put(data[i], data[i])
	}

	tl := treelist.New[int64, int64](compare.Any[int64])
	for i := 0; i < size; i++ {
		tl.Put(data[i], data[i])
	}

	b.Run("skiplist", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink = sk.Index(int64(i % size))
		}
	})

	b.Run("treelist", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sink = tl.Index(int64(i % size))
		}
	})
}

func TestTreeHeightCompare(t *testing.T) {
	sizes := []int{1000, 5000, 10000, 50000}
	seeds := []int64{12345, 67890}

	for _, size := range sizes {
		for _, seed := range seeds {
			data := newBenchDataWithSeed(size, seed)

			sk := New[int64, int64](compare.Any[int64])
			for i := 0; i < size; i++ {
				sk.Put(data[i], data[i])
			}

			tl := treelist.New[int64, int64](compare.Any[int64])
			for i := 0; i < size; i++ {
				tl.Put(data[i], data[i])
			}

			it := indextree.New(compare.Any[int64])
			for i := 0; i < size; i++ {
				it.Put(data[i], data[i])
			}
			itStats := it.BenchmarkStats()

			ta := avl.New[int64, int64](compare.AnyEx[int64])
			for i := 0; i < size; i++ {
				ta.Put(data[i], data[i])
			}
			taH := int(ta.Height())

			t.Logf("size=%d seed=%d: skiplist[h=%d] treelist[h=%d] indextree[h=%d avg=%.2f rot/op=%.2f] avl[h=%d]",
				size, seed, sk.Height(), 0, itStats.Height, itStats.AvgDepth,
				float64(itStats.SingleRotations+itStats.DoubleRotations)/float64(size), taH)
		}
	}
}

func TestTreeOperations(t *testing.T) {
	data := newBenchDataWithSeed(10000, 12345)

	sk := New[int64, int64](compare.Any[int64])
	tl := treelist.New[int64, int64](compare.Any[int64])
	it := indextree.New(compare.Any[int64])
	ta := avl.New[int64, int64](compare.AnyEx[int64])

	for i := 0; i < 5000; i++ {
		sk.Put(data[i], data[i])
		tl.Put(data[i], data[i])
		it.Put(data[i], data[i])
		ta.Put(data[i], data[i])
	}

	for i := 5000; i < 10000; i++ {
		sk.Put(data[i], data[i])
		tl.Put(data[i], data[i])
		it.Put(data[i], data[i])
		ta.Put(data[i], data[i])
	}

	t.Logf("skiplist size=%d", sk.Size())
	t.Logf("treelist size=%d", tl.Size())
	t.Logf("indextree size=%d", it.Size())
	t.Logf("avl size=%d", ta.Size())
}

func benchmarkSetTrees(size int) (*SkipList[int, int], *SkipList[int, int]) {
	tree1 := New[int, int](compare.Any[int])
	tree2 := New[int, int](compare.Any[int])

	for i := 0; i < size; i += 2 {
		tree1.Put(i, i)
	}

	for i := 1; i < size; i += 2 {
		tree2.Put(i, i)
	}

	for i := size / 4; i < size+size/4; i += 4 {
		tree1.Put(i, i)
		tree2.Put(i, i)
	}

	return tree1, tree2
}

func BenchmarkSetOperations(b *testing.B) {
	tree1, tree2 := benchmarkSetTrees(1 << 15)

	b.Run("intersection/skiplist", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			result := tree1.Intersection(tree2)
			if result.Size() == 0 {
				b.Fatal("unexpected empty intersection")
			}
		}
	})

	b.Run("union/skiplist", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			result := tree1.UnionSets(tree2)
			if result.Size() == 0 {
				b.Fatal("unexpected empty union")
			}
		}
	})

	b.Run("difference/skiplist", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			result := tree1.DifferenceSets(tree2)
			if result.Size() == 0 {
				b.Fatal("unexpected empty difference")
			}
		}
	})
}

func TestIteratorSeekForForce(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var data []int64

	for i := 0; i < 200000; i++ {
		data = append(data, r.Int63())
	}

	tree := New[int64, int64](compare.Any[int64])
	for _, v := range data {
		tree.Put(v, v)
	}

	for i := 0; i < 10000; i++ {
		v := data[r.Intn(len(data))]
		iter := tree.Iterator()
		iter.SeekGE(v)
		if iter.Valid() {
			iter.Next()
		}
	}

	t.Log("skiplist", tree.Size())
}
