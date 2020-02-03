package btree

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// this BPlusTree requires all keys to be unique
type LeafNode struct {
	keys         []int
	values       []int
	rightSibling *LeafNode
}

func (leaf LeafNode) getValue(key int) (int, error) {
	if len(leaf.keys) != len(leaf.values) {
		return 0, errors.New("LeafNode's keys and values should have similar number of items")
	}
	index := sort.SearchInts(leaf.keys, key)
	if index == len(leaf.keys) || leaf.keys[index] != key {
		return 0, errors.New(fmt.Sprintf("Key %d not found", key))
	}
	return leaf.values[index], nil
}

type InternalNode struct {
	isLeaf bool
	keys   []int
	// children can be a slice of pointers to LeafNode or InternalNode
	children []interface{}
}

func (node *InternalNode) searchPossibleLeafNode(key int) (*LeafNode, error) {
	if len(node.keys)+1 != len(node.children) {
		return nil, errors.New(fmt.Sprintf("There is an internal node in failed state: %d keys and %d children", len(node.keys), len(node.children)))
	}
	chosenIndex := sort.SearchInts(node.keys, key)
	switch child := node.children[chosenIndex].(type) {
	case *LeafNode:
		return child, nil
	case *InternalNode:
		return child.searchPossibleLeafNode(key)
	default:
		return nil, errors.New(fmt.Sprintf("Class of a node should be LeafNode or InternalNode insted of %s", reflect.TypeOf(child).String()))
	}
}

type BPlusTree struct {
	root   *InternalNode
	degree int
}

func (tree *BPlusTree) search(key int) (*LeafNode, error) {
	possibleLeaf, error := tree.root.searchPossibleLeafNode(key)
	if error != nil {
		return nil, error
	}
	_, error = possibleLeaf.getValue(key)
	if error != nil {
		return nil, error
	}
	return possibleLeaf, nil
}
