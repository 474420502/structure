package avl

import (
	"log"
	"sort"
	"testing"

	"github.com/474420502/random"
	utils "github.com/474420502/structure"
	"github.com/474420502/structure/compare"
)

func TestNextPrev(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for i := 0; i < 10; i++ {
		tree.Set(i, i)
	}

	iter := tree.Iterator()
	iter.SeekToFirst()
	if !iter.Valid() {
		panic("")
	}

	if iter.Value() != 0 {
		panic("")
	}

	for i := 0; i < 10; i++ {

		if !iter.Valid() || iter.Value() != i {
			panic("")
		}
		log.Println(iter.view())
		iter.Next()
	}

	iter.Prev()
	log.Println(iter.Valid(), iter.Key())
	if !iter.Valid() && iter.Value() != 9 {
		panic("")
	}

	for i := 9; i >= 0; i-- {
		if iter.Value() != i {
			panic("")
		}
		iter.Prev()
	}

	iter.SeekGE(5)
	for i := 5; i >= 0; i-- {
		if iter.Value() != i {
			panic("")
		}
		iter.Prev()
	}

	iter.SeekGT(5)
	for i := 6; i < 10; i++ {
		if iter.Value() != i {
			log.Panic(iter.Value())
		}
		iter.Next()
	}

	iter.SeekLT(5)
	for i := 4; i >= 0; i-- {
		if iter.Value() != i {
			log.Panic(iter.Value())
		}
		iter.Prev()
	}
}

func TestDefault(t *testing.T) {

	// │           ┌── 92(2)
	// │           │   └── 85(1)
	// │       ┌── 72(3)
	// │       │   └── 70(2)
	// │       │       └── 69(1)
	// │   ┌── 64(4)
	// │   │   │   ┌── 56(2)
	// │   │   │   │   └── 43(1)
	// │   │   └── 41(3)
	// │   │       │   ┌── 39(1)
	// │   │       └── 32(2)
	// └── 27(5)
	//     │       ┌── 20(1)
	//     │   ┌── 18(2)
	//     └── 16(3)
	//         │   ┌── 13(1)
	//         └── 4(2)
	//             └── 0(1)

	var data []int = []int{4, 0, 41, 27, 64, 13, 16, 32, 18, 39, 56, 70, 20, 43, 72, 92, 85, 69}
	tree := New[int, int](compare.AnyEx[int])
	for _, v := range data {
		tree.Put(v, v)
	}
	// log.Println(tree.view())

	iter := tree.Iterator()

	iter.SeekLE(93) //
	if !iter.Valid() && iter.Key() != 92 {
		panic("SeekLE 93")
	}

	iter.Prev()
	if !iter.Valid() && iter.Key() != 85 {
		panic("Prev 85")
	}

	iter.SeekLT(27)
	if !iter.Valid() && iter.Key() != 20 {
		panic("SeekLT 27")
	}
	a := iter.Clone()

	iter.SeekGE(65)
	if !iter.Valid() && iter.Key() != 69 {
		panic("SeekGE 69")
	}

	a.Next()
	if !a.Valid() && a.Key() != 27 {
		panic("a.Next() 20 -> 27")
	}

	a.Prev()
	if !a.Valid() && a.Key() != 20 {
		panic("a.Prev() 20 <- 27")
	}

	if !iter.Valid() && iter.Key() != 69 {
		panic("SeekGE 69")
	}

	iter.SeekToFirst()
	if !iter.Valid() && iter.Key() != 0 {
		panic("SeekToFirst")
	}
	iter.Prev() // 0 ->  无穷小
	if iter.Valid() {
		panic("0 ->>>> min")
	}
	iter.Prev()
	if iter.Valid() {
		log.Println(iter.Key())
		panic("min ->>>> min")
	}
	iter.Next()
	if !iter.Valid() {
		panic("min ->>>> 0")
	}

	iter.SeekToLast()
	if !iter.Valid() && iter.Key() != 92 {
		panic("SeekToFirst")
	}

	iter.Next()
	if iter.Valid() {
		panic("92 ->>>> max")
	}
	iter.Next()
	if iter.Valid() {
		panic("max ->>>> max")
	}

	iter.Prev()
	if !iter.Valid() {
		panic("max ->>>> 92")
	}

}

