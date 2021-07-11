package indextree

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestIndex(t *testing.T) {
	tree := New(compare.Int)
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
		// log.Println(tree.debugString(true))
	}

	// log.Println(tree.debugString(true))

	for i := 0; i < 10; i++ {
		if _, v := tree.Index(int64(i)); v != i {
			t.Error("index error", v, i)
		}
	}

	t.Error(tree.Index(10))
}
