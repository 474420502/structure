package realbst

import (
	"log"
	"math/rand"
	"sort"
	"testing"

	"github.com/474420502/structure/compare"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestCase1(t *testing.T) {
	tree := New(compare.Int)
	var priority []int
	for i := 0; i < 100; i++ {
		v := rand.Intn(100)
		if tree.Put(v, v) {
			priority = append(priority, v)
		}

	}
	sort.Ints(priority)
	log.Println(tree.debugString())
	log.Println(priority, tree.Get(46), tree.Center.Key)
}
