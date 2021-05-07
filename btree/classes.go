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
	tree    *BPlusTree
	depth   int
	indices []int
	nodes   []interface{}
}

func newLeafNode(keys []int, values []int) *LeafNode {
	return &LeafNode{keys: keys, values: values}
}

func newInternalNode(keys []int, children []interface{}) *InternalNode {
	return &InternalNode{keys: keys, children: children}
}

func newBPlusTree(degree int) *BPlusTree {
	return &BPlusTree{root: newLeafNode([]int{}, []int{}), degree: degree}
}

func newCursor(tree *BPlusTree) *Cursor {
	return &Cursor{tree: tree, depth: -1, indices: []int{}, nodes: []interface{}{}}
}
