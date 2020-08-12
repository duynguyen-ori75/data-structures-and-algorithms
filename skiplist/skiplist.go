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
	return &SkipList{head: &NewNode{}}
}

func (list *SkipList) Search(key int) (int, error) {
	node := list.searchNode(key)
	if node.key != key {
		return 0, fmt.Errorf("Key %d not found", key)
	}
	return node.value, nil
}
