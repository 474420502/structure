package treequeue

import (
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
)

func TestPrioritySurface(t *testing.T) {
	q := New[int, string](compare.Any[int])
	if q.Len() != 0 || q.Size() != 0 {
		t.Fatal("new priority queue should be empty")
	}

	q.Put(3, "c1")
	q.Put(3, "c2")
	q.Put(1, "a")
	q.Put(2, "b")

	if q.Len() != 4 {
		t.Fatalf("Len = %d want 4", q.Len())
	}
	if _, ok := q.Get(3); !ok {
		t.Fatal("Get(3) should be present")
	}
	if got := q.Gets(3); len(got) != 2 {
		t.Fatalf("Gets(3) len = %d want 2", len(got))
	}
	if got := q.Gets(99); len(got) != 0 {
		t.Fatalf("Gets(99) len = %d want 0", len(got))
	}

	if q.InsertIfAbsent(3, "x") {
		t.Fatal("InsertIfAbsent must reject an existing key")
	}
	if !q.InsertIfAbsent(4, "d") {
		t.Fatal("InsertIfAbsent must insert a missing key")
	}
	if !q.Upsert(3, "c3") {
		t.Fatal("Upsert should replace the first matching value")
	}

	// Traverse must be non-decreasing by key and cover every node.
	var prev int = -1 << 30
	count := 0
	q.Traverse(func(k int, v string) bool {
		if k < prev {
			t.Fatalf("Traverse out of order: %d after %d", k, prev)
		}
		prev = k
		count++
		return true
	})
	if count != q.Len() {
		t.Fatalf("Traverse visited %d nodes, Len = %d", count, q.Len())
	}
	if len(q.Values()) != q.Len() {
		t.Fatalf("Values len = %d want %d", len(q.Values()), q.Len())
	}

	// Index and negative index agree with Values().
	values := q.Values()
	for i := 0; i < q.Len(); i++ {
		if got := q.Index(i); got != values[i] {
			t.Fatalf("Index(%d) = %v want %v", i, got, values[i])
		}
		if got := q.Index(i - q.Len()); got != values[i] {
			t.Fatalf("Index(%d) (negative) = %v want %v", i-q.Len(), got, values[i])
		}
	}

	// RemoveIndex removes by position.
	before := q.Len()
	if got := q.RemoveIndex(0); got != values[0] {
		t.Fatalf("RemoveIndex(0) = %v want %v", got, values[0])
	}
	if q.Len() != before-1 {
		t.Fatalf("Len after RemoveIndex = %d want %d", q.Len(), before-1)
	}

	// PopHead / PopTail pop the extremes.
	head, ok := q.PopHead()
	if !ok {
		t.Fatal("PopHead should succeed on non-empty queue")
	}
	tail, ok := q.PopTail()
	if !ok {
		t.Fatal("PopTail should succeed on non-empty queue")
	}
	if len(q.Values()) != q.Len() {
		t.Fatal("value count drift")
	}
	_ = head
	_ = tail

	q.Clear()
	if q.Len() != 0 {
		t.Fatal("Clear should empty the queue")
	}
	if _, ok := q.PopHead(); ok {
		t.Fatal("PopHead on empty queue should be false")
	}
	if _, ok := q.PopTail(); ok {
		t.Fatal("PopTail on empty queue should be false")
	}
}

