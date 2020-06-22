package hashing

import (
	"fmt"
	"hash"
	"hash/crc32"
	"rand"
	"sort"
)

/**
 * Random-string generator
 */
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const nameLength = 10

func randomString(length int) {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

/**
 * Node information
 */
type NodeConfig struct {
	nodeNames []string
	hashFunc  hash.Hash32
}

type Node struct {
	name    string // a representative of real-life IP address:port
	kvstore map[int]int
	config  NodeConfig
}

func NewNode(name string, config NodeConfig) *Node {
	return &Node{name: name, kvstore: make(map[int]int), config: config}
}

func (node Node) Receive(senderName string, message string) error {
	fmt.Printf("Receiving from %s - message %s\n", senderName, message)
	return nil
}

/**
 * Infrastructure information
 */
type ClusterInfras struct {
	nodes []*Node
	// easy to simulate real-world system
	isNodeDeath []bool
}

func NewInfras(numberOfNodes int) ClusterInfras {
	var infras ClusterInfras
	// initialize share node config
	nodeConfig := NodeConfig{hashFunc: crc32.NewIEEE()}
	for index := 0; index < numberOfNodes; index++ {
		nodeConfig.nodeNames = append(nodeConfig.nodeNames, randomString(nameLength))
	}
	sort.Strings(nodeConfig.nodeNames)
	// initialize nodes and the cluster
	for _, nodeName := range nodeConfig.nodeNames {
		infras.nodes = append(infras.nodes, NewNode(nodeName, nodeConfig))
	}
	// initialize state of nodes
	infras.isNodeDeath = make([]bool, numberOfNodes)
	for index := 0; index < numberOfNodes; index++ {
		infras.isNodeDeath[index] = false
	}
	return infras
}

func (infras *ClusterInfras) SendMessage(senderName string, receiverName string, message string) error {
	senderIdx := sort.Search(len(infras.nodes), func(index int) bool {
		return infras.nodes[index].name >= senderName
	})
	receiverIdx := sort.Search(len(infras.nodes), func(index int) bool {
		return infras.nodes[index].name >= receiverName
	})
	if senderIdx >= len(infras.nodes) || infras.nodes[senderIdx].name != senderName {
		return fmt.Errorf("Can't find node: %s", senderName)
	}
	if receiverIdx >= len(infras.nodes) || infras.nodes[receiverIdx].name != receiverName {
		return fmt.Errorf("Can't find node: %s", receiverName)
	}
	// check state of nodes
	if infras.isNodeDeath[senderIdx] || infras.isNodeDeath[receiverIdx] {
		return fmt.Errorf("One of two nodes is death (%s -> %s)", senderName, receiverName)
	}
	// send message
	return infras.nodes[receiverIdx].Receive(senderName, message)
}
