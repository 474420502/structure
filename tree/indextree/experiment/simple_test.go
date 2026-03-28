package experiment

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestSimplePut(t *testing.T) {
	tree := New[int64, int](compare.Any[int64])

	tree.Put(1, 1)
	if tree.Size() != 1 {
		t.Errorf("Expected size 1, got %d", tree.Size())
	}

	tree.Put(2, 2)
	if tree.Size() != 2 {
		t.Errorf("Expected size 2, got %d", tree.Size())
	}

	tree.Put(3, 3)
	if tree.Size() != 3 {
		t.Errorf("Expected size 3, got %d", tree.Size())
	}
}
