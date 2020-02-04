package btree

import (
	"fmt"
	"strings"
)

func arrayToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}

func insertInt(slice []int, index int, newElement int) []int {
	return append(slice[:index], append([]int{newElement}, slice[index:]...)...)
}

func insertInterface(slice []interface{}, index int, newElement interface{}) []interface{} {
	return append(slice[:index], append([]interface{}{newElement}, slice[index:]...)...)
}

func newLeafNode(keys []int, values []int, parent *InternalNode) *LeafNode {
	return &LeafNode{keys: keys, values: values, parent: parent}
}

func newBPlusTree(degree int) *BPlusTree {
	return &BPlusTree{root: newLeafNode([]int{}, []int{}, nil), degree: degree}
}
