package hashing

import (
	"fmt"
	"reflect"
	"sort"
)

/**
 * Node information
 */
type NodeConfig struct {
	nodeNames []string
	network   *ClusterInfras
}

type Node struct {
	name    string // a representative of real-life IP address:port
	kvstore map[string]int
	config  NodeConfig
}

func NewNode(name string, config NodeConfig) *Node {
	return &Node{name: name, kvstore: make(map[string]int), config: config}
}

func (node Node) ConsistentHash(key string) int {
	nodeIndex := sort.Search(len(node.config.nodeNames), func(index int) bool {
		return hashSum(node.config.nodeNames[index]) >= hashSum(key)
	})
	return nodeIndex % len(node.config.nodeNames)
}

func (node Node) Receive(senderName string, message interface{}) interface{} {
	fmt.Printf("Receiving from %s - message %s\n", senderName, message)
	switch request := message.(type) {
	case ReadRequest:
		return node.Read(request)
	case WriteRequest:
		return node.Write(request)
	default:
		return fmt.Errorf("Can't decode the message - its type is '%s'", reflect.TypeOf(request).String())
	}
	return nil
}

/**
 * Infrastructure information - contains abstract functions to simulate network behavior
 */
type ClusterInfras struct {
	nodes []*Node
	// easy to simulate real-world system
	isNodeDeath []bool
}

func NewInfras(numberOfNodes int) (*ClusterInfras, error) {
	var infras ClusterInfras
	if numberOfNodes <= 0 {
		return nil, fmt.Errorf("Number of Nodes have to be bigger than 0")
	}
	// initialize share node config
	nodeConfig := NodeConfig{network: &infras}
	for index := 0; index < numberOfNodes; index++ {
		nodeConfig.nodeNames = append(nodeConfig.nodeNames, randomString())
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
	return &infras, nil
}

func (infras *ClusterInfras) SendMessage(senderName string, receiverName string, message interface{}) interface{} {
	if senderName == receiverName {
		return fmt.Errorf("Sender and receiver must be different")
	}
	senderIdx := infras.nodes[0].ConsistentHash(senderName)
	receiverIdx := infras.nodes[0].ConsistentHash(receiverName)
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
