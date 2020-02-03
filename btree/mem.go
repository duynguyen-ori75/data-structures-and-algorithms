package btree

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// helper stuffs
func insertElement(slice []int, index int, newElement int) []int {
	return append(slice[:index], append([]int{newElement}, slice[index:]...)...)
}

// LeafNode struct
type LeafNode struct {
	keys         []int
	values       []int
	rightSibling *LeafNode
	parent       *InternalNode
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

func (leaf *LeafNode) insertValue(key int, value int, degree int) error {
	index := sort.SearchInts(leaf.keys, key)
	if index < len(leaf.keys) && leaf.keys[index] == key {
		return errors.New("All keys should be unique")
	}
	leaf.keys = insertElement(leaf.keys, index, key)
	leaf.values = insertElement(leaf.values, index, value)
	if len(leaf.keys) <= 2*degree-1 {
		return nil
	}
	sibling := &LeafNode{keys: leaf.keys[degree:], values: leaf.values[degree:], rightSibling: leaf.rightSibling, parent: leaf.parent}
	leaf.rightSibling, leaf.keys, leaf.values = sibling, leaf.keys[:degree], leaf.values[:degree]
	if leaf.parent == nil {
		parent := &InternalNode{keys: []int{sibling.keys[0]}}
		parent.children = append(parent.children, leaf, sibling)
		leaf.parent, sibling.parent = parent, parent
		return nil
	}
	return leaf.parent.insertKey(sibling.keys[0], sibling, degree)
}

// InternalNode struct
type InternalNode struct {
	parent *InternalNode
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
		return nil, errors.New(fmt.Sprintf("Class of a node should be LeafNode or InternalNode instead of %s", reflect.TypeOf(child).String()))
	}
}

// this newChild should be on the right size of newKey
func (node *InternalNode) insertKey(newKey int, newChild interface{}, degree int) error {
	index := sort.SearchInts(node.keys, newKey)
	node.keys = insertElement(node.keys, index, newKey)
	node.children = append(node.children[:index], append([]interface{}{newChild}, node.children[index:]...)...)
	if len(node.keys) <= 2*degree-1 {
		return nil
	}
	sibling := &InternalNode{parent: node.parent, keys: node.keys[degree:], children: node.children[degree:]}
	node.keys, node.children = node.keys[:degree], node.children[:degree+1]
	if node.parent == nil {
		parent := &InternalNode{keys: []int{sibling.keys[0]}}
		parent.children = append(parent.children, node, sibling)
		node.parent, sibling.parent = parent, parent
		return nil
	}
	return node.parent.insertKey(sibling.keys[0], sibling, degree)
}

// our BPlusTree
type BPlusTree struct {
	root   interface{}
	degree int
}

func (tree *BPlusTree) search(key int) (*LeafNode, error) {
	var leaf *LeafNode
	switch root := tree.root.(type) {
	case *InternalNode:
		var err error
		leaf, err = root.searchPossibleLeafNode(key)
		if err != nil {
			return nil, err
		}
	case *LeafNode:
		leaf = root
	default:
		return nil, errors.New(fmt.Sprintf("Class of a node should be LeafNode or InternalNode insted of %s", reflect.TypeOf(root).String()))
	}
	_, err := leaf.getValue(key)
	if err != nil {
		return nil, err
	}
	return leaf, nil
}

func (tree *BPlusTree) insert(key int, value int) error {
	leaf, err := tree.search(key)
	if err != nil {
		return err
	}
	err = leaf.insertValue(key, value, tree.degree)
	if err != nil {
		return err
	}
	if leaf.parent == nil {
		tree.root = leaf
		return nil
	}
	newParent := leaf.parent
	for newParent.parent != nil {
		newParent = newParent.parent
	}
	tree.root = newParent
	return nil
}
