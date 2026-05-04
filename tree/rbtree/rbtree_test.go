package rbtree

import (
	"math"
	"math/rand"
	"sort"
	"testing"

	"github.com/474420502/structure/compare"
)

func TestBasicOperations(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		if !tree.Put(i, i*10) {
			t.Fatalf("expected first insert for %d to succeed", i)
		}
	}

	if tree.Put(5, 999) {
		t.Fatal("duplicate Put should not insert")
	}

	if inserted := tree.Set(5, 999); inserted {
		t.Fatal("Set should report false when overwriting")
	}

	value, ok := tree.Get(5)
	if !ok || value != 999 {
		t.Fatalf("expected updated value 999 for key 5, got %v, %v", value, ok)
	}

	if replaced := tree.Upsert(5, 1234); !replaced {
		t.Fatal("Upsert should report replacement on existing key")
	}

	for i := 0; i < 10; i += 2 {
		if _, ok := tree.Remove(i); !ok {
			t.Fatalf("expected remove(%d) to succeed", i)
		}
	}

	for i := 0; i < 10; i++ {
		value, ok := tree.Get(i)
		if i%2 == 0 {
			if ok {
				t.Fatalf("expected key %d to be removed, got value %d", i, value)
			}
		} else if !ok {
			t.Fatalf("expected key %d to remain present", i)
		}
	}

	validateTree(t, tree, map[int]int{1: 10, 3: 30, 5: 1234, 7: 70, 9: 90})
}

func TestIteratorOperations(t *testing.T) {
	tree := New[int, int](compare.Any[int])
	for i := 0; i < 20; i += 2 {
		tree.Put(i, i*2)
	}

	iter := tree.Iterator()
	iter.SeekToFirst()
	for i := 0; i < 20; i += 2 {
		if !iter.Valid() || iter.Key() != i || iter.Value() != i*2 {
			t.Fatalf("expected iterator at key %d, got valid=%v key=%d", i, iter.Valid(), iter.Key())
		}
		iter.Next()
	}
	if iter.Valid() {
		t.Fatal("iterator should be invalid after last element")
	}

	iter.SeekGE(6)
	if !iter.Valid() || iter.Key() != 6 {
		t.Fatalf("SeekGE(6) should land on key 6, got valid=%v key=%d", iter.Valid(), iter.Key())
	}

	if exact := iter.SeekGE(7); exact {
		t.Fatal("SeekGE(7) should report inexact match")
	}
	if !iter.Valid() || iter.Key() != 8 {
		t.Fatalf("SeekGE(7) should land on key 8, got valid=%v key=%d", iter.Valid(), iter.Key())
	}

	if exact := iter.SeekLT(8); !exact {
		t.Fatal("SeekLT(8) should report that key 8 existed")
	}
	if !iter.Valid() || iter.Key() != 6 {
		t.Fatalf("SeekLT(8) should land on key 6, got valid=%v key=%d", iter.Valid(), iter.Key())
	}

	iter.SeekToLast()
	if !iter.Valid() || iter.Key() != 18 {
		t.Fatalf("SeekToLast should land on key 18, got valid=%v key=%d", iter.Valid(), iter.Key())
	}
	iter.Prev()
	if !iter.Valid() || iter.Key() != 16 {
		t.Fatalf("Prev from last should land on key 16, got valid=%v key=%d", iter.Valid(), iter.Key())
	}

	clone := iter.Clone()
	clone.Prev()
	if !iter.Valid() || iter.Key() != 16 {
		t.Fatal("moving clone should not affect original iterator")
	}
	if !clone.Valid() || clone.Key() != 14 {
		t.Fatalf("clone Prev should land on key 14, got valid=%v key=%d", clone.Valid(), clone.Key())
	}
}

func TestRandomizedAgainstMap(t *testing.T) {
	random := rand.New(rand.NewSource(12345))
	tree := New[int, int](compare.Any[int])
	expected := make(map[int]int)

	for step := 0; step < 5000; step++ {
		key := random.Intn(256)
		value := random.Intn(4096)

		switch random.Intn(4) {
		case 0:
			tree.Put(key, value)
			if _, exists := expected[key]; !exists {
				expected[key] = value
			}
		case 1:
			tree.Set(key, value)
			expected[key] = value
		case 2:
			tree.Remove(key)
			delete(expected, key)
		default:
			got, ok := tree.Get(key)
			want, exists := expected[key]
			if ok != exists {
				t.Fatalf("step %d: Get(%d) existence mismatch: got=%v want=%v", step, key, ok, exists)
			}
			if ok && got != want {
				t.Fatalf("step %d: Get(%d)=%d want %d", step, key, got, want)
			}
		}

		if step%100 == 0 {
			validateTree(t, tree, expected)
		}
	}

	validateTree(t, tree, expected)
}

