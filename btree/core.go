package btree

import (
	"errors"
	"reflect"
	"sort"
)

func (cursor *Cursor) Search(key int) int {
	result, chosenIndex := -1, -1
	loopNode := cursor.tree.root

	for {
		cursor.depth++
		switch node := loopNode.(type) {
		case *InternalNode:
			chosenIndex := sort.SearchInts(node.keys, key)
			loopNode = node.children[chosenIndex]
		case *LeafNode:
			chosenIndex := sort.SearchInts(node.keys, key)
			if node.keys[chosenIndex] == key {
				result = node.values[chosenIndex]
			}
		default:
			chosenIndex = -1
		}
		cursor.indices[cursor.depth] = chosenIndex
		if chosenIndex < 0 || reflect.TypeOf(loopNode).String() == "LeafNode" {
			break
		}
	}

	return result
}

func (cursor *Cursor) rebalance() error {
	return nil
}

func (cursor *Cursor) Insert(key int, value int) error {
	if cursor.depth < 0 {
		found := cursor.Search(key)
		if found >= 0 {
			return errors.New("Insert error: Key exists")
		}
	}

	leaf, _ := cursor.nodes[cursor.depth].(LeafNode)
	insertInt(leaf.keys, cursor.indices[cursor.depth], key)
	insertInt(leaf.values, cursor.indices[cursor.depth], value)

	if len(leaf.keys) >= cursor.tree.degree {
		return cursor.rebalance()
	}
	return nil
}

func (cursor *Cursor) Delete(key int) error {
	if cursor.depth < 0 {
		found := cursor.Search(key)
		if found < 0 {
			return errors.New("Delete error: Key not exists")
		}
	}

	leaf, _ := cursor.nodes[cursor.depth].(LeafNode)
	removeInt(leaf.keys, cursor.indices[cursor.depth])
	removeInt(leaf.values, cursor.indices[cursor.depth])

	if len(leaf.keys) < cursor.tree.degree/2 {
		return cursor.rebalance()
	}
	return nil
}
