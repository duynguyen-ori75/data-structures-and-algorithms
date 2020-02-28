package btree

import (
	"reflect"
	"testing"
)

func newTestTree() *BPlusTree {
	root := &InternalNode{}
	root.keys = append(root.keys, 2, 5, 9)
	root.children = append(
		root.children,
		newLeafNode([]int{1, 2}, []int{6, 9}, root),
		newLeafNode([]int{3, 4, 5}, []int{10, 4, 7}, root),
		newLeafNode([]int{6, 9}, []int{12, 2}, root),
		newLeafNode([]int{14, 15}, []int{2, 8}, root),
	)
	for idx := 0; idx < 3; idx++ {
		root.children[idx].(*LeafNode).rightSibling = root.children[idx+1].(*LeafNode)
	}
	return &BPlusTree{degree: 2, root: root}
}

func TestHelper(t *testing.T) {
	a := []int{}
	a = insertInt(a, 0, 5)
	a = insertInt(a, 1, 4)
	a = insertInt(a, 0, 2)
	a = insertInt(a, 2, 1)
	if !reflect.DeepEqual(a, []int{2, 5, 1, 4}) {
		t.Errorf("Expected array should be %s instead of %s", arrayToString(a), arrayToString([]int{2, 5, 1, 4}))
	}
	a = removeInt(a, 1)
	if !reflect.DeepEqual(a, []int{2, 1, 4}) {
		t.Errorf("Expected array should be %s instead of %s", arrayToString(a), arrayToString([]int{2, 1, 4}))
	}
	a = removeInt(removeInt(removeInt(a, 2), 0), 0)
	if len(a) != 0 {
		t.Errorf("Slice after removing all elements should be empty. Current size %d", len(a))
	}
}

func TestLeafNodeGetValue(t *testing.T) {
	node := newLeafNode([]int{2, 5, 7}, []int{1, 2}, nil)
	val, err := node.getValue(5)
	if err == nil {
		t.Error("Should raise invalid state exception")
	}
	node.values = append(node.values, 7)
	val, err = node.getValue(4)
	if err == nil || err.Error() != "Key 4 not found" {
		t.Errorf("Should raise key not found instead of: %s", err.Error())
	}
	val, err = node.getValue(5)
	if err != nil || val != 2 {
		t.Error("Should get value correctly without any problem")
	}
}

func TestSearch(t *testing.T) {
	tree := newTestTree()
	node, err := tree.search(7)
	if err != nil || !reflect.DeepEqual(node.keys, []int{6, 9}) {
		t.Errorf("Serch does not return correct LeafNode. Keys of found node: %s", arrayToString(node.keys))
	}
	node, err = tree.search(1)
	if err != nil || node.keys[0] != 1 {
		t.Errorf("Search does not return correct LeafNode. Keys of found node: %s", arrayToString(node.keys))
	}
}

func TestLeafNodeInsert(t *testing.T) {
	testLeafNode := newLeafNode([]int{3}, []int{10}, nil)
	err := testLeafNode.insertValue(1, 5, 1)
	if err != nil {
		t.Error("Test leaf node should insert value without any exception")
	}
	if testLeafNode.parent == nil || testLeafNode.rightSibling == nil {
		t.Error("Parent node and sibling node should not be nil")
	}
	if !reflect.DeepEqual(testLeafNode.parent.keys, []int{3}) {
		t.Errorf("Parent node's keys are not correct. Should be [3] instead of %s", arrayToString(testLeafNode.parent.keys))
	}
	if !reflect.DeepEqual(testLeafNode.keys, []int{1}) {
		t.Errorf("Current node's keys are not correct. Should be [1] instead of %s", arrayToString(testLeafNode.keys))
	}
	if !reflect.DeepEqual(testLeafNode.rightSibling.keys, []int{3}) {
		t.Errorf("Current node's keys are not correct. Should be [1] instead of %s", arrayToString(testLeafNode.keys))
	}
	if !reflect.DeepEqual(testLeafNode.values, []int{5}) {
		t.Errorf("Current node's keys are not correct. Should be [5] instead of %s", arrayToString(testLeafNode.keys))
	}
	if !reflect.DeepEqual(testLeafNode.rightSibling.values, []int{10}) {
		t.Errorf("Current node's keys are not correct. Should be [10] instead of %s", arrayToString(testLeafNode.keys))
	}
}

