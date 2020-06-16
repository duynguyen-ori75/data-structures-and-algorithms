package btree

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	leaf := newLeafNode([]int{1, 4, 4, 4, 4, 7, 8}, []int{4, 5, 7, 2, 3, 1, 5}, nil)
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

}
