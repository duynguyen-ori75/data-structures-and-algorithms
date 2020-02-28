package btree

import (
	"errors"
	"fmt"
	"sort"
)

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
	leaf.keys, leaf.values = insertInt(leaf.keys, index, key), insertInt(leaf.values, index, value)
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

func (leaf *LeafNode) deleteKey(key int, degree int) error {
	index := sort.SearchInts(leaf.keys, key)
	if index >= len(leaf.keys) || leaf.keys[index] != key {
		return fmt.Errorf("Key %d not found", key)
	}
	leaf.keys, leaf.values = removeInt(leaf.keys, index), removeInt(leaf.values, index)
	return nil
}
