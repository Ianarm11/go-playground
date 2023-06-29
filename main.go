package main

import (
	"fmt"
	a "go-playground/go-playground/api"
	h "go-playground/go-playground/handlers"
	"net/http"
	//finance "go-playground/go-playground/finance"
)

const portNumber = ":8080"

func main() {
	//Call to handlers to set up handlers
	h.SetHandlers()
	a.SetApiHandlers()

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
	//finance.Finance()

}
