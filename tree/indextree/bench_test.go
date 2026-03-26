package indextree

import (
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func newBenchDataWithSeed(n int, seed int64) []int64 {
	random := rand.New(rand.NewSource(seed))
	keys := make([]int64, n)
	for i := 0; i < n; i++ {
		keys[i] = random.Int63()
	}
	return keys
}

func BenchmarkPut(b *testing.B) {
	data := newBenchDataWithSeed(4000000, 12345)

	b.Run("indextree", func(b *testing.B) {
		tree := New(compare.Any[int64])
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tree.Put(data[i], data[i])
		}
	})

	b.Run("avl", func(b *testing.B) {
		tree := avl.New[int64, int64](compare.AnyEx[int64])
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tree.Put(data[i], data[i])
		}
	})
}

func BenchmarkPutSequential(b *testing.B) {
	b.Run("indextree", func(b *testing.B) {
		tree := New(compare.Any[int])
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tree.Put(i, i)
		}
	})

	b.Run("avl", func(b *testing.B) {
		tree := avl.New[int, int](compare.AnyEx[int])
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tree.Put(i, i)
		}
	})
}

func BenchmarkGet(b *testing.B) {
	data := newBenchDataWithSeed(1000000, 12345)

	tree1 := New(compare.Any[int64])
	for i := 0; i < 1000000; i++ {
		tree1.Put(data[i], data[i])
	}

	tree2 := avl.New[int64, int64](compare.AnyEx[int64])
	for i := 0; i < 1000000; i++ {
		tree2.Put(data[i], data[i])
	}

	b.Run("indextree", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tree1.Get(data[i%1000000])
		}
	})

	b.Run("avl", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tree2.Get(data[i%1000000])
		}
	})
}

func BenchmarkRemove(b *testing.B) {
	data := newBenchDataWithSeed(1000000, 12345)

	b.Run("indextree", func(b *testing.B) {
		b.StopTimer()
		tree := New(compare.Any[int64])
		for i := 0; i < 1000000; i++ {
			tree.Put(data[i], data[i])
		}
		b.StartTimer()
		for i := 0; i < b.N && i < 1000000; i++ {
			tree.Remove(data[i])
		}
	})

	b.Run("avl", func(b *testing.B) {
		b.StopTimer()
		tree := avl.New[int64, int64](compare.AnyEx[int64])
		for i := 0; i < 1000000; i++ {
			tree.Put(data[i], data[i])
		}
		b.StartTimer()
		for i := 0; i < b.N && i < 1000000; i++ {
			tree.Remove(data[i])
		}
	})
}

func BenchmarkPutStringKey(b *testing.B) {
	tree := New(compare.ArrayAny[[]byte])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put([]byte(string(rune(i))), i)
	}
}

func BenchmarkIndexOnly(b *testing.B) {
	data := newBenchDataWithSeed(1000000, 12345)

	tree := New(compare.Any[int64])
	for i := 0; i < 1000000; i++ {
		tree.Put(data[i], data[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Index(int64(i % 1000000))
	}
}

func BenchmarkIndexOfOnly(b *testing.B) {
	data := newBenchDataWithSeed(1000000, 12345)

	tree := New(compare.Any[int64])
	for i := 0; i < 1000000; i++ {
		tree.Put(data[i], data[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.IndexOf(data[i%1000000])
	}
}

func TestHeight(t *testing.T) {
	itree := New(compare.Any[int64])
	avltree := avl.New[int64, int64](compare.AnyEx[int64])

	for i := 0; i < 5000; i++ {
		v := rand.Int63()
		itree.Put(v, v)
		avltree.Put(v, v)

		if itree.Size() != int64(avltree.Size()) {
			t.Fatal("size mismatch")
		}

		h1 := itree.hight()
		h2 := int(avltree.Height())
		if h1 != h2 && h1 != h2+1 {
			t.Logf("height diff: indextree=%d avl=%d", h1, h2)
		}
	}
}

func TestPutAndIndexFromDataFile(t *testing.T) {
	var data []int64
	if testutils.LoadData("BenchmarkPut", data) {
		tree := New(compare.Any[int64])
		for _, v := range data {
			tree.Put(v, v)
		}
		tree.check()
	}
}