func TestSeekFor(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for i := 0; i < 20; i += 2 {
		tree.Set(i, i)
	}

	log.Println(tree.view())

	iter := tree.Iterator()
	iter.SeekLE(7)
	if !iter.Valid() { // Key == 6
		t.Error("SeekLE return is error")
	}
	for i := 6; i >= 0; i -= 2 {
		if iter.Value() != i {
			panic("")
		}
		iter.Prev()
	}

	iter.SeekLE(7)
	if !iter.Valid() {
		t.Error("SeekLE return is error")
	}
	for i := 6; i < 20; i += 2 {
		if iter.Value() != i {
			panic("")
		}
		iter.Next()
	}

	iter.SeekLE(6)
	if !iter.Valid() { // key == 6
		t.Error("SeekLE return is error")
	}
	for i := 6; i < 20; i += 2 {
		if iter.Value() != i {
			panic("")
		}
		iter.Next()
	}

	iter.SeekGE(7)
	if !iter.Valid() { // Key == 8
		t.Error("SeekGE return is error")
	}
	for i := 8; i < 20; i += 2 {
		if iter.Value() != i {
			panic("")
		}
		iter.Next()
	}

	iter.SeekGE(8)
	if !iter.Valid() { // Key == 8
		t.Error("SeekGE return is error")
	}
	for i := 8; i < 20; i += 2 {
		if iter.Value() != i {
			panic("")
		}
		iter.Next()
	}
}

func TestIteratorForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree := New[int, int](compare.AnyEx[int])
		var priority []int
		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			if tree.Put(v, v) {
				priority = append(priority, v)
			}

		}
		sort.Slice(priority, func(i, j int) bool {
			return priority[i] < priority[j]
		})

		s := rand.Intn(100)

		// log.Println(priority, idx, len(priority))

		iter := tree.Iterator()
		iter.SeekGE(s)

		idx := sort.Search(len(priority), func(i int) bool {
			return priority[i] >= s
		})

		if idx == len(priority) {
			if iter.Valid() {
				log.Panicln(iter.Value(), priority)
			}
		} else {
			for i := idx; i < int(tree.Size()); i++ {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Next()
			}
			iter.SeekGE(s)
			for i := idx; i >= 0; i-- {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Prev()
			}
		}

		iter.SeekLE(s)
		idx = sort.Search(len(priority), func(i int) bool {
			return priority[i] > s
		})

		if idx-1 < 0 {
			if iter.Valid() {
				log.Panicln(iter.Value(), priority, idx-1, s)
			}
		} else {
			for i := idx - 1; i < int(tree.Size()); i++ {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Next()
			}
			iter.SeekLE(s)
			for i := idx - 1; i >= 0; i-- {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Prev()
			}
		}

		iter.SeekToFirst()
		for i := 0; i < len(priority); i++ {
			if iter.Value() != priority[i] {
				log.Panic("")
			}
			iter.Next()
		}

		iter.SeekToLast()
		for i := len(priority) - 1; i >= 0; i-- {
			if iter.Value() != priority[i] {
				log.Panic("")
			}
			iter.Prev()
		}

	}
}

func TestIteratorForce2(t *testing.T) {
	rand := random.New(1683989312052736623)
	for n := 0; n < 2000; n++ {
		tree := New[int, int](compare.AnyEx[int])
		var priority []int
		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			if tree.Put(v, v) {
				priority = append(priority, v)
			}

		}
		sort.Slice(priority, func(i, j int) bool {
			return priority[i] < priority[j]
		})

		s := rand.Intn(100)

		// log.Println(priority, idx, len(priority))

		iter := tree.Iterator()
		iter.SeekGT(s)

		idx := sort.Search(len(priority), func(i int) bool {
			return priority[i] > s
		})

		if idx == len(priority) {
			if iter.Valid() {
				log.Panicln(iter.Value(), priority)
			}
		} else {
			for i := idx; i < int(tree.Size()); i++ {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Next()
			}

			iter.SeekGT(s)
			for i := idx; i >= 0; i-- {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Prev()
			}
		}

		iter.SeekLT(s)
		idx = sort.Search(len(priority), func(i int) bool {
			return priority[i] >= s
		})
		// if priority[idx] != s {
		// 	idx--
		// }

		if idx-1 < 0 {
			if iter.Valid() {
				log.Panicln(iter.Value(), priority, idx-1, s)
			}
		} else {
			for i := idx - 1; i < int(tree.Size()); i++ {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Next()
			}
			iter.SeekLT(s)
			for i := idx - 1; i >= 0; i-- {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Prev()
			}
		}

		iter.SeekToFirst()
		for i := 0; i < len(priority); i++ {
			if iter.Value() != priority[i] {
				log.Panic("")
			}
			iter.Next()
		}

		iter.SeekToLast()
		for i := len(priority) - 1; i >= 0; i-- {
			if iter.Value() != priority[i] {
				log.Panic("")
			}
			iter.Prev()
		}

	}
}

