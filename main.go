package main

import (
	data_structures "data-structures"
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func main() {
	//http.HandleFunc("/", Home)
	//http.HandleFunc("/about", About)

	//_ = http.ListenAndServe(portNumber, nil)

	list := data_structures.List{}
	list.Push(2)
	list.Push(3)
	list.Push(1)
	list.Push(400)
	list.Push("Hello")
	list.DisplayTail()
	list.Display()
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, parsingFilesErr := template.ParseFiles("../src/templates/" + tmpl)
	if parsingFilesErr == nil {
		executeErr := parsedTemplate.Execute(w, nil)
		if executeErr != nil {
			fmt.Println("Error in executing templates. Error Output: ")
			fmt.Println(executeErr)
			return
		}
	} else {
		fmt.Println("Error in parsing templates. Error Output: ")
		fmt.Println(parsingFilesErr)
	}
}
