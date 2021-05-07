package btree

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

func (cursor *Cursor) rebalance() error {
	for cursor.depth > 0 {
		parentNode := cursor.nodes[cursor.depth-1].(InternalNode)
		currentIdx := cursor.indices[cursor.depth-1]

		switch node := cursor.nodes[cursor.depth].(type) {
		case *InternalNode:
			if len(node.keys) < cursor.tree.degree/2 {
				if cursor.indices[cursor.depth-1] > 0 {
					sibling := parentNode.children[currentIdx-1].(InternalNode)
					lastIdx := len(sibling.keys) - (len(sibling.keys)-len(node.keys))/2
					if len(sibling.keys)+len(node.keys) < cursor.tree.degree {
						lastIdx = 0
					}
					sibling.keys = append(sibling.keys, parentNode.keys[currentIdx-1])
					node.keys = append(sibling.keys[lastIdx:], node.keys...)
					node.children = append(sibling.children[lastIdx:], node.children...)
					if lastIdx > 0 {
						parentNode.keys[currentIdx-1] = sibling.keys[lastIdx-1]
						sibling.keys = sibling.keys[:lastIdx-1]
						sibling.children = sibling.children[:lastIdx]
					} else {
						parentNode.keys = removeInt(parentNode.keys, currentIdx-1)
						parentNode.children = removeInterface(parentNode.children, currentIdx-1)
					}
				} else {
					sibling := parentNode.children[currentIdx+1].(InternalNode)
					firstIdx := (len(sibling.keys) - len(node.keys)) / 2
					if len(sibling.keys)+len(node.keys) < cursor.tree.degree {
						firstIdx = len(sibling.keys)
					}
					node.keys = append(append(node.keys, parentNode.keys[currentIdx]), sibling.keys[:firstIdx]...)
					node.children = append(node.children, sibling.children[:firstIdx]...)

					if firstIdx < len(sibling.keys) {
						parentNode.keys[currentIdx-1] = node.keys[len(node.children)]
						node.keys = node.keys[:len(node.children)]
						sibling.keys = sibling.keys[firstIdx:]
						sibling.children = sibling.children[firstIdx:]
					} else {
						parentNode.keys = removeInt(parentNode.keys, currentIdx+1)
						parentNode.children = removeInterface(parentNode.children, currentIdx+1)
					}
				}
			} else if len(node.keys) > cursor.tree.degree {
				// split current internal page in half
				half := len(node.keys) / 2
				sibling := newInternalNode(node.keys[:half-1], node.children[:half])
				parentNode.keys = insertInt(parentNode.keys, currentIdx, node.keys[half-1])
				parentNode.children = insertInterface(parentNode.children, currentIdx, sibling)
				node.keys, node.children = node.keys[half:], node.children[half:]
			} else {
				return nil
			}
			break
		case *LeafNode:
			if len(node.keys) < cursor.tree.degree/2 {
				// current node has less entries than the required degree, try to borrow some from its siblings
				if cursor.indices[cursor.depth-1] > 0 {
					sibling := parentNode.children[currentIdx-1].(LeafNode)
					lastIdx := len(sibling.keys) - (len(sibling.keys)-len(node.keys))/2
					if len(sibling.keys)+len(node.keys) < cursor.tree.degree {
						parentNode.keys = removeInt(parentNode.keys, currentIdx-1)
						parentNode.children = removeInterface(parentNode.children, currentIdx-1)
						lastIdx = 0
					}
					node.keys = append(sibling.keys[lastIdx:], node.keys...)
					node.values = append(sibling.values[lastIdx:], node.values...)
					sibling.keys = sibling.keys[:lastIdx]
					sibling.values = sibling.values[:lastIdx]
				} else {
					sibling := parentNode.children[currentIdx+1].(LeafNode)
					firstIdx := (len(sibling.keys) - len(node.keys)) / 2
					if len(sibling.keys)+len(node.keys) < cursor.tree.degree {
						parentNode.keys = removeInt(parentNode.keys, currentIdx)
						parentNode.children = removeInterface(parentNode.children, currentIdx)
						firstIdx = len(sibling.keys)
					}
					node.keys = append(node.keys, sibling.keys[:firstIdx]...)
					node.values = append(node.values, sibling.values[:firstIdx]...)
					sibling.keys = sibling.keys[firstIdx:]
					sibling.values = sibling.values[firstIdx:]
				}
			} else if len(node.keys) > cursor.tree.degree {
				// current node has more entries than tree's degree, split current page in half
				half := len(node.keys) / 2
				sibling := newLeafNode(node.keys[:half], node.values[:half])
				node.keys, node.values = node.keys[half:], node.values[half:]
				parentNode.keys = insertInt(parentNode.keys, currentIdx, sibling.keys[half-1])
				parentNode.children = insertInterface(parentNode.children, currentIdx, sibling)
			} else {
				return nil
			}
			break
		default:
			return fmt.Errorf("rebalance ops found invalid node - its type is %s", reflect.TypeOf(cursor.nodes[cursor.depth]).String())
		}
	}
	return nil
}

func (cursor *Cursor) Search(key int) int {
	result, chosenIndex := -1, -1
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
		default:
			chosenIndex = -1
		}

		cursor.indices = append(cursor.indices, max(chosenIndex, 0))
		if chosenIndex < 0 || isLeafNode(loopNode) {
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

	leaf, _ := cursor.nodes[cursor.depth].(*LeafNode)
	leaf.keys = removeInt(leaf.keys, cursor.indices[cursor.depth])
	leaf.values = removeInt(leaf.values, cursor.indices[cursor.depth])

	if len(leaf.keys) < cursor.tree.degree/2 {
		return cursor.rebalance()
	}
	return nil
}

func (cursor *Cursor) Reset() {
	cursor.depth = -1
	cursor.indices = nil
	cursor.nodes = nil
}
