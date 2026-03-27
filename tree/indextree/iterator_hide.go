package indextree

func newIterator[T any](tree *Tree[T]) *Iterator[T] {
	hight := tree.hight()
	if hight == 0 {
		hight = 1
	}
	iter := &Iterator[T]{
		tree: tree,
		idx:  -1,
	}
	iter.stack = make([]nodeDir[T], hight+1)
	return iter
}

func (iter *Iterator[T]) down(cmp int8) bool {

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

func (iter *Iterator[T]) up(cmp int8) bool {

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
				if int(iter.idx)+1 < len(iter.stack) {
					iter.idx++
					ndir := &iter.stack[iter.idx]
					ndir.N = iter.cur
					ndir.D = cmp
				}
				iter.cur = nil
			}

			return false
		}
		idx--
	}

}

func (iter *Iterator[T]) seekEqual(key T, LessAndGreater int8) bool {

	iter.cur = iter.tree.getRoot()
	if iter.cur == nil {
		return false
	}

	iter.idx = -1

	for {
		cmp := iter.tree.compare(iter.cur.Key, key)

		var dir int8
		if cmp < 0 {
			dir = 1
		} else {
			dir = 0
		}

		if cmp == 0 {
			iter.pos = iter.tree.IndexOf(key)
			return true
		}

		if !iter.down(dir) {
			if dir == LessAndGreater {
				iter.up(LessAndGreater)
			}
			iter.pos = iter.tree.IndexOf(key)
			return false
		}
	}
}

func (iter *Iterator[T]) seekThan(key T, LessAndGreater int8) bool {

	iter.cur = iter.tree.getRoot()
	if iter.cur == nil {
		return false
	}

	iter.idx = -1

	for {

		cmp := iter.tree.compare(iter.cur.Key, key)

		var dir int8
		if cmp < 0 {
			dir = 1
		} else {
			dir = 0
		}

		if cmp > 0 {
			iter.move(LessAndGreater)
			iter.pos = iter.tree.IndexOf(iter.cur.Key)
			return true
		}

		if !iter.down(dir) {
			if dir == LessAndGreater {
				iter.up(LessAndGreater)
			}
			if iter.cur != nil {
				iter.pos = iter.tree.IndexOf(iter.cur.Key)
			}
			return false
		}
	}
}

func (iter *Iterator[T]) move(cmp int8) {
	rcmp := ^cmp + 2
	if iter.down(cmp) {
		for iter.down(rcmp) {

		}
	} else {
		iter.up(cmp)
	}

}
