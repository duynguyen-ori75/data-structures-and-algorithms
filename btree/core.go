package btree

import (
	"errors"
	"sort"
)

func (cursor *Cursor) Search(key int) int {
	result, chosenIndex := -1, 1
	loopNode := cursor.tree.root

	for {
		cursor.depth++
		switch node := loopNode.(type) {
		case *InternalNode:
			chosenIndex = sort.SearchInts(node.keys, key)
			loopNode = node.children[chosenIndex]
			cursor.nodes = append(cursor.nodes, node)
		case *LeafNode:
			chosenIndex = sort.SearchInts(node.keys, key)
			if chosenIndex < len(node.keys) && node.keys[chosenIndex] == key {
				result = node.values[chosenIndex]
			}
			cursor.nodes = append(cursor.nodes, node)
		}

		cursor.indices = append(cursor.indices, max(chosenIndex, 0))
		if isLeafNode(cursor.nodes[cursor.depth]) {
			break
		}
	}

	return result
}

func (cursor *Cursor) Insert(key int, value int) error {
	if cursor.depth < 0 {
		found := cursor.Search(key)
		if found >= 0 {
			return errors.New("Insert error: Key exists")
		}
	}

	leaf, _ := cursor.nodes[cursor.depth].(*LeafNode)
	leaf.keys = insertInt(leaf.keys, cursor.indices[cursor.depth], key)
	leaf.values = insertInt(leaf.values, cursor.indices[cursor.depth], value)

	if len(leaf.keys) > cursor.tree.degree*2 {
		cursor.rebalance()
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

	leaf, _ := cursor.nodes[cursor.depth].(*LeafNode)
	leaf.keys = removeInt(leaf.keys, cursor.indices[cursor.depth])
	leaf.values = removeInt(leaf.values, cursor.indices[cursor.depth])

	if len(leaf.keys) < cursor.tree.degree {
		cursor.rebalance()
	}
	return nil
}

func (cursor *Cursor) Reset() {
	cursor.depth = -1
	cursor.indices = nil
	cursor.nodes = nil
}
