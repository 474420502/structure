package btree

import (
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/474420502/structure/compare"
)

func TestBTreeBasic(t *testing.T) {
	tree := New[int, string](compare.Any[int])
	if !tree.Empty() || tree.Len() != 0 || tree.Size() != 0 || tree.Height() != 0 {
		t.Fatal("new tree should be empty")
	}
	if _, ok := tree.Get(1); ok {
		t.Fatal("Get on empty tree should be false")
	}
	if _, ok := tree.Delete(1); ok {
		t.Fatal("Delete on empty tree should be false")
	}
	if _, _, ok := tree.Min(); ok {
		t.Fatal("Min on empty tree should be false")
	}
	if _, _, ok := tree.Max(); ok {
		t.Fatal("Max on empty tree should be false")
	}

	if !tree.Put(1, "a") {
		t.Fatal("Put should insert a missing key")
	}
	if tree.Put(1, "b") {
		t.Fatal("Put should preserve an existing value")
	}
	if v, _ := tree.Get(1); v != "a" {
		t.Fatalf("Put changed the value to %q", v)
	}
	if tree.Set(1, "b") {
		t.Fatal("Set should report false when overwriting")
	}
	if v, _ := tree.Get(1); v != "b" {
		t.Fatalf("Set value = %q want b", v)
	}
	if !tree.Upsert(1, "c") {
		t.Fatal("Upsert should report replacement")
	}
	if tree.Upsert(2, "d") {
		t.Fatal("Upsert should report false when inserting")
	}
	if tree.InsertIfAbsent(2, "x") {
		t.Fatal("InsertIfAbsent should reject an existing key")
	}

	if k, v, ok := tree.Min(); !ok || k != 1 || v != "c" {
		t.Fatalf("Min = %d,%q,%v", k, v, ok)
	}
	if k, v, ok := tree.Max(); !ok || k != 2 || v != "d" {
		t.Fatalf("Max = %d,%q,%v", k, v, ok)
	}

	if v, ok := tree.Delete(2); !ok || v != "d" {
		t.Fatalf("Delete(2) = %q,%v", v, ok)
	}
	if _, ok := tree.Delete(2); ok {
		t.Fatal("second Delete should be false")
	}
	if tree.Len() != 1 {
		t.Fatalf("Len = %d want 1", tree.Len())
	}

	tree.Clear()
	if !tree.Empty() || tree.Height() != 0 {
		t.Fatal("Clear should empty the tree")
	}
}

func TestBTreeStringKeys(t *testing.T) {
	tree := New[string, int](compare.ArrayAny[string])
	for _, k := range []string{"b", "a", "aa", "ab", "z"} {
		tree.Set(k, len(k))
	}
	var got []string
	tree.Traverse(func(k string, v int) bool { got = append(got, k); return true })
	want := []string{"a", "aa", "ab", "b", "z"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("order = %v want %v", got, want)
	}
}

func validateTree(t *testing.T, tree *Tree[int, int], expected map[int]int) {
	t.Helper()

	if tree.size != len(expected) {
		t.Fatalf("size = %d want %d", tree.size, len(expected))
	}
	if len(expected) == 0 {
		if len(tree.root.items) != 0 || len(tree.root.children) != 0 {
			t.Fatal("empty tree should have an empty root")
		}
		return
	}

	var keys []int
	tree.Traverse(func(k, v int) bool {
		if expected[k] != v {
			t.Fatalf("value mismatch for %d: got %d want %d", k, v, expected[k])
		}
		keys = append(keys, k)
		return true
	})
	want := make([]int, 0, len(expected))
	for k := range expected {
		want = append(want, k)
	}
	sort.Ints(want)
	if !reflect.DeepEqual(keys, want) {
		t.Fatalf("traversal mismatch: got %v want %v", keys, want)
	}

	leafDepth := -1
	count := 0
	var check func(n *node[int, int], depth int, low, high *int)
	check = func(n *node[int, int], depth int, low, high *int) {
		if n != tree.root {
			if len(n.items) < tree.degree-1 || len(n.items) > tree.maxItems() {
				t.Fatalf("node item count %d out of [%d,%d]", len(n.items), tree.degree-1, tree.maxItems())
			}
		} else if len(n.items) > tree.maxItems() {
			t.Fatalf("root item count %d > %d", len(n.items), tree.maxItems())
		}
		for i := range n.items {
			if i > 0 && n.items[i-1].Key >= n.items[i].Key {
				t.Fatalf("node keys not sorted: %d >= %d", n.items[i-1].Key, n.items[i].Key)
			}
			if low != nil && n.items[i].Key <= *low {
				t.Fatalf("key %d <= lower bound %d", n.items[i].Key, *low)
			}
			if high != nil && n.items[i].Key >= *high {
				t.Fatalf("key %d >= upper bound %d", n.items[i].Key, *high)
			}
		}
		count += len(n.items)
		if len(n.children) == 0 {
			if leafDepth == -1 {
				leafDepth = depth
			} else if leafDepth != depth {
				t.Fatalf("leaves at different depths: %d and %d", leafDepth, depth)
			}
			return
		}
		if len(n.children) != len(n.items)+1 {
			t.Fatalf("children %d != items+1 %d", len(n.children), len(n.items)+1)
		}
		for i := range n.children {
			var lo, hi *int
			if i > 0 {
				lo = &n.items[i-1].Key
			}
			if i < len(n.items) {
				hi = &n.items[i].Key
			}
			check(n.children[i], depth+1, lo, hi)
		}
	}
	check(tree.root, 0, nil, nil)
	if count != tree.size {
		t.Fatalf("counted %d items, size = %d", count, tree.size)
	}
}

