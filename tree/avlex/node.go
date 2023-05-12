package avlex

import (
	"fmt"
)

type Node[KEY any, VALUE any] struct {
	Key      KEY
	Value    VALUE
	Height   int
	Children [2]*Node[KEY, VALUE]
}

func (node *Node[KEY, VALUE]) String() string {
	return fmt.Sprintf("%v(%v)", node.Key, node.Value)
}

func (node *Node[KEY, VALUE]) updateHeight() bool {
	lh, rh := getHeight(node.Children[0])+1, getHeight(node.Children[1])+1
	if lh > rh {
		if node.Height != lh {
			node.Height = lh
			return true
		}
	} else {
		if node.Height != rh {
			node.Height = rh
			return true
		}
	}

	return false
}

func (node *Node[KEY, VALUE]) updateHeightOneChild(child int) {
	node.Height = getHeight(node.Children[child]) + 1
}

func (node *Node[KEY, VALUE]) rebalance(parent *Node[KEY, VALUE], child int) bool {

	lh, rh := getHeight(node.Children[0]), getHeight(node.Children[1])

	diff := lh - rh

	if diff > 1 {
		sub := node.Children[0]
		if getHeight(sub.Children[1]) > getHeight(sub.Children[0]) {
			rightRotateWithLeft(parent, child)
		} else {
			rightRotate(parent, child)
		}
		return true
	} else if diff < -1 {
		sub := node.Children[1]
		if getHeight(sub.Children[0]) > getHeight(sub.Children[1]) {
			leftRotateWithRight(parent, child)
		} else {
			leftRotate(parent, child)
		}
		return true
	} else {
		if lh > rh {
			if node.Height != lh+1 {
				node.Height = lh + 1
				return true
			}
		} else {
			if node.Height != rh+1 {
				node.Height = rh + 1
				return true
			}
		}

	}

	return false
}

func newNode[KEY any, VALUE any]() *Node[KEY, VALUE] {
	return &Node[KEY, VALUE]{
		Height: 1,
	}
}

func getHeight[KEY, VALUE any](cur *Node[KEY, VALUE]) int {
	if cur == nil {
		return 0
	}
	return cur.Height
}

func updateHeight[KEY, VALUE any](cur *Node[KEY, VALUE]) {
	lh, rh := getHeight(cur.Children[0]), getHeight(cur.Children[1])
	if lh > rh {
		cur.Height = lh + 1
	} else {
		cur.Height = rh + 1
	}
}

func leftRotateWithRight[KEY, VALUE any](parent *Node[KEY, VALUE], child int) {
	cur := parent.Children[child] //
	sub := cur.Children[1]        // right
	subsub := sub.Children[0]     // right->left

	parent.Children[child] = subsub

	children := subsub.Children
	subsub.Children[0] = cur
	subsub.Children[1] = sub

	cur.Children[1] = children[0]
	sub.Children[0] = children[1]

	cur.updateHeightOneChild(0)
	sub.updateHeightOneChild(1)
	subsub.updateHeightOneChild(0)

}

func rightRotateWithLeft[KEY, VALUE any](parent *Node[KEY, VALUE], child int) {
	cur := parent.Children[child] //
	sub := cur.Children[0]        // left
	subsub := sub.Children[1]     // left->right

	parent.Children[child] = subsub

	children := subsub.Children
	subsub.Children[1] = cur
	subsub.Children[0] = sub

	cur.Children[0] = children[1]
	sub.Children[1] = children[0]

	cur.updateHeightOneChild(1)
	sub.updateHeightOneChild(0)
	subsub.updateHeightOneChild(1)
}

func leftRotate[KEY, VALUE any](parent *Node[KEY, VALUE], child int) {
	cur := parent.Children[child]
	sub := cur.Children[1]

	parent.Children[child] = sub

	cur.Children[1] = sub.Children[0]
	sub.Children[0] = cur

	cur.updateHeightOneChild(1)
	sub.updateHeightOneChild(0)
}

func rightRotate[KEY, VALUE any](parent *Node[KEY, VALUE], child int) {
	cur := parent.Children[child]
	sub := cur.Children[0]

	parent.Children[child] = sub

	cur.Children[0] = sub.Children[1]
	sub.Children[1] = cur

	cur.updateHeightOneChild(0)
	sub.updateHeightOneChild(1)
}

func view[KEY, VALUE any](root *Node[KEY, VALUE]) (result string) {
	result = "\n"
	if root == nil {
		result += "└── nil"
		return
	}
	var nmap = make(map[*Node[KEY, VALUE]]int)
	outputfordebug(nmap, root, "", true, &result)
	return
}

func (tree *Tree[KEY, VALUE]) checkHeightTree(root *Node[KEY, VALUE]) bool {
	errorNode := 0
	var check func(root *Node[KEY, VALUE]) = nil
	check = func(root *Node[KEY, VALUE]) {
		if root == nil {
			return
		}

		height := 0
		lh, rh := getHeight(root.Children[0]), getHeight(root.Children[1])
		if lh > rh {
			height = lh
		} else {
			height = rh
		}

		if root.Height != height+1 {
			errorNode += 1
		}

		if root.Children[0] != nil {
			if tree.Compare(root.Key, root.Children[0].Key) != 0 {
				panic("error tree tree.Compare(root.Key, root.Children[0].Key) != 0")
			}
		}

		if root.Children[1] != nil {
			if tree.Compare(root.Key, root.Children[1].Key) != 1 {
				panic("error tree tree.Compare(root.Key, root.Children[1].Key) != 1 ")
			}
		}

		check(root.Children[0])
		check(root.Children[1])
	}
	check(root)
	return errorNode == 0
}

func outputfordebug[KEY, VALUE any](nmap map[*Node[KEY, VALUE]]int, node *Node[KEY, VALUE], prefix string, isTail bool, str *string) {
	if v, ok := nmap[node]; ok {
		if v > 2 {
			return
		}
		nmap[node] = v + 1
	} else {
		nmap[node] = 1
	}

	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		outputfordebug(nmap, node.Children[1], newPrefix, false, str)
	}

	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	*str += fmt.Sprintf("%v(%d)", node.Key, node.Height) + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		outputfordebug(nmap, node.Children[0], newPrefix, true, str)
	}
}
