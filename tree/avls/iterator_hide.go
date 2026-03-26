package avls

import "fmt"

func newIterator[KEY, VALUE any](tree *Tree[KEY, VALUE]) *Iterator[KEY, VALUE] {
	hight := tree.Height()
	iter := &Iterator[KEY, VALUE]{
		tree: tree,
		idx:  -1,
	}
	iter.stack = make([]NodeDir[KEY, VALUE], hight)
	return iter
}

func (iter *Iterator[KEY, VALUE]) down(cmp int8) bool {

	if iter.cur == nil {
		if iter.idx > -1 {
			ndir := &iter.stack[iter.idx]
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
	ndir := &iter.stack[iter.idx]
	ndir.N = iter.cur
	ndir.D = cmp
	iter.cur = iter.cur.Children[cmp]
	return true
}

func (iter *Iterator[KEY, VALUE]) up(cmp int8) bool {

	idx := iter.idx
	for {

		if idx > -1 {
			ndir := &iter.stack[idx]
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
				ndir := &iter.stack[iter.idx]
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
	candidateStack := make([]NodeDir[KEY, VALUE], len(iter.stack))

	for cur != nil {
		cmp := iter.tree.Compare(cur.Key, key)

		satisfies := false
		moveDir := int8(0)

		switch lessAndGreater {
		case 1:
			switch {
			case cmp == 0:
				satisfies = true
				moveDir = 0
			case cmp < 0:
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
			case cmp == 1:
				satisfies = true
				moveDir = 1
			case cmp < 0:
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
			copy(candidateStack, iter.stack)
		}

		next := cur.Children[moveDir]
		if next == nil {
			break
		}

		pathIdx++
		iter.stack[pathIdx] = NodeDir[KEY, VALUE]{N: cur, D: moveDir}
		cur = next
	}

	iter.cur = candidate
	iter.idx = candidateIdx
	if candidate != nil {
		copy(iter.stack, candidateStack)
	}

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
	result = fmt.Sprintf("%v  current: %v", iter.stack, iter.Key())
	return result
}
