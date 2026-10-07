package utils

import (
	"testing"

	"github.com/474420502/structure/compare"
	"github.com/474420502/structure/map/orderedmap"
	treequeue "github.com/474420502/structure/queue/priority"
	"github.com/474420502/structure/set/treeset"
	"github.com/474420502/structure/tree/avl"
	"github.com/474420502/structure/tree/avls"
	"github.com/474420502/structure/tree/btree"
	"github.com/474420502/structure/tree/indextree"
	"github.com/474420502/structure/tree/rbtree"
	"github.com/474420502/structure/tree/skiplist"
	"github.com/474420502/structure/tree/treelist"
)

// TestNewWithInference is living documentation for the preferred generic
// construction mechanism: the key type is named once inside the comparator and
// the value type is inferred from a prototype value. Every ordered container
// therefore builds an [int -> string] instance without repeating int in the
// type arguments.
func TestNewWithInference(t *testing.T) {
	check := func(name string, value string, ok bool) {
		t.Helper()
		if !ok || value != "ok" {
			t.Fatalf("%s: got %q, %v", name, value, ok)
		}
	}

	rbt := rbtree.NewWith(compare.Any[int], "")
	rbt.Put(1, "ok")
	v, ok := rbt.Get(1)
	check("rbtree", v, ok)

	avlt := avl.NewWith(compare.Any[int], "")
	avlt.Put(2, "ok")
	v, ok = avlt.Get(2)
	check("avl", v, ok)

	avlst := avls.NewWith(compare.Any[int], "")
	avlst.Put(3, "ok")
	v, ok = avlst.Get(3)
	check("avls", v, ok)

	bt := btree.NewWith(compare.Any[int], "")
	bt.Put(4, "ok")
	v, ok = bt.Get(4)
	check("btree", v, ok)

	sl := skiplist.NewWith(compare.Any[int], "")
	sl.Put(5, "ok")
	v, ok = sl.Get(5)
	check("skiplist", v, ok)

	tl := treelist.NewWith(compare.Any[int], "")
	tl.Put(6, "ok")
	v, ok = tl.Get(6)
	check("treelist", v, ok)

	it := indextree.NewWith(compare.Any[int], "")
	it.Put(7, "ok")
	v, ok = it.Get(7)
	check("indextree", v, ok)

	ts := treeset.NewWith(compare.Any[int], "")
	ts.Add(8, "ok")
	v, ok = ts.Get(8)
	check("treeset", v, ok)

	pq := treequeue.NewWith(compare.Any[int], "")
	pq.Put(9, "ok")
	v, ok = pq.Get(9)
	check("priority", v, ok)

	om := orderedmap.NewWith(compare.Any[int], "")
	om.Put(10, "ok")
	v, ok = om.Get(10)
	check("orderedmap", v, ok)
}
