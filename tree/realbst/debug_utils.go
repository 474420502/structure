package realbst

import (
	"fmt"
)

// func output(node *Node, prefix string, isTail bool, str *string) {

// 	if node.Children[1] != nil {
// 		newPrefix := prefix
// 		if isTail {
// 			newPrefix += "\033[34m│   \033[0m"
// 		} else {
// 			newPrefix += "    "
// 		}
// 		output(node.Children[1], newPrefix, false, str)
// 	}
// 	*str += prefix
// 	if isTail {
// 		*str += "\033[34;40m└── \033[0m"
// 	} else {
// 		*str += "\033[31;40m┌── \033[0m"
// 	}

// 	*str += "(" + fmt.Sprintf("%v", node.Key) + "->" + fmt.Sprintf("%v", node.Value) + ")" + "\n"

// 	if node.Children[0] != nil {
// 		newPrefix := prefix
// 		if isTail {
// 			newPrefix += "    "
// 		} else {
// 			newPrefix += "\033[31m│   \033[0m"
// 		}
// 		output(node.Children[0], newPrefix, true, str)
// 	}

// }

func outputfordebug(node *Node, prefix string, isTail bool, str *string, start, end *Node) {

	if node == nil {
		return
	}

	if node.Children[1] != end {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		outputfordebug(node.Children[1], newPrefix, false, str, node, end)
	}

	*str += prefix
	if isTail {
		*str += "\033[34m└── \033[0m"
	} else {
		*str += "\033[31m┌── \033[0m"
	}

	k := node.Key

	*str += fmt.Sprintf("%v", k) + "\n"

	if node.Children[0] != start {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		outputfordebug(node.Children[0], newPrefix, true, str, start, node)
	}
}

func (tree *Tree) debugString() string {
	str := "realbst\n"
	root := tree.Center
	if root == nil {
		return str + "nil"
	}

	outputfordebug(root, "", true, &str, nil, nil)

	return str
}

// func lookTree(root *Node) string {
// 	str := "\n"
// 	if root == nil {
// 		return str + "nil"
// 	}
// 	outputfordebug(root, "", true, &str, nil. nil)
// 	return str
// }

// func (tree *IndexTree) debugLookNode(cur *Node) {
// 	var temp interface{} = cur.Key
// 	cur.Key = []byte(fmt.Sprintf("\033[32m%s\033[0m", cur.Key))
// 	log.Println(tree.debugString(true))
// 	cur.Key = temp
// }
