package btree

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"
)

const ENABLE_LOG bool = false
const NO_KEYS int = 100

func TestInsertAndSearch(t *testing.T) {
	if !ENABLE_LOG {
		log.SetOutput(ioutil.Discard)
	}
	tree := newBPlusTree(1)
	cursor := newCursor(tree)
	result := make(map[int]int)

	for idx := 1; idx <= NO_KEYS; idx++ {
		result[idx] = rand.Int()%1000 + 1
		err := cursor.Insert(idx, result[idx])
		if err != nil {
			t.Error(fmt.Sprintf("Should not throw any exception. Found: %s", err.Error()))
		}
		cursor.Reset()
	}

	for idx := 1; idx <= NO_KEYS; idx++ {
		value := cursor.Search(idx)
		if value < 0 || value != result[idx] {
			t.Error(fmt.Sprintf("Key %d is either not found or mapped to wrong value. Found %d - expected %d", idx, value, result[idx]))
		}
		cursor.Reset()
	}
}

func TestDelete(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	tree := newBPlusTree(1)
	cursor := newCursor(tree)
	result := make(map[int]int)
	removed := make(map[int]bool)

	for idx := 1; idx <= NO_KEYS; idx++ {
		result[idx] = rand.Int()%1000 + 1
		err := cursor.Insert(idx, result[idx])
		if err != nil {
			t.Error(fmt.Sprintf("Should not throw any exception. Found: %s", err.Error()))
		}
		cursor.Reset()
	}

	if ENABLE_LOG {
		log.SetOutput(os.Stdout)
	}

	for idx := 1; idx <= NO_KEYS/2; idx++ {
		item := -1
		for item <= 0 || removed[item] {
			item = rand.Int()%NO_KEYS + 1
		}
		removed[item] = true
		log.Printf("--------Remove key %d------------\n", item)
		err := cursor.Delete(item)
		if err != nil {
			t.Error(fmt.Sprintf("Delete %d. Should not throw any exception. Found: %s", item, err.Error()))
		}
		cursor.Reset()
	}

	for key, _ := range result {
		if !removed[key] {
			value := cursor.Search(key)
			if value < 0 || value != result[key] {
				t.Error(fmt.Sprintf("Key %d is either not found or mapped to wrong value. Found %d - expected %d", key, value, result[key]))
			}
			cursor.Reset()
		}
	}
}
