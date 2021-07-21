package treelist

import (
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
	level := Level0 / 10

	// t.StopTimer()
	for i := 0; i < level; i++ {

		tree := New()
		tree.compare = compare.BytesLen
		for i := 0; i < level; i += rand.Intn(10) + 5 {
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
	t.Error(TreeListCountTime.Nanoseconds()/int64(level), "ns/op")
}
