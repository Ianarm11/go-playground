package handlers

import (
	"fmt"
	"go-playground/go-playground/api"
	Api "go-playground/go-playground/api"
	Constants "go-playground/go-playground/constants"
	Models "go-playground/go-playground/models"
	"html/template"
	"net/http"

	mux "github.com/gorilla/mux"
)

var Serve http.Handler

func SetHandlers(r *mux.Router) {
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/home", Home).Methods("GET")
	r.HandleFunc("/aboutme/", AboutMe)
	r.HandleFunc("/posts/", Posts)
	r.HandleFunc("/posts/{title}/", Post)
	r.HandleFunc("/createpost/", CreatePost)

	Api.SetApiHandlers(r)

	Serve = r
}

func Home(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles(Constants.HomeTemplate))
	temp.Execute(w, nil)
}

func AboutMe(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles(Constants.AboutMeTemplate))
	temp.Execute(w, nil)
}

func Posts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp := template.Must(template.ParseFiles(Constants.PostsTemplate))

		apiUrl := Constants.LocalUrl + Constants.GetPostsApiUrl

		response, err := http.Get(apiUrl)
		if err != nil {
			fmt.Println("Error: GET request in Posts endpoint")
		}
		defer response.Body.Close()

		posts := api.DecodePosts(response)

		temp.Execute(w, posts)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("templates/post.page.go.tmpl"))
	var post Models.Post

	if r.Method == "GET" {
		title := mux.Vars(r)["title"]

		//Validation to check if title is in DB, right now any page will be created
		apiUrl := Constants.LocalUrl + Constants.GetPostUrl + title

		response, err := http.Get(apiUrl)
		if err != nil {
			fmt.Println("Error: GET request in Previews endpoint")
		}
		defer response.Body.Close()

		post = Api.DecodePost(response)
		temp.Execute(w, post)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("templates/newpost.page.go.tmpl"))

	temp.Execute(w, nil)
}
