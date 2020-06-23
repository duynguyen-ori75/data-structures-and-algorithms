package hashing

import (
	//"log"
	"strings"
	"testing"
)

func TestUtilities(t *testing.T) {
	testStr := randomString()
	if len(testStr) != nameLength {
		t.Errorf("Length of random generated string should be %d", nameLength)
	}
	for _, char := range testStr {
		if !strings.ContainsRune(testStr, char) {
			t.Errorf("Senerated string (%s) contains unexpected char, which is %c", testStr, char)
		}
	}
}