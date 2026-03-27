package avl

import "sort"

type BenchmarkStats struct {
	SingleRotations int
	DoubleRotations int
	Height          int
	AvgDepth        float64
	P50Depth        int
	P95Depth        int
}

func (tree *Tree[KEY, VALUE]) shapeStats() (height int, avgDepth float64, p50Depth int, p95Depth int) {
	root := tree.getRoot()
	if root == nil {
		return 0, 0, 0, 0
	}

	type depthNode struct {
		node  *Node[KEY, VALUE]
		depth int
	}

	queue := []depthNode{{node: root, depth: 1}}
	totalDepth := 0
	count := 0
	depths := make([]int, 0, tree.size)

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		count++
		totalDepth += item.depth
		depths = append(depths, item.depth)
		if item.depth > height {
			height = item.depth
		}

		if left := item.node.Children[0]; left != nil {
			queue = append(queue, depthNode{node: left, depth: item.depth + 1})
		}
		if right := item.node.Children[1]; right != nil {
			queue = append(queue, depthNode{node: right, depth: item.depth + 1})
		}
	}

	sort.Ints(depths)
	p50Depth = depths[(len(depths)-1)/2]
	p95Depth = depths[((len(depths)-1)*95)/100]

	return height, float64(totalDepth) / float64(count), p50Depth, p95Depth
}
