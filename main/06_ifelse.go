package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("8 is divisible by 2 or 7 is divisible by 2")
	}

	if number := 9; number < 0 {
		fmt.Println("9 is less than zero")
	} else if number < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is more than 10 or equal to 10")
	}
}
