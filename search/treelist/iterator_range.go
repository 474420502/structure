package treelist

type RangeDirection int

const (
	// Forward start to end
	Forward RangeDirection = 0
	// Reverse end to start
	Reverse RangeDirection = 1
)

// IteratorRange the iterator for easy to range the data
type IteratorRange struct {
	tree         *Tree
	siter, eiter nodePoint
	dir          RangeDirection
}

type SliceIndex struct {
	*Slice
	Index int64
}

// SetDirection set iterator range direction. default Forward(start to end)
func (ir *IteratorRange) Range(do func(cur *SliceIndex) bool) {
	const (
		L = 0
		R = 1
	)

	var cur *treeNode
	var end *treeNode
	var dir int
	var idx int64

	if ir.dir == Forward {
		cur = ir.siter.cur
		idx = ir.siter.idx
		end = ir.eiter.cur
		dir = R
	} else {
		cur = ir.eiter.cur
		idx = ir.eiter.idx
		end = ir.siter.cur
		dir = L
	}

	if cur != nil {
		for {
			if !do(&SliceIndex{Slice: &cur.Slice, Index: idx}) || cur == end {
				break
			}
			cur = cur.Direct[dir]
			idx++
		}
	}

}

// SetDirection set iterator range direction. default Forward(start to end)
func (ir *IteratorRange) SetDirection(dir RangeDirection) {
	ir.dir = dir
}

// SetDirection set iterator range direction
func (ir *IteratorRange) Direction() RangeDirection {
	return ir.dir
}

// Size get range size
func (ir *IteratorRange) Size() int64 {
	if ir.dir == Forward {
		return ir.eiter.idx - ir.siter.idx
	}
	return ir.siter.idx - ir.eiter.idx
}

// GE2LE [s,e] start with GE, end with LE. (like Seek**)
func (ir *IteratorRange) GE2LE(start, end []byte) {

	const (
		L = 0
		R = 1
	)

	cur, idx, dir := ir.tree.seekNodeWithIndex(start)
	ir.siter.idx = idx
	if dir > 0 {
		cur = cur.Direct[R]
		ir.siter.idx++
	}
	ir.siter.cur = cur

	cur, idx, dir = ir.tree.seekNodeWithIndex(end)
	ir.eiter.idx = idx
	if dir < 0 {
		cur = cur.Direct[L]
		ir.eiter.idx--
	}
	ir.eiter.cur = cur

}

// GE2LE (s,e] start with GT, end with LE. (like Seek**)
func (ir *IteratorRange) GT2LE(start, end []byte) {

	const (
		L = 0
		R = 1
	)

	cur, idx, dir := ir.tree.seekNodeWithIndex(start)
	ir.siter.idx = idx
	if dir >= 0 {
		cur = cur.Direct[R]
		ir.siter.idx++
	}
	ir.siter.cur = cur

	cur, idx, dir = ir.tree.seekNodeWithIndex(end)
	ir.eiter.idx = idx
	if dir < 0 {
		cur = cur.Direct[L]
		ir.eiter.idx--
	}
	ir.eiter.cur = cur

}

// GE2LT [s,e) start with GE, end with LT. (like Seek**)
func (ir *IteratorRange) GE2LT(start, end []byte) {

	const (
		L = 0
		R = 1
	)

	cur, idx, dir := ir.tree.seekNodeWithIndex(start)
	ir.siter.idx = idx
	if dir > 0 {
		cur = cur.Direct[R]
		ir.siter.idx++
	}
	ir.siter.cur = cur

	cur, idx, dir = ir.tree.seekNodeWithIndex(end)
	ir.eiter.idx = idx
	if dir <= 0 {
		cur = cur.Direct[L]
		ir.eiter.idx--
	}
	ir.eiter.cur = cur

}

// GE2LT (s,e) start with GT, end with LT. (like Seek**)
func (ir *IteratorRange) GT2LT(start, end []byte) {

	const (
		L = 0
		R = 1
	)

	cur, idx, dir := ir.tree.seekNodeWithIndex(start)
	ir.siter.idx = idx
	if dir >= 0 {
		cur = cur.Direct[R]
		ir.siter.idx++
	}
	ir.siter.cur = cur

	cur, idx, dir = ir.tree.seekNodeWithIndex(end)
	ir.eiter.idx = idx
	if dir <= 0 {
		cur = cur.Direct[L]
		ir.eiter.idx--
	}
	ir.eiter.cur = cur

}
