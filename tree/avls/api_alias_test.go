package avls

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestSemanticAliases(t *testing.T) {
	tree := New[int, string](compare.AnyEx[int])

	if !tree.InsertIfAbsent(1, "a") {
		t.Fatal("InsertIfAbsent should insert a missing key")
	}
	if tree.InsertIfAbsent(1, "b") {
		t.Fatal("InsertIfAbsent should not add duplicates")
	}
	if tree.Len() != 1 {
		t.Fatalf("InsertIfAbsent should not grow size for duplicates, got %d", tree.Len())
	}

	if tree.Upsert(2, "b") {
		t.Fatal("Upsert should report false when inserting a new key")
	}
	if !tree.Upsert(2, "c") {
		t.Fatal("Upsert should report true when replacing an existing key")
	}
	if value, _ := tree.Get(2); value != "c" {
		t.Fatalf("Upsert should replace the value, got %v", value)
	}
	if tree.Len() != 2 {
		t.Fatalf("Len should report 2, got %d", tree.Len())
	}

	if tree.Put(2, "shadow") != true {
		t.Fatal("legacy Put should still allow duplicate insertion")
	}
	if tree.Len() != 3 {
		t.Fatalf("legacy Put should still grow size for duplicates, got %d", tree.Len())
	}

	if value, ok := tree.Delete(1); !ok || value != "a" {
		t.Fatalf("Delete should return removed value and ok=true, got (%v, %v)", value, ok)
	}
	if _, ok := tree.Delete(999); ok {
		t.Fatal("Delete should report ok=false for a missing key")
	}
}