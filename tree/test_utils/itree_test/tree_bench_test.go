package indextreetest

import (
	"testing"

	"github.com/474420502/structure/compare"
)

const Level0 = 100000
const Level1 = 1000000
const Level2 = 5000000
const Level3 = 10000000
const Level4 = 50000000
const Level5 = 100000000

func init() {
	// debug.SetGCPercent(800)
}

func BenchmarkPutEx(b *testing.B) {
	b.StopTimer()
	tree := New(compare.Int)
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Put(i, i)
	}
}

func BenchmarkIndex(b *testing.B) {
	b.StopTimer()
	tree := New(compare.Int)

	for i := 0; i < Level3; i++ {
		tree.Put(i, i)
	}

	b.ResetTimer()
	b.StartTimer()

	b.N = Level3
	for i := 0; i < Level3; i++ {
		tree.Index(int64(i))
	}
}

func BenchmarkOp(b *testing.B) {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var x int
	for i := 0; i < b.N; i++ {
		x = a[9]
	}
	b.Log(x)
}
