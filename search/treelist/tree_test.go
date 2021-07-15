package treelist

import (
	"bytes"
	"log"
	"strconv"
	"testing"

	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestIndex(t *testing.T) {
	tree := New()
	for i := 0; i < 100; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)

	}

}

func TestRank(t *testing.T) {
	tree := New()
	tree.compare = compare.BytesLen

	for i := 0; i < 100; i++ {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
		// log.Println(tree.debugString(true))
	}

	// log.Println(tree.debugString(true))

	for i := 0; i < 100; i++ {
		k := []byte(strconv.Itoa(i))
		if v := tree.IndexOf(k); v != int64(i) {
			t.Error("index error", i, "rank", v)
		}
	}

}

func TestRemove1(t *testing.T) {
	tree := New()
	for _, i := range testutils.TestedArray {
		v := []byte(strconv.Itoa(i))
		if !tree.Put(v, v) {
			log.Println("equal key", i)
		}
	}

	for _, v := range tree.Slices() {
		r := tree.Remove(v.Key)
		if bytes.Compare(r.Key, v.Key) != 0 {
			log.Println("remove error check it", string(r.Key), string(v.Key))
		}
	}

	if tree.Size() != 0 {
		t.Error(tree.Slices())
	}
}

func TestRemove2(t *testing.T) {
	tree := New()
	for _, i := range testutils.TestedBigArray {
		v := []byte(strconv.Itoa(i))
		if !tree.Put(v, v) {
			log.Println("equal key", i)
		}
	}

	if tree.Size() != int64(len(testutils.TestedBigArray)-4) {
		t.Error(tree.Size(), tree.Slices())
	}

	for _, v := range tree.Slices() {
		tree.Remove(v.Key)
	}

	if tree.Size() != 0 {
		t.Error(tree.Slices())
	}
}

func TestRange(t *testing.T) {
	tree := New()
	for i := 0; i < 100; i += 2 {
		v := []byte(strconv.Itoa(i))
		tree.Put(v, v)
	}

	start := []byte(strconv.Itoa(0)) // 41 63
	end := []byte(strconv.Itoa(63))
	log.Println(tree.debugString(false))
	tree.RemoveRange(start, end)
	log.Println(tree.debugString(true))

}
