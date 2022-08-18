package searchtree

import (
	"github.com/474420502/structure/tree/treelist"
)

type Index[KEY, VALUE any] struct {
	tree *treelist.Tree[KEY, VALUE]
}
