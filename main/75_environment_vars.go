package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println()

	for _, arg := range os.Environ() {
		args := strings.SplitN(arg, "=", 2)
		fmt.Println(args[0])
	}
}
