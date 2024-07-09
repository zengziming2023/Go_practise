package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response Status:", response.Status)

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
