package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check3(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755)
	check3(err)
	defer func() {
		fmt.Println("remove all subdir.")
		err := os.RemoveAll("subdir")
		if err != nil {
			fmt.Printf("Error removing subdir: %v\n", err)
		}
	}()

	createEmptyFile := func(name string) {
		d := []byte("")
		check3(os.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file1")
	err = os.MkdirAll("subdir/parent/child", 0755)
	check3(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check3(err)
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check3(err)

	c, err = os.ReadDir(".")
	check3(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check3(err)

	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
