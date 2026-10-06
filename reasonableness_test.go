package utils

import (
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/474420502/structure/compare"
	priority "github.com/474420502/structure/queue/priority"
	"github.com/474420502/structure/set/treeset"
	"github.com/474420502/structure/tree/avl"
	"github.com/474420502/structure/tree/avls"
	"github.com/474420502/structure/tree/indextree"
	"github.com/474420502/structure/tree/rbtree"
	"github.com/474420502/structure/tree/skiplist"
	"github.com/474420502/structure/tree/treelist"
)

func isSortedInts(values []int) bool {
	for i := 1; i < len(values); i++ {
		if values[i-1] > values[i] {
			return false
		}
	}
	return true
}

// TestMonotoneHeightReasonableness checks that sorted insertion stays within a
// sane height bound and never degenerates into a list.
func TestMonotoneHeightReasonableness(t *testing.T) {
	const n = 50000
	maxHeight := int(4*math.Log2(float64(n+1))) + 4

	avlTree := avl.New[int, int](compare.Any[int])
	avlsTree := avls.New[int, int](compare.Any[int])
	rbTree := rbtree.New[int, int](compare.Any[int])
	indexTree := indextree.New[int](compare.Any[int])

	for i := 0; i < n; i++ {
		avlTree.Put(i, i)
		avlsTree.Put(i, i)
		rbTree.Put(i, i)
		indexTree.Put(i, i)
	}

	if h := int(avlTree.Height()); h <= 0 || h > maxHeight {
		t.Fatalf("avl height %d out of range (0,%d]", h, maxHeight)
	}
	if h := int(avlsTree.Height()); h <= 0 || h > maxHeight {
		t.Fatalf("avls height %d out of range (0,%d]", h, maxHeight)
	}
	if h := rbTree.Height(); h <= 0 || h > maxHeight {
		t.Fatalf("rbtree height %d out of range (0,%d]", h, maxHeight)
	}
	if h := indexTree.BenchmarkStats().Height; h <= 0 || h > maxHeight {
		t.Fatalf("indextree height %d out of range (0,%d]", h, maxHeight)
	}

	for _, tree := range []struct {
		name string
		len  func() int
	}{
		{"avl", avlTree.Len},
		{"avls", avlsTree.Len},
		{"rbtree", rbTree.Len},
		{"indextree", indexTree.Len},
	} {
		if got := tree.len(); got != n {
			t.Fatalf("%s Len = %d want %d", tree.name, got, n)
		}
	}
}

type intMapOps struct {
	name   string
	set    func(k, v int)
	del    func(k int) bool
	length func() int
	keys   func() []int
}

func runIntMapWorkload(t *testing.T, ops intMapOps) {
	t.Run(ops.name, func(t *testing.T) {
		random := rand.New(rand.NewSource(4242))
		ref := map[int]int{}

		for step := 0; step < 4000; step++ {
			key := random.Intn(256)
			if random.Intn(4) == 0 {
				ok := ops.del(key)
				_, exists := ref[key]
				if ok != exists {
					t.Fatalf("step %d: delete(%d) ok=%v exists=%v", step, key, ok, exists)
				}
				delete(ref, key)
			} else {
				value := random.Intn(100000)
				ops.set(key, value)
				ref[key] = value
			}
			if step%500 == 0 {
				checkIntMapState(t, ops, ref)
			}
		}
		checkIntMapState(t, ops, ref)
	})
}

func checkIntMapState(t *testing.T, ops intMapOps, ref map[int]int) {
	t.Helper()
	if got := ops.length(); got != len(ref) {
		t.Fatalf("%s: len = %d want %d", ops.name, got, len(ref))
	}
	want := make([]int, 0, len(ref))
	for k := range ref {
		want = append(want, k)
	}
	sort.Ints(want)

	got := ops.keys()
	if !isSortedInts(got) {
		t.Fatalf("%s: keys are not sorted", ops.name)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s: keys = %v want %v", ops.name, got, want)
	}
}

