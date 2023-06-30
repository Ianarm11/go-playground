package handlers

import (
	"fmt"
	api "go-playground/go-playground/api"
	Constants "go-playground/go-playground/constants"
	url "go-playground/go-playground/urlservice"
	"html/template"
	"net/http"
)

func SetHandlers() {
	SetStaticHandlers()
	SetDynamicHandlers()
}

func SetStaticHandlers() {
	//http.HandleFunc("/", Home)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/aboutme/", AboutMe)
	http.HandleFunc("/posts", Previews)
}

func SetDynamicHandlers() {
	urls := url.GetUrls()

	for _, url := range urls {
		fmt.Println(url)
		http.HandleFunc(url.Url, DynamicHandler)
	}
}

func DynamicHandler(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles(Constants.DynamicTemplate))

	//Get the id from url
	id := r.Header.Get("Url")
	fmt.Println("Post Id: " + id)

	//Make GET request to get title, date, and body
	response, err := http.Get("http://localhost:8080/getpost/" + id)
	if err != nil {
		fmt.Println("Error: GET request in Post endpoint")
	}
	defer response.Body.Close()

	post := api.DecodePost(response)
	temp.Execute(w, post)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Info: hitting Home endpoint")
	temp := template.Must(template.ParseFiles(Constants.HomeTemplate))
	temp.Execute(w, nil)
}

func AboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Info: hitting AboutMe endpoint")
	temp := template.Must(template.ParseFiles(Constants.AboutMeTemplate))
	temp.Execute(w, nil)
}

func Previews(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Info: hitting Previews endpoint")
	temp := template.Must(template.ParseFiles(Constants.PreviewTemplate))

	//Make a GET request to receive data
	//Will be a lists of Previews (obj) that store the title, date, and summary
	response, err := http.Get(Constants.LocalUrl + Constants.GetPreviewsApiUrl)
	if err != nil {
		fmt.Println("Error: GET request in Previews endpoint")
	}
	defer response.Body.Close()

	preview := api.DecodePreview(response)
	fmt.Println(preview)
	temp.Execute(w, preview)
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Info: hitting Post endpoint")
	temp := template.Must(template.ParseFiles("templates/post.page.tmpl"))

	//Get the id from url
	id := r.Header.Get("id")
	fmt.Println("Post Id: " + id)

	//Make GET request to get title, date, and body
	response, err := http.Get("http://localhost:8080/getpost/" + id)
	if err != nil {
		fmt.Println("Error: GET request in Post endpoint")
	}
	defer response.Body.Close()

	post := api.DecodePost(response)
	temp.Execute(w, post)
}
