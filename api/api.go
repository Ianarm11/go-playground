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

func DecodePosts(response *http.Response) []Posts {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: Not reading response body")
	}
	var posts []Posts
	json.Unmarshal([]byte(body), &posts)
	return posts
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
		packet := PostsPacket{Title: "Bronze Age Mindset", Date: 03122022, Summary: "A journey past the bugmen's cruel world.", Id: "1", Url: "bronzeagepervert"}
		packet2 := PostsPacket{Title: "Zero To One", Date: 07052023, Summary: "Going zero to one is good.", Id: "2", Url: "zerotoone"}

		var packets []PostsPacket
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
		packet := PostPacket{Title: "Bronze Age Mindset", Date: 03122022, Body: "A journey past the bugmen's cruel world.", Id: "1", Url: "bronzeagepervert"}
		packet2 := PostPacket{Title: "Zero To One", Date: 07052023, Body: "Going zero to one is good.", Id: "2", Url: "zerotoone"}

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

type PostsPacket struct {
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

type Posts struct {
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
