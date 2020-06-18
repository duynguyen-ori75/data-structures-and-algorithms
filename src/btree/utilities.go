package btree

import (
	"fmt"
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

func newLeafNode(keys []int, values []int, leftSibling *LeafNode, rightSibling *LeafNode, parent *InternalNode) *LeafNode {
	return &LeafNode{keys: keys, values: values, leftSibling: leftSibling, rightSibling: rightSibling, parent: parent}
}

func newInternalNode(keys []int, children []interface{}) *InternalNode {
	return &InternalNode{keys: keys, children: children}
}

func newBPlusTree(degree int) *BPlusTree {
	return &BPlusTree{root: newLeafNode([]int{}, []int{}, nil, nil, nil), degree: degree}
}
