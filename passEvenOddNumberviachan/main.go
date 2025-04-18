package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan string)

	wg.Add(2)

	go func() {
		wg.Wait()
		close(c)
	}()

	go printEven(&wg, c)
	go printOdd(&wg, c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func printEven(wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			c <- fmt.Sprintf("%d is even", i)
		}
	}
}

func printOdd(wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			c <- fmt.Sprintf("%d is odd", i)
		}
	}
}
