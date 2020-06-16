package btree

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

func (leaf LeafNode) Search(key int) ([]int, error) {
	if len(leaf.keys) != len(leaf.values) {
		return nil, errors.New("LeafNode's keys and values should have similar number of items")
	}
	index := sort.SearchInts(leaf.keys, key)
	if index == len(leaf.keys) || leaf.keys[index] != key {
		return nil, errors.New(fmt.Sprintf("Key %d not found", key))
	}
	var result []int
	for ; leaf.keys[index] == key; index++ {
		result = append(result, leaf.values[index])
	}
	return result, nil
}

func (leaf *LeafNode) Insert(key int, value int, degree int) error {
	index := sort.SearchInts(leaf.keys, key)
	if index < len(leaf.keys) && leaf.keys[index] == key {
		return errors.New("All keys should be unique")
	}
	leaf.keys, leaf.values = insertInt(leaf.keys, index, key), insertInt(leaf.values, index, value)
	// number of keys is still lower than the maximum number
	if len(leaf.keys) < degree {
		return nil
	}
	numberOfKeys := degree / 2
	// otherwise, split the LeafNode and create new InternalNode
	rightSibling := newLeafNode(leaf.keys[numberOfKeys:], leaf.values[numberOfKeys:], leaf, leaf.rightSibling, leaf.parent)
	leaf.rightSibling, leaf.keys, leaf.values = rightSibling, leaf.keys[:numberOfKeys], leaf.values[:numberOfKeys]
	// if parent node is nil -> create new parent node
	if leaf.parent == nil {
		parent := newInternalNode([]int{rightSibling.keys[0]}, []interface{}{leaf, rightSibling})
		leaf.parent, rightSibling.parent = parent, parent
		return nil
	}
	// else, insert the split key into the parent
	return leaf.parent.Insert(rightSibling.keys[0], rightSibling, degree)
}

func (leaf *LeafNode) Delete(key int, degree int) error {
	index := sort.SearchInts(leaf.keys, key)
	if index >= len(leaf.keys) || leaf.keys[index] != key {
		return fmt.Errorf("Key %d not found", key)
	}
	leaf.keys, leaf.values = removeInt(leaf.keys, index), removeInt(leaf.values, index)
	// if len(leaf.keys) == 0 -> delete corresponding key in parent and self-destruct
	if len(leaf.keys) == 0 {
		if leaf.leftSibling != nil {
			leaf.leftSibling.rightSibling = leaf.rightSibling
		}
		if leaf.rightSibling != nil {
			leaf.rightSibling.leftSibling = leaf.leftSibling
		}
		if leaf.parent != nil {
			leaf.parent = nil
			return leaf.parent.Delete(key, degree)
		}
		return nil
	}
	numberOfKeys := degree / 2
	// number of keys is higher than the minimum degree
	// or this leaf node does not have a right sibling to either borrow a key or to be merged with it
	if len(leaf.keys) >= numberOfKeys || leaf.rightSibling == nil {
		// if first index == 0 -> have to update corresponding key in its parent node
		if index == 0 && leaf.parent != nil {
			parentIndex := sort.SearchInts(leaf.parent.keys, key)
			leaf.parent.keys[parentIndex] = leaf.keys[0]
		}
		return nil
	}
	rightSibling := leaf.rightSibling
	borrowedKey := rightSibling.keys[0]
	// borrow a key of right sibling if it has enough keys
	if len(rightSibling.keys) > numberOfKeys {
		leaf.keys, rightSibling.keys = append(leaf.keys, borrowedKey), removeInt(rightSibling.keys, 0)
		leaf.values, rightSibling.values = append(leaf.values, rightSibling.values[0]), removeInt(rightSibling.values, 0)
		// update the corresponding key in the parent node
		parentIndex := sort.SearchInts(leaf.parent.keys, borrowedKey)
		log.Printf("Parent's keys: %s - try to delete key %d at index %d", arrayToString(leaf.parent.keys), borrowedKey, parentIndex)
		leaf.parent.keys[parentIndex] = rightSibling.keys[0]
		return nil
	}
	// otherwise, merge this leaf node with its sibling
	leaf.rightSibling, leaf.keys, leaf.values =
		rightSibling.rightSibling, append(leaf.keys, rightSibling.keys...), append(leaf.values, rightSibling.values...)
	// rebalance the parent node
	return leaf.parent.Delete(key, degree)
}
