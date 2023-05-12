package avlbest

import (
	"log"

	"github.com/474420502/structure/compare"
)

type Tree[KEY, VALUE any] struct {
	Center  *Node[KEY, VALUE]
	Compare compare.Compare[KEY]
	// HeightDiff int
	zero VALUE
	Size int
	st   *Stack[Node[KEY, VALUE]]
}

func NewTree[KEY, VALUE any](Compare compare.Compare[KEY]) *Tree[KEY, VALUE] {

	tree := &Tree[KEY, VALUE]{
		Center:  &Node[KEY, VALUE]{Height: 0},
		Compare: Compare,
		st:      NewStack[Node[KEY, VALUE]](),
	}

	// tree.Center.Children[0] = tree.Center
	return tree
}

func (tree *Tree[KEY, VALUE]) View() (result string) {
	result = "\n"
	if tree.Center.Children[1] == nil {
		result += "└── nil"
		return
	}

	var nmap = make(map[*Node[KEY, VALUE]]int)
	outputfordebug(nmap, tree.Center.Children[1], "", true, &result)
	return
}

func (tree *Tree[KEY, VALUE]) check() (result string) {
	if !tree.checkHeightTree(tree.Center.Children[1]) {
		log.Panic("height error")
	}
	return
}

func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool {

	tree.st.Clear()
	cmp := 1
	cur := tree.Center.Children[cmp]
	tree.st.Push(tree.Center, cmp)

	for {

		if cur == nil {

			node := newNode(key, value)
			lastdir := tree.st.Pop()
			parent := lastdir.Node
			child := lastdir.Child

			parent.Children[child] = node
			tree.Size += 1

			if parent.Children[^child+2] == nil {
				parent.Height += 1
				tree.st.PopFast(1)
				for fixParent := tree.st.Pop(); fixParent != nil; fixParent = tree.st.Pop() {
					fixCur := fixParent.Node.Children[fixParent.Child]
					if !fixCur.rebalance(fixParent.Node, fixParent.Child) {
						break
					}
				}
			}

			return false
		}

		cmp = tree.Compare(cur.Key, key)
		if cmp < 0 {
			cur.Key = key
			cur.Value = value
			return true
		} else {
			tree.st.Push(cur, cmp)
			cur = cur.Children[cmp]
		}

	}

}

func (tree *Tree[KEY, VALUE]) Remove(key KEY) bool {
	tree.st.Clear()
	cmp := 1
	cur := tree.Center.Children[cmp]
	tree.st.Push(tree.Center, cmp)

	for {

		if cur == nil {
			return false
		}

		cmp = tree.Compare(cur.Key, key)

		// if fmt.Sprintf("%v", key) == "24" && fmt.Sprintf("%v", cur.Key) == "24" && tree.Size == 8 {
		// 	log.Println()
		// }
		if cmp < 0 {

			tree.Size -= 1
			lastdir := tree.st.Peek()
			parent := lastdir.Node
			child := lastdir.Child
			// log.Println(tree.View(), key)
			if cur.Children[0] == nil {
				parent.Children[child] = cur.Children[1]
				tree.st.PopFast(1)
			} else if cur.Children[1] == nil {
				parent.Children[child] = cur.Children[0]
				tree.st.PopFast(1)
			} else {

				tree.st.Push(cur, child)
				replacer := cur.Children[child]
				rchild := ^child + 2

				for replacer.Children[rchild] != nil {
					tree.st.Push(replacer, rchild)
					replacer = replacer.Children[rchild]
				}

				fixParent := tree.st.Pop()
				fixParent.Node.Children[fixParent.Child] = replacer.Children[child]
				// fixParent.Node.updateHeightOneChild(^fixParent.Child + 2)

				cur.Key = replacer.Key
				cur.Value = replacer.Value

			}
			// log.Println(tree.View())
			for fixParent := tree.st.Pop(); fixParent != nil; fixParent = tree.st.Pop() {
				fixCur := fixParent.Node.Children[fixParent.Child]
				if !fixCur.rebalance(fixParent.Node, fixParent.Child) {
					break
				}
			}
			// log.Println(tree.View())
			return true
		} else {
			tree.st.Push(cur, cmp)
			cur = cur.Children[cmp]
		}
	}

}
