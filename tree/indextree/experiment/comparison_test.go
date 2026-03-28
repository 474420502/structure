package experiment

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/tree/indextree"
)

func TestCompareRotations(t *testing.T) {
	avltree := New[int64, int](compare.Any[int64])
	mainTree := indextree.New(compare.Any[int64])
	defaultTree := NewIndexTreeDefault[int64, int](compare.Any[int64])
	shift3 := NewShiftTolerance[int64, int](compare.Any[int64], 3)

	for i := 0; i < 10000; i++ {
		avltree.Put(int64(i), i)
		mainTree.Put(int64(i), i)
		defaultTree.Put(int64(i), i)
		shift3.Put(int64(i), i)
	}

	avlStats := avltree.BenchmarkStats()
	defaultStats := defaultTree.BenchmarkStats()
	shift3Stats := shift3.BenchmarkStats()

	fmt.Printf("\n=== After 10000 inserts ===\n")
	fmt.Printf("AVL: Height=%d, SingleRot=%d, DoubleRot=%d, TotalRot=%d\n",
		avlStats.Height, avlStats.SingleRotations, avlStats.DoubleRotations,
		avlStats.SingleRotations+avlStats.DoubleRotations)
	fmt.Printf("IndexTree(main): Size=%d\n", mainTree.Size())
	fmt.Printf("ShiftToleranceTree(indextree-default): Height=%d, SingleRot=%d, DoubleRot=%d, TotalRot=%d\n",
		defaultStats.Height, defaultStats.SingleRotations, defaultStats.DoubleRotations,
		defaultStats.SingleRotations+defaultStats.DoubleRotations)
	fmt.Printf("ShiftToleranceTree(shift=3): Height=%d, SingleRot=%d, DoubleRot=%d, TotalRot=%d\n",
		shift3Stats.Height, shift3Stats.SingleRotations, shift3Stats.DoubleRotations,
		shift3Stats.SingleRotations+shift3Stats.DoubleRotations)
}

func TestDifferentShifts(t *testing.T) {
	testCases := []struct {
		name  string
		shift int64
	}{
		{"indextree-default", 2},
		{"shift=3", 3},
		{"shift=4", 4},
		{"shift=5", 5},
	}

	avltree := New[int64, int](compare.Any[int64])
	for i := 0; i < 10000; i++ {
		avltree.Put(int64(i), i)
	}
	avlStats := avltree.BenchmarkStats()

	fmt.Printf("\n=== Sequential Insert Shift Comparison (AVL Height=%d) ===\n", avlStats.Height)
	fmt.Printf("%-22s %10s %10s %10s\n", "Config", "Height", "SingleRot", "DoubleRot")

	for _, tc := range testCases {
		tree := NewShiftTolerance[int64, int](compare.Any[int64], tc.shift)
		for i := 0; i < 10000; i++ {
			tree.Put(int64(i), i)
		}
		stats := tree.BenchmarkStats()
		fmt.Printf("%-22s %10d %10d %10d\n", tc.name, stats.Height, stats.SingleRotations, stats.DoubleRotations)
	}
}

func TestRandomInsert(t *testing.T) {
	r := rand.New(rand.NewSource(12345))
	keys := make([]int64, 10000)
	for i := range keys {
		keys[i] = r.Int63()
	}

	avltree := New[int64, int](compare.Any[int64])
	for _, key := range keys {
		avltree.Put(key, int(key))
	}
	avlStats := avltree.BenchmarkStats()

	fmt.Printf("\n=== Random Insert (10000 keys, AVL Height=%d) ===\n", avlStats.Height)
	fmt.Printf("%-22s %10s %10s %10s\n", "Config", "Height", "SingleRot", "DoubleRot")
	fmt.Printf("%-22s %10d %10d %10d\n", "AVL", avlStats.Height, avlStats.SingleRotations, avlStats.DoubleRotations)

	for _, shift := range []int64{2, 3, 4, 5} {
		tree := NewShiftTolerance[int64, int](compare.Any[int64], shift)
		for _, key := range keys {
			tree.Put(key, int(key))
		}
		stats := tree.BenchmarkStats()
		name := fmt.Sprintf("shift=%d", shift)
		if shift == 2 {
			name = "indextree-default"
		}
		fmt.Printf("%-22s %10d %10d %10d\n", name, stats.Height, stats.SingleRotations, stats.DoubleRotations)
	}
}

