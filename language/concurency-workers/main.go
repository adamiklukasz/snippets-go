package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("start i=%#v\n", id)
		time.Sleep(1 * time.Second)
		fmt.Printf("stop i=%#v\n", id)
		results <- j * 2
	}
}

func main() {
	const (
		numJobs    = 5
		numWorkers = 3
	)

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a < numJobs; a++ {
		j := <-results
		fmt.Printf("j=%#v\n", j)
	}

}
