package treelist

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/474420502/random"
	utils "github.com/474420502/structure"
	"github.com/474420502/structure/compare"
	avl "github.com/474420502/structure/tree/avl"
	indextree "github.com/474420502/structure/tree/itree"

	testutils "github.com/474420502/structure/tree/test_utils"
)

const Level0 = 100000
const Level1 = 1000000
const Level2 = 5000000
const Level3 = 10000000
const Level4 = 50000000
const Level5 = 100000000

func init() {
	// debug.SetGCPercent(800)
}

func BenchmarkPut(b *testing.B) {
	b.StopTimer()
	tree := New(compare.ArrayAny[[]byte])
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

}

func BenchmarkPut2(b *testing.B) {
	b.StopTimer()
	tree := New(compare.ArrayAny[[]byte])

	var data [][]byte
	for i := 0; i < Level1; i++ {
		data = append(data, utils.Rangdom(8, 32))
	}

	b.ResetTimer()
	b.StartTimer()

	var idx = 0
	for i := 0; i < b.N; i++ {
		if idx >= len(data) {
			idx = 0
		}
		tree.Put(data[idx], data[idx])
		idx++
	}

	b.Log(tree.Size())
}

func BenchmarkPut3(b *testing.B) {
	var data []int64

	if !testutils.LoadData("BenchmarkPut", data) {
		for i := 0; i < 4000000; i++ {
			v := rand.Int63()
			data = append(data, v)
		}
		testutils.SaveData("BenchmarkPut", data)
	}

	b.ResetTimer()

	rand.Seed(time.Now().Unix())
	start := int(rand.Int63n(500000))

	b.Run("pre", func(b *testing.B) {
		tree := avl.New(compare.Any[int64])

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v := data[i+start]
			tree.Put(v, v)
		}
		// b.Log(tree.Size())
	})

	b.Run("treelist", func(b *testing.B) {
		tree := New(compare.Any[int64])

		b.ResetTimer()

		// b.N = 100
		for i := 0; i < b.N; i++ {
			v := data[i+start]
			tree.Put(v, v)
		}

	})

	b.Run("indextree", func(b *testing.B) {
		tree := indextree.New(compare.Any[int64])

		b.ResetTimer()

		// b.N = 100
		for i := 0; i < b.N; i++ {
			v := data[i+start]
			tree.Put(v, v)
		}

	})

	b.Run("avl", func(b *testing.B) {
		tree := avl.New(compare.Any[int64])

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v := data[i+start]
			tree.Put(v, v)
		}

	})

}

func BenchmarkIndex(b *testing.B) {
	b.StopTimer()
	tree := New(compare.ArrayAny[[]byte])

	for i := 0; i < Level1; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

	b.ResetTimer()
	b.StartTimer()

	// b.N = Level3
	for i := 0; i < Level1; i++ {
		tree.Index(int64(i))
	}
}

func TestRemoveRange(t *testing.T) {
	// rand.Seed(time.Now().UnixNano())
	var TreeListCountTime time.Duration = 0
	level := Level0 / 100

	// t.StopTimer()
	for i := 0; i < level; i++ {

		tree := New(compare.ArrayAny[[]byte])
		tree.compare = compare.ArrayLenAny[[]byte]
		for i := 0; i < level; i += rand.Intn(10) + 10 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
		}

		s := rand.Intn(level)
		e := rand.Intn(level)
		if s > e {
			temp := s
			s = e
			e = temp
		}

		now := time.Now()
		tree.RemoveRange([]byte(strconv.Itoa(s)), []byte(strconv.Itoa(e)))
		TreeListCountTime += time.Since(now)
	}
	t.Log(TreeListCountTime.Nanoseconds()/int64(level), "ns/op")
}

// func TestTrimBadBench(t *testing.T) {
// 	seed := time.Now().UnixNano()
// 	log.Println(seed)
// 	rand.Seed(seed)
// 	// rand.Seed(time.Now().UnixNano())
// 	var TreeListCountTime time.Duration = 0
// 	level := Level0 / 100

// 	// t.StopTimer()
// 	for i := 0; i < level; i++ {

// 		// tree := New(compare.CompareBytes[[]byte])
// 		// tree.compare = compare.CompareBytesLen[[]byte]
// 		treeEx := New(compare.CompareBytes[[]byte])
// 		treeEx.compare = compare.CompareBytesLen[[]byte]
// 		for i := 0; i < level; i += rand.Intn(10) + 1 {
// 			v := []byte(strconv.Itoa(i))
// 			// tree.Put(v, v)
// 			treeEx.Put(v, v)
// 		}

