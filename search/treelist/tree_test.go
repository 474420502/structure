package treelist

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strconv"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
	avl "github.com/474420502/structure/tree/avl"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestIndexForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree := New()
		tree.compare = compare.BytesLen
		var arr []int = make([]int, 0, 50)
		for i := 0; i < 50; i++ {
			r := rand.Intn(1000)

			v := []byte(strconv.Itoa(r))
			if tree.Put(v, v) {
				arr = append(arr, r)
			}
		}
		sort.Ints(arr)
		for i, v := range arr {
			vv := []byte(strconv.Itoa(v))
			s := tree.Index(int64(i))
			if bytes.Compare(s.Key, vv) != 0 {
				t.Error("error ", string(vv), string(s.Key))
			}
		}
	}

}

func TestIndex(t *testing.T) {
	tree := New()
	tree.compare = compare.BytesLen
	for i := 0; i < 100; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)

		s := tree.Index(int64(i))
		if bytes.Compare(s.Key, v) != 0 {
			t.Error(s, v)
		}
		iter := tree.Iterator()
		iter.Seek(v)
		if iter.Index() != int64(i) {
			t.Error("iterator index error")
		}
		if bytes.Compare(iter.Key(), v) != 0 {
			t.Error("iterator key error")
		}
	}

}

func TestRank(t *testing.T) {
	tree := New()
	tree.compare = compare.BytesLen

	for i := 0; i < 100; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

	for i := 0; i < 100; i++ {
		k := []byte(strconv.Itoa(i))
		if v := tree.IndexOf(k); v != int64(i) {
			t.Error("index error", i, "rank", v)
		}
	}

}

func TestRemove1(t *testing.T) {
	tree := New()
	for _, i := range testutils.TestedArray {
		v := []byte(strconv.Itoa(i))
		if !tree.Put(v, v) {
			log.Println("equal key", i)
		}
	}

	for _, v := range tree.Slices() {
		r := tree.Remove(v.Key)
		if bytes.Compare(r.Key, v.Key) != 0 {
			log.Println("remove error check it", string(r.Key), string(v.Key))
		}
	}

	if tree.Size() != 0 {
		t.Error(tree.Slices())
	}
}

func TestRemove2(t *testing.T) {
	tree := New()
	for _, i := range testutils.TestedBigArray {
		v := []byte(strconv.Itoa(i))
		if !tree.Put(v, v) {
			log.Println("equal key", i)
		}
	}

	if tree.Size() != int64(len(testutils.TestedBigArray)-4) {
		t.Error(tree.Size(), tree.Slices())
	}

	for _, v := range tree.Slices() {
		tree.Remove(v.Key)
	}

	if tree.Size() != 0 {
		t.Error(tree.Slices())
	}
}

func TestRemoveNode(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 1000; n++ {

		tree := New()
		tree.compare = compare.BytesLen

		var dmap map[int]int = make(map[int]int)

		for i := 0; i < 200; i += rand.Intn(4) + 1 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			dmap[i] = i
		}

		for i := 0; i < 200; i += rand.Intn(4) + 1 {
			v := []byte(strconv.Itoa(i))
			if _, ok := dmap[i]; ok {
				r := tree.Remove(v)
				if r == nil {
					t.Error(r)
				}
				if bytes.Compare(r.Key, v) != 0 {
					t.Error()
				}
			} else {
				r := tree.Remove(v)
				if r != nil {
					t.Error(r)
				}
			}
		}
	}
}

func TestRemoveForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 1000; n++ {

		tree := New()
		tree.compare = compare.BytesLen

		var priority [][]byte

		for i := 0; i < 500; i += rand.Intn(8) + 1 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			priority = append(priority, v)
		}

		sort.Slice(priority, func(i, j int) bool {
			return tree.compare(priority[i], priority[j]) < 0
		})

		for tree.Size() != 0 {

			i := rand.Intn(int(tree.Size()))
			tree.RemoveIndex(int64(i))
			priority = append(priority[0:i], priority[i+1:]...)

			if tree.Size() != 0 && rand.Intn(2) == 0 {
				s := rand.Int63n(tree.Size())
				e := rand.Int63n(tree.Size())
				if s > e {
					s, e = e, s
				}
				tree.RemoveRangeByIndex(s, e)
				priority = append(priority[0:s], priority[e+1:]...)
			}

			var idx = 0
			tree.Traverse(func(s *Slice) bool {
				if tree.compare(s.Key, priority[idx]) != 0 {
					panic("")
				}
				idx++
				return true
			})

			if rand.Intn(2) == 0 {
				v := []byte(strconv.Itoa(rand.Intn(100)))
				if tree.Put(v, v) {
					priority = append(priority, v)
				}
				sort.Slice(priority, func(i, j int) bool {
					return tree.compare(priority[i], priority[j]) < 0
				})
			}

			tree.check()
		}
	}
}

