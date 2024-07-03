package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("i = ", i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println("j = ", j)
	}

	for k := range 3 {
		fmt.Println("k = ", k)
	}

	for {
		fmt.Println("loop")
		break
	}

	for m := range 6 {
		if m%2 == 0 {
			continue
		}
		fmt.Println("m = ", m)
	}

}
