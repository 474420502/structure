package itree

import (
	"log"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

func TestCase(t *testing.T) {

	r := random.New(1683745535466910261)

	tree := New[int, int](compare.AnyEx[int])
	var s []int
	count := 40
	for i := 0; i < count; i++ {
		v := r.Intn(100)

		if tree.Put(v, v) {
			s = append(s, v)
		}

	}
	log.Printf("%v len:%d", s, len(s))
	tree.check()

	// log.Println(tree.view())

	// tree.Trim(70, 70)

	iter := tree.Iterator()

	iter.SeekLE(93) //
	log.Println(iter.Vaild())
	log.Println(iter.Key(), iter.Value())
	iter.Prev()
	log.Println(iter.Vaild())
	log.Println(iter.Key(), iter.Value())

	iter.SeekLT(27)
	log.Println(iter.Vaild())
	log.Println(iter.Key(), iter.Value())
	a := iter.Clone()

	iter.SeekGE(65)
	log.Println(iter.Vaild())
	log.Println(iter.Key(), iter.Value())

	log.Println(a.Vaild())
	log.Println(a.Key(), a.Value())

	log.Println(iter, a)

	iter.SeekToFirst()
	log.Println(iter.Vaild())
	log.Println(iter.Key(), iter.Value())

	iter.SeekToLast()

	iter.Prev()
	log.Println(iter.Vaild())
	log.Println(iter.Key(), iter.Value())
	iter.Prev()
	log.Println(iter.Vaild())
	log.Println(iter.Key(), iter.Value())
}

func TestCasePut(t *testing.T) {

	r := random.New()

	for n := 0; n < 10000; n++ {
		tree := New[int, int](compare.AnyEx[int])
		count := r.Intn(100)
		for i := 0; i < count; i++ {
			v := r.Intn(100)
			tree.Put(v, v)
			tree.check()
		}
		// log.Println(tree.View())
	}

	// log.Println(tree.View())
}

func TestCaseR(t *testing.T) {
	log.SetFlags(log.Llongfile)

	r := random.New(1684430584192166479)

	for nn := 0; nn < 1000; nn++ {

		tree := New[int, int](compare.AnyEx[int])
		count := r.Intn(50) + 50

		var checkv []int

		for i := 0; i < count; i++ {
			v := r.Intn(64)
			if !tree.Put(v, v) {
				checkv = append(checkv, v)
			}
			tree.check()
		}
		// tree.check()
		var isDebug = false
		// log.Println(tree.View())
		// log.Println("remove:", rv, "remove list:", checkv)
		for i := 0; i < r.Intn(8); i++ {
			var rv int = r.Intn(64)
			tree.Remove(rv)
			tree.check()
		}

		// log.Println(tree.Remove(rv))
		// log.Println(tree.View())
		for _, v := range checkv {
			// if v == 38 {
			// 	log.Println("")
			// }

			beforeView := tree.view()
			if _, ok := tree.Remove(v); ok {
				if isDebug {
					log.Print(beforeView, " key: ", v, "", tree.view(), "\n")
				}
			}
			tree.check()
		}
	}
	// log.Println(tree.View())
}

func TestCaseMainIndex(t *testing.T) {
	log.SetFlags(log.Llongfile)

	r := random.New(1684430584192166479)

	tree := New[int, int](compare.AnyEx[int])
	count := r.Intn(40)
	for i := 0; i < count; i++ {
		v := r.Intn(100)
		tree.Put(v, v)
		tree.check()
	}

	log.Println(tree.view())
	// tree.TrimByIndex(1, -1)
	tree.RemoveRange(72, 75)
	log.Println(tree.Index(29))
	log.Println(tree.RemoveIndex(29))
	tree.check()
	log.Println(tree.view())
}

// func TestCompareOther(t *testing.T) {

// 	r := random.New(t.Name())

// 	for n := 0; n < 500; n++ {
// 		tree := New[int, int](compare.AnyEx[int])
// 		tree2 := avl.New(compare.Any[int])
// 		var removelist []int
// 		count := r.Intn(64)
// 		for i := 0; i < count; i++ {
// 			v := r.Intn(64)
// 			if tree.Put(v, v) {
// 				removelist = append(removelist, v)
// 			}
// 			tree2.Put(v, v)

// 			r1 := tree.view()
// 			r2 := tree2.View()
// 			// tree.check()
// 			if r1 != r2 {
// 				panic("Put")
// 			}

// 			if tree.size != uint(tree2.Size()) {
// 				panic("Size")
// 			}
// 		}

// 		for _, v := range removelist {

// 			_, ok := tree2.Remove(v)
// 			_, ok2 := tree.Remove(v)
// 			if ok != ok2 {
// 				panic("ok")
// 			}

// 			// r1 := tree.View()
// 			// r2 := tree2.View()
// 			// // tree.check()
// 			// if r1 != r2 {
// 			// 	panic("Remove")
// 			// }

// 			if tree.size != uint(tree2.Size()) {
// 				panic("Remove Size")
// 			}
// 		}
// 	}

// }
