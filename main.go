package main

import (
	"fmt"
	a "go-playground/go-playground/api"
	h "go-playground/go-playground/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	h.SetHandlers()
	a.SetApiHandlers()

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
