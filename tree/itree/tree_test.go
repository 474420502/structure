package indextree

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"testing"
	"time"

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

func TestIndexForce(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		tree := New(compare.Int)

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

func TestRankForce(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		tree := New(compare.Int)
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

func TestRemoveRange(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	tree := New(compare.Int)
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
			sarr = append(sarr[0:r1], sarr[r2:len(sarr)]...)

			result1 := fmt.Sprintf("%v", sarr)
			result2 := fmt.Sprintf("%v", tree.Values())
			if result1 != result2 {
				t.Error(result1)
				t.Error(result2)
			}
		}
	}
}

func TestTrim(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	tree := New(compare.Int)
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

	tree := New(compare.Int)

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
}

func TestRemoveRangeIndexForce(t *testing.T) {

	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {

		var priority []int
		tree1 := New(compare.Int)
		tree2 := New(compare.Int)

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

	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {

		tree1 := New(compare.Int)
		tree2 := New(compare.Int)
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

	tree := New(compare.Int)

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
	tree.Traverse(func(k interface{}, v interface{}) bool {
		result = append(result, k)
		return true
	})

	s := fmt.Sprintf("%v", result)
	if s != "[8 9]" {
		t.Error()
	}
}

func TestAllForce(t *testing.T) {

	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {

		tree1 := New(compare.Int)
		var dict map[int]int = make(map[int]int)

		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			dict[v] = i
			tree1.PutCover(v, i)

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
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		tree1 := New(compare.Int)
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
			if v2, ok := tree2[k.(int)]; !ok || v1 != v2 {
				panic("")
			}

			if v3, ok := tree1.Get(k); !ok || v1 != v3 {
				panic("")
			}

			if rand.Intn(2) == 0 {
				tree1.Remove(k)
				delete(tree2, k.(int))
			} else {
				tree1.RemoveIndex(int64(i))
				delete(tree2, k.(int))
			}

			if rand.Intn(2) == 0 {
				v := rand.Intn(100)
				if rand.Intn(2) == 0 {
					tree1.Put(v, v)
				} else {
					tree1.PutCover(v, v)
				}
				tree2[v] = v
			}

			tree1.check()
		}

	}
}
