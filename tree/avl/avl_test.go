package avl

import (
	"log"
	"testing"

	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestPutGet(t *testing.T) {
	tree := New(compare.Int)
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	// log.Println(tree.String())

	for i := 0; i < tree.Size(); i++ {
		if v, b := tree.Get(i); !b || v != i {
			t.Error("error", b, v)
		}
	}

	tree.Clear()
	for _, i := range testutils.TestedArray {
		tree.Put(i, i)
	}

	if tree.Size() != len(testutils.TestedArray) {
		t.Error(tree.Values())
	}

	vs := tree.Values()
	if vs[0] != 1 || vs[tree.Size()-1] != 99 {
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

	if tree.Size() != len(testutils.TestedBigArray)-4 {
		t.Error(tree.Size(), tree.Values())
	}

	for _, v := range tree.Values() {
		tree.Remove(v)
	}

	if tree.Size() != 0 {
		t.Error(tree.Values())
	}
}

func TestRemove1(t *testing.T) {
	tree := New(compare.Int)
	for _, i := range testutils.TestedArray {
		if !tree.Put(i, i) {
			log.Println("equal key", i)
		}
	}

	if tree.Size() != len(testutils.TestedArray) {
		t.Error(tree.Size(), tree.Values())
	}

	// log.Println(tree.debugString())
	for _, v := range tree.Values() {
		tree.Remove(v)

		log.Println(tree.debugString())
	}

	if tree.Size() != 0 {
		t.Error(tree.Values())
	}
}
