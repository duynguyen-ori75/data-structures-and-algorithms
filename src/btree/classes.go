package btree

type LeafNode struct {
	keys         []int
	values       []int
	leftSibling  *LeafNode
	rightSibling *LeafNode
	parent       *InternalNode
}

type InternalNode struct {
	keys     []int
	children []interface{}
	parent   *InternalNode
}

type BPlusTree struct {
	root   interface{}
	degree int
}
