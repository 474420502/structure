package avl

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestNextPrev(t *testing.T) {
	tree := New(compare.Int)
	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	// log.Println(tree.String())

	iter := tree.Iterator()
	iter.SeekToFirst()
	if !iter.Vaild() {
		panic("")
	}

	if iter.Value() != 0 {
		panic("")
	}

	for i := 0; i < 10; i++ {
		if iter.Value() != i {
			panic("")
		}
		iter.Next()
	}

	iter.Prev()
	if !iter.Vaild() && iter.Value() != 9 {
		panic("")
	}

	for i := 9; i >= 0; i-- {
		if iter.Value() != i {
			panic("")
		}
		iter.Prev()
	}

	iter.Seek(5)
	for i := 5; i >= 0; i-- {
		if iter.Value() != i {
			panic("")
		}
		iter.Prev()
	}

	iter.Seek(5)
	for i := 5; i < 10; i++ {
		if iter.Value() != i {
			panic("")
		}
		iter.Next()
	}
}

func TestSeekFor(t *testing.T) {
	tree := New(compare.Int)
	for i := 0; i < 20; i += 2 {
		tree.Put(i, i)
	}

	// log.Println(tree.String())

	iter := tree.Iterator()
	iter.SeekForPrev(7) // Key == 6
	for i := 6; i >= 0; i -= 2 {
		if iter.Value() != i {
			panic("")
		}
		iter.Prev()
	}

	iter.SeekForPrev(7)
	for i := 6; i < 20; i += 2 {
		if iter.Value() != i {
			panic("")
		}
		iter.Next()
	}

	iter.SeekForPrev(6)
	for i := 6; i < 20; i += 2 {
		if iter.Value() != i {
			panic("")
		}
		iter.Next()
	}

	iter.SeekForNext(7)
	for i := 8; i < 20; i += 2 {
		if iter.Value() != i {
			panic("")
		}
		iter.Next()
	}
}
