package main

import (
	"fmt"
	"os"
	"sync"
)

var m = map[string]int{}

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("error while reading dir")
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	numRoutines := 3
	fileCh := make(chan string)

	// Start 3 goroutines
	for range numRoutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for name := range fileCh {
				fileData, err := os.ReadFile("." + name)
				if err != nil {
					fmt.Println("error while reading the file:", name)
					continue
				}
				for _, v := range fileData {
					mu.Lock()
					m[string(v)] = m[string(v)] + 1
					mu.Unlock()
				}
			}
		}()
	}

	// Send file names to goroutines
	for _, v := range files {
		fileCh <- v.Name()
	}
	close(fileCh)

	wg.Wait()

	fmt.Println("Character counts:", m)
}
