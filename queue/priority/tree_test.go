package treequeue

import (
	"log"
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
)

func TestCase1(t *testing.T) {
	q := New(compare.Int)

	for i := 0; i < 10; i++ {
		v := rand.Intn(10)
		q.Put(v, i)
	}

	log.Println(q.debugString(true))
	q.Remove(0)
	q.Remove(7)
	log.Println(q.debugString(true))
}
