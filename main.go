package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)

	/*list := data_structures.List{}
	list.Push(2)
	list.Push(3)
	list.Push(1)
	list.Push(400)
	list.ReverseList(list.GetHead())*/

	/*queue := data_structures.Queue{}
	queue = queue.Enqueue(1)
	queue = queue.Enqueue(3)
	queue = queue.Enqueue(60)
	queue = queue.Dequeue()
	queue = queue.Enqueue(4)
	queue.Display()*/
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