func TestRange(t *testing.T) {
	// tree := New()
	// tree.compare = compare.BytesLen
	// for i := 0; i < 100; i += 4 {
	// 	v := []byte(strconv.Itoa(i))
	// 	tree.Put(v, v)
	// }
	// log.Println(tree.debugString(true))
	rand := random.New(t.Name())
	for n := 0; n < 1000; n++ {
		startkey := rand.Intn(200)
		endkey := rand.Intn(200)
		if startkey > endkey {
			temp := startkey
			startkey = endkey
			endkey = temp
		}
		tree := New()
		tree.compare = compare.BytesLen
		avltree := avl.New(compare.Int)

		for i := 0; i < 200; i += rand.Intn(8) + 2 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			avltree.Cover(i, i)
		}
		// tree.rcount = 0
		start := []byte(strconv.Itoa(startkey))
		end := []byte(strconv.Itoa(endkey))

		tree.RemoveRange(start, end)

		for i := startkey; i <= endkey; i++ {
			avltree.Remove(i)
		}

		if tree.Size() != int64(avltree.Size()) {
			t.Error(avltree.Height(), avltree.Size(), avltree)
			t.Error(tree.Size(), tree.debugString(true))
			return
		}

		avltree.Traverse(func(k, v interface{}) bool {
			key := []byte(strconv.Itoa(v.(int)))
			if _, ok := tree.Get(key); !ok {
				t.Error("tree is error")
			}
			return true
		})

		iter := tree.Iterator()
		iter.SeekToLast()
		root := tree.getRoot()
		if root == nil {
			if tree.root.Direct[0] != nil {
				log.Panicln(tree.root.Direct[0], tree.debugString(true))
			}
			if tree.root.Direct[1] != nil {
				log.Panicln(tree.root.Direct[1], tree.debugString(true))
			}
		} else {
			hand := root
			for hand.Children[1] != nil {
				hand = hand.Children[1]
			}
			if hand != iter.cur {
				log.Panicln(root, hand, iter.cur, tree.debugString(true))
			}
		}

		iter.SeekToFirst()
		if root != nil {
			hand := root
			for hand.Children[0] != nil {
				hand = hand.Children[0]
			}
			if hand != iter.cur {
				log.Panicln(hand, iter.cur, tree.debugString(true))
			}
		}

		for iter.Valid() {
			if _, ok := tree.Get(iter.Key()); !ok {
				log.Println("start:", startkey, "end:", endkey)
				// log.Println(srctree)
				log.Println(tree.debugString(true))
				log.Println("not ok", string(iter.Key()))
				panic("")
			}
			iter.Next()
		}
	}
}

func TestRangeReturn(t *testing.T) {

	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		startkey := rand.Intn(300)
		endkey := rand.Intn(300)
		if startkey > endkey {
			temp := startkey
			startkey = endkey
			endkey = temp
		}

		tree := New()
		tree.compare = compare.BytesLen
		avltree := avl.New(compare.Int)

		for i := 100; i < 200; i += rand.Intn(8) + 2 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			avltree.Cover(i, i)
		}
		// tree.rcount = 0
		start := []byte(strconv.Itoa(startkey))
		end := []byte(strconv.Itoa(endkey))

		size1 := int64(avltree.Size())
		size2 := tree.Size()

		result := tree.RemoveRange(start, end)
		for i := startkey; i <= endkey; i++ {
			avltree.Remove(i)
		}

		if !result {
			// log.Println(tree.debugString(true))
			if !(size1 == size2 && tree.Size() == size2 && int64(avltree.Size()) == size1) {
				t.Error("RemoveRange return error")
			}
		}
	}
}

func TestHeadTail(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 1000; n++ {

		tree := New()
		tree.compare = compare.BytesLen

		var min, max int
		for i := 0; i < 500; i += rand.Intn(8) + 1 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			if min > i {
				min = i
			}
			if max < i {
				max = i
			}

			if compare.BytesLen(tree.Head().Key, []byte(strconv.Itoa(min))) != 0 {
				t.Error("test the seed")
			}

			if compare.BytesLen(tree.Tail().Key, []byte(strconv.Itoa(max))) != 0 {
				t.Error("test the seed")
			}
		}
	}
}

