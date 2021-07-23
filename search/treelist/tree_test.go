package treelist

import (
	"bytes"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/474420502/structure/compare"
	avl "github.com/474420502/structure/tree/avl"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func init() {
	log.SetFlags(log.Llongfile)
}

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
	// tree := New()
	// tree.compare = compare.BytesLen
	// for i := 0; i < 100; i += 4 {
	// 	v := []byte(strconv.Itoa(i))
	// 	tree.Put(v, v)
	// }
	// log.Println(tree.debugString(true))
	seed := time.Now().UnixNano()
	log.Println(seed)
	rand.Seed(seed)
	for n := 0; n < 100000; n++ {
		startkey := rand.Intn(100)
		endkey := rand.Intn(100)
		if startkey > endkey {
			temp := startkey
			startkey = endkey
			endkey = temp
		}
		tree := New()
		tree.compare = compare.BytesLen
		avltree := avl.New(compare.Int)
		for i := 0; i < 100; i += 4 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			avltree.Put(i, i)
		}
		// tree.rcount = 0
		start := []byte(strconv.Itoa(startkey)) // 41 63
		end := []byte(strconv.Itoa(endkey))

		// srctree := tree.debugString(true)

		// log.Println(srctree)
		// log.Println("start:", startkey, "end:", endkey)
		tree.RemoveRange(start, end)
		// tree.debugString(true)
		// log.Println("rcount", tree.rcount, tree.getHeight(), tree.Size())
		// log.Println(tree.debugString(true))

		for i := startkey; i <= endkey; i++ {
			avltree.Remove(i)
		}

		if tree.Size() != int64(avltree.Size()) {
			t.Error(avltree.Height(), avltree.Size(), avltree)
			t.Error(tree.Size(), tree.debugString(true))
			return
		}

		avltree.Traverse(func(k, v interface{}) bool {
			key := []byte(strconv.Itoa(v.(int)))
			if _, ok := tree.Get(key); !ok {
				t.Error("tree is error")
			}
			return true
		})

		iter := tree.Iterator()
		iter.SeekToFirst()
		for iter.Valid() {
			if _, ok := tree.Get(iter.Value()); !ok {
				log.Println("start:", startkey, "end:", endkey)
				// log.Println(srctree)
				log.Println(tree.debugString(true))
				log.Println("not ok", string(iter.Value()))
				panic("")
			}
			iter.Next()
		}
	}
}
