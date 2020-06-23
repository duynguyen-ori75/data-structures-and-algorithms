package hashing

import (
	"sort"
)

func (node Node) Read(r ReadRequest) ReadResponse {
	if val, ok := node.kvstore[r.key]; ok {
		return ReadResponse{value: val, err: nil}
	}
	expectedNodeIndex := sort.Search(len(node.config.nodeNames), func(index int) bool {
		return hashSum(node.config.nodeNames[index]) >= hashSum(r.key)
	})
	expectedNodeIndex %= len(node.config.nodeNames)
	return node.Send(node.config.nodeNames[expectedNodeIndex], r).(ReadResponse)
}
