package heap

import (
	"bytes"
	"encoding/binary"
	"sort"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

func TestHeapGrowSlimming(t *testing.T) {

	rand := random.New(t.Name())

	for ii := 0; ii < 2000; ii++ {

		h := New(compare.AnyDesc[int])
		var results []int
		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			results = append(results, v)
			h.Put(v)
		}
		sort.Slice(results, func(i, j int) bool {
			return results[i] > results[j]
		})

		if h.Size() != 100 || h.Empty() {
			t.Error("size != 100")
		}

		for i := 0; !h.Empty(); i++ {
			v, _ := h.Pop()
			if results[i] != v {
				t.Error("heap is error")
			}
		}

		if h.Size() != 0 {
			t.Error("size != 0")
		}

		h.Put(1)
		h.Put(5)
		h.Put(2)

		if r, _ := h.Top(); r != 5 {
			t.Error("top is not equal to 5")
		}

		h.Clear()
		h.Reset()

		if !h.Empty() {
			t.Error("clear reborn is error")
		}
	}

}

func TestHeapPushTopPop(t *testing.T) {
	h := New(compare.AnyDesc[int])
	l := []int{9, 5, 15, 2, 3}
	ol := []int{15, 9, 5, 3, 2}
	for _, v := range l {
		h.Put(v)
	}

	for _, tv := range ol {
		if v, isfound := h.Top(); isfound {
			if !(isfound && v == tv) {
				t.Error(v)
			}
		}

		if v, isfound := h.Pop(); isfound {
			if !(isfound && v == tv) {
				t.Error(v)
			}
		}
	}

	if h.Size() != 0 {
		t.Error("heap size is not equals to zero")
	}

	h.Clear()

	l = []int{3, 5, 2, 7, 1}

	for _, v := range l {
		h.Put(v)
	}

	sort.Slice(l, func(i, j int) bool {
		return l[i] > l[j]
	})

	for i := 0; !h.Empty(); i++ {
		v, _ := h.Pop()
		if l[i] != v {
			t.Error("heap is error")
		}
	}
}

// 做新研究 没实际意义
func TestCase(t *testing.T) {

	rand := random.New(t.Name())

	var buf = bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.BigEndian, []byte("12313"))
	if err != nil {
		panic(err)
	}

	var source []int
	for i := 0; i < 1000; i++ {
		source = append(source, i)
	}

	for n := 0; n < 10; n++ {
		min := New(compare.AnyDesc[int])
		max := New(compare.Any[int])

		rand.Shuffle(len(source), func(i, j int) {
			source[i], source[j] = source[j], source[i]
		})

		for i := 0; i < rand.Intn(90)+100; i++ {
			v := source[i]
			min.Put(v)
			max.Put(v)
		}

		minlist := min.elements[0:min.size]
		maxlist := max.elements[0:max.size]

		// log.Println(min.debugString())
		// log.Println(minlist)
		// log.Println(max.debugString())
		// log.Println(maxlist)

		var count = 0
		for i := range minlist {
			if minlist[i] == maxlist[i] {
				// log.Println(i, v)
				count++
			}

		}
	}
}

// func BenchmarkPush(b *testing.B) {
// 	h := New(compare.CompareAny[int])
// 	b.N = 40000000
// 	var results []int
// 	for i := 0; i < b.N; i++ {
// 		results = append(results, rand.Int())
// 	}

// 	b.ResetTimer()
// 	for _, v := range results {
// 		h.Put(v)
// 	}
// 	if h.Size() != b.N {
// 		b.Error(h.Size())
// 	}
// }

// func Int(k1, k2 interface{}) int {
// 	c1 := k1.(int)
// 	c2 := k2.(int)
// 	switch {
// 	case c1 > c2:
// 		return -1
// 	case c1 < c2:
// 		return 1
// 	default:
// 		return 0
// 	}
// }

// func TestPush(t *testing.T) {

// 	for i := 0; i < 1000000; i++ {
// 		h := New(Int)

// 		gods := binaryheap.NewWithIntComparator()
// 		for c := 0; c < 20; c++ {
// 			v := randomdata.Number(0, 100)
// 			h.Push(v)
// 			gods.Push(v)
// 		}

// 		r1 := fmt.Sprintf("%v",h.Values())
// 		r2 := fmt.Sprintf("%v",gods.Values())
// 		if r1 != r2 {
// 			t.Error(r1)
// 			t.Error(r2)
// 			break
// 		}
// 	}

// }

// func TestPop(t *testing.T) {

// 	for i := 0; i < 200000; i++ {
// 		h := New(compare.CompareAnyDesc[int])

// 		// m := make(map[int]int)
// 		gods := binaryheap.NewWithIntComparator()
// 		for c := 0; c < 40; c++ {
// 			v := randomdata.Number(0, 100)
// 			// if _, ok := m[v]; !ok {
// 			h.Put(v)
// 			gods.Push(v)
// 			// 	m[v] = v
// 			// }

// 		}

// 		// t.Error(h.Values())
// 		// t.Error(gods.Values())
// 		for c := 0; c < randomdata.Number(5, 10); c++ {
// 			v1, _ := h.Pop()
// 			v2, _ := gods.Pop()

// 			if v1 != v2 {
// 				t.Error(h.Values(), v1)
// 				t.Error(gods.Values(), v2)
// 				return
// 			}
// 		}

// 		r1 := fmt.Sprintf("%v",h.Values())
// 		r2 := fmt.Sprintf("%v",gods.Values())
// 		if r1 != r2 {
// 			t.Error(r1)
// 			t.Error(r2)
// 			break
// 		}
// 	}
// }

// func BenchmarkPush(b *testing.B) {

// 	l := loadTestData()

// 	b.ResetTimer()
// 	execCount := 50
// 	b.N = len(l) * execCount

// 	for c := 0; c < execCount; c++ {
// 		b.StopTimer()
// 		h := New(compare.CompareAny[int])
// 		b.StartTimer()
// 		for _, v := range l {
// 			h.Put(v)
// 		}
// 	}
// }

// func BenchmarkPop(b *testing.B) {

// 	h := New(compare.CompareAnyDesc[int])

// 	l := loadTestData()

// 	b.ResetTimer()
// 	execCount := 50
// 	b.N = len(l) * execCount

// 	for c := 0; c < execCount; c++ {
// 		b.StopTimer()
// 		for _, v := range l {
// 			h.Put(v)
// 		}
// 		b.StartTimer()
// 		for h.size != 0 {
// 			h.Pop()
// 		}
// 	}
// }

// func BenchmarkGodsPop(b *testing.B) {

// 	h := binaryheap.NewWithIntComparator()

// 	l := loadTestData()

// 	b.ResetTimer()
// 	execCount := 10
// 	b.N = len(l) * execCount

// 	for c := 0; c < execCount; c++ {
// 		b.StopTimer()
// 		for _, v := range l {
// 			h.Push(v)
// 		}
// 		b.StartTimer()
// 		for h.Size() != 0 {
// 			h.Pop()
// 		}
// 	}

// }

// func BenchmarkGodsPush(b *testing.B) {
// 	l := loadTestData()

// 	b.ResetTimer()
// 	execCount := 50
// 	b.N = len(l) * execCount

// 	for c := 0; c < execCount; c++ {
// 		b.StopTimer()
// 		h := binaryheap.NewWith(Int)
// 		b.StartTimer()
// 		for _, v := range l {
// 			h.Push(v)
// 		}
// 	}
// }
