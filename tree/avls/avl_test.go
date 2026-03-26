package avls

import (
	"log"
	"sort"
	"testing"

	random "github.com/474420502/random"
	"github.com/474420502/structure/compare"
	testutils "github.com/474420502/structure/tree/test_utils"
)

type duplicateEntry struct {
	key   int
	value int
	id    int
}

func TestPut(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for _, i := range []int{5, 8, 11, 10, 11, 50, 1, 99} {
		tree.Put(i, i)
	}

	if int(tree.Size()) != 8 {
		t.Fatalf("size mismatch: %v", tree.Values())
	}

	values := tree.Values()
	if values[4] != 11 || values[5] != 11 {
		t.Fatalf("duplicate key should be preserved: %v", values)
	}

	log.Println(tree.view())
}

func TestPutDuplicateLifecycle(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for _, value := range []int{5, 5, 5, 3, 7} {
		tree.Put(value, value)
	}

	if int(tree.Size()) != 5 {
		t.Fatalf("size mismatch: %v", tree.Values())
	}
 
	values := tree.Values()
	if got := len(values); got != 5 {
		t.Fatalf("unexpected value count: %d %v", got, values)
	}

	count5 := 0
	for _, value := range values {
		if value == 5 {
			count5++
		}
	}
	if count5 != 3 {
		t.Fatalf("duplicate count mismatch: %v", values)
	}

	for remaining := 2; remaining >= 0; remaining-- {
		if value, ok := tree.Remove(5); !ok || value != 5 {
			t.Fatalf("remove duplicate failed: ok=%v value=%v", ok, value)
		}
		if int(tree.Size()) != remaining+2 {
			t.Fatalf("size after remove mismatch: %d %v", tree.Size(), tree.Values())
		}
		_, ok := tree.Get(5)
		if remaining == 0 {
			if ok {
				t.Fatalf("duplicate key should be exhausted: %v", tree.Values())
			}
		} else if !ok {
			t.Fatalf("duplicate key should still exist: %v", tree.Values())
		}
	}

	if _, ok := tree.Remove(5); ok {
		t.Fatalf("remove should fail after all duplicates are removed: %v", tree.Values())
	}

	
}

func TestSetDoesNotDuplicateKey(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	if !tree.Set(5, 50) {
		t.Fatal("first set should insert")
	}
	if tree.Set(5, 55) {
		t.Fatal("second set should update existing key")
	}

	if int(tree.Size()) != 1 {
		t.Fatalf("set should not duplicate key: %v", tree.Values())
	}

	value, ok := tree.Get(5)
	if !ok || value != 55 {
		t.Fatalf("set should update value: ok=%v value=%v", ok, value)
	}
}

func TestRemoveDuplicateKeyWithDistinctValues(t *testing.T) {
	tree := New[int, string](compare.AnyEx[int])
	for _, item := range []struct {
		key   int
		value string
	}{
		{5, "alpha"},
		{5, "beta"},
		{5, "gamma"},
		{3, "left"},
		{7, "right"},
	} {
		tree.Put(item.key, item.value)
	}

	if int(tree.Size()) != 5 {
		t.Fatalf("unexpected size: %d %v", tree.Size(), tree.Values())
	}

	removed := make(map[string]int)
	for i := 0; i < 3; i++ {
		value, ok := tree.Remove(5)
		if !ok {
			t.Fatalf("expected duplicate key to be removable at step %d: %v", i, tree.Values())
		}
		removed[value]++
	}

	if len(removed) != 3 {
		t.Fatalf("expected all duplicate values to be removable exactly once: %#v", removed)
	}
	for _, expected := range []string{"alpha", "beta", "gamma"} {
		if removed[expected] != 1 {
			t.Fatalf("missing removed value %q: %#v", expected, removed)
		}
	}

	if int(tree.Size()) != 2 {
		t.Fatalf("unexpected size after duplicate removals: %d %v", tree.Size(), tree.Values())
	}
	if _, ok := tree.Get(5); ok {
		t.Fatalf("key should be exhausted after removing all duplicates: %v", tree.Values())
	}

	values := tree.Values()
	if got := len(values); got != 2 {
		t.Fatalf("unexpected remaining values: %v", values)
	}
	remaining := map[string]bool{}
	for _, value := range values {
		remaining[value] = true
	}
	if !remaining["left"] || !remaining["right"] {
		t.Fatalf("non-duplicate neighbors should remain intact: %v", values)
	}
}

