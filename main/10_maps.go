package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v := m["k1"]
	fmt.Println("v:", v)

	v2, success2 := m["k2"]
	fmt.Println("v2:", v2, success2)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	v3, success3 := m["k2"]
	fmt.Println("v3:", v3, success3)

	n := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
