package linkedlist

import (
	"fmt"
	"log"
	"testing"

	"github.com/474420502/random"
)

func TestPush(t *testing.T) {
	l := New[int]()
	for i := 0; i < 5; i++ {
		l.Push(i)
	}
	var result string
	result = fmt.Sprintf("%v", l.Values())
	if result != "[0 1 2 3 4]" {
		t.Error(result)
	}

	l.Push(0)
	result = fmt.Sprintf("%v", l.Values())
	if result != "[0 1 2 3 4 0]" {
		t.Error(result)
	}
}

func TestPushFront(t *testing.T) {
	l := New[int]()
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	var result string
	result = fmt.Sprintf("%v", l.Values())
	if result != "[4 3 2 1 0]" {
		t.Error(result)
	}

	l.PushFront(0)
	result = fmt.Sprintf("%v", l.Values())
	if result != "[0 4 3 2 1 0]" {
		t.Error(result)
	}
}

func TestPushBack(t *testing.T) {
	l := New[int]()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	var result string
	result = fmt.Sprintf("%v", l.Values())
	if result != "[0 1 2 3 4]" {
		t.Error(result)
	}

	l.PushBack(0)
	result = fmt.Sprintf("%v", l.Values())
	if result != "[0 1 2 3 4 0]" {
		t.Error(result)
	}
}

func TestPopFront(t *testing.T) {
	l := New[int]()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	// var result string

	for i := 4; i >= 0; i-- {
		if v, ok := l.PopFront(); ok {
			if v != i {
				t.Error("[4 3 2 1 0] PopFront value should be ", i, ", but is ", v)
			}
		} else {
			t.Error("PopFront is not ok")
		}

		if l.Size() != uint(i) {
			t.Error("l.Size() is error, is", l.Size())
		}
	}
}

func TestPopBack(t *testing.T) {
	l := New[int]()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	// var result string

	for i := 0; i < 5; i++ {
		if v, ok := l.PopBack(); ok {
			if v != i {
				t.Error("[4 3 2 1 0] PopFront value should be ", i, ", but is ", v)
			}
		} else {
			t.Error("PopFront is not ok")
		}

		if l.Size() != uint(5-i-1) {
			t.Error("l.Size() is error, is", l.Size())
		}
	}

}

func TestIndex(t *testing.T) {
	l := New[int]()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	if v, ok := l.Index(4); ok {
		if v != 0 {
			t.Error("[4 3 2 1 0] Index 4 value is 0, but v is ", v)
		}
	} else {
		t.Error("not ok is error")
	}

	if v, ok := l.Index(1); ok {
		if v != 3 {
			t.Error("[4 3 2 1 0] Index 1 value is 3, but v is ", v)
		}
	} else {
		t.Error("not ok is error")
	}

	if v, ok := l.Index(0); ok {
		if v != 4 {
			t.Error("[4 3 2 1 0] Index 1 value is 4, but v is ", v)
		}
	} else {
		t.Error("not ok is error")
	}

	if _, ok := l.Index(5); ok {
		t.Error("[4 3 2 1 0] Index 5, out of range,ok = true is error")
	}

}

func TestTraversal(t *testing.T) {
	l := New[uint]()
	for i := 0; i < 5; i++ {
		l.PushFront(uint(i))
	}

	var result []interface{}

	l.Traverse(func(v uint) bool {
		result = append(result, v)
		return true
	})

	if fmt.Sprintf("%v", result) != "[4 3 2 1 0]" {
		t.Error(result)
	}

	l.PushBack(7, 8)
	result = nil
	l.Traverse(func(v uint) bool {
		result = append(result, v)
		return true
	})

	if fmt.Sprintf("%v", result) != "[4 3 2 1 0 7 8]" {
		t.Error(result)
	}
}

func TestIterator(t *testing.T) {
	ll := New[int]()
	for i := 0; i < 10; i++ {
		ll.PushFront(i)
	}

	iter := ll.Iterator()

	for i := 0; iter.Vaild(); i++ {
		if iter.Value() != 9-i {
			t.Error("iter.Next() ", iter.Value(), "is not equal ", 9-i)
		}
		iter.Next()
	}

	if iter.cur != iter.ll.tail {
		t.Error("current point is not equal tail ", iter.ll.tail)
	}

	for i := 0; iter.Vaild(); i++ {
		if iter.Value() != i {
			t.Error("iter.Prev() ", iter.Value(), "is not equal ", i)
		}
		iter.Prev()
	}

	iter.ToTail()
	if !iter.Vaild() {
		t.Error(iter.Value())
	}

	iter.ToHead()
	if !iter.Vaild() {
		t.Error(iter.Value())
	}

}