func TestTreeInsert(t *testing.T) {
	tree := newTestTree()
	err := tree.insert(1, 2)
	if err == nil {
		t.Error("Key already exists exception should be raised")
	}
	err = tree.insert(7, 3)
	if err != nil {
		t.Errorf("Tree should insert correctly. Exception: %s", err.Error())
	}
	thirdChild := tree.root.(*InternalNode).children[2].(*LeafNode)
	if !reflect.DeepEqual(thirdChild.keys, []int{6, 7, 9}) {
		t.Errorf("Third child should have keys [6, 7, 9] instead of %s", arrayToString(thirdChild.keys))
	}
	if !reflect.DeepEqual(thirdChild.values, []int{12, 3, 2}) {
		t.Errorf("Third child should have keys [12, 3, 2] instead of %s", arrayToString(thirdChild.values))
	}
}

func TestTreeInsertFromScratch(t *testing.T) {
	tree := newBPlusTree(1)
	err := tree.insert(1, 5)
	if err != nil {
		t.Error("Insert shouldn't raise any exception")
	}
	_, ok := tree.root.(*LeafNode)
	if !ok {
		t.Error("Type of 1-node tree's root should be a LeafNode")
	}
	err = tree.insert(2, 5)
	if err != nil {
		t.Error("Insert shouldn't raise any exception")
	}
	root, ok := tree.root.(*InternalNode)
	if !ok {
		t.Error("Type of 2-node tree's root should be a InternalNode")
	}
	if len(root.children) != 2 {
		t.Error("2-node tree should have two LeafNodes and 1 InternalNode")
	}
	err = tree.insert(3, 2)
	if err != nil {
		t.Error("Insert shouldn't raise any error")
	}
	if !reflect.DeepEqual(tree.root.(*InternalNode).keys, []int{3}) {
		t.Errorf("Root's keys should be [3] instead of %s", arrayToString(tree.root.(*InternalNode).keys))
	}
	leftChild, okLeft := tree.root.(*InternalNode).children[0].(*InternalNode)
	rightChild, okRight := tree.root.(*InternalNode).children[1].(*InternalNode)
	if !okLeft || !okRight {
		t.Error("Two children should be both InternalNodes")
	}
	if !reflect.DeepEqual(leftChild.keys, []int{2}) || !reflect.DeepEqual(rightChild.keys, []int{3}) {
		t.Errorf("Left child's keys and right child's keys should be [2]/[3] instead of %s/%s",
			arrayToString(leftChild.keys), arrayToString(rightChild.keys))
	}
}

func TestLeafNodeRemoval(t *testing.T) {
	testLeafNode := newLeafNode([]int{2, 5, 4}, []int{10, 2, 1}, nil)
	err := testLeafNode.deleteKey(5, 2)
	if err != nil || !reflect.DeepEqual(testLeafNode.keys, []int{2, 4}) {
		t.Errorf("Delete operation should work correctly. Expected keys are [2, 4] instead of %s", arrayToString(testLeafNode.keys))
	}
	err = testLeafNode.deleteKey(2, 2)
	if err != nil {
		t.Error("Delete ops shouldn't raise any exception here")
	}
	err = testLeafNode.deleteKey(4, 2)
	if err != nil {
		t.Error("Delete ops shouldn't raise any exception here")
	}
	if len(testLeafNode.keys) != 0 {
		t.Error("After 3 ops, leafNode should not have any key")
	}
	err = testLeafNode.deleteKey(1, 1)
	if err == nil || err.Error() != "Key 1 not found" {
		t.Error("Delete key from empty node should raise key not found exception")
	}
}

func TestTreeRemoval(t *testing.T) {
	tree := newTestTree()
	err := tree.delete(9)
	if err != nil {
		t.Errorf("Delete ops should not raise exception. Error: %s", err.Error())
	}
	node, err := tree.search(9)
	if err != nil {
		t.Errorf("Search ops should not raise exception. Error: %s", err.Error())
	}
	if len(node.keys) != 1 || node.keys[0] != 6 {
		t.Errorf("LeafNode's keys are not correct. Should be [6] instead of %s", arrayToString(node.keys))
	}
}