func TestIndexTreeDefaultHeightGap(t *testing.T) {
	r := rand.New(rand.NewSource(12345))
	keys := make([]int64, 10000)
	for i := range keys {
		keys[i] = r.Int63()
	}

	avltree := New[int64, int](compare.Any[int64])
	for _, key := range keys {
		avltree.Put(key, int(key))
	}
	avlStats := avltree.BenchmarkStats()

	defaultTree := NewIndexTreeDefault[int64, int](compare.Any[int64])
	for _, key := range keys {
		defaultTree.Put(key, int(key))
	}
	stats := defaultTree.BenchmarkStats()

	if stats.Height > avlStats.Height+2 {
		t.Fatalf("indextree-default tolerance tree height %d exceeds avl height %d by more than 2", stats.Height, avlStats.Height)
	}

	if stats.SingleRotations+stats.DoubleRotations >= avlStats.SingleRotations+avlStats.DoubleRotations {
		t.Fatalf("indextree-default tolerance tree rotations %d should stay below avl rotations %d", stats.SingleRotations+stats.DoubleRotations, avlStats.SingleRotations+avlStats.DoubleRotations)
	}
}

func TestIndexTreeDefaultMatchesMainPackage(t *testing.T) {
	testCases := []struct {
		name string
		keys []int64
	}{
		{name: "sequential-1000", keys: sequentialKeys(1000)},
		{name: "random-1000-seed-12345", keys: randomKeys(1000, 12345)},
		{name: "random-5000-seed-67890", keys: randomKeys(5000, 67890)},
	}

	for _, tc := range testCases {
		mainTree := indextree.New(compare.Any[int64])
		defaultTree := NewIndexTreeDefault[int64, int](compare.Any[int64])

		for _, key := range tc.keys {
			mainTree.Put(key, int(key))
			defaultTree.Put(key, int(key))
		}

		mainStats := mainTree.BenchmarkStats()
		defaultStats := defaultTree.BenchmarkStats()

		if int64(defaultTree.Size()) != mainTree.Size() {
			t.Fatalf("%s: size mismatch default=%d main=%d", tc.name, defaultTree.Size(), mainTree.Size())
		}

		if defaultStats.Height != mainStats.Height {
			t.Fatalf("%s: height mismatch default=%d main=%d", tc.name, defaultStats.Height, mainStats.Height)
		}

		if defaultStats.SingleRotations != mainStats.SingleRotations {
			t.Fatalf("%s: single rotations mismatch default=%d main=%d", tc.name, defaultStats.SingleRotations, mainStats.SingleRotations)
		}

		if defaultStats.DoubleRotations != mainStats.DoubleRotations {
			t.Fatalf("%s: double rotations mismatch default=%d main=%d", tc.name, defaultStats.DoubleRotations, mainStats.DoubleRotations)
		}

		if defaultStats.AvgDepth != mainStats.AvgDepth {
			t.Fatalf("%s: avg depth mismatch default=%f main=%f", tc.name, defaultStats.AvgDepth, mainStats.AvgDepth)
		}

		if defaultStats.P50Depth != mainStats.P50Depth || defaultStats.P95Depth != mainStats.P95Depth {
			t.Fatalf("%s: percentile depth mismatch default=(%d,%d) main=(%d,%d)", tc.name, defaultStats.P50Depth, defaultStats.P95Depth, mainStats.P50Depth, mainStats.P95Depth)
		}
	}
}

func sequentialKeys(size int) []int64 {
	keys := make([]int64, size)
	for index := range keys {
		keys[index] = int64(index)
	}
	return keys
}

func randomKeys(size int, seed int64) []int64 {
	rng := rand.New(rand.NewSource(seed))
	keys := make([]int64, size)
	for index := range keys {
		keys[index] = rng.Int63()
	}
	return keys
}