func TestSetWithExistingDuplicatesUpdatesSingleNode(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for _, value := range []int{10, 20, 30} {
		tree.Put(5, value)
	}

	if tree.Set(5, 99) {
		t.Fatal("set should update an existing duplicate key instead of inserting")
	}

	if int(tree.Size()) != 3 {
		t.Fatalf("set should not increase size when duplicates already exist: %d %v", tree.Size(), tree.Values())
	}

	values := tree.Values()
	counts := map[int]int{}
	for _, value := range values {
		counts[value]++
	}
	if counts[99] != 1 || counts[10]+counts[20]+counts[30] != 2 {
		t.Fatalf("set should update exactly one duplicate node: %v", values)
	}

	for i := 0; i < 3; i++ {
		if _, ok := tree.Remove(5); !ok {
			t.Fatalf("expected duplicate key to remain removable at step %d: %v", i, tree.Values())
		}
	}
	if _, ok := tree.Remove(5); ok {
		t.Fatalf("no duplicate key should remain after removing all of them: %v", tree.Values())
	}
}

func TestDuplicateKeyGetSetRemoveStability(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for _, item := range []struct {
		key   int
		value int
	}{
		{4, 400},
		{5, 100},
		{5, 200},
		{5, 300},
		{6, 600},
	} {
		tree.Put(item.key, item.value)
	}

	if value, ok := tree.Get(5); !ok || value != 100 {
		t.Fatalf("get should return the oldest remaining duplicate: ok=%v value=%v all=%v", ok, value, tree.Values())
	}

	if tree.Set(5, 999) {
		t.Fatal("set should update an existing duplicate key")
	}

	if int(tree.Size()) != 5 {
		t.Fatalf("set should not change size: %d %v", tree.Size(), tree.Values())
	}

	values := tree.Values()
	expected := []int{400, 999, 200, 300, 600}
	for i, value := range expected {
		if values[i] != value {
			t.Fatalf("unexpected values after set: %v", values)
		}
	}

	if value, ok := tree.Get(5); !ok || value != 999 {
		t.Fatalf("get should track the same duplicate node that set updated: ok=%v value=%v all=%v", ok, value, tree.Values())
	}

	for _, expectedValue := range []int{999, 200, 300} {
		value, ok := tree.Remove(5)
		if !ok || value != expectedValue {
			t.Fatalf("remove should consume duplicates from the same frontier as get: want=%d got=(%v,%v) all=%v", expectedValue, ok, value, tree.Values())
		}
	}

	if _, ok := tree.Get(5); ok {
		t.Fatalf("all duplicates should be exhausted: %v", tree.Values())
	}
}

func TestMultisetReferenceForce(t *testing.T) {
	rand := random.New(t.Name())

	for round := 0; round < 300; round++ {
		tree := New[int, int](compare.AnyEx[int])
		var reference []duplicateEntry
		nextID := 1

		for step := 0; step < 150; step++ {
			key := rand.Intn(12)
			value := rand.Intn(1000)

			switch rand.Intn(4) {
			case 0:
				tree.Put(key, value)
				reference = append(reference, duplicateEntry{key: key, value: value, id: nextID})
				nextID++
			case 1:
				inserted := tree.Set(key, value)
				if idx := firstReferenceIndex(reference, key); idx >= 0 {
					if inserted {
						t.Fatalf("set should update existing key: key=%d ref=%v", key, reference)
					}
					reference[idx].value = value
				} else {
					if !inserted {
						t.Fatalf("set should insert missing key: key=%d ref=%v", key, reference)
					}
					reference = append(reference, duplicateEntry{key: key, value: value, id: nextID})
					nextID++
				}
			case 2:
				gotValue, gotOK := tree.Get(key)
				if idx := firstReferenceIndex(reference, key); idx >= 0 {
					if !gotOK || gotValue != reference[idx].value {
						t.Fatalf("get mismatch: key=%d got=(%v,%v) want=%v ref=%v values=%v", key, gotOK, gotValue, reference[idx].value, reference, tree.Values())
					}
				} else if gotOK {
					t.Fatalf("get should miss: key=%d got=%v ref=%v values=%v", key, gotValue, reference, tree.Values())
				}
			case 3:
				gotValue, gotOK := tree.Remove(key)
				if idx := firstReferenceIndex(reference, key); idx >= 0 {
					wantValue := reference[idx].value
					reference = append(reference[:idx], reference[idx+1:]...)
					if !gotOK || gotValue != wantValue {
						t.Fatalf("remove mismatch: key=%d got=(%v,%v) want=%v ref=%v values=%v", key, gotOK, gotValue, wantValue, reference, tree.Values())
					}
				} else if gotOK {
					t.Fatalf("remove should miss: key=%d got=%v ref=%v values=%v", key, gotValue, reference, tree.Values())
				}
			}

			expectedValues := referenceValues(reference)
			actualValues := tree.Values()
			if len(actualValues) != len(expectedValues) {
				t.Fatalf("size mismatch: got=%d want=%d ref=%v values=%v", len(actualValues), len(expectedValues), reference, actualValues)
			}
			for i, value := range expectedValues {
				if actualValues[i] != value {
					t.Fatalf("values mismatch at %d: got=%v want=%v ref=%v values=%v", i, actualValues, expectedValues, reference, actualValues)
				}
			}
			if int(tree.Size()) != len(reference) {
				t.Fatalf("reported size mismatch: got=%d want=%d ref=%v values=%v", tree.Size(), len(reference), reference, actualValues)
			}
		}
	}
}

