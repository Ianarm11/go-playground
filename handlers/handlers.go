package handlers

import (
	"fmt"
	api "go-playground/go-playground/api"
	Constants "go-playground/go-playground/constants"
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

	api.SetApiHandlers(r)

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
	temp := template.Must(template.ParseFiles(Constants.PostsTemplate))

	apiUrl := Constants.LocalUrl + Constants.GetPostsApiUrl

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error: GET request in Posts endpoint")
	}
	defer response.Body.Close()

	posts := api.DecodePosts(response)
	fmt.Println(posts)

	temp.Execute(w, posts)
}

func Post(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("templates/post.page.go.tmpl"))

	title := mux.Vars(r)["title"]

	apiUrl := Constants.LocalUrl + Constants.GetPostUrl + title

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error: GET request in Previews endpoint")
	}
	defer response.Body.Close()

	post := api.DecodePost(response)

	temp.Execute(w, post)
}
