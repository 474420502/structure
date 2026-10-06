package orderedmap

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/474420502/structure/compare"
)

func TestOrderedMapOperations(t *testing.T) {
	m := New[int, string](compare.Any[int])

	if !m.IsEmpty() || m.Len() != 0 || m.Size() != 0 {
		t.Fatal("new map should be empty")
	}

	if !m.InsertIfAbsent(2, "two") {
		t.Fatal("InsertIfAbsent should insert missing key")
	}
	if m.InsertIfAbsent(2, "TWO") {
		t.Fatal("InsertIfAbsent should not overwrite")
	}
	if v, ok := m.Get(2); !ok || v != "two" {
		t.Fatalf("Get(2) = %q,%v want two,true", v, ok)
	}

	if v := m.Upsert(2, "TWO"); !v {
		t.Fatal("Upsert should replace existing key")
	}
	if v := m.Upsert(3, "three"); v {
		t.Fatal("Upsert should report false when inserting")
	}
	if v, _ := m.Get(2); v != "TWO" {
		t.Fatalf("Upsert value = %q want TWO", v)
	}

	if !m.Put(1, "one") {
		t.Fatal("Put should insert")
	}
	if m.Put(1, "ONE") {
		t.Fatal("Put should preserve existing value")
	}
	if v, _ := m.Get(1); v != "one" {
		t.Fatalf("Put overwrote value: %q", v)
	}

	if !m.Contains(1) || m.Contains(99) {
		t.Fatal("Contains mismatch")
	}

	if v, ok := m.Delete(2); !ok || v != "TWO" {
		t.Fatalf("Delete(2) = %q,%v", v, ok)
	}
	if _, ok := m.Delete(2); ok {
		t.Fatal("Delete missing should be false")
	}
	if _, ok := m.Remove(3); !ok {
		t.Fatal("Remove(3) should succeed")
	}

	if got := m.Keys(); !reflect.DeepEqual(got, []int{1}) {
		t.Fatalf("Keys = %v want [1]", got)
	}
	if got := m.Values(); !reflect.DeepEqual(got, []string{"one"}) {
		t.Fatalf("Values = %v want [one]", got)
	}

	m.Clear()
	if !m.IsEmpty() || m.Len() != 0 {
		t.Fatal("Clear should empty the map")
	}
}

func TestOrderedMapIndexOperations(t *testing.T) {
	m := New[string, int](compare.ArrayAny[string])
	for _, k := range []string{"c", "a", "b", "d"} {
		m.Set(k, len(k))
	}

	if got := m.IndexOf("a"); got != 0 {
		t.Fatalf("IndexOf(a) = %d want 0", got)
	}
	if got := m.IndexOf("d"); got != 3 {
		t.Fatalf("IndexOf(d) = %d want 3", got)
	}
	if got := m.IndexOf("z"); got >= 0 {
		t.Fatalf("IndexOf(missing) = %d want negative", got)
	}

	if k, _ := m.Index(1); k != "b" {
		t.Fatalf("Index(1).key = %q want b", k)
	}

	k, v, ok := m.RemoveIndex(0)
	if !ok || k != "a" || v != 1 {
		t.Fatalf("RemoveIndex(0) = %q,%d,%v", k, v, ok)
	}
	if m.Len() != 3 {
		t.Fatalf("Len after RemoveIndex = %d want 3", m.Len())
	}
	if _, _, ok := m.RemoveIndex(99); ok {
		t.Fatal("RemoveIndex out of range should be false")
	}
	if _, _, ok := m.RemoveIndex(-1); ok {
		t.Fatal("RemoveIndex negative should be false")
	}
}

