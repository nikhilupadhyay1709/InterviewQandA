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

	go func() {
		defer wg.Done()
		for num < 10 {
			mu.Lock()
			if num%2 != 0 {
				fmt.Print(num, ",")
				num++
			}
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for num < 10 {
			mu.Lock()
			if num%2 == 0 {
				fmt.Print(num, ",")
				num++
			}
			mu.Unlock()
		}
	}()

	wg.Wait()
}
