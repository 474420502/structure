package experiment

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestExperimentTreeSurface(t *testing.T) {
	tree := New[int, int](compare.Any[int])
	for i := 0; i < 20; i++ {
		tree.Put(i, i*10)
	}
	if tree.Len() != 20 || tree.Size() != 20 {
		t.Fatalf("size = %d/%d want 20", tree.Len(), tree.Size())
	}
	if v, ok := tree.Get(7); !ok || v != 70 {
		t.Fatalf("Get(7) = %d,%v want 70,true", v, ok)
	}

	tree.Set(7, 71)
	if v, _ := tree.Get(7); v != 71 {
		t.Fatalf("Set value = %d want 71", v)
	}
	if !tree.Upsert(7, 72) {
		t.Fatal("Upsert should replace an existing key")
	}
	if tree.Upsert(100, 1000) {
		t.Fatal("Upsert should report false when inserting")
	}
	if !tree.InsertIfAbsent(101, 1010) {
		t.Fatal("InsertIfAbsent should insert a missing key")
	}
	if tree.InsertIfAbsent(7, 0) {
		t.Fatal("InsertIfAbsent should reject an existing key")
	}

	tree.Remove(7)
	tree.Delete(8)
	if tree.Len() != 20 {
		t.Fatalf("Len = %d want 20", tree.Len())
	}

	var previous = -1
	visited := 0
	tree.Traverse(func(k, v int) bool {
		if k <= previous {
			t.Fatalf("Traverse out of order at %d after %d", k, previous)
		}
		previous = k
		visited++
		return true
	})
	if visited != tree.Len() {
		t.Fatalf("Traverse visited %d want %d", visited, tree.Len())
	}
	if len(tree.Values()) != tree.Len() {
		t.Fatal("Values length mismatch")
	}
	if tree.Height() <= 0 {
		t.Fatal("Height should be positive")
	}

	iter := tree.Iterator()
	iter.SeekToFirst()
	if !iter.Valid() || iter.Key() != 0 {
		t.Fatalf("SeekToFirst key = %d", iter.Key())
	}
	iter.SeekToLast()
	if !iter.Valid() || iter.Key() != 101 {
		t.Fatalf("SeekToLast key = %d want 101", iter.Key())
	}
	if exact := iter.SeekGE(9); !exact || iter.Key() != 9 {
		t.Fatalf("SeekGE(9) exact=%v key=%d", exact, iter.Key())
	}
	if exact := iter.SeekLE(9); !exact || iter.Key() != 9 {
		t.Fatalf("SeekLE(9) exact=%v key=%d", exact, iter.Key())
	}
	if exact := iter.SeekGT(9); !exact || iter.Key() != 10 {
		t.Fatalf("SeekGT(9) exact=%v key=%d", exact, iter.Key())
	}
	clone := iter.Clone()
	clone.Prev()
	if iter.Key() != 10 || clone.Key() != 9 {
		t.Fatalf("Clone mismatch: orig=%d clone=%d", iter.Key(), clone.Key())
	}

	tree.ResetBenchmarkStats()
	stats := tree.BenchmarkStats()
	if stats.Height <= 0 {
		t.Fatal("BenchmarkStats height should be positive")
	}

	tree.Clear()
	if tree.Len() != 0 {
		t.Fatal("Clear should empty the tree")
	}
}

func TestExperimentNewEx(t *testing.T) {
	tree := NewEx[int, int](compare.Any[int], 1)
	for i := 0; i < 32; i++ {
		tree.Put(i, i)
	}
	if tree.Len() != 32 {
		t.Fatalf("Len = %d want 32", tree.Len())
	}
	tree.ResetBenchmarkStats()
	_ = tree.BenchmarkStats()
}

func TestShiftToleranceSurface(t *testing.T) {
	tree := NewShiftTolerance[int, int](compare.Any[int], 2)
	for i := 0; i < 64; i++ {
		tree.Put(i*3%64, i)
	}
	if tree.Size() != 64 {
		t.Fatalf("Size = %d want 64", tree.Size())
	}
	if _, ok := tree.Get(10); !ok {
		t.Fatal("Get(10) should be present")
	}
	if tree.Height() <= 0 {
		t.Fatal("Height should be positive")
	}
	tree.ResetBenchmarkStats()
	if stats := tree.BenchmarkStats(); stats.Height <= 0 {
		t.Fatal("BenchmarkStats height should be positive")
	}
	if defaultTree := NewIndexTreeDefault[int, int](compare.Any[int]); defaultTree == nil {
		t.Fatal("NewIndexTreeDefault should return a tree")
	}
}
