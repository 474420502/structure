package indextree

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
	testutils "github.com/474420502/structure/tree/test_utils"
)

// var data []int64 = func() []int64 {
// 	var r []int64
// 	for i := 0; i < 2000000; i++ {
// 		r = append(r, rand.Int63())
// 	}
// 	return r
// }()

type DefaultCompareType interface {
	int | int64 | int32 | int8 | float32 | float64 | uint8 | uint | uint32 | uint64
}

func CompareAny[T DefaultCompareType](k1, k2 T) int {

	switch {
	case k1 > k2:
		return 1
	case k1 < k2:
		return -1
	default:
		return 0
	}
}

func BenchmarkPut(b *testing.B) {
	var data []int64

	if !testutils.LoadData("BenchmarkPut", data) {
		for i := 0; i < 4000000; i++ {
			v := rand.Int63()
			data = append(data, v)
		}
		testutils.SaveData("BenchmarkPut", data)
	}

	b.ResetTimer()

	rand.Seed(time.Now().Unix())
	start := int(rand.Int63n(500000))

	b.Run("pre", func(b *testing.B) {
		tree := avl.New(CompareAny[int64])

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v := data[i+start]
			tree.Put(v, v)
		}
		// b.Log(tree.Size())
	})

	b.Run("indextree", func(b *testing.B) {
		tree := New(CompareAny[int64])

		b.ResetTimer()

		// b.N = 100
		for i := 0; i < b.N; i++ {
			v := data[i+start]
			tree.Put(v, v)
		}

	})

	b.Run("avl", func(b *testing.B) {
		tree := avl.New(CompareAny[int64])

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v := data[i+start]
			tree.Put(v, v)
		}

	})

}

func TestHeight(t *testing.T) {
	itree := New(CompareAny[int64])
	avltree := avl.New(CompareAny[int64])

	var diffcount = 0
	for i := 0; i < 5000; i++ {
		v := rand.Int63()
		itree.Put(v, v)
		avltree.Put(v, v)

		if itree.Size() != avltree.Size() {
			log.Panic()
		}

		if h1, h2 := itree.hight(), avltree.Height(); math.Abs(float64(h1-h2)) >= 1 {
			diffcount++
			log.Println("height:", h1, h2, "diff:", diffcount, h1-h2)
		}
	}
}

func BenchmarkPut2(b *testing.B) {

	tree := New(compare.ArrayAny[[]byte])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}
}

func BenchmarkAvlPut(b *testing.B) {

	tree := avl.New(CompareAny[int64])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := rand.Int63()
		tree.Put(v, v)
	}
	b.Log(tree.Size())
}

func TestCase10(t *testing.T) {

	// for _, v := range []int{0, 131, 756, 459, 533} {
	// 	tree.Put(v, v)
	// 	t.Error(tree.debugString(false))
	// }
	r := random.New()
	for n := 0; n < 10000; n++ {
		tree := New(compare.Any[int])
		for i := 0; i < 1000; i++ {
			v := r.Intn(100)
			tree.Put(v, v)
			// t.Error(tree.debugString(false))
		}
	}

}

func estDiffHight(t *testing.T) {
	// tree := New(compare.Int64)
	// avltree := avl.New(compare.Int64)

	// for n := 0; n < 100000; n++ {
	// 	for i := 0; i < 1000; i++ {
	// 		v := rand.Int63n(3000)
	// 		avltree.Put(v, v)
	// 		tree.Put(v, v)
	// 	}

	// 	if avltree.Height()-tree.hight() > 1 {
	// 		log.Println(avltree.Height() - tree.hight())
	// 	}

	// 	tree.Clear()
	// 	avltree.Clear()
	// }

}
