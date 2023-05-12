package avlbest

import (
	"log"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

func TestCase(t *testing.T) {

	// 	2023/05/11 03:10:37
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

	r := random.New(1683745535466910261)

	tree := NewTree[int, int](compare.AnyEx[int])
	count := 20
	for i := 0; i < count; i++ {
		v := r.Intn(100)
		tree.Put(v, v)
		log.Println(tree.View(), v)
		tree.check()
	}
	// tree.check()

	log.Println(tree.View())
}

// func TestCompareOther(t *testing.T) {

// 	r := random.New(t.Name())

// 	for n := 0; n < 10000; n++ {
// 		tree := NewTree[int, int](compare.AnyEx[int])
// 		tree2 := avl.New(compare.Any[int])
// 		count := r.Intn(128)
// 		for i := 0; i < count; i++ {
// 			v := r.Intn(128)
// 			tree.Put(v, v)
// 			tree2.Put(v, v)
// 			r1 := tree.View()
// 			r2 := tree2.View()
// 			// tree.check()
// 			if r1 != r2 {
// 				panic("")
// 			}
// 		}
// 	}

// }

func TestCaseR(t *testing.T) {
	log.SetFlags(log.Llongfile)

	r := random.New()

	for nn := 0; nn < 100000; nn++ {

		tree := NewTree[int, int](compare.AnyEx[int])
		count := 40

		var checkv []int

		for i := 0; i < count; i++ {
			v := r.Intn(100)
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
			var rv int = r.Intn(100)
			tree.Remove(rv)
		}

		// log.Println(tree.Remove(rv))
		// log.Println(tree.View())
		for _, v := range checkv {
			// if v == 38 {
			// 	log.Println("")
			// }

			beforeView := tree.View()
			if ok := tree.Remove(v); ok {
				if isDebug {
					log.Print(beforeView, " key: ", v, "", tree.View(), "\n")
				}
			}
			tree.check()
		}
	}
	// log.Println(tree.View())
}

func BenchmarkMark(b *testing.B) {
	rand := random.New(1683721792150515321)
	// rand := random.New()
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