func TestRemoveHeadTail(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 1000; n++ {

		tree := New()
		tree.compare = compare.BytesLen

		var min, max int
		for i := 0; i < rand.Intn(500); i += rand.Intn(4) + 1 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			if min > i {
				min = i
			}
			if max < i {
				max = i
			}
		}

		// log.Println(n)
		// if n == 306 {
		// 	log.Println(tree.Size())
		// }

		tree.RemoveHead()
		if s := tree.Size(); s > 0 {
			if compare.BytesLen(tree.Head().Key, []byte(strconv.Itoa(min))) == 0 {
				t.Error("test the seed")
			}

			if s == 1 {
				if compare.BytesLen(tree.Head().Key, tree.Tail().Key) != 0 {
					t.Error(n, "head is should be equal to tail")
				}
			}
		} else {
			if tree.Head() != nil && tree.Tail() != nil {
				t.Error("head tail should be nil", n)
			}

		}

		tree.RemoveTail()
		if s := tree.Size(); s > 0 {
			if compare.BytesLen(tree.Tail().Key, []byte(strconv.Itoa(max))) == 0 {
				t.Error("test the seed")
			}

			if s == 1 {
				if compare.BytesLen(tree.Head().Key, tree.Tail().Key) != 0 {
					t.Error(n, "head is should be equal to tail")
				}
			}
		} else {
			if tree.Head() != nil && tree.Tail() != nil {
				t.Error("head tail should be nil", n)
			}
		}

	}
}

func TestRemoveRangeIndex(t *testing.T) {

	tree := New()
	tree.compare = compare.BytesLen

	v := []byte(strconv.Itoa(0))
	tree.Put(v, v)
	tree.RemoveRangeByIndex(0, 0)
	if tree.Size() != 0 {
		t.Error()
	}

	for i := 0; i < 10; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}
	tree.RemoveRangeByIndex(-1, 20)
	if tree.Size() != 0 {
		t.Error()
	}

	for i := 0; i < 10; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

	tree.RemoveRangeByIndex(0, tree.Size()-2)
	if tree.Size() != 1 || string(tree.Index(0).Key) == "0" {
		t.Error()
	}
}

func TestRemoveRangeIndexForce(t *testing.T) {

	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		tree := New()
		tree.compare = compare.BytesLen
		tree2 := New()
		tree2.compare = compare.BytesLen

		for i := 0; i < 200; i += rand.Intn(8) + 1 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			tree2.Put(v, v)
		}

		s := rand.Int63n(tree.Size())
		e := rand.Int63n(tree.Size())
		if s > e {
			s, e = e, s
		}

		size := tree.Size()
		// log.Println(tree.debugString(true))

		tree.RemoveRangeByIndex(s, e)
		skey := tree2.index(s).Key
		ekey := tree2.Index(e).Key
		tree2.RemoveRange(skey, ekey)
		// log.Println(tree.debugString(true), s, e)
		if e-s+1 != size-tree.Size() && tree.Size() != tree2.Size() {
			log.Panic(e, s, tree.Size(), size)
		}

		iter1 := tree.Iterator()
		iter2 := tree2.Iterator()

		iter1.SeekToFirst()
		iter2.SeekToFirst()

		for iter1.Valid() {
			if tree.compare(iter1.Key(), iter2.Key()) != 0 {
				panic("")
			}
			iter1.Next()
			iter2.Next()
		}

		iter1.SeekToLast()
		iter2.SeekToLast()

		for iter1.Valid() {
			if tree.compare(iter1.Key(), iter2.Key()) != 0 {
				panic("")
			}
			iter1.Prev()
			iter2.Prev()
		}

		// log.Println()
	}
}

func TestTrimIndexForce(t *testing.T) {

	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		tree := New()
		tree.compare = compare.BytesLen
		tree2 := New()
		tree2.compare = compare.BytesLen

		for i := 0; i < 200; i += rand.Intn(4) + 1 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			tree2.Put(v, v)
		}

		s := rand.Int63n(tree.Size())
		e := rand.Int63n(tree.Size())
		if s > e {
			s, e = e, s
		}

		size := tree.Size()
		// log.Println(tree.debugString(true))

		tree.TrimByIndex(s, e)
		skey := tree2.index(s).Key
		ekey := tree2.Index(e).Key
		tree2.Trim(skey, ekey)
		// log.Println(tree.debugString(true), s, e)
		// log.Println(tree2.debugString(true), s, e)
		if e-s+1 != tree.Size() && tree.Size() != tree2.Size() {
			log.Panic(e, s, tree.Size(), size)
		}

		iter1 := tree.Iterator()
		iter2 := tree2.Iterator()

		iter1.SeekToFirst()
		iter2.SeekToFirst()

		for iter1.Valid() {
			if tree.compare(iter1.Key(), iter2.Key()) != 0 {
				panic("")
			}
			iter1.Next()
			iter2.Next()
		}

		iter1.SeekToLast()
		iter2.SeekToLast()

		for iter1.Valid() {
			if tree.compare(iter1.Key(), iter2.Key()) != 0 {
				panic("")
			}
			iter1.Prev()
			iter2.Prev()
		}

		// log.Println()
	}
}

