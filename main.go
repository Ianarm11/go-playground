package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hey! How's it going?")
		if err == nil {
			fmt.Println(w, err)
		}
		fmt.Println(fmt.Sprintln("Number of Bytes: ", n))
		fmt.Printf("The Number of Bytes: %d\n", n)
	})
	_ = http.ListenAndServe(portNumber, nil)
}
