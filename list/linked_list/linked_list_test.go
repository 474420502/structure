package linkedlist

import (
	"fmt"
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

func TestRemove(t *testing.T) {
	l := New[int]()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	l.Remove(0)
	var result string
	result = fmt.Sprintf("%v", l.Values())
	if result != "[3 2 1 0]" {
		t.Error("should be [3 2 1 0] but result is", result)
	}

	l.Remove(3)
	result = fmt.Sprintf("%v", l.Values())
	if result != "[3 2 1]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	l.Remove(2)
	result = fmt.Sprintf("%v", l.Values())
	if result != "[3 2]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	l.Remove(1)
	result = fmt.Sprintf("%v", l.Values())
	if result != "[3]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	l.Remove(0)
	result = fmt.Sprintf("%v", l.Values())
	if result != "[]" && l.Size() == 0 && len(l.Values()) == 0 {
		t.Error("should be [] but result is", result, "Size is", l.Size())
	}

	if _, rvalue := l.Remove(3); rvalue != false {
		t.Error("l is empty")
	}

}

func TestRemoveIf(t *testing.T) {
	l := New[int]()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	if result, ok := l.RemoveIf(func(idx uint, value int) RemoveState {
		if value == 0 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {
		if result[0] != 0 {
			t.Error("result should is", 0)
		}
	} else {
		t.Error("should be ok")
	}

	if result, ok := l.RemoveIf(func(idx uint, value int) RemoveState {
		if value == 4 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {
		if result[0] != 4 {
			t.Error("result should is", 4)
		}
	} else {
		t.Error("should be ok")
	}

	var result string
	result = fmt.Sprintf("%v", l.Values())
	if result != "[3 2 1]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	if result, ok := l.RemoveIf(func(idx uint, value int) RemoveState {
		if value == 4 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {
		t.Error("should not be ok and result is nil")
	} else {
		if result != nil {
			t.Error("should be nil")
		}
	}

	result = fmt.Sprintf("%v", l.Values())
	if result != "[3 2 1]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	l.RemoveIf(func(idx uint, value int) RemoveState {
		if value == 3 || value == 2 || value == 1 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	})

	result = fmt.Sprintf("%v", l.Values())
	if result != "[]" {
		t.Error("result should be [], but now result is", result)
	}

	if results, ok := l.RemoveIf(func(idx uint, value int) RemoveState {
		if value == 3 || value == 2 || value == 1 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {
		t.Error("why  ok")
	} else {
		if results != nil {
			t.Error(results)
		}
	}

}

func TestRemoveIf2(t *testing.T) {
	l := New[int]()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	// 只删除一个
	if result, ok := l.RemoveIf(func(idx uint, value int) RemoveState {
		if value == 0 {
			return RemoveAndBreak
		}
		return UnremoveAndContinue
	}); ok {
		if result[0] != 0 {
			t.Error("result should is", 0)
		}
	} else {
		t.Error("should be ok")
	}

	var resultstr string
	resultstr = fmt.Sprintf("%v", l.Values())
	if resultstr != "[4 3 2 1 4 3 2 1 0]" {
		t.Error("result should is", resultstr)
	}

	// 只删除多个
	if result, ok := l.RemoveIf(func(idx uint, value int) RemoveState {
		if value == 4 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {

		resultstr = fmt.Sprintf("%v", result)
		if resultstr != "[4 4]" {
			t.Error("result should is", result)
		}

		resultstr = fmt.Sprintf("%v", l.Values())
		if resultstr != "[3 2 1 3 2 1 0]" {
			t.Error("result should is", resultstr)
		}

	} else {
		t.Error("should be ok")
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

	for i := 0; iter.Next(); i++ {
		if iter.Value() != 9-i {
			t.Error("iter.Next() ", iter.Value(), "is not equal ", 9-i)
		}
	}

	if iter.cur != iter.ll.tail {
		t.Error("current point is not equal tail ", iter.ll.tail)
	}

	for i := 0; iter.Prev(); i++ {
		if iter.Value() != i {
			t.Error("iter.Prev() ", iter.Value(), "is not equal ", i)
		}
	}
}

func TestCircularIterator(t *testing.T) {
	ll := New[int]()
	for i := 0; i < 10; i++ {
		ll.PushFront(i)
	}

	iter := ll.CircularIterator()

	for i := 0; i != 10; i++ {
		iter.Next()
		if iter.Value() != 9-i {
			t.Error("iter.Next() ", iter.Value(), "is not equal ", 9-i)
		}
	}

	if iter.cur != iter.ll.tail.prev {
		t.Error("current point is not equal tail ", iter.ll.tail.prev)
	}

	if iter.Next() {
		if iter.Value() != 9 {
			t.Error("iter.Value() != ", iter.Value())
		}
	}

	iter.ToTail()
	for i := 0; i != 10; i++ {
		iter.Prev()
		if iter.Value() != i {
			t.Error("iter.Prev() ", iter.Value(), "is not equal ", i)
		}
	}

	if iter.cur != iter.ll.head.next {
		t.Error("current point is not equal tail ", iter.ll.tail.prev)
	}

	if iter.Prev() {
		if iter.Value() != 0 {
			t.Error("iter.Value() != ", iter.Value())
		}
	}
}

func TestContains(t *testing.T) {
	ll := New[int]()
	for i := 0; i < 10; i++ {
		ll.Push(i)
	}

	for i := 0; i < 10; i++ {
		if !ll.Contains(i) {
			t.Error(i)
		}
	}

	for i := 10; i < 20; i++ {
		if ll.Contains(i) {
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
				for i := 0; i <= int(idx); i++ {
					iter.Next()
				}
				iter.InsertFront(v)

				var i = uint(0)
				l.Traverse(func(value int) bool {
					if i == idx {
						if value != v {
							panic("")
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
		citer := l.CircularIterator()
		for citer.Next() {
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

	// log.Println(l.String(), iter.Value()) // Head -> 4 -> 3(Remove) and Cur to 2  "[4 2 1 0]" 2
	if l.String() != "[4 2 1 0]" || iter.Value() != 2 {
		t.Error(l.String(), iter.Value())
	}

	iter.Move(-1) // 4
	// log.Println(l.String(), iter.Value())
	iter.RemoveToPrev() // Head
	iter.Move(3)        // 1

	if l.String() != "[2 1 0]" || iter.Value() != 0 {
		t.Error(l.String(), iter.Value())
	}

	for iter.RemoveToNext() {

	}

	if l.String() != "[2 1]" || iter.Value() != 0 {
		t.Error(l.String(), iter.Value())
	}

	for iter.RemoveToPrev() { // can not remove tail or head

	}

	if l.String() != "[2 1]" || iter.Value() != 0 {
		t.Error(l.String(), iter.Value())
	}

	iter.Move(-1)             // 1
	for iter.RemoveToPrev() { // remove all

	}

	if l.String() != "[]" || iter.Value() != 0 { // default int == 0
		t.Error(l.String(), iter.Value())
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
