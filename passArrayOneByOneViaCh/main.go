package main

import (
	"fmt"
	"sync"
)

func createArray(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := []int{10, 20, 30, 40, 50}
	for _, v := range arr {
		c <- v
	}
	close(c)
}

func printArray(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	res := make([]int, 0)
	for v := range c {
		res = append(res, v)
	}
	fmt.Println("Array:", res)
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(2)

	go createArray(c, &wg)
	go printArray(c, &wg)

	wg.Wait()
}
