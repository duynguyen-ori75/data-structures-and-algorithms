package skiplist

import (
	"errors"
	"fmt"
)

type Node struct {
	up    *Node
	down  *Node
	right *Node
	key   int
	value int
}

type SkipListPointers struct {
	maxLevel int
	head     *Node
}

func NewSkipListPointers() *SkipListPointers {
	return &SkipListPointers{maxLevel: 0, head: &Node{}}
}

func (list *SkipListPointers) Search(key int) (int, error) {
	node := list.searchNode(key)
	if node.key != key {
		return 0, fmt.Errorf("Key %d not found", key)
	}
	return node.value, nil
}

func (list *SkipListPointers) Insert(key int, value int) error {
	if key <= 0 {
		return errors.New("All keys should be positive")
	}
	latestHeads, node := list.getRightMostNodes(key)
	if node.key == key {
		return fmt.Errorf("Key %d already exists", key)
	}
	// current node should be at level 1 (node.down == nil -> level 1)
	nextNodeHeight := getNewHeight()
	for ; list.maxLevel < nextNodeHeight; list.maxLevel++ {
		newHead := &Node{up: nil, down: list.head, right: nil}
		list.head.up = newHead
		list.head, latestHeads = newHead, append(latestHeads, newHead)
	}
	var lowerNode *Node
	for _, latestHead := range latestHeads[:nextNodeHeight+1] {
		newNode := &Node{key: key, value: value, right: latestHead.right, down: lowerNode}
		if lowerNode != nil {
			lowerNode.up = newNode
		}
		latestHead.right, lowerNode = newNode, newNode
	}
	return nil
}

func (list *SkipListPointers) Remove(key int) error {
	latestHeads, _ := list.getRightMostNodes(key - 1)
	for currentLevel, latestHead := range latestHeads {
		node := latestHead.right
		if node == nil || node.key != key {
			if currentLevel == 0 {
				return fmt.Errorf("Key %d not found", key)
			}
			break
		}
		latestHead.right = node.right
	}
	return nil
}
