package main

import "fmt"

func intSec() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func main() {
	initSec := intSec()
	fmt.Println(initSec())
	fmt.Println(initSec())
	fmt.Println(initSec())

	intSec2 := intSec()
	fmt.Println(intSec2())
}
