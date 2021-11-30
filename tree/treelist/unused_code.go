package treelist

// func (tree *Tree) rankNode(key []byte) (*Node, int64) {
// 	const L = 0
// 	const R = 1

// 	cur := tree.getRoot()
// 	var offset int64 = getSize(cur.Children[L])
// 	for cur != nil {
// 		c := tree.compare(key, cur.Key)
// 		switch {
// 		case c < 0:
// 			cur = cur.Children[L]
// 			offset -= getSize(cur.Children[R]) + 1
// 		case c > 0:
// 			cur = cur.Children[R]
// 			offset += getSize(cur.Children[L]) + 1
// 		default:
// 			return cur, offset
// 		}
// 	}
// 	return nil, -1
// }

// func (tree *Tree) getNodeWithIndex(key []byte) (node *Node, idx int64) {
// 	const L = 0
// 	const R = 1

// 	cur := tree.getRoot()
// 	var offset int64 = getSize(cur.Children[L])
// 	for cur != nil {
// 		c := tree.compare(key, cur.Key)
// 		switch {
// 		case c < 0:
// 			cur = cur.Children[L]
// 			if cur != nil {
// 				offset -= getSize(cur.Children[L]) + 1
// 			} else {
// 				return nil, -1
// 			}
// 		case c > 0:
// 			cur = cur.Children[R]
// 			if cur != nil {
// 				offset += getSize(cur.Children[L]) + 1
// 			} else {
// 				return nil, -1
// 			}
// 		default:
// 			return cur, offset
// 		}
// 	}
// 	return nil, -1
// }
