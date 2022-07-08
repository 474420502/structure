package treelist

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/474420502/random"
	utils "github.com/474420502/structure"
	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestIteratorRange(t *testing.T) {
	tree := New[[]byte, []byte](compare.ArrayAny[[]byte])
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

		iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
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
		iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
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
		iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
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
		iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
			result = append(result, string(cur.Key))
			return true
		})
		if fmt.Sprintf("%v", result) != "[c6 c4 c1 a5 a3 a1]" {
			t.Error(fmt.Sprintf("%v", result))
		}
	}()

}

func TestIteratorRangeForce(t *testing.T) {
	r := random.New()

	for n := 0; n < 1000; n++ {

		tree := New[[]byte, []byte](compare.ArrayAny[[]byte])
		var result [][]byte
		var start, end []byte
		for i := 0; i < 10; i++ {
			k := utils.Rangdom(1, 32, r)
			result = append(result, k)
			tree.Put(k, k)
		}

		start = result[r.Intn(len(result))]
		end = result[r.Intn(len(result))]
		if bytes.Compare(start, end) > 0 {
			start, end = end, start
		}
		// log.Println(string(start), string(end))
		// log.Println(tree.debugString(false))

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.GE2LT(start, end) // a4 <= key < c4
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				// log.Println("iter1:", string(cur.Key), cur.Index)
				result = append(result, string(cur.Key))
				return true
			})

			// log.Println("------------------------------------")
			var result2 []string
			iter2 := tree.Iterator()
			iter2.SeekGE(start)
			for iter2.Valid() {
				if iter2.Compare(end) >= 0 {
					break
				}
				result2 = append(result2, string(iter2.Key()))
				// log.Println("iter2:", string(iter2.Key()), iter2.Index())
				iter2.Next()
			}
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", result2) {
				panic("")
			}
			if iter.Size() != int64(len(result)) || iter.Size() != int64(len(result2)) {
				log.Println("start:", string(start), "end:", string(end))
				log.Println("siter:", string(iter.siter.cur.Key), "eiter:", string(iter.eiter.cur.Key), iter.siter.idx, iter.eiter.idx)
				log.Println("range size:", iter.Size(), "result len:", len(result))
				log.Println(result)
				log.Println()
			}

		}()

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.GT2LT(start, end) // a4 <= key < c4
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				// log.Println("iter1:", string(cur.Key), cur.Index)
				result = append(result, string(cur.Key))
				return true
			})

			// log.Println("------------------------------------")
			var result2 []string
			iter2 := tree.Iterator()
			iter2.SeekGT(start)
			for iter2.Valid() {
				if iter2.Compare(end) >= 0 {
					break
				}
				result2 = append(result2, string(iter2.Key()))
				// log.Println("iter2:", string(iter2.Key()), iter2.Index())
				iter2.Next()
			}
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", result2) {
				panic("")
			}
			if iter.Size() != int64(len(result)) || iter.Size() != int64(len(result2)) {
				log.Println("range size:", iter.Size(), "result len:", len(result))
				log.Println("start:", string(start), "end:", string(end))
				log.Println("siter:", string(iter.siter.cur.Key), "eiter:", string(iter.eiter.cur.Key), iter.siter.idx, iter.eiter.idx)
				log.Println(result)
				log.Println()
			}
		}()

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.GE2LE(start, end) // a4 <= key < c4
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				// log.Println("iter1:", string(cur.Key), cur.Index)
				result = append(result, string(cur.Key))
				return true
			})

			// log.Println("------------------------------------")
			var result2 []string
			iter2 := tree.Iterator()
			iter2.SeekGE(start)
			for iter2.Valid() {
				if iter2.Compare(end) > 0 {
					break
				}
				result2 = append(result2, string(iter2.Key()))
				// log.Println("iter2:", string(iter2.Key()), iter2.Index())
				iter2.Next()
			}
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", result2) {
				panic("")
			}
			if iter.Size() != int64(len(result)) || iter.Size() != int64(len(result2)) {
				log.Println("start:", string(start), "end:", string(end))
				log.Println("siter:", string(iter.siter.cur.Key), "eiter:", string(iter.eiter.cur.Key), iter.siter.idx, iter.eiter.idx)
				log.Println("range size:", iter.Size(), "result len:", len(result))
				log.Println(result)
				log.Println()
			}
		}()

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.GT2LE(start, end) // a4 <= key < c4
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				// log.Println("iter1:", string(cur.Key), cur.Index)
				result = append(result, string(cur.Key))
				return true
			})

			// log.Println("------------------------------------")
			var result2 []string
			iter2 := tree.Iterator()
			iter2.SeekGT(start)
			for iter2.Valid() {
				if iter2.Compare(end) > 0 {
					break
				}
				result2 = append(result2, string(iter2.Key()))
				// log.Println("iter2:", string(iter2.Key()), iter2.Index())
				iter2.Next()
			}
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", result2) {
				panic("")
			}
			if iter.Size() != int64(len(result)) || iter.Size() != int64(len(result2)) {
				log.Println("start:", string(start), "end:", string(end))
				log.Println("siter:", string(iter.siter.cur.Key), "eiter:", string(iter.eiter.cur.Key), iter.siter.idx, iter.eiter.idx)
				log.Println("range size:", iter.Size(), "result len:", len(result))
				log.Println(result)
				log.Println()
			}
		}()

		// -------------------------------------------------------
		// 简单的越界错误测试
		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.GE2LT([]byte("a4"), []byte("c4")) // a4 <= key < c4
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				result = append(result, string(cur.Key))
				return true
			})
			// if fmt.Sprintf("%v", result) != "[a5 c1]" {
			// 	t.Error(fmt.Sprintf("%v", result))
			// }
		}()

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.GT2LT([]byte("a4"), []byte("c9"))
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				result = append(result, string(cur.Key))
				return true
			})
			// if fmt.Sprintf("%v", result) != "[a5 c1 c4 c6]" {
			// 	t.Error(fmt.Sprintf("%v", result))
			// }
		}()

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.GE2LE([]byte("a0"), []byte("c9"))
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				result = append(result, string(cur.Key))
				return true
			})
			// if fmt.Sprintf("%v", result) != "[a1 a3 a5 c1 c4 c6]" {
			// 	t.Error(fmt.Sprintf("%v", result))
			// }
		}()

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.SetDirection(Reverse)
			iter.GT2LE([]byte("a0"), []byte("c9"))
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				result = append(result, string(cur.Key))
				return true
			})
			// if fmt.Sprintf("%v", result) != "[c6 c4 c1 a5 a3 a1]" {
			// 	t.Error(fmt.Sprintf("%v", result))
			// }
		}()
	}
}

