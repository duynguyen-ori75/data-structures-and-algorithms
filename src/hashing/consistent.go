package hashing

import (
	"fmt"
	"math/rand"
)

/**
 * All RPC structs definition
 */
type ReadRequest struct {
	key string
}

type ReadResponse struct {
	value int
	err   error
}

type WriteRequest struct {
	key   string
	value int
}

type WriteResponse struct {
	err error
}

type AddNodeRequest struct {
	name string
}

/**
 * @brief      Read request to any node
 *
 * @param      request  The read request
 *
 * @return     The read response from the cluster
 */
func (node Node) Read(request ReadRequest) ReadResponse {
	if val, ok := node.kvstore[request.key]; ok {
		return ReadResponse{value: val, err: nil}
	}
	expectedNodeIndex := node.FindExpectedNode(request.key)
	if node.name == node.config.nodeNames[expectedNodeIndex] {
		return ReadResponse{err: fmt.Errorf("Key %s does not exist", request.key)}
	}
	response := node.config.network.SendMessage(node.name, node.config.nodeNames[expectedNodeIndex], request)
	if readResp, ok := response.(ReadResponse); ok {
		return readResp
	}
	return ReadResponse{err: response.(error)}
}

/**
 * @brief      Write(Update) request to any node
 *
 * @param      request  The write request
 *
 * @return     The write response from the cluster
 */
func (node *Node) Write(request WriteRequest) WriteResponse {
	expectedNodeIndex := node.FindExpectedNode(request.key)
	if node.name == node.config.nodeNames[expectedNodeIndex] {
		node.kvstore[request.key] = request.value
		return WriteResponse{err: nil}
	}
	response := node.config.network.SendMessage(node.name, node.config.nodeNames[expectedNodeIndex], request)
	if writeResp, ok := response.(WriteResponse); ok {
		return writeResp
	}
	return WriteResponse{err: response.(error)}
}

/**
 * @brief      Adds a node to the cluster
 *
 * @param      request  The add node request - contain new node's name (aka its IP in real-world)
 *
 * @return     The add node response
 */
func (node *Node) AddNode(request AddNodeRequest) {
	expectedInsertIndex := node.FindExpectedNode(request.name)
	if request.name != node.config.nodeNames[expectedInsertIndex] {
		// the current node has not inserted new node yet -> insert and gossip to some random nodes
		node.config.nodeNames = insertString(node.config.nodeNames, expectedInsertIndex, request.name)
		// gossip protocol: pick random 5 nodes and forward the request to these nodes
		for time := 0; time < 5; time++ {
			forwardedIndex := expectedInsertIndex
			for forwardedIndex == expectedInsertIndex {
				forwardedIndex = rand.Intn(len(node.config.nodeNames))
			}
			forwardedNode := node.config.nodeNames[forwardedIndex]
			node.config.network.SendMessage(node.name, forwardedNode, request)
		}
	}
	// the node was already inserted in this node (visited before) => do nothing
}