func TestCircularIterator(t *testing.T) {
	ll := New[int]()
	for i := 0; i < 10; i++ {
		ll.PushFront(i)
	}

	iter := ll.CircularIterator()

	for i := 0; i != 10; i++ {
		if iter.Value() != 9-i {
			t.Error("iter.Next() ", iter.Value(), "is not equal ", 9-i)
		}
		iter.Next()
	}

	if !iter.Vaild() {
		t.Error("should be true", iter.Value(), iter.ll.String())
	}

	if iter.Next(); iter.Vaild() {
		if iter.Value() != 8 {
			t.Error("iter.Value() != ", iter.Value())
		}
	}

	iter.ToTail()
	for i := 0; i != 9; i++ {
		if iter.Value() != i {
			t.Error("iter.Prev() ", iter.Value(), "is not equal ", i)
		}
		iter.Prev()
	}

	if iter.cur != iter.ll.head.next {
		t.Error("current point is not equal tail ", iter.ll.tail.prev)
	}

	if iter.Prev(); iter.Vaild() {
		if iter.Value() != 0 {
			t.Error("iter.Value() != ", iter.Value())
		}
	}

	iter.ToTail()
	iter.Next()
	if iter.Value() != 9 {
		t.Error()
	}

	iter.ToHead()
	iter.Prev()
	if iter.Value() != 0 {
		t.Error()
	}
}

func TestContains(t *testing.T) {
	ll := New[int]()
	for i := 0; i < 10; i++ {
		ll.Push(i)
	}

	for i := 0; i < 10; i++ {
		if ll.Contains(i) == 0 {
			t.Error(i)
		}
	}

	for i := 10; i < 20; i++ {
		if ll.Contains(i) > 0 {
			t.Error(i)
		}
	}

	if v, _ := ll.Front(); v != 0 {
		t.Error(v)
	}

	if v, _ := ll.Back(); v != 9 {
		t.Error(v)
	}

	ll.Clear()
	if !ll.Empty() {
		t.Error("not Empty?")
	}

	if v, ok := ll.Front(); ok {
		t.Error(v)
	}
}

func TestForce(t *testing.T) {

	rand := random.New(t.Name())
	l := New[int]()
	// "[4 3 2 1 0]"
	for n := 0; n < 2000; n++ {

		for i := 0; i < 200; i++ {

			if rand.Bool() {
				l.PushFront(rand.Intn(1000))
			} else {
				l.PushBack(rand.Intn(1000))
			}

			if rand.Bool() {
				idx := uint(rand.Intn(int(l.Size())))
				v := rand.Intn(1000)

				iter := l.Iterator()
				for i := 0; i < int(idx); i++ {
					iter.Next()
				}
				iter.InsertFront(v)

				var i = uint(0)
				l.Traverse(func(value int) bool {
					if i == idx {
						if value != v {
							log.Panic(l.String())
						}
					}
					i++
					return true
				})
			}
		}

		idx := rand.Intn(int(l.Size()))
		v, ok := l.Index(idx)
		if !ok {
			panic("")
		}

		var i = 0
		l.Traverse(func(value int) bool {
			if i == idx {
				if value != v {
					panic(l.String())
				}
			}
			i++
			return true
		})

		var lsize = l.Size()
		var result1 = make([]int, lsize)
		var result2 = make([]int, lsize)

		for citer := l.CircularIterator(); citer.Vaild(); citer.Next() {
			if len(result1) != int(lsize) {
				result1 = append(result1, citer.Value().(int))
			} else if len(result2) != int(lsize) {
				result2 = append(result2, citer.Value().(int))
			} else {
				break
			}
		}

		if fmt.Sprintf("%v", result1) != fmt.Sprintf("%v", result2) {
			panic("")
		}

		l.Clear()
	}
}

func TestIteratorInsert(t *testing.T) {

	l := New[int]()
	l.Push(1)

	iter := l.Iterator()
	// iter.Next() // next to 1

	if !iter.Vaild() {
		panic("")
	}
	iter.InsertFront(2, 3, 5)

	// log.Println(l.String(), iter.Value()) // [2 3 5 1] 1
	if l.String() != "[2 3 5 1]" && iter.Value() != 1 {
		t.Error("InsertFront error")
	}

	iter.InsertBack(20, 30, 50)

	// log.Println(l.String(), iter.Value()) // [2 3 5 1 20 30 50] 1
	if l.String() != "[2 3 5 1 20 30 50]" && iter.Value() != 1 {
		t.Error("InsertBack error")
	}

	iter.Next()
	// log.Println(iter.Value())
	if iter.Value() != 20 {
		t.Error("")
	}

	iter.Prev()
	iter.Prev()
	if iter.Value() != 5 {
		t.Error("")
	}

}

func TestCircularIteratorInsert(t *testing.T) {

	l := New[int]()
	l.Push(1)

	iter := l.CircularIterator()
	iter.Next() // next to 1

	iter.InsertFront(2, 3, 5)

	// log.Println(l.String(), iter.Value()) // [2 3 5 1] 1
	if l.String() != "[2 3 5 1]" && iter.Value() != 1 {
		t.Error("InsertFront error")
	}

	iter.InsertBack(20, 30, 50)

	// log.Println(l.String(), iter.Value()) // [2 3 5 1 20 30 50] 1
	if l.String() != "[2 3 5 1 20 30 50]" && iter.Value() != 1 {
		t.Error("InsertBack error")
	}

	iter.Next()
	// log.Println(iter.Value())
	if iter.Value() != 20 {
		t.Error("")
	}

	iter.Prev()
	iter.Prev()
	if iter.Value() != 5 {
		t.Error("")
	}

}

