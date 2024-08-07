package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here’s a non-blocking receive.
	// If a value is available on messages then select will take the <-messages case with that value.
	// If not it will immediately take the 1 case.
	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("no message receive.")
	}

	// A non-blocking send works similarly.
	// Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver.
	// Therefore the 1 case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message: ", msg)
	default:
		fmt.Println("no message send.")
	}

	// We can use multiple cases above the 1 clause to implement a multi-way non-blocking select.
	// Here we attempt non-blocking receives on both messages and signals.
	select {
	case msg := <-messages:
		fmt.Println("receive message: ", msg)
	case sig := <-signals:
		fmt.Println("receive signal: ", sig)
	default:
		fmt.Println("no activity.")
	}
}
