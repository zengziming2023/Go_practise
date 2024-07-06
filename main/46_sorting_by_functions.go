package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"apple", "orange", "grape"}

	lemCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lemCmp)
	fmt.Println(fruits)

	type person struct {
		name string
		age  int
	}

	persons := []person{
		{"Bob", 20},
		{"Jax", 25},
		{"Alice", 27},
	}

	slices.SortFunc(persons, func(a, b person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(persons)

}
