package treelist

import (
	"log"
	"strings"
	"testing"

	testutils "github.com/474420502/structure/tree/test_utils"
)

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
		v := string(iter.Value())
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
	iter.SeekToPrev([]byte("1"))
	for iter.Valid() {
		v := string(iter.Value())
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
