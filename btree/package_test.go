package btree

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := newBPlusTree(2)
	cursor := newCursor(tree)
	result := make(map[int]int)

	for idx := 1; idx <= 20; idx++ {
		result[idx] = rand.Int()%1000 + 1
		err := cursor.Insert(idx, result[idx])
		if err != nil {
			t.Error(fmt.Sprintf("Should not throw any exception. Found: %s", err.Error()))
		}
		cursor.Reset()
	}

	for idx := 1; idx <= 20; idx++ {
		value := cursor.Search(idx)
		if value < 0 || value != result[idx] {
			t.Error(fmt.Sprintf("Key %d is either not found or mapped to wrong value. Found %d - expected %d", idx, value, result[idx]))
		}
		cursor.Reset()
	}
}
