package btree

import (
	"errors"
	"fmt"
	"reflect"
)

// this should return the leaf node that the key should belong to
func (tree *BPlusTree) Search(key int) (*LeafNode, error) {
	switch root := tree.root.(type) {
	case *InternalNode:
		return root.Search(key)
	case *LeafNode:
		return root, nil
	default:
		return nil, fmt.Errorf("Class of a node should be neither LeafNode or InternalNode insted of %s", 
			reflect.TypeOf(root).String())
	}
	return nil, errors.New("This line should not be executed")
}

func (tree *BPlusTree) Insert(key int, value int) error {
	leaf, err := tree.Search(key)
	if err != nil {
		return err
	}
	err = leaf.Insert(key, value, tree.degree)
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

func (tree *BPlusTree) Delete(key int) error {
	leaf, err := tree.Search(key)
	if err != nil {
		return err
	}
	return leaf.Delete(key, tree.degree)
}
