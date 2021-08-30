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
	set := treeset.New(compare.Int)

	for n := 0; n < 2000; n++ {

		if !set.Empty() {
			panic("")
		}

		if !hm.Empty() {
			panic("")
		}

		for i := 0; i < 200; i++ {
			v := rand.Intn(100)
			hm.Put(v, v)
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

		for _, v := range hm.Values() {
			if ok := set.Contains(v); !ok {
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
				hm.Put(k, k)
				set.Add(k)
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
	}

	// t.Error(hm.Get(4))
}