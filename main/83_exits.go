package main

import (
	"fmt"
	"os"
)

func main() {
	// defer not run when call Exit.
	defer fmt.Println("!")
	os.Exit(1)
}
