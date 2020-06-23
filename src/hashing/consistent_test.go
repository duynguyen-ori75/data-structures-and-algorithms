package hashing

import (
	//"log"
	"math/rand"
	"sort"
	"strings"
	"testing"
)

func TestUtilities(t *testing.T) {
	// random string generator
	testStr := randomString()
	if len(testStr) != nameLength {
		t.Errorf("Length of random generated string should be %d", nameLength)
	}
	for _, char := range testStr {
		if !strings.ContainsRune(testStr, char) {
			t.Errorf("Senerated string (%s) contains unexpected char, which is %c", testStr, char)
		}
	}
	// hash sum - check the determinism of our hash function
	if hashSum("aaaa") != hashSum("aaaa") || hashSum("9xvef84723xas") != hashSum("9xvef84723xas") {
		t.Error("Hash function should be deterministic")
	}
}

func TestConsistentHash(t *testing.T) {
	infras, err := NewInfras(100)
	if err != nil {
		t.Errorf("Should not raise any exception. Meet: %s", err)
	}
	if len(infras.nodes) != 100 || len(infras.isNodeDeath) != 100 {
		t.Errorf("Number of nodes should be 100")
	}
	// check if nodes are sorted asc using hash value of node names
	less := func(i, j int) bool {
		return hashSum(infras.nodes[i].name) <= hashSum(infras.nodes[j].name)
	}
	if !sort.SliceIsSorted(infras.nodes, less) {
		t.Error("All nodes should be sorted in order of their names")
	}
	// check the determinism and correctness of Consistent Hashing function
	firstNode := infras.nodes[0]
	testKeys := []string{"abcxyz", "12jahjsad", "ahjjsad124", "dumpKey", "w98e87as7cc", infras.nodes[10].name, firstNode.name}
	for _, testKey := range testKeys {
		expectedNodeIdx := firstNode.ConsistentHash(testKey)
		if expectedNodeIdx != firstNode.ConsistentHash(testKey) {
			t.Error("Consistent hashing function should be deterministic")
		}
		// find correct node in the ring to store the key-value
		prevNodeIdx := (expectedNodeIdx + 99) % 100
		flag := (hashSum(infras.nodes[prevNodeIdx].name) < hashSum(testKey)) ||
			(hashSum(infras.nodes[expectedNodeIdx].name) >= hashSum(testKey))
		if !flag {
			t.Errorf("Key %s (hashSum is %d) is found at index %d, before node whose hash sum is %d",
				testKey, hashSum(testKey), expectedNodeIdx, hashSum(infras.nodes[expectedNodeIdx].name))
		}
		// test write and read
		request := WriteRequest{key: testKey, value: rand.Int()}
		writeResp := firstNode.Write(request)
		if writeResp.err != nil {
			t.Errorf("Should not raise any exception. Meet: %s", writeResp.err)
		}
		if _, ok := infras.nodes[expectedNodeIdx].kvstore[testKey]; !ok {
			t.Errorf("%s should be insert to node %d", testKey, expectedNodeIdx)
		}
		response := firstNode.Read(ReadRequest{key: testKey})
		if response.err != nil {
			t.Errorf("Read existing key %s should not raise exception", testKey)
		}
		if response.value != request.value {
			t.Errorf("Returned value is wrong, expected %d - found %d", request.value, response.value)
		}
	}
	response := firstNode.Read(ReadRequest{key: "not-exist"})
	if response.err == nil {
		t.Error("Read non-existing key should raise exception")
	}
}
