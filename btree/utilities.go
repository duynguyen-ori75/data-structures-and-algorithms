package btree

import (
	"fmt"
	"strings"
)

func insertInt(slice []int, index int, newElement int) []int {
	return append(slice[:index], append([]int{newElement}, slice[index:]...)...)
}

func insertInterface(slice []interface{}, index int, newElement interface{}) []interface{} {
	return append(slice[:index], append([]interface{}{newElement}, slice[index:]...)...)
}

func newLeafNode(keys []int, values []int, parent *InternalNode) *LeafNode {
	return &LeafNode{keys: keys, values: values, parent: parent}
}

func arrayToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}