func TestPutGet(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for i := 0; i < 100; i++ {
		tree.Set(i, i)
	}

	log.Println(tree.view())

	for i := 0; i < int(tree.Size()); i++ {
		if v, b := tree.Get(i); !b || v != i {
			t.Error("error", b, v)
			log.Println(b, v)
		}
	}

	tree.Clear()
	for _, i := range testutils.TestedArray {
		tree.Set(i, i)
	}

	if int(tree.Size()) != len(testutils.TestedArray) {
		t.Error(tree.Values())
		log.Println(tree.view())
	}

	vs := tree.Values()
	if vs[0] != 1 || vs[int(tree.Size())-1] != 99 {
		t.Error(tree.Values())
		log.Println(tree.view())
		log.Println(tree.Values())
	}
}

func TestRemove2(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for _, i := range testutils.TestedBigArray {
		if !tree.Set(i, i) {
			// log.Println("equal key", i)
		}
	}

	if int(tree.Size()) != len(testutils.TestedBigArray)-4 {
		t.Error(int(tree.Size()), tree.Values())
	}

	for _, v := range tree.Values() {
		tree.Remove(v)
	}

	if int(tree.Size()) != 0 {
		t.Error(tree.Values())
	}
}

func TestRemove1(t *testing.T) {
	tree := New[int, int](compare.AnyEx[int])
	for _, i := range testutils.TestedArray {
		if !tree.Set(i, i) {
			// log.Println("equal key", i)
		}
	}

	if int(int(tree.Size())) != len(testutils.TestedArray) {
		t.Error(int(tree.Size()), tree.Values())
	}

	// log.Println(tree.debugString())
	for _, v := range tree.Values() {
		tree.Remove(v)
	}

	if int(tree.Size()) != 0 {
		t.Error(tree.Values())
	}
}

func TestForce(t *testing.T) {
	rand := random.New(t.Name())

	tree := New[int, int](compare.AnyEx[int])
	for n := 0; n < 2000; n++ {

		var priority []int
		for i := 0; i < 100; i++ {
			v := rand.Intn(100)
			tree.Put(v, v)
			priority = append(priority, v)
		}

		if int(tree.Size()) != len(priority) {
			panic("")
		}

		sort.Slice(priority, func(i, j int) bool {
			return priority[i] < priority[j]
		})

		for i, v := range tree.Values() {
			if priority[i] != v {
				panic("")
			}
		}

		for i := 0; i < 40; i++ {

			v := rand.Intn(100)

			if _, ok := tree.Get(v); ok {

				rv, ok := tree.Remove(v)
				if !ok || rv != v {
					panic("")
				}

				if idx := sort.SearchInts(priority, v); idx == len(priority) {
					panic("")
				} else {
					priority = append(priority[:idx], priority[idx+1:]...)
				}

			}
		}

		var i = 0
		tree.Traverse(func(k int, v int) bool {
			if priority[i] != v {
				panic("")
			}
			i++
			return true
		})

		tree.Clear()

	}
}

// func TestCaseX(t *testing.T) {
// 	New[int,int](compare.TimeDescEx[time.Time])
// }

func BenchmarkPut(b *testing.B) {
	rand := random.New(1683721792150515321)

	tree := New[int, int](compare.AnyEx[int])
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		v := rand.Int()
		tree.Put(v, v)
		// tree.check()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v := rand.Int()
		tree.Put(v, v)
	}
	// b.Log(tree.rotateCount)
}

func BenchmarkRemove(b *testing.B) {
	rand := random.New(1683721792150515321)
	tree := New[int, int](compare.AnyEx[int])
	var removelist []int
	var ri = 0

	for i := 0; i < b.N; i++ {
		if tree.Size() == 0 {
			b.StopTimer()
			removelist = nil
			ri = 0
			for i := 0; i < 100; i++ {
				v := rand.Intn(100)
				if tree.Put(v, v) {
					removelist = append(removelist, v)
				}

				if i%25 == 0 {
					removelist = append(removelist, rand.Intn(100))
				}
				// tree.check()
			}
			b.StartTimer()
		}

		v := removelist[ri]
		tree.Remove(v)
		ri += 1
	}

}

func firstReferenceIndex(reference []duplicateEntry, key int) int {
	sorted := sortedReference(reference)
	for _, entry := range sorted {
		if entry.key == key {
			for idx := range reference {
				if reference[idx].id == entry.id {
					return idx
				}
			}
		}
	}
	return -1
}

func referenceValues(reference []duplicateEntry) []int {
	sorted := sortedReference(reference)
	values := make([]int, 0, len(sorted))
	for _, entry := range sorted {
		values = append(values, entry.value)
	}
	return values
}

func sortedReference(reference []duplicateEntry) []duplicateEntry {
	sorted := append([]duplicateEntry(nil), reference...)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].key < sorted[j].key
	})
	return sorted
}
