package main

import (
	"fmt"
	h "go-playground/go-playground/handlers"
	"net/http"

	mux "github.com/gorilla/mux"
)

const portNumber = ":8080"

func main() {
	r := mux.NewRouter()
	h.SetHandlers(r)

	err := http.ListenAndServe(portNumber, r)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
