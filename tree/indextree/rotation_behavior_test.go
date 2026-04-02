package indextree

import (
	"testing"

	"github.com/474420502/structure/compare"
)

func TestDoubleRotationOnPutRightLeft(t *testing.T) {
	tree := New(compare.Any[int])
	for _, key := range []int{1, 3, 2} {
		tree.Put(key, key)
	}

	stats := tree.BenchmarkStats()
	if stats.DoubleRotations != 1 {
		t.Fatalf("expected one double rotation for RL insertion, got single=%d double=%d", stats.SingleRotations, stats.DoubleRotations)
	}
	if stats.SingleRotations != 0 {
		t.Fatalf("expected no single rotation for RL insertion, got single=%d", stats.SingleRotations)
	}
	root := tree.getRoot()
	if root == nil || root.Key != 2 {
		t.Fatalf("expected root key 2 after RL double rotation, got %#v", root)
	}
}

func TestDoubleRotationOnPutLeftRight(t *testing.T) {
	tree := New(compare.Any[int])
	for _, key := range []int{3, 1, 2} {
		tree.Put(key, key)
	}

	stats := tree.BenchmarkStats()
	if stats.DoubleRotations != 1 {
		t.Fatalf("expected one double rotation for LR insertion, got single=%d double=%d", stats.SingleRotations, stats.DoubleRotations)
	}
	if stats.SingleRotations != 0 {
		t.Fatalf("expected no single rotation for LR insertion, got single=%d", stats.SingleRotations)
	}
	root := tree.getRoot()
	if root == nil || root.Key != 2 {
		t.Fatalf("expected root key 2 after LR double rotation, got %#v", root)
	}
}

func TestSequentialPutUsesOnlySingleRotations(t *testing.T) {
	tree := New(compare.Any[int])
	for key := 0; key < 1000; key++ {
		tree.Put(key, key)
	}

	stats := tree.BenchmarkStats()
	if stats.DoubleRotations != 0 {
		t.Fatalf("expected sequential insertions to avoid double rotations, got single=%d double=%d", stats.SingleRotations, stats.DoubleRotations)
	}
	if stats.SingleRotations == 0 {
		t.Fatalf("expected sequential insertions to trigger single rotations")
	}
}

func TestRandomPutTriggersDoubleRotations(t *testing.T) {
	data := newBenchDataWithSeed(10000, 12345)
	tree := New(compare.Any[int64])
	for _, key := range data {
		tree.Put(key, key)
	}

	stats := tree.BenchmarkStats()
	if stats.DoubleRotations == 0 {
		t.Fatalf("expected random insertions to trigger double rotations, got single=%d double=%d", stats.SingleRotations, stats.DoubleRotations)
	}
	if stats.SingleRotations <= stats.DoubleRotations {
		t.Fatalf("expected outer-heavy cases to remain common, got single=%d double=%d", stats.SingleRotations, stats.DoubleRotations)
	}
}