func TestDefaultSeek(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for i := 0; i < 10; i += 2 {
		tree.Set(i, i)
	}
	utils.Expect("[0 2 4 6 8]", tree.Values())
	iter := tree.Iterator()

	// 测试 SeekLE 和 SeekLT 在树为空时的情况
	emptyTree := New[int, int](compare.AnyEx[int])
	emptyIter := emptyTree.Iterator()

	utils.Expect("false false", emptyIter.SeekLE(-1), emptyIter.Valid())
	utils.Expect("false false", emptyIter.SeekLT(-1), emptyIter.Valid())

	// 测试 SeekGE 和 SeekGT 在树为空时的情况
	utils.Expect("false false", emptyIter.SeekGE(10), emptyIter.Valid())
	utils.Expect("false false", emptyIter.SeekGT(10), emptyIter.Valid())

	// 测试在树不为空时的情况
	utils.Expect("false false", iter.SeekLE(-1), iter.Valid())
	utils.Expect("false false", iter.SeekLT(-1), iter.Valid())

	utils.Expect("false false", iter.SeekGE(10), iter.Valid())
	utils.Expect("false false", iter.SeekGT(10), iter.Valid())

	utils.Expect("true true 6", iter.SeekLE(6), iter.Valid(), iter.Value())
	utils.Expect("true true 4", iter.SeekLT(6), iter.Valid(), iter.Value())
	utils.Expect("false true 4", iter.SeekLE(5), iter.Valid(), iter.Value())
	utils.Expect("false true 4", iter.SeekLT(5), iter.Valid(), iter.Value())
	utils.Expect("true true 6", iter.SeekGE(6), iter.Valid(), iter.Value())
	utils.Expect("true true 8", iter.SeekGT(6), iter.Valid(), iter.Value())
	utils.Expect("false true 6", iter.SeekGE(5), iter.Valid(), iter.Value())
	utils.Expect("false true 6", iter.SeekGT(5), iter.Valid(), iter.Value())

	// 测试在树中间插入新元素后的情况
	tree.Set(5, 5)
	utils.Expect("[0 2 4 5 6 8]", tree.Values())
	iter = tree.Iterator()
	utils.Expect("true true 5", iter.SeekLE(5), iter.Valid(), iter.Value())
	utils.Expect("true true 4", iter.SeekLT(5), iter.Valid(), iter.Value())
	utils.Expect("true true 5", iter.SeekGE(5), iter.Valid(), iter.Value())
	utils.Expect("true true 6", iter.SeekGT(5), iter.Valid(), iter.Value())

	// 测试在树头部插入新元素后的情况
	tree.Set(-2, -2)
	utils.Expect("[-2 0 2 4 5 6 8]", tree.Values())
	iter = tree.Iterator()
	utils.Expect("true true -2", iter.SeekLE(-2), iter.Valid(), iter.Value())
	utils.Expect("false false", iter.SeekLT(-3), iter.Valid())
	utils.Expect("true true -2", iter.SeekGE(-2), iter.Valid(), iter.Value())
	utils.Expect("false true -2", iter.SeekGT(-3), iter.Valid(), iter.Value())

	// 测试在树尾部插入新元素后的情况
	tree.Set(10, 10)
	utils.Expect("[-2 0 2 4 5 6 8 10]", tree.Values())
	iter = tree.Iterator()
	utils.Expect("false true 10", iter.SeekLE(11), iter.Valid(), iter.Value())
	utils.Expect("true true 8", iter.SeekLT(10), iter.Valid(), iter.Value())
	utils.Expect("false false", iter.SeekGE(11), iter.Valid())
	utils.Expect("false false", iter.SeekGT(11), iter.Valid())

	// 测试从中间删除元素后的情况
	tree.Remove(4)
	utils.Expect("[-2 0 2 5 6 8 10]", tree.Values())
	iter = tree.Iterator()
	utils.Expect("true true 5", iter.SeekLE(5), iter.Valid(), iter.Value())
	utils.Expect("true true 2", iter.SeekLT(5), iter.Valid(), iter.Value())
	utils.Expect("true true 5", iter.SeekGE(5), iter.Valid(), iter.Value())
	utils.Expect("true true 6", iter.SeekGT(5), iter.Valid(), iter.Value())

	// 测试从头部删除元素后的情况
	tree.Remove(-2)
	utils.Expect("[0 2 5 6 8 10]", tree.Values())
	iter = tree.Iterator()
	utils.Expect("false false", iter.SeekLE(-1), iter.Valid())
	utils.Expect("false false", iter.SeekLT(-1), iter.Valid())
	utils.Expect("true true 0", iter.SeekGE(0), iter.Valid(), iter.Value())
	utils.Expect("true true 2", iter.SeekGT(0), iter.Valid(), iter.Value())

	// 插入重复元素
	tree.Set(5, 5)
	utils.Expect("[0 2 5 6 8 10]", tree.Values())
	iter = tree.Iterator()
	utils.Expect("true true 5", iter.SeekLE(5), iter.Valid(), iter.Value())
	utils.Expect("true true 2", iter.SeekLT(5), iter.Valid(), iter.Value())
	utils.Expect("true true 5", iter.SeekGE(5), iter.Valid(), iter.Value())
	utils.Expect("true true 6", iter.SeekGT(5), iter.Valid(), iter.Value())

	// 测试查找比最大值更大的元素
	utils.Expect("false true 10", iter.SeekLE(11), iter.Valid(), iter.Value())
	utils.Expect("false true 10", iter.SeekLT(11), iter.Valid(), iter.Value())
	utils.Expect("false false", iter.SeekGE(11), iter.Valid())
	utils.Expect("false false", iter.SeekGT(11), iter.Valid())

	// 测试查找比最小值更小的元素
	utils.Expect("false false", iter.SeekLE(-1), iter.Valid())
	utils.Expect("false false", iter.SeekLT(-1), iter.Valid())
	utils.Expect("false true 0", iter.SeekGE(-1), iter.Valid(), iter.Value())
	utils.Expect("false true 0", iter.SeekGT(-1), iter.Valid(), iter.Value())
}

