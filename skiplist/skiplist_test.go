package skiplist

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"testing"
)

const BENCHMARK_LIMIT = 1000
const KEY_MAXVAL = 100000000

func arrayToString(a []int) string {
	return strings.Replace(fmt.Sprint(a), " ", ",", -1)
}

func TestSkipListPointersCorrectness(t *testing.T) {
	list := NewSkipListPointers()
	expected := make(map[int]int)
	expectedKeys := []int{}
	for index := 0; index < 500; index++ {
		newKey, newVal := rand.Intn(1000)+1, rand.Intn(1000)
		if _, ok := expected[newKey]; !ok {
			expected[newKey], expectedKeys = newVal, append(expectedKeys, newKey)
			sort.Ints(expectedKeys)
		}
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
	}
	if list.head.getColumnHeight() != list.maxLevel {
		t.Errorf("Height of head column (which is %d) should be equal to maxLevel(%d)", list.head.getColumnHeight(), list.maxLevel)
	}
	for index := 0; index < 200; index++ {
		chosenKeyIndex := rand.Intn(len(expectedKeys))
		nodeBefore, currentNode, removedKey := list.getLevelZeroHead(), list.searchNode(expectedKeys[chosenKeyIndex]), expectedKeys[chosenKeyIndex]
		if chosenKeyIndex > 0 {
			nodeBefore = list.searchNode(expectedKeys[chosenKeyIndex-1])
		}
		list.Remove(removedKey)
		expectedKeys = append(expectedKeys[:chosenKeyIndex], expectedKeys[chosenKeyIndex+1:]...)
		level := 0
		for nodeBefore != nil && currentNode != nil {
			if nodeBefore.right != currentNode.right {
				if nodeBefore.right == nil {
					t.Errorf("NodeBefore(%d) should point to %d instead of the end of the SkipList", nodeBefore.key, currentNode.right.key)
				} else if currentNode.right == nil {
					t.Errorf("NodeBefore(%d) should point to the end of the SkipList instead of key %d", nodeBefore.key, nodeBefore.right.key)
				} else {
					t.Errorf("Remove key %d does not work correctly. Next of nodeBefore(%d - level %d) should be %d instead of %d",
						removedKey, nodeBefore.key, level, currentNode.right.key, nodeBefore.right.key)
				}
			}
			nodeBefore, currentNode, level = nodeBefore.up, currentNode.up, level+1
		}
		_, err := list.Search(removedKey)
		if err == nil {
			t.Errorf("After the removal, key %d should not be found", removedKey)
		}
	}
}

func TestSkipListInsert(t *testing.T) {
	list := NewSkipList()
	err := list.Insert(0, 2)
	if err == nil {
		t.Error("Should raise exception here")
	}
	err = list.Insert(2, 5)
	if err != nil {
		t.Errorf("Should not raise exception. Meet: %s", err)
	}
	err = list.Insert(2, 10)
	if err == nil {
		t.Error("Should raise exception here")
	}
	list.Insert(6, 2)
	list.Insert(3, 1)
	list.Insert(15, 4)
	val, err := list.Search(4)
	if err == nil {
		t.Error("Should raise exception here")
	}
	val, err = list.Search(15)
	if err != nil {
		t.Errorf("Should not raise exception here. Meet: %s", err)
	}
	if val != 4 {
		t.Errorf("Returned value should be 4 instead of %d", val)
	}
	val, err = list.Search(3)
	if err != nil {
		t.Errorf("Should not raise exception here. Meet: %s", err)
	}
	if val != 1 {
		t.Errorf("Returned value should be 1 instead of %d", val)
	}
}

func TestSkipListRemove(t *testing.T) {
	list := NewSkipList()
	expectedMap := make(map[int]int)
	for idx := 1; idx < 100; idx++ {
		key, value := idx, rand.Intn(100)
		list.Insert(key, value)
		expectedMap[key] = value
	}
	if !sort.IntsAreSorted(list.getFirstLevelKeys()) {
		t.Errorf("First-level keys should be sorted. Found: %s", arrayToString(list.getFirstLevelKeys()))
	}
	for idx := 1; idx < 50; idx++ {
		selectedKey := rand.Intn(100)
		expectedVal, expectedOk := expectedMap[selectedKey]
		val, err := list.Search(selectedKey)
		if err != nil {
			if expectedOk {
				t.Errorf("Key %d should exists with value: %d", selectedKey, expectedVal)
			}
			err = list.Remove(selectedKey)
			if err == nil {
				t.Error("Delete not existing key should throw exception")
			}
		} else {
			if !expectedOk {
				t.Errorf("Key %d should already be deleted. Found value: %d", selectedKey, val)
			} else {
				if expectedVal != val {
					t.Errorf("expectedVal should equal to val. Found (%d, %d), respectively", expectedVal, val)
				}
				delete(expectedMap, selectedKey)
				err = list.Remove(selectedKey)
				if err != nil {
					t.Errorf("Delete existing key should not raise exception. Meet: %s", err)
				}
				val, err = list.Search(selectedKey)
				if err == nil {
					t.Error("Delete not existing key should throw exception")
				}
			}
		}
	}
}

/**
 * Benchmark two versions
 */

func prepareTestData() ([]int, int) {
	keyList, insertedKey := []int{}, make(map[int]bool)
	for idx := 0; idx < BENCHMARK_LIMIT; idx++ {
		newKey := rand.Intn(KEY_MAXVAL)
		for _, found := insertedKey[newKey]; found; newKey = rand.Intn(KEY_MAXVAL) {
		}
		insertedKey[newKey] = true
		keyList = append(keyList, newKey)
	}
	return keyList, BENCHMARK_LIMIT
}

func BenchmarkSkipListPointer(t *testing.B) {
	insertedKeys, length := prepareTestData()
	t.ResetTimer()
	for test := 0; test < t.N; test++ {
		list := NewSkipListPointers()
		for _, key := range insertedKeys {
			list.Insert(key, rand.Intn(length))
		}
		deletedKey := 10
		for idx := 0; idx < length/10; idx++ {
			list.Remove(deletedKey)
			deletedKey += rand.Intn(length/100) + 1
		}
	}
}

func BenchmarkSkipListArray(t *testing.B) {
	insertedKeys, length := prepareTestData()
	t.ResetTimer()
	for test := 0; test < t.N; test++ {
		list := NewSkipList()
		for _, key := range insertedKeys {
			list.Insert(key, rand.Intn(length))
		}
		deletedKey := 10
		for idx := 0; idx < length/10; idx++ {
			list.Remove(deletedKey)
			deletedKey += rand.Intn(length/100) + 1
		}
	}
}
