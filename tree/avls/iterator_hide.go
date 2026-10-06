package avls

import "fmt"

func newIterator[KEY, VALUE any](tree *Tree[KEY, VALUE]) *Iterator[KEY, VALUE] {
	return &Iterator[KEY, VALUE]{
		tree: tree,
		idx:  -1,
	}
}

// dir returns a writable slot for path index i, spilling to the heap only for
// trees deeper than the inline backing array.
func (iter *Iterator[KEY, VALUE]) dir(i int8) *NodeDir[KEY, VALUE] {
	if int(i) < iteratorStackSize {
		return &iter.backing[i]
	}
	j := int(i) - iteratorStackSize
	for len(iter.overflow) <= j {
		iter.overflow = append(iter.overflow, NodeDir[KEY, VALUE]{})
	}
	return &iter.overflow[j]
}

func (iter *Iterator[KEY, VALUE]) down(cmp int8) bool {

	if iter.cur == nil {
		if iter.idx > -1 {
			ndir := iter.dir(iter.idx)
			rcmp := ^cmp + 2
			if ndir.D == rcmp {
				iter.cur = ndir.N
				iter.idx--
				return true
			}
		}
		return false
	}

	if iter.cur.Children[cmp] == nil {
		return false
	}
	iter.idx += 1
	ndir := iter.dir(iter.idx)
	ndir.N = iter.cur
	ndir.D = cmp
	iter.cur = iter.cur.Children[cmp]
	return true
}

func (iter *Iterator[KEY, VALUE]) up(cmp int8) bool {

	idx := iter.idx
	for {

		if idx > -1 {
			ndir := iter.dir(idx)
			rcmp := ^cmp + 2
			if ndir.D == rcmp {
				idx--
				iter.idx = idx
				iter.cur = ndir.N
				return true
			}
		} else {
			if iter.cur != nil {
				iter.idx++
				ndir := iter.dir(iter.idx)
				ndir.N = iter.cur
				ndir.D = cmp
				iter.cur = nil
			}

			return false
		}
		idx--
	}

}

func (iter *Iterator[KEY, VALUE]) seekEqual(key KEY, lessAndGreater int8) bool {
	return iter.seekBound(key, lessAndGreater, false)
}

func (iter *Iterator[KEY, VALUE]) seekThan(key KEY, lessAndGreater int8) bool {
	return iter.seekBound(key, lessAndGreater, true)
}

func (iter *Iterator[KEY, VALUE]) seekBound(key KEY, lessAndGreater int8, strict bool) bool {
	cur := iter.tree.getRoot()
	if cur == nil {
		iter.cur = nil
		iter.idx = -1
		return false
	}

	iter.idx = -1
	pathIdx := int8(-1)
	exact := false

	var candidate *Node[KEY, VALUE]
	candidateIdx := int8(-1)

	for cur != nil {
		cmp := iter.tree.Compare(cur.Key, key)

		satisfies := false
		moveDir := int8(0)

		switch lessAndGreater {
		case 1:
			switch {
			case cmp > 0:
				satisfies = true
				moveDir = 0
			case cmp == 0:
				exact = true
				if !strict {
					satisfies = true
					moveDir = 0
				} else {
					moveDir = 1
				}
			default:
				moveDir = 1
			}
		case 0:
			switch {
			case cmp < 0:
				satisfies = true
				moveDir = 1
			case cmp == 0:
				exact = true
				if !strict {
					satisfies = true
					moveDir = 1
				} else {
					moveDir = 0
				}
			default:
				moveDir = 0
			}
		}

		if satisfies {
			candidate = cur
			candidateIdx = pathIdx
		}

		next := cur.Children[moveDir]
		if next == nil {
			break
		}

		pathIdx++
		*iter.dir(pathIdx) = NodeDir[KEY, VALUE]{N: cur, D: moveDir}
		cur = next
	}

	// seekBound only pushes onto the path, so the candidate's ancestor prefix
	// is still intact; no snapshot or copy is needed.
	iter.cur = candidate
	iter.idx = candidateIdx

	return exact
}

// move move to (left or right) cmp == 0 or 1 , 0 == left , 1 == right
func (iter *Iterator[KEY, VALUE]) move(cmp int8) {
	rcmp := ^cmp + 2
	// iter.push(rcmp)
	if iter.down(cmp) {
		for iter.down(rcmp) {

		}
	} else {
		iter.up(cmp)
	}

}

func (iter *Iterator[KEY, VALUE]) view() (result string) {
	n := int(iter.idx) + 1
	if n < 0 {
		n = 0
	}
	if n > iteratorStackSize {
		n = iteratorStackSize
	}
	result = fmt.Sprintf("%v  current: %v", iter.backing[:n], iter.Key())
	return result
}
