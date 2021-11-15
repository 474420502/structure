package treelist

import (
	"fmt"
	"testing"

	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestIteratorRange(t *testing.T) {
	tree := New()
	for _, v := range testutils.TestedBytesSimlpe {
		tree.Put(v, v)
	}

	//	│       ┌── c6
	//	│   ┌── c4
	//	└── c1
	//		│   ┌── a5
	//		└── a3
	//			└── a1
	// [a1 a3 a5 c1 c4 c6]

	func() {
		var result []string
		iter := tree.IteratorRange()
		iter.GE2LT([]byte("a4"), []byte("c4")) // a4 <= key < c4
		iter.Range(func(cur *SliceIndex) bool {
			result = append(result, string(cur.Key))
			return true
		})
		if fmt.Sprintf("%v", result) != "[a5 c1]" {
			t.Error(fmt.Sprintf("%v", result))
		}
	}()

	func() {
		var result []string
		iter := tree.IteratorRange()
		iter.GT2LT([]byte("a4"), []byte("c9"))
		iter.Range(func(cur *SliceIndex) bool {
			result = append(result, string(cur.Key))
			return true
		})
		if fmt.Sprintf("%v", result) != "[a5 c1 c4 c6]" {
			t.Error(fmt.Sprintf("%v", result))
		}
	}()

	func() {
		var result []string
		iter := tree.IteratorRange()
		iter.GE2LE([]byte("a0"), []byte("c9"))
		iter.Range(func(cur *SliceIndex) bool {
			result = append(result, string(cur.Key))
			return true
		})
		if fmt.Sprintf("%v", result) != "[a1 a3 a5 c1 c4 c6]" {
			t.Error(fmt.Sprintf("%v", result))
		}
	}()

	func() {
		var result []string
		iter := tree.IteratorRange()
		iter.SetDirection(Reverse)
		iter.GT2LE([]byte("a0"), []byte("c9"))
		iter.Range(func(cur *SliceIndex) bool {
			result = append(result, string(cur.Key))
			return true
		})
		if fmt.Sprintf("%v", result) != "[c6 c4 c1 a5 a3 a1]" {
			t.Error(fmt.Sprintf("%v", result))
		}
	}()

}
