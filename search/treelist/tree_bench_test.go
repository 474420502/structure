package treelist

import (
	"strconv"
	"testing"
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

func BenchmarkPut(b *testing.B) {
	b.StopTimer()
	tree := New()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}
}

func BenchmarkIndex(b *testing.B) {
	b.StopTimer()
	tree := New()

	for i := 0; i < Level3; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

	b.ResetTimer()
	b.StartTimer()

	b.N = Level3
	for i := 0; i < Level3; i++ {
		tree.Index(int64(i))
	}
}
