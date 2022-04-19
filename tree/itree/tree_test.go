package indextree

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestGet(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
		if _, ok := tree.Get(i); !ok {
			t.Error("not ok", i)
		}
	}
}

func TestIndexForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree := New(compare.Any[int])

		var arr []int = make([]int, 0, 50)
		for i := 0; i < 50; i++ {
			r := rand.Intn(1000)

			if tree.Put(r, r) {
				arr = append(arr, r)
			}
		}
		sort.Ints(arr)
		for i, v := range arr {
			key, _ := tree.Index(int64(i))
			if key != v {
				t.Error("error ", v, key)
			}
		}
	}

}

func TestIndex(t *testing.T) {
	tree := New(compare.Any[int])
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

func TestRankForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree := New(compare.Any[int])
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
			if idx != int64(i) {
				t.Error("error ", i, idx)
			}
		}
	}
}

func TestRank(t *testing.T) {
	tree := New(compare.Any[int])
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
	tree := New(compare.Any[int])
	for _, i := range testutils.TestedArray {
		if !tree.Put(i, i) {
			log.Println("equal key", i)
		}
	}

	for _, v := range tree.Values() {
		if tree.Remove(v.(int)) != v {
			t.Error("remove error check it")
		}
	}

	if tree.Size() != 0 {
		t.Error(tree.Values())
	}
}

func TestRemove2(t *testing.T) {
	tree := New(compare.Any[int])
	for _, i := range testutils.TestedBigArray {
		if !tree.Put(i, i) {
			log.Println("equal key", i)
		}
	}

	if tree.Size() != int64(len(testutils.TestedBigArray)-4) {
		t.Error(tree.Size(), tree.Values())
	}

	for _, v := range tree.Values() {
		tree.Remove(v.(int))
	}

	if tree.Size() != 0 {
		t.Error(tree.Values())
	}
}

