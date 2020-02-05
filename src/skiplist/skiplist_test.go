package skiplist

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func arrayToString(a []int) string {
	return strings.Replace(fmt.Sprint(a), " ", ",", -1)
}

func TestSkipListCorrectness(t *testing.T) {
	list := NewSkipList()
	expected := make(map[int]int)
	expectedKeys := []int{}
	for index := 0; index < 100; index++ {
		newKey, newVal := rand.Int()%10+1, rand.Int()%1000
		if _, ok := expected[newKey]; !ok {
			expected[newKey], expectedKeys = newVal, append(expectedKeys, newKey)
			sort.Ints(expectedKeys)
		}
		//fmt.Printf("Try to insert (%d, %d)\n", newKey, newVal)
		err := list.Insert(newKey, newVal)
		if err != nil {
			value, err := list.Search(newKey)
			if err != nil {
				t.Errorf("Key %d should be found in the SkipList. Meet exception: %s", newKey, err.Error())
			}
			if value != expected[newKey] {
				t.Errorf("Key %d should be mapped to value %d instead of %d", newKey, expected[newKey], value)
			}
		} else {
			if !reflect.DeepEqual(expectedKeys, list.getFirstLevelKeys()) {
				t.Errorf("Expected keys should be %s instead of %s", arrayToString(expectedKeys), arrayToString(list.getFirstLevelKeys()))
			}
			insertedCol := list.searchNode(newKey)
			if insertedCol.key != newKey {
				t.Errorf("Search operation does not return correct column. Search %d but found %d", newKey, insertedCol.key)
			}
			if insertedCol.getColumnHeight() > list.maxLevel {
				t.Errorf("A height of a column should never be higher than the height of a SkipList")
			}
		}
		//list.logAllList()
	}
}
