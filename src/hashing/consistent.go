package hashing

import "fmt"

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
	expectedNodeIndex := node.ConsistentHash(request.key)
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
	expectedNodeIndex := node.ConsistentHash(request.key)
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