func TestIteratorRangeForce2(t *testing.T) {
	r := random.New()

	for n := 0; n < 1000; n++ {

		tree := New[[]byte, []byte](compare.ArrayAny[[]byte])
		var result [][]byte
		var start, end []byte
		for i := 0; i < 10; i++ {
			k := utils.Rangdom(1, 32, r)
			result = append(result, k)
			tree.Put(k, k)
		}

		start = result[r.Intn(len(result))]
		end = result[r.Intn(len(result))]
		if bytes.Compare(start, end) > 0 {
			start, end = end, start
		}
		// log.Println(string(start), string(end))
		// log.Println(tree.debugString(false))

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.SetDirection(Reverse)
			iter.GE2LT(start, end) // a4 <= key < c4
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				// log.Println("iter1:", string(cur.Key), cur.Index)
				result = append(result, string(cur.Key))
				return true
			})

			// log.Println("------------------------------------")
			var result2 []string
			iter2 := tree.Iterator()
			iter2.SeekLT(end)
			for iter2.Valid() {
				if iter2.Compare(start) < 0 {
					break
				}
				result2 = append(result2, string(iter2.Key()))
				// log.Println("iter2:", string(iter2.Key()), iter2.Index())
				iter2.Prev()
			}

			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", result2) {
				panic("")
			}

			if iter.Size() != int64(len(result)) || iter.Size() != int64(len(result2)) {
				log.Println("start:", string(start), "end:", string(end))
				log.Println("siter:", string(iter.siter.cur.Key), "eiter:", string(iter.eiter.cur.Key), iter.siter.idx, iter.eiter.idx)
				log.Println("range size:", iter.Size(), "result len:", len(result))
				log.Println(result)
				log.Println()
			}

		}()

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.SetDirection(Reverse)
			iter.GT2LT(start, end) // a4 <= key < c4
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				// log.Println("iter1:", string(cur.Key), cur.Index)
				result = append(result, string(cur.Key))
				return true
			})

			// log.Println("------------------------------------")
			var result2 []string
			iter2 := tree.Iterator()
			iter2.SeekLT(end)
			for iter2.Valid() {
				if iter2.Compare(start) <= 0 {
					break
				}
				result2 = append(result2, string(iter2.Key()))
				// log.Println("iter2:", string(iter2.Key()), iter2.Index())
				iter2.Prev()
			}

			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", result2) {
				panic("")
			}

			if iter.Size() != int64(len(result)) || iter.Size() != int64(len(result2)) {
				log.Println("range size:", iter.Size(), "result len:", len(result))
				log.Println("start:", string(start), "end:", string(end))
				log.Println("siter:", string(iter.siter.cur.Key), "eiter:", string(iter.eiter.cur.Key), iter.siter.idx, iter.eiter.idx)
				log.Println(result)
				log.Println()
			}
		}()

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.SetDirection(Reverse)
			iter.GE2LE(start, end) // a4 <= key < c4
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				// log.Println("iter1:", string(cur.Key), cur.Index)
				result = append(result, string(cur.Key))
				return true
			})

			// log.Println("------------------------------------")
			var result2 []string
			iter2 := tree.Iterator()
			iter2.SeekLE(end)
			for iter2.Valid() {
				if iter2.Compare(start) < 0 {
					break
				}
				result2 = append(result2, string(iter2.Key()))
				// log.Println("iter2:", string(iter2.Key()), iter2.Index())
				iter2.Prev()
			}

			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", result2) {
				panic("")
			}

			if iter.Size() != int64(len(result)) || iter.Size() != int64(len(result2)) {
				log.Println("start:", string(start), "end:", string(end))
				log.Println("siter:", string(iter.siter.cur.Key), "eiter:", string(iter.eiter.cur.Key), iter.siter.idx, iter.eiter.idx)
				log.Println("range size:", iter.Size(), "result len:", len(result))
				log.Println(result)
				log.Println()
			}
		}()

		func() {
			var result []string
			iter := tree.IteratorRange()
			iter.SetDirection(Reverse)
			iter.GT2LE(start, end) // a4 <= key < c4
			iter.Range(func(cur *SliceIndex[[]byte, []byte]) bool {
				// log.Println("iter1:", string(cur.Key), cur.Index)
				result = append(result, string(cur.Key))
				return true
			})

			// log.Println("------------------------------------")
			var result2 []string
			iter2 := tree.Iterator()
			iter2.SeekLE(end)
			for iter2.Valid() {
				if iter2.Compare(start) <= 0 {
					break
				}
				result2 = append(result2, string(iter2.Key()))
				// log.Println("iter2:", string(iter2.Key()), iter2.Index())
				iter2.Prev()
			}

			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", result2) {
				panic("")
			}

			if iter.Size() != int64(len(result)) || iter.Size() != int64(len(result2)) {
				log.Println("start:", string(start), "end:", string(end))
				log.Println("siter:", string(iter.siter.cur.Key), "eiter:", string(iter.eiter.cur.Key), iter.siter.idx, iter.eiter.idx)
				log.Println("range size:", iter.Size(), "result len:", len(result))
				log.Println(result)
				log.Println()
			}
		}()
	}
}
