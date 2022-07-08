package treelist

type RangeDirection int

const (
	// Forward start to end
	Forward RangeDirection = 0
	// Reverse end KEYo start
	Reverse RangeDirection = 1
)

// IteratorRange the iterator for easy to range the data
type IteratorRange[KEY any, VALUE any] struct {
	tree         *Tree[KEY, VALUE]
	siter, eiter hNode[KEY, VALUE]
	dir          RangeDirection
}

type SliceIndex[KEY any, VALUE any] struct {
	*Slice[KEY, VALUE]
	Index int64
}

// SetDirection set iterator range direction. default Forward(start to end)
func (ir *IteratorRange[KEY, VALUE]) Range(do func(cur *SliceIndex[KEY, VALUE]) bool) {

	if ir.siter.idx > ir.eiter.idx {
		return
	}

	const (
		L = 0
		R = 1
	)

	var cur *treeNode[KEY, VALUE]
	var end *treeNode[KEY, VALUE]
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

	for {
		if !do(&SliceIndex[KEY, VALUE]{Slice: &cur.Slice, Index: idx}) || cur == end {
			break
		}
		cur = cur.Direct[dir]
		idx++
	}

}

// SetDirection set iterator range direction. default Forward(start to end)
func (ir *IteratorRange[KEY, VALUE]) SetDirection(dir RangeDirection) {
	ir.dir = dir
}

// SetDirection set iterator range direction
func (ir *IteratorRange[KEY, VALUE]) Direction() RangeDirection {
	return ir.dir
}

// Size get range size
func (ir *IteratorRange[KEY, VALUE]) Size() int64 {

	if ir.siter.cur == nil || ir.eiter.cur == nil || ir.siter.idx > ir.eiter.idx {
		return 0
	}

	return ir.eiter.idx - ir.siter.idx + 1

}

// GE2LE [s,e] start with GE, end with LE. (like Seek**)
func (ir *IteratorRange[KEY, VALUE]) GE2LE(start, end KEY) {

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
func (ir *IteratorRange[KEY, VALUE]) GT2LE(start, end KEY) {

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
func (ir *IteratorRange[KEY, VALUE]) GE2LT(start, end KEY) {

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
func (ir *IteratorRange[KEY, VALUE]) GT2LT(start, end KEY) {

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