func TestBTreeRandomizedAgainstMap(t *testing.T) {
	for _, degree := range []int{2, 3, 16} {
		tree := NewWithDegree[int, int](compare.Any[int], degree)
		ref := map[int]int{}
		random := rand.New(rand.NewSource(int64(degree)*7919 + 1))

		for step := 0; step < 20000; step++ {
			key := random.Intn(512)
			value := random.Intn(100000)
			switch random.Intn(5) {
			case 0:
				if tree.Put(key, value) {
					if _, exists := ref[key]; !exists {
						ref[key] = value
					}
				}
			case 1:
				tree.Set(key, value)
				ref[key] = value
			case 2:
				tree.Upsert(key, value)
				ref[key] = value
			case 3:
				got, ok := tree.Delete(key)
				want, exists := ref[key]
				if ok != exists {
					t.Fatalf("degree %d step %d: Delete(%d) ok=%v want %v", degree, step, key, ok, exists)
				}
				if ok && got != want {
					t.Fatalf("degree %d step %d: Delete(%d)=%d want %d", degree, step, key, got, want)
				}
				delete(ref, key)
			default:
				got, ok := tree.Get(key)
				want, exists := ref[key]
				if ok != exists || (ok && got != want) {
					t.Fatalf("degree %d step %d: Get(%d)=%d,%v want %d,%v", degree, step, key, got, ok, want, exists)
				}
			}
			if step%500 == 0 {
				validateTree(t, tree, ref)
			}
		}
		validateTree(t, tree, ref)
	}
}

func TestBTreeMonotoneHeight(t *testing.T) {
	const n = 100000
	tree := New[int, int](compare.Any[int])
	for i := 0; i < n; i++ {
		tree.Put(i, i)
	}
	height := tree.Height()
	maxHeight := int(math.Log(float64(n+1))/math.Log(float64(tree.degree))) + 2
	if height <= 0 || height > maxHeight {
		t.Fatalf("height %d out of (0,%d]", height, maxHeight)
	}
	if k, _, _ := tree.Min(); k != 0 {
		t.Fatalf("Min key = %d want 0", k)
	}
	if k, _, _ := tree.Max(); k != n-1 {
		t.Fatalf("Max key = %d want %d", k, n-1)
	}
}

func TestBTreeDeleteEdgeCases(t *testing.T) {
	tree := NewWithDegree[int, int](compare.Any[int], 2)
	for i := 0; i < 64; i++ {
		tree.Put(i, i)
	}
	for i := 0; i < 64; i++ {
		if v, ok := tree.Delete(i); !ok || v != i {
			t.Fatalf("Delete(%d) = %d,%v", i, v, ok)
		}
		remaining := map[int]int{}
		for j := i + 1; j < 64; j++ {
			remaining[j] = j
		}
		validateTree(t, tree, remaining)
	}
	if !tree.Empty() {
		t.Fatal("tree should be empty after deleting every key")
	}
}