func TestTrimIndex(t *testing.T) {

	tree := New()
	tree.compare = compare.BytesLen

	v := []byte(strconv.Itoa(0))
	tree.Put(v, v)
	tree.TrimByIndex(0, 0)
	if tree.Size() != 1 {
		t.Error()
	}

	for i := 0; i < 10; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

	if tree.Size() != 10 {
		t.Error()
	}
	tree.TrimByIndex(8, 9)
	if tree.Size() != 2 {
		t.Error()
	}
	if tree.IndexOf([]byte(strconv.Itoa(8))) != 0 {
		t.Error()
	}
	if tree.IndexOf([]byte(strconv.Itoa(9))) != 1 {
		t.Error()
	}

	var result []string
	tree.Traverse(func(s *Slice) bool {
		result = append(result, string(s.Key))
		return true
	})

	s := fmt.Sprintf("%v", result)
	if s != "[8 9]" {
		t.Error()
	}
}

func TestIntersectionSlice(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree1 := New()
		tree1.compare = compare.BytesLen
		tree2 := New()
		tree2.compare = compare.BytesLen

		var table1 map[string]bool = make(map[string]bool)
		var table2 map[string]bool = make(map[string]bool)

		for i := 0; i < 1000; i += rand.Intn(100) + 10 {
			v := []byte(strconv.Itoa(i))
			table1[string(v)] = true
			tree1.Put(v, v)

		}

		for i := 0; i < 1000; i += rand.Intn(100) + 10 {
			v := []byte(strconv.Itoa(i))
			table2[string(v)] = true
			tree2.Put(v, v)
		}

		var result1 []string
		for _, s := range tree1.intersectionSlice(tree2) {
			result1 = append(result1, string(s.Key))
		}

		var result2 []string
		for k := range table2 {
			if _, ok := table1[k]; ok {
				result2 = append(result2, k)
			}
		}
		sort.Slice(result2, func(i, j int) bool {
			return tree1.compare([]byte(result2[i]), []byte(result2[j])) < 0
		})

		if fmt.Sprintf("%v", result1) != fmt.Sprintf("%v", result2) {
			log.Panic(result1, result2)
		}
	}
}

func TestIntersection(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree1 := New()
		tree1.compare = compare.BytesLen
		tree2 := New()
		tree2.compare = compare.BytesLen

		var table1 map[string]bool = make(map[string]bool)
		var table2 map[string]bool = make(map[string]bool)

		for i := 0; i < 1000; i += rand.Intn(100) + 1 {
			v := []byte(strconv.Itoa(i))
			table1[string(v)] = true
			tree1.Put(v, v)

		}

		for i := 0; i < 1000; i += rand.Intn(100) + 1 {
			v := []byte(strconv.Itoa(i))
			table2[string(v)] = true
			tree2.Put(v, v)
		}

		var result1 []string
		tree1.Intersection(tree2).Traverse(func(s *Slice) bool {
			result1 = append(result1, string(s.Key))
			return true
		})

		var result2 []string
		for k := range table2 {
			if _, ok := table1[k]; ok {
				result2 = append(result2, k)
			}
		}
		sort.Slice(result2, func(i, j int) bool {
			return tree1.compare([]byte(result2[i]), []byte(result2[j])) < 0
		})

		if fmt.Sprintf("%v", result1) != fmt.Sprintf("%v", result2) {
			log.Panic(result1, result2)
		}
	}
}

func TestUnionSetSlice(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree1 := New()
		tree1.compare = compare.BytesLen
		tree2 := New()
		tree2.compare = compare.BytesLen

		var table1 map[string]bool = make(map[string]bool)
		// var table2 map[string]bool = make(map[string]bool)

		for i := 0; i < 1000; i += rand.Intn(100) + 10 {
			v := []byte(strconv.Itoa(i))
			table1[string(v)] = true
			tree1.Put(v, v)

		}

		for i := 0; i < 1000; i += rand.Intn(100) + 10 {
			v := []byte(strconv.Itoa(i))
			table1[string(v)] = true
			tree2.Put(v, v)
		}

		var result1 []string
		for _, s := range tree1.unionSetSlice(tree2) {
			result1 = append(result1, string(s.Key))
		}

		var result2 []string
		for k := range table1 {
			result2 = append(result2, k)
		}
		sort.Slice(result2, func(i, j int) bool {
			return tree1.compare([]byte(result2[i]), []byte(result2[j])) < 0
		})

		if fmt.Sprintf("%v", result1) != fmt.Sprintf("%v", result2) {
			log.Panic(result1, result2)
		}
	}
}

