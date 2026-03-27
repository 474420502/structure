package skiplist

import (
	"math/rand"
	"sort"
	"sync"
	"testing"

	"github.com/474420502/structure/compare"
)

func TestBasicPutAndGet(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 100; i++ {
		if !tree.Put(i, i*10) {
			t.Errorf("Put failed for key %d", i)
		}
	}

	if tree.Size() != 100 {
		t.Errorf("expected size 100, got %d", tree.Size())
	}

	for i := 0; i < 100; i++ {
		v, ok := tree.Get(i)
		if !ok {
			t.Errorf("Get failed for key %d", i)
		}
		if v != i*10 {
			t.Errorf("expected value %d, got %d", i*10, v)
		}
	}

	_, ok := tree.Get(999)
	if ok {
		t.Error("Get should return false for non-existent key")
	}
}

func TestPutDuplicate(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	tree.Put(1, 100)
	if tree.Put(1, 200) {
		t.Error("PutDuplicate should return false when updating existing key")
	}

	v, _ := tree.Get(1)
	if v != 200 {
		t.Errorf("expected value 200, got %d", v)
	}
}

func TestPutDuplicateWithCallback(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	tree.Put(1, 100)
	var oldValue int
	tree.PutDuplicate(1, 200, func(s *Slice[int, int]) {
		oldValue = s.Value
	})

	if oldValue != 100 {
		t.Errorf("callback should receive old value 100, got %d", oldValue)
	}

	v, _ := tree.Get(1)
	if v != 200 {
		t.Errorf("expected value 200, got %d", v)
	}
}

func TestRemove(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	for i := 0; i < 100; i++ {
		removed := tree.Remove(i)
		if removed == nil {
			t.Errorf("Remove failed for key %d", i)
		}
		if removed.Key != i || removed.Value != i {
			t.Errorf("Remove returned wrong value for key %d", i)
		}
	}

	if tree.Size() != 0 {
		t.Errorf("expected size 0 after remove, got %d", tree.Size())
	}

	removed := tree.Remove(999)
	if removed != nil {
		t.Error("Remove should return nil for non-existent key")
	}
}

func TestRemoveHead(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	tree.Put(5, 50)
	tree.Put(3, 30)
	tree.Put(7, 70)

	removed := tree.RemoveHead()
	if removed == nil || removed.Key != 3 {
		t.Errorf("RemoveHead should return smallest key, got %v", removed)
	}

	removed = tree.RemoveHead()
	if removed == nil || removed.Key != 5 {
		t.Errorf("RemoveHead should return next smallest key, got %v", removed)
	}

	removed = tree.RemoveHead()
	if removed == nil || removed.Key != 7 {
		t.Errorf("RemoveHead should return last key, got %v", removed)
	}

	if tree.Size() != 0 {
		t.Errorf("expected size 0, got %d", tree.Size())
	}

	removed = tree.RemoveHead()
	if removed != nil {
		t.Error("RemoveHead on empty tree should return nil")
	}
}

func TestRemoveTail(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	tree.Put(5, 50)
	tree.Put(3, 30)
	tree.Put(7, 70)

	removed := tree.RemoveTail()
	if removed == nil || removed.Key != 7 {
		t.Errorf("RemoveTail should return largest key, got %v", removed)
	}

	removed = tree.RemoveTail()
	if removed == nil || removed.Key != 5 {
		t.Errorf("RemoveTail should return next largest key, got %v", removed)
	}

	removed = tree.RemoveTail()
	if removed == nil || removed.Key != 3 {
		t.Errorf("RemoveTail should return last key, got %v", removed)
	}

	if tree.Size() != 0 {
		t.Errorf("expected size 0, got %d", tree.Size())
	}

	removed = tree.RemoveTail()
	if removed != nil {
		t.Error("RemoveTail on empty tree should return nil")
	}
}

func TestRemoveIndex(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	removed := tree.RemoveIndex(0)
	if removed == nil || removed.Key != 0 {
		t.Errorf("RemoveIndex(0) should remove key 0, got %v", removed)
	}

	removed = tree.RemoveIndex(8)
	if removed == nil || removed.Key != 9 {
		t.Errorf("RemoveIndex(8) should remove key 9, got %v", removed)
	}

	removed = tree.RemoveIndex(-1)
	if removed != nil {
		t.Error("RemoveIndex(-1) should return nil")
	}

	removed = tree.RemoveIndex(100)
	if removed != nil {
		t.Error("RemoveIndex(100) should return nil")
	}
}

