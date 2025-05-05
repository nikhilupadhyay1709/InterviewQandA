package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	num := 1
	wg.Add(2)
	go worker(&wg, &mu, &num)
	go worker(&wg, &mu, &num)
	wg.Wait()
}

func worker(wg *sync.WaitGroup, mu *sync.Mutex, num *int) {
	defer wg.Done()
	for {
		mu.Lock()
		if *num >= 1000 {
			mu.Unlock()
			break
		}
		fmt.Print(" ", *num)
		*num++
		mu.Unlock()
	}
}
