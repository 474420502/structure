package treequeue

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/474420502/structure/compare"
)

func TestCase1(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(1628580888818576051)

	q := New(compare.Int)

	for i := 0; i < 20; i++ {
		v := rand.Intn(10)
		q.Put(v, i)
		log.Println(q.debugStringWithValue())
	}

	log.Println(q.debugStringWithValue())
	log.Println(q.Gets(2))
	q.Remove(2)
	q.check()
	log.Println(q.debugStringWithValue())
	log.Println(q.Gets(2))
}

type tKey struct {
	Key   float64
	Value int
}

func (k *tKey) String() string {
	return fmt.Sprintf("(%v,%v)", int(k.Key), k.Value)
}

func TestExtractForce(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*tKey
		var offset = 0.00001
		for i := 0; i < 200; i++ {
			v := rand.Intn(1000)
			queue.Put(v, i)
			priority = append(priority, &tKey{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key < priority[j].Key
		})

		var start = rand.Intn(100)
		var end = rand.Intn(110)
		if start > end {
			start, end = end, start
		}

		var selects []int

		var s int = -1
		for i, v := range priority {
			if int(v.Key) >= start {
				s = i
				break
			}
		}

		if s >= 0 {
			for _, v := range priority[s:] {
				if int(v.Key) > end {
					break
				}
				selects = append(selects, v.Value)
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
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*tKey
		var offset = 0.00001
		for i := 0; i < 200; i++ {
			v := rand.Intn(1000)
			queue.Put(v, i)
			priority = append(priority, &tKey{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key < priority[j].Key
		})

		var start = rand.Intn(int(queue.Size()))
		var end = rand.Intn(int(queue.Size()))
		if start > end {
			start, end = end, start
		}

		var selects []int
		for _, v := range priority[start : end+1] {
			selects = append(selects, v.Value)
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
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		queue := New(compare.Int)
		var priority []*tKey
		var offset = 0.00001
		for i := 0; i < 20; i++ {
			v := rand.Intn(100)
			queue.Put(v, i)
			priority = append(priority, &tKey{float64(v) + offset, i})
			offset += 0.00001
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i].Key < priority[j].Key
		})

		for queue.Size() > 15 {

			ridx := rand.Intn(len(priority))
			k := priority[ridx].Key //TODO: 必须选择第一个.
			// log.Println(priority)
			for ridx > 0 {
				ridx--
				if float64(int(k)) != float64(int(priority[ridx].Key)) {
					ridx++
					break
				}
			}

			priority = append(priority[0:ridx], priority[ridx+1:]...)

			queue.Remove(int(k))
			queue.check()
			// var start = rand.Intn(int(queue.Size()))
			// var end = rand.Intn(int(queue.Size()))
			// if start > end {
			// 	start, end = end, start
			// }

			var selectValues []int
			var selectKeys []float64
			for i, v := range priority {
				selectValues = append(selectValues, v.Value)
				selectKeys = append(selectKeys, v.Key)
				qk, qv := queue.Index(int64(i))
				if qv != v.Value || qk != int(v.Key) {
					panic("")
				}
			}

			for i := 0; i < 5; i++ {
				idx := int64(rand.Intn(len(selectKeys)))
				k := selectKeys[idx]
				// if k == 39.00003 {
				// 	log.Println()
				// }
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

func TestIndexOfForce(t *testing.T) {

}