func TestRemoveRangeByIndex(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	tree.RemoveRangeByIndex(2, 5)
	if tree.Size() != 6 {
		t.Errorf("expected size 6, got %d", tree.Size())
	}

	for i := 2; i <= 5; i++ {
		if _, ok := tree.Get(i); ok {
			t.Errorf("key %d should be removed", i)
		}
	}

	tree.RemoveRangeByIndex(-1, 100)
	if tree.Size() != 0 {
		t.Errorf("expected size 0, got %d", tree.Size())
	}
}

func TestClear(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	tree.Clear()

	if tree.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", tree.Size())
	}

	_, ok := tree.Get(50)
	if ok {
		t.Error("Get should return false after clear")
	}
}

func TestHead(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	head := tree.Head()
	if head != nil {
		t.Error("Head on empty tree should return nil")
	}

	tree.Put(5, 50)
	tree.Put(3, 30)
	tree.Put(7, 70)

	head = tree.Head()
	if head == nil || head.Key != 3 {
		t.Errorf("Head should return smallest key, got %v", head)
	}
}

func TestTail(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	tail := tree.Tail()
	if tail != nil {
		t.Error("Tail on empty tree should return nil")
	}

	tree.Put(5, 50)
	tree.Put(3, 30)
	tree.Put(7, 70)

	tail = tree.Tail()
	if tail == nil || tail.Key != 7 {
		t.Errorf("Tail should return largest key, got %v", tail)
	}
}

func TestIndex(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 100; i++ {
		tree.Put(i, i*10)
	}

	for i := 0; i < 100; i++ {
		s := tree.Index(int64(i))
		if s == nil || s.Key != i {
			t.Errorf("Index(%d) should return key %d, got %v", i, i, s)
		}
	}

	s := tree.Index(-1)
	if s != nil {
		t.Error("Index(-1) should return nil")
	}

	s = tree.Index(100)
	if s != nil {
		t.Error("Index(100) should return nil")
	}
}

func TestIndexOf(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	for i := 0; i < 100; i++ {
		idx := tree.IndexOf(i)
		if idx != int64(i) {
			t.Errorf("IndexOf(%d) should return %d, got %d", i, i, idx)
		}
	}

	idx := tree.IndexOf(999)
	if idx != -1 {
		t.Errorf("IndexOf(999) should return -1, got %d", idx)
	}
}

func TestHeight(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	if tree.Height() < 1 {
		t.Error("Height should be at least 1")
	}

	for i := 0; i < 1000; i++ {
		tree.Put(i, i)
	}

	h := tree.Height()
	if h < 1 || h > MaxLevel {
		t.Errorf("Height should be between 1 and %d, got %d", MaxLevel, h)
	}
}

func TestTraverse(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	var result []int
	tree.Traverse(func(s *Slice[int, int]) bool {
		result = append(result, s.Key)
		return true
	})

	if len(result) != 10 {
		t.Errorf("Traverse should visit 10 nodes, got %d", len(result))
	}

	for i := 0; i < 10; i++ {
		if result[i] != i {
			t.Errorf("Traverse order wrong: expected %d at index %d, got %d", i, i, result[i])
		}
	}
}

func TestTraverseBreak(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	var count int
	tree.Traverse(func(s *Slice[int, int]) bool {
		count++
		if s.Key == 5 {
			return false
		}
		return true
	})

	if count != 6 {
		t.Errorf("Traverse should stop after 6 visits, got %d", count)
	}
}

func TestSlice(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i*10)
	}

	slices := tree.Slice()

	if len(slices) != 10 {
		t.Errorf("Slice should have 10 elements, got %d", len(slices))
	}

	for i := 0; i < 10; i++ {
		if slices[i].Key != i || slices[i].Value != i*10 {
			t.Errorf("Slice[%d] = {%d, %d}, expected {%d, %d}", i, slices[i].Key, slices[i].Value, i, i*10)
		}
	}
}

func TestIterator(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	iter := tree.Iterator()
	var result []int

	for iter.Valid() {
		result = append(result, iter.Key())
		iter.Next()
	}

	if len(result) != 10 {
		t.Errorf("Iterator should traverse 10 nodes, got %d", len(result))
	}

	for i := 0; i < 10; i++ {
		if result[i] != i {
			t.Errorf("Iterator order wrong: expected %d at index %d, got %d", i, i, result[i])
		}
	}
}

