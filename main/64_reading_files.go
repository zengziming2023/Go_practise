package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(file))

	openFile, err := os.Open("/tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := openFile.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := openFile.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := openFile.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = openFile.Seek(4, io.SeekCurrent)
	check(err)

	_, err = openFile.Seek(-10, io.SeekEnd)
	//check(err)

	o3, err := openFile.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(openFile, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = openFile.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(openFile)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	openFile.Close()
}
