/* HTTP handlers. Executing templates should be the only logic here. Any backend logic should be done somewhere else */
package handlers

import (
	"fmt"
	api "go-playground/go-playground/api"
	"html/template"
	"net/http"
)

func SetHandlers() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/home/", Home)
	http.HandleFunc("/aboutme/", AboutMe)
	http.HandleFunc("/posts/", Previews)
}

func Home(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("templates/home.page.tmpl"))
	temp.Execute(w, nil)
}

func AboutMe(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("templates/aboutme.page.tmpl"))
	temp.Execute(w, nil)
}

func Previews(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("templates/posts.page.tmpl"))

	//Make a GET request to receive data
	//Will be a lists of Previews (obj) that store the title, date, and summary
	response, err := http.Get("http://localhost:8080/getpreviews")
	if err != nil {
		fmt.Println("Error in GET request for Previews")
	}
	defer response.Body.Close()

	post := api.DecodeData(response)
	fmt.Println(post.Title)
	temp.Execute(w, post)
}
