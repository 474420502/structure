package treequeue

import "fmt"

func output(node *Node, prefix string, isTail bool, str *string) {

	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		output(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34;40m└── \033[0m"
	} else {
		*str += "\033[31;40m┌── \033[0m"
	}

	*str += "(" + fmt.Sprintf("%v", node.Key) + "->" + fmt.Sprintf("%v", node.Value) + ")" + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		output(node.Children[0], newPrefix, true, str)
	}

}

func outputfordebug(node *Node, prefix string, isTail bool, str *string) {

	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		outputfordebug(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34m└── \033[0m"
	} else {
		*str += "\033[31m┌── \033[0m"
	}

	suffix := "("
	parentv := ""
	if node.Parent == nil {
		parentv = "nil"
	} else {
		parentv = fmt.Sprintf("%v", node.Parent.Key)
	}

	// suffix += parentv + "|" + fmt.Sprintf("%v",node.Size) + " " + ldirect + "<->" + rdirect + ")"
	suffix += parentv + "|" + fmt.Sprintf("%v", node.Size) + ")"
	// suffix = ""
	k := node.Key

	*str += fmt.Sprintf("%v", k) + suffix + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		outputfordebug(node.Children[0], newPrefix, true, str)
	}
}

func outputfordebugNoSuffix(node *Node, prefix string, isTail bool, str *string) {

	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		outputfordebugNoSuffix(node.Children[1], newPrefix, false, str)
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
		outputfordebugNoSuffix(node.Children[0], newPrefix, true, str)
	}
}

func (tree *Tree) debugString(isSuffix bool) string {
	str := "BinarayList\n"
	root := tree.getRoot()
	if root == nil {
		return str + "nil"
	}

	if isSuffix {
		outputfordebug(root, "", true, &str)
	} else {
		outputfordebugNoSuffix(root, "", true, &str)
	}

	return str
}

// func (tree *IndexTree) debugLookNode(cur *Node) {
// 	var temp interface{} = cur.Key
// 	cur.Key = []byte(fmt.Sprintf("\033[32m%s\033[0m", cur.Key))
// 	log.Println(tree.debugString(true))
// 	cur.Key = temp
// }
