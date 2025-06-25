package main

import (
	"fmt"
	"math"
	"sync"
)

func createNumbers(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func calculateSqrt(in <-chan int, out chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		out <- math.Sqrt(float64(num))
	}
	close(out)
}

func printResults(in <-chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	for sqrt := range in {
		fmt.Printf("Square root of %d is %.2f\n", i, sqrt)
		i++
	}
}

func main() {
	nchan := make(chan int)
	sqChan := make(chan float64)
	var wg sync.WaitGroup

	wg.Add(3)

	go createNumbers(nchan, &wg)
	go calculateSqrt(nchan, sqChan, &wg)
	go printResults(sqChan, &wg)

	wg.Wait()
}
