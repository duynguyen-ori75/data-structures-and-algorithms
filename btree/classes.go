package btree

/**
 * Purposes of siblings:
 * - Borrow key & merging operator
 * - Clear reference to other LeafNodes -> garbage collection
 */
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
