package arraylist

import (
	"container/list"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestCasePushRemove(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		l := New()
		count := rand.Intn(50)
		var carray []int
		for i := 0; i < count; i++ {
			v := rand.Intn(1000)
			carray = append(carray, v)
			l.Push(v)

			if len(carray) != int(l.Size()) {
				log.Println(carray, l.String())
			}
		}

		iter := l.Iterator()
		var idx = 0
		for iter.Next() {
			if iter.Value() != carray[idx] {
				log.Println(idx, carray, len(carray), l.String(), l.Size())
				log.Panic("存在错误", iter.Value(), carray[idx])
			}
			idx++
		}

		for i := 0; i < count; i++ {
			idx = rand.Intn(count - i)
			carray = append(carray[0:idx], carray[idx+1:count-i]...)
			l.Remove(idx)

			r1 := fmt.Sprintf("%v", carray)
			r2 := l.String()
			if r1 != r2 {
				t.Error(r1, r2)
			}
		}
	}
}

func TestCasePushPop(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		l := New()
		count := rand.Intn(50)
		var gl = list.New()
		for i := 0; i < count; i++ {
			v := rand.Intn(1000)
			if i%2 == 0 {
				l.PushFront(v)
				gl.PushFront(v)
			} else {
				l.PushBack(v)
				gl.PushBack(v)
			}
		}

		// log.Println(gl.Back().Value, l.Back())
		if gl.Len() != 0 {
			if gl.Back().Value != l.Back() {
				t.Error("list back error")
			}
		}

		if gl.Len() != 0 {
			if gl.Front().Value != l.Front() {
				t.Error("list front error")
			}
		}

		glen := gl.Len()
		for i := 0; i < glen; i++ {

			if i%3 == 0 {
				v1 := gl.Remove(gl.Back())
				if v2, ok := l.PopBack(); !ok {
					t.Error("list pop back error")
				} else {
					if v2 != v1 {
						log.Panicln("popback error", i, glen)
					}
				}
			} else {
				v1 := gl.Remove(gl.Front())
				if v2, ok := l.PopFront(); !ok {
					t.Error("list Front back error")
				} else {
					if v2 != v1 {
						t.Error("popFront error")
					}
				}
			}
		}
	}
}

func TestCaseContains(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		l := New()
		count := rand.Intn(50)
		var temp map[int]bool = make(map[int]bool)
		for i := 0; i < count; i++ {
			v := rand.Intn(1000)
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
			temp[v] = true
		}
		for v := range temp {
			if !l.Contains(v) {
				t.Error("?")
			}
		}
	}
}

func TestCaseCircularIterator(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(seed)
	rand.Seed(seed)

	for n := 0; n < 2000; n++ {
		l := New()
		gl := list.New()
		count := rand.Intn(50)

		for i := 0; i < count; i++ {
			v := rand.Intn(1000)
			if i%2 == 0 {
				l.PushFront(v)
				gl.PushFront(v)
			} else {
				l.PushBack(v)
				gl.PushBack(v)
			}
		}

		gcur := gl.Front()
		cur := l.CircularIterator()
		cur.Next()
		for gcur != nil {
			if gcur.Value != cur.Value() {
				t.Error("?")
			}
			gcur = gcur.Next()
			cur.Next()
		}

		gcur = gl.Front()
		for gcur != nil {
			if gcur.Value != cur.Value() {
				t.Error("?")
			}
			gcur = gcur.Next()
			cur.Next()
		}

		gcur = gl.Front()
		for gcur != nil {
			if gcur.Value != cur.Value() {
				t.Error("?")
			}
			gcur = gcur.Next()
			cur.Next()
		}

	}
}
