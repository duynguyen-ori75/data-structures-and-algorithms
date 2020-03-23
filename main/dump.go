package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func formula(value int) int {
	//return value
	return value ^ value%12345678
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func normalAggregation(n int) int {
	defer timeTrack(time.Now(), "Normal aggregation")
	result := 0
	for element := 0; element <= n; element++ {
		result += formula(element)
	}
	return result
}

func goroutinesAggregation(n int) int {
	defer timeTrack(time.Now(), "Goroutines aggregation")
	var result int64 = 0
	var wg sync.WaitGroup
	wg.Add(n)
	for element := 1; element <= n; element++ {
		go func(element int) {
			atomic.AddInt64(&result, int64(formula(element)))
			wg.Done()
		}(element)
	}
	wg.Wait()
	return int(result)
}

func generatorAggregation(n int) int {
	defer timeTrack(time.Now(), "Generator aggregation")
	numberOfWorkers := 4
	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)
	allAlements, summary := make(chan int), make(chan int, numberOfWorkers)
	// int generator
	go func(n int) {
		for element := 1; element <= n; element++ {
			allAlements <- element
		}
		for idx := 0; idx < numberOfWorkers; idx++ {
			allAlements <- -1
		}
	}(n)
	// 4 goroutines receiving int and doing aggregation
	for numRoutine := 0; numRoutine < numberOfWorkers; numRoutine++ {
		go func() {
			tmpResult := 0
			for {
				nextElement := <-allAlements
				if nextElement < 0 {
					break
				}
				tmpResult += formula(nextElement)
			}
			summary <- tmpResult
			wg.Done()
		}()
	}
	wg.Wait()
	// aggregate all values
	result := 0
	for numRoutine := 0; numRoutine < numberOfWorkers; numRoutine++ {
		result += <-summary
	}
	return result
}

func main() {
	fmt.Println(normalAggregation(1000000))
	fmt.Println(goroutinesAggregation(1000000))
	fmt.Println(generatorAggregation(1000000))
}
