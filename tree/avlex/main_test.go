package avlex

import (
	"log"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
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

	log.Println(tree.View())
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

	for nn := 0; nn < 10000; nn++ {

		tree := NewTree[int, int](compare.AnyEx[int])
		count := 100

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
			if result, ok := tree.Remove(v); ok {
				if isDebug {
					log.Print(beforeView, "remove value: ", result, " key: ", v, "", tree.View(), "\n")
				}
			}
			tree.check()
		}
	}
	// log.Println(tree.View())
}

func Benchmark(b *testing.B) {

}

func BenchmarkMark(b *testing.B) {
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
