package indextree

import (
	"log"
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestGet(t *testing.T) {
	tree := New(compare.Int)
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
		if _, ok := tree.Get(i); !ok {
			t.Error("not ok", i)
		}
	}
}

func TestIndex(t *testing.T) {
	tree := New(compare.Int)
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
		// log.Println(tree.debugString(true))
	}

	// log.Println(tree.debugString(true))

	for i := 0; i < 100; i++ {
		if _, v := tree.Index(int64(i)); v != i {
			t.Error("index error", v, i)
		}
	}

}

func TestRank(t *testing.T) {
	tree := New(compare.Int)
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
		// log.Println(tree.debugString(true))
	}

	// log.Println(tree.debugString(true))

	for i := 0; i < 100; i++ {
		if v := tree.IndexOf(i); v != int64(i) {
			t.Error("index error", i, "rank", v)
		}
	}
	tree.IndexOf(100)
}

func TestRemove1(t *testing.T) {
	tree := New(compare.Int)
	for _, i := range testutils.TestedArray {
		if !tree.Put(i, i) {
			log.Println("equal key", i)
		}
	}

	for _, v := range tree.Values() {
		if tree.Remove(v) != v {
			t.Error("remove error check it")
		}
	}

	if tree.Size() != 0 {
		t.Error(tree.Values())
	}
}

func TestRemove2(t *testing.T) {
	tree := New(compare.Int)
	for _, i := range testutils.TestedBigArray {
		if !tree.Put(i, i) {
			log.Println("equal key", i)
		}
	}

	if tree.Size() != int64(len(testutils.TestedBigArray)-4) {
		t.Error(tree.Size(), tree.Values())
	}

	for _, v := range tree.Values() {
		tree.Remove(v)
	}

	if tree.Size() != 0 {
		t.Error(tree.Values())
	}
}

func TestRemove3(t *testing.T) {
	tree := New(compare.Int)
	for n := 0; n < 1000; n++ {
		tree.Clear()
		for i := 0; i < 100; i += rand.Intn(3) + 1 {
			tree.Put(i, i)
			// log.Println(tree.debugString(true))
		}

		for i := 0; i < 10; i += rand.Intn(3) + 1 {
			v := rand.Intn(100)
			if _, ok := tree.Get(v); ok {
				if tree.Remove(v) == nil {
					t.Error("remove error")
				}
			} else {
				if tree.Remove(v) != nil {
					t.Error("remove error")
				}
			}
		}

	}

}
