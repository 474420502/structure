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
	tree := New(compare.Bytes)
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

}

func BenchmarkPut2(b *testing.B) {
	b.StopTimer()
	tree := New(compare.Bytes)

	var data [][]byte
	for i := 0; i < Level0; i++ {
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

	b.Log(len(data))
}

func BenchmarkIndex(b *testing.B) {
	b.StopTimer()
	tree := New(compare.Bytes)

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

		tree := New(compare.Bytes)
		tree.compare = compare.BytesLen
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

// 		// tree := New(compare.Bytes)
// 		// tree.compare = compare.BytesLen
// 		treeEx := New(compare.Bytes)
// 		treeEx.compare = compare.BytesLen
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

	// 	tree := New(compare.Bytes)
	// 	tree.compare = compare.BytesLen
	// 	treeEx := New(compare.Bytes)
	// 	treeEx.compare = compare.BytesLen
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

		tree1 := New(compare.Bytes)
		tree1.compare = compare.BytesLen
		tree2 := New(compare.Bytes)
		tree2.compare = compare.BytesLen

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
