package avl

const L = 0
const R = 1

func newIterator(tree *Tree) *Iterator {
	hight := tree.Height()
	iter := &Iterator{
		tree:  tree,
		idx:   -1,
		stack: make([]*Node, hight),
	}
	return iter
}

// Iterator tree iterator
type Iterator struct {
	tree *Tree

	cur   *Node
	stack []*Node

	idx int8
}

func (iter *Iterator) Key() interface{} {
	return iter.cur.Key
}

func (iter *Iterator) Value() interface{} {
	return iter.cur.Value
}

// SeekToFirst seek to first item
func (iter *Iterator) SeekToFirst() {
	iter.cur = iter.tree.Root
	iter.idx = -1
	if iter.cur != nil {
		for iter.cur.Children[L] != nil {
			iter.push()
			iter.cur = iter.cur.Children[L]
		}
	}
}

// SeekToFirst seek to last item
func (iter *Iterator) SeekToLast() {
	iter.cur = iter.tree.Root
	iter.idx = -1

	if iter.cur != nil {
		for iter.cur.Children[R] != nil {
			iter.push()
			iter.cur = iter.cur.Children[R]
		}
	}
}

func (iter *Iterator) SeekLE(key interface{}) bool {

	iter.idx = -1
	iter.cur = iter.tree.Root
	if iter.cur == nil {
		return false
	}

	for {
		switch c := iter.tree.Compare(key, iter.cur.Key); c {
		case -1:
			if iter.cur.Children[L] == nil {
				if !iter.lpop() {
					iter.push()
					iter.cur = nil
				}
				return false
			}

			iter.push()
			iter.cur = iter.cur.Children[L]
		case 1:
			if iter.cur.Children[R] == nil {
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[R]
		case 0:
			return true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}
}

func (iter *Iterator) SeekLT(key interface{}) bool {

	iter.idx = -1
	iter.cur = iter.tree.Root
	if iter.cur == nil {
		return false
	}

	for {
		switch c := iter.tree.Compare(key, iter.cur.Key); c {
		case -1:
			if iter.cur.Children[L] == nil {
				if !iter.lpop() {
					iter.push()
					iter.cur = nil
				}
				return false
			}

			iter.push()
			iter.cur = iter.cur.Children[L]
		case 1:
			if iter.cur.Children[R] == nil {
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[R]
		case 0:
			iter.Prev()
			return true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}
}

func (iter *Iterator) SeekGE(key interface{}) bool {
	iter.idx = -1
	iter.cur = iter.tree.Root
	if iter.cur == nil {
		return false
	}

	for {
		switch c := iter.tree.Compare(key, iter.cur.Key); c {
		case -1:
			if iter.cur.Children[L] == nil {
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[L]
		case 1:
			if iter.cur.Children[R] == nil {
				if !iter.rpop() {
					iter.push()
					iter.cur = nil
				}
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[R]
		case 0:
			return true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}
}

func (iter *Iterator) SeekGT(key interface{}) bool {
	iter.idx = -1
	iter.cur = iter.tree.Root
	if iter.cur == nil {
		return false
	}

	for {
		switch c := iter.tree.Compare(key, iter.cur.Key); c {
		case -1:
			if iter.cur.Children[L] == nil {
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[L]
		case 1:
			if iter.cur.Children[R] == nil {
				if !iter.rpop() {
					iter.push()
					iter.cur = nil
				}
				return false
			}
			iter.push()
			iter.cur = iter.cur.Children[R]
		case 0:
			iter.Next()
			return true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}
}

func (iter *Iterator) Vaild() bool {
	return iter.cur != nil
}

func (iter *Iterator) Next() {

	if iter.cur == nil || iter.cur.Children[R] == nil {
		// rpop
		if !iter.rpop() {
			iter.push()
			iter.cur = nil
		}
		return
	}

	iter.push()
	iter.cur = iter.cur.Children[R]

	if iter.cur != nil {
		for iter.cur.Children[L] != nil {
			iter.push()
			iter.cur = iter.cur.Children[L]
		}
		return
	}
}

func (iter *Iterator) Prev() {

	if iter.cur == nil || iter.cur.Children[L] == nil {
		// lpop
		if !iter.lpop() {
			iter.push()
			iter.cur = nil
		}
		return
	}

	iter.push()
	iter.cur = iter.cur.Children[L]

	if iter.cur != nil {
		for iter.cur.Children[R] != nil {
			iter.push()
			iter.cur = iter.cur.Children[R]
		}
		return
	}
}

func (iter *Iterator) push() {
	iter.idx++
	iter.stack[iter.idx] = iter.cur
}

func (iter *Iterator) lpop() bool {

	idx := iter.idx
	cur := iter.cur
	var p *Node

	for idx >= 0 {
		p = iter.stack[idx]
		idx--
		if p.Children[R] == cur {
			iter.cur = p
			iter.idx = idx
			return true
		}
		cur = p
	}
	return false
}

func (iter *Iterator) rpop() bool {

	idx := iter.idx
	cur := iter.cur
	var p *Node

	for idx >= 0 {
		p = iter.stack[idx]
		idx--
		if p.Children[L] == cur {
			iter.cur = p
			iter.idx = idx
			return true
		}
		cur = p
	}

	return false
}

// Clone 复制一个当前迭代的iterator. 用于复位
func (iter *Iterator) Clone() *Iterator {
	return &Iterator{tree: iter.tree, cur: iter.cur, idx: iter.idx}
}
