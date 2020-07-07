package lockfree

import (
	//"fmt"
	"errors"
	"sync"
)

type SingleLockQueue struct {
	mux  sync.Mutex
	head *Node
	tail *Node
}

func NewSingleLockQueue() *SingleLockQueue {
	nullNode := &Node{value: 0, next: nil}
	return &SingleLockQueue{head: nullNode, tail: nullNode}
}

func (q *SingleLockQueue) Push(val int) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.tail.next = &Node{value: val}
	q.tail = q.tail.next
}

func (q *SingleLockQueue) Pop() (int, error) {
	q.mux.Lock()
	defer q.mux.Unlock()
	if q.head.next == nil {
		return -1, errors.New("The queue is empty")
	}
	headVal := q.head.next.value
	q.head = q.head.next
	return headVal, nil
}

type TwoLockQueue struct {
	headMux sync.Mutex
	tailMux sync.Mutex
	head    *Node
	tail    *Node
}

func NewTwoLockQueue() *TwoLockQueue {
	nullNode := &Node{value: 0, next: nil}
	return &TwoLockQueue{head: nullNode, tail: nullNode}
}

func (q *TwoLockQueue) Push(val int) {
	q.tailMux.Lock()
	defer q.tailMux.Unlock()
	q.tail.next = &Node{value: val}
	q.tail = q.tail.next
}

func (q *TwoLockQueue) Pop() (int, error) {
	q.headMux.Lock()
	defer q.headMux.Unlock()
	if q.head.next == nil {
		return -1, errors.New("The queue is empty")
	}
	headVal := q.head.next.value
	q.head = q.head.next
	return headVal, nil
}

// for testing only
func (q TwoLockQueue) size() int {
	count := -1
	for tmp := (*Node)(q.head); tmp != nil; tmp = tmp.next {
		count++
	}
	return count
}
