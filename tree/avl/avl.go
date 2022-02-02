package avl

import (
	"fmt"

	"github.com/474420502/structure/compare"
)

const HeightDiff = 1

type Node[T any] struct {
	Children [2]*Node[T]
	Parent   *Node[T]
	height   int
	Key      T
	Value    interface{}
}

func (n *Node[T]) String() string {
	if n == nil {
		return "nil"
	}

	p := "nil"
	if n.Parent != nil {
		p = fmt.Sprintf("%v", n.Parent.Value)
	}
	return fmt.Sprintf("%v", n.Value) + "(" + p + "|" + fmt.Sprintf("%v", n.height) + ")"
}

type Tree[T any] struct {
	Root    *Node[T]
	size    int64
	Compare compare.Compare[T]
}

func New[T any](Compare compare.Compare[T]) *Tree[T] {
	return &Tree[T]{Compare: Compare}
}

func (tree *Tree[T]) String() string {
	if tree.size == 0 {
		return ""
	}
	str := "AVLTree\n"
	output(tree.Root, "", true, &str)

	return str
}

func (tree *Tree[T]) Size() int64 {
	return tree.size
}

func (tree *Tree[T]) Height() int {
	if tree.Root == nil {
		return 0
	}
	return tree.Root.height + 1
}

// Iterator must call Seek*.
func (tree *Tree[T]) Iterator() *Iterator[T] {
	return newIterator(tree)
}

func (tree *Tree[T]) Remove(key T) (interface{}, bool) {

	if n, ok := tree.getNode(key); ok {

		tree.size--
		if tree.size == 0 {
			tree.Root = nil
			return n.Value, true
		}

		left := getHeight(n.Children[0])
		right := getHeight(n.Children[1])

		if left == -1 && right == -1 {
			p := n.Parent
			p.Children[getRelationship(n)] = nil
			tree.fixRemoveHeight(p)
			return n.Value, true
		}

		var cur *Node[T]
		if left > right {
			cur = n.Children[0]
			for cur.Children[1] != nil {
				cur = cur.Children[1]
			}

			cleft := cur.Children[0]
			cur.Parent.Children[getRelationship(cur)] = cleft
			if cleft != nil {
				cleft.Parent = cur.Parent
			}

		} else {
			cur = n.Children[1]
			for cur.Children[0] != nil {
				cur = cur.Children[0]
			}

			cright := cur.Children[1]
			cur.Parent.Children[getRelationship(cur)] = cright

			if cright != nil {
				cright.Parent = cur.Parent
			}
		}

		cparent := cur.Parent
		// 修改为interface 交换
		n.Value, cur.Value = cur.Value, n.Value
		n.Key, cur.Key = cur.Key, n.Key

		// 考虑到刚好替换的节点是 被替换节点的孩子节点的时候, 从自身修复高度
		if cparent == n {
			tree.fixRemoveHeight(n)
		} else {
			tree.fixRemoveHeight(cparent)
		}

		return cur.Value, true
	}

	return nil, false
}

func (tree *Tree[T]) Clear() {
	tree.size = 0
	tree.Root = nil
}

// Values 返回先序遍历的值
func (tree *Tree[T]) Values() []interface{} {
	var mszie int64 = 0
	if tree.Root != nil {
		mszie = tree.size
	}
	result := make([]interface{}, 0, mszie)
	tree.Traverse(func(k T, v interface{}) bool {
		result = append(result, v)
		return true
	})
	return result
}

func (tree *Tree[T]) Get(key T) (interface{}, bool) {
	n, ok := tree.getNode(key)
	if ok {
		return n.Value, true
	}
	return n, false
}

