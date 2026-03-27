package indextree

import (
	"sort"
	"testing"

	"github.com/474420502/structure/compare"
)

func TestTreeNew(t *testing.T) {
	tree := New(compare.Any[int])
	if tree == nil {
		t.Error("New should return non-nil tree")
	}
	if tree.Size() != 0 {
		t.Error("new tree should have size 0")
	}
}

func TestTreeSize(t *testing.T) {
	tree := New(compare.Any[int])
	if tree.Size() != 0 {
		t.Error("empty tree should have size 0")
	}

	tree.Put(1, 100)
	if tree.Size() != 1 {
		t.Error("tree with one node should have size 1")
	}
}

func TestTreeClear(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}
	tree.Clear()
	if tree.Size() != 0 {
		t.Error("tree should have size 0 after Clear")
	}
}

func TestTreeGet(t *testing.T) {
	tree := New(compare.Any[int])
	tree.Put(1, 100)
	tree.Put(3, 300)

	if _, ok := tree.Get(2); ok {
		t.Error("Get non-existent key should return false")
	}
	if v, ok := tree.Get(1); !ok || v != 100 {
		t.Error("Get existing key should return correct value")
	}
}

func TestTreePutDuplicate(t *testing.T) {
	tree := New(compare.Any[int])
	if !tree.Put(1, 100) {
		t.Error("Put new key should return true")
	}
	if tree.Put(1, 200) {
		t.Error("Put duplicate key should return false")
	}
	if v, _ := tree.Get(1); v != 100 {
		t.Error("original value should not be overwritten by Put")
	}
}

func TestTreeSetDuplicate(t *testing.T) {
	tree := New(compare.Any[int])
	if !tree.Set(1, 100) {
		t.Error("Set new key should return false")
	}
	if !tree.Set(1, 200) {
		t.Error("Set duplicate key should return true")
	}
	if v, _ := tree.Get(1); v != 200 {
		t.Error("value should be overwritten by Set")
	}
}

func TestTreeRemove(t *testing.T) {
	tree := New(compare.Any[int])
	tree.Put(1, 100)

	if tree.Remove(2) != nil {
		t.Error("Remove non-existent should return nil")
	}
	if tree.Remove(1) != 100 {
		t.Error("Remove should return correct value")
	}
	if tree.Size() != 0 {
		t.Error("size should be 0 after removing last node")
	}
}

func TestTreeTraverse(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	var keys []int
	tree.Traverse(func(k int, v interface{}) bool {
		keys = append(keys, k)
		return true
	})

	if len(keys) != 100 {
		t.Errorf("Traverse should visit 100 nodes, got %d", len(keys))
	}

	for i := 0; i < len(keys)-1; i++ {
		if keys[i] >= keys[i+1] {
			t.Error("Traverse should visit keys in ascending order")
		}
	}
}

func TestTreeTraverseEarlyExit(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	count := 0
	tree.Traverse(func(k int, v interface{}) bool {
		count++
		if k >= 10 {
			return false
		}
		return true
	})

	if count != 11 {
		t.Errorf("Traverse with early exit should visit 11 nodes, got %d", count)
	}
}

func TestTreeValues(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	values := tree.Values()
	if len(values) != 100 {
		t.Errorf("Values should return 100 elements, got %d", len(values))
	}

	for i := 0; i < len(values)-1; i++ {
		if values[i].(int) >= values[i+1].(int) {
			t.Error("Values should be in ascending order")
		}
	}
}

func TestTreeIndex(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	for i := 0; i < 100; i++ {
		key, val := tree.Index(int64(i))
		if key != i || val != i*2 {
			t.Errorf("Index(%d) = (%d, %d), want (%d, %d)", i, key, val, i, i*2)
		}
	}
}

func TestTreeIndexOf(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	for i := 0; i < 100; i++ {
		idx := tree.IndexOf(i)
		if idx != int64(i) {
			t.Errorf("IndexOf(%d) = %d, want %d", i, idx, i)
		}
	}
}

