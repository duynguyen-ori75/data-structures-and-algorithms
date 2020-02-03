package btree

import (
	"fmt"
	"testing"
)

func newLeafNode(keys []int, values []int) *LeafNode {
	result := &LeafNode{keys: keys, values: values}
	return result
}

func newTestTree() *BPlusTree {
	root := &InternalNode{}
	root.keys = append(root.keys, 2, 5, 9)
	root.children = append(
		root.children,
		newLeafNode([]int{1, 2}, []int{6, 9}),
		newLeafNode([]int{3, 4, 5}, []int{10, 4, 7}),
		newLeafNode([]int{6, 9}, []int{12, 2}),
		newLeafNode([]int{14, 15}, []int{2, 8}),
	)
	return &BPlusTree{degree: 2, root: root}
}

func TestLeafNodeGetValue(t *testing.T) {
	node := newLeafNode([]int{2, 5, 7}, []int{1, 2})
	val, err := node.getValue(5)
	if err == nil {
		t.Error("Should raise invalid state exception")
	}
	node.values = append(node.values, 7)
	val, err = node.getValue(4)
	if err == nil || err.Error() != "Key 4 not found" {
		t.Error(fmt.Sprintf("Should raise key not found instead of: %s", err.Error()))
	}
	val, err = node.getValue(5)
	if err != nil || val != 2 {
		t.Error("Should get value correctly without any problem")
	}
}

func TestSearch(t *testing.T) {
	tree := newTestTree()
	node, err := tree.search(7)
	if err == nil || err.Error() != "Key 7 not found" {
		t.Error(fmt.Sprintf("Should raise key 7 not found instead of: %s", err.Error()))
	}
	node, err = tree.search(1)
	if err != nil || node.keys[0] != 1 {
		t.Error("Search does not return correct LeafNode")
	}
}
