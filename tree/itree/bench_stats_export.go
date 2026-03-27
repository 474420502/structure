package itree

type BenchmarkStats struct {
	SingleRotations int
	DoubleRotations int
	Height          int
	AvgDepth        float64
	P50Depth        int
	P95Depth        int
}

func (tree *Tree[KEY, VALUE]) ResetBenchmarkStats() {
	tree.resetRotationStats()
}

func (tree *Tree[KEY, VALUE]) BenchmarkStats() BenchmarkStats {
	height, avgDepth, p50Depth, p95Depth := tree.shapeStats()
	single, double := tree.rotationStats()
	return BenchmarkStats{
		SingleRotations: single,
		DoubleRotations: double,
		Height:          height,
		AvgDepth:        avgDepth,
		P50Depth:        p50Depth,
		P95Depth:        p95Depth,
	}
}