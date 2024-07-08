package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered Error: ", r)
		}
	}()

	mayPanic()
	fmt.Println("after mayPanic")
}