func TestIteratorRemove(t *testing.T) {
	l := New[int]()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	iter := l.Iterator()
	iter.Move(2)
	iter.RemoveToNext()

	// log.Println(l.String(), iter.Value()) //   4 -> 2(Remove) and Cur to 1  "[4 3 1 0]" 1
	if l.String() != "[4 3 1 0]" || iter.Value() != 1 {
		t.Error(l.String(), iter.Value())
	}

	iter.Move(-1)       // 3
	iter.RemoveToPrev() // 3(remove) -> 4
	iter.Move(2)        // 0

	if l.String() != "[4 1 0]" || iter.Value() != 0 {
		t.Error(l.String(), iter.Value())
	}

	for iter.Vaild() {
		iter.RemoveToNext()
	}

	if l.String() != "[4 1]" || iter.Value() != 0 {
		t.Error(l.String(), iter.Value())
	}

	for iter.Vaild() { // can not remove tail or head
		iter.RemoveToPrev()
	}

	if l.String() != "[4 1]" || iter.Value() != 0 {
		t.Error(l.String(), iter.Value())
	}

	iter.Move(-1)      // 1
	for iter.Vaild() { // remove all
		iter.RemoveToPrev()
	}

	if l.String() != "[]" || iter.Value() != 0 { // default int == 0
		t.Error(l.String(), iter.Value())
	}

	l.Clear()
	for i := 0; i < 5; i++ {
		l.PushFront(i) // "[4 3 2 1 0]" cur:0
	}

	iter.ToHead() // 4
	if iter.Value() != 4 {
		t.Error(iter.Value())
	}
	iter.SetValue(1)
	if l.String() != "[1 3 2 1 0]" {
		t.Error(l.String())
	}

	other := l.Iterator()
	other.ToTail()

	iter.Swap(other)

	if l.String() != "[0 3 2 1 1]" {
		t.Error(l.String())
	}

	// log.Println(l.String(), l.Size())
}

func TestCircularIteratorIteratorRemove(t *testing.T) {
	l := New[int]()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	iter := l.CircularIterator()
	iter.Move(2)
	iter.RemoveToNext()

	// log.Println(l.String(), iter.Value()) //   4 -> 2(Remove) and Cur to 1  "[4 3 1 0]" 1
	if l.String() != "[4 3 1 0]" || iter.Value() != 1 {
		t.Error(l.String(), iter.Value())
	}

	iter.Move(-1) // 3

	// log.Println(l.String(), iter.Value())
	iter.RemoveToPrev() // 4

	iter.Move(3) // 4

	if l.String() != "[4 1 0]" || iter.Value() != 4 {
		t.Error(l.String(), iter.Value())
	}

	var result []string = []string{"[1 0]", "[0]", "[]"}
	for i := 0; iter.Vaild(); i++ {
		iter.RemoveToNext()
		if l.String() != result[i] {
			t.Error(l.String())
		}
	}

	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	iter.Move(10) // "[4 3 2 1 0]" cur:0

	result = []string{"[3 2 1 0]", "[3 2 1]", "[3 2]", "[3]", "[]"}
	for i := 0; iter.Vaild(); i++ { // can not remove tail or head
		iter.RemoveToPrev()
		if l.String() != result[i] {
			t.Error(l.String())
		}
	}

	for i := 0; i < 5; i++ {
		l.PushFront(i) // "[4 3 2 1 0]" cur:0
	}

	iter.ToHead() // 4
	if iter.Value() != 4 {
		t.Error(iter.Value())
	}
	iter.SetValue(1)
	if l.String() != "[1 3 2 1 0]" {
		t.Error(l.String())
	}

	other := l.CircularIterator()
	other.ToTail()

	iter.Swap(other)

	if l.String() != "[0 3 2 1 1]" {
		t.Error(l.String())
	}

	// log.Println(l.String(), l.Size())
}

// func BenchmarkPushBack(b *testing.B) {

// 	ec := 5
// 	cs := 2000000
// 	b.N = cs * ec

// 	for c := 0; c < ec; c++ {
// 		l := New[int]()
// 		for i := 0; i < cs; i++ {
// 			l.PushBack(i)
// 		}
// 	}
// }

// func BenchmarkPushFront(b *testing.B) {

// 	ec := 5
// 	cs := 2000000
// 	b.N = cs * ec

// 	for c := 0; c < ec; c++ {
// 		l := New[int]()
// 		for i := 0; i < cs; i++ {
// 			l.PushFront(i)
// 		}
// 	}

// }

// func BenchmarkInsert(b *testing.B) {

// 	ec := 10
// 	cs := 1000
// 	b.N = cs * ec

// 	for c := 0; c < ec; c++ {
// 		l := New[int]()
// 		for i := 0; i < cs; i++ {
// 			ridx := randomdata.Number(0, int(l.Size())+1)
// 			l.Insert(uint(ridx), i)
// 		}
// 	}
// }
