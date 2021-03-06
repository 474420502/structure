package treeset

const nL = 0
const nR = 1

func newIterator[T any](tree *Tree[T]) *Iterator[T] {
	hight := tree.Height()
	iter := &Iterator[T]{
		tree:  tree,
		idx:   -1,
		stack: make([]*aNode[T], hight),
	}
	return iter
}

type Iterator[T any] struct {
	tree *Tree[T]

	cur   *aNode[T]
	stack []*aNode[T]

	idx int8
}

func (iter *Iterator[T]) Value() interface{} {
	return iter.cur.Key
}

func (iter *Iterator[T]) SeekToFirst() {
	iter.cur = iter.tree.Root
	iter.idx = -1
	if iter.cur != nil {
		for iter.cur.Children[nL] != nil {
			iter.push()
			iter.cur = iter.cur.Children[nL]
		}
	}
}

func (iter *Iterator[T]) SeekToLast() {
	iter.cur = iter.tree.Root
	iter.idx = -1

	if iter.cur != nil {
		for iter.cur.Children[nR] != nil {
			iter.push()
			iter.cur = iter.cur.Children[nR]
		}
	}
}

func (iter *Iterator[T]) SeekLE(key T) bool {

	iter.idx = -1

	for iter.cur = iter.tree.Root; iter.cur != nil; {
		switch c := iter.tree.Compare(key, iter.cur.Key); c {
		case -1:
			if iter.cur.Children[nL] == nil {
				if !iter.lpop() {
					iter.push()
					iter.cur = nil
				}
				return false
			}

			iter.push()
			iter.cur = iter.cur.Children[nL]
		case 1:
			if iter.cur.Children[nR] == nil {
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[nR]
		case 0:
			return true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}

	return false
}

func (iter *Iterator[T]) SeekLT(key T) bool {

	iter.idx = -1

	for iter.cur = iter.tree.Root; iter.cur != nil; {
		switch c := iter.tree.Compare(key, iter.cur.Key); c {
		case -1:
			if iter.cur.Children[nL] == nil {
				if !iter.lpop() {
					iter.push()
					iter.cur = nil
				}
				return false
			}

			iter.push()
			iter.cur = iter.cur.Children[nL]
		case 1:
			if iter.cur.Children[nR] == nil {
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[nR]
		case 0:
			iter.Prev()
			return true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}

	return false
}

func (iter *Iterator[T]) SeekGE(key T) bool {
	iter.idx = -1

	for iter.cur = iter.tree.Root; iter.cur != nil; {
		switch c := iter.tree.Compare(key, iter.cur.Key); c {
		case -1:
			if iter.cur.Children[nL] == nil {
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[nL]
		case 1:
			if iter.cur.Children[nR] == nil {
				if !iter.rpop() {
					iter.push()
					iter.cur = nil
				}
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[nR]
		case 0:
			return true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}

	return false
}

func (iter *Iterator[T]) SeekGT(key T) bool {
	iter.idx = -1

	for iter.cur = iter.tree.Root; iter.cur != nil; {
		switch c := iter.tree.Compare(key, iter.cur.Key); c {
		case -1:
			if iter.cur.Children[nL] == nil {
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[nL]
		case 1:
			if iter.cur.Children[nR] == nil {
				if !iter.rpop() {
					iter.push()
					iter.cur = nil
				}
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[nR]
		case 0:
			iter.Next()
			return true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}

	return false
}

func (iter *Iterator[T]) Vaild() bool {
	return iter.cur != nil
}

func (iter *Iterator[T]) Next() {

	if iter.cur == nil || iter.cur.Children[nR] == nil {
		// rpop
		if !iter.rpop() {
			iter.push()
			iter.cur = nil
		}
		return
	}

	iter.push()
	iter.cur = iter.cur.Children[nR]

	if iter.cur != nil {
		for iter.cur.Children[nL] != nil {
			iter.push()
			iter.cur = iter.cur.Children[nL]
		}
		return
	}
}

func (iter *Iterator[T]) Prev() {

	if iter.cur == nil || iter.cur.Children[nL] == nil {
		// lpop
		if !iter.lpop() {
			iter.push()
			iter.cur = nil
		}
		return
	}

	iter.push()
	iter.cur = iter.cur.Children[nL]

	if iter.cur != nil {
		for iter.cur.Children[nR] != nil {
			iter.push()
			iter.cur = iter.cur.Children[nR]
		}
		return
	}
}

func (iter *Iterator[T]) push() {
	iter.idx++
	iter.stack[iter.idx] = iter.cur
}

func (iter *Iterator[T]) lpop() bool {

	idx := iter.idx
	cur := iter.cur
	var p *aNode[T]

	for idx >= 0 {
		p = iter.stack[idx]
		idx--
		if p.Children[nR] == cur {
			iter.cur = p
			iter.idx = idx
			return true
		}
		cur = p
	}
	return false
}

func (iter *Iterator[T]) rpop() bool {

	idx := iter.idx
	cur := iter.cur
	var p *aNode[T]

	for idx >= 0 {
		p = iter.stack[idx]
		idx--
		if p.Children[nL] == cur {
			iter.cur = p
			iter.idx = idx
			return true
		}
		cur = p
	}

	return false
}
