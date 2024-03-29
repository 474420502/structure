package treelist

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/474420502/random"
	testutils "github.com/474420502/structure/tree/test_utils"
)

func TestIteratorIndexForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree := New()
		tree.compare = compareBytesLen
		var plist []int
		for i := 0; i < 200; i += rand.Intn(8) + 2 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			plist = append(plist, i)
		}
		plen := len(plist)
		i := rand.Intn(plen)

		iter := tree.Iterator()
		iter.SeekGE([]byte(strconv.Itoa(plist[i])))
		if iter.Index() != int64(i) {
			t.Error()
		}

		iter.SeekToFirst()
		for iter.Valid() {

			if tree.compare([]byte(strconv.Itoa(plist[i])), iter.Key()) == 0 {
				break
			}
			iter.Next()
		}
		if iter.Index() != int64(i) {
			log.Panic()
		}

		iter.SeekToLast()
		for iter.Valid() {
			if tree.compare([]byte(strconv.Itoa(plist[i])), iter.Key()) == 0 {
				break
			}
			iter.Prev()
		}
		if iter.Index() != int64(i) {
			log.Panic()
		}
	}
}

func TestSeekRand(t *testing.T) {

	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		tree := New()
		tree.compare = compareBytesLen
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
		iter.SeekGE(mid)
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

		iter.SeekGE(mid)
		if i < plen-1 {
			v := plist[i+1]
			iter.Next()
			if bytes.Compare(iter.Key(), []byte(strconv.Itoa(v))) != 0 {
				t.Error("seek error", string(iter.Key()), plist, v)
			}

			// log.Println(v - 1)
			p := []byte(strconv.Itoa(v - 1))
			iter.SeekLE(p)
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

func TestSeekByIndexForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {

		tree := New()
		tree.compare = compareBytesLen
		var plist []int
		for i := 0; i < 200; i += rand.Intn(8) + 4 {
			v := []byte(strconv.Itoa(i))
			tree.Put(v, v)
			plist = append(plist, i)
		}

		sort.Ints(plist)
		for i := 0; i < 5; i++ {
			iter := tree.Iterator()
			idx := int64(rand.Intn(len(plist)))
			iter.SeekByIndex(idx)
			if !iter.Valid() {
				t.Error()
				panic(idx)
			}
			if tree.compare([]byte(strconv.Itoa(plist[idx])), iter.Key()) != 0 {
				log.Panicln(idx, strconv.Itoa(plist[idx]), iter.Key())
			}

			citer := iter.Clone()
			for x := idx + 1; x < tree.Size(); x++ {
				iter.Next()
				if tree.compare([]byte(strconv.Itoa(plist[x])), iter.Key()) != 0 {
					log.Panicln(x, strconv.Itoa(plist[x]), iter.Key())
				}
			}

			for x := idx - 1; x >= 0; x-- {
				citer.Prev()
				if tree.compare([]byte(strconv.Itoa(plist[x])), citer.Key()) != 0 {
					log.Panicln(x, strconv.Itoa(plist[x]), citer.Key())
				}
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
	iter.SeekGE([]byte("wor"))
	log.Println(string(iter.cur.Key))
	var checkresult []string
	for iter.Valid() {
		v := string(iter.Key())
		checkresult = append(checkresult, v)
		if !strings.HasPrefix(v, "wor") {
			t.Error(v)
		}
		iter.Next()
	}

	if len(checkresult) != 2 {
		t.Error("checkresult len != 2", checkresult)
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
	iter.SeekLE([]byte("1"))
	for iter.Valid() {
		v := string(iter.Key())
		if strings.HasPrefix(v, "1") {
			result = append(result, v)
		} else {
			break
		}
		iter.Next()
	}

	iter.SeekLE([]byte("0"))
	if iter.Valid() {
		panic("0 not has prev")
	}

	iter.SeekGE([]byte("12"))
	if iter.Valid() {
		if string(iter.Key()) != "14" {
			panic("SeekGE error")
		}
	}

	iter.SeekGE([]byte("999"))
	if iter.Valid() {
		panic("999 not has next")
	}

	for i, v := range correctResult {
		if result[i] != v {
			t.Error("seek error")
		}
	}
}

func TestSeekRange(t *testing.T) {
	tree := New()
	for _, v := range testutils.TestedBytesSimlpe {
		tree.Put(v, v)
	}

	//	│       ┌── c6
	//	│   ┌── c4
	//	└── c1
	//		│   ┌── a5
	//		└── a3
	//			└── a1

	iter := tree.Iterator()
	if !iter.SeekGE([]byte("a3")) { // 由于a5存在
		t.Error("SeekLT errror")
	}

	var result []string
	for ; iter.Valid(); iter.Next() {
		k := string(iter.Key())
		result = append(result, k)
		if k == "c4" {
			break
		}
	}

	if fmt.Sprintf("%v", result) != "[a3 a5 c1 c4]" {
		t.Error()
	}

}

func TestSeekDirect(t *testing.T) {
	tree := New()
	for _, v := range testutils.TestedBytesSimlpe {
		tree.Put(v, v)
	}

	//	│       ┌── c6
	//	│   ┌── c4
	//	└── c1
	//		│   ┌── a5
	//		└── a3
	//			└── a1

	iter := tree.Iterator()
	if !iter.SeekLT([]byte("a5")) { // 由于a5存在
		t.Error("SeekLT errror")
	}
	if iter.Valid() {
		k := string(iter.Key())
		if k != "a3" {
			t.Error("SeekLT errror")
		}
	}

	if !iter.SeekGT([]byte("a5")) { // 由于a5存在
		t.Error("SeekGT errror")
	}
	if iter.Valid() {
		k := string(iter.Key())
		if k != "c1" {
			t.Error("SeekGT errror")
		}
	}

	if iter.SeekLT([]byte("a4")) { // 由于a5不存在
		t.Error("SeekLT errror")
	}
	if iter.Valid() {
		k := string(iter.Key())
		if k != "a3" {
			t.Error("SeekLT errror")
		}
	}

	if iter.SeekGT([]byte("a4")) { // 由于a5不存在
		t.Error("SeekGT errror")
	}
	if iter.Valid() {
		k := string(iter.Key())
		if k != "a5" {
			t.Error("SeekGT errror")
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

func TestIteratorSeekForForce(t *testing.T) {
	r := random.New()

	for n := 0; n < 2000; n++ {
		tree := New()
		var plist [][]byte
		for i := 0; i < 200; i++ {
			v := []byte(strconv.Itoa(r.Intn(200)))
			if tree.Put(v, v) {
				plist = append(plist, v)
			}
		}

		sort.Slice(plist, func(i, j int) bool {
			return tree.compare(plist[i], plist[j]) < 0
		})

		iter := tree.Iterator()
		iter.SeekToFirst()

		for i, key := range plist {
			if !iter.Valid() {
				t.Error("")
			}
			if iter.Index() != int64(i) {
				t.Error("")
			}
			if iter.Compare(key) != 0 {
				t.Error("")
			}
			iter.Next()
		}

		for i := 0; i < 5; i++ {
			idx := rand.Intn(len(plist)-2) + 1

			skey := plist[idx]

			nkey := plist[idx+1]
			pkey := plist[idx-1]

			tree.Remove(skey)

			iter := tree.Iterator()
			iter.SeekLE(skey)
			if string(iter.Key()) != string(pkey) {
				panic("")
			}

			iter.SeekGE(skey)
			if string(iter.Key()) != string(nkey) {
				panic("")
			}

			tree.Put(skey, skey)
		}

	}
}
