package heap

import (
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
)

func BenchmarkPut(b *testing.B) {
	tree := New(compare.CompareAny[int64])
	for i := 0; i < b.N; i++ {
		v := rand.Int63()
		tree.Put(v)
	}
}
