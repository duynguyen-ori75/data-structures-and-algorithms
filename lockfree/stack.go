package lockfree

import (
    "errors"
    "sync/atomic"
    "unsafe"
)

type Node struct {
    value int
    next *Node
}


// Normal stack functions
type Stack struct {
    head *Node
}


func (s *Stack) push(val int) {
    newNode := &Node{value: val, next: s.head}
    s.head = newNode
}

func (s *Stack) pop() (int, error) {
    if s.head == nil {
        return -1, errors.New("Stack is empty, can't pop")
    }
    popNode := s.head
    s.head = s.head.next
    return popNode.value, nil
}

func (s *Stack) size() int {
    count := 0
    for tmp := s.head; tmp != nil; tmp = tmp.next {
        count ++
    }
    return count
}


// Lock-free Stack functions
// this is a naive implementation which would fail ABA problem
type LockFreeStack struct {
    head unsafe.Pointer
    nilNode *Node
}

func newLockFreeStack() *LockFreeStack {
    result := &LockFreeStack{nilNode: &Node{value: -1}}
    result.head = unsafe.Pointer(result.nilNode)
    return result
}

// for testing only
func (s *LockFreeStack) size() int {
    count := 0
    for tmp := (*Node)(s.head); tmp != s.nilNode; tmp = tmp.next {
        count ++
    }
    return count
}

func (s *LockFreeStack) lfPush(val int) {
    for {
        currentHead := atomic.LoadPointer(&s.head)
        newNode := &Node{value: val, next: (*Node)(currentHead)}
        if atomic.CompareAndSwapPointer(&s.head, currentHead, unsafe.Pointer(newNode)) {
            break
        }
    }
}

func (s *LockFreeStack) lfPop() (int, error) {
    for {
        currentHead := atomic.LoadPointer(&s.head)
        if (*Node)(currentHead) == s.nilNode {
            return -1, errors.New("Stack is empty, can't pop")
        }
        if atomic.CompareAndSwapPointer(
            &s.head,
            currentHead,
            unsafe.Pointer((*Node)(currentHead).next)) {
            return (*Node)(currentHead).value, nil
        }
    }
}