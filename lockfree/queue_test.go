package lockfree

import (
	//"fmt"
	"math/rand"
	"sync"
	"testing"
)

const benchNumThreads = 8

func TestTwoLockQueue(t *testing.T) {
	q := NewTwoLockQueue()
	_, err := q.Pop()
	if err == nil {
		t.Error("Pop from empty queue should raise exception")
	}
	q.Push(1)
	q.Push(3)
	q.Push(5)
	q.Push(2)
	b := [...]int{1, 3, 5, 2}
	for idx := 0; idx < 4; idx++ {
		val, err := q.Pop()
		if err != nil {
			t.Error("Shouldn't raise any exception, receive:", err)
		}
		if val != b[idx] {
			t.Errorf("Value at index %d should be %d instead of %d", idx, b[idx], val)
		}
	}
}

func TestTwoLockQueue_Concurrency(t *testing.T) {
	var wg sync.WaitGroup
	q := NewTwoLockQueue()
	wg.Add(benchNumThreads)
	for thread := 0; thread < benchNumThreads; thread++ {
		go func() {
			for val := 0; val < numberOfPushes/benchNumThreads; val++ {
				q.Push(val)
			}
			for time := 0; time < numberOfPops/benchNumThreads; time++ {
				_, err := q.Pop()
				if err != nil {
					t.Errorf("Should not raise exception here. Meet: %s", err)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	queueSize := q.size()
	if queueSize != numberOfPushes-numberOfPops {
		t.Errorf("Expected size is %d, found %d elements", numberOfPushes-numberOfPops, queueSize)
	}
}

func BenchmarkSingleLockQueue_8_threads(t *testing.B) {
	// intialize
	var wg sync.WaitGroup
	actions := []int{}
	for val := 0; val < numberOfPushes; val++ {
		actions = append(actions, rand.Intn(100))
	}
	t.ResetTimer()
	for times := 0; times < t.N; times++ {
		q := NewSingleLockQueue()
		wg.Add(benchNumThreads)
		for thread := 0; thread < benchNumThreads; thread++ {
			go func() {
				// 70% push - 30% pop
				for val := 0; val < numberOfPushes; val++ {
					if actions[val] < 70 {
						q.Push(val)
					} else {
						q.Pop()
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkTwoLockQueue_8_threads(t *testing.B) {
	// intialize
	var wg sync.WaitGroup
	actions := []int{}
	for val := 0; val < numberOfPushes; val++ {
		actions = append(actions, rand.Intn(100))
	}
	t.ResetTimer()
	for times := 0; times < t.N; times++ {
		q := NewTwoLockQueue()
		wg.Add(benchNumThreads)
		for thread := 0; thread < benchNumThreads; thread++ {
			go func() {
				// 70% push - 30% pop
				for val := 0; val < numberOfPushes; val++ {
					if actions[val] < 70 {
						q.Push(val)
					} else {
						q.Pop()
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
