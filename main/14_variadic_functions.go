package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, "   ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	res := sum(1, 2)
	fmt.Println(res)

	res = sum(1, 2, 3)
	fmt.Println(res)

	res = sum(1, 2, 3, 4)
	fmt.Println(res)

	nums := []int{1, 2, 3, 4, 5}
	res = sum(nums...)
	fmt.Println(res)
}