func TestIteratorSeekToFirst(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 5; i < 10; i++ {
		tree.Put(i, i)
	}

	iter := tree.Iterator()
	iter.SeekToFirst()

	if !iter.Valid() || iter.Key() != 5 {
		t.Errorf("SeekToFirst should position at key 5, got %v", iter.Key())
	}
}

func TestIteratorSeekToLast(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	iter := tree.Iterator()
	iter.SeekToLast()

	if !iter.Valid() || iter.Key() != 9 {
		t.Errorf("SeekToLast should position at key 9, got %v", iter.Key())
	}
}

func TestIteratorPrev(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	iter := tree.Iterator()
	iter.SeekToLast()

	var result []int
	for iter.Valid() {
		result = append(result, iter.Key())
		iter.Prev()
	}

	if len(result) != 10 {
		t.Errorf("Iterator.Prev should traverse 10 nodes, got %d", len(result))
	}

	for i := 0; i < 10; i++ {
		if result[i] != 9-i {
			t.Errorf("Iterator.Prev order wrong: expected %d at index %d, got %d", 9-i, i, result[i])
		}
	}
}

func TestIteratorSeekGE(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		tree.Put(i, i)
	}

	iter := tree.Iterator()

	if !iter.SeekGE(5) {
		t.Error("SeekGE(5) should find key 5")
	}
	if iter.Key() != 5 {
		t.Errorf("SeekGE(5) should return key 5, got %d", iter.Key())
	}

	if !iter.SeekGE(6) {
		t.Error("SeekGE(6) should find key 7")
	}
	if iter.Key() != 7 {
		t.Errorf("SeekGE(6) should return key 7, got %d", iter.Key())
	}

	if !iter.SeekGE(9) {
		t.Error("SeekGE(9) should find key 9")
	}
	if iter.Key() != 9 {
		t.Errorf("SeekGE(9) should return key 9, got %d", iter.Key())
	}

	if iter.SeekGE(10) {
		t.Error("SeekGE(10) should not find any key")
	}
}

func TestIteratorSeekGT(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		tree.Put(i, i)
	}

	iter := tree.Iterator()

	if !iter.SeekGT(5) {
		t.Error("SeekGT(5) should find key 7")
	}
	if iter.Key() != 7 {
		t.Errorf("SeekGT(5) should return key 7, got %d", iter.Key())
	}

	if !iter.SeekGT(6) {
		t.Error("SeekGT(6) should find key 7")
	}
	if iter.Key() != 7 {
		t.Errorf("SeekGT(6) should return key 7, got %d", iter.Key())
	}

	if !iter.SeekGT(7) {
		t.Error("SeekGT(7) should find key 9")
	}
	if iter.Key() != 9 {
		t.Errorf("SeekGT(7) should return key 9, got %d", iter.Key())
	}

	if iter.SeekGT(9) {
		t.Error("SeekGT(9) should not find any key")
	}
}

func TestIteratorSeekLE(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		tree.Put(i, i)
	}

	iter := tree.Iterator()

	if !iter.SeekLE(5) {
		t.Error("SeekLE(5) should find key 5")
	}
	if iter.Key() != 5 {
		t.Errorf("SeekLE(5) should return key 5, got %d", iter.Key())
	}

	if !iter.SeekLE(6) {
		t.Error("SeekLE(6) should find key 5")
	}
	if iter.Key() != 5 {
		t.Errorf("SeekLE(6) should return key 5, got %d", iter.Key())
	}

	if !iter.SeekLE(9) {
		t.Error("SeekLE(9) should find key 9")
	}
	if iter.Key() != 9 {
		t.Errorf("SeekLE(9) should return key 9, got %d", iter.Key())
	}

	if !iter.SeekLE(100) {
		t.Error("SeekLE(100) should find key 9")
	}
	if iter.Key() != 9 {
		t.Errorf("SeekLE(100) should return key 9, got %d", iter.Key())
	}
}

func TestIteratorSeekLT(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		tree.Put(i, i)
	}

	iter := tree.Iterator()

	if !iter.SeekLT(5) {
		t.Error("SeekLT(5) should find key 3")
	}
	if iter.Key() != 3 {
		t.Errorf("SeekLT(5) should return key 3, got %d", iter.Key())
	}

	if !iter.SeekLT(6) {
		t.Error("SeekLT(6) should find key 5")
	}
	if iter.Key() != 5 {
		t.Errorf("SeekLT(6) should return key 5, got %d", iter.Key())
	}

	if !iter.SeekLT(7) {
		t.Error("SeekLT(7) should find key 5")
	}
	if iter.Key() != 5 {
		t.Errorf("SeekLT(7) should return key 5, got %d", iter.Key())
	}

	if iter.SeekLT(1) {
		t.Error("SeekLT(1) should not find any key")
	}
}

