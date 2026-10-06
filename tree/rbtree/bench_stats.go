package rbtree

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
	if tree.root == nil {
		return 0, 0, 0, 0
	}

	type depthNode struct {
		node  *Node[KEY, VALUE]
		depth int
	}

	queue := []depthNode{{node: tree.root, depth: 1}}
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

		if item.node.Left != nil {
			queue = append(queue, depthNode{node: item.node.Left, depth: item.depth + 1})
		}
		if item.node.Right != nil {
			queue = append(queue, depthNode{node: item.node.Right, depth: item.depth + 1})
		}
	}

	sort.Ints(depths)
	p50Depth = depths[(len(depths)-1)/2]
	p95Depth = depths[((len(depths)-1)*95)/100]

	return height, float64(totalDepth) / float64(count), p50Depth, p95Depth
}
