package main

import (
	"fmt"
	"time"
)

// Worker function processes tasks from the jobs channel.
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d is processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2      // Send result back
	}
}

func main() {
	numJobs := 5      // Number of tasks
	numWorkers := 3   // Number of workers

	jobs := make(chan int, numJobs)    // Task queue
	results := make(chan int, numJobs) // Results queue

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channel when done sending tasks

	// Collect results
	for r := 1; r <= numJobs; r++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}
