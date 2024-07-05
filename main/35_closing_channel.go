package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job: ", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := range 3 {
		jobs <- j
		fmt.Println("sent job", j)
		time.Sleep(time.Millisecond)
	}
	close(jobs)

	isDone := <-done
	fmt.Println(isDone)

	_, ok := <-jobs
	fmt.Println("receive more jobs: ", ok)
}
