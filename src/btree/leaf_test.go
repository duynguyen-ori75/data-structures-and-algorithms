package btree

import (
	//"log"
	"reflect"
	"testing"
)

func TestLeafNode_Search(t *testing.T) {
	leaf := newLeafNode([]int{1, 4, 7, 8}, []int{4, 5, 1, 5}, nil, nil, nil)
	val, err := leaf.Search(2)
	if err == nil {
		t.Error("Search for non-existant key should return error")
	}
	val, err = leaf.Search(4)
	if err != nil {
		t.Error("Search of existing key should be good")
	}
	if val != 5 {
		t.Errorf("Expected result should be 5 instead of %d", val)
	}
}

func TestLeafNode_Insert(t *testing.T) {
	leaf, degree := newLeafNode([]int{}, []int{}, nil, nil, nil), 3
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

func TestLeafNode_Delete(t *testing.T) {
	// initialize test
	parent, degree := newInternalNode([]int{5}, nil), 4
	sibling := newLeafNode([]int{5, 8, 10}, []int{9, 4, 5}, nil, nil, parent)
	leaf := newLeafNode([]int{1, 3, 4}, []int{5, 3, 12}, nil, sibling, parent)
	sibling.leftSibling = leaf

	// start testing
	err := leaf.Delete(6, degree)
	if err == nil {
		t.Error("Delete non-existant key should raise exception")
	}
	err = leaf.Delete(3, degree)
	if err != nil {
		t.Error("Delete existing key should be fine")
	}
	if !reflect.DeepEqual(leaf.keys, []int{1, 4}) || !reflect.DeepEqual(sibling.keys, []int{5, 8, 10}) {
		t.Errorf("All keys should be correct. Leaf's keys: %s - Sibling's keys: %s", arrayToString(leaf.keys), arrayToString(sibling.keys))
	}
	if !reflect.DeepEqual(leaf.values, []int{5, 12}) || !reflect.DeepEqual(sibling.values, []int{9, 4, 5}) {
		t.Errorf("All values should be correct")
	}
	err = leaf.Delete(1, degree)
	if err != nil {
		t.Error("Delete existing key should be fine")
	}
	if !reflect.DeepEqual(leaf.keys, []int{4, 5}) || !reflect.DeepEqual(sibling.keys, []int{8, 10}) {
		t.Errorf("Leaf should borrow key correctly. Leaf's keys: %s - Sibling's keys: %s", arrayToString(leaf.keys), arrayToString(sibling.keys))
	}
	if !reflect.DeepEqual(leaf.values, []int{12, 9}) || !reflect.DeepEqual(sibling.values, []int{4, 5}) {
		t.Errorf("All values should be relocated correctly. Leaf's values: %s - Sibling's values: %s",
			arrayToString(leaf.values), arrayToString(sibling.values))
	}

	err = leaf.Delete(5, degree)
	if err != nil {
		t.Error("Delete existing key should be fine")
	}
	if leaf.rightSibling != nil {
		t.Error("Two leaf nodes should be merged into one now")
	}
	if !reflect.DeepEqual(leaf.keys, []int{4, 8, 10}) || !reflect.DeepEqual(leaf.values, []int{12, 4, 5}) {
		t.Error("Last-standing leaf node should have correct keys and values")
	}
}
