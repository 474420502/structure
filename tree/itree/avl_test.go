package itree

import (
	"fmt"
	"log"
	"sort"
	"testing"

	random "github.com/474420502/random"
	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestPutGet(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
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
	tree := New[int, int](compare.AnyEx[int])
	for _, i := range testutils.TestedBigArray {
		if !tree.Set(i, i) {
			// log.Println("equal key", i)
		}
	}

	if int(tree.Size()) != len(testutils.TestedBigArray)-4 {
		t.Error(int(tree.Size()), tree.Values())
	}

	for _, v := range tree.Values() {
		tree.Remove(v)
	}

	if int(tree.Size()) != 0 {
		t.Error(tree.Values())
	}
}

func TestRemove1(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
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
		tree.Remove(v)
	}

	if int(tree.Size()) != 0 {
		t.Error(tree.Values())
	}
}

func TestForce(t *testing.T) {
	rand := random.New(t.Name())

	tree := New[int, int](compare.AnyEx[int])
	for n := 0; n < 5000; n++ {

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

		// tree.check()

		sort.Slice(priority, func(i, j int) bool {
			return priority[i] < priority[j]
		})

		for i, v := range tree.Values() {
			if priority[i] != v {
				log.Println(tree.view())
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
		tree.Traverse(func(k int, v int) bool {
			if priority[i] != v {
				panic("")
			}
			i++
			return true
		})

		tree.Clear()

	}
}

func TestIndex(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
		// log.Println(tree.debugString(true))
	}

	// log.Println(tree.debugString(true))

	for i := 0; i < 100; i++ {
		if v := tree.Index(i); v != i {
			t.Error("index error", v, i)
		}
	}

}

func TestRankForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree := New[int, int](compare.AnyEx[int])
		var arr []int = make([]int, 0, 50)
		for i := 0; i < 50; i++ {
			r := rand.Intn(1000)
			if tree.Put(r, r) {
				arr = append(arr, r)
			}
		}
		sort.Ints(arr)
		for i, v := range arr {
			idx := tree.IndexOf(v)
			if idx != i {
				t.Error("error ", i, idx)
			}
		}
	}
}

func TestRank(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
		// log.Println(tree.debugString(true))
	}

	// log.Println(tree.debugString(true))

	for i := 0; i < 100; i++ {
		if v := tree.IndexOf(i); v != i {
			t.Error("index error", i, "rank", v)
		}
	}
	tree.IndexOf(100)
}

func TestTrim(t *testing.T) {
	rand := random.New(t.Name())
	tree := New[int, int](compare.AnyEx[int])
	for n := 0; n < 2000; n++ {
		tree.Clear()
		var sarr []int
		for i := 0; i < 100; i += rand.Intn(3) + 1 {
			tree.Put(i, i)
			sarr = append(sarr, i)
		}

		sort.Ints(sarr)

		// log.Println(tree.debugString(true))

		for i := 0; i < 5; i++ {
			if len(sarr) == 0 {
				return
			}
			s := rand.Intn(len(sarr))
			e := rand.Intn(len(sarr))
			if s > e {
				s, e = e, s
			}
			r1 := sarr[s]
			r2 := sarr[e]

			sarr = sarr[s:e]
			sarr = append(sarr, r2)
			tree.Trim(r1, r2)
			// log.Println(tree.debugString(true))

			result1 := fmt.Sprintf("%v", sarr)
			result2 := fmt.Sprintf("%v", tree.Values())
			if result1 != result2 {
				t.Error(result1)
				t.Error(result2)
			}
		}
	}
}

func TestTrimIndex(t *testing.T) {

	tree := New[int, int](compare.AnyEx[int])

	tree.Put(0, 0)
	tree.TrimByIndex(0, 0)
	if tree.Size() != 1 {
		t.Error()
	}

	for i := 0; i < 10; i++ {
		v := i
		tree.Put(v, v)
	}

	if tree.Size() != 10 {
		t.Error()
	}
	tree.TrimByIndex(8, 9)
	if tree.Size() != 2 {
		t.Error()
	}
	if tree.IndexOf(8) != 0 {
		t.Error()
	}
	if tree.IndexOf(9) != 1 {
		t.Error()
	}

	var result []interface{}
	tree.Traverse(func(k int, v int) bool {
		result = append(result, k)
		return true
	})

	s := fmt.Sprintf("%v", result)
	if s != "[8 9]" {
		t.Error()
	}
}

func BenchmarkIndex(b *testing.B) {
	b.StopTimer()
	tree := New[int, int](compare.AnyEx[int])
	for i := 0; i < 1000; i++ {
		tree.Put(i, i)
	}
	b.StartTimer()

	r := random.New()
	s := tree.Size()
	for i := 0; i < b.N; i++ {
		tree.Index(r.Intn(s))
	}
}

func BenchmarkPut(b *testing.B) {
	rand := random.New(1683721792150515321)
	tree := New[int, int](compare.AnyEx[int])
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
	// b.Log(tree.rotateCount)
}

func BenchmarkRemove(b *testing.B) {
	rand := random.New(1683721792150515321)
	tree := New[int, int](compare.AnyEx[int])
	var removelist []int
	var ri = 0

	for i := 0; i < b.N; i++ {
		if tree.Size() == 0 {
			b.StopTimer()
			removelist = nil
			ri = 0
			for i := 0; i < 1000; i++ {
				v := rand.Intn(1000)
				if tree.Put(v, v) {
					removelist = append(removelist, v)
				}

				if i%25 == 0 {
					removelist = append(removelist, rand.Intn(1000))
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
