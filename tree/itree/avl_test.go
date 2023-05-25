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

	func() {

		defer func() {
			if err := recover(); err != nil {
				if err.(string) != "out of index size: 100 index: 100" {
					log.Panic(err)
				}
			}
		}()

		tree.RemoveIndex(100)
	}()

	func() {

		tree := New[int, int](compare.AnyEx[int])
		for i := 0; i < 100; i++ {
			tree.Put(i, i)
		}

		tree.RemoveIndex(99)
		defer func() {
			if err := recover(); err != nil {
				if err.(string) != "out of index size: 99 index: 99" {
					log.Panic(err)
				}
			}
		}()

		tree.RemoveIndex(99)
	}()

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

func TestRemoveRange(t *testing.T) {
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

		for i := 0; i < 5; i++ {
			s := rand.Intn(100)
			e := s + rand.Intn(20) + 15
			// log.Println(tree.view())
			tree.RemoveRange(s, e)
			r1 := sort.Search(len(sarr), func(i int) bool { return sarr[i] >= s })
			r2 := sort.Search(len(sarr), func(i int) bool { return sarr[i] > e })
			sarr = append(sarr[0:r1], sarr[r2:]...)

			result1 := fmt.Sprintf("%v", sarr)
			result2 := fmt.Sprintf("%v", tree.Values())
			if result1 != result2 {
				t.Error(result1)
				t.Error(result2)
			}

			tree.check()
		}
	}
}

func TestRemoveRangeIndex(t *testing.T) {

	tree := New[int, int](compare.AnyEx[int])

	v := 0
	tree.Put(v, v)
	tree.RemoveRangeByIndex(0, 0)
	if tree.Size() != 0 {
		t.Error(tree.view())
	}

	for i := 0; i < 10; i++ {
		v := i
		tree.Put(v, v)
	}
	tree.RemoveRangeByIndex(-1, 20)
	if tree.Size() != 0 {
		t.Error()
	}

	for i := 0; i < 10; i++ {
		v := i
		tree.Put(v, v)
	}

	tree.RemoveRangeByIndex(0, tree.Size()-2)
	k := tree.Index(0)
	if tree.Size() != 1 || k == 0 {
		t.Error()
	}

	for i := 0; i < 10; i++ {
		v := i
		tree.Set(v, v)
	}
	tree.RemoveRangeByIndex(0, -10)
	if tree.Size() != 10 {
		t.Error()
	}
}

func TestRemoveRangeIndexForce(t *testing.T) {

	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		var priority []int
		tree1 := New[int, int](compare.AnyEx[int])
		tree2 := New[int, int](compare.AnyEx[int])

		for i := 0; i < 200; i += rand.Intn(8) + 1 {
			v := i
			tree1.Put(v, v)
			tree2.Put(v, v)
			priority = append(priority, v)
		}

		s := rand.Intn(tree1.Size())
		e := rand.Intn(tree1.Size())
		if s > e {
			s, e = e, s
		}

		size := tree1.Size()
		// log.Println(tree.debugString(true))

		tree1.RemoveRangeByIndex(s, e)
		skey := tree2.Index(s)
		ekey := tree2.Index(e)
		tree2.RemoveRange(skey, ekey)
		priority = append(priority[0:s], priority[e+1:]...)

		if int(tree1.Size()) != len(priority) {
			log.Panic(tree1.Size(), len(priority))
		}
		if e-s+1 != size-tree1.Size() && tree1.Size() != tree2.Size() {
			log.Panic(e, s, tree1.Size(), size)
		}

		for i := 0; i < 200; i += rand.Intn(8) + 1 {
			v := i
			tree1.Put(v, v)
			tree2.Put(v, v)
		}

		tree1.check()
		tree2.check()
	}
}

func TestSimpleForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree1 := New[int, int](compare.AnyEx[int])
		tree2 := make(map[int]int)

		for i := 0; i < 40; i++ {
			v := rand.Intn(100)
			tree1.Put(v, v)
			tree2[v] = v
		}

		for k, v2 := range tree2 {
			if v1, ok := tree1.Get(k); !ok || v1 != v2 {
				panic("")
			}
		}

		for len(tree2) != 0 {
			i := rand.Intn(int(tree1.Size()))

			v1 := tree1.Index(i)
			if v2, ok := tree2[v1]; !ok || v1 != v2 {
				panic("")
			}

			if v3, ok := tree1.Get(v1); !ok || v1 != v3 {
				panic("")
			}

			if rand.Intn(2) == 0 {
				tree1.Remove(v1)
				delete(tree2, v1)
			} else {
				tree1.RemoveIndex(i)
				delete(tree2, v1)
			}

			if rand.Intn(2) == 0 {
				v := rand.Intn(100)
				if rand.Intn(2) == 0 {
					tree1.Put(v, v)
				} else {
					tree1.Set(v, v)
				}
				tree2[v] = v
			}

			tree1.check()
		}

	}
}

func TestSplit(t *testing.T) {
	rand := random.New(1684961987626654826)

	for n := 0; n < 10000; n++ {

		tree1 := New[int, int](compare.AnyEx[int])
		var priority []int
		for i := 0; i < 80; i++ {
			v := rand.Intn(200)
			if tree1.Put(v, v) {
				priority = append(priority, v)
			}

		}

		srcSize := tree1.Size()

		sort.Ints(priority)

		skey := rand.Intn(220)

		// idx := sort.Search(len(priority), func(i int) bool {
		// 	return priority[i] <= skey
		// })
		// if idx < len(priority) {
		// 	if priority[idx] != skey {
		// 		idx = idx + 1
		// 	}
		// }
		// log.Println(priority[idx])

		var idx = -1
		for i, v := range priority {
			if skey < v {
				break
			}
			idx = i
		}

		// log.Println(skey, priority[idx])
		tree2 := tree1.Split(skey)
		tree1.check()
		tree2.check()

		if tree1.Size()+tree2.Size() != srcSize {
			log.Println(tree1.view(), tree2.view(), skey)
			panic("tree1 size != tree2 size")
		}

		// log.Println(tree1.Values(), tree2.Values(), skey, idx)
		idx = idx + 1
		for i, v := range priority[0:idx] {
			if v1 := tree1.Index(i); v1 != v {
				log.Println(priority[0:idx], tree1.Values())
				log.Panicln(v1, v)
			}
		}
		// log.Println(tree2.view())
		for i, v := range priority[idx:] {
			// log.Println(priority[idx])
			if v2 := tree2.Index(i); v2 != v {
				log.Println(tree2.Size())
				log.Println(tree1.view(), tree2.view(), skey)
				log.Println(priority[idx:], tree2.Values())
				log.Panicln(v2, v)
			}
		}
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
