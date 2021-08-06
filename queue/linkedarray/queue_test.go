package arrayqueue

import (
	"testing"

	testutils "github.com/474420502/structure"
)

func TestCasePut(t *testing.T) {
	q := New()

	for i := 0; i < 16; i++ {
		q.Push(i)
	}

	if len(q.Values()) != int(q.size) {
		panic("Values error")
	}

	if q.cap != int64(cap(q.data)) {
		panic("cap error")
	}

	if q.Peek() != 0 && q.Peek() != q.Index(0) {
		panic("check error")
	}
	// log.Println(q.Values(), "cap:", q.cap, cap(q.data))
}

func TestCasePop(t *testing.T) {
	var result []interface{}
	q := New()

	for i := 0; i < 16; i++ {
		q.Push(i)
	}

	result = q.Values()
	q.Traverse(func(idx int64, value interface{}) bool {
		if result[idx] != value {
			panic("check error")
		}
		return true
	})

	for i := 0; i < 8; i++ {
		q.Pop()
	}

	result = q.Values()
	q.Traverse(func(idx int64, value interface{}) bool {
		if result[idx] != value {
			panic("check error")
		}
		return true
	})

	if q.Peek() == 0 {
		panic("check")
	}

	if len(q.Values()) != int(q.size) {
		panic("Values error")
	}

	if q.cap != int64(cap(q.data)) {
		panic("cap error")
	}

	for i := 0; i < 4; i++ {
		q.Pop()
	}

	if len(q.Values()) != int(q.Size()) {
		panic("Values error")
	}

	result = q.Values()
	q.Traverse(func(idx int64, value interface{}) bool {
		if result[idx] != value {
			panic("check error")
		}
		return true
	})

	q.Push(100)
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
		q.Push(i)
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
