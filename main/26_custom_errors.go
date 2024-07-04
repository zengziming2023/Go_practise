package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg int
	msg string
}

func (err *argError) Error() string {
	return fmt.Sprintf("%d - %s", err.arg, err.msg)
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f1(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println("err match argError success.")
		fmt.Println(ae.arg)
		fmt.Println(ae.msg)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
