package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	fmt.Println(len(s))

	for i, c := range s {
		fmt.Printf("%d: %c\n", i, c)
		fmt.Printf("%d: %x\n", i, c)
	}

	fmt.Println("Rune Count: ", utf8.RuneCountInString(s))
}
