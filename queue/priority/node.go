package treequeue

import (
	"fmt"
)

type Node[KEY any, VALUE any] struct {
	Key      KEY
	Value    VALUE
	Size     int
	Children [2]*Node[KEY, VALUE]
}

func (node *Node[KEY, VALUE]) String() string {
	return fmt.Sprintf("%v(%v)", node.Key, node.Value)
}

func (node *Node[KEY, VALUE]) updateSize() {
	shouldsize := getSize(node.Children[0]) + getSize(node.Children[1]) + 1
	if node.Size != shouldsize {
		node.Size = shouldsize
	}
}

func (node *Node[KEY, VALUE]) getSize() int {
	if node == nil {
		return 0
	}
	return node.Size
}

func newNode[KEY any, VALUE any]() *Node[KEY, VALUE] {
	return &Node[KEY, VALUE]{
		Size: 1,
	}
}

// inline
func getSize[KEY, VALUE any](cur *Node[KEY, VALUE]) int {
	if cur == nil {
		return 0
	}
	return cur.Size
}

func getMaybeHeight(size int) (height uint) {
	for size != 0 {
		size = size >> 1
		height++
	}
	return height
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

	cur.updateSize()
	sub.updateSize()
	subsub.updateSize()

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

	cur.updateSize()
	sub.updateSize()
	subsub.updateSize()
}

func leftRotate[KEY, VALUE any](parent *Node[KEY, VALUE], child int) {
	cur := parent.Children[child]
	sub := cur.Children[1]

	parent.Children[child] = sub

	cur.Children[1] = sub.Children[0]
	sub.Children[0] = cur

	cur.updateSize()
	sub.updateSize()
}

func rightRotate[KEY, VALUE any](parent *Node[KEY, VALUE], child int) {
	cur := parent.Children[child]
	sub := cur.Children[0]

	parent.Children[child] = sub

	cur.Children[0] = sub.Children[1]
	sub.Children[1] = cur

	cur.updateSize()
	sub.updateSize()
}

func (root *Node[KEY, VALUE]) view() (result string) {
	result = "\n"
	if root == nil {
		result += "└── nil"
		return
	}
	var nmap = make(map[*Node[KEY, VALUE]]int)
	outputfordebug(nmap, root, "", true, &result, nil)
	return
}

func (tree *Tree[KEY, VALUE]) checkSizeTree(root *Node[KEY, VALUE]) *Node[KEY, VALUE] {

	var check func(root *Node[KEY, VALUE]) *Node[KEY, VALUE] = nil
	check = func(root *Node[KEY, VALUE]) *Node[KEY, VALUE] {
		if root == nil {
			return nil
		}

		lsize, rsize := getSize(root.Children[0]), getSize(root.Children[1])

		if root.Size != lsize+rsize+1 {
			return root
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

		result := check(root.Children[0])
		if result != nil {
			return result
		}
		return check(root.Children[1])
	}

	return check(root)
}

func outputfordebug[KEY, VALUE any](nmap map[*Node[KEY, VALUE]]int, node *Node[KEY, VALUE], prefix string, isTail bool, str *string, target *Node[KEY, VALUE]) {
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
		outputfordebug(nmap, node.Children[1], newPrefix, false, str, target)
	}

	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	if target != nil && target == node {
		*str += fmt.Sprintf("%v(%d)*", node.Key, node.Size) + "\n"
	} else {
		*str += fmt.Sprintf("%v(%d)", node.Key, node.Size) + "\n"
	}

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		outputfordebug(nmap, node.Children[0], newPrefix, true, str, target)
	}
}
