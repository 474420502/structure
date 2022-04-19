package heap

import (
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
)

func BenchmarkPut(b *testing.B) {
	h := New(compare.Any[int64])
	for i := 0; i < b.N; i++ {
		v := rand.Int63()
		h.Put(v)
	}
}

func BenchmarkPop(b *testing.B) {
	b.StopTimer()
	h := New(compare.Any[int64])
	for i := 0; i < b.N; i++ {
		v := rand.Int63()
		h.Put(v)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h.Pop()
	}
}
