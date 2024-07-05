package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		isDone := false
		for !isDone {
			select {
			case <-done:
				fmt.Println("ticker done")
				isDone = true
			case t := <-ticker.C:
				fmt.Println("Tick at: ", t)
			}
		}
		fmt.Println("go func end.")
	}()

	time.Sleep(2 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
