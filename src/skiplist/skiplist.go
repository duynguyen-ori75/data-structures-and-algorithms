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

type SkipList struct {
	maxLevel int
	// top-left head
	head *Node
}

func NewSkipList() *SkipList {
	return &SkipList{maxLevel: 0, head: &Node{}}
}

func (list *SkipList) Search(key int) (int, error) {
	node := list.searchNode(key)
	if node.key != key {
		return 0, fmt.Errorf("Key %d not found", key)
	}
	return node.value, nil
}

func (list *SkipList) Insert(key int, value int) error {
	if key <= 0 {
		return errors.New("All keys should be positive")
	}
	latestHeads, node := make([]*Node, list.maxLevel+1), list.head
	for currentLevel := list.maxLevel; currentLevel >= 0; currentLevel-- {
		for node.right != nil && node.right.key <= key {
			node = node.right
		}
		latestHeads[currentLevel] = node
		if currentLevel > 0 {
			node = node.down
		}
	}
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
