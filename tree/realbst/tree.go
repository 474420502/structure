package realbst

import (
	"fmt"
	"log"

	"github.com/474420502/structure/compare"
)

type Node struct {
	Children [2]*Node
	Key      interface{}
	Value    interface{}
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", n.Key)
}

type Tree struct {
	Center   *Node
	Head     *Node
	Tail     *Node
	LowSize  int64
	HighSize int64
	Size     int64
	Compare  compare.Compare
}

func New(comp compare.Compare) *Tree {
	return &Tree{
		Compare: comp,
	}
}

func (tree *Tree) Get(key interface{}) interface{} {
	const L = 0
	const R = 1

	center := tree.Center
	if center == nil {
		return nil
	}

	var start, end *Node
	for {
		c := tree.Compare(key, center.Key)
		switch {
		case c < 0:
			if center.Children[L] == start {
				return nil
			}
			end = center
			center = center.Children[L]

		case c > 0:
			if center.Children[R] == end {
				return nil
			}
			start = center
			center = center.Children[R]
		default:

			return center.Value
		}
	}
}

func (tree *Tree) Put(key, value interface{}) bool {

	const L = 0
	const R = 1

	center := tree.Center
	if center == nil {
		tree.Center = &Node{Key: key, Value: value}
		tree.Head = tree.Center
		tree.Tail = tree.Center
		tree.Size++
		return true
	}

	var start, end *Node
	for {
		c := tree.Compare(key, center.Key)
		switch {
		case c < 0:
			if center.Children[L] == start {

				node := &Node{Key: key, Value: value}
				if start == nil {
					tree.Head = node
				} else {
					node.Children[L] = center.Children[L]
				}

				node.Children[R] = center
				center.Children[L] = node

				tree.Size++
				if tree.Compare(key, tree.Center.Key) > 0 {
					tree.HighSize++
				} else {
					tree.LowSize++
				}

				offset := tree.HighSize - tree.LowSize
				if abs(offset) > 1 {
					log.Println(tree.debugString())
					tree.pickCenter(offset)
					log.Println(tree.debugString())
				}

				return true
			}

			end = center
			center = center.Children[L]

		case c > 0:

			if center.Children[R] == end {

				node := &Node{Key: key, Value: value}
				if end == nil {
					tree.Tail = node
				} else {
					node.Children[R] = center.Children[R]
				}

				node.Children[L] = center
				center.Children[R] = node

				tree.Size++
				if tree.Compare(key, tree.Center.Key) > 0 {
					tree.HighSize++
				} else {
					tree.LowSize++
				}

				offset := tree.HighSize - tree.LowSize
				if abs(offset) > 1 {
					log.Println(tree.debugString())
					tree.pickCenter(offset)
					log.Println(tree.debugString())
				}

				return true
			}

			start = center
			center = center.Children[R]
		default:
			return false
		}
	}
}

func (tree *Tree) pickCenter(offset int64) {
	const L = 0
	const R = 1

	if offset > 0 {
		center := tree.Center
		tree.Center = tree.pickRight(center.Children[R], center)

		tree.LowSize++
		tree.HighSize--
	} else {
		tree.Center = tree.pickLeft(tree.Center.Children[L], tree.Center)
		tree.LowSize--
		tree.HighSize++
	}
	log.Println(tree.Center, tree.Center.Children[L], tree.Center.Children[R])

}

func (tree *Tree) pickRight(center, start *Node) *Node {

	if center.Children[0] == start {
		return center
	}

	return tree.pickRight(center.Children[0], start)
}

func (tree *Tree) pickLeft(center, end *Node) *Node {
	if center.Children[1] == end {
		return center
	}

	return tree.pickLeft(center.Children[1], end)
}
