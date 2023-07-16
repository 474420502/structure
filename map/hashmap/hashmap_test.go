package hashmap

import (
	"log"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/set/treeset"
)

func TestForce(t *testing.T) {

	rand := random.New()
	hm := New()
	set := treeset.New[int, int](compare.AnyEx[int])

	for n := 0; n < 2000; n++ {

		if !set.Empty() {
			panic("")
		}

		if !hm.Empty() {
			panic("")
		}

		for i := 0; i < 200; i++ {
			v := rand.Intn(100)
			hm.Set(v, v)
			set.Add(v, v)
		}

		for _, v := range set.Values() {
			if _, ok := hm.Get(v); !ok {
				log.Panicln(set.String())
			}
		}

		for _, k := range hm.Keys() {
			if ok := set.Contains(k.(int)); !ok {
				panic("")
			}
		}

		for _, v := range hm.Values() {
			if ok := set.Contains(v.(int)); !ok {
				panic("")
			}
		}

		for i := 0; hm.Size() != 0 && i < 120; i++ {
			k := rand.Intn(100)
			if rand.OneOf64n(2) {
				hm.Remove(k)
				set.Remove(k)
			}

			if rand.OneOf64n(3) {
				hm.Set(k, k)
				set.Add(k, k)
			}
		}

		for _, k := range hm.Keys() {
			if ok := set.Contains(k.(int)); !ok {
				panic("")
			}
		}

		for _, v := range hm.Values() {
			if ok := set.Contains(v.(int)); !ok {
				panic("")
			}
		}

		hm.Clear()
		set.Clear()
	}

}

func TestCoverPut(t *testing.T) {

	rand := random.New()
	hm := New()

	for n := 0; n < 2000; n++ {

		for i := 0; i < 100; i++ {
			v := rand.Intn(200)
			hm.Put(v, v) // 不覆盖 所以kv永远相等
		}

		for _, s := range hm.Slices() {
			if s.Key != s.Value {
				panic("")
			}
		}

	}

}
