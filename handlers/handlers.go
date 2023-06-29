/* HTTP handlers. Executing templates should be the only logic here. Any backend logic should be done somewhere else */
package handlers

import (
	"fmt"
	api "go-playground/go-playground/api"
	url "go-playground/go-playground/urlservice"
	"html/template"
	"net/http"
)

func SetHandlers() {
	SetStaticHandlers()
	SetDynamicHandlers()
}

func SetStaticHandlers() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/aboutme/", AboutMe)
	http.HandleFunc("/posts/", Previews)
}

func SetDynamicHandlers() {
	urls := url.GetUrls()

	for _, url := range urls {
		fmt.Println(url)
		http.HandleFunc(url.Url, DynamicHandler)
	}
}

func DynamicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("we are hitting it baby")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Info: hitting Home endpoint")
	temp := template.Must(template.ParseFiles("templates/home.page.tmpl"))
	temp.Execute(w, nil)
}

func AboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Info: hitting AboutMe endpoint")
	temp := template.Must(template.ParseFiles("templates/aboutme.page.tmpl"))
	temp.Execute(w, nil)
}

func Previews(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Info: hitting Previews endpoint")
	temp := template.Must(template.ParseFiles("templates/preview.page.tmpl"))

	//Make a GET request to receive data
	//Will be a lists of Previews (obj) that store the title, date, and summary
	response, err := http.Get("http://localhost:8080/getpreviews")
	if err != nil {
		fmt.Println("Error: GET request in Previews endpoint")
	}
	defer response.Body.Close()

	preview := api.DecodePreview(response)
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
