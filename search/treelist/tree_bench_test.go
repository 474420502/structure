package treelist

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

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
	tree := New()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}
}

func BenchmarkIndex(b *testing.B) {
	b.StopTimer()
	tree := New()

	for i := 0; i < Level3; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

	b.ResetTimer()
	b.StartTimer()

	b.N = Level3
	for i := 0; i < Level3; i++ {
		tree.Index(int64(i))
	}
}

func TestRemoveRange(t *testing.T) {
	// rand.Seed(time.Now().UnixNano())
	var TreeListCountTime time.Duration = 0
	level := Level0 / 100

	// t.StopTimer()
	for i := 0; i < level; i++ {

		tree := New()
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

// 		// tree := New()
// 		// tree.compare = compare.BytesLen
// 		treeEx := New()
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
	seed := time.Now().UnixNano()
	log.Println(seed)
	rand.Seed(seed)
	// rand.Seed(time.Now().UnixNano())
	var TreeListCountTime time.Duration = 0
	level := Level0 / 100

	// t.StopTimer()
	for i := 0; i < level; i++ {

		tree := New()
		tree.compare = compare.BytesLen
		treeEx := New()
		treeEx.compare = compare.BytesLen
		for i := 0; i < level; i += rand.Intn(10) + 1 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			treeEx.Put(v, v)
		}

		s := rand.Intn(level)
		e := rand.Intn(level)
		if s > e {
			temp := s
			s = e
			e = temp
		}
		// log.Println(i)
		// if i == 81 {
		// 	log.Println()
		// }
		now := time.Now()
		ss := []byte(strconv.Itoa(s))
		ee := []byte(strconv.Itoa(e))
		tree.trimBad(ss, ee)
		a := tree.hashString()
		treeEx.Trim(ss, ee)
		b := treeEx.hashString()
		if a != b {
			log.Println(tree.debugString(true))
			log.Println(treeEx.debugString(true))
			t.Error(string(ss), string(ee))
		}
		TreeListCountTime += time.Since(now)
		// log.Println(tree.debugString(true), s, e)
	}
	t.Log(TreeListCountTime.Nanoseconds()/int64(level), "ns/op")
}
