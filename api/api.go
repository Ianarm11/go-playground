package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	mux "github.com/gorilla/mux"
)

func SetApiHandlers(r *mux.Router) {
	r.HandleFunc("/getposts/", GetPosts)
	r.HandleFunc("/getpost/{title}", GetPost)
}

func DecodePreview(response *http.Response) Preview {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: Not reading response body")
	}
	var Preview Preview
	json.Unmarshal([]byte(body), &Preview)
	return Preview
}

func DecodePost(response *http.Response) Post {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: Not reading response body")
	}
	var Post Post
	json.Unmarshal([]byte(body), &Post)
	return Post
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		//Make DB call to get data. Using dummy data now
		packet := PreviewPacket{Title: "Bronze Age Mindset", Date: 03122022, Summary: "A journey past the bugmen's cruel world.", Id: "1", Url: "bronzeagepervert"}

		//Send out the encoded data
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(packet)
		return
	} else {
		fmt.Println("Error: Not a GET request in GetPreviews api call")
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//title := mux.Vars(r)["title"]

		//Make DB call to get data. Using dummy data now
		packet := PostPacket{Title: "Bronze Age Mindset", Date: 03122022, Body: "A journey past the bugmen's cruel world.", Id: "1", Url: "bronzeagepervert"}

		//Send out the encoded data
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(packet)
		return
	} else {
		fmt.Println("Error: Not a GET request in GetPreviews api call")
	}
}

type PreviewPacket struct {
	Title   string `json:"title"`
	Date    int    `json:"date"`
	Summary string `json:"summary"`
	Id      string `json:"id"`
	Url     string `json:"url"`
}

type PostPacket struct {
	Title string `json:"title"`
	Date  int    `json:"date"`
	Body  string `json:"body"`
	Id    string `josn:"id"`
	Url   string `json:"url"`
}

type Preview struct {
	Title   string
	Date    int
	Summary string
	Id      string
	Url     string
}

type Post struct {
	Title string
	Date  int
	Body  string
	Id    string
	Url   string
}
