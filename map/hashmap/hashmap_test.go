package hashmap

import (
	"testing"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/random"
	"github.com/474420502/structure/set/treeset"
)

func TestForce(t *testing.T) {

	r := random.New()
	hm := New()
	set := treeset.New(compare.Int)

	for n := 0; n < 2000; n++ {
		for i := 0; i < 2000; i++ {
			v := r.Intn(100)
			hm.Put(v, v)
			set.Add(v)
		}

		set.Values()
	}

	// t.Error(hm.Get(4))
}
