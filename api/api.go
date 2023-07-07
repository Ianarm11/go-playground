package api

import (
	"encoding/json"
	"fmt"
	Constants "go-playground/go-playground/constants"
	"io/ioutil"
	"net/http"

	mux "github.com/gorilla/mux"
)

func SetApiHandlers(r *mux.Router) {
	r.HandleFunc("/getposts/", GetPosts)
	r.HandleFunc("/getpost/{title}", GetPost)
	r.HandleFunc("/newurl", NewUrl).Methods("POST")
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

func NewUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//Do error handling
		r.ParseForm()

		//Get value(s) from form
		//Initial Validations!!
		url := r.FormValue("url")
		//Do validation(s)

		//Deserialize it into a struct? Probably
		// newPost := Post{
		// 	Title: "test title",
		// 	Url:   url,
		// 	Id:    "test guid",
		// 	Date:  12032023,
		// 	Body:  "test body",
		// }

		//Send to "DB" (post request)
		//DB Validations!!!

		//Make POST request to Post page
		// postBody, err := json.Marshal(newPost)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//responseBody := bytes.NewBuffer(postBody)

		apiUrl := Constants.LocalUrl + "posts/" + url + "/"
		http.Get(apiUrl)
		return
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
