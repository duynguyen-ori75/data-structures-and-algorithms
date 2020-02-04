package btree

type LeafNode struct {
	keys         []int
	values       []int
	rightSibling *LeafNode
	parent       *InternalNode
}

type InternalNode struct {
	parent *InternalNode
	keys   []int
	// children can be a slice of pointers to LeafNode or InternalNode
	children []interface{}
}

type BPlusTree struct {
	root   interface{}
	degree int
}
