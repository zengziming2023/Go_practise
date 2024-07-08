package main

import (
	"fmt"
	"strconv"
)

func main() {
	float, err := strconv.ParseFloat("1.23456789", 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(float)

	parseInt, err := strconv.ParseInt("123", 0, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	parseInt, err = strconv.ParseInt("0x1c8", 0, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
