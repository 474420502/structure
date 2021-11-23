package avl

import (
	"log"
	"sort"
	"testing"

	random "github.com/474420502/random"
	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestCase(t *testing.T) {
	tree := New(compare.Int)
	r := random.New()
	for i := 0; i < 100; i++ {
		v := r.Intn(100)
		tree.Put(v, v)
	}
}

func TestPutGet(t *testing.T) {
	tree := New(compare.Int)
	for i := 0; i < 100; i++ {
		tree.Set(i, i)
	}

	// log.Println(tree.String())

	for i := 0; i < tree.Size(); i++ {
		if v, b := tree.Get(i); !b || v != i {
			t.Error("error", b, v)
		}
	}

	tree.Clear()
	for _, i := range testutils.TestedArray {
		tree.Set(i, i)
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
		if !tree.Set(i, i) {
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
		if !tree.Set(i, i) {
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

func TestForce(t *testing.T) {
	rand := random.New(t.Name())

	tree := New(compare.Int)
	for n := 0; n < 2000; n++ {

		var priority []int
		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			if tree.Put(v, v) {
				priority = append(priority, v)
			}
		}

		if tree.Size() != len(priority) {
			panic("")
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i] < priority[j]
		})

		for i, v := range tree.Values() {
			if priority[i] != v {
				panic("")
			}
		}

		for i := 0; i < 40; i++ {

			v := rand.Intn(100)

			if _, ok := tree.Get(v); ok {

				rv, ok := tree.Remove(v)
				if !ok || rv != v {
					panic("")
				}

				if idx := sort.SearchInts(priority, v); idx == len(priority) {
					panic("")
				} else {
					priority = append(priority[:idx], priority[idx+1:]...)
				}

			}
		}

		var i = 0
		tree.Traverse(func(k, v interface{}) bool {
			if priority[i] != v {
				panic("")
			}
			i++
			return true
		})

		tree.Clear()

	}
}
