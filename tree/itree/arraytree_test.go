package itree

// type NodeArray struct {
// 	Key  int
// 	Next *NodeArray
// }

// type TreeArray struct {
// 	data   []*NodeArray
// 	size   int
// 	lrsize [2]int
// }

// func NewTreeArray() *TreeArray {
// 	tree := &TreeArray{
// 		data: make([]*NodeArray, 100000),
// 		size: 0,
// 	}
// 	return tree
// }

// func (ta *TreeArray) insert(cur *NodeArray, key *NodeArray) {
// 	if cur.Next == nil {
// 		cur.Next = key
// 		return
// 	}
// 	ta.insert(cur.Next, key)
// }

// func (ta *TreeArray) Put(key int) {
// 	ta.put(key, 0, 0, -1)
// }

// func (ta *TreeArray) put(key int, parent, child int, cmp int) {
// 	if cap(ta.data) <= child {
// 		ta.size += 1
// 		ta.lrsize[cmp]++

// 		cur := ta.data[parent]
// 		node := &NodeArray{Key: key}
// 		if cur.Next == nil {
// 			cur.Next = node
// 		} else {
// 			ta.insert(ta.data[parent], node)
// 		}
// 		return
// 	}

// 	cur := ta.data[child]
// 	if cur == nil {
// 		ta.data[child] = &NodeArray{Key: key}
// 		ta.size += 1
// 		return
// 	}

// 	if cur.Key > key {
// 		ta.put(key, child, child*2+1, 0)
// 	} else if cur.Key < key {
// 		ta.put(key, child, child*2+2, 1)
// 	} else {
// 		return
// 	}

// }

// func BenchmarkArrayPut(b *testing.B) {
// 	rand := random.New(1683721792150515321)
// 	tree := NewTreeArray()
// 	b.StopTimer()
// 	for i := 0; i < 10000; i++ {
// 		v := rand.Int()
// 		tree.Put(v)
// 		// tree.check()
// 	}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		v := rand.Int()
// 		tree.Put(v)
// 	}
// 	b.Log(tree.lrsize, tree.size)
// }
