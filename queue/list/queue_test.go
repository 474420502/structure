package listqueue

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/474420502/random"
)

func TestForcePush(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		queue1 := New()
		queue2 := list.New()

		for i := 0; i < 10; i += 1 {

			v := rand.Intn(100)
			if rand.Intn(2) == 0 {
				queue1.PushBack(v)
				queue2.PushBack(v)
			} else {
				queue1.PushFront(v)
				queue2.PushFront(v)
			}

			if queue1.Front().Value() != queue2.Front().Value {
				panic(fmt.Errorf("%d,%d", queue1.Front().Value(), queue2.Front().Value))
			}

			if queue1.Back().Value() != queue2.Back().Value {
				panic("")
			}

			if queue1.Size() != int64(queue2.Len()) {
				panic("")
			}

		}

		e1 := queue1.Front()
		e2 := queue2.Front()
		for e1 != nil {
			if e1.Value() != e2.Value {
				panic(fmt.Errorf("%d,%d", e1.Value(), e2.Value))
			}
			e1 = e1.Next()
			e2 = e2.Next()
		}

		if e1 != nil {
			panic("")
		}

		for n2 := 0; n2 < 50; n2++ {

			// if n == 3 && n2 == 1 {
			// 	log.Println("")
			// }
			// log.Println(n, n2)

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
					front2 := queue2.Front()
					front1 := queue1.PopFront()
					if front2.Value != front1 {
						panic(fmt.Errorf("%d,%d", front1, front2.Value))
					}
					queue2.Remove(front2)
				}
			} else {
				break
			}

		}

		for queue1.Size() != 0 {
			if queue1.Front().Value() != queue2.Front().Value {
				panic(fmt.Errorf("%d,%d", queue1.Front().Value(), queue2.Front().Value))
			}

			if queue1.Back().Value() != queue2.Back().Value {
				panic("")
			}

			if queue1.Size() != int64(queue2.Len()) {
				panic("")
			}

			if rand.Intn(2) == 0 {
				queue1.PopBack()
				queue2.Remove(queue2.Back())
			} else {
				queue1.PopFront()
				queue2.Remove(queue2.Front())
			}

			if rand.Intn(3) == 0 {
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

	}
}
