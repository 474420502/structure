package avl

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

// seekE seek to the key that (less or greater) than or equal to
func (iter *Iterator[KEY, VALUE]) seekEqual(key KEY, LessAndGreater int8) bool {

	iter.cur = iter.tree.getRoot()
	if iter.cur == nil {
		return false
	}

	iter.idx = -1

	for {
		cmp := iter.tree.Compare(iter.cur.Key, key)
		if cmp < 0 {
			return true
		}

		if !iter.down(int8(cmp)) {
			if int8(cmp) == LessAndGreater {
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

		if cmp < 0 {
			iter.move(LessAndGreater)
			return true
		}

		if !iter.down(int8(cmp)) {
			if int8(cmp) == LessAndGreater {
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
	result = fmt.Sprintf("%v  current: %v", iter.stack, iter.Key())
	return result
}
