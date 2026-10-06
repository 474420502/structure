package heap

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/474420502/structure/compare"
)

func TestHeapEdges(t *testing.T) {
	h := New[int](compare.Any[int])
	if !h.Empty() || h.Len() != 0 {
		t.Fatal("new heap should be empty")
	}
	if _, ok := h.Top(); ok {
		t.Fatal("Top on empty heap should be false")
	}
	if _, ok := h.Pop(); ok {
		t.Fatal("Pop on empty heap should be false")
	}

	h.Put(2)
	h.Put(1)
	if v, ok := h.Top(); !ok || v != 1 {
		t.Fatalf("Top = %d,%v want 1,true", v, ok)
	}

	h.Reset()
	if !h.Empty() || h.Len() != 0 {
		t.Fatal("Reset should empty the heap")
	}

	h.Put(5)
	h.Clear()
	if !h.Empty() || h.Len() != 0 {
		t.Fatal("Clear should empty the heap")
	}
	if v, ok := h.Top(); ok || v != 0 {
		t.Fatalf("Top after Clear = %d,%v", v, ok)
	}
}

func TestHeapRandomizedOrdering(t *testing.T) {
	random := rand.New(rand.NewSource(7))
	h := New[int](compare.Any[int])
	values := make([]int, 0, 2000)
	for i := 0; i < 2000; i++ {
		v := random.Intn(100000)
		values = append(values, v)
		h.Put(v)
	}
	sort.Ints(values)

	var got []int
	for {
		v, ok := h.Pop()
		if !ok {
			break
		}
		got = append(got, v)
	}
	if len(got) != len(values) {
		t.Fatalf("popped %d values want %d", len(got), len(values))
	}
	for i := range values {
		if got[i] != values[i] {
			t.Fatalf("heap order mismatch at %d: got %d want %d", i, got[i], values[i])
		}
	}
}

func TestHeapDescendingOrdering(t *testing.T) {
	h := New[int](compare.AnyDesc[int])
	for _, v := range []int{3, 1, 4, 1, 5, 9, 2, 6} {
		h.Put(v)
	}
	previous := 1 << 30
	count := 0
	for {
		v, ok := h.Pop()
		if !ok {
			break
		}
		if v > previous {
			t.Fatalf("descending heap yielded %d after %d", v, previous)
		}
		previous = v
		count++
	}
	if count != 8 {
		t.Fatalf("popped %d values want 8", count)
	}
}
