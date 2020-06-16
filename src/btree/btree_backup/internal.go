package btree

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

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

// this newChild should sit next to newKey
func (node *InternalNode) insertKey(newKey int, newChild interface{}, degree int) error {
	index := sort.SearchInts(node.keys, newKey)
	node.keys = insertInt(node.keys, index, newKey)
	node.children = insertInterface(node.children, index, newChild)
	if len(node.keys) <= 2*degree-1 {
		return nil
	}
	sibling := &InternalNode{parent: node.parent, keys: node.keys[degree:], children: node.children[degree:]}
	node.keys, node.children, node.rightSibling = node.keys[:degree], node.children[:degree+1], sibling
	if node.parent == nil {
		parent := &InternalNode{keys: []int{sibling.keys[0]}}
		parent.children = append(parent.children, node, sibling)
		node.parent, sibling.parent = parent, parent
		return nil
	}
	return node.parent.insertKey(sibling.keys[0], sibling, degree)
}
