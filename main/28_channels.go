package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("put msg into channel")
		// data into channel
		messages <- "message in"
	}()

	time.Sleep(time.Second)
	// get data from channel
	fmt.Println("waiting for message")
	msg := <-messages
	fmt.Println("get message: ", msg)
}
