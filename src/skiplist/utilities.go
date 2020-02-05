package skiplist

import (
	"fmt"
	"math/rand"
)

func getNewHeight() int {
	result := 0
	for rand.Int()%2 == 0 {
		result += 1
	}
	return result
}

func (node *Node) getColumnHeight() int {
	result := 0
	for ; node.up != nil; node = node.up {
	}
	for ; node.down != nil; node = node.down {
		result++
	}
	return result
}

func (list *SkipList) getFirstLevelKeys() []int {
	current, result := list.head, []int{}
	for ; current.down != nil; current = current.down {
	}
	if current.right == nil {
		return result
	}
	for current = current.right; current != nil; current = current.right {
		result = append(result, current.key)
	}
	return result
}

func (list *SkipList) logAllList() {
	current := list.head
	for ; current.down != nil; current = current.down {
	}
	fmt.Println("======")
	for current = current.right; current != nil; current = current.right {
		fmt.Printf("Key %d - value %d - height %d\n", current.key, current.value, current.getColumnHeight())
	}
	fmt.Println("======")
}

func (list *SkipList) searchNode(key int) *Node {
	node := list.head
	for node != nil {
		for node.right != nil && node.right.key <= key {
			node = node.right
		}
		if node.down == nil {
			break
		}
		node = node.down
	}
	return node
}
