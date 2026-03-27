package itree

import (
	"fmt"
	"sort"
)

var IsDebugString = false

func debugOutput[KEY, VALUE any](node *Node[KEY, VALUE], prefix string, isTail bool, str *string) {
	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		debugOutput(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34;40m└── \033[0m"
	} else {
		*str += "\033[31;40m┌── \033[0m"
	}

	*str += fmt.Sprintf("(%v->%v)", node.Key, node.Value) + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		debugOutput(node.Children[0], newPrefix, true, str)
	}
}

func debugOutputWithSuffix[KEY, VALUE any](node *Node[KEY, VALUE], prefix string, isTail bool, str *string, deep int) {
	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		debugOutputWithSuffix(node.Children[1], newPrefix, false, str, deep+1)
	}
	*str += prefix
	if isTail {
		*str += "\033[34m└── \033[0m"
	} else {
		*str += "\033[31m┌── \033[0m"
	}

	suffix := "("
	suffix += fmt.Sprintf("|%v|%d", node.Size, deep) + ")"
	k := node.Key

	*str += fmt.Sprintf("%v", k) + suffix + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		debugOutputWithSuffix(node.Children[0], newPrefix, true, str, deep+1)
	}
}

func debugOutputNoSuffix[KEY, VALUE any](node *Node[KEY, VALUE], prefix string, isTail bool, str *string) {
	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		debugOutputNoSuffix(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34m└── \033[0m"
	} else {
		*str += "\033[31m┌── \033[0m"
	}

	k := node.Key

	*str += fmt.Sprintf("%v", k) + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		debugOutputNoSuffix(node.Children[0], newPrefix, true, str)
	}
}

func (tree *Tree[KEY, VALUE]) debugString(isSuffix bool) string {
	str := "Tree\n"
	root := tree.getRoot()
	if root == nil {
		return str + "nil"
	}

	if isSuffix {
		debugOutputWithSuffix(root, "", true, &str, 1)
	} else {
		debugOutputNoSuffix(root, "", true, &str)
	}

	return str
}

func (tree *Tree[KEY, VALUE]) String() string {
	return tree.debugString(IsDebugString)
}

func (tree *Tree[KEY, VALUE]) resetRotationStats() {
	tree.singleRotations = 0
	tree.doubleRotations = 0
}

func (tree *Tree[KEY, VALUE]) rotationStats() (single int, double int) {
	return tree.singleRotations, tree.doubleRotations
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
	depths := make([]int, 0, root.Size)

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
