package arraytree

import (
	"fmt"
	"log"

	"github.com/474420502/structure/compare"
)

type Tree struct {
	compare compare.Compare
	size    int64
	datas   []Node
}

type Slice struct {
	Key   interface{}
	Value interface{}
}

type Node struct {
	Size int64
	Slice
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Key)
}

func New() *Tree {
	tree := &Tree{compare: compare.Int}
	tree.datas = make([]Node, 1<<3-1)
	return tree
}

func (tree *Tree) Put(key, value interface{}) {
	dlen := len(tree.datas)
	var idx int = 0
	var cur *Node
	for {
		log.Println(idx)
		if idx < dlen {
			cur = &tree.datas[idx]
			if cur.Size == 0 {
				// tree.datas[idx] = &Node{Size: 1, Slice: Slice{Key: key, Value: value}}
				cur.Size = 1
				cur.Slice.Key = key
				cur.Slice.Value = value
				tree.size++
				tree.fixPutSize(idx)
				return
			}
		} else {
			idx = (idx - 1) >> 1
			mov := idx
			log.Println(mov)
			// var height int64 = 2
			var limitsize int64 = 1 << 2
			var movepath []*Node = []*Node{&tree.datas[idx]}
			for {
				idx = getParentIndex(idx)
				cur := &tree.datas[idx]
				movepath = append(movepath, cur)
				if cur.Size < limitsize-1 {
					// TODO: mov
					for cur.Size != 0 {
						lidx, ridx := getChildrenIndex(idx)
						lc := &tree.datas[lidx]
						rc := &tree.datas[ridx]
						if lc.Size < rc.Size {
							idx = lidx
							cur = lc
						} else {
							idx = ridx
							cur = rc
						}
						movepath = append(movepath, cur)
					}

					moved := movepath[len(movepath)-1]
					for i := len(movepath) - 2; i >= 0; i-- {
						cur := movepath[i]
						moved.Slice = cur.Slice
						moved = cur
					}

					moved.Slice.Key = key
					moved.Slice.Value = value
					moved.Size = 1
					tree.fixPutSize(idx)
					log.Println(movepath, moved)
					return

				}

				limitsize = limitsize << 1
			}

		}

		c := tree.compare(key, cur.Key)
		switch {
		case c < 0:
			idx = idx<<1 + 1
		case c > 0:
			idx = idx<<1 + 2
		default:
			tree.datas[idx].Value = value
			return
		}
	}
}

func getParentIndex(idx int) int {
	return (idx - 1) >> 1
}

func getChildrenIndex(idx int) (lidx, ridx int) {
	idx = idx << 1
	return idx + 1, idx + 2
}

func (tree *Tree) fixPutSize(idx int) {
	for {
		idx = (idx - 1) >> 1
		if idx >= 0 {
			tree.datas[idx].Size++
		} else {
			return
		}
	}
}