func TestPriorityIterators(t *testing.T) {
	q := New[int, int](compare.Any[int])
	for i := 1; i <= 8; i++ {
		q.Put(i, i*10)
	}

	iter := q.Iterator()
	iter.SeekToFirst()
	for i := 1; i <= 8; i++ {
		if !iter.Valid() || iter.Key() != i || iter.Value() != i*10 {
			t.Fatalf("forward at %d: valid=%v key=%d", i, iter.Valid(), iter.Key())
		}
		iter.Next()
	}
	if iter.Valid() {
		t.Fatal("iterator should be invalid after final Next")
	}

	iter.SeekToLast()
	if iter.Key() != 8 {
		t.Fatalf("SeekToLast key = %d want 8", iter.Key())
	}
	iter.Prev()
	if iter.Key() != 7 {
		t.Fatalf("Prev key = %d want 7", iter.Key())
	}

	if !iter.SeekGEExact(5) || iter.Key() != 5 {
		t.Fatalf("SeekGEExact(5) key = %d", iter.Key())
	}
	if iter.SeekGEExact(50) {
		t.Fatal("SeekGEExact(50) should report no exact match")
	}
	if !iter.SeekLEExact(5) || iter.Key() != 5 {
		t.Fatalf("SeekLEExact(5) key = %d", iter.Key())
	}
	if iter.SeekLEExact(0) {
		t.Fatal("SeekLEExact(0) should report no exact match")
	}
	if !iter.SeekGTExact(5) || iter.Key() != 6 {
		t.Fatalf("SeekGTExact(5) key = %d", iter.Key())
	}
	if !iter.SeekLTExact(5) || iter.Key() != 4 {
		t.Fatalf("SeekLTExact(5) key = %d", iter.Key())
	}

	// The plain Seek* helpers position without returning the exact flag.
	iter.SeekGE(3)
	if !iter.Valid() || iter.Key() != 3 {
		t.Fatalf("SeekGE(3) key = %d", iter.Key())
	}
	iter.SeekLE(3)
	if !iter.Valid() || iter.Key() != 3 {
		t.Fatalf("SeekLE(3) key = %d", iter.Key())
	}
	iter.SeekGT(3)
	if !iter.Valid() || iter.Key() != 4 {
		t.Fatalf("SeekGT(3) key = %d", iter.Key())
	}
	iter.SeekLT(3)
	if !iter.Valid() || iter.Key() != 2 {
		t.Fatalf("SeekLT(3) key = %d", iter.Key())
	}

	iter.SeekGE(3)
	clone := iter.Clone()
	clone.Next()
	if iter.Key() != 3 || clone.Key() != 4 {
		t.Fatalf("Clone leaked state: orig=%d clone=%d", iter.Key(), clone.Key())
	}
	if !iter.Valid() {
		t.Fatal("Valid alias should agree with Vaild")
	}
	if iter.Vaild() != iter.Valid() {
		t.Fatal("Vaild and Valid disagree")
	}
}

func TestPrioritySliceAccessors(t *testing.T) {
	s := &Slice[int]{key: 7, value: "v"}
	if s.Key() != 7 {
		t.Fatalf("Key = %d want 7", s.Key())
	}
	if s.Value() != "v" {
		t.Fatalf("Value = %v want v", s.Value())
	}
	s.SetValue("w")
	if s.Value() != "w" {
		t.Fatalf("Value after SetValue = %v want w", s.Value())
	}
	if s.String() != "(7,w)" {
		t.Fatalf("String = %q want (7,w)", s.String())
	}
}

func TestPriorityRandomizedAgainstMultiset(t *testing.T) {
	random := rand.New(rand.NewSource(99))
	q := New[int, int](compare.Any[int])
	// This container is a multimap: Put always inserts, Delete removes one.
	counts := map[int]int{}
	total := 0

	for step := 0; step < 5000; step++ {
		key := random.Intn(128)
		if random.Intn(3) == 0 {
			_, ok := q.Delete(key)
			if ok != (counts[key] > 0) {
				t.Fatalf("step %d: Delete(%d) ok=%v count=%d", step, key, ok, counts[key])
			}
			if ok {
				counts[key]--
				total--
			}
		} else {
			q.Put(key, random.Intn(10000))
			counts[key]++
			total++
		}
	}

	if q.Len() != total {
		t.Fatalf("Len = %d want %d", q.Len(), total)
	}
	for k, c := range counts {
		if c <= 0 {
			continue
		}
		if _, ok := q.Get(k); !ok {
			t.Fatalf("Get(%d) missing for count %d", k, c)
		}
		if got := len(q.Gets(k)); got != c {
			t.Fatalf("Gets(%d) len = %d want %d", k, got, c)
		}
	}
}