func (tree *Tree[T]) getNode(key T) (*Node[T], bool) {

	for n := tree.Root; n != nil; {
		switch c := tree.Compare(key, n.Key); c {
		case -1:
			n = n.Children[0]
		case 1:
			n = n.Children[1]
		case 0:
			return n, true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}
	return nil, false
}

// Set Put And value Set Node.Value
func (tree *Tree[T]) Set(key T, value interface{}) bool {

	if tree.size == 0 {
		tree.size++
		tree.Root = &Node[T]{Key: key, Value: value}
		return false
	}

	for cur, c := tree.Root, 0; ; {
		c = tree.Compare(key, cur.Key)
		if c == -1 {
			if cur.Children[0] == nil {
				tree.size++
				cur.Children[0] = &Node[T]{Key: key, Value: value}
				cur.Children[0].Parent = cur
				if cur.height == 0 {
					tree.fixPutHeight(cur)
				}
				return false
			}
			cur = cur.Children[0]
		} else if c == 1 {
			if cur.Children[1] == nil {
				tree.size++
				cur.Children[1] = &Node[T]{Key: key, Value: value}
				cur.Children[1].Parent = cur
				if cur.height == 0 {
					tree.fixPutHeight(cur)
				}
				return false
			}
			cur = cur.Children[1]
		} else {
			cur.Key = key
			cur.Value = value
			return true
		}
	}
}

// Put Put into Tree . if key exists, not cover. and return false. else return true
func (tree *Tree[T]) Put(key T, value interface{}) bool {

	if tree.size == 0 {
		tree.size++
		tree.Root = &Node[T]{Key: key, Value: value}
		return true
	}

	for cur, c := tree.Root, 0; ; {
		c = tree.Compare(key, cur.Key)
		if c == -1 {
			if cur.Children[0] == nil {
				tree.size++
				cur.Children[0] = &Node[T]{Key: key, Value: value}
				cur.Children[0].Parent = cur
				if cur.height == 0 {
					tree.fixPutHeight(cur)
				}
				return true
			}
			cur = cur.Children[0]
		} else if c == 1 {
			if cur.Children[1] == nil {
				tree.size++
				cur.Children[1] = &Node[T]{Key: key, Value: value}
				cur.Children[1].Parent = cur
				if cur.height == 0 {
					tree.fixPutHeight(cur)
				}
				return true
			}
			cur = cur.Children[1]
		} else {
			return false
		}
	}
}

// type TraversalMethod int

// const (
// 	// L = left R = right D = Value(dest)
// 	_ TraversalMethod = iota
// 	//DLR 先值 然后左递归 右递归 下面同理
// 	DLR
// 	//LDR 先从左边有序访问到右边 从小到大
// 	LDR
// 	// LRD 同理
// 	LRD
// 	// DRL 同理
// 	DRL
// 	// RDL 先从右边有序访问到左边 从大到小
// 	RDL
// 	// RLD 同理
// 	RLD
// )

// Traverse 遍历的方法 默认是LDR 从小到大 Compare 为 l < r
func (tree *Tree[T]) Traverse(every func(k T, v interface{}) bool) {
	if tree.Root == nil {
		return
	}

	var traverasl func(cur *Node[T]) bool
	traverasl = func(cur *Node[T]) bool {
		if cur == nil {
			return true
		}
		if !traverasl(cur.Children[0]) {
			return false
		}
		if !every(cur.Key, cur.Value) {
			return false
		}
		if !traverasl(cur.Children[1]) {
			return false
		}
		return true
	}
	traverasl(tree.Root)
}

func (tree *Tree[T]) lrrotate(cur *Node[T]) {

	const l = 1
	const r = 0

	movparent := cur.Children[l]
	mov := movparent.Children[r]

	mov.Value, cur.Value = cur.Value, mov.Value //交换值达到, 相对位移
	mov.Key, cur.Key = cur.Key, mov.Key

	if mov.Children[l] != nil {
		movparent.Children[r] = mov.Children[l]
		movparent.Children[r].Parent = movparent
		//movparent.children[r].child = l
	} else {
		movparent.Children[r] = nil
	}

	if mov.Children[r] != nil {
		mov.Children[l] = mov.Children[r]
		//mov.children[l].child = l
	} else {
		mov.Children[l] = nil
	}

	if cur.Children[r] != nil {
		mov.Children[r] = cur.Children[r]
		mov.Children[r].Parent = mov
	} else {
		mov.Children[r] = nil
	}

	cur.Children[r] = mov
	mov.Parent = cur

	mov.height = getMaxChildrenHeight(mov) + 1
	movparent.height = getMaxChildrenHeight(movparent) + 1
	cur.height = getMaxChildrenHeight(cur) + 1
}

func (tree *Tree[T]) rlrotate(cur *Node[T]) {

	const l = 0
	const r = 1

	movparent := cur.Children[l]
	mov := movparent.Children[r]

	mov.Value, cur.Value = cur.Value, mov.Value //交换值达到, 相对位移
	mov.Key, cur.Key = cur.Key, mov.Key

	if mov.Children[l] != nil {
		movparent.Children[r] = mov.Children[l]
		movparent.Children[r].Parent = movparent
	} else {
		movparent.Children[r] = nil
	}

	if mov.Children[r] != nil {
		mov.Children[l] = mov.Children[r]
	} else {
		mov.Children[l] = nil
	}

	if cur.Children[r] != nil {
		mov.Children[r] = cur.Children[r]
		mov.Children[r].Parent = mov
	} else {
		mov.Children[r] = nil
	}

	cur.Children[r] = mov
	mov.Parent = cur

	mov.height = getMaxChildrenHeight(mov) + 1
	movparent.height = getMaxChildrenHeight(movparent) + 1
	cur.height = getMaxChildrenHeight(cur) + 1
}

func (tree *Tree[T]) rrotate(cur *Node[T]) {

	const l = 0
	const r = 1
	// 1 right 0 left
	mov := cur.Children[l]

	mov.Value, cur.Value = cur.Value, mov.Value //交换值达到, 相对位移
	mov.Key, cur.Key = cur.Key, mov.Key

	//  mov.children[l]不可能为nil
	mov.Children[l].Parent = cur
	cur.Children[l] = mov.Children[l]

	// 解决mov节点孩子转移的问题
	if mov.Children[r] != nil {
		mov.Children[l] = mov.Children[r]
	} else {
		mov.Children[l] = nil
	}

	if cur.Children[r] != nil {
		mov.Children[r] = cur.Children[r]
		mov.Children[r].Parent = mov
	} else {
		mov.Children[r] = nil
	}

	// 连接转移后的节点 由于mov只是与cur交换值,parent不变
	cur.Children[r] = mov

	mov.height = getMaxChildrenHeight(mov) + 1
	cur.height = getMaxChildrenHeight(cur) + 1
}

func (tree *Tree[T]) lrotate(cur *Node[T]) {

	const l = 1
	const r = 0

	mov := cur.Children[l]

	mov.Value, cur.Value = cur.Value, mov.Value //交换值达到, 相对位移
	mov.Key, cur.Key = cur.Key, mov.Key

	// 不可能为nil
	mov.Children[l].Parent = cur
	cur.Children[l] = mov.Children[l]

	if mov.Children[r] != nil {
		mov.Children[l] = mov.Children[r]
	} else {
		mov.Children[l] = nil
	}

	if cur.Children[r] != nil {
		mov.Children[r] = cur.Children[r]
		mov.Children[r].Parent = mov
	} else {
		mov.Children[r] = nil
	}

	cur.Children[r] = mov

	mov.height = getMaxChildrenHeight(mov) + 1
	cur.height = getMaxChildrenHeight(cur) + 1
}

func getMaxAndChildrenHeight[T any](cur *Node[T]) (h1, h2, maxh int) {
	h1 = getHeight(cur.Children[0])
	h2 = getHeight(cur.Children[1])
	if h1 > h2 {
		maxh = h1
	} else {
		maxh = h2
	}

	return
}

func getMaxChildrenHeight[T any](cur *Node[T]) int {
	h1 := getHeight(cur.Children[0])
	h2 := getHeight(cur.Children[1])
	if h1 > h2 {
		return h1
	}
	return h2
}

func getHeight[T any](cur *Node[T]) int {
	if cur == nil {
		return -1
	}
	return cur.height
}

func (tree *Tree[T]) fixRemoveHeight(cur *Node[T]) {
	for {

		lefth, rigthh, lrmax := getMaxAndChildrenHeight(cur)

		// 判断当前节点是否有变化, 如果没变化的时候, 不需要往上修复
		curheight := lrmax + 1
		cur.height = curheight

		// 计算高度的差值 绝对值大于2的时候需要旋转
		diff := lefth - rigthh
		if diff < -HeightDiff {
			r := cur.Children[1] // 根据左旋转的右边节点的子节点 左右高度选择旋转的方式
			if getHeight(r.Children[0]) > getHeight(r.Children[1]) {
				tree.lrrotate(cur)
			} else {
				tree.lrotate(cur)
			}
		} else if diff > HeightDiff {
			l := cur.Children[0]
			if getHeight(l.Children[1]) > getHeight(l.Children[0]) {
				tree.rlrotate(cur)
			} else {
				tree.rrotate(cur)
			}
		} else {
			if cur.height == curheight {
				return
			}
		}

		if cur.Parent == nil {
			return
		}

		cur = cur.Parent
	}

}

func (tree *Tree[T]) fixPutHeight(cur *Node[T]) {

	for {

		lefth := getHeight(cur.Children[0])
		rigthh := getHeight(cur.Children[1])

		// 计算高度的差值 绝对值大于2的时候需要旋转
		diff := lefth - rigthh
		if diff < -HeightDiff {
			r := cur.Children[1] // 根据左旋转的右边节点的子节点 左右高度选择旋转的方式
			if getHeight(r.Children[0]) > getHeight(r.Children[1]) {
				tree.lrrotate(cur)
			} else {
				tree.lrotate(cur)
			}
		} else if diff > HeightDiff {
			l := cur.Children[0]
			if getHeight(l.Children[1]) > getHeight(l.Children[0]) {
				tree.rlrotate(cur)
			} else {
				tree.rrotate(cur)
			}

		} else {
			// 选择一个child的最大高度 + 1为 高度
			if lefth > rigthh {
				cur.height = lefth + 1
			} else {
				cur.height = rigthh + 1
			}
		}

		if cur.Parent == nil || cur.height < cur.Parent.height {
			return
		}
		cur = cur.Parent
	}
}

func output[T any](node *Node[T], prefix string, isTail bool, str *string) {

	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	*str += fmt.Sprintf("%v", node.Key) + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Children[0], newPrefix, true, str)
	}

}

func outputfordebug[T any](node *Node[T], prefix string, isTail bool, str *string) {

	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		outputfordebug(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	suffix := "("
	parentv := ""
	if node.Parent == nil {
		parentv = "nil"
	} else {
		parentv = fmt.Sprintf("%v", node.Parent.Value)
	}
	suffix += parentv + "|" + fmt.Sprintf("%v", node.height) + ")"
	*str += fmt.Sprintf("%v", node.Value) + suffix + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		outputfordebug(node.Children[0], newPrefix, true, str)
	}
}

func (tree *Tree[T]) debugString() string {
	if tree.size == 0 {
		return ""
	}
	str := "AVLTree\n"
	outputfordebug(tree.Root, "", true, &str)
	return str
}

func getRelationship[T any](cur *Node[T]) int {
	if cur.Parent.Children[1] == cur {
		return 1
	}
	return 0
}
