package indextree

import "fmt"

const (
	color_1 = "\033[31m%v\033[0m"
	color_2 = "\033[32m%v\033[0m"
	color_3 = "\033[33m%v\033[0m"
	color_4 = "\033[34m%v\033[0m"
	color_5 = "\033[35m%v\033[0m"
	color_6 = "\033[36m%v\033[0m"
	color_7 = "\033[37m%v\033[0m"
	color_8 = "\033[38m%v\033[0m"
	color_9 = "\033[39m%v\033[0m"
)

type colorNode[T any] struct {
	Node  *Node[T]
	Color string
}

func output[T any](node *Node[T], prefix string, isTail bool, str *string) {

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

func outputfordebug[T any](node *Node[T], prefix string, isTail bool, str *string, deep int) {

	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		outputfordebug(node.Children[1], newPrefix, false, str, deep+1)
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
	suffix += parentv + fmt.Sprintf("|%v|%d", node.Size, deep) + ")"
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
		outputfordebug(node.Children[0], newPrefix, true, str, deep+1)
	}
}

func outputfordebugNoSuffix[T any](node *Node[T], prefix string, isTail bool, str *string) {

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

func (tree *Tree[T]) debugString(isSuffix bool) string {

	str := "IndexTree\n"
	root := tree.getRoot()
	if root == nil {
		return str + "nil"
	}

	if isSuffix {
		outputfordebug(root, "", true, &str, 1)
	} else {
		outputfordebugNoSuffix(root, "", true, &str)
	}

	return str
}

func lookTree[T any](root *Node[T]) string {
	str := "\n"
	if root == nil {
		return str + "nil"
	}
	outputfordebugNoSuffix(root, "", true, &str)
	return str
}

// func (tree *IndexTree) debugLookNode[T any](cur *Node[T]) {
// 	var temp interface{} = cur.Key
// 	cur.Key = []byte(fmt.Sprintf("\033[32m%s\033[0m", cur.Key))
// 	log.Println(tree.debugString(true))
// 	cur.Key = temp
// }
