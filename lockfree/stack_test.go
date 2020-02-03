package lockfree

import (
    "testing"
    "math/rand"
    "sync"
)

func TestNormalStack(t *testing.T) {
    s := Stack{}
    _, err := s.pop()
    if err == nil {
        t.Error("Pop from empty stack should raise exception")
    }
    s.push(1)
    s.push(3)
    s.push(5)
    s.push(2)
    b := [...]int {2, 5, 3, 1}
    for idx := 0; idx < 4; idx ++ {
        val, err := s.pop()
        if err != nil {
            t.Error("Shouldn't raise any exception, receive:", err)
        }
        if val != b[idx] {
            t.Errorf("Value at index %d should be %d instead of %d", idx, b[idx], val)
        }
    }
}

func TestLockFreeStack(t *testing.T) {
    s := newLockFreeStack()
    var wg sync.WaitGroup
    wg.Add(32)
    for idx := 0; idx < 32; idx ++ {
        go func(s *LockFreeStack) {
            defer wg.Done()
            for times := 0; times < 100; times ++ {
                s.lfPush(rand.Intn(100))
            }
        }(s)
    }
    wg.Wait()
    if s.size() != 3200 {
        t.Errorf("Wrong size %d", s.size())
    }
    for times := 0; times < 32; times ++ {
        curSize, expectedSize := s.size(), int(32 * 100 - times * 100)
        if curSize != expectedSize {
            t.Errorf("Step %d: Current stack should have %d elements. Actual current size is %d",
                     times, expectedSize, curSize)
        }
        wg.Add(10)
        for idx := 0; idx < 10; idx ++ {
            go func(s *LockFreeStack) {
                defer wg.Done()
                for noPop := 0; noPop < 10; noPop ++ {
                    _, err := s.lfPop()
                    if err != nil {
                        t.Error("Good pop shouldn't raise any exception")
                    }
                }
            }(s)
        }
        wg.Wait()
    }
}