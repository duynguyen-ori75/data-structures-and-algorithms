package btree

import (
	//"log"
	"reflect"
	"testing"
)

/**
 * @brief      initialize InternalNode with the following structure, should works with degree 3
 *
 * 				   5
 * 			   /      \
 * 			 4          10
 * 		   /   \      /    \
 * 		  1   4,5   7,8    10
 *
 * @return     the parent InternalNode and the maximum degree of BTree
 */
func newTestInternalNode() (*InternalNode, int) {
	// initialize top-most InternalNode and its children (two InternalNodes)
	parent, degree := newInternalNode([]int{5}, []interface{}{
		newInternalNode([]int{4}, nil),
		newInternalNode([]int{10}, nil),
	}), 3
	leftChild, rightChild := parent.children[0].(*InternalNode), parent.children[1].(*InternalNode)
	leftChild.parent, rightChild.parent = parent, parent

	// initialize left child
	leftChild.children = []interface{}{
		newLeafNode([]int{1}, []int{6}, nil, nil, leftChild),
		newLeafNode([]int{4, 5}, []int{2, 10}, nil, nil, leftChild),
	}
	leftChild.children[0].(*LeafNode).rightSibling, leftChild.children[1].(*LeafNode).leftSibling =
		leftChild.children[1].(*LeafNode), leftChild.children[0].(*LeafNode)

	// initialize right child
	rightChild.children = []interface{}{
		newLeafNode([]int{7, 8}, []int{5, 1}, nil, nil, rightChild),
		newLeafNode([]int{10}, []int{1}, nil, nil, rightChild),
	}
	rightChild.children[0].(*LeafNode).rightSibling, rightChild.children[1].(*LeafNode).leftSibling =
		rightChild.children[1].(*LeafNode), rightChild.children[0].(*LeafNode)

	// return data
	return parent, degree
}

func TestInternalNode_Search(t *testing.T) {
	// initialize test
	node, _ := newTestInternalNode()

	leaf, err := node.Search(4)
	if err != nil {
		t.Error("Should not raise exception here")
	}
	if leaf != node.children[0].(*InternalNode).children[1] {
		t.Errorf("The internal node does not give correct leaf node. Expected leaf node [4,5] - get %s", arrayToString(leaf.keys))
	}

	leaf, err = node.Search(9)
	if err != nil {
		t.Error("Should not raise exception here")
	}
	if leaf != node.children[1].(*InternalNode).children[0] {
		t.Errorf("The internal node does not give correct leaf node. Expected leaf node [6,7] - get %s", arrayToString(leaf.keys))
	}

	leaf, err = node.Search(10)
	if err != nil {
		t.Error("Should not raise exception here")
	}
	if leaf != node.children[1].(*InternalNode).children[1] {
		t.Errorf("The internal node does not give correct leaf node. Expected leaf node [10] - get %s", arrayToString(leaf.keys))
	}
}

func TestInternalNode_Insert(t *testing.T) {
	node, degree := newInternalNode([]int{}, []interface{}{2}), 3

	err := node.Insert(3, 4, degree)
	if err != nil {
		t.Error("Should not raise exception here")
	}
	err = node.Insert(3, 10, degree)
	if err == nil {
		t.Error("Should raise exception here")
	}

	err = node.Insert(1, 5, degree)
	if err != nil {
		t.Error("Should not raise exception here")
	}
	if !reflect.DeepEqual(node.keys, []int{1, 3}) {
		t.Errorf("Expected keys are [1, 3], get %s", arrayToString(node.keys))
	}
	if !reflect.DeepEqual(node.children, []interface{}{2, 5, 4}) {
		t.Errorf("Expected children are [5, 4], get %s", arrayToString(node.keys))
	}

	err = node.Insert(7, 12, degree)
	if err != nil {
		t.Error("Should not raise exception here")
	}
	if node.parent == nil {
		t.Error("New parent node should be created")
	}
	if !reflect.DeepEqual(node.keys, []int{1}) {
		t.Errorf("Expected keys are [1], get %s", arrayToString(node.keys))
	}
	if !reflect.DeepEqual(node.parent.keys, []int{3}) {
		t.Errorf("Expected keys are [3], get %s", arrayToString(node.parent.keys))
	}
	rightSibling := node.parent.children[1].(*InternalNode)
	if !reflect.DeepEqual(rightSibling.keys, []int{7}) {
		t.Errorf("Expected keys are [7], get %s", arrayToString(rightSibling.keys))
	}
	if node.parent != rightSibling.parent {
		t.Error("Two internal nodes should have the same parent node")
	}

	err = rightSibling.Insert(4, 1, degree)
	if err != nil {
		t.Error("Should not raise exception here")
	}
	if !reflect.DeepEqual(rightSibling.keys, []int{4, 7}) {
		t.Errorf("Expected keys are [4, 7], get %s", arrayToString(rightSibling.keys))
	}

	err = rightSibling.Insert(8, 4, degree)
	if err != nil {
		t.Error("Should not raise exception here")
	}
	if !reflect.DeepEqual(rightSibling.keys, []int{4}) {
		t.Errorf("Expected keys are [4], get %s", arrayToString(rightSibling.keys))
	}
	if !reflect.DeepEqual(rightSibling.parent.keys, []int{3, 7}) {
		t.Errorf("Expected keys are [3, 7], get %s", arrayToString(rightSibling.keys))
	}
	rightMostSib := node.parent.children[2].(*InternalNode)
	if !reflect.DeepEqual(rightMostSib.keys, []int{8}) {
		t.Errorf("Expected keys are [8], get %s", arrayToString(rightMostSib.keys))
	}
}
