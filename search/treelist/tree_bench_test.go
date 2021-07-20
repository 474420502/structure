package treelist

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
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
	var TreeListCountTime time.Duration = 0

	for i := 0; i < Level0/10; i++ {

		tree := New()
		for i := 0; i < Level0; i += 10 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
		}

		s := rand.Intn(Level0)
		e := rand.Intn(Level0)
		if s > e {
			temp := s
			s = e
			e = temp
		}

		now := time.Now()
		tree.RemoveRange([]byte(strconv.Itoa(s)), []byte(strconv.Itoa(e)))
		TreeListCountTime += time.Since(now)
	}

	t.Error(TreeListCountTime.Nanoseconds()/Level0*10, "ns/op")
}
