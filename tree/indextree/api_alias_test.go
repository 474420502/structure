package indextree

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestSemanticAliases(t *testing.T) {
	tree := New(compare.Any[int])

	if !tree.InsertIfAbsent(1, "a") {
		t.Fatal("InsertIfAbsent should insert a missing key")
	}
	if tree.InsertIfAbsent(1, "b") {
		t.Fatal("InsertIfAbsent should not overwrite an existing key")
	}
	if value, _ := tree.Get(1); value != "a" {
		t.Fatalf("existing value should remain unchanged, got %v", value)
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

	if value, ok := tree.Delete(1); !ok || value != "a" {
		t.Fatalf("Delete should return removed value and ok=true, got (%v, %v)", value, ok)
	}
	if _, ok := tree.Delete(999); ok {
		t.Fatal("Delete should report ok=false for a missing key")
	}
}