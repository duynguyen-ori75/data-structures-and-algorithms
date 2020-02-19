package lsm

import (
	"os"
	"testing"
)

func TestSSTable(t *testing.T) {
	// clean database folder
	os.RemoveAll("./db")
	// start testing
	table, err := NewTable(16, 16)
	if err != nil {
		t.Errorf("Pop from empty stack should raise exception. Exception: %s", err.Error())
	}
	table.Update("hello", "world")
	table.Flush()
}
