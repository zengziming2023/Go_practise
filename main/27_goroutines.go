package main

import (
	"fmt"
	"time"
)

func f2(from string) {
	for i := range 3 {
		fmt.Println(from, " : ", i)
	}
}

func main() {
	f2("direct")

	go f2("goroutine")

	go func(msg string) {
		time.Sleep(time.Millisecond)
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
