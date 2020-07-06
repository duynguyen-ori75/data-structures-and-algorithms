package lsm

import (
	"os"
	"testing"
)

func TestSSTable(t *testing.T) {
	// clean database folder
	os.RemoveAll("./db")
	// start testing
	table, err := NewTable(8, 8)
	if err != nil {
		t.Errorf("New table shouldn't raise exception. Exception: %s", err.Error())
	}
	err = table.Update("hello", "world")
	if err != nil {
		t.Errorf("Insert shouldn't raise exception. Exception: %s", err.Error())
	}
	err = table.Update("abcxyz", "123456")
	if err != nil {
		t.Errorf("Insert shouldn't raise exception. Exception: %s", err.Error())
	}
	table.Flush()
	testTable, err := NewTableFromID(table.id.String())
	if err != nil {
		t.Errorf("Create new sstable from uuid shouldn't raise exception. Exception: %s", err.Error())
	}
	if testTable.keySize != table.keySize {
		t.Errorf("Expected keySize is %d, found %d", table.keySize, testTable.keySize)
	}
	if testTable.valueSize != table.valueSize {
		t.Errorf("Expected valueSize is %d, found %d", table.valueSize, testTable.valueSize)
	}
	if len(testTable.data) != 2 {
		t.Errorf("Expected data should be a map of 1 tuple, found %d tuple(s)", len(testTable.data))
	}
	val, err := testTable.Read("hello")
	if err != nil {
		t.Errorf("Read should work instead of raising exception: %s", err.Error())
	}
	if val != "world" {
		t.Errorf("Expected value is 'world', found %s", val)
	}
	val, err = testTable.Read("abcxyz")
	if err != nil {
		t.Errorf("Read should work instead of raising exception: %s", err.Error())
	}
	if val != "123456" {
		t.Errorf("Expected value is '123456', found %s", val)
	}
}