func TestIteratorSeekByIndex(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	iter := tree.Iterator()

	iter.SeekByIndex(5)
	if !iter.Valid() || iter.Key() != 5 {
		t.Errorf("SeekByIndex(5) should position at key 5, got %v", iter.Key())
	}

	iter.SeekByIndex(-1)
	if iter.Valid() {
		t.Error("SeekByIndex(-1) should not be valid")
	}

	iter.SeekByIndex(10)
	if iter.Valid() {
		t.Error("SeekByIndex(10) should not be valid")
	}
}

func TestIteratorClone(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	iter1 := tree.Iterator()
	iter1.SeekToFirst()
	iter1.Next()
	iter1.Next()

	iter2 := iter1.Clone()
	iter2.Next()

	if iter1.Key() == iter2.Key() {
		t.Error("Clone should create independent iterator")
	}
}

func TestIteratorIndex(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 10; i++ {
		tree.Put(i, i)
	}

	iter := tree.Iterator()
	iter.SeekToFirst()

	for i := 0; i < 10; i++ {
		if iter.Index() != int64(i) {
			t.Errorf("Iterator index should be %d, got %d", i, iter.Index())
		}
		iter.Next()
	}
}

func TestIntersection(t *testing.T) {
	tree1 := New[int, int](compare.Any[int])
	tree2 := New[int, int](compare.Any[int])

	for i := 1; i < 5; i++ {
		tree1.Put(i, i)
	}

	for i := 2; i < 6; i++ {
		tree2.Put(i, i)
	}

	result := tree1.Intersection(tree2)
	slices := result.Slices()

	if len(slices) != 3 {
		t.Errorf("Intersection should have 3 elements, got %d", len(slices))
	}

	expected := []int{2, 3, 4}
	for i, s := range slices {
		if s.Key != expected[i] {
			t.Errorf("Intersection[%d] = %d, expected %d", i, s.Key, expected[i])
		}
	}
}

func TestUnionSets(t *testing.T) {
	tree1 := New[int, int](compare.Any[int])
	tree2 := New[int, int](compare.Any[int])

	for i := 1; i < 5; i++ {
		tree1.Put(i, i)
	}

	for i := 3; i < 7; i++ {
		tree2.Put(i, i)
	}

	result := tree1.UnionSets(tree2)
	slices := result.Slices()

	if len(slices) != 6 {
		t.Errorf("Union should have 6 elements, got %d", len(slices))
	}

	for i := 0; i < len(slices)-1; i++ {
		if slices[i].Key >= slices[i+1].Key {
			t.Error("Union result should be sorted")
		}
	}
}

func TestDifferenceSets(t *testing.T) {
	tree1 := New[int, int](compare.Any[int])
	tree2 := New[int, int](compare.Any[int])

	for i := 1; i < 5; i++ {
		tree1.Put(i, i)
	}

	for i := 3; i < 7; i++ {
		tree2.Put(i, i)
	}

	result := tree1.DifferenceSets(tree2)
	slices := result.Slices()

	if len(slices) != 2 {
		t.Errorf("Difference should have 2 elements, got %d", len(slices))
	}

	if slices[0].Key != 1 || slices[1].Key != 2 {
		t.Errorf("Difference = [%d, %d], expected [1, 2]", slices[0].Key, slices[1].Key)
	}
}

func TestTrim(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	tree.Trim(25, 75)

	if tree.Size() != 51 {
		t.Errorf("Trim(25, 75) should leave 51 elements, got %d", tree.Size())
	}

	if _, ok := tree.Get(24); ok {
		t.Error("Key 24 should be trimmed")
	}
	if _, ok := tree.Get(25); !ok {
		t.Error("Key 25 should remain")
	}
	if _, ok := tree.Get(75); !ok {
		t.Error("Key 75 should remain")
	}
	if _, ok := tree.Get(76); ok {
		t.Error("Key 76 should be trimmed")
	}
}

func TestTrimByIndex(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	tree.TrimByIndex(25, 75)

	if tree.Size() != 51 {
		t.Errorf("TrimByIndex(25, 75) should leave 51 elements, got %d", tree.Size())
	}

	first := tree.Index(0)
	if first.Key != 25 {
		t.Errorf("First element should be key 25, got %d", first.Key)
	}

	last := tree.Index(tree.Size() - 1)
	if last.Key != 75 {
		t.Errorf("Last element should be key 75, got %d", last.Key)
	}
}

