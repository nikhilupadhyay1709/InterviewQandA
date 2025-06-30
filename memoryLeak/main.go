package main

import (
	"fmt"
	"math/rand"
	"time"
)

var leakedData [][]byte // this global slice causes memory leak

func main() {
	for {
		leakMemory()
		time.Sleep(100 * time.Millisecond)
	}
}

func leakMemory() {
	// Allocate 1MB slice
	data := make([]byte, 1024*1024)

	// Fill it with random data to simulate usage
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}

	// Fix: Clear old references
	if len(leakedData) > 1000 {
		leakedData = nil // or: leakedData = leakedData[:0]
	}

	// "Leak" it by appending to a global slice
	leakedData = append(leakedData, data)

	fmt.Printf("Leaked memory: %d MB\n", len(leakedData))
}