// 		s := rand.Intn(level)
// 		e := rand.Intn(level)
// 		if s > e {
// 			temp := s
// 			s = e
// 			e = temp
// 		}
// 		// log.Println(i)
// 		// if i == 81 {
// 		// 	log.Println()
// 		// }
// 		ss := []byte(strconv.Itoa(s))
// 		ee := []byte(strconv.Itoa(e))
// 		now := time.Now()
// 		treeEx.trimBad(ss, ee)
// 		TreeListCountTime += time.Since(now)
// 		// log.Println(tree.debugString(true), s, e)
// 	}
// 	t.Log(TreeListCountTime.Nanoseconds()/int64(level), "ns/op")
// }

func TestTrimBench(t *testing.T) {
	// rand := random.New(t.Name())
	// // rand.Seed(time.Now().UnixNano())
	// var TreeListCountTime time.Duration = 0
	// level := Level0 / 100

	// // t.StopTimer()
	// for i := 0; i < level; i++ {

	// 	tree := New(compare.CompareBytes[[]byte])
	// 	tree.compare = compare.CompareBytesLen[[]byte]
	// 	treeEx := New(compare.CompareBytes[[]byte])
	// 	treeEx.compare = compare.CompareBytesLen[[]byte]
	// 	for i := 0; i < level; i += rand.Intn(10) + 1 {
	// 		v := []byte(strconv.Itoa(i))
	// 		tree.Put(v, v)
	// 		treeEx.Put(v, v)
	// 	}

	// 	s := rand.Intn(level)
	// 	e := rand.Intn(level)
	// 	if s > e {
	// 		temp := s
	// 		s = e
	// 		e = temp
	// 	}
	// 	// log.Println(i)
	// 	// if i == 81 {
	// 	// 	log.Println()
	// 	// }
	// 	now := time.Now()
	// 	ss := []byte(strconv.Itoa(s))
	// 	ee := []byte(strconv.Itoa(e))
	// 	tree.trimBad(ss, ee)
	// 	a := tree.hashString()
	// 	treeEx.Trim(ss, ee)
	// 	b := treeEx.hashString()
	// 	if a != b {
	// 		log.Println(tree.debugString(true))
	// 		log.Println(treeEx.debugString(true))
	// 		t.Error(string(ss), string(ee))
	// 	}
	// 	TreeListCountTime += time.Since(now)
	// 	// log.Println(tree.debugString(true), s, e)
	// }
	// t.Log(TreeListCountTime.Nanoseconds()/int64(level), "ns/op")
}

// 优化交集测试
func estIntersectionP(t *testing.T) {
	rand := random.New(t.Name())

	var cost1 time.Duration
	var cost2 time.Duration

	for n := 0; n < 2000; n++ {

		var table1 map[string]bool = make(map[string]bool)
		var table2 map[string]bool = make(map[string]bool)

		tree1 := New(compare.ArrayAny[[]byte])
		tree1.compare = compare.ArrayLenAny[[]byte]
		tree2 := New(compare.ArrayAny[[]byte])
		tree2.compare = compare.ArrayLenAny[[]byte]

		for i := 0; i < 100000; i += rand.Intn(1000) + 1 {
			v := []byte(strconv.Itoa(i))
			// now := time.Now()
			table1[string(v)] = true
			// cost2 += time.Since(now)

			// now = time.Now()
			tree1.Put(v, v)
			// cost1 += time.Since(now)
		}

		for i := 0; i < 100000; i += rand.Intn(1000) + 1 {
			v := []byte(strconv.Itoa(i))
			// now := time.Now()
			table2[string(v)] = true
			// cost2 += time.Since(now)

			// now = time.Now()
			tree2.Put(v, v)
			// cost1 += time.Since(now)
		}

		// var a1 []*Slice
		// var a2 []*Slice

		// tree1.Traverse(func(s *Slice) bool {
		// 	a2 = append(a2, s)
		// 	return true
		// })

		// tree2.Traverse(func(s *Slice) bool {
		// 	a1 = append(a1, s)
		// 	return true
		// })

		now := time.Now()
		tree1.intersectionSlice(tree2)
		cost1 += time.Since(now)

		now = time.Now()
		var result []string
		for k := range table2 {
			if _, ok := table1[k]; ok {
				result = append(result, k)
			}
		}
		cost2 += time.Since(now)
	}

	log.Println(cost1, cost2)
}
