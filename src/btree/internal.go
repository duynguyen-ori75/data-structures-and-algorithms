package btree

import (
	"log"
)

func (node *InternalNode) Insert(newKey int, newChild interface{}, degree int) error {
	log.Println("InternalNode - Insert", newKey, newChild, degree)
	return nil
}

func (node *InternalNode) Delete(newKey int, degree int) error {
	log.Println("InternalNode - Delete", newKey, degree)
	return nil
}
