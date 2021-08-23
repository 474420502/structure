//go:build !test

package liststack

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkPush(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		v := rand.Int63()
		s.Push(v)
	}
}

func BenchmarkPushPop(b *testing.B) {

	s := New()

	b.StopTimer()

	for i := 0; i < 1000000; i++ {
		v := rand.Int63()
		s.Push(v)
	}

	b.StartTimer()

	count := 0
	op := 0
	for i := 0; i < b.N; i++ {

		if count == 0 {
			if rand.Intn(2) == 0 {
				op = 0
				count = rand.Intn(1000) + 100
			} else {
				op = 1
				count = rand.Intn(2000) + 200
			}
		} else {
			if op == 0 {
				v := rand.Int63()
				s.Push(v)
				count--
			} else {
				s.Pop()
				count--
			}
		}

	}
}

func TestBenchmarkPop(b *testing.T) {
	s := New()

	N := int64(1000000)

	for i := int64(0); i < N; i++ {
		v := rand.Int63()
		s.Push(v)
	}

	now := time.Now()

	for i := int64(0); i < N; i++ {
		_, ok := s.Pop()
		if ok == false {
			panic("")
		}
	}

	b.Log(time.Since(now).Nanoseconds(), float64(time.Since(now).Nanoseconds())/float64(N))
	// b.Fatal()
}
