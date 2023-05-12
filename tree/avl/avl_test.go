package avl

import (
	"log"
	"sort"
	"testing"

	random "github.com/474420502/random"
	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestPutGet(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Set(i, i)
	}

	// log.Println(tree.String())

	for i := 0; i < int(tree.Size()); i++ {
		if v, b := tree.Get(i); !b || v != i {
			t.Error("error", b, v)
		}
	}

	tree.Clear()
	for _, i := range testutils.TestedArray {
		tree.Set(i, i)
	}

	if int(tree.Size()) != len(testutils.TestedArray) {
		t.Error(tree.Values())
	}

	vs := tree.Values()
	if vs[0] != 1 || vs[int(tree.Size())-1] != 99 {
		t.Error(tree.Values())
	}
}

func TestRemove2(t *testing.T) {
	tree := New(compare.Any[int])
	for _, i := range testutils.TestedBigArray {
		if !tree.Set(i, i) {
			// log.Println("equal key", i)
		}
	}

	if int(tree.Size()) != len(testutils.TestedBigArray)-4 {
		t.Error(int(tree.Size()), tree.Values())
	}

	for _, v := range tree.Values() {
		tree.Remove(v.(int))
	}

	if int(tree.Size()) != 0 {
		t.Error(tree.Values())
	}
}

func TestRemove1(t *testing.T) {
	tree := New(compare.Any[int])
	for _, i := range testutils.TestedArray {
		if !tree.Set(i, i) {
			// log.Println("equal key", i)
		}
	}

	if int(int(tree.Size())) != len(testutils.TestedArray) {
		t.Error(int(tree.Size()), tree.Values())
	}

	// log.Println(tree.debugString())
	for _, v := range tree.Values() {
		tree.Remove(v.(int))
		log.Println(tree.debugString())
	}

	if int(tree.Size()) != 0 {
		t.Error(tree.Values())
	}
}

func TestForce(t *testing.T) {
	rand := random.New(t.Name())

	tree := New(compare.Any[int])
	for n := 0; n < 2000; n++ {

		var priority []int
		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			if tree.Put(v, v) {
				priority = append(priority, v)
			}
		}

		if int(tree.Size()) != len(priority) {
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
		tree.Traverse(func(k int, v interface{}) bool {
			if priority[i] != v {
				panic("")
			}
			i++
			return true
		})

		tree.Clear()

	}
}

// func TestCaseX(t *testing.T) {
// 	New(compare.TimeDesc[time.Time])
// }

func BenchmarkPut(b *testing.B) {
	rand := random.New(1683721792150515321)

	tree := New(compare.Any[int])
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		v := rand.Int()
		tree.Put(v, v)
		// tree.check()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v := rand.Int()
		tree.Put(v, v)
	}
}

func BenchmarkRemove(b *testing.B) {
	rand := random.New(1683721792150515321)
	tree := New(compare.Any[int])
	var removelist []int
	var ri = 0

	for i := 0; i < b.N; i++ {
		if tree.Size() == 0 {
			b.StopTimer()
			removelist = nil
			ri = 0
			for i := 0; i < 100; i++ {
				v := rand.Intn(100)
				if tree.Put(v, v) {
					removelist = append(removelist, v)
				}

				if i%25 == 0 {
					removelist = append(removelist, rand.Intn(100))
				}
				// tree.check()
			}
			b.StartTimer()
		}

		v := removelist[ri]
		tree.Remove(v)
		ri += 1
	}

}
