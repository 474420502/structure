package treelist

import (
	"bytes"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestSeekRand(t *testing.T) {

	seed := time.Now().UnixNano()
	log.Println(seed)
	rand.Seed(seed)
	for n := 0; n < 2000; n++ {

		tree := New()
		tree.compare = compare.BytesLen
		var plist []int
		for i := 0; i < 200; i += rand.Intn(4) + 1 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			plist = append(plist, i)
		}

		plen := len(plist)
		i := rand.Intn(plen)
		m := plist[i]
		iter := tree.Iterator()
		mid := []byte(strconv.Itoa(m))
		iter.Seek(mid)
		if bytes.Compare(iter.Key(), mid) != 0 {
			t.Error("seek error")
		}
		if i > 0 {
			v := plist[i-1]
			iter.Prev()
			if bytes.Compare(iter.Key(), []byte(strconv.Itoa(v))) != 0 {
				t.Error("seek error", string(iter.Key()), plist, v)
			}
		} else {
			v := plist[i]
			iter.Prev()
			if iter.Valid() {
				t.Error(v)
			}
		}

		iter.Seek(mid)
		if i < plen-1 {
			v := plist[i+1]
			iter.Next()
			if bytes.Compare(iter.Key(), []byte(strconv.Itoa(v))) != 0 {
				t.Error("seek error", string(iter.Key()), plist, v)
			}

			// log.Println(v - 1)
			p := []byte(strconv.Itoa(v - 1))
			iter.SeekForPrev(p)
			if iter.Valid() {
				if bytes.Compare(iter.Key(), []byte(strconv.Itoa(m))) != 0 {
					log.Panicln("seek error key:", string(iter.Key()), plist, "mid:", m, string(p))
				}
			}

		} else {
			v := plist[i]
			iter.Next()
			if iter.Valid() {
				t.Error(v)
			}
		}
	}

}

func TestSeek(t *testing.T) {

	tree := New()
	for _, v := range testutils.TestedBytesWords {
		tree.Put(v, v)
		// log.Println(tree.debugString(true))
	}

	iter := tree.Iterator()
	iter.Seek([]byte("wor"))
	log.Println(string(iter.cur.Key))
	for iter.Valid() {
		v := string(iter.Key())
		if !strings.HasPrefix(v, "wor") {
			t.Error(v)
		}
		iter.Next()
	}

	tree.Clear()
	for _, v := range testutils.TestedBytes {
		tree.Put(v, v)
		// log.Println(tree.debugString(true))
	}

	//
	// │       ┌── 99(8|1)
	// │   ┌── 8(5|3)
	// │   │   └── 50(8|1)
	// └── 5(|8)
	// 	   │   ┌── 14(11|1)
	// 	   └── 11(5|4)
	// 		   └── 10(11|2)
	// 			   └── 1(10|1)
	//

	var correctResult = []string{"1", "10", "11", "14"}
	var result []string
	iter = tree.Iterator()
	iter.SeekForPrev([]byte("1"))
	for iter.Valid() {
		v := string(iter.Key())
		if strings.HasPrefix(v, "1") {
			result = append(result, v)
		} else {
			break
		}
		iter.Next()
	}
	for i, v := range correctResult {
		if result[i] != v {
			t.Error("seek error")
		}
	}
}

func TestFirstLast(t *testing.T) {
	tree := New()
	for _, v := range testutils.TestedBytes {
		tree.Put(v, v)
		log.Println(string(tree.root.Direct[0].Key), string(tree.root.Direct[1].Key))
	}

	iter := tree.Iterator()
	iter.SeekToLast()
	for iter.Valid() {
		iter.Value()
		iter.Prev()
	}

	iter.SeekToFirst()
	for iter.Valid() {
		iter.Value()
		iter.Next()
	}
}