func TestUnionSet(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree1 := New()
		tree1.compare = compare.BytesLen
		tree2 := New()
		tree2.compare = compare.BytesLen

		var table1 map[string]bool = make(map[string]bool)
		// var table2 map[string]bool = make(map[string]bool)

		for i := 0; i < 1000; i += rand.Intn(100) + 1 {
			v := []byte(strconv.Itoa(i))
			table1[string(v)] = true
			tree1.Put(v, v)

		}

		for i := 0; i < 1000; i += rand.Intn(100) + 1 {
			v := []byte(strconv.Itoa(i))
			table1[string(v)] = true
			tree2.Put(v, v)
		}

		var result1 []string
		tree1.UnionSets(tree2).Traverse(func(s *Slice) bool {
			result1 = append(result1, string(s.Key))
			return true
		})

		var result2 []string
		for k := range table1 {
			result2 = append(result2, k)
		}
		sort.Slice(result2, func(i, j int) bool {
			return tree1.compare([]byte(result2[i]), []byte(result2[j])) < 0
		})

		if fmt.Sprintf("%v", result1) != fmt.Sprintf("%v", result2) {
			log.Panic(result1, result2)
		}
	}
}

func TestDifferenceSets(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree1 := New()
		tree1.compare = compare.BytesLen
		tree2 := New()
		tree2.compare = compare.BytesLen

		var table1 map[string]bool = make(map[string]bool)
		var table2 map[string]bool = make(map[string]bool)

		for i := 0; i < 1000; i += rand.Intn(100) + 1 {
			v := []byte(strconv.Itoa(i))
			table1[string(v)] = true
			tree1.Put(v, v)

		}

		for i := 0; i < 1000; i += rand.Intn(100) + 1 {
			v := []byte(strconv.Itoa(i))
			table2[string(v)] = true
			tree2.Put(v, v)
		}

		var result1 []string
		tree1.DifferenceSets(tree2).Traverse(func(s *Slice) bool {
			result1 = append(result1, string(s.Key))
			return true
		})

		var result2 []string
		for k := range table1 {
			if _, ok := table2[k]; !ok {
				result2 = append(result2, k)
			}
		}

		sort.Slice(result2, func(i, j int) bool {
			return tree1.compare([]byte(result2[i]), []byte(result2[j])) < 0
		})

		if fmt.Sprintf("%v", result1) != fmt.Sprintf("%v", result2) {
			log.Panic(result1, result2)
		}
	}
}

func TestCover(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		tree1 := New()
		tree1.compare = compare.BytesLen
		var dict map[int]int = make(map[int]int)
		for i := 0; i < 20; i++ {

			dict[i] = i
			tree1.Put([]byte(strconv.Itoa(i)), i)
		}

		for i := 0; i < 40; i++ {
			tree1.Put([]byte(strconv.Itoa(i)), rand.Intn(100)+20)
		}

		for k, i1 := range dict {
			if i2, ok := tree1.Get([]byte(strconv.Itoa(k))); !ok {
				panic("")
			} else if i2 != i1 {
				panic("")
			}
		}

		dict = make(map[int]int)
		tree1.Clear()

		for i := 0; i < 200; i++ {
			v := rand.Intn(100)
			dict[v] = i
			tree1.Cover([]byte(strconv.Itoa(v)), i)
		}

		for k, i1 := range dict {
			if i2, ok := tree1.Get([]byte(strconv.Itoa(k))); !ok {
				panic("")
			} else if i2 != i1 {
				panic("")
			}
		}

		dict = make(map[int]int)
		tree1.Clear()

		for i := 0; i < 200; i++ {
			v := rand.Intn(100)
			dict[v] = i
			tree1.PutDuplicate([]byte(strconv.Itoa(v)), i, func(exists *Slice) {
				exists.Value = i
			})
		}

		for k, i1 := range dict {
			if i2, ok := tree1.Get([]byte(strconv.Itoa(k))); !ok {
				panic("")
			} else if i2 != i1 {
				panic("")
			}
		}

	}
}
