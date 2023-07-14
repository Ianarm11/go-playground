package api

import (
	"encoding/json"
	"fmt"
	Constants "go-playground/go-playground/constants"
	Models "go-playground/go-playground/models"
	Util "go-playground/go-playground/utilities"
	"io/ioutil"
	"net/http"

	mux "github.com/gorilla/mux"
)

func SetApiHandlers(r *mux.Router) {
	r.HandleFunc("/getposts/", GetPosts)
	r.HandleFunc("/getpost/{title}", GetPost)
	r.HandleFunc("/newurl", NewUrl).Methods("POST")
	r.HandleFunc("/moreposts", MorePosts)
}

func DecodePosts(response *http.Response) []Models.Posts {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: Not reading response body")
	}
	var posts []Models.Posts
	json.Unmarshal([]byte(body), &posts)
	return posts
}

func DecodePost(response *http.Response) Models.Post {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: Not reading response body")
	}
	var Post Models.Post
	json.Unmarshal([]byte(body), &Post)
	return Post
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		//Make DB call to get data. Using dummy data now
		packet := Models.PostsPacket{Title: "Bronze Age Mindset", Date: 03122022, Summary: "A journey past the bugmen's cruel world.", Id: "1", Url: "bronzeagepervert"}
		packet2 := Models.PostsPacket{Title: "Zero To One", Date: 07052023, Summary: "Going zero to one is good.", Id: "2", Url: "zerotoone"}

		var packets []Models.PostsPacket
		packets = append(packets, packet)
		packets = append(packets, packet2)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(packets)
		return
	} else {
		fmt.Println("Error: Not a GET request in GetPreviews api call")
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		title := mux.Vars(r)["title"]

		w.Header().Set("Content-Type", "application/json")

		//Make DB call to get data. Using dummy data now
		packet := Models.PostPacket{Title: "Bronze Age Mindset", Date: 03122022, Body: "A journey past the bugmen's cruel world.", Id: "1", Url: "bronzeagepervert"}
		packet2 := Models.PostPacket{Title: "Zero To One", Date: 07052023, Body: "Going zero to one is good.", Id: "2", Url: "zerotoone"}

		if title == packet2.Url {
			json.NewEncoder(w).Encode(packet2)
		} else if title == packet.Url {
			json.NewEncoder(w).Encode(packet)
		}

		return
	} else {
		fmt.Println("Error: Not a GET request in GetPreviews api call")
	}
}

func NewUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//Do error handling
		r.ParseForm()

		//Get value(s) from form
		//Initial Validations!!
		url := r.FormValue("url")
		//Do validation(s)

		//Write to DB
		Util.WritePostToDatabase(url)

		var redirectUrl string
		redirectUrl = Constants.LocalUrl + "posts/" + url + "/"

		fmt.Println(redirectUrl)

		http.Redirect(w, r, redirectUrl, 303)
	}
}

func MorePosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HTMX must be working!")

	if r.Method == "GET" {
		fmt.Println("GET request made?")
	}

	if r.Method == "POST" {
		fmt.Println("It's a post request!")
	}

}
