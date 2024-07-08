package main

import (
	"bufio"
	"fmt"
	"os"
)

func check2(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check2(err)

	create, err := os.Create("/tmp/dat2")
	check2(err)

	defer create.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := create.Write(d2)
	check2(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := create.WriteString("writes\n")
	check2(err)
	fmt.Printf("wrote %d bytes\n", n3)

	create.Sync()
	w := bufio.NewWriter(create)
	n4, err := w.WriteString("buffered\n")
	check2(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
