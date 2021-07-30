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
	log.Println(seed)
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
	log.Println(seed)
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
	log.Println(seed)
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
	log.Println(seed)
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
