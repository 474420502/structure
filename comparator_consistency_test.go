package utils

import (
	"reflect"
	"testing"

	"github.com/474420502/structure/compare"
	arraylist "github.com/474420502/structure/list/array_list"
	priority "github.com/474420502/structure/queue/priority"
	"github.com/474420502/structure/set/treeset"
	"github.com/474420502/structure/tree/avl"
	"github.com/474420502/structure/tree/avls"
	"github.com/474420502/structure/tree/btree"
	"github.com/474420502/structure/tree/heap"
	"github.com/474420502/structure/tree/indextree"
	"github.com/474420502/structure/tree/rbtree"
	"github.com/474420502/structure/tree/skiplist"
	"github.com/474420502/structure/tree/treelist"
)

var consistencyKeys = []int{3, 1, 4, 2, 5}

func assertAscending(t *testing.T, name string, got []int) {
	t.Helper()
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s: comparator order mismatch: got %v want %v", name, got, want)
	}
}

// TestStandardComparatorAcrossContainers proves that one standard comparator
// drives the same ascending order in every ordered container, including the
// former index-style trees (avl, avls, treeset, priority).
func TestStandardComparatorAcrossContainers(t *testing.T) {
	t.Run("avl", func(t *testing.T) {
		tree := avl.New[int, int](compare.Any[int])
		for _, k := range consistencyKeys {
			tree.Put(k, k)
		}
		var got []int
		tree.Traverse(func(k, v int) bool { got = append(got, k); return true })
		assertAscending(t, "avl", got)
		if _, ok := tree.Delete(3); !ok || tree.Len() != 4 {
			t.Fatal("avl Delete/Len mismatch")
		}
	})

	t.Run("avls", func(t *testing.T) {
		tree := avls.New[int, int](compare.Any[int])
		for _, k := range consistencyKeys {
			tree.Put(k, k)
		}
		var got []int
		tree.Traverse(func(k, v int) bool { got = append(got, k); return true })
		assertAscending(t, "avls", got)
	})

	t.Run("treeset", func(t *testing.T) {
		set := treeset.New[int, int](compare.Any[int])
		for _, k := range consistencyKeys {
			set.Set(k, k)
		}
		var got []int
		set.Traverse(func(k, v int) bool { got = append(got, k); return true })
		assertAscending(t, "treeset", got)
	})

	t.Run("priority", func(t *testing.T) {
		queue := priority.New[int, int](compare.Any[int])
		for _, k := range consistencyKeys {
			queue.Put(k, k)
		}
		var got []int
		queue.Traverse(func(k, v int) bool { got = append(got, k); return true })
		assertAscending(t, "priority", got)
	})

	t.Run("indextree", func(t *testing.T) {
		tree := indextree.New[int](compare.Any[int])
		for _, k := range consistencyKeys {
			tree.Put(k, k)
		}
		var got []int
		tree.Traverse(func(k, v int) bool { got = append(got, k); return true })
		assertAscending(t, "indextree", got)
	})

	t.Run("btree", func(t *testing.T) {
		tree := btree.New[int, int](compare.Any[int])
		for _, k := range consistencyKeys {
			tree.Put(k, k)
		}
		var got []int
		tree.Traverse(func(k, v int) bool { got = append(got, k); return true })
		assertAscending(t, "btree", got)
	})

	t.Run("rbtree", func(t *testing.T) {
		tree := rbtree.New[int, int](compare.Any[int])
		for _, k := range consistencyKeys {
			tree.Put(k, k)
		}
		var got []int
		tree.Traverse(func(k, v int) bool { got = append(got, k); return true })
		assertAscending(t, "rbtree", got)
	})

	t.Run("skiplist", func(t *testing.T) {
		list := skiplist.New[int, int](compare.Any[int])
		for _, k := range consistencyKeys {
			list.Put(k, k)
		}
		var got []int
		list.Traverse(func(s *skiplist.Slice[int, int]) bool { got = append(got, s.Key); return true })
		assertAscending(t, "skiplist", got)
	})

	t.Run("treelist", func(t *testing.T) {
		tree := treelist.New[int, int](compare.Any[int])
		for _, k := range consistencyKeys {
			tree.Put(k, k)
		}
		var got []int
		tree.Traverse(func(s *treelist.Slice[int, int]) bool { got = append(got, s.Key); return true })
		assertAscending(t, "treelist", got)
	})

	t.Run("heap", func(t *testing.T) {
		h := heap.New[int](compare.Any[int])
		for _, k := range consistencyKeys {
			h.Put(k)
		}
		var got []int
		for {
			v, ok := h.Pop()
			if !ok {
				break
			}
			got = append(got, v)
		}
		assertAscending(t, "heap", got)
	})

	t.Run("list-contains", func(t *testing.T) {
		l := arraylist.New[int](compare.Any[int])
		for _, k := range consistencyKeys {
			l.PushBack(k)
		}
		if count := l.Contains(3, 5); count != 2 {
			t.Fatalf("array list Contains count = %d, want 2", count)
		}
	})
}
