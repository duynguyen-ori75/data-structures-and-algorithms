package lockfree

import (
	//"fmt"
	"math/rand"
	"sync"
	"testing"
)

const benchNumThreads = 8

func QueueTestFactory(t *testing.T, q Queue) {
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

func QueueConcurrencyTestFactory(t *testing.T, q Queue) {
	var wg sync.WaitGroup
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

func actionsGenerator() []int {
	result := []int{}
	for val := 0; val < numberOfPushes; val++ {
		result = append(result, rand.Intn(100))
	}
	return result
}

func QueueBenchmarkFactory(t *testing.B, q Queue, actions []int) {
	// intialize
	var wg sync.WaitGroup
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

func TestSingleLockQueue(t *testing.T) { QueueTestFactory(t, NewSingleLockQueue()) }
func TestTwoLockQueue(t *testing.T) { QueueTestFactory(t, NewTwoLockQueue()) }
func TestLockFreeQueue(t *testing.T) { QueueTestFactory(t, NewLockFreeQueue()) }

func TestSingleLockQueue_Concurrency(t *testing.T) { QueueConcurrencyTestFactory(t, NewSingleLockQueue()) }
func TestTwoLockQueue_Concurrency(t *testing.T) { QueueConcurrencyTestFactory(t, NewTwoLockQueue()) }
func TestLockFreeQueue_Concurrency(t *testing.T) { QueueConcurrencyTestFactory(t, NewLockFreeQueue()) }

func BenchmarkSingleLockQueue(t *testing.B) {
	actions := actionsGenerator()
	t.ResetTimer()
	for times := 0; times < t.N; times++ {
		QueueBenchmarkFactory(t, NewSingleLockQueue(), actions)
	}
}

func BenchmarkTwoLockQueue(t *testing.B) {
	actions := actionsGenerator()
	t.ResetTimer()
	for times := 0; times < t.N; times++ {
		QueueBenchmarkFactory(t, NewTwoLockQueue(), actions)
	}
}

func BenchmarkLockFreeQueue(t *testing.B) {
	actions := actionsGenerator()
	t.ResetTimer()
	for times := 0; times < t.N; times++ {
		QueueBenchmarkFactory(t, NewLockFreeQueue(), actions)
	}
}
