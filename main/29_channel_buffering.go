package main

import "fmt"

func main() {
	message := make(chan string, 2)

	message <- "hello"
	fmt.Println("send hello success")
	message <- "world"
	fmt.Println("send world success")

	fmt.Println("get message: ", <-message)
	fmt.Println("get message: ", <-message)
}
