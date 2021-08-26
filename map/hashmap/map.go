package hashmap

type Node struct {
	Key        int
	Prev, Next *Node
}

type HashMap struct {
	data []interface{}
	size int
}

func New() {

}
