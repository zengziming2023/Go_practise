package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for index, num2 := range nums {
		if index == 2 {
			fmt.Println("index = ", index, ", num = ", num2)
		}
	}

	kvs := map[string]string{"key1": "value1", "key2": "value2"}
	fmt.Println(kvs)
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for index, c := range "go" {
		fmt.Printf("%d -> %c\n", index, c)
	}
}