// TestRandomWorkloadReasonableness drives every ordered map through the same
// random insert/delete workload and checks size and ordering against a Go map.
func TestRandomWorkloadReasonableness(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		avlTree := avl.New[int, int](compare.Any[int])
		runIntMapWorkload(t, intMapOps{
			name:   "avl",
			set:    func(k, v int) { avlTree.Set(k, v) },
			del:    func(k int) bool { _, ok := avlTree.Delete(k); return ok },
			length: avlTree.Len,
			keys: func() []int {
				var result []int
				avlTree.Traverse(func(k, v int) bool { result = append(result, k); return true })
				return result
			},
		})

		avlsTree := avls.New[int, int](compare.Any[int])
		runIntMapWorkload(t, intMapOps{
			name:   "avls",
			set:    func(k, v int) { avlsTree.Set(k, v) },
			del:    func(k int) bool { _, ok := avlsTree.Delete(k); return ok },
			length: avlsTree.Len,
			keys: func() []int {
				var result []int
				avlsTree.Traverse(func(k, v int) bool { result = append(result, k); return true })
				return result
			},
		})

		set := treeset.New[int, int](compare.Any[int])
		runIntMapWorkload(t, intMapOps{
			name:   "treeset",
			set:    func(k, v int) { set.Set(k, v) },
			del:    func(k int) bool { _, ok := set.Delete(k); return ok },
			length: set.Len,
			keys: func() []int {
				var result []int
				set.Traverse(func(k, v int) bool { result = append(result, k); return true })
				return result
			},
		})

		indexTree := indextree.New[int](compare.Any[int])
		runIntMapWorkload(t, intMapOps{
			name:   "indextree",
			set:    func(k, v int) { indexTree.Set(k, v) },
			del:    func(k int) bool { _, ok := indexTree.Delete(k); return ok },
			length: indexTree.Len,
			keys: func() []int {
				var result []int
				indexTree.Traverse(func(k, v int) bool { result = append(result, k); return true })
				return result
			},
		})

		rbTree := rbtree.New[int, int](compare.Any[int])
		runIntMapWorkload(t, intMapOps{
			name:   "rbtree",
			set:    func(k, v int) { rbTree.Set(k, v) },
			del:    func(k int) bool { _, ok := rbTree.Delete(k); return ok },
			length: rbTree.Len,
			keys: func() []int {
				var result []int
				rbTree.Traverse(func(k, v int) bool { result = append(result, k); return true })
				return result
			},
		})

		skip := skiplist.New[int, int](compare.Any[int])
		runIntMapWorkload(t, intMapOps{
			name:   "skiplist",
			set:    func(k, v int) { skip.Set(k, v) },
			del:    func(k int) bool { _, ok := skip.Delete(k); return ok },
			length: skip.Len,
			keys: func() []int {
				var result []int
				skip.Traverse(func(s *skiplist.Slice[int, int]) bool { result = append(result, s.Key); return true })
				return result
			},
		})

		treeList := treelist.New[int, int](compare.Any[int])
		runIntMapWorkload(t, intMapOps{
			name:   "treelist",
			set:    func(k, v int) { treeList.Set(k, v) },
			del:    func(k int) bool { _, ok := treeList.Delete(k); return ok },
			length: treeList.Len,
			keys: func() []int {
				var result []int
				treeList.Traverse(func(s *treelist.Slice[int, int]) bool { result = append(result, s.Key); return true })
				return result
			},
		})
	})
}

// TestDuplicateKeyReasonableness checks the multiset containers keep every
// duplicate and remove exactly one per Delete call.
func TestDuplicateKeyReasonableness(t *testing.T) {
	t.Run("avls", func(t *testing.T) {
		tree := avls.New[int, int](compare.Any[int])
		for i := 0; i < 5; i++ {
			tree.Put(7, i)
		}
		if tree.Len() != 5 {
			t.Fatalf("avls Len = %d want 5", tree.Len())
		}
		seen := 0
		tree.Traverse(func(k, v int) bool {
			if k == 7 {
				seen++
			}
			return true
		})
		if seen != 5 {
			t.Fatalf("avls duplicate run = %d want 5", seen)
		}
		if _, ok := tree.Delete(7); !ok || tree.Len() != 4 {
			t.Fatalf("avls Delete left Len = %d want 4", tree.Len())
		}
	})

	t.Run("priority", func(t *testing.T) {
		queue := priority.New[int, int](compare.Any[int])
		for i := 0; i < 5; i++ {
			queue.Put(7, i)
		}
		if queue.Len() != 5 {
			t.Fatalf("priority Len = %d want 5", queue.Len())
		}
		if got := len(queue.Gets(7)); got != 5 {
			t.Fatalf("priority Gets(7) = %d want 5", got)
		}
		if _, ok := queue.Delete(7); !ok || queue.Len() != 4 {
			t.Fatalf("priority Delete left Len = %d want 4", queue.Len())
		}
	})
}
