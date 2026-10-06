package arraylist

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestLenAlias(t *testing.T) {
	l := New[int](compare.Any[int])
	if l.Len() != 0 {
		t.Fatalf("empty list Len should be 0, got %d", l.Len())
	}
	l.PushBack(1, 2, 3)
	if l.Len() != 3 {
		t.Fatalf("Len should report 3, got %d", l.Len())
	}
}
