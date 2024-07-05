package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := range 2 {
		fmt.Println("round: ", i, " start")
		select {
		case msg := <-c1:
			fmt.Println("receive c1: ", msg)
		case msg := <-c2:
			fmt.Println("receive c2: ", msg)
		}
		fmt.Println("round: ", i, " end")
	}

}
