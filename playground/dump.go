package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const modulo int64 = 12345678

func power(x int64, y int64) int64 {
	var result int64 = 1
	for times := int64(0); times < y; times++ {
		result = (result * x) % modulo
	}
	return result
	/*
		if y <= 1 {
			return int64(math.Pow(float64(x), float64(y)))
		}
		half := power(x, y/2) % modulo
		half = (half * half) % modulo
		if y%2 == 0 {
			return half
		}
		return (half * x) % modulo
	*/
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func normalAggregation(n int64) int64 {
	defer timeTrack(time.Now(), "Normal aggregation")
	var result int64 = 0
	for element := int64(1); element <= n; element++ {
		result = (result + power(element, element)) % modulo
	}
	return result
}

func goroutinesAggregation(n int64) int64 {
	defer timeTrack(time.Now(), "Goroutines aggregation")
	var result int64 = 0
	var wg sync.WaitGroup
	wg.Add(int(n))
	for element := int64(1); element <= n; element++ {
		go func(element int64) {
			for {
				oldResult := result
				expectedResult := (oldResult + power(element, element)) % int64(modulo)
				if atomic.CompareAndSwapInt64(&result, oldResult, expectedResult) {
					break
				}
			}
			wg.Done()
		}(element)
	}
	wg.Wait()
	return result
}

func workerAggregation(n int64, allAlements chan int64, numberOfWorkers int) int64 {
	defer timeTrack(time.Now(), "Worker aggregation")
	summary := make(chan int64)
	// 4 goroutines receiving int and doing aggregation
	for numRoutine := 0; numRoutine < numberOfWorkers; numRoutine++ {
		go func() {
			var tmpResult int64 = 0
			for nextElement := range allAlements {
				tmpResult = (tmpResult + power(nextElement, nextElement)) % modulo
			}
			summary <- tmpResult
		}()
	}
	// aggregate all values
	var result int64 = 0
	for idx := 0; idx < numberOfWorkers; idx++ {
		result = (result + <-summary) % modulo
	}
	return result
}

func main() {
	var n int64 = 10000
	fmt.Println(normalAggregation(n))
	fmt.Println(goroutinesAggregation(n))

	// worker method
	numberOfWorkers := 16
	allAlements := make(chan int64, n)
	// int generator
	go func(n int64) {
		for element := int64(1); element <= n; element++ {
			allAlements <- element
		}
		close(allAlements)
	}(n)
	fmt.Println(workerAggregation(n, allAlements, numberOfWorkers))
}
