package main

import (
	"fmt"
	"time"
)

func worker(done chan<- bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	doneChan := make(chan bool, 1)
	go worker(doneChan)
	fmt.Println("wait...")
	isDone := <-doneChan
	fmt.Println("isDone:", isDone)
}
