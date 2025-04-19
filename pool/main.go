package main

import (
	"fmt"
	"time"
)

// Worker function processes tasks from the jobs channel.
func worker(id int, jobs <-chan int, res chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d is processing job %d\n", id, job)
		time.Sleep(time.Second) 
		res <- job * 2
	}
}

func main() {
	njobs := 5      
	nworkers := 3 
	jobs := make(chan int, njobs) // Task queue
	res := make(chan int, njobs)  // res queue

	// Start workers
	for w := 1; w <= nworkers; w++ {
		go worker(w, jobs, res)
	}

	// Send jobs to the jobs channel
	for j := 1; j <= njobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect res
	for r := 1; r <= njobs; r++ {
		fmt.Printf("Result: %d\n", <-res)
	}
}
