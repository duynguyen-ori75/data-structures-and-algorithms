package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

const MODULO int64 = 12345678

func power(x int64, y int64) int64 {
	if y <= 1 {
		return int64(math.Pow(float64(x), float64(y)))
	}
	half := power(x, y / 2) % MODULO
	half = (half * half) % MODULO
	if y % 2 == 0 {
		return half
	} else {
		return (half * x) % MODULO
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func normalAggregation(n int64) int64 {
	defer timeTrack(time.Now(), "Normal aggregation")
	var result int64 = 0
	for element := int64(1); element <= n; element++ {
		result = (result + power(element, element)) % MODULO
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
				expectedResult := (oldResult + power(element, element)) % int64(MODULO)
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

func generatorAggregation(n int64) int64 {
	defer timeTrack(time.Now(), "Generator aggregation")
	numberOfWorkers := 4
	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)
	allAlements, summary := make(chan int64), make(chan int64, numberOfWorkers)
	// int generator
	go func(n int64) {
		for element := int64(1); element <= n; element++ {
			allAlements <- element
		}
		close(allAlements)
	}(n)
	// 4 goroutines receiving int and doing aggregation
	for numRoutine := 0; numRoutine < numberOfWorkers; numRoutine++ {
		go func() {
			var tmpResult int64 = 0
			for nextElement := range allAlements {
				tmpResult = (tmpResult + power(nextElement, nextElement)) % MODULO
			}
			summary <- tmpResult
			wg.Done()
		}()
	}
	wg.Wait()
	// aggregate all values
	var result int64 = 0
	for numRoutine := 0; numRoutine < numberOfWorkers; numRoutine++ {
		result = (result + <-summary) % MODULO
	}
	return result
}

func main() {
	fmt.Println(normalAggregation(10000))
	fmt.Println(goroutinesAggregation(10000))
	fmt.Println(generatorAggregation(10000))
}
