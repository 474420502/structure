package arraytree

import (
	"log"
	"testing"

	"github.com/474420502/structure/compare"
	indextree "github.com/474420502/structure/tree/itree"
)

func TestCase(t *testing.T) {
	tree := New()
	tl := indextree.New(compare.Int)
	for i := 0; i < 1<<3-1; i++ {
		tl.Put(i, i)
		log.Println(tl.String())
	}

	for i := 0; i < 1000; i++ {
		tree.Put(i, i)

	}
}
