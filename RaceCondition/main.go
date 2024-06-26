package main

import (
	"fmt"
	"sync"
)

// race condition
func main() {

	wg := &sync.WaitGroup{}

	var score = []int{0}

	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		fmt.Println("One R")
		score = append(score, 1)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Two R")
		score = append(score, 2)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Three R")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()

	fmt.Println("Score:", score)
}

// exit status 66 : race condition was...

func NoRaceCondition() {

	wg := &sync.WaitGroup{}

	mut := &sync.Mutex{}
	var score = []int{0}

	wg.Add(1)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()	
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
		
	wg.Wait()

	fmt.Println("Score:", score)
}
