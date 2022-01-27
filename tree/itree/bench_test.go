package indextree

import (
	"log"
	"math/rand"
	"strconv"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
)

// var data []int64 = func() []int64 {
// 	var r []int64
// 	for i := 0; i < 2000000; i++ {
// 		r = append(r, rand.Int63())
// 	}
// 	return r
// }()

func BenchmarkPut(b *testing.B) {

	tree := New(compare.Int64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := rand.Int63()
		tree.Put(v, v)
	}
	b.Log(tree.Size())
}

func BenchmarkPut2(b *testing.B) {
	b.ResetTimer()
	tree := New(compare.Bytes)

	for i := 0; i < b.N; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

}

func BenchmarkAvlPut(b *testing.B) {

	tree := avl.New(compare.Int64)
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
		tree := New(compare.Int)
		for i := 0; i < 1000; i++ {
			v := r.Intn(100)
			tree.Put(v, v)
			// t.Error(tree.debugString(false))
		}
	}

}

func estDiffHight(t *testing.T) {
	tree := New(compare.Int64)
	avltree := avl.New(compare.Int64)

	for n := 0; n < 100000; n++ {
		for i := 0; i < 1000; i++ {
			v := rand.Int63n(3000)
			avltree.Put(v, v)
			tree.Put(v, v)
		}

		if avltree.Height()-tree.hight() > 1 {
			log.Println(avltree.Height() - tree.hight())
		}

		tree.Clear()
		avltree.Clear()
	}

}