// func TestCompareSimilarForce(t *testing.T) {
// 	tree1 := treelist.New()
// 	tree2 := New(compare.ArrayAny[[]byte])

// 	rand := random.New()

// 	for i := 0; i < 1000; i++ {

// 		var buf []byte
// 		for n := 1; n < 64; n++ {
// 			buf = append(buf, byte(rand.Intn(256)))
// 		}
// 		is := tree1.Set(buf, buf)
// 		if tree2.Set(buf, buf) != is {
// 			t.Error("tree1 Set is not equal to tree2")
// 			panic(nil)
// 		}

// 		iter1 := tree1.Iterator()
// 		iter2 := tree2.Iterator()

// 		for n := 1; n < 10; n++ {
// 			buf = append(buf, byte(rand.Intn(256)))
// 		}

// 		if iter1.SeekGE(buf) != iter2.SeekGE(buf) {
// 			t.Error("SeekGE")
// 			panic(nil)
// 		}

// 		for iter1.Valid() && iter2.Vaild() {
// 			if !bytes.Equal(iter1.Key(), iter2.Key()) {
// 				t.Error("SeekGE")
// 				panic(nil)
// 			}
// 			iter1.Next()
// 			iter2.Next()
// 		}

// 		if iter1.SeekGE(buf) != iter2.SeekGE(buf) {
// 			t.Error("SeekGE")
// 			panic(nil)
// 		}

// 		for iter1.Valid() && iter2.Vaild() {
// 			if !bytes.Equal(iter1.Key(), iter2.Key()) {
// 				t.Error("SeekGE")
// 				panic(nil)
// 			}
// 			iter1.Prev()
// 			iter2.Prev()
// 		}

// 		if iter1.SeekLE(buf) != iter2.SeekLE(buf) {
// 			t.Error("SeekLE")
// 			panic(nil)
// 		}

// 		for iter1.Valid() && iter2.Vaild() {
// 			if !bytes.Equal(iter1.Key(), iter2.Key()) {
// 				t.Error("SeekLE")
// 				panic(nil)
// 			}
// 			iter1.Next()
// 			iter2.Next()
// 		}

// 		if iter1.SeekLE(buf) != iter2.SeekLE(buf) {
// 			t.Error("SeekLE")
// 			panic(nil)
// 		}

// 		for iter1.Valid() && iter2.Vaild() {
// 			if !bytes.Equal(iter1.Key(), iter2.Key()) {
// 				t.Error("SeekLE")
// 				panic(nil)
// 			}
// 			iter1.Prev()
// 			iter2.Prev()
// 		}

// 	}
// }
