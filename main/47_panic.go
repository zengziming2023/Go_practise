package main

import (
	"fmt"
	"os"
)

func main() {
	//panic("a problem")

	file, err := os.Create("/temp/file")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(file)
}
