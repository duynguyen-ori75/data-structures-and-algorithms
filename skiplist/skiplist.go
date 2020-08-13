package skiplist

import (
	"errors"
	"fmt"
)

type NewNode struct {
	key   int
	value int
	level int
	next  []*NewNode
}

type SkipList struct {
	root *NewNode
}

func NewSkipList() *SkipList {
	return &SkipList{root: &NewNode{next: []*NewNode{nil}}}
}

func (list *SkipList) Search(key int) (int, error) {
	node := list.searchNode(key)
	if node.key != key {
		return 0, fmt.Errorf("Key %d not found", key)
	}
	return node.value, nil
}

func (list *SkipList) Insert(key int, value int) error {
	//fmt.Println("--------------Try to insert (", key, value, ")---------------")
	if key <= 0 {
		return errors.New("All keys should be positive")
	}
	latestHeads, node := list.getRightMostNodes(key)
	if node.key == key {
		return fmt.Errorf("Key %d already exists", key)
	}
	newNodeHeight := getNewHeight()
	newNode := &NewNode{key: key, value: value, level: newNodeHeight, next: make([]*NewNode, newNodeHeight+1)}
	if newNodeHeight > list.root.level {
		list.root.next = append(list.root.next, make([]*NewNode, newNodeHeight-list.root.level)...)
		for idx := 0; idx < newNodeHeight-list.root.level; idx++ {
			latestHeads = append(latestHeads, list.root)
		}
		list.root.level = newNodeHeight
	}
	//fmt.Println(latestHeads)
	for level, latestHead := range latestHeads {
		//fmt.Println(level, latestHead)
		if level > newNodeHeight {
			//fmt.Println("Exit now")
			break
		}
		newNode.next[level], latestHead.next[level] = latestHead.next[level], newNode
	}
	return nil
}