func TestString(t *testing.T) {
	tree := New[int, int](compare.Any[int])
	if tree.String() != "skiplist" {
		t.Error("String() should return 'skiplist'")
	}
}

func TestNewWithMaxLevel(t *testing.T) {
	tree := NewWithMaxLevel[int, int](8, compare.Any[int])

	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	if tree.Height() > 8 {
		t.Errorf("Height should not exceed max level 8, got %d", tree.Height())
	}

	tree2 := NewWithMaxLevel[int, int](0, compare.Any[int])
	if tree2.Height() < 1 {
		t.Error("NewWithMaxLevel(0) should default to at least level 1")
	}

	tree3 := NewWithMaxLevel[int, int](100, compare.Any[int])
	for i := 0; i < 100; i++ {
		tree3.Put(i, i)
	}
	if tree3.Height() > MaxLevel {
		t.Errorf("Height should not exceed MaxLevel %d, got %d", MaxLevel, tree3.Height())
	}
}

func TestConcurrentPut(t *testing.T) {
	tree := New[int, int](compare.Any[int])
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			tree.Put(key, key)
		}(i)
	}

	wg.Wait()

	if tree.Size() != 100 {
		t.Errorf("expected size 100 after concurrent puts, got %d", tree.Size())
	}
}

func TestConcurrentGet(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 100; i++ {
		tree.Put(i, i)
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			tree.Get(key)
		}(i)
	}

	wg.Wait()
}

func TestConcurrentPutAndGet(t *testing.T) {
	tree := New[int, int](compare.Any[int])
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			tree.Put(key, key)
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			tree.Get(key)
		}(i)
	}

	wg.Wait()
}

func TestRandomOperations(t *testing.T) {
	r := rand.New(rand.NewSource(12345))
	tree := New[int, int](compare.Any[int])
	var expected map[int]int = make(map[int]int)

	for i := 0; i < 10000; i++ {
		key := r.Intn(1000)
		value := r.Intn(10000)

		tree.Put(key, value)
		expected[key] = value

		if r.Intn(10) == 0 {
			idx := r.Intn(1000)
			if s := tree.Index(int64(idx)); s != nil {
				idx--
				for idx >= 0 {
					if _, ok := tree.Get(idx); ok {
						break
					}
					idx--
				}
			}
		}
	}

	for key, expectedValue := range expected {
		value, ok := tree.Get(key)
		if !ok {
			t.Errorf("key %d should exist", key)
		}
		if value != expectedValue {
			t.Errorf("key %d: expected %d, got %d", key, expectedValue, value)
		}
	}
}

func TestIteratorSeekRandom(t *testing.T) {
	r := rand.New(rand.NewSource(54321))
	tree := New[int, int](compare.Any[int])

	var keys []int
	for i := 0; i < 1000; i++ {
		key := r.Intn(10000)
		keys = append(keys, key)
		tree.Put(key, key)
	}

	sort.Ints(keys)

	for _, key := range keys[:100] {
		iter := tree.Iterator()
		iter.SeekGE(key)
		if !iter.Valid() {
			t.Errorf("SeekGE(%d) should find a key", key)
		}

		iter.SeekLE(key)
		if !iter.Valid() {
			t.Errorf("SeekLE(%d) should find a key", key)
		}
	}
}

func TestSet(t *testing.T) {
	tree := New[int, int](compare.Any[int])

	if !tree.Set(1, 100) {
		t.Error("Set on new key should return true")
	}

	if tree.Set(1, 200) {
		t.Error("Set on existing key should return false")
	}

	v, _ := tree.Get(1)
	if v != 200 {
		t.Errorf("expected value 200, got %d", v)
	}
}

func TestStress(t *testing.T) {
	r := rand.New(rand.NewSource(67890))
	tree := New[int, int](compare.Any[int])

	for i := 0; i < 50000; i++ {
		key := r.Intn(100000)
		op := r.Intn(10)

		switch op {
		case 0, 1, 2:
			tree.Put(key, key)
		case 3, 4:
			tree.Get(key)
		case 5:
			tree.Remove(key)
		case 6:
			iter := tree.Iterator()
			if iter.Valid() {
				iter.Next()
			}
		case 7:
			iter := tree.Iterator()
			iter.SeekToFirst()
		case 8:
			iter := tree.Iterator()
			iter.SeekToLast()
		case 9:
			tree.Size()
		}
	}
}
