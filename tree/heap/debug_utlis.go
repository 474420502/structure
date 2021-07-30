package heap

import "fmt"

func (tree *Tree) outputfordebug(idx int, prefix string, isTail bool, str *string) {

	ridx := idx<<1 + 2
	if ridx < tree.size && tree.elements[ridx] != nil {
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

	*str += fmt.Sprintf("%v", tree.elements[idx]) + "\n"

	lidx := idx<<1 + 1

	if lidx < tree.size && tree.elements[lidx] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		tree.outputfordebug(lidx, newPrefix, true, str)
	}
}

func (tree *Tree) debugString() string {
	str := "Heap\n"
	if tree.elements[0] == nil {
		return str + "nil"
	}

	tree.outputfordebug(0, "", true, &str)

	return str
}