func TestSequentialHeightBound(t *testing.T) {
	const n = 50000
	tree := New[int, int](compare.Any[int])
	for i := 0; i < n; i++ {
		tree.Put(i, i)
	}

	stats := tree.BenchmarkStats()
	maxHeight := int(math.Ceil(2 * math.Log2(float64(n+1))))
	if stats.Height > maxHeight {
		t.Fatalf("sequential insert height %d exceeds red-black upper bound %d", stats.Height, maxHeight)
	}
}

func validateTree(t *testing.T, tree *Tree[int, int], expected map[int]int) {
	t.Helper()

	if int(tree.Size()) != len(expected) {
		t.Fatalf("size mismatch: got %d want %d", tree.Size(), len(expected))
	}

	if len(expected) == 0 {
		if tree.root != nil {
			t.Fatal("expected empty tree root to be nil")
		}
		return
	}

	if tree.root == nil {
		t.Fatal("expected non-empty tree to have a root")
	}
	if tree.root.Parent != nil {
		t.Fatal("root parent must be nil")
	}
	if tree.root.Color != black {
		t.Fatal("root must be black")
	}

	verifyNode(t, tree.root, nil, nil)

	wantKeys := make([]int, 0, len(expected))
	for key := range expected {
		wantKeys = append(wantKeys, key)
	}
	sort.Ints(wantKeys)

	gotKeys := make([]int, 0, len(expected))
	tree.Traverse(func(key int, value int) bool {
		wantValue, ok := expected[key]
		if !ok {
			t.Fatalf("unexpected key %d in traversal", key)
		}
		if value != wantValue {
			t.Fatalf("value mismatch for key %d: got %d want %d", key, value, wantValue)
		}
		gotKeys = append(gotKeys, key)
		return true
	})

	if len(gotKeys) != len(wantKeys) {
		t.Fatalf("traversal size mismatch: got %d want %d", len(gotKeys), len(wantKeys))
	}
	for i := range wantKeys {
		if gotKeys[i] != wantKeys[i] {
			t.Fatalf("traversal key mismatch at %d: got %d want %d", i, gotKeys[i], wantKeys[i])
		}
	}

	iterKeys := make([]int, 0, len(expected))
	iter := tree.Iterator()
	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		iterKeys = append(iterKeys, iter.Key())
	}
	for i := range wantKeys {
		if iterKeys[i] != wantKeys[i] {
			t.Fatalf("iterator key mismatch at %d: got %d want %d", i, iterKeys[i], wantKeys[i])
		}
	}
}

func verifyNode(t *testing.T, current *Node[int, int], min *int, max *int) int {
	t.Helper()

	if current == nil {
		return 1
	}

	if min != nil && current.Key <= *min {
		t.Fatalf("bst order violated: key %d <= min %d", current.Key, *min)
	}
	if max != nil && current.Key >= *max {
		t.Fatalf("bst order violated: key %d >= max %d", current.Key, *max)
	}

	if current.Left != nil && current.Left.Parent != current {
		t.Fatalf("left parent link broken at key %d", current.Key)
	}
	if current.Right != nil && current.Right.Parent != current {
		t.Fatalf("right parent link broken at key %d", current.Key)
	}

	if current.Color == red {
		if colorOf(current.Left) == red || colorOf(current.Right) == red {
			t.Fatalf("red node %d has red child", current.Key)
		}
	}

	leftBlackHeight := verifyNode(t, current.Left, min, &current.Key)
	rightBlackHeight := verifyNode(t, current.Right, &current.Key, max)
	if leftBlackHeight != rightBlackHeight {
		t.Fatalf("black-height mismatch at key %d: left=%d right=%d", current.Key, leftBlackHeight, rightBlackHeight)
	}

	if current.Color == black {
		return leftBlackHeight + 1
	}
	return leftBlackHeight
}