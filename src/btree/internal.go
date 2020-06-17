package btree

import (
	"fmt"
	"log"
	"reflect"
	"sort"
)

func (node InternalNode) Search(key int) (*LeafNode, error) {
	if len(node.keys)+1 != len(node.children) {
		return nil, fmt.Errorf("There is an internal node in failed state: %d keys and %d children", len(node.keys), len(node.children))
	}
	chosenIndex := sort.SearchInts(node.keys, key)
	switch child := node.children[chosenIndex].(type) {
	case *LeafNode:
		return child, nil
	case *InternalNode:
		return child.Search(key)
	default:
		return nil, fmt.Errorf("Class of a node should be LeafNode or InternalNode instead of %s", reflect.TypeOf(child).String())
	}
}

func (node *InternalNode) Insert(newKey int, newChild interface{}, degree int) error {
	log.Println("InternalNode - Insert", newKey, newChild, degree)
	return nil
}

func (node *InternalNode) Delete(key int, degree int) error {
	log.Println("InternalNode - Delete", key, degree)
	return nil
}
