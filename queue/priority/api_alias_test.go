package treequeue

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestSemanticAliases(t *testing.T) {
	queue := New[int, string](compare.AnyEx[int])

	if !queue.InsertIfAbsent(1, "a") {
		t.Fatal("InsertIfAbsent should insert a missing key")
	}
	if queue.InsertIfAbsent(1, "b") {
		t.Fatal("InsertIfAbsent should not add duplicates")
	}
	if queue.Len() != 1 {
		t.Fatalf("InsertIfAbsent should not grow size for duplicates, got %d", queue.Len())
	}

	if queue.Upsert(2, "b") {
		t.Fatal("Upsert should report false when inserting a new key")
	}
	if !queue.Upsert(2, "c") {
		t.Fatal("Upsert should report true when replacing an existing key")
	}
	if value, _ := queue.Get(2); value != "c" {
		t.Fatalf("Upsert should replace the value, got %v", value)
	}
	if queue.Len() != 2 {
		t.Fatalf("Len should report 2, got %d", queue.Len())
	}

	if !queue.Put(2, "shadow") {
		t.Fatal("legacy Put should still allow duplicate insertion")
	}
	if queue.Len() != 3 {
		t.Fatalf("legacy Put should still grow size for duplicates, got %d", queue.Len())
	}
	if values := queue.Gets(2); len(values) != 2 {
		t.Fatalf("legacy duplicate behavior should remain available, got %v", values)
	}

	if value, ok := queue.Delete(1); !ok || value != "a" {
		t.Fatalf("Delete should return removed value and ok=true, got (%v, %v)", value, ok)
	}
	if _, ok := queue.Delete(999); ok {
		t.Fatal("Delete should report ok=false for a missing key")
	}
}