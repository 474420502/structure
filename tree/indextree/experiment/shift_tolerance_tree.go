package experiment

import (
	"sort"

	"github.com/474420502/structure/compare"
)

type ShiftToleranceTree[KEY, VALUE any] struct {
	Center           *ShiftToleranceNode[KEY, VALUE]
	Compare          compare.Compare[KEY]
	size             uint
	zero             VALUE
	differenceHeight int8
	singleRotations  int
	doubleRotations  int
	shift            int64
}

type ShiftToleranceNode[KEY any, VALUE any] struct {
	Key      KEY
	Value    VALUE
	Height   int8
	Size     int64
	Balance  int64
	Parent   *ShiftToleranceNode[KEY, VALUE]
	Children [2]*ShiftToleranceNode[KEY, VALUE]
}

type heightLimit struct {
	rootsize   int64
	bottomsize int64
}

func NewShiftTolerance[KEY, VALUE any](cmp compare.Compare[KEY], shift int64) *ShiftToleranceTree[KEY, VALUE] {
	tree := &ShiftToleranceTree[KEY, VALUE]{
		Center:           &ShiftToleranceNode[KEY, VALUE]{Height: 0},
		Compare:          cmp,
		differenceHeight: 1,
		shift:            shift,
	}
	tree.Center.Children[0] = tree.Center
	return tree
}

func NewIndexTreeDefault[KEY, VALUE any](cmp compare.Compare[KEY]) *ShiftToleranceTree[KEY, VALUE] {
	return NewShiftTolerance[KEY, VALUE](cmp, 2)
}

func (tree *ShiftToleranceTree[KEY, VALUE]) Size() uint {
	return tree.size
}

func (tree *ShiftToleranceTree[KEY, VALUE]) Height() int8 {
	if tree.size == 0 {
		return 0
	}
	return tree.getRoot().Height
}

func (tree *ShiftToleranceTree[KEY, VALUE]) getRoot() *ShiftToleranceNode[KEY, VALUE] {
	return tree.Center.Children[1]
}

