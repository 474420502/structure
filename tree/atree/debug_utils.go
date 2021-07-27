package arraytree

import "fmt"

func (tree *Tree) output(idx int, prefix string, isTail bool, str *string) {

	if idx >= len(tree.datas) {
		return
	}

	node := tree.datas[idx]
	lidx, ridx := getChildrenIndex(idx)
	lc := tree.datas[lidx]
	rc := tree.datas[ridx]

	if rc.Size != 0 {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		tree.output(ridx, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34;40m└── \033[0m"
	} else {
		*str += "\033[31;40m┌── \033[0m"
	}

	*str += "(" + fmt.Sprintf("%v", node.Key) + "->" + fmt.Sprintf("%v", node.Value) + ")" + "\n"

	if lc.Size != 0 {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		tree.output(lidx, newPrefix, true, str)
	}

}

func (tree *Tree) outputfordebug(idx int, prefix string, isTail bool, str *string) {

	if idx >= len(tree.datas) {
		return
	}

	node := tree.datas[idx]
	lidx, ridx := getChildrenIndex(idx)
	// lc := tree.datas[lidx]

	if ridx < len(tree.datas) {

		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		tree.outputfordebug(ridx, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34m└── \033[0m"
	} else {
		*str += "\033[31m┌── \033[0m"
	}

	suffix := "("
	parentv := ""
	pidx := getParentIndex(idx)
	if pidx < 0 {
		parentv = "nil"
	} else {
		parentv = fmt.Sprintf("%v", tree.datas[pidx].Key)
	}

	// suffix += parentv + "|" + fmt.Sprintf("%v",node.Size) + " " + ldirect + "<->" + rdirect + ")"
	suffix += parentv + "|" + fmt.Sprintf("%v", node.Size) + ")"
	// suffix = ""
	k := node.Key

	*str += fmt.Sprintf("%v", k) + suffix + "\n"

	if lidx < len(tree.datas) {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		tree.outputfordebug(lidx, newPrefix, true, str)
	}
}

func (tree *Tree) outputfordebugNoSuffix(idx int, prefix string, isTail bool, str *string) {

	if idx >= len(tree.datas) {
		return
	}

	node := tree.datas[idx]
	lidx, ridx := getChildrenIndex(idx)
	lc := tree.datas[lidx]
	rc := tree.datas[ridx]

	if rc.Size != 0 {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		tree.outputfordebugNoSuffix(ridx, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34m└── \033[0m"
	} else {
		*str += "\033[31m┌── \033[0m"
	}

	k := node.Key

	*str += fmt.Sprintf("%v", k) + "\n"

	if lc.Size != 0 {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		tree.outputfordebugNoSuffix(lidx, newPrefix, true, str)
	}
}

func (tree *Tree) debugString(isSuffix bool) string {
	str := "BinarayList\n"
	idx := 0
	root := &tree.datas[idx]
	if root.Size == 0 {
		return str + "nil"
	}

	if isSuffix {
		tree.outputfordebug(idx, "", true, &str)
	} else {
		tree.outputfordebugNoSuffix(idx, "", true, &str)
	}

	str += "\n["
	dlen := len(tree.datas) - 1
	for i, n := range tree.datas {
		str += fmt.Sprintf("%v", n.Key)
		if i != dlen {
			str += " "
		}

	}
	str += "]"

	return str
}

// func (tree *IndexTree) debugLookNode(cur *Node) {
// 	var temp interface{} = cur.Key
// 	cur.Key = []byte(fmt.Sprintf("\033[32m%s\033[0m", cur.Key))
// 	log.Println(tree.debugString(true))
// 	cur.Key = temp
// }
