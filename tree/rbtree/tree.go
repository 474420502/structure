package rbtree

import "github.com/474420502/structure/compare"

const (
	red   = true
	black = false
)

type Node[KEY any, VALUE any] struct {
	Key    KEY
	Value  VALUE
	Color  bool
	Left   *Node[KEY, VALUE]
	Right  *Node[KEY, VALUE]
	Parent *Node[KEY, VALUE]
}

type Tree[KEY any, VALUE any] struct {
	root            *Node[KEY, VALUE]
	compare         compare.Compare[KEY]
	zero            VALUE
	size            int
	singleRotations int
}

func New[KEY any, VALUE any](comp compare.Compare[KEY]) *Tree[KEY, VALUE] {
	return &Tree[KEY, VALUE]{compare: comp}
}

func (tree *Tree[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	return tree.insert(key, value, false)
}

func (tree *Tree[KEY, VALUE]) InsertIfAbsent(key KEY, value VALUE) bool {
	return tree.Put(key, value)
}

func (tree *Tree[KEY, VALUE]) Set(key KEY, value VALUE) bool {
	return tree.insert(key, value, true)
}

func (tree *Tree[KEY, VALUE]) Upsert(key KEY, value VALUE) bool {
	if current := tree.getNode(key); current != nil {
		current.Value = value
		return true
	}
	tree.insert(key, value, false)
	return false
}

func (tree *Tree[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	if current := tree.getNode(key); current != nil {
		return current.Value, true
	}
	return tree.zero, false
}

func (tree *Tree[KEY, VALUE]) Remove(key KEY) (VALUE, bool) {
	target := tree.getNode(key)
	if target == nil {
		return tree.zero, false
	}

	removedValue := target.Value
	y := target
	yOriginalColor := y.Color
	var replacement *Node[KEY, VALUE]
	var replacementParent *Node[KEY, VALUE]

	if target.Left == nil {
		replacement = target.Right
		replacementParent = target.Parent
		tree.transplant(target, target.Right)
	} else if target.Right == nil {
		replacement = target.Left
		replacementParent = target.Parent
		tree.transplant(target, target.Left)
	} else {
		y = minimum(target.Right)
		yOriginalColor = y.Color
		replacement = y.Right
		if y.Parent == target {
			replacementParent = y
			if replacement != nil {
				replacement.Parent = y
			}
		} else {
			replacementParent = y.Parent
			tree.transplant(y, y.Right)
			y.Right = target.Right
			y.Right.Parent = y
		}

		tree.transplant(target, y)
		y.Left = target.Left
		y.Left.Parent = y
		y.Color = target.Color
	}

	tree.size--
	if yOriginalColor == black {
		tree.deleteFix(replacement, replacementParent)
	}
	if tree.root != nil {
		tree.root.Color = black
	}

	return removedValue, true
}

func (tree *Tree[KEY, VALUE]) Delete(key KEY) (VALUE, bool) {
	return tree.Remove(key)
}

func (tree *Tree[KEY, VALUE]) Clear() {
	tree.root = nil
	tree.size = 0
	tree.singleRotations = 0
}

func (tree *Tree[KEY, VALUE]) Size() int64 {
	return int64(tree.size)
}

func (tree *Tree[KEY, VALUE]) Len() int {
	return tree.size
}

func (tree *Tree[KEY, VALUE]) Height() int {
	height, _, _, _ := tree.shapeStats()
	return height
}

func (tree *Tree[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE] {
	return &Iterator[KEY, VALUE]{tree: tree}
}

func (tree *Tree[KEY, VALUE]) Traverse(every func(KEY, VALUE) bool) {
	var walk func(current *Node[KEY, VALUE]) bool
	walk = func(current *Node[KEY, VALUE]) bool {
		if current == nil {
			return true
		}
		if !walk(current.Left) {
			return false
		}
		if !every(current.Key, current.Value) {
			return false
		}
		return walk(current.Right)
	}
	walk(tree.root)
}

func (tree *Tree[KEY, VALUE]) Values() []VALUE {
	if tree.size == 0 {
		return nil
	}
	result := make([]VALUE, 0, tree.size)
	tree.Traverse(func(_ KEY, value VALUE) bool {
		result = append(result, value)
		return true
	})
	return result
}

func (tree *Tree[KEY, VALUE]) ResetBenchmarkStats() {
	tree.singleRotations = 0
}

func (tree *Tree[KEY, VALUE]) BenchmarkStats() BenchmarkStats {
	height, avgDepth, p50Depth, p95Depth := tree.shapeStats()
	return BenchmarkStats{
		SingleRotations: tree.singleRotations,
		DoubleRotations: 0,
		Height:          height,
		AvgDepth:        avgDepth,
		P50Depth:        p50Depth,
		P95Depth:        p95Depth,
	}
}

func (tree *Tree[KEY, VALUE]) insert(key KEY, value VALUE, overwrite bool) bool {
	var parent *Node[KEY, VALUE]
	current := tree.root

	for current != nil {
		parent = current
		cmp := tree.compare(key, current.Key)
		switch {
		case cmp < 0:
			current = current.Left
		case cmp > 0:
			current = current.Right
		default:
			if overwrite {
				current.Value = value
			}
			return false
		}
	}

	inserted := &Node[KEY, VALUE]{
		Key:    key,
		Value:  value,
		Color:  red,
		Parent: parent,
	}

	if parent == nil {
		tree.root = inserted
	} else if tree.compare(key, parent.Key) < 0 {
		parent.Left = inserted
	} else {
		parent.Right = inserted
	}

	tree.size++
	tree.insertFix(inserted)
	tree.root.Color = black
	return true
}

func (tree *Tree[KEY, VALUE]) insertFix(current *Node[KEY, VALUE]) {
	for current != tree.root && colorOf(current.Parent) == red {
		if current.Parent == current.Parent.Parent.Left {
			uncle := current.Parent.Parent.Right
			if colorOf(uncle) == red {
				setColor(current.Parent, black)
				setColor(uncle, black)
				setColor(current.Parent.Parent, red)
				current = current.Parent.Parent
				continue
			}

			if current == current.Parent.Right {
				current = current.Parent
				tree.leftRotate(current)
			}

			setColor(current.Parent, black)
			setColor(current.Parent.Parent, red)
			tree.rightRotate(current.Parent.Parent)
		} else {
			uncle := current.Parent.Parent.Left
			if colorOf(uncle) == red {
				setColor(current.Parent, black)
				setColor(uncle, black)
				setColor(current.Parent.Parent, red)
				current = current.Parent.Parent
				continue
			}

			if current == current.Parent.Left {
				current = current.Parent
				tree.rightRotate(current)
			}

			setColor(current.Parent, black)
			setColor(current.Parent.Parent, red)
			tree.leftRotate(current.Parent.Parent)
		}
	}
}

func (tree *Tree[KEY, VALUE]) deleteFix(current *Node[KEY, VALUE], parent *Node[KEY, VALUE]) {
	for current != tree.root && colorOf(current) == black {
		if parent == nil {
			break
		}

		if current == parent.Left {
			sibling := parent.Right
			if colorOf(sibling) == red {
				setColor(sibling, black)
				setColor(parent, red)
				tree.leftRotate(parent)
				sibling = parent.Right
			}

			if colorOf(leftOf(sibling)) == black && colorOf(rightOf(sibling)) == black {
				setColor(sibling, red)
				current = parent
				parent = current.Parent
				continue
			}

			if colorOf(rightOf(sibling)) == black {
				setColor(leftOf(sibling), black)
				setColor(sibling, red)
				tree.rightRotate(sibling)
				sibling = parent.Right
			}

			setColor(sibling, colorOf(parent))
			setColor(parent, black)
			setColor(rightOf(sibling), black)
			tree.leftRotate(parent)
			current = tree.root
			parent = nil
		} else {
			sibling := parent.Left
			if colorOf(sibling) == red {
				setColor(sibling, black)
				setColor(parent, red)
				tree.rightRotate(parent)
				sibling = parent.Left
			}

			if colorOf(rightOf(sibling)) == black && colorOf(leftOf(sibling)) == black {
				setColor(sibling, red)
				current = parent
				parent = current.Parent
				continue
			}

			if colorOf(leftOf(sibling)) == black {
				setColor(rightOf(sibling), black)
				setColor(sibling, red)
				tree.leftRotate(sibling)
				sibling = parent.Left
			}

			setColor(sibling, colorOf(parent))
			setColor(parent, black)
			setColor(leftOf(sibling), black)
			tree.rightRotate(parent)
			current = tree.root
			parent = nil
		}
	}
	setColor(current, black)
}

func (tree *Tree[KEY, VALUE]) leftRotate(pivot *Node[KEY, VALUE]) {
	if pivot == nil || pivot.Right == nil {
		return
	}

	right := pivot.Right
	pivot.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = pivot
	}

	right.Parent = pivot.Parent
	if pivot.Parent == nil {
		tree.root = right
	} else if pivot == pivot.Parent.Left {
		pivot.Parent.Left = right
	} else {
		pivot.Parent.Right = right
	}

	right.Left = pivot
	pivot.Parent = right
	tree.singleRotations++
}

func (tree *Tree[KEY, VALUE]) rightRotate(pivot *Node[KEY, VALUE]) {
	if pivot == nil || pivot.Left == nil {
		return
	}

	left := pivot.Left
	pivot.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = pivot
	}

	left.Parent = pivot.Parent
	if pivot.Parent == nil {
		tree.root = left
	} else if pivot == pivot.Parent.Left {
		pivot.Parent.Left = left
	} else {
		pivot.Parent.Right = left
	}

	left.Right = pivot
	pivot.Parent = left
	tree.singleRotations++
}

func (tree *Tree[KEY, VALUE]) transplant(target *Node[KEY, VALUE], replacement *Node[KEY, VALUE]) {
	if target.Parent == nil {
		tree.root = replacement
	} else if target == target.Parent.Left {
		target.Parent.Left = replacement
	} else {
		target.Parent.Right = replacement
	}
	if replacement != nil {
		replacement.Parent = target.Parent
	}
}

func (tree *Tree[KEY, VALUE]) getNode(key KEY) *Node[KEY, VALUE] {
	current := tree.root
	for current != nil {
		cmp := tree.compare(key, current.Key)
		switch {
		case cmp < 0:
			current = current.Left
		case cmp > 0:
			current = current.Right
		default:
			return current
		}
	}
	return nil
}

func minimum[KEY any, VALUE any](current *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}

func maximum[KEY any, VALUE any](current *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	for current != nil && current.Right != nil {
		current = current.Right
	}
	return current
}

func successor[KEY any, VALUE any](current *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	if current == nil {
		return nil
	}
	if current.Right != nil {
		return minimum(current.Right)
	}
	parent := current.Parent
	for parent != nil && current == parent.Right {
		current = parent
		parent = parent.Parent
	}
	return parent
}

func predecessor[KEY any, VALUE any](current *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	if current == nil {
		return nil
	}
	if current.Left != nil {
		return maximum(current.Left)
	}
	parent := current.Parent
	for parent != nil && current == parent.Left {
		current = parent
		parent = parent.Parent
	}
	return parent
}

func colorOf[KEY any, VALUE any](current *Node[KEY, VALUE]) bool {
	if current == nil {
		return black
	}
	return current.Color
}

func setColor[KEY any, VALUE any](current *Node[KEY, VALUE], color bool) {
	if current != nil {
		current.Color = color
	}
}

func leftOf[KEY any, VALUE any](current *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	if current == nil {
		return nil
	}
	return current.Left
}

func rightOf[KEY any, VALUE any](current *Node[KEY, VALUE]) *Node[KEY, VALUE] {
	if current == nil {
		return nil
	}
	return current.Right
}