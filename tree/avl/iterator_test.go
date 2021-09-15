package avl

import (
	"log"
	"sort"
	"testing"

	random "github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

func TestNextPrev(t *testing.T) {
	tree := New(compare.Int)
	for i := 0; i < 10; i++ {
		tree.Cover(i, i)
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
		tree.Cover(i, i)
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

func TestIteratorForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree := New(compare.Int)
		var priority []int
		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			if tree.Put(v, v) {
				priority = append(priority, v)
			}

		}
		sort.Slice(priority, func(i, j int) bool {
			return priority[i] < priority[j]
		})

		s := rand.Intn(100)

		// log.Println(priority, idx, len(priority))

		iter := tree.Iterator()
		iter.SeekForNext(s)

		idx := sort.Search(len(priority), func(i int) bool {
			return priority[i] >= s
		})

		if idx == len(priority) {
			if iter.Vaild() {
				log.Panicln(iter.Value(), priority)
			}
		} else {
			for i := idx; i < tree.Size(); i++ {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Next()
			}
			iter.SeekForNext(s)
			for i := idx; i >= 0; i-- {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Prev()
			}
		}

		iter.SeekForPrev(s)
		idx = sort.Search(len(priority), func(i int) bool {
			return priority[i] > s
		})

		if idx-1 < 0 {
			if iter.Vaild() {
				log.Panicln(iter.Value(), priority, idx-1, s)
			}
		} else {
			for i := idx - 1; i < tree.Size(); i++ {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Next()
			}
			iter.SeekForPrev(s)
			for i := idx - 1; i >= 0; i-- {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Prev()
			}
		}

		iter.SeekToFirst()
		for i := 0; i < len(priority); i++ {
			if iter.Value() != priority[i] {
				log.Panic("")
			}
			iter.Next()
		}

		iter.SeekToLast()
		for i := len(priority) - 1; i >= 0; i-- {
			if iter.Value() != priority[i] {
				log.Panic("")
			}
			iter.Prev()
		}

	}
}
