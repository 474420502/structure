package avlex

import (
	"log"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/avl"
)

func TestCase(t *testing.T) {

	r := random.New(1683745535466910261)

	tree := NewTree[int, int](compare.AnyEx[int])
	count := 20
	for i := 0; i < count; i++ {
		v := r.Intn(100)
		tree.Put(v, v)
	}
	tree.check()

	log.Println(tree.view())
}

func TestCasePut(t *testing.T) {

	r := random.New()

	for n := 0; n < 500000; n++ {
		tree := NewTree[int, int](compare.AnyEx[int])
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

	r := random.New(t.Name())

	for nn := 0; nn < 100000; nn++ {

		tree := NewTree[int, int](compare.AnyEx[int])
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

func TestCompareOther(t *testing.T) {

	r := random.New(t.Name())

	for n := 0; n < 50000; n++ {
		tree := NewTree[int, int](compare.AnyEx[int])
		tree2 := avl.New(compare.Any[int])
		var removelist []int
		count := r.Intn(64)
		for i := 0; i < count; i++ {
			v := r.Intn(64)
			if tree.Put(v, v) {
				removelist = append(removelist, v)
			}
			tree2.Put(v, v)

			r1 := tree.view()
			r2 := tree2.View()
			// tree.check()
			if r1 != r2 {
				panic("Put")
			}

			if tree.Size != uint(tree2.Size()) {
				panic("Size")
			}
		}

		for _, v := range removelist {

			_, ok := tree2.Remove(v)
			_, ok2 := tree.Remove(v)
			if ok != ok2 {
				panic("ok")
			}

			// r1 := tree.View()
			// r2 := tree2.View()
			// // tree.check()
			// if r1 != r2 {
			// 	panic("Remove")
			// }

			if tree.Size != uint(tree2.Size()) {
				panic("Remove Size")
			}
		}
	}

}

func BenchmarkPut(b *testing.B) {
	rand := random.New(1683721792150515321)

	b.StopTimer()
	tree := NewTree[int, int](compare.AnyEx[int])
	for i := 0; i < 10000; i++ {
		v := rand.Int()
		tree.Put(v, v)
		// tree.check()
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v := rand.Int()
		tree.Put(v, v)
		// tree.check()
	}

}

func BenchmarkRemove(b *testing.B) {
	rand := random.New(1683721792150515321)
	tree := NewTree[int, int](compare.AnyEx[int])

	var removelist []int
	var ri = 0
	for i := 0; i < b.N; i++ {

		if tree.Size == 0 {
			b.StopTimer()
			removelist = nil
			ri = 0
			for i := 0; i < 100; i++ {
				v := rand.Intn(100)
				if tree.Put(v, v) {
					removelist = append(removelist, v)
				}

				if i%25 == 0 {
					removelist = append(removelist, rand.Intn(100))
				}
				// tree.check()
			}
			b.StartTimer()
		}

		v := removelist[ri]
		tree.Remove(v)
		ri += 1

	}

}
