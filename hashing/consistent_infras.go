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
	config  *NodeConfig
}

func NewNode(name string, network *ClusterInfras, nodeNames []string) *Node {
	return &Node{name: name, kvstore: make(map[string]int),
		config: &NodeConfig{network: network, nodeNames: append([]string{}, nodeNames...)}}
}

func (node Node) FindExpectedNode(key string) int {
	nodeIndex := sort.Search(len(node.config.nodeNames), func(index int) bool {
		return hashSum(node.config.nodeNames[index]) >= hashSum(key)
	})
	return nodeIndex % len(node.config.nodeNames)
}

func (node Node) Receive(senderName string, message interface{}) interface{} {
	switch request := message.(type) {
	case ReadRequest:
		return node.Read(request)
	case WriteRequest:
		return node.Write(request)
	case AddNodeRequest:
		node.AddNode(request)
	default:
		return fmt.Errorf("Can't decode the message - its type is '%s'", reflect.TypeOf(request).String())
	}
	return nil
}

/**
 * Infrastructure information - contains abstract functions to simulate network behavior
 * Its main responsibility is to simulate real world network system, therefore it should be used in tests only
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
	var nodeNames []string
	for index := 0; index < numberOfNodes; index++ {
		nodeNames = append(nodeNames, randomString())
	}
	sort.Slice(nodeNames, func(i, j int) bool {
		return hashSum(nodeNames[i]) <= hashSum(nodeNames[j])
	})
	// initialize nodes in the cluster - and sort them in asc order (easy to test)
	for _, nodeName := range nodeNames {
		infras.nodes = append(infras.nodes, NewNode(nodeName, &infras, nodeNames))
	}
	// initialize state of nodes
	infras.isNodeDeath = make([]bool, numberOfNodes)
	for index := 0; index < numberOfNodes; index++ {
		infras.isNodeDeath[index] = false
	}
	return &infras, nil
}

func (infras *ClusterInfras) FindExpectedNode(key string) int {
	nodeIndex := sort.Search(len(infras.nodes), func(index int) bool {
		return hashSum(infras.nodes[index].name) >= hashSum(key)
	})
	return nodeIndex % len(infras.nodes)
}

func (infras *ClusterInfras) SendMessage(senderName string, receiverName string, message interface{}) interface{} {
	if senderName == receiverName {
		return fmt.Errorf("Sender and receiver must be different")
	}
	senderIdx := infras.FindExpectedNode(senderName)
	receiverIdx := infras.FindExpectedNode(receiverName)
	if senderIdx >= len(infras.nodes) || infras.nodes[senderIdx].name != senderName {
		return fmt.Errorf("Can't find sender node: %s", senderName)
	}
	if receiverIdx >= len(infras.nodes) || infras.nodes[receiverIdx].name != receiverName {
		return fmt.Errorf("Can't find receiver node: %s", receiverName)
	}
	// check state of nodes
	if infras.isNodeDeath[senderIdx] || infras.isNodeDeath[receiverIdx] {
		return fmt.Errorf("One of two nodes is death (%s -> %s)", senderName, receiverName)
	}
	// send message
	return infras.nodes[receiverIdx].Receive(senderName, message)
}

/**
 * @brief      Insert new node into current network cluster (should be run after gossiping all nodes)
 *
 * @param      newNodeName  The new node name
 *
 * @return     Golang error || inserted index (Int)
 */
func (infras *ClusterInfras) AddNewNode(newNodeName string) interface{} {
	expectedIndex := infras.nodes[0].FindExpectedNode(newNodeName)
	if infras.nodes[expectedIndex].name == newNodeName {
		return fmt.Errorf("Node %s already exists", newNodeName)
	}
	infras.isNodeDeath = insertBool(infras.isNodeDeath, expectedIndex, false)
	infras.nodes = insertNode(infras.nodes, expectedIndex,
		NewNode(newNodeName, infras, infras.nodes[0].config.nodeNames))
	return expectedIndex
}
