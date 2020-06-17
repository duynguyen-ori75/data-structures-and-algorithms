package btree

import (
	//"log"
	//"reflect"
	"fmt"
	"testing"
)

/**
 * @brief      initialize InternalNode with the following structure
 *
 * 				   5
 * 			   /      \
 * 			 3          8
 * 		   /   \      /   \
 * 		  1   4,5   6,7   10
 *
 * @return     the parent InternalNode and the maximum degree of BTree
 */
func newTestInternalNode() (*InternalNode, int) {
	// initialize top-most InternalNode and its children (two InternalNodes)
	parent, degree := newInternalNode([]int{5}, []interface{}{
		newInternalNode([]int{3}, nil),
		newInternalNode([]int{8}, nil),
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
		newLeafNode([]int{6, 7}, []int{5, 1}, nil, nil, rightChild),
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
		fmt.Errorf("The internal node does not give correct leaf node. Expected leaf node [4,5] - get %s", arrayToString(leaf.keys))
	}
	leaf, err = node.Search(8)
	if err != nil {
		t.Error("Should not raise exception here")
	}
	if leaf != node.children[1].(*InternalNode).children[0] {
		fmt.Errorf("The internal node does not give correct leaf node. Expected leaf node [6,7] - get %s", arrayToString(leaf.keys))
	}

}
