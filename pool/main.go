package main

import (
	"fmt"
	"time"
)

// Worker function processes tasks from the jobs channel.
func worker(id int, jobs <-chan int, res chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d is processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		res <- job * 2          // Send result back
	}
}

func main() {
	njobs := 5      // Number of tasks
	numWorkers := 3 // Number of workers

	jobs := make(chan int, njobs) // Task queue
	res := make(chan int, njobs)  // res queue

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, res)
	}

	// Send jobs to the jobs channel
	for j := 1; j <= njobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channel when done sending tasks

	// Collect res
	for r := 1; r <= njobs; r++ {
		fmt.Printf("Result: %d\n", <-res)
	}
}
