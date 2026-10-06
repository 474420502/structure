package hashset

import (
	"sort"
	"testing"
)

func TestHashSetBehaviors(t *testing.T) {
	set := New[int]()
	if !set.Empty() || set.Size() != 0 || set.Len() != 0 {
		t.Fatal("new set should be empty")
	}

	set.Add(1, 2, 3, 3, 2)
	if set.Len() != 3 {
		t.Fatalf("Len = %d want 3", set.Len())
	}
	if !set.Contains(1) || !set.Contains(2) || !set.Contains(3) {
		t.Fatal("Contains should find inserted items")
	}
	if set.Contains(4) {
		t.Fatal("Contains should be false for missing item")
	}

	// Values returns an unordered snapshot.
	ints := set.Values()
	sort.Ints(ints)
	if len(ints) != 3 || ints[0] != 1 || ints[1] != 2 || ints[2] != 3 {
		t.Fatalf("Values = %v want [1 2 3]", ints)
	}

	if s := set.String(); s != "[1,2,3]" && s != "[1,3,2]" && s != "[2,1,3]" && s != "[2,3,1]" && s != "[3,1,2]" && s != "[3,2,1]" {
		t.Fatalf("String = %q unexpected", s)
	}

	set.Remove(2, 99)
	if set.Contains(2) || set.Len() != 2 {
		t.Fatalf("Remove failed: len=%d", set.Len())
	}
	if set.Empty() {
		t.Fatal("set should not be empty")
	}

	set.Clear()
	if !set.Empty() || set.Len() != 0 {
		t.Fatal("Clear should empty the set")
	}
}