func (tree *ShiftToleranceTree[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	cur := tree.getRoot()
	if cur == nil {
		tree.Center.Children[1] = &ShiftToleranceNode[KEY, VALUE]{
			Key:    key,
			Value:  value,
			Height: 1,
			Size:   1,
			Parent: tree.Center,
		}
		tree.size++
		return true
	}

	const left = 0
	const right = 1

	for {
		cmp := tree.Compare(cur.Key, key)
		switch {
		case cmp < 0:
			if cur.Children[right] != nil {
				cur = cur.Children[right]
			} else {
				node := &ShiftToleranceNode[KEY, VALUE]{Parent: cur, Key: key, Value: value, Height: 1, Size: 1}
				cur.Children[right] = node
				tree.size++
				tree.fixPut(cur)
				return true
			}
		case cmp > 0:
			if cur.Children[left] != nil {
				cur = cur.Children[left]
			} else {
				node := &ShiftToleranceNode[KEY, VALUE]{Parent: cur, Key: key, Value: value, Height: 1, Size: 1}
				cur.Children[left] = node
				tree.size++
				tree.fixPut(cur)
				return true
			}
		default:
			return false
		}
	}
}

func (tree *ShiftToleranceTree[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	cur := tree.get(key, tree.getRoot())
	if cur == nil {
		return tree.zero, false
	}
	return cur.Value, true
}

func (tree *ShiftToleranceTree[KEY, VALUE]) get(key KEY, cur *ShiftToleranceNode[KEY, VALUE]) *ShiftToleranceNode[KEY, VALUE] {
	if cur == nil {
		return nil
	}
	cmp := tree.Compare(cur.Key, key)
	if cmp < 0 {
		return cur
	}
	return tree.get(key, cur.Children[cmp])
}

func (tree *ShiftToleranceTree[KEY, VALUE]) fixPutSize(cur *ShiftToleranceNode[KEY, VALUE]) {
	for cur != nil && cur != tree.Center {
		cur.Size = getSizeTolerance(cur.Children[0]) + getSizeTolerance(cur.Children[1]) + 1
		cur.updateBalance()
		cur.updateHeight()
		cur = cur.Parent
	}
}

func (tree *ShiftToleranceTree[KEY, VALUE]) fixPut(cur *ShiftToleranceNode[KEY, VALUE]) {
	cur.Size++
	cur.updateBalance()
	cur.updateHeight()
	if cur.Size == 3 {
		tree.fixPutSize(cur.Parent)
		return
	}

	var height int64 = 2
	cur = cur.Parent

	for cur != nil && cur != tree.Center {
		cur.Size++
		cur.updateBalance()
		cur.updateHeight()
		parent := cur.Parent

		limitSize := tree.getHeightLimit(height)
		if cur.Size < limitSize.rootsize {
			balance := cur.Balance
			if balance < 0 {
				if -balance >= limitSize.bottomsize {
					tree.sizeRRotate(cur)
					tree.fixPutSize(parent)
					return
				}
			} else {
				if balance >= limitSize.bottomsize {
					tree.sizeLRotate(cur)
					tree.fixPutSize(parent)
					return
				}
			}
		}

		height++
		cur = parent
	}
}

func (tree *ShiftToleranceTree[KEY, VALUE]) getHeightLimit(height int64) *heightLimit {
	root2nsize := int64(1) << height
	return &heightLimit{
		rootsize:   root2nsize,
		bottomsize: (root2nsize >> tree.shift) + 1,
	}
}

func getSizeTolerance[KEY, VALUE any](cur *ShiftToleranceNode[KEY, VALUE]) int64 {
	if cur == nil {
		return 0
	}
	return cur.Size
}

func getHeightTolerance[KEY, VALUE any](cur *ShiftToleranceNode[KEY, VALUE]) int8 {
	if cur == nil {
		return 0
	}
	return cur.Height
}

func (node *ShiftToleranceNode[KEY, VALUE]) updateHeight() {
	lh := getHeightTolerance(node.Children[0])
	rh := getHeightTolerance(node.Children[1])
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

func (node *ShiftToleranceNode[KEY, VALUE]) updateBalance() {
	node.Balance = getSizeTolerance(node.Children[0]) - getSizeTolerance(node.Children[1])
}

func getChildrenSizeTolerance[KEY, VALUE any](cur *ShiftToleranceNode[KEY, VALUE]) (int64, int64) {
	return getSizeTolerance(cur.Children[0]), getSizeTolerance(cur.Children[1])
}

func (tree *ShiftToleranceTree[KEY, VALUE]) sizeLRotate(cur *ShiftToleranceNode[KEY, VALUE]) {
	const left = 0
	llsize, lrsize := getChildrenSizeTolerance(cur.Children[left])
	if llsize < lrsize {
		tree.doubleRotations++
		tree.lrotate(cur.Children[left])
		tree.rrotate(cur)
		return
	}
	tree.singleRotations++
	tree.rrotate(cur)
}

func (tree *ShiftToleranceTree[KEY, VALUE]) sizeRRotate(cur *ShiftToleranceNode[KEY, VALUE]) {
	const right = 1
	llsize, lrsize := getChildrenSizeTolerance(cur.Children[right])
	if llsize > lrsize {
		tree.doubleRotations++
		tree.rrotate(cur.Children[right])
		tree.lrotate(cur)
		return
	}
	tree.singleRotations++
	tree.lrotate(cur)
}

func (tree *ShiftToleranceTree[KEY, VALUE]) lrotate(cur *ShiftToleranceNode[KEY, VALUE]) *ShiftToleranceNode[KEY, VALUE] {
	const left = 1
	const right = 0
	mov := cur.Children[left]
	movright := mov.Children[right]

	if cur.Parent.Children[left] == cur {
		cur.Parent.Children[left] = mov
	} else {
		cur.Parent.Children[right] = mov
	}
	mov.Parent = cur.Parent

	if movright != nil {
		cur.Children[left] = movright
		movright.Parent = cur
	} else {
		cur.Children[left] = nil
	}

	mov.Children[right] = cur
	cur.Parent = mov

	cur.Size = getSizeTolerance(cur.Children[0]) + getSizeTolerance(cur.Children[1]) + 1
	mov.Size = getSizeTolerance(mov.Children[0]) + getSizeTolerance(mov.Children[1]) + 1
	cur.updateBalance()
	mov.updateBalance()
	cur.updateHeight()
	mov.updateHeight()

	return mov
}

func (tree *ShiftToleranceTree[KEY, VALUE]) rrotate(cur *ShiftToleranceNode[KEY, VALUE]) *ShiftToleranceNode[KEY, VALUE] {
	const left = 0
	const right = 1
	mov := cur.Children[left]
	movright := mov.Children[right]

	if cur.Parent.Children[left] == cur {
		cur.Parent.Children[left] = mov
	} else {
		cur.Parent.Children[right] = mov
	}
	mov.Parent = cur.Parent

	if movright != nil {
		cur.Children[left] = movright
		movright.Parent = cur
	} else {
		cur.Children[left] = nil
	}

	mov.Children[right] = cur
	cur.Parent = mov

	cur.Size = getSizeTolerance(cur.Children[0]) + getSizeTolerance(cur.Children[1]) + 1
	mov.Size = getSizeTolerance(mov.Children[0]) + getSizeTolerance(mov.Children[1]) + 1
	cur.updateBalance()
	mov.updateBalance()
	cur.updateHeight()
	mov.updateHeight()

	return mov
}

func (tree *ShiftToleranceTree[KEY, VALUE]) ResetBenchmarkStats() {
	tree.singleRotations = 0
	tree.doubleRotations = 0
}

func (tree *ShiftToleranceTree[KEY, VALUE]) BenchmarkStats() BenchmarkStats {
	height, avgDepth, p50Depth, p95Depth := tree.shapeStats()
	return BenchmarkStats{
		SingleRotations: tree.singleRotations,
		DoubleRotations: tree.doubleRotations,
		Height:          height,
		AvgDepth:        avgDepth,
		P50Depth:        p50Depth,
		P95Depth:        p95Depth,
	}
}

func (tree *ShiftToleranceTree[KEY, VALUE]) shapeStats() (height int, avgDepth float64, p50Depth int, p95Depth int) {
	root := tree.getRoot()
	if root == nil {
		return 0, 0, 0, 0
	}

	type depthNode struct {
		node  *ShiftToleranceNode[KEY, VALUE]
		depth int
	}

	queue := []depthNode{{node: root, depth: 1}}
	totalDepth := 0
	count := 0
	depths := make([]int, 0, tree.size)

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		count++
		totalDepth += item.depth
		depths = append(depths, item.depth)
		if item.depth > height {
			height = item.depth
		}

		if left := item.node.Children[0]; left != nil {
			queue = append(queue, depthNode{node: left, depth: item.depth + 1})
		}
		if right := item.node.Children[1]; right != nil {
			queue = append(queue, depthNode{node: right, depth: item.depth + 1})
		}
	}

	sort.Ints(depths)
	p50Depth = depths[(len(depths)-1)/2]
	p95Depth = depths[((len(depths)-1)*95)/100]

	return height, float64(totalDepth) / float64(count), p50Depth, p95Depth
}
