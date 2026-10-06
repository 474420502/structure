package treeset

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

// seekE seek to the key that (less or greater) than or equal to
func (iter *Iterator[KEY, VALUE]) seekEqual(key KEY, LessAndGreater int8) bool {

	iter.cur = iter.tree.getRoot()
	if iter.cur == nil {
		return false
	}
	iter.idx = -1

	for {
		cmp := iter.tree.Compare(iter.cur.Key, key)
		if cmp == 0 {
			return true
		}

		dir := int8(0)
		if cmp < 0 {
			dir = 1
		}

		if !iter.down(dir) {
			if dir == LessAndGreater {
				iter.up(LessAndGreater)
			}
			return false
		}
	}
}

// SeekLT seek to the key that (less or greater) than
func (iter *Iterator[KEY, VALUE]) seekThan(key KEY, LessAndGreater int8) bool {

	iter.cur = iter.tree.getRoot()
	if iter.cur == nil {
		return false
	}
	iter.idx = -1

	for {

		cmp := iter.tree.Compare(iter.cur.Key, key)

		if cmp == 0 {
			iter.move(LessAndGreater)
			return true
		}

		dir := int8(0)
		if cmp < 0 {
			dir = 1
		}

		if !iter.down(dir) {
			if dir == LessAndGreater {
				iter.up(LessAndGreater)
			}
			return false
		}
	}
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
