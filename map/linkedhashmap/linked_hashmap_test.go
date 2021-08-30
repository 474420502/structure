package linkedhashmap

import (
	"container/list"
	"log"
	"reflect"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/set/treeset"
)

func TestPush(t *testing.T) {
	lhm := New()
	lhm.PushFront(1, "1")
	lhm.PushBack("2", 2)
	var values []interface{}
	values = lhm.Values()

	var testType reflect.Type

	if testType = reflect.TypeOf(values[0]); testType.String() != "string" {
		t.Error(testType)
	}

	if testType = reflect.TypeOf(values[1]); testType.String() != "int" {
		t.Error(testType)
	}

	// 1 2
	lhm.PushFront(4, "4") // 4 1 2
	lhm.PushBack("3", 3)  // 4 1 2 3

	if lhm.String() != "[4 1 2 3]" {
		t.Error(lhm.String())
	}

	lhm.Put(5, 5)
	if lhm.String() != "[4 1 2 3 5]" {
		t.Error(lhm.String())
	}
}

func TestBase(t *testing.T) {
	lhm := New()
	for i := 0; i < 10; i++ {
		lhm.PushBack(i, i)
	}

	if lhm.Empty() {
		t.Error("why lhm Enpty, check it")
	}

	if lhm.Size() != 10 {
		t.Error("why lhm Size != 10, check it")
	}

	lhm.Clear()
	if !lhm.Empty() {
		t.Error("why lhm Clear not Empty, check it")
	}

	if lhm.Size() != 0 {
		t.Error("why lhm Size != 0, check it")
	}
}

func TestGet(t *testing.T) {
	lhm := New()
	for i := 0; i < 10; i++ {
		lhm.PushBack(i, i)
	}

	for i := 0; i < 10; i++ {
		lhm.PushBack(i, i)
	}

	if lhm.Size() != 10 {
		t.Error("why lhm Size != 10, check it")
	}

	for i := 0; i < 10; i++ {
		if v, ok := lhm.Get(i); !ok || v != i {
			t.Error("ok is ", ok, " get value is ", v)
		}
	}
}

func TestRemove(t *testing.T) {
	lhm := New()
	for i := 0; i < 10; i++ {
		lhm.PushBack(i, i)
	}

	var resultStr = "[0 1 2 3 4 5 6 7 8 9]"
	for i := 0; i < 10; i++ {
		if lhm.String() != resultStr {
			t.Error(lhm.String(), resultStr)
		}

		lhm.Remove(i)
		if lhm.Size() != 9-i {
			t.Error("why lhm Size != ", uint(9-i), ", check it")
		}

		resultStr = resultStr[0:1] + resultStr[3:]
	}

	if lhm.Size() != 0 {
		t.Error(lhm.Size())
	}

	for i := 0; i < 10; i++ {
		lhm.PushFront(i, i)
	}

	for i := 0; i < 10; i++ {
		if i >= 5 {
			lhm.Remove(i)
		}
	}

	if lhm.String() != "[4 3 2 1 0]" {
		t.Error(lhm.String())
	}

	// RemoveIndex [4 3 2 1 0]

}

func TestForce(t *testing.T) {
	rand := random.New()
	hm := New()
	set := treeset.New(compare.Int)
	l := list.New()

	for n := 0; n < 2000; n++ {

		if !set.Empty() {
			panic("")
		}

		if !hm.Empty() {
			panic("")
		}

		for i := 0; i < 20; i++ {
			v := rand.Intn(20)
			if rand.Bool() {
				if hm.Put(v, v) {
					l.PushBack(v)
				}
			} else {
				if hm.PushFront(v, v) {
					l.PushFront(v)
				}
			}
			set.Add(v)
		}

		for _, v := range set.Values() {
			if _, ok := hm.Get(v); !ok {
				log.Panicln(set.String())
			}
		}

		for _, k := range hm.Keys() {
			if ok := set.Contains(k); !ok {
				panic("")
			}
		}

		if set.Size() != hm.Size() {
			panic("")
		}

		for _, v := range hm.Values() {
			if ok := set.Contains(v); !ok {
				panic("")
			}
		}

		cur := l.Front()
		for _, v := range hm.Values() {
			if cur.Value != v {
				log.Println(cur.Value, v)
			}
			cur = cur.Next()
		}

		if cur != nil {
			panic("")
		}

		for i := 0; hm.Size() != 0 && i < 120; i++ {
			k := rand.Intn(100)
			if rand.OneOf64n(2) {
				if _, ok := hm.Remove(k); ok {
					cur := l.Front()
					for cur != nil {
						if cur.Value == k {
							l.Remove(cur)
							break
						}
						cur = cur.Next()
					}
				}
				set.Remove(k)
			}

			if rand.OneOf64n(3) {
				if rand.OneOf64n(2) {
					hm.Put(k, k)
					set.Add(k)
				} else {
					hm.PushFront(k, k)
					set.Add(k)
				}
			}

			if set.Size() != hm.Size() {
				panic("")
			}

		}

		for _, k := range hm.Keys() {
			if ok := set.Contains(k); !ok {
				panic("")
			}
		}

		for _, v := range hm.Values() {
			if ok := set.Contains(v); !ok {
				panic("")
			}
		}

		hm.Clear()
		set.Clear()
		l = l.Init()
	}
}
