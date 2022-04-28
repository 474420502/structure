package arrayqueue

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/474420502/random"
	testutils "github.com/474420502/structure"
)

func TestCasePut(t *testing.T) {
	q := New[int]()

	for i := 0; i < 16; i++ {
		q.PushBack(i)
	}

	if len(q.Values()) != int(q.size) {
		panic("Values error")
	}

	if q.cap != int64(cap(q.data)) {
		panic("cap error")
	}

	if q.Front() != 0 && q.Front() != q.Index(0) {
		panic("check error")
	}
	// log.Println(q.Values(), "cap:", q.cap, cap(q.data))
}

func TestCasePop(t *testing.T) {
	var result []int
	q := New[int]()

	for i := 0; i < 16; i++ {
		q.PushBack(i)
	}

	result = q.Values()
	q.Traverse(func(idx int64, value int) bool {
		if result[idx] != value {
			panic("check error")
		}
		return true
	})

	for i := 0; i < 8; i++ {
		q.PopFront()
	}

	result = q.Values()
	q.Traverse(func(idx int64, value int) bool {
		if result[idx] != value {
			panic("check error")
		}
		return true
	})

	if q.Front() == 0 {
		panic("check")
	}

	if len(q.Values()) != int(q.size) {
		panic(fmt.Errorf("Values error: %d %d", len(q.Values()), int(q.size)))
	}

	if q.cap != int64(cap(q.data)) {
		panic("cap error")
	}

	for i := 0; i < 4; i++ {
		q.PopFront()
	}

	if len(q.Values()) != int(q.Size()) {
		panic("Values error")
	}

	result = q.Values()
	q.Traverse(func(idx int64, value int) bool {
		if result[idx] != value {
			panic("check error")
		}
		return true
	})

	q.PushBack(100)
	if q.Index(q.Size()-1) != 100 {
		panic("check error")
	}

	err := testutils.TryPanic(func() {
		q.Index(q.Size())
	})

	if err == nil {
		panic("check")
	}

	for i := 0; i < 16; i++ {
		q.PushBack(i)
	}

	if q.Index(q.Size()-1) != 15 {
		panic("check error")
	}

	if q.Index(q.Size()-2) != 14 {
		panic("check error")
	}

	err = testutils.TryPanic(func() {
		q.Index(q.Size())
	})
	if err == nil {
		panic("check")
	}
}

func TestForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		queue1 := New[int]()
		queue2 := list.New()

		for i := 0; i < 10; i += 1 {

			v := rand.Intn(100)
			if i%2 == 0 {
				queue1.PushBack(v)
				queue2.PushBack(v)
			} else {
				queue1.PushFront(v)
				queue2.PushFront(v)
			}

			if queue1.Front() != queue2.Front().Value {
				panic(fmt.Errorf("%d,%d", queue1.Front(), queue2.Front().Value))
			}

			if queue1.Back() != queue2.Back().Value {
				panic("")
			}

			if queue1.Size() != int64(queue2.Len()) {
				panic("")
			}

		}

		var i int64 = 0
		for e := queue2.Front(); e != nil; e = e.Next() {
			if e.Value != queue1.Index(i) {
				panic("")
			}
			i++
		}

		e := queue2.Front()
		queue1.Traverse(func(idx int64, value int) bool {
			if e.Value != value {
				panic(fmt.Errorf("%d,%d,%d", idx, e.Value, value))
			}
			e = e.Next()
			return true
		})

		if e != nil {
			panic("")
		}

		for n := 0; n < 50; n++ {

			if queue1.Size() != int64(queue2.Len()) {
				panic("")
			}

			if queue1.Size() != 0 {
				if rand.Int()%2 == 0 {
					back := queue2.Back()
					if queue1.PopBack() != back.Value {
						panic("")
					}
					queue2.Remove(back)
				}
			} else {
				break
			}

			for x := 0; x < 10; x++ {
				if rand.Intn(2) == 0 {
					v := rand.Intn(100)
					if rand.Intn(2) == 0 {
						queue1.PushBack(v)
						queue2.PushBack(v)
					} else {
						queue1.PushFront(v)
						queue2.PushFront(v)
					}
				}
			}

			if queue1.Size() != int64(queue2.Len()) {
				panic("")
			}

			if queue1.Size() != 0 {
				if rand.Int()%2 == 0 {
					front := queue2.Front()
					if front.Value != queue1.PopFront() {
						panic("")
					}
					queue2.Remove(front)
				}
			} else {
				break
			}

		}

		e = queue2.Front()
		queue1.Traverse(func(idx int64, value int) bool {
			if e.Value != value {
				panic(fmt.Errorf("%d,%d,%d", idx, e.Value, value))
			}
			e = e.Next()
			return true
		})

		e = queue2.Front()
		for _, value := range queue1.Values() {
			if e.Value != value {
				panic(fmt.Errorf(" %d,%d", e.Value, value))
			}
			e = e.Next()
		}

		if queue1.Front() != queue2.Front().Value {
			panic(fmt.Errorf("%d,%d", queue1.Front(), queue2.Front().Value))
		}

		if queue1.Back() != queue2.Back().Value {
			panic("")
		}

		if queue1.Size() != int64(queue2.Len()) {
			panic("")
		}

	}
}
