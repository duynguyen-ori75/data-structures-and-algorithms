package btree

import (
	"fmt"
	//"log"
	"reflect"
	"sort"
)

func (node InternalNode) validation() error {
	//log.Printf("DFS to node %s\n", arrayToString(node.keys))
	if len(node.keys)+1 != len(node.children) {
		return fmt.Errorf("There is an internal node in failed state: %d keys and %d children", len(node.keys), len(node.children))
	}
	return nil
}

func (node InternalNode) Search(key int) (*LeafNode, error) {
	err := node.validation()
	if err != nil {
		return nil, err
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
	err := node.validation()
	if err != nil {
		return err
	}
	index := sort.SearchInts(node.keys, key)
	if index < len(node.keys) && node.keys[index] == key {
		return fmt.Errorf("InternalNode: All keys should be unique. Try to insert %d into %s", key, arrayToString(node.keys))
	}
	node.keys, node.children = insertInt(node.keys, index, key), insertInterface(node.children, index+1, newChild)
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
		node.parent = newInternalNode([]int{}, []interface{}{node})
	}
	rightSibl.parent = node.parent
	return node.parent.Insert(moveUpKey, rightSibl, degree)
}

func (node *InternalNode) Delete(key int, degree int) error {
	err := node.validation()
	if err != nil {
		return err
	}
	index := sort.SearchInts(node.keys, key)
	if index == len(node.keys) || node.keys[index] != key {
		return fmt.Errorf("InternalNode: Can't delete key %d in node %s", key, arrayToString(node.keys))
	}
	node.keys, node.children = removeInt(node.keys, index), removeInterface(node.children, index+1)
	// validate number of keys - should be not lower than the min degree (max degree / 2)
	if len(node.keys) >= degree/2 || node.parent == nil {
		return nil
	}
	/**
	 * Otherwise, there are two cases:
	 * - len(currentNode + rightSibling + 1) < degree  // 1 here mean the corresponding key in parent node
	 * - rotate key and child (node borrows key from parent, and the parent borrows a key from rightSibl)
	 */
	childIndex := sort.SearchInts(node.parent.keys, key)
	rightSibl := node.parent.children[childIndex+1].(*InternalNode)
	if len(node.keys)+len(rightSibl.keys)+1 < degree {
		node.keys = append(node.keys, node.parent.keys[childIndex])
		node.keys = append(node.keys, rightSibl.keys...)
		node.children = append(node.children, rightSibl.children...)
		return node.parent.Delete(childIndex, degree)
	}
	// scenario 3
	// update node.keys first
	node.keys, node.children = append(node.keys, node.parent.keys[childIndex]), append(node.children, rightSibl.children[0])
	// update respective key in node.parent
	node.parent.keys[childIndex] = rightSibl.keys[0]
	// delete first entry of right sibling
	rightSibl.keys, rightSibl.children = removeInt(rightSibl.keys, 0), removeInterface(rightSibl.children, 0)
	return nil
}