func TestRemove3(t *testing.T) {
	tree := New(compare.Any[int])
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

func TestRemoveRange(t *testing.T) {
	rand := random.New(t.Name())
	tree := New(compare.Any[int])
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

func TestTrim(t *testing.T) {
	rand := random.New(t.Name())
	tree := New(compare.Any[int])
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

func TestRemoveRangeIndex(t *testing.T) {

	tree := New(compare.Any[int])

	v := 0
	tree.Put(v, v)
	tree.RemoveRangeByIndex(0, 0)
	if tree.Size() != 0 {
		t.Error()
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
	k, _ := tree.Index(0)
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
		tree1 := New(compare.Any[int])
		tree2 := New(compare.Any[int])

		for i := 0; i < 200; i += rand.Intn(8) + 1 {
			v := i
			tree1.Put(v, v)
			tree2.Put(v, v)
			priority = append(priority, v)
		}

		s := rand.Int63n(tree1.Size())
		e := rand.Int63n(tree1.Size())
		if s > e {
			s, e = e, s
		}

		size := tree1.Size()
		// log.Println(tree.debugString(true))

		tree1.RemoveRangeByIndex(s, e)
		skey, _ := tree2.Index(s)
		ekey, _ := tree2.Index(e)
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

func TestTrimIndexForce(t *testing.T) {

	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		tree1 := New(compare.Any[int])
		tree2 := New(compare.Any[int])
		var priority []int

		for i := 0; i < 200; i += rand.Intn(4) + 1 {
			v := i
			tree1.Put(v, v)
			tree2.Put(v, v)
			priority = append(priority, v)
		}

		s := rand.Int63n(tree1.Size())
		e := rand.Int63n(tree1.Size())
		if s > e {
			s, e = e, s
		}

		size := tree1.Size()

		tree1.TrimByIndex(s, e)
		skey := tree2.index(s).Key
		ekey := tree2.index(e).Key
		tree2.Trim(skey, ekey)
		priority = priority[s : e+1]

		if int(tree1.Size()) != len(priority) {
			log.Panic(tree1.Size(), len(priority))
		}
		if e-s+1 != tree1.Size() && tree1.Size() != tree2.Size() {
			log.Panic(e, s, tree1.Size(), size)
		}

		tree1.check()
		tree2.check()
		// log.Println()
	}
}

func TestTrimIndex(t *testing.T) {

	tree := New(compare.Any[int])

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
	tree.Traverse(func(k int, v interface{}) bool {
		result = append(result, k)
		return true
	})

	s := fmt.Sprintf("%v", result)
	if s != "[8 9]" {
		t.Error()
	}
}

func TestAllForce(t *testing.T) {

	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		tree1 := New(compare.Any[int])
		var dict map[int]int = make(map[int]int)

		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			dict[v] = i
			tree1.Set(v, i)

		}

		for k, v := range dict {
			if r, ok := tree1.Get(k); !ok || r != v {
				panic("")
			}
		}

		tree1.check()

		// log.Println()
	}
}

func TestSimpleForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree1 := New(compare.Any[int])
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

			k, v1 := tree1.Index(int64(i))
			if v2, ok := tree2[k]; !ok || v1 != v2 {
				panic("")
			}

			if v3, ok := tree1.Get(k); !ok || v1 != v3 {
				panic("")
			}

			if rand.Intn(2) == 0 {
				tree1.Remove(k)
				delete(tree2, k)
			} else {
				tree1.RemoveIndex(int64(i))
				delete(tree2, k)
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

func TestSplitContain(t *testing.T) {
	rand := random.New(t.Name())
	var showlist []interface{} = make([]interface{}, 2)
	defer func() {
		if err := recover(); err != nil {
			log.Panicln(showlist...)
		}
	}()

	for n := 0; n < 1000; n++ {

		tree1 := New(compare.Any[int])
		var priority []int
		for i := 0; i < 80; i++ {
			v := rand.Intn(200)
			if tree1.Put(v, v) {
				priority = append(priority, v)
			}

			// log.Println()
		}

		sort.Ints(priority)

		skey := rand.Intn(220)
		// log.Println("current n:", n)
		// if n == 5 {
		// 	log.Println()
		// }
		// showlist[0] = tree1.debugString(true)
		// showlist[1] = skey
		var idx = -1
		for i, v := range priority {
			if skey < v {
				break
			}
			idx = i
		}
		// log.Println(tree1.debugString(true))

		tree2 := tree1.SplitContain(skey)
		tree1.check()
		tree2.check()

		// log.Println(priority, skey, "idx:", idx)
		// log.Println(tree1.Values(), tree2.Values())
		// log.Println(tree1.debugString(true), tree2.debugString(true))

		idx = idx + 1
		for i, v := range priority[0:idx] {
			if v1, _ := tree1.Index(int64(i)); v1 != v {
				log.Println(priority[0:idx], tree1.Values())
				log.Panicln(v1, v)
			}
		}

		for i, v := range priority[idx:] {
			if v2, _ := tree2.Index(int64(i)); v2 != v {
				log.Println(priority[idx:], tree2.Values())
				log.Panicln(v2, v)
			}
		}
	}

}

func TestSplit(t *testing.T) {
	rand := random.New(t.Name())
	var showlist []interface{} = make([]interface{}, 2)
	defer func() {
		if err := recover(); err != nil {
			log.Panicln(showlist...)
		}
	}()

	for n := 0; n < 1000; n++ {

		tree1 := New(compare.Any[int])
		var priority []int
		for i := 0; i < 80; i++ {
			v := rand.Intn(200)
			if tree1.Put(v, v) {
				priority = append(priority, v)
			}

		}

		sort.Ints(priority)

		skey := rand.Intn(220)

		var idx = -1
		for i, v := range priority {
			if skey <= v {
				break
			}
			idx = i
		}

		tree2 := tree1.Split(skey)
		tree1.check()
		tree2.check()

		// log.Println(tree1.Values(), tree2.Values(), skey, idx)
		idx = idx + 1
		for i, v := range priority[0:idx] {
			if v1, _ := tree1.Index(int64(i)); v1 != v {
				log.Println(priority[0:idx], tree1.Values())
				log.Panicln(v1, v)
			}
		}

		for i, v := range priority[idx:] {
			if v2, _ := tree2.Index(int64(i)); v2 != v {
				log.Println(priority[idx:], tree2.Values())
				log.Panicln(v2, v)
			}
		}
	}

}

func TestCase(t *testing.T) {
	tree1 := New(compare.Any[int])

	for i := 0; i < 100; i++ {
		v := rand.Intn(1000)
		tree1.Put(v, v)
	}
	// log.Println(tree1.debugString(false))
}
