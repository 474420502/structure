package arraytree

import (
	"math/rand"
	"testing"
)

type Node struct {
	Parent   *Node
	Children [2]*Node
	Val      int
}

type SortedStack struct {
	Root *Node
	Top  *Node
}

func Constructor() SortedStack {
	return SortedStack{}
}

func (this *SortedStack) Push(val int) {

	if this.Root == nil {
		this.Root = &Node{Val: val}
		this.Top = this.Root
		return
	}

	if this.Top.Val > val {
		node := &Node{Val: val, Parent: this.Top}
		this.Top.Children[0] = node
		this.Top = node
		return
	}

	cur := this.Root
	for {
		if val < cur.Val {
			if cur.Children[0] == nil {
				node := &Node{Val: val, Parent: cur}
				cur.Children[0] = node
				return
			}
			cur = cur.Children[0]
		} else {
			if cur.Children[1] == nil {
				cur.Children[1] = &Node{Val: val, Parent: cur}
				return
			}
			cur = cur.Children[1]
		}
	}
}

func (this *SortedStack) Pop() {
	if this.Root == nil {
		return
	}

	var parent *Node = this.Top.Parent
	var next *Node
	if this.Top.Children[0] == nil {
		next = this.Top.Children[1]
	} else {
		next = this.Top.Children[0]
	}

	if parent == nil {
		this.Root = next
	} else {
		parent.Children[0] = next
	}

	if next != nil {
		next.Parent = parent
		for next.Children[0] != nil {
			next = next.Children[0]
		}
		this.Top = next
	} else {
		if parent == nil {
			this.Top = this.Root
		} else {
			this.Top = parent
		}
	}
}

func (this *SortedStack) Peek() int {
	if this.Top == nil {
		return -1
	}
	return this.Top.Val
}

func (this *SortedStack) IsEmpty() bool {
	return this.Root == nil
}

var ops = []string{"SortedStack", "peek", "peek", "pop", "isEmpty", "peek", "push", "pop", "push", "peek", "push", "peek", "pop", "pop", "push", "isEmpty", "push", "peek", "isEmpty", "push", "peek", "peek", "isEmpty", "push", "isEmpty", "peek", "isEmpty", "pop", "peek", "pop", "push", "push", "isEmpty", "pop", "isEmpty", "peek", "push", "pop", "push", "push", "isEmpty", "pop", "pop", "push", "peek", "isEmpty", "pop", "peek", "push", "push", "peek", "isEmpty", "isEmpty", "isEmpty", "isEmpty", "isEmpty", "push", "push", "push", "push", "push", "peek", "peek", "isEmpty", "push"}

func BenchmarkCase1(t *testing.B) {
	stack := Constructor()

	for n := 0; n < t.N; n++ {
		stack.Push(rand.Intn(100000))
	}
}
