package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an integer")
	boolPtr := flag.Bool("fork", false, "a boolean")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string")

	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	// go build 73_command_line_flags.go
	// ./73_command_line_flags -word=opt -numb=7 -fork -svar=flag
	// ./73_command_line_flags -word=opt
	// ./73_command_line_flags -word=opt a1 a2 a3
	// ./73_command_line_flags -word=opt a1 a2 a3 -numb=7
	// ./73_command_line_flags -h
	// ./73_command_line_flags -wat
}
