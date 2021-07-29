package arraytree

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/474420502/structure/compare"
	indextree "github.com/474420502/structure/tree/itree"
)

func TestCase(t *testing.T) {
	log.Printf("%d,%d,%d", 0b0101010, 0b0101010, 0b0111011)
	log.Printf("%d,%d,%d", 0b1110000, 0b1101000, 0b1010001)

	tree := New()
	tl := indextree.New(compare.Int)
	for i := 0; i < 1<<5-1; i++ {
		tl.Put(i, i)
		log.Println(tl.String())
	}

	for i := 0; i < 1000; i++ {
		tree.Put(i, i)
		log.Println(tree.debugString(true))
	}
}

func TestCopy(t *testing.T) {

	var data []interface{} = make([]interface{}, 1000050)
	for i := 0; i < 1000050; i++ {
		data[i] = i
	}

	now := time.Now()
	a := data[0:1000]
	for i := 0; i < 100; i++ {
		n := rand.Intn(50)
		for x := 0; x < 1000; x++ {
			copy(data[n:1000000+n], a)
		}

	}
	r := time.Since(now).Nanoseconds() / 100
	t.Error(r)
}
