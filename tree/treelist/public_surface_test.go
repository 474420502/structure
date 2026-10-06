package treelist

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestTreeListPublicSurface(t *testing.T) {
	tree := New[int, int](compare.Any[int])
	if tree.Tail() != nil || tree.Head() != nil {
		t.Fatal("empty tree Head/Tail should be nil")
	}
	if _, ok := tree.Get(1); ok {
		t.Fatal("Get on empty tree should be false")
	}
	if got := tree.RemoveIndex(0); got != nil {
		t.Fatal("RemoveIndex on empty tree should be nil")
	}

	for i := 0; i < 20; i++ {
		tree.Set(i, i*10)
	}

	if tail := tree.Tail(); tail == nil || tail.Key != 19 {
		t.Fatalf("Tail = %v want key 19", tail)
	}
	if head := tree.Head(); head == nil || head.Key != 0 {
		t.Fatalf("Head = %v want key 0", head)
	}
	if s := tree.Index(3); s == nil || s.Key != 3 {
		t.Fatalf("Index(3) = %v", s)
	}
	if got := tree.IndexOf(0); got != 0 {
		t.Fatalf("IndexOf(0) = %d want 0", got)
	}
	if got := tree.IndexOf(19); got != 19 {
		t.Fatalf("IndexOf(19) = %d want 19", got)
	}

	// Rotation counters are part of the public surface.
	_ = tree.SingleRotations()
	_ = tree.DoubleRotations()
	tree.ResetRotations()
	if tree.SingleRotations() != 0 || tree.DoubleRotations() != 0 {
		t.Fatal("ResetRotations should zero the counters")
	}

	// Iterator Slice exposes the current element.
	iter := tree.Iterator()
	iter.SeekToFirst()
	if !iter.Valid() {
		t.Fatal("iterator should be valid after SeekToFirst")
	}
	if slice := iter.Slice(); slice == nil || slice.Key != 0 {
		t.Fatalf("iterator Slice = %v", slice)
	}
	if iter.Index() != 0 {
		t.Fatalf("iterator Index = %d want 0", iter.Index())
	}

	// Range iterator direction round-trips.
	r := tree.IteratorRange()
	r.SetDirection(Forward)
	if r.Direction() != Forward {
		t.Fatalf("Direction = %v want Forward", r.Direction())
	}
	r.SetDirection(Reverse)
	if r.Direction() != Reverse {
		t.Fatalf("Direction = %v want Reverse", r.Direction())
	}
	// GE2LE positions the bounds without changing the configured direction.
	r.SetDirection(Forward)
	r.GE2LE(2, 5)
	if r.Direction() != Forward {
		t.Fatalf("GE2LE changed direction to %v", r.Direction())
	}
	count := 0
	r.Range(func(cur *SliceIndex[int, int]) bool {
		count++
		return true
	})
	if count != 4 {
		t.Fatalf("range [2,5] visited %d want 4", count)
	}

	// RemoveIndex returns the removed slice.
	removed := tree.RemoveIndex(0)
	if removed == nil || removed.Key != 0 {
		t.Fatalf("RemoveIndex(0) = %v", removed)
	}
	if tree.Len() != 19 {
		t.Fatalf("Len after RemoveIndex = %d want 19", tree.Len())
	}
	if s := new(Slice[int, int]); s.String() == "" {
		t.Fatal("Slice.String should not be empty")
	}
}
