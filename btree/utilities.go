package btree

import (
	"fmt"
	"log"
	"strings"
)

func arrayToString(a []int) string {
	return strings.Replace(fmt.Sprint(a), " ", ",", -1)
}

func interfacesToString(a []interface{}) string {
	return strings.Replace(fmt.Sprint(a), " ", ",", -1)
}

func insertInt(slice []int, index int, newElement int) []int {
	return append(slice[:index], append([]int{newElement}, slice[index:]...)...)
}

func insertInterface(slice []interface{}, index int, newElement interface{}) []interface{} {
	return append(slice[:index], append([]interface{}{newElement}, slice[index:]...)...)
}

func removeInt(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func removeInterface(slice []interface{}, index int) []interface{} {
	return append(slice[:index], slice[index+1:]...)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func isLeafNode(node interface{}) bool {
	switch node.(type) {
	case LeafNode, *LeafNode:
		return true
	default:
		return false
	}
}

func debugTree(tree BPlusTree) {
	type queueItem struct {
		depth int
		node  interface{}
	}
	var top queueItem
	var queue []queueItem
	queue = append(queue, queueItem{depth: 0, node: tree.root})
	log.Println("===============Debug B+tree================")
	for len(queue) > 0 {
		top, queue = queue[0], queue[1:]
		log.Printf("Layer %d: %+v\n", top.depth, top.node)
		if node, ok := top.node.(*InternalNode); ok {
			for _, child := range node.children {
				queue = append(queue, queueItem{depth: top.depth + 1, node: child})
			}
		}
	}
	log.Println("===========================================")
}
