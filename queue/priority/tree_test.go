package treequeue

import (
	"fmt"
	"log"
	"sort"
	"testing"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/random"
)

type dSlice struct {
	key   interface{}
	value interface{}
}

func (s *dSlice) String() string {
	return fmt.Sprintf("(%v,%v)", int(s.key.(float64)), s.value)
}

func (s *dSlice) Key() interface{} {
	return s.key
}

func (s *dSlice) Value() interface{} {
	return s.value
}

func (s *dSlice) SetValue(v interface{}) {
	s.value = v
}

func TestCase1(t *testing.T) {
	rand := random.New(t.Name())
	q := New(compare.Int)

	for i := 0; i < 20; i++ {
		v := rand.Intn(10)
		q.Put(v, i)
		// log.Println(q.debugStringWithValue())
	}

	q.Remove(2)
	q.check()

}

func TestExtractForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*dSlice
		var offset = 0.00001
		for i := 0; i < 200; i++ {
			v := rand.Intn(1000)
			queue.Put(v, i)
			priority = append(priority, &dSlice{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key().(float64) < priority[j].Key().(float64)
		})

		var start = rand.Intn(100)
		var end = rand.Intn(110)
		if start > end {
			start, end = end, start
		}

		var selects []int
		var s int = -1
		for i, v := range priority {
			if int(v.Key().(float64)) >= start {
				s = i
				break
			}
		}

		if s >= 0 {
			for _, v := range priority[s:] {
				if int(v.Key().(float64)) > end {
					break
				}
				selects = append(selects, v.Value().(int))
			}
		}

		queue.Extract(start, end)
		queue.check()

		r1 := fmt.Sprintf("%v", queue.Values())
		r2 := fmt.Sprintf("%v", selects)
		if r1 != r2 {
			log.Println(n, start, end)
			log.Println(priority)
			log.Panicln(r1, r2)
		}

		for i := 0; i < 10; i++ {
			v := rand.Intn(1000)
			queue.Put(v, i)
		}

		queue.check()
	}
}

func TestExtractIndexForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*dSlice
		var offset = 0.00001
		for i := 0; i < 200; i++ {
			v := rand.Intn(1000)
			queue.Put(v, i)
			priority = append(priority, &dSlice{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key().(float64) < priority[j].Key().(float64)
		})

		var start = rand.Intn(int(queue.Size()))
		var end = rand.Intn(int(queue.Size()))
		if start > end {
			start, end = end, start
		}

		var selects []int
		for _, v := range priority[start : end+1] {
			selects = append(selects, v.Value().(int))
		}

		queue.ExtractByIndex(int64(start), int64(end))
		queue.check()

		r1 := fmt.Sprintf("%v", queue.Values())
		r2 := fmt.Sprintf("%v", selects)
		if r1 != r2 {
			log.Println(n, start, end)
			log.Println(priority)
			log.Panicln(r1, r2)
		}

		for i := 0; i < 10; i++ {
			v := rand.Intn(1000)
			queue.Put(v, i)
		}

		queue.check()
	}
}

func TestRemoveForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*dSlice
		var offset = 0.00001
		for i := 0; i < 40; i++ {
			v := rand.Intn(100)
			queue.Put(v, i)
			priority = append(priority, &dSlice{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key().(float64) < priority[j].Key().(float64)
		})

		for queue.Size() > 15 {

			ridx := rand.Intn(len(priority))
			k := priority[ridx].Key().(float64) //TODO: 必须选择第一个.

			for ridx > 0 {
				ridx--
				if float64(int(k)) != float64(int(priority[ridx].Key().(float64))) {
					ridx++
					break
				}
			}

			priority = append(priority[0:ridx], priority[ridx+1:]...)

			queue.Remove(int(k))
			queue.check()

			var selectValues []int
			var selectKeys []float64
			for i, v := range priority {
				selectValues = append(selectValues, v.Value().(int))
				selectKeys = append(selectKeys, v.Key().(float64))
				s := queue.Index(int64(i))
				if s.Value() != v.Value() || s.Key() != int(v.Key().(float64)) {
					panic("")
				}
			}

			for i := 0; i < 5; i++ {
				idx := int64(rand.Intn(len(selectKeys)))
				k := selectKeys[idx]
				qidx := queue.IndexOf(int(k))

				if qidx != idx {
					if int(k) != int(selectKeys[idx-1]) {
						log.Println(queue.debugStringWithValue())
						log.Panicln(k, selectKeys[idx-1], idx, qidx, selectKeys)
					}
				}

				if int(selectKeys[qidx]) != int(k) {
					panic("")
				}

			}

			queue.check()

			r1 := fmt.Sprintf("%v", queue.Values())
			r2 := fmt.Sprintf("%v", selectValues)
			if r1 != r2 {
				log.Println(priority, ridx, int(k))
				log.Println(queue.debugStringWithValue())
				log.Println(r1)
				log.Println(r2)
				log.Panicln()
			}

		}

		for i := 0; i < 5; i++ {
			v := rand.Intn(1000)
			queue.Put(v, i)
		}

		queue.check()
	}
}

func TestRemoveRangeForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*dSlice
		var offset = 0.00001
		for i := 0; i < 20; i++ {
			v := rand.Intn(100)
			queue.Put(v, i)
			priority = append(priority, &dSlice{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key().(float64) < priority[j].Key().(float64)
		})

		var sidx = rand.Intn(int(queue.Size()))
		var eidx = rand.Intn(int(queue.Size()))
		if sidx > eidx {
			sidx, eidx = eidx, sidx
		}

		var start = priority[sidx]
		var end = priority[eidx]
		for sidx > 0 && int(priority[sidx-1].Key().(float64)) == int(start.Key().(float64)) {
			sidx--
		}
		for eidx != len(priority) && int(priority[eidx].Key().(float64)) == int(end.Key().(float64)) {
			eidx++
		}

		src := queue.debugStringWithValue()
		queue.RemoveRange(int(start.Key().(float64)), int(end.Key().(float64)))
		if eidx < len(priority) {
			priority = append(priority[0:sidx], priority[eidx:]...)
		} else {
			priority = priority[0:sidx]
		}

		var selectValues []int
		var selectKeys []float64
		for _, v := range priority {
			selectValues = append(selectValues, v.Value().(int))
			selectKeys = append(selectKeys, v.Key().(float64))
		}

		queue.check()

		r1 := fmt.Sprintf("%v", queue.Values())
		r2 := fmt.Sprintf("%v", selectValues)
		if r1 != r2 {
			log.Println(src)
			log.Println(queue.debugStringWithValue(), start.Key(), end.Key(), sidx, eidx)
			log.Println(r1)
			log.Println(r2)
			log.Panicln()
		}

	}
}

func TestRemoveRangeByIndexForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*dSlice
		var offset = 0.00001
		for i := 0; i < 20; i++ {
			v := rand.Intn(100)
			queue.Put(v, i)
			priority = append(priority, &dSlice{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key().(float64) < priority[j].Key().(float64)
		})

		var sidx = rand.Intn(int(queue.Size()))
		var eidx = rand.Intn(int(queue.Size()))
		if sidx > eidx {
			sidx, eidx = eidx, sidx
		}

		src := queue.debugStringWithValue()
		queue.RemoveRangeByIndex(int64(sidx), int64(eidx))
		eidx++
		if eidx < len(priority) {
			priority = append(priority[0:sidx], priority[eidx:]...)
		} else {
			priority = priority[0:sidx]
		}

		var selectValues []int
		var selectKeys []float64
		for _, v := range priority {
			selectValues = append(selectValues, v.Value().(int))
			selectKeys = append(selectKeys, v.Key().(float64))
		}

		queue.check()

		r1 := fmt.Sprintf("%v", queue.Values())
		r2 := fmt.Sprintf("%v", selectValues)
		if r1 != r2 {
			log.Println(src)
			log.Println(queue.debugStringWithValue(), sidx, eidx)
			log.Println(r1)
			log.Println(r2)
			log.Panicln()
		}

	}
}

func TestPutGetsRemoveIndexForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*dSlice
		var offset = 0.00001
		for i := 0; i < 40; i++ {
			v := rand.Intn(100)
			queue.Put(v, v)
			priority = append(priority, &dSlice{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key().(float64) < priority[j].Key().(float64)
		})

		for i := 0; i < 5; i++ {
			idx := rand.Intn(len(priority))
			v := queue.RemoveIndex(int64(idx))
			if int(priority[idx].Key().(float64)) != v.Key().(int) {
				log.Panicln(int(priority[idx].Key().(float64)), v.Key())
			}

			if idx == len(priority)-1 {
				priority = priority[0:idx]
			} else {
				priority = append(priority[0:idx], priority[idx+1:]...)
			}
		}

		var same map[int]int = make(map[int]int)
		for i := 0; i < len(priority)-1; i++ {
			v1 := priority[i]
			v2 := priority[i+1]

			key1 := int(v1.Key().(float64))
			key2 := int(v2.Key().(float64))

			if key1 == key2 {
				if _, ok := same[key1]; !ok {
					same[key1] = 1
				}
				same[key1]++
			}
		}

		for k, count := range same {
			r := queue.Gets(k)
			if len(r) != count {
				panic("")
			}

			for _, v := range r {
				if k != v.Key().(int) {
					panic("")
				}
			}
		}
	}

}

func TestHeadTailForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*dSlice
		var offset = 0.00001
		for i := 0; i < 40; i++ {
			v := rand.Intn(100)
			queue.Put(v, v)
			priority = append(priority, &dSlice{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key().(float64) < priority[j].Key().(float64)
		})

		for {

			if rand.Intn(2) == 0 {
				hslice := queue.Head()
				if hslice == nil {
					break
				}

				h1 := queue.Index(0)
				if h1.Key() != hslice.Key() || h1.Value() != hslice.Value() {
					panic("")
				}

				rslice := queue.RemoveHead()
				if rslice.Value() != hslice.Value() {
					panic("")
				}

			} else {
				tslice := queue.Tail()
				if tslice == nil {
					break
				}

				t1 := queue.Index(queue.Size() - 1)
				if t1.Key() != tslice.Key() || t1.Value() != tslice.Value() {
					panic("")
				}

				src := queue.Values()

				rslice := queue.RemoveTail()
				if rslice.Value() != tslice.Value() {
					log.Panicln(src, rslice.Value(), tslice.Value())
				}
			}
		}

	}

}

func TestForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*dSlice
		var offset = 0.00001

		for i := 0; i < 40; i++ {
			v := rand.Intn(100)
			queue.Put(v, v)
			priority = append(priority, &dSlice{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key().(float64) < priority[j].Key().(float64)
		})

		for i := 0; i < 5; i++ {
			idx := rand.Intn(len(priority))
			v := queue.RemoveIndex(int64(idx))
			if int(priority[idx].Key().(float64)) != v.Key().(int) {
				panic("")
			}

			if idx == len(priority)-1 {
				priority = priority[0:idx]
			} else {
				priority = append(priority[0:idx], priority[idx+1:]...)
			}
		}

		var same map[int]int = make(map[int]int)
		for i := 0; i < len(priority)-1; i++ {
			v1 := priority[i]
			v2 := priority[i+1]

			key1 := int(v1.Key().(float64))
			key2 := int(v2.Key().(float64))

			if key1 == key2 {
				if _, ok := same[key1]; !ok {
					same[key1] = 1
				}
				same[key1]++
			}
		}

		for k, count := range same {
			r := queue.Gets(k)
			if len(r) != count {
				panic("")
			}

			for _, v := range r {
				if k != v.Key().(int) {
					panic("")
				}
			}
		}

		var sidx = rand.Intn(int(queue.Size()))
		var eidx = rand.Intn(int(queue.Size()))
		if sidx > eidx {
			sidx, eidx = eidx, sidx
		}

		src := queue.debugStringWithValue()
		queue.RemoveRangeByIndex(int64(sidx), int64(eidx))
		eidx++
		if eidx < len(priority) {
			priority = append(priority[0:sidx], priority[eidx:]...)
		} else {
			priority = priority[0:sidx]
		}

		// var selectValues []int
		var selectKeys []int
		for _, v := range priority {
			// selectValues = append(selectValues, v.Value().(int))
			selectKeys = append(selectKeys, int(v.Key().(float64)))
		}

		r1 := fmt.Sprintf("%v", queue.Values())
		r2 := fmt.Sprintf("%v", selectKeys)
		if r1 != r2 {
			log.Println(src)
			log.Println(queue.debugStringWithValue(), sidx, eidx)
			log.Println(r1)
			log.Println(r2)
			log.Panicln()
		}

		for i := 0; i < 40; i++ {
			v := rand.Intn(200)
			queue.Put(v, v)
			priority = append(priority, &dSlice{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key().(float64) < priority[j].Key().(float64)
		})

		sidx = rand.Intn(int(queue.Size()))
		eidx = rand.Intn(int(queue.Size()))
		if sidx > eidx {
			sidx, eidx = eidx, sidx
		}

		var start = priority[sidx]
		var end = priority[eidx]
		for sidx > 0 && int(priority[sidx-1].Key().(float64)) == int(start.Key().(float64)) {
			sidx--
		}
		for eidx != len(priority) && int(priority[eidx].Key().(float64)) == int(end.Key().(float64)) {
			eidx++
		}

		src = queue.debugStringWithValue()
		queue.RemoveRange(int(start.Key().(float64)), int(end.Key().(float64)))
		if eidx < len(priority) {
			priority = append(priority[0:sidx], priority[eidx:]...)
		} else {
			priority = priority[0:sidx]
		}

		// var selectValues []int
		selectKeys = []int{}
		for _, v := range priority {
			// selectValues = append(selectValues, v.Value().(int))
			selectKeys = append(selectKeys, int(v.Key().(float64)))
		}

		r1 = fmt.Sprintf("%v", queue.Values())
		r2 = fmt.Sprintf("%v", selectKeys)
		if r1 != r2 {
			log.Println(src)
			log.Println(queue.debugStringWithValue(), start.Key(), end.Key(), sidx, eidx)
			log.Println(r1)
			log.Println(r2)
			log.Panicln()
		}

		queue.check()

		if queue.Size() != 0 && len(priority) != 0 {
			sidx = rand.Intn(int(queue.Size()))
			v2 := priority[sidx]
			v3 := queue.Index(int64(sidx))

			if int(v2.Key().(float64)) != v3.Key().(int) {
				log.Println(queue.Values())
				log.Println(priority)
				log.Panicln(v3.Value(), v2.Value())
			}
		}

		queue.Clear()
		queue.check()
	}
}
