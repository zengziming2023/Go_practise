package main

import (
	"embed"
	"fmt"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	fmt.Println(fileString)
	fmt.Println(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	fmt.Println(string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	fmt.Println(string(content2))

	// mkdir -p folder
	// echo "hello go" > folder/single_file.txt
	// echo "123" > folder/file1.hash
	// echo "456" > folder/file2.hash
}
