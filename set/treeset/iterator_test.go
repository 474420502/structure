package treeset

import (
	"log"
	"sort"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

func TestIteratorForce(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree := New(compare.Int)
		var priority []int
		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			if tree.Add(v) {
				priority = append(priority, v)
			}

		}
		sort.Slice(priority, func(i, j int) bool {
			return priority[i] < priority[j]
		})

		s := rand.Intn(100)

		// log.Println(priority, idx, len(priority))

		iter := tree.Iterator()
		iter.SeekGE(s)

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
			iter.SeekGE(s)
			for i := idx; i >= 0; i-- {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Prev()
			}
		}

		iter.SeekLE(s)
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
			iter.SeekLE(s)
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

func TestIteratorForce2(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 2000; n++ {
		tree := New(compare.Int)
		var priority []int
		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			if tree.Add(v) {
				priority = append(priority, v)
			}

		}
		sort.Slice(priority, func(i, j int) bool {
			return priority[i] < priority[j]
		})

		s := rand.Intn(100)

		// log.Println(priority, idx, len(priority))

		iter := tree.Iterator()
		iter.SeekGT(s)

		idx := sort.Search(len(priority), func(i int) bool {
			return priority[i] > s
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
			iter.SeekGT(s)
			for i := idx; i >= 0; i-- {
				if priority[i] != iter.Value() {
					panic("")
				}
				iter.Prev()
			}
		}

		iter.SeekLT(s)
		idx = sort.Search(len(priority), func(i int) bool {
			return priority[i] >= s
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
			iter.SeekLT(s)
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