func TestOrderedMapIteratorNavigation(t *testing.T) {
	m := New[int, int](compare.Any[int])
	for i := 0; i < 10; i++ {
		m.Set(i, i*10)
	}

	iter := m.Iterator()
	iter.SeekToFirst()
	if !iter.Valid() || iter.Key() != 0 || iter.Value() != 0 {
		t.Fatalf("SeekToFirst landed on %d", iter.Key())
	}
	for i := 0; i < 10; i++ {
		if !iter.Valid() || iter.Key() != i || iter.Index() != int64(i) {
			t.Fatalf("forward step %d: key=%d index=%d", i, iter.Key(), iter.Index())
		}
		iter.Next()
	}
	if iter.Valid() {
		t.Fatal("iterator should be invalid after final Next")
	}

	iter.SeekToLast()
	if !iter.Valid() || iter.Key() != 9 {
		t.Fatalf("SeekToLast landed on %d", iter.Key())
	}
	iter.Prev()
	if !iter.Valid() || iter.Key() != 8 {
		t.Fatalf("Prev landed on %d", iter.Key())
	}

	if exact := iter.SeekGE(5); !exact || iter.Key() != 5 {
		t.Fatalf("SeekGE(5) exact=%v key=%d", exact, iter.Key())
	}
	if exact := iter.SeekGE(15); exact || iter.Valid() {
		t.Fatalf("SeekGE(15) should be inexact and invalid, exact=%v valid=%v", exact, iter.Valid())
	}
	if exact := iter.SeekLE(5); !exact || iter.Key() != 5 {
		t.Fatalf("SeekLE(5) exact=%v key=%d", exact, iter.Key())
	}
	if exact := iter.SeekLE(-1); exact || iter.Valid() {
		t.Fatalf("SeekLE(-1) should be inexact and invalid")
	}
	if exact := iter.SeekGT(5); !exact || iter.Key() != 6 {
		t.Fatalf("SeekGT(5) exact=%v key=%d want true,6", exact, iter.Key())
	}
	if exact := iter.SeekLT(5); !exact || iter.Key() != 4 {
		t.Fatalf("SeekLT(5) exact=%v key=%d want true,4", exact, iter.Key())
	}

	iter.SeekByIndex(3)
	if !iter.Valid() || iter.Key() != 3 {
		t.Fatalf("SeekByIndex(3) key=%d", iter.Key())
	}

	clone := iter.Clone()
	clone.Next()
	if iter.Key() != 3 || clone.Key() != 4 {
		t.Fatalf("Clone should not affect original: orig=%d clone=%d", iter.Key(), clone.Key())
	}

	empty := New[int, int](compare.Any[int]).Iterator()
	empty.SeekToFirst()
	if empty.Valid() {
		t.Fatal("empty iterator should be invalid")
	}
}

func TestOrderedMapRandomizedAgainstMap(t *testing.T) {
	random := rand.New(rand.NewSource(20240517))
	m := New[int, int](compare.Any[int])
	expected := map[int]int{}

	for step := 0; step < 20000; step++ {
		key := random.Intn(512)
		value := random.Intn(100000)
		switch random.Intn(5) {
		case 0:
			if m.InsertIfAbsent(key, value) {
				if _, exists := expected[key]; !exists {
					expected[key] = value
				}
			}
		case 1:
			m.Set(key, value)
			expected[key] = value
		case 2:
			m.Upsert(key, value)
			expected[key] = value
		case 3:
			_, ok := m.Delete(key)
			_, want := expected[key]
			if ok != want {
				t.Fatalf("step %d: Delete(%d) ok=%v want %v", step, key, ok, want)
			}
			delete(expected, key)
		default:
			got, ok := m.Get(key)
			want, exists := expected[key]
			if ok != exists || (ok && got != want) {
				t.Fatalf("step %d: Get(%d)=%d,%v want %d,%v", step, key, got, ok, want, exists)
			}
		}
	}

	keys := make([]int, 0, len(expected))
	for k := range expected {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	if m.Len() != len(expected) {
		t.Fatalf("Len = %d want %d", m.Len(), len(expected))
	}
	if got := m.Keys(); !reflect.DeepEqual(got, keys) {
		t.Fatalf("Keys ordering mismatch: got %d keys want %d", len(got), len(keys))
	}
}

func TestOrderedMapStringKeys(t *testing.T) {
	m := New[string, int](compare.ArrayAny[string])
	for _, k := range []string{"b", "a", "aa", "ab"} {
		m.Set(k, len(k))
	}
	want := []string{"a", "aa", "ab", "b"}
	if got := m.Keys(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Keys = %v want %v", got, want)
	}
}
