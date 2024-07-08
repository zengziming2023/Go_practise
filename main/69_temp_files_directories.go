package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check4(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	temp, err := os.CreateTemp("", "smaple")
	check4(err)

	defer temp.Close()
	defer os.Remove(temp.Name())

	fmt.Println(temp.Name())

	_, err = temp.Write([]byte{1, 2, 3, 4, 5})
	check4(err)

	dname, err := os.MkdirTemp("", "sampleDir")
	check4(err)
	fmt.Println(dname)
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	fmt.Println(fname)
	err = os.WriteFile(fname, []byte{1, 2, 3, 4, 5}, 0666)
	check4(err)

}
