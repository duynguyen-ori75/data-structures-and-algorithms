package btree

import (
	"fmt"
	"log"
	"reflect"
	"sort"
)

func (node InternalNode) Search(key int) (*LeafNode, error) {
	log.Printf("DFS to node %s\n", arrayToString(node.keys))
	if len(node.keys)+1 != len(node.children) {
		return nil, fmt.Errorf("There is an internal node in failed state: %d keys and %d children", len(node.keys), len(node.children))
	}
	chosenIndex := sort.SearchInts(node.keys, key+1)
	switch child := node.children[chosenIndex].(type) {
	case *LeafNode:
		return child, nil
	case *InternalNode:
		return child.Search(key)
	default:
		return nil, fmt.Errorf("Class of a node should be LeafNode or InternalNode instead of %s", reflect.TypeOf(child).String())
	}
}

func (node *InternalNode) Insert(key int, newChild interface{}, degree int) error {
	index := sort.SearchInts(node.keys, key)
	if index < len(node.keys) && node.keys[index] == key {
		return fmt.Errorf("InternalNode: All keys should be unique. Try to insert %d into %s", key, arrayToString(node.keys))
	}
	node.keys, node.children = insertInt(node.keys, index, key), insertInterface(node.children, index, newChild)
	// number of keys is still lower than the maximum number
	if len(node.keys) < degree {
		return nil
	}
	// otherwise, split the internal node, move up the middle key into parent
	numberOfKeys := degree / 2 // numberOfKeys is also the index of key that we should push to the parent
	moveUpKey := node.keys[numberOfKeys]
	rightSibl := newInternalNode(node.keys[numberOfKeys+1:], node.children[numberOfKeys+1:])
	node.keys, node.children = node.keys[:numberOfKeys], node.children[:numberOfKeys+1]
	// if parent is nil -> this internal node is the top-most node -> create new internal node on top of it
	if node.parent == nil {
		node.parent = newInternalNode([]int{}, []interface{}{})
	}
	rightSibl.parent = node.parent
	return node.parent.Insert(moveUpKey, rightSibl, degree)
}

func (node *InternalNode) Delete(key int, degree int) error {
	log.Println("InternalNode - Delete", key, degree)
	return nil
}
