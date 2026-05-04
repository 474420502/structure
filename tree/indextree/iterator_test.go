package indextree

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestIteratorBasic(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	iter := tree.Iterator()

	iter.SeekToFirst()
	if !iter.Valid() {
		t.Error("iterator should be valid after SeekToFirst")
	}
	if iter.Key() != 0 {
		t.Error("first key should be 0")
	}
	if iter.Value() != 0 {
		t.Error("first value should be 0")
	}

	iter.SeekToLast()
	if !iter.Valid() {
		t.Error("iterator should be valid after SeekToLast")
	}
	if iter.Key() != 99 {
		t.Error("last key should be 99")
	}
}

func TestIteratorTraversal(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	iter := tree.Iterator()
	var count int64
	var prevKey int = -1

	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		if iter.Key() != prevKey+1 {
			t.Error("keys should be in ascending order")
		}
		prevKey = iter.Key()
		count++
	}

	if count != 100 {
		t.Errorf("expected 100 iterations, got %d", count)
	}
}

func TestIteratorPrevNext(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 10; i++ {
		tree.Put(i, i*2)
	}

	iter := tree.Iterator()
	iter.SeekToFirst()

	for i := 0; i < 10; i++ {
		if !iter.Valid() {
			t.Errorf("iterator should be valid at index %d", i)
		}
		if iter.Key() != i {
			t.Errorf("expected key %d, got %d", i, iter.Key())
		}
		iter.Next()
	}

	if iter.Valid() {
		t.Error("iterator should be invalid after end")
	}
}

func TestIteratorClone(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 10; i++ {
		tree.Put(i, i*2)
	}

	iter1 := tree.Iterator()
	iter1.SeekToFirst()
	iter1.Next()
	iter1.Next()

	iter2 := iter1.Clone()

	if !iter2.Valid() || iter2.Key() != iter1.Key() || iter2.Value() != iter1.Value() {
		t.Error("clone should have same state as original")
	}

	iter2.Next()
	if iter1.Key() == iter2.Key() {
		t.Error("modifying clone should not affect original")
	}
}

func TestIteratorSeekByIndex(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	iter := tree.Iterator()

	iter.SeekByIndex(0)
	if !iter.Valid() || iter.Key() != 0 {
		t.Error("SeekByIndex(0) should return key 0")
	}

	iter.SeekByIndex(50)
	if !iter.Valid() || iter.Key() != 50 {
		t.Error("SeekByIndex(50) should return key 50")
	}

	iter.SeekByIndex(99)
	if !iter.Valid() || iter.Key() != 99 {
		t.Error("SeekByIndex(99) should return key 99")
	}
}

func TestIteratorIndex(t *testing.T) {
	tree := New(compare.Any[int])
	for i := 0; i < 100; i++ {
		tree.Put(i, i*2)
	}

	iter := tree.Iterator()
	iter.SeekByIndex(50)
	if iter.Index() != 50 {
		t.Errorf("expected index 50, got %d", iter.Index())
	}

	iter.SeekToFirst()
	if iter.Index() != 0 {
		t.Errorf("expected index 0, got %d", iter.Index())
	}

	iter.SeekToLast()
	if iter.Index() != 99 {
		t.Errorf("expected index 99, got %d", iter.Index())
	}
}

func TestIteratorEmptyTree(t *testing.T) {
	tree := New(compare.Any[int])
	iter := tree.Iterator()

	iter.SeekToFirst()
	if iter.Valid() {
		t.Error("iterator on empty tree should not be valid")
	}

	iter.SeekToLast()
	if iter.Valid() {
		t.Error("iterator on empty tree should not be valid")
	}

	iter.SeekByIndex(0)
	if iter.Valid() {
		t.Error("iterator on empty tree should not be valid")
	}
}

func TestIteratorSingleElement(t *testing.T) {
	tree := New(compare.Any[int])
	tree.Put(42, 84)

	iter := tree.Iterator()

	iter.SeekToFirst()
	if !iter.Valid() || iter.Key() != 42 || iter.Value() != 84 {
		t.Error("SeekToFirst failed on single element")
	}

	iter.SeekToLast()
	if !iter.Valid() || iter.Key() != 42 || iter.Value() != 84 {
		t.Error("SeekToLast failed on single element")
	}
}

func TestIteratorSeekBounds(t *testing.T) {
	tree := New(compare.Any[int])
	for _, value := range []int{0, 2, 4, 6, 8, 10} {
		tree.Put(value, value)
	}

	iter := tree.Iterator()

	if exact := iter.SeekGE(5); exact {
		t.Fatal("SeekGE(5) should report inexact match")
	}
	if !iter.Valid() || iter.Key() != 6 || iter.Index() != 3 {
		t.Fatalf("SeekGE(5) should land on key 6 at index 3, got valid=%v key=%d index=%d", iter.Valid(), iter.Key(), iter.Index())
	}

	if exact := iter.SeekGT(6); !exact {
		t.Fatal("SeekGT(6) should report exact match before moving to successor")
	}
	if !iter.Valid() || iter.Key() != 8 || iter.Index() != 4 {
		t.Fatalf("SeekGT(6) should land on key 8 at index 4, got valid=%v key=%d index=%d", iter.Valid(), iter.Key(), iter.Index())
	}

	if exact := iter.SeekLE(5); exact {
		t.Fatal("SeekLE(5) should report inexact match")
	}
	if !iter.Valid() || iter.Key() != 4 || iter.Index() != 2 {
		t.Fatalf("SeekLE(5) should land on key 4 at index 2, got valid=%v key=%d index=%d", iter.Valid(), iter.Key(), iter.Index())
	}

	if exact := iter.SeekLT(4); !exact {
		t.Fatal("SeekLT(4) should report exact match before moving to predecessor")
	}
	if !iter.Valid() || iter.Key() != 2 || iter.Index() != 1 {
		t.Fatalf("SeekLT(4) should land on key 2 at index 1, got valid=%v key=%d index=%d", iter.Valid(), iter.Key(), iter.Index())
	}

	if exact := iter.SeekGE(11); exact || iter.Valid() {
		t.Fatalf("SeekGE(11) should be invalid, got exact=%v valid=%v", exact, iter.Valid())
	}
	if iter.Index() != tree.Size() {
		t.Fatalf("SeekGE(11) should set index to size, got %d", iter.Index())
	}

	if exact := iter.SeekLT(-1); exact || iter.Valid() {
		t.Fatalf("SeekLT(-1) should be invalid, got exact=%v valid=%v", exact, iter.Valid())
	}
	if iter.Index() != -1 {
		t.Fatalf("SeekLT(-1) should set index to -1, got %d", iter.Index())
	}
}
