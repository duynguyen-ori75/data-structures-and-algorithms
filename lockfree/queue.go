package lockfree

import (
	//"fmt"
	"errors"
	"sync"
	"sync/atomic"
	"unsafe"
)

/**
 * Implementation of a simple MPMC queue using a single mutex for both head and tail
 */
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


/**
 * Implementation of a MPMC queue using two mutexes, one for head pointer and the other for the tail
 */
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

/**
 * Implementation of a MPMC lock-free queue using 
 */
type LockFreeQueue struct {
	head    unsafe.Pointer
	tail    unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
	nullNode := unsafe.Pointer(&Node{value: 0, next: nil})
	return &LockFreeQueue{head: nullNode, tail: nullNode}
}

func (q *LockFreeQueue) Push(val int) {
	newNode := &Node{value: val}
	for {
		currentTail := atomic.LoadPointer(&q.tail)
		desiredModifiedOldTail := (*Node)(currentTail)
		if desiredModifiedOldTail.next == nil {
			desiredModifiedOldTail.next = newNode
			if atomic.CompareAndSwapPointer(&q.tail, currentTail, unsafe.Pointer(desiredModifiedOldTail)) {
				atomic.CompareAndSwapPointer(&q.tail, unsafe.Pointer(desiredModifiedOldTail), unsafe.Pointer(newNode))
				break
			}
		}
	}
}

func (q* LockFreeQueue) Pop() (int, error) {
	for {
		currentHead := atomic.LoadPointer(&q.head)
		if currentHead == atomic.LoadPointer(&q.tail) {
			return -1, errors.New("Stack is empty, can't pop")
		}
		expectedNewHead := ((*Node)(currentHead)).next
		if atomic.CompareAndSwapPointer(&q.head, currentHead, unsafe.Pointer(expectedNewHead)) {
			return ((*Node)(currentHead)).value, nil
		}
	}
}