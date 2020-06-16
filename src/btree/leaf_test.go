package btree

import (
	"log"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	leaf := newLeafNode([]int{1, 4, 4, 4, 4, 7, 8}, []int{4, 5, 7, 2, 3, 1, 5}, nil, nil)
	val, err := leaf.Search(2)
	if err == nil {
		t.Error("Search for non-existant key should return error")
	}
	val, err = leaf.Search(4)
	if err != nil {
		t.Error("Search of existing key should be good")
	}
	if !reflect.DeepEqual(val, []int{5, 7, 2, 3}) {
		t.Errorf("Expected array should be %s instead of %s", arrayToString([]int{5, 7, 2, 3}), arrayToString(val))
	}
}

func TestInsert(t *testing.T) {
	leaf, degree := newLeafNode([]int{}, []int{}, nil, nil), 3
	log.Printf("Inserting with degree %d\n", degree)
	if leaf.Insert(1, 2, degree) != nil {
		t.Error("Insert (1, 2) should not raise exception")
	}
	if leaf.Insert(3, 4, degree) != nil {
		t.Error("Insert (3, 4) should not raise exception")
	}
	if leaf.rightSibling != nil {
		t.Error("There should be no sibling after inserting 2 keys")
	}
	if leaf.Insert(2, 9, degree) != nil {
		t.Error("Insert (2, 9) should not raise exception")
	}
	if leaf.rightSibling == nil {
		t.Error("There should be a sibling after inserting 3 keys")
	}
	if !reflect.DeepEqual(leaf.keys, []int{1}) {
		t.Errorf("Expected array should be %s instead of %s", arrayToString([]int{1}), arrayToString(leaf.keys))
	}
	if !reflect.DeepEqual(leaf.values, []int{2}) {
		t.Errorf("Expected array should be %s instead of %s", arrayToString([]int{2}), arrayToString(leaf.values))
	}
	if !reflect.DeepEqual(leaf.rightSibling.keys, []int{2, 3}) {
		t.Errorf("Expected array should be %s instead of %s", arrayToString([]int{2, 3}), arrayToString(leaf.rightSibling.keys))
	}
	if !reflect.DeepEqual(leaf.rightSibling.values, []int{9, 4}) {
		t.Errorf("Expected array should be %s instead of %s", arrayToString([]int{9, 4}), arrayToString(leaf.rightSibling.values))
	}
}
