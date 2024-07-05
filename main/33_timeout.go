package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case result := <-c1:
		fmt.Println("receive c1: ", result)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case result := <-c2:
		fmt.Println("receive c2: ", result)
	case <-time.After(time.Second * 4):
		fmt.Println("timeout")
	}
}