func TestTreeConsistency(t *testing.T) {
	tree := New(compare.Any[int])
	var keys []int
	for i := 0; i < 100; i++ {
		k := i * 2
		if tree.Put(k, k) {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)

	for i, k := range keys {
		idx := tree.IndexOf(k)
		if idx != int64(i) {
			t.Errorf("IndexOf(%d) = %d, want %d", k, idx, i)
		}

		key, _ := tree.Index(int64(i))
		if key != k {
			t.Errorf("Index(%d) = %d, want %d", i, key, k)
		}
	}
}

func TestTreeSplit(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	splitKey := 55
	rightTree := tree.Split(splitKey)

	if tree.Size() != 55 {
		t.Errorf("left tree should have 55 elements, got %d", tree.Size())
	}

	if rightTree.Size() != 45 {
		t.Errorf("right tree should have 45 elements, got %d", rightTree.Size())
	}

	for i := 0; i < 55; i++ {
		if _, ok := tree.Get(i); !ok {
			t.Errorf("left tree should contain key %d", i)
		}
	}

	for i := 55; i < 100; i++ {
		if _, ok := rightTree.Get(i); !ok {
			t.Errorf("right tree should contain key %d", i)
		}
	}
}

func TestTreeSplitContain(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	splitKey := 55
	rightTree := tree.SplitContain(splitKey)

	if tree.Size() != 56 {
		t.Errorf("left tree should have 56 elements, got %d", tree.Size())
	}

	if rightTree.Size() != 44 {
		t.Errorf("right tree should have 44 elements, got %d", rightTree.Size())
	}

	for i := 0; i <= 55; i++ {
		if _, ok := tree.Get(i); !ok {
			t.Errorf("left tree should contain key %d", i)
		}
	}

	for i := 56; i < 100; i++ {
		if _, ok := rightTree.Get(i); !ok {
			t.Errorf("right tree should contain key %d", i)
		}
	}

	if _, ok := rightTree.Get(55); ok {
		t.Error("right tree should not contain split key 55")
	}
}

func TestTreeTrim(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	tree.Trim(25, 75)

	if tree.Size() != 51 {
		t.Errorf("Trim(25, 75) should keep 51 elements, got %d", tree.Size())
	}

	for i := 0; i < 25; i++ {
		if _, ok := tree.Get(i); ok {
			t.Errorf("key %d should be trimmed", i)
		}
	}

	for i := 25; i <= 75; i++ {
		if _, ok := tree.Get(i); !ok {
			t.Errorf("key %d should be kept", i)
		}
	}

	for i := 76; i < 100; i++ {
		if _, ok := tree.Get(i); ok {
			t.Errorf("key %d should be trimmed", i)
		}
	}
}

func TestTreeRemoveRange(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	tree.RemoveRange(25, 75)

	if tree.Size() != 49 {
		t.Errorf("RemoveRange(25, 75) should leave 49 elements, got %d", tree.Size())
	}

	for i := 0; i < 25; i++ {
		if _, ok := tree.Get(i); !ok {
			t.Errorf("key %d should remain", i)
		}
	}

	for i := 25; i <= 75; i++ {
		if _, ok := tree.Get(i); ok {
			t.Errorf("key %d should be removed", i)
		}
	}

	for i := 76; i < 100; i++ {
		if _, ok := tree.Get(i); !ok {
			t.Errorf("key %d should remain", i)
		}
	}
}

func TestTreeCheckAfterOperations(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 500; i++ {
		tree.Put(i, i)
		tree.check()
	}

	for i := 0; i < 200; i++ {
		tree.Remove(i)
		tree.check()
	}

	for i := 0; i < 200; i++ {
		tree.Set(i, i)
		tree.check()
	}
}

func TestTreeIteratorConsistency(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	var iterValues []int
	iter := tree.Iterator()
	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		iterValues = append(iterValues, iter.Key())
	}

	var traverseValues []int
	tree.Traverse(func(k int, v interface{}) bool {
		traverseValues = append(traverseValues, k)
		return true
	})

	if len(iterValues) != len(traverseValues) {
		t.Error("Iterator and Traverse should return same number of elements")
	}

	for i := 0; i < len(iterValues); i++ {
		if iterValues[i] != traverseValues[i] {
			t.Error("Iterator and Traverse should return same sequence")
		}
	}
}
