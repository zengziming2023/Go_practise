package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File {
	fmt.Println("creating file", path)
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return fp
}

func writeFile(fp *os.File) {
	fmt.Println("Writing file...")
	fmt.Fprintln(fp, "data")
}

func closeFile(fp *os.File) {
	fmt.Println("Closing file...")
	err := fp.Close()
	if err != nil {
		fmt.Fprintf(fp, "error: %v\n", err)
		os.Exit(1)
	}

}

func main() {
	fp := createFile("/tmp/defer.txt")
	defer closeFile(fp)
	writeFile(fp)
}
