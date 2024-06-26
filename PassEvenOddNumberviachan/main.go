package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(2)

	go func() {
		wg.Wait()
		close(ch)
	}()

	go printEven(&wg, ch)
	go printOdd(&wg, ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}

func printEven(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			ch <- fmt.Sprintf("%d is even", i)
		}
	}
}

func printOdd(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			ch <- fmt.Sprintf("%d is odd", i)
		}
	}
}
