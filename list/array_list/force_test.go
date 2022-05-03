package arraylist

import (
	"container/list"
	"fmt"
	"log"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

func TestCasePushRemove(t *testing.T) {
	rand := random.New(t.Name())

	for n := 0; n < 2000; n++ {
		l := New(compare.Any[int])
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
		for ; iter.Vaild(); iter.Next() {
			if iter.Value() != carray[idx] {
				log.Println(idx, carray, len(carray), l.String(), l.Size())
				log.Panic("存在错误", iter.Value(), carray[idx])
			}
			idx++
		}

		for i := 0; i < count; i++ {
			idx = rand.Intn(count - i)
			carray = append(carray[0:idx], carray[idx+1:count-i]...)
			l.Remove(uint(idx))

			r1 := fmt.Sprintf("%v", carray)
			r2 := l.String()
			if r1 != r2 {
				t.Error(r1, r2)
			}
		}
	}
}

func TestCasePushPop(t *testing.T) {
	rand := random.New(t.Name())

	for n := 0; n < 2000; n++ {
		l := New(compare.Any[int])
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
			v, _ := l.Back()
			if gl.Back().Value != v {
				t.Error("list back error")
			}
		}

		if gl.Len() != 0 {
			v, _ := l.Front()
			if gl.Front().Value != v {
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
	rand := random.New(t.Name())

	for n := 0; n < 2000; n++ {
		l := New(compare.Any[int])
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
			if l.Contains(v) == 0 {
				t.Error("?")
			}
		}
	}
}

func TestCaseCircularIterator(t *testing.T) {
	rand := random.New(t.Name())

	for n := 0; n < 2000; n++ {
		l := New(compare.Any[int])
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

func StringSrcList(l *list.List) string {
	if l.Len() == 0 {
		return "[]"
	}
	var content []byte
	content = append(content, '[')
	for e := l.Front(); e != nil; e = e.Next() {
		content = append(content, []byte(fmt.Sprintf("%v ", e.Value))...)
	}
	content[len(content)-1] = ']'
	return string(content)
}

type _Iterator[T any] interface {
	Next()
	Value() T
	Vaild() bool
}

func StringIerator[T comparable](iter _Iterator[T], stringSize int) string {

	if stringSize == 0 {
		return "[]"
	}

	var content []byte
	content = append(content, '[')
	for i := 0; i < stringSize; i++ {
		content = append(content, []byte(fmt.Sprintf("%v ", iter.Value()))...)
		iter.Next()
	}

	content[len(content)-1] = ']'
	return string(content)
}

func TestIteratorGCompare(t *testing.T) {
	r := random.New()

	for i := 0; i < 100; i++ {
		al := New(compare.Any[int])
		l := list.New()

		for i := 0; i < r.Intn(3)+2; i++ {
			v1 := r.Int()
			al.PushFront(v1)
			l.PushFront(v1)
			v2 := r.Int()
			al.PushBack(v2)
			l.PushBack(v2)

			if al.String() != StringSrcList(l) {
				t.Error(al.String(), StringSrcList(l))
			}
		}

		r.Execute(1, 4, func() {
			var e1 *Iterator[int]
			var e2 *list.Element

			size := al.Size()
			if r.Bool() {
				e1 = al.Iterator()
				e1.ToHead()
				e2 = l.Front()

				r.Execute(0, int(size)-1, func() {
					e1.Next()
					e2 = e2.Next()
				})

			} else {
				e1 = al.Iterator()
				e1.ToTail()
				e2 = l.Back()

				r.Execute(0, int(size)-1, func() {
					e1.Prev()
					e2 = e2.Prev()
				})
			}

			v := r.Int()
			e1.SetValue(v)
			e2.Value = v

			idx := r.Intn(int(size))
			e1.IndexTo(uint(idx))
			e2 = l.Front()
			r.Execute(idx, idx, func() {
				e2 = e2.Next()
			})

			if e1.Value() != e2.Value {
				t.Error("e1.Value() != e2.Value", e1.Value(), e2.Value)
			}

			if r.Bool() {
				e2swap := e2.Next()
				if e2swap != nil {
					l.MoveBefore(e2swap, e2)
					e1swap := al.Iterator()
					e1swap.IndexTo(e1.Index())
					e1swap.Next()
					e1.Swap(e1swap)
				}
			} else {
				e2swap := e2.Prev()
				if e2swap != nil {
					l.MoveBefore(e2, e2swap)
					e1swap := al.Iterator()
					e1swap.IndexTo(e1.Index())
					e1swap.Prev()
					e1.Swap(e1swap)
				}
			}

			idx = r.Intn(int(size))
			e1.IndexTo(uint(idx))
			e2 = l.Front()
			r.Execute(idx, idx, func() {
				e2 = e2.Next()
			})

			if r.Bool() {
				e1.RemoveToNext()
				e2next := e2.Next()
				l.Remove(e2)
				if e2next != nil {
					if e2next.Value != e1.Value() {
						t.Error()
					}
				}

			} else {
				e1.RemoveToPrev()
				e2prev := e2.Prev()
				l.Remove(e2)
				if e2prev != nil {
					if e2prev.Value != e1.Value() {
						t.Error()
					}
				}
			}

		})

		if al.String() != StringSrcList(l) {
			t.Error(al.String(), StringSrcList(l))
		}

		iter := al.Iterator()
		for i := uint(0); i < al.Size(); i++ {
			iter.Next()
		}

		for i := uint(0); i < al.Size(); i++ {
			iter.Prev()
		}

		if s := StringIerator[int](iter, int(al.Size())); s != StringSrcList(l) {
			t.Error(s, StringSrcList(l))
		}

	}
}

func TestCIteratorGCompare(t *testing.T) {
	r := random.New()

	for i := 0; i < 100; i++ {
		al := New(compare.Any[int])
		l := list.New()

		for i := 0; i < r.Intn(3)+2; i++ {
			v1 := r.Int()
			al.PushFront(v1)
			l.PushFront(v1)
			v2 := r.Int()
			al.PushBack(v2)
			l.PushBack(v2)

			if al.String() != StringSrcList(l) {
				t.Error(al.String(), StringSrcList(l))
			}
		}

		r.Execute(1, 4, func() {
			var e1 *CircularIterator[int]
			var e2 *list.Element

			size := al.Size()
			if r.Bool() {
				e1 = al.CircularIterator()
				e1.ToHead()
				e2 = l.Front()

				r.Execute(0, int(size)-1, func() {
					e1.Next()
					e2 = e2.Next()
				})

			} else {
				e1 = al.CircularIterator()
				e1.ToTail()
				e2 = l.Back()

				r.Execute(0, int(size)-1, func() {
					e1.Prev()
					e2 = e2.Prev()
				})
			}

			v := r.Int()
			e1.SetValue(v)
			e2.Value = v

			idx := r.Intn(int(size))
			e1.IndexTo(uint(idx))
			e2 = l.Front()
			r.Execute(idx, idx, func() {
				e2 = e2.Next()
			})

			if e1.Value() != e2.Value {
				t.Error("e1.Value() != e2.Value", e1.Value(), e2.Value)
			}

			if r.Bool() {
				e2swap := e2.Next()
				if e2swap != nil {
					l.MoveBefore(e2swap, e2)
					e1swap := al.CircularIterator()
					e1swap.IndexTo(e1.Index())
					e1swap.Next()
					e1.Swap(e1swap)
				}
			} else {
				e2swap := e2.Prev()
				if e2swap != nil {
					l.MoveBefore(e2, e2swap)
					e1swap := al.CircularIterator()
					e1swap.IndexTo(e1.Index())
					e1swap.Prev()
					e1.Swap(e1swap)
				}
			}

			idx = r.Intn(int(size))
			e1.IndexTo(uint(idx))
			e2 = l.Front()
			r.Execute(idx, idx, func() {
				e2 = e2.Next()
			})

			if r.Bool() {
				e1.RemoveToNext()
				e2next := e2.Next()
				l.Remove(e2)
				if e2next != nil {
					if e2next.Value != e1.Value() {
						t.Error()
					}
				}

			} else {
				e1.RemoveToPrev()
				e2prev := e2.Prev()
				l.Remove(e2)
				if e2prev != nil {
					if e2prev.Value != e1.Value() {
						t.Error()
					}
				}
			}

		})

		if al.String() != StringSrcList(l) {
			t.Error(al.String(), StringSrcList(l))
		}

		iter := al.CircularIterator()
		for i := uint(0); i < al.Size(); i++ {
			iter.Next()
		}

		for i := uint(0); i < al.Size(); i++ {
			iter.Prev()
		}

		if s := StringIerator[int](iter, int(al.Size())); s != StringSrcList(l) {
			t.Error(s, StringSrcList(l))
		}

	}
}
