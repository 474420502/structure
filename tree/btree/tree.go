// Package btree implements a generic in-memory B-tree.
//
// It uses the repository-wide standard comparator contract: a negative result
// means k1 < k2, a positive result means k1 > k2, and zero means equality.
// Nodes hold up to 2*degree-1 items, which keeps the tree shallow and makes
// both point lookups and in-order iteration cache friendly.
package btree

import "github.com/474420502/structure/compare"

// defaultDegree is the minimum degree used by New.
const defaultDegree = 16

// Item is a key/value pair stored in a node.
type Item[K any, V any] struct {
	Key   K
	Value V
}

type node[K any, V any] struct {
	items    []Item[K, V]
	children []*node[K, V]
}

// Tree is a generic in-memory B-tree ordered map.
type Tree[K any, V any] struct {
	root    *node[K, V]
	degree  int
	compare compare.Compare[K]
	size    int
	zero    V
}

// New creates a B-tree with the default minimum degree.
func New[K any, V any](comp compare.Compare[K]) *Tree[K, V] {
	return NewWithDegree[K, V](comp, defaultDegree)
}

// NewWithDegree creates a B-tree with the given minimum degree.
func NewWithDegree[K any, V any](comp compare.Compare[K], degree int) *Tree[K, V] {
	if degree < 2 {
		degree = 2
	}
	tree := &Tree[K, V]{
		degree:  degree,
		compare: comp,
	}
	tree.root = tree.newNode()
	return tree
}

func (t *Tree[K, V]) maxItems() int { return 2*t.degree - 1 }

// newNode returns a node with room for a full node's worth of items, which
// avoids repeated slice growth while a node fills up.
func (t *Tree[K, V]) newNode() *node[K, V] {
	return &node[K, V]{items: make([]Item[K, V], 0, t.maxItems())}
}

// Len returns the number of stored keys.
func (t *Tree[K, V]) Len() int { return t.size }

// Size returns the number of stored keys.
func (t *Tree[K, V]) Size() int { return t.size }

// Empty reports whether the tree has no entries.
func (t *Tree[K, V]) Empty() bool { return t.size == 0 }

// Clear removes every entry.
func (t *Tree[K, V]) Clear() {
	t.root = t.newNode()
	t.size = 0
}

// Height returns the number of node levels; a non-empty tree is at least 1.
func (t *Tree[K, V]) Height() int {
	if t.size == 0 {
		return 0
	}
	height := 1
	n := t.root
	for len(n.children) > 0 {
		height++
		n = n.children[0]
	}
	return height
}

