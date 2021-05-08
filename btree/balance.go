package btree

import (
	"log"
)

func (cursor *Cursor) balance_root() {
	var newRootNode *InternalNode
	switch root := cursor.nodes[0].(type) {
	case *InternalNode:
		half := len(root.keys) / 2
		sibling := newInternalNode(root.keys[:half], root.children[:half+1])
		newRootNode = newInternalNode([]int{root.keys[half]}, []interface{}{sibling, root})
		root.keys, root.children = root.keys[half+1:], root.children[half+1:]
	case *LeafNode:
		half := len(root.keys) / 2
		sibling := newLeafNode(root.keys[:half+1], root.values[:half+1])
		newRootNode = newInternalNode([]int{root.keys[half]}, []interface{}{sibling, root})
		root.keys, root.values = root.keys[half+1:], root.values[half+1:]
	}
	cursor.tree.root = newRootNode
}

func (cursor *Cursor) balance_nonroot() {
	for cursor.depth > 0 {
		parentNode := cursor.nodes[cursor.depth-1].(*InternalNode)
		currentIdx := cursor.indices[cursor.depth-1]

		switch node := cursor.nodes[cursor.depth].(type) {
		case *InternalNode:
			if len(node.keys) < cursor.tree.degree {
				if cursor.indices[cursor.depth-1] > 0 {
					sibling := parentNode.children[currentIdx-1].(*InternalNode)
					lastIdx := len(sibling.keys) - (len(sibling.keys)-len(node.keys))/2
					if len(sibling.keys)+len(node.keys) <= cursor.tree.degree*2 {
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
					sibling := parentNode.children[currentIdx+1].(*InternalNode)
					firstIdx := (len(sibling.keys) - len(node.keys)) / 2
					if len(sibling.keys)+len(node.keys) <= cursor.tree.degree*2 {
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
			} else if len(node.keys) > cursor.tree.degree*2 {
				// split current internal page in half
				half := len(node.keys) / 2
				sibling := newInternalNode(node.keys[:half], node.children[:half+1])
				parentNode.keys = insertInt(parentNode.keys, currentIdx, node.keys[half])
				parentNode.children = insertInterface(parentNode.children, currentIdx, sibling)
				node.keys, node.children = node.keys[half+1:], node.children[half+1:]
			} else {
				return
			}
		case *LeafNode:
			if len(node.keys) < cursor.tree.degree {
				// current node has less entries than the required degree, try to borrow some from its siblings
				if cursor.indices[cursor.depth-1] > 0 {
					sibling := parentNode.children[currentIdx-1].(*LeafNode)
					lastIdx := len(sibling.keys) - (len(sibling.keys)-len(node.keys))/2
					if len(sibling.keys)+len(node.keys) <= cursor.tree.degree*2 {
						parentNode.keys = removeInt(parentNode.keys, currentIdx-1)
						parentNode.children = removeInterface(parentNode.children, currentIdx-1)
						lastIdx = 0
					}
					node.keys = append(sibling.keys[lastIdx:], node.keys...)
					node.values = append(sibling.values[lastIdx:], node.values...)
					sibling.keys = sibling.keys[:lastIdx]
					sibling.values = sibling.values[:lastIdx]
				} else {
					sibling := parentNode.children[currentIdx+1].(*LeafNode)
					firstIdx := (len(sibling.keys) - len(node.keys)) / 2
					if len(sibling.keys)+len(node.keys) <= cursor.tree.degree*2 {
						parentNode.keys = removeInt(parentNode.keys, currentIdx)
						parentNode.children = removeInterface(parentNode.children, currentIdx)
						firstIdx = len(sibling.keys)
					}
					node.keys = append(node.keys, sibling.keys[:firstIdx]...)
					node.values = append(node.values, sibling.values[:firstIdx]...)
					sibling.keys = sibling.keys[firstIdx:]
					sibling.values = sibling.values[firstIdx:]
				}
			} else if len(node.keys) > cursor.tree.degree/2 {
				// current node has more entries than tree's degree, split current page in half
				half := len(node.keys) / 2
				sibling := newLeafNode(node.keys[:half], node.values[:half])
				node.keys, node.values = node.keys[half:], node.values[half:]
				parentNode.keys = insertInt(parentNode.keys, currentIdx, sibling.keys[half-1])
				parentNode.children = insertInterface(parentNode.children, currentIdx, sibling)
			} else {
				return
			}
		}
		cursor.depth--
		log.Printf("-----after balance depth %d----------\n", cursor.depth+1)
		debugTree(*cursor.tree)
	}
	return
}

func (cursor *Cursor) rebalance() {
	log.Println("------------------------Rebalancing-----------------------")
	debugTree(*cursor.tree)
	if cursor.depth > 0 {
		cursor.balance_nonroot()
	}
	debugTree(*cursor.tree)
	if cursor.depth == 0 {
		node := cursor.nodes[0]
		switch root := node.(type) {
		case *InternalNode:
			if len(root.keys) > cursor.tree.degree*2 {
				cursor.balance_root()
			}
		case *LeafNode:
			if len(root.keys) >= cursor.tree.degree*2 {
				cursor.balance_root()
			}
		}
	}
	debugTree(*cursor.tree)
}
