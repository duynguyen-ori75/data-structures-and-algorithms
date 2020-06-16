package btree

type LeafNode struct {
	keys         []int
	values       []int
	rightSibling *LeafNode
	parent       *InternalNode
}

type InternalNode struct {
	keys         []int
	children     []interface{}
	rightSibling *InternalNode
	parent       *InternalNode
}

type BPlusTree struct {
	root   interface{}
	degree int
}
