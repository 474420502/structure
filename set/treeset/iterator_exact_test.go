package treeset

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestIteratorSeekExactAliases(t *testing.T) {
	set := New[int, int](compare.AnyEx[int])
	for _, value := range []int{2, 4, 6, 8} {
		set.Set(value, value)
	}

	iter := set.Iterator()
	if !iter.SeekGEExact(6) || !iter.Valid() || iter.Value() != 6 {
		t.Fatalf("SeekGEExact(6) should exact-match key 6, got valid=%v value=%v", iter.Valid(), func() any { if iter.Valid() { return iter.Value() }; return nil }())
	}
	if iter.SeekGEExact(5) || !iter.Valid() || iter.Value() != 6 {
		t.Fatalf("SeekGEExact(5) should return false and land on 6, got valid=%v value=%v", iter.Valid(), func() any { if iter.Valid() { return iter.Value() }; return nil }())
	}
	if iter.SeekGEExact(9) || iter.Valid() {
		t.Fatalf("SeekGEExact(9) should return false and invalidate iterator, got valid=%v", iter.Valid())
	}

	if !iter.SeekLEExact(6) || !iter.Valid() || iter.Value() != 6 {
		t.Fatalf("SeekLEExact(6) should exact-match key 6, got valid=%v value=%v", iter.Valid(), func() any { if iter.Valid() { return iter.Value() }; return nil }())
	}
	if iter.SeekLEExact(5) || !iter.Valid() || iter.Value() != 4 {
		t.Fatalf("SeekLEExact(5) should return false and land on 4, got valid=%v value=%v", iter.Valid(), func() any { if iter.Valid() { return iter.Value() }; return nil }())
	}
	if iter.SeekLEExact(1) || iter.Valid() {
		t.Fatalf("SeekLEExact(1) should return false and invalidate iterator, got valid=%v", iter.Valid())
	}

	if !iter.SeekGTExact(6) || !iter.Valid() || iter.Value() != 8 {
		t.Fatalf("SeekGTExact(6) should report exact key and land on 8, got valid=%v value=%v", iter.Valid(), func() any { if iter.Valid() { return iter.Value() }; return nil }())
	}
	if iter.SeekGTExact(5) || !iter.Valid() || iter.Value() != 6 {
		t.Fatalf("SeekGTExact(5) should return false and land on 6, got valid=%v value=%v", iter.Valid(), func() any { if iter.Valid() { return iter.Value() }; return nil }())
	}
	if !iter.SeekGTExact(8) || iter.Valid() {
		t.Fatalf("SeekGTExact(8) should report exact key and invalidate iterator after end, got valid=%v", iter.Valid())
	}

	if !iter.SeekLTExact(6) || !iter.Valid() || iter.Value() != 4 {
		t.Fatalf("SeekLTExact(6) should report exact key and land on 4, got valid=%v value=%v", iter.Valid(), func() any { if iter.Valid() { return iter.Value() }; return nil }())
	}
	if iter.SeekLTExact(5) || !iter.Valid() || iter.Value() != 4 {
		t.Fatalf("SeekLTExact(5) should return false and land on 4, got valid=%v value=%v", iter.Valid(), func() any { if iter.Valid() { return iter.Value() }; return nil }())
	}
	if !iter.SeekLTExact(2) || iter.Valid() {
		t.Fatalf("SeekLTExact(2) should report exact key and invalidate iterator before start, got valid=%v", iter.Valid())
	}
}