// Get returns the value stored for key.
func (t *Tree[K, V]) Get(key K) (V, bool) {
	n := t.root
	for {
		// search is inlined here: it is small but was not inlinable as a
		// method, and Get is the hottest read path.
		lo, hi := 0, len(n.items)
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			if t.compare(n.items[mid].Key, key) < 0 {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if lo < len(n.items) && t.compare(n.items[lo].Key, key) == 0 {
			return n.items[lo].Value, true
		}
		if len(n.children) == 0 {
			return t.zero, false
		}
		n = n.children[lo]
	}
}

// Put inserts key only when it is absent; an existing value is preserved.
func (t *Tree[K, V]) Put(key K, value V) bool { return t.insert(key, value, false) }

// InsertIfAbsent is the preferred explicit name for insert-only writes.
func (t *Tree[K, V]) InsertIfAbsent(key K, value V) bool { return t.insert(key, value, false) }

// Set inserts key or overwrites the existing value.
func (t *Tree[K, V]) Set(key K, value V) bool { return t.insert(key, value, true) }

// Upsert ensures the key exists with value and reports whether an existing
// value was replaced.
func (t *Tree[K, V]) Upsert(key K, value V) bool {
	if n, i, found := t.find(key); found {
		n.items[i].Value = value
		return true
	}
	t.insert(key, value, false)
	return false
}

// Delete removes key and returns the previous value when present.
func (t *Tree[K, V]) Delete(key K) (V, bool) {
	value, removed := t.remove(t.root, key)
	if removed {
		t.size--
		if len(t.root.items) == 0 && len(t.root.children) > 0 {
			t.root = t.root.children[0]
		}
	}
	return value, removed
}

// Remove is the historical alias for Delete.
func (t *Tree[K, V]) Remove(key K) (V, bool) { return t.Delete(key) }

// Traverse visits every entry in ascending key order.
func (t *Tree[K, V]) Traverse(do func(K, V) bool) {
	var walk func(n *node[K, V]) bool
	walk = func(n *node[K, V]) bool {
		for i := range n.items {
			if len(n.children) > 0 && !walk(n.children[i]) {
				return false
			}
			if !do(n.items[i].Key, n.items[i].Value) {
				return false
			}
		}
		if len(n.children) > 0 {
			return walk(n.children[len(n.items)])
		}
		return true
	}
	walk(t.root)
}

// Values returns every value in ascending key order.
func (t *Tree[K, V]) Values() []V {
	if t.size == 0 {
		return nil
	}
	result := make([]V, 0, t.size)
	t.Traverse(func(_ K, value V) bool {
		result = append(result, value)
		return true
	})
	return result
}

// Min returns the smallest entry.
func (t *Tree[K, V]) Min() (K, V, bool) {
	if t.size == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	n := t.root
	for len(n.children) > 0 {
		n = n.children[0]
	}
	return n.items[0].Key, n.items[0].Value, true
}

// Max returns the largest entry.
func (t *Tree[K, V]) Max() (K, V, bool) {
	if t.size == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	n := t.root
	for len(n.children) > 0 {
		n = n.children[len(n.children)-1]
	}
	return n.items[len(n.items)-1].Key, n.items[len(n.items)-1].Value, true
}

// lowerBound returns the first index whose key is >= key. Callers compare the
// item at that index for equality; the binary search keeps wide nodes cheap.
func (t *Tree[K, V]) lowerBound(n *node[K, V], key K) int {
	lo, hi := 0, len(n.items)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if t.compare(n.items[mid].Key, key) < 0 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// find returns the node and index holding key.
func (t *Tree[K, V]) find(key K) (*node[K, V], int, bool) {
	n := t.root
	for {
		lo, hi := 0, len(n.items)
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			if t.compare(n.items[mid].Key, key) < 0 {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if lo < len(n.items) && t.compare(n.items[lo].Key, key) == 0 {
			return n, lo, true
		}
		if len(n.children) == 0 {
			return nil, 0, false
		}
		n = n.children[lo]
	}
}

func (t *Tree[K, V]) insert(key K, value V, overwrite bool) bool {
	r := t.root
	var inserted bool
	if len(r.items) == t.maxItems() {
		s := t.newNode()
		s.children = []*node[K, V]{r}
		t.root = s
		t.splitChild(s, 0)
		inserted = t.insertNonFull(s, key, value, overwrite)
	} else {
		inserted = t.insertNonFull(r, key, value, overwrite)
	}
	if inserted {
		t.size++
	}
	return inserted
}

func (t *Tree[K, V]) insertNonFull(n *node[K, V], key K, value V, overwrite bool) bool {
	i := t.lowerBound(n, key)
	found := i < len(n.items) && t.compare(n.items[i].Key, key) == 0
	if found {
		if overwrite {
			n.items[i].Value = value
		}
		return false
	}

	if len(n.children) == 0 {
		n.items = append(n.items, Item[K, V]{})
		copy(n.items[i+1:], n.items[i:])
		n.items[i] = Item[K, V]{Key: key, Value: value}
		return true
	}

	if len(n.children[i].items) == t.maxItems() {
		t.splitChild(n, i)
		switch t.compare(key, n.items[i].Key) {
		case 0:
			if overwrite {
				n.items[i].Value = value
			}
			return false
		case 1:
			i++
		}
	}
	return t.insertNonFull(n.children[i], key, value, overwrite)
}

func (t *Tree[K, V]) splitChild(parent *node[K, V], i int) {
	child := parent.children[i]
	mid := len(child.items) / 2
	midItem := child.items[mid]

	right := t.newNode()
	right.items = append(right.items, child.items[mid+1:]...)
	if len(child.children) > 0 {
		right.children = append([]*node[K, V](nil), child.children[mid+1:]...)
		child.children = child.children[:mid+1]
	}
	child.items = child.items[:mid]

	parent.items = append(parent.items, Item[K, V]{})
	copy(parent.items[i+1:], parent.items[i:])
	parent.items[i] = midItem

	parent.children = append(parent.children, nil)
	copy(parent.children[i+2:], parent.children[i+1:])
	parent.children[i+1] = right
}

func (t *Tree[K, V]) remove(n *node[K, V], key K) (V, bool) {
	i := t.lowerBound(n, key)
	found := i < len(n.items) && t.compare(n.items[i].Key, key) == 0

	if found {
		if len(n.children) == 0 {
			value := n.items[i].Value
			n.items = append(n.items[:i], n.items[i+1:]...)
			return value, true
		}
		return t.removeInternal(n, i, key)
	}

	if len(n.children) == 0 {
		return t.zero, false
	}

	if len(n.children[i].items) < t.degree {
		t.fill(n, i)
	}
	if i > len(n.items) {
		return t.remove(n.children[i-1], key)
	}
	return t.remove(n.children[i], key)
}

func (t *Tree[K, V]) removeInternal(n *node[K, V], i int, key K) (V, bool) {
	removed := n.items[i].Value
	left := n.children[i]
	right := n.children[i+1]

	if len(left.items) >= t.degree {
		pred := maxItem(left)
		n.items[i] = pred
		t.remove(left, pred.Key)
		return removed, true
	}
	if len(right.items) >= t.degree {
		succ := minItem(right)
		n.items[i] = succ
		t.remove(right, succ.Key)
		return removed, true
	}

	// Merge left, the separator and right into left. The separator carried the
	// key we are deleting, so the key now lives inside the merged child.
	t.merge(n, i)
	return t.remove(left, key)
}

func minItem[K any, V any](n *node[K, V]) Item[K, V] {
	for len(n.children) > 0 {
		n = n.children[0]
	}
	return n.items[0]
}

func maxItem[K any, V any](n *node[K, V]) Item[K, V] {
	for len(n.children) > 0 {
		n = n.children[len(n.children)-1]
	}
	return n.items[len(n.items)-1]
}

func (t *Tree[K, V]) fill(n *node[K, V], i int) {
	if i > 0 && len(n.children[i-1].items) >= t.degree {
		t.borrowFromPrev(n, i)
		return
	}
	if i < len(n.items) && len(n.children[i+1].items) >= t.degree {
		t.borrowFromNext(n, i)
		return
	}
	if i < len(n.items) {
		t.merge(n, i)
		return
	}
	t.merge(n, i-1)
}

func (t *Tree[K, V]) borrowFromPrev(n *node[K, V], i int) {
	child := n.children[i]
	sibling := n.children[i-1]

	child.items = append(child.items, Item[K, V]{})
	copy(child.items[1:], child.items[:len(child.items)-1])
	child.items[0] = n.items[i-1]

	if len(child.children) > 0 {
		child.children = append(child.children, nil)
		copy(child.children[1:], child.children[:len(child.children)-1])
		child.children[0] = sibling.children[len(sibling.children)-1]
		sibling.children = sibling.children[:len(sibling.children)-1]
	}

	n.items[i-1] = sibling.items[len(sibling.items)-1]
	sibling.items = sibling.items[:len(sibling.items)-1]
}

func (t *Tree[K, V]) borrowFromNext(n *node[K, V], i int) {
	child := n.children[i]
	sibling := n.children[i+1]

	child.items = append(child.items, n.items[i])
	if len(child.children) > 0 {
		child.children = append(child.children, sibling.children[0])
		sibling.children = sibling.children[1:]
	}

	n.items[i] = sibling.items[0]
	sibling.items = sibling.items[1:]
}

func (t *Tree[K, V]) merge(n *node[K, V], i int) {
	child := n.children[i]
	sibling := n.children[i+1]

	child.items = append(child.items, n.items[i])
	child.items = append(child.items, sibling.items...)
	if len(child.children) > 0 {
		child.children = append(child.children, sibling.children...)
	}

	n.items = append(n.items[:i], n.items[i+1:]...)
	n.children = append(n.children[:i+1], n.children[i+2:]...)
}
