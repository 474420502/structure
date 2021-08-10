package treequeue

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/474420502/structure/compare"
)

func TestCase1(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(1628580888818576051)

	q := New(compare.Int)

	for i := 0; i < 20; i++ {
		v := rand.Intn(10)
		q.Put(v, i)
		log.Println(q.debugStringWithValue())
	}

	log.Println(q.debugStringWithValue())
	log.Println(q.Gets(2))
	q.Remove(2)
	q.check()
	log.Println(q.debugStringWithValue())
	log.Println(q.Gets(2))
}

func TestForce(t *testing.T) {
	seed := time.Now().UnixNano()
	log.Println(t.Name(), seed)
	rand.Seed(seed)

	q := New(compare.Int)

	for i := 0; i < 20; i++ {
		v := rand.Intn(10)
		q.Put(v, i)
	}

	log.Println(q.debugString(true))
	q.Extract(3, 5)
	log.Println(q.debugString(true))
}
