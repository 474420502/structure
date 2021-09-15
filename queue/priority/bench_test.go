package treequeue

import (
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
)

func BenchmarkPut(b *testing.B) {
	tree := New(compare.Int64)
	for i := 0; i < b.N; i++ {
		v := rand.Int63()
		tree.Put(v, v)
	}
}

func BenchmarkAvlPut(b *testing.B) {
	tree := avl.New(compare.Int64)
	for i := 0; i < b.N; i++ {
		v := rand.Int63()
		tree.Cover(v, v)
	}
}
