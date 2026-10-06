package heap

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestLenAlias(t *testing.T) {
	h := New[int](compare.Any[int])
	if h.Len() != 0 {
		t.Fatalf("empty heap Len should be 0, got %d", h.Len())
	}
	h.Put(2)
	h.Put(1)
	if h.Len() != 2 {
		t.Fatalf("Len should report 2, got %d", h.Len())
	}
}
