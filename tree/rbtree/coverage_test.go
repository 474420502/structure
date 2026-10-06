package rbtree

import (
	"reflect"
	"testing"

	"github.com/474420502/structure/compare"
)

func TestRBTreeEmptyAndBoundaries(t *testing.T) {
	tree := New[int, int](compare.Any[int])
	if tree.Len() != 0 || tree.Size() != 0 || tree.Height() != 0 {
		t.Fatal("new rbtree should be empty")
	}
	if _, ok := tree.Get(1); ok {
		t.Fatal("Get on empty tree should be false")
	}
	if _, ok := tree.Delete(1); ok {
		t.Fatal("Delete on empty tree should be false")
	}

	iter := tree.Iterator()
	iter.SeekToFirst()
	if iter.Valid() {
		t.Fatal("empty iterator should be invalid")
	}
	iter.SeekToLast()
	if iter.Valid() {
		t.Fatal("empty iterator should be invalid")
	}
	if iter.SeekGE(1) || iter.SeekGT(1) || iter.SeekLE(1) || iter.SeekLT(1) {
		t.Fatal("seek on empty tree should report no exact match")
	}

	for i := 0; i < 10; i++ {
		tree.Put(i, i*10)
	}
	if tree.Height() <= 0 {
		t.Fatal("Height should be positive")
	}

	// Exact-match flags follow the standard contract.
	if exact := iter.SeekGE(5); !exact || iter.Key() != 5 {
		t.Fatalf("SeekGE(5) exact=%v key=%d", exact, iter.Key())
	}
	if exact := iter.SeekGT(5); !exact || iter.Key() != 6 {
		t.Fatalf("SeekGT(5) exact=%v key=%d", exact, iter.Key())
	}
	if exact := iter.SeekLE(5); !exact || iter.Key() != 5 {
		t.Fatalf("SeekLE(5) exact=%v key=%d", exact, iter.Key())
	}
	if exact := iter.SeekLT(5); !exact || iter.Key() != 4 {
		t.Fatalf("SeekLT(5) exact=%v key=%d", exact, iter.Key())
	}
	if exact := iter.SeekGE(100); exact || iter.Valid() {
		t.Fatalf("SeekGE(100) exact=%v valid=%v", exact, iter.Valid())
	}
	if exact := iter.SeekLE(-1); exact || iter.Valid() {
		t.Fatalf("SeekLE(-1) exact=%v valid=%v", exact, iter.Valid())
	}

	// Values are returned in ascending key order.
	if got := tree.Values(); !reflect.DeepEqual(got, []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}) {
		t.Fatalf("Values = %v", got)
	}

	clone := tree.Iterator()
	clone.SeekToLast()
	other := clone.Clone()
	other.Prev()
	if clone.Key() != 9 || other.Key() != 8 {
		t.Fatalf("Clone mismatch: %d vs %d", clone.Key(), other.Key())
	}

	tree.Clear()
	if tree.Len() != 0 || tree.Height() != 0 {
		t.Fatal("Clear should empty the tree")
	}
}

func TestRBTreeWriteSemantics(t *testing.T) {
	tree := New[int, string](compare.Any[int])

	if !tree.Put(1, "a") {
		t.Fatal("Put should insert")
	}
	if tree.Put(1, "b") {
		t.Fatal("Put should preserve the existing value")
	}
	if v, _ := tree.Get(1); v != "a" {
		t.Fatalf("Put changed value to %q", v)
	}

	if tree.Set(1, "b") {
		// Set overwrites existing and reports false (no new insert).
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

	if !tree.InsertIfAbsent(3, "e") {
		t.Fatal("InsertIfAbsent should insert")
	}
	if tree.InsertIfAbsent(3, "f") {
		t.Fatal("InsertIfAbsent should reject an existing key")
	}

	if v, ok := tree.Delete(2); !ok || v != "d" {
		t.Fatalf("Delete(2) = %q,%v", v, ok)
	}
	if _, ok := tree.Delete(2); ok {
		t.Fatal("second Delete should be false")
	}
}
