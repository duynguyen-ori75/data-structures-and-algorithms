package lockfree

import (
	//"fmt"
	"errors"
	"sync"
	"sync/atomic"
	"unsafe"
)

type Queue interface {
	Push(int)
	Pop() (int, error)
	// size() is for testing only
	size() int
}

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

// for testing only
func (q SingleLockQueue) size() int {
	count := -1
	for tmp := q.head; tmp != nil; tmp = tmp.next {
		count++
	}
	return count
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

func (q TwoLockQueue) size() int {
	count := -1
	for tmp := q.head; tmp != nil; tmp = tmp.next {
		count++
	}
	return count
}

/**
 * Implementation of a MPMC lock-free queue using 
 */
type LFNode struct {
	value int
	next  unsafe.Pointer
}
type LockFreeQueue struct {
	head    unsafe.Pointer
	tail    unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
	nullNode := unsafe.Pointer(&LFNode{value: 0, next: nil})
	return &LockFreeQueue{head: nullNode, tail: nullNode}
}

func (q *LockFreeQueue) Push(val int) {
	newNode := &LFNode{value: val}
	for {
		currentTail := q.tail
		nextNode := ((*LFNode)(currentTail)).next
		if q.tail == currentTail && nextNode == nil {
			if atomic.CompareAndSwapPointer(&((*LFNode)(currentTail).next), nil, unsafe.Pointer(newNode)) {
				break
			}
		} else {
			atomic.CompareAndSwapPointer(&q.tail, currentTail, nextNode)
		}
	}
}

func (q* LockFreeQueue) Pop() (int, error) {
	for {
		currentHead, currentTail := q.head, q.tail
		if currentHead == currentTail {
			// tail is falling behind -> advance tail node
			if ((*LFNode)(currentTail)).next != nil {
				atomic.CompareAndSwapPointer(&q.tail, currentTail, ((*LFNode)(currentTail)).next)
				continue
			}
			return -1, errors.New("Stack is empty, can't pop")
		}
		expectedNewHead := ((*LFNode)(currentHead)).next
		if expectedNewHead != nil {
			if atomic.CompareAndSwapPointer(&q.head, currentHead, unsafe.Pointer(expectedNewHead)) {
				return ((*LFNode)(expectedNewHead)).value, nil
			}
		}
	}
}

func (q LockFreeQueue) size() int {
	count := -1
	for tmp := q.head; tmp != nil; tmp = ((*LFNode)(tmp)).next {
		count++
	}
	return count
}