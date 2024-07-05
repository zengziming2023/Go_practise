package main

import (
	"fmt"
	"time"
)

func worker2(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("worker init:", id)
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for wid := range 3 {
		go worker2(wid, jobs, results)
	}

	for j := range numJobs {
		jobs <- j
	}
	close(jobs)
	for range numJobs {
		fmt.Println(<-results)
	}
	fmt.Println("finished all jobs, and get all results")
}
