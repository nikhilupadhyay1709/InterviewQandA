package main

import (
	"fmt"
	"sync"
)

func createArray(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := 1
	c <- arr
	close(c)
}

func printArray(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for arr := range c {
		fmt.Println("Received array:", arr)
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(2)

	go createArray(c, &wg)
	go printArray(c, &wg)

	wg.Wait()
}
