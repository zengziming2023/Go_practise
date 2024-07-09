package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello2(w http.ResponseWriter, req *http.Request) {
	context := req.Context()
	fmt.Println("hello func start")
	defer fmt.Println("hello func end")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-context.Done():
		err := context.Err()
		fmt.Println("err: ", err)
		serverError := http.StatusInternalServerError
		http.Error(w, http.StatusText(serverError), serverError)

	}
}

func main() {
	http.HandleFunc("/hello2", hello2)
	http.ListenAndServe(":8090", nil)
}
