package btree

import (
	"fmt"
)

func (node *InternalNode) Insert(newKey int, newChild interface{}, degree int) error {
	fmt.Println(newKey, newChild, degree)
	return nil
}
