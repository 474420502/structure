package hashmap

import "testing"

func TestSemanticAliases(t *testing.T) {
	hm := New()

	if !hm.InsertIfAbsent("a", 1) {
		t.Fatal("InsertIfAbsent should insert a missing key")
	}
	if hm.InsertIfAbsent("a", 2) {
		t.Fatal("InsertIfAbsent should not overwrite an existing key")
	}
	if value, _ := hm.Get("a"); value != 1 {
		t.Fatalf("existing value should remain unchanged, got %v", value)
	}

	if hm.Upsert("b", 2) {
		t.Fatal("Upsert should report false when inserting a new key")
	}
	if !hm.Upsert("b", 3) {
		t.Fatal("Upsert should report true when replacing an existing key")
	}
	if value, _ := hm.Get("b"); value != 3 {
		t.Fatalf("Upsert should replace the value, got %v", value)
	}

	if hm.Len() != 2 {
		t.Fatalf("Len should report 2, got %d", hm.Len())
	}

	if value, ok := hm.Delete("a"); !ok || value != 1 {
		t.Fatalf("Delete should return removed value and ok=true, got (%v, %v)", value, ok)
	}
	if _, ok := hm.Delete("missing"); ok {
		t.Fatal("Delete should report ok=false for a missing key")
	}
}