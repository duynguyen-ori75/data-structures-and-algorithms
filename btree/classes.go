package btree

type LeafNode struct {
	keys   []int
	values []int
}

type InternalNode struct {
	keys     []int
	children []interface{}
}

type BPlusTree struct {
	root   interface{}
	degree int
}

type Cursor struct {
	depth   int
	indices []int
	nodes   []interface{}
}
