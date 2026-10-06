package hashset

import "testing"

func TestLenAlias(t *testing.T) {
	set := New[int]()
	if set.Len() != 0 {
		t.Fatalf("empty set Len should be 0, got %d", set.Len())
	}
	set.Add(1, 2, 3)
	if set.Len() != 3 {
		t.Fatalf("Len should report 3, got %d", set.Len())
	}
	set.Remove(2)
	if set.Len() != 2 {
		t.Fatalf("Len should report 2 after remove, got %d", set.Len())
	}
}
