package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SetApiHandlers() {
	http.HandleFunc("/getpreviews/", GetPreviews)
}

func DecodeData(response *http.Response) DumbPost {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in reading response body")
	}
	var post DumbPost
	json.Unmarshal([]byte(body), &post)
	return post
}

func GetPreviews(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//Log entry point
		fmt.Println("Hit getpreviews endpoint")

		//Make DB call to get data. Using dummy data now
		dumbData1 := DummyPost{Title: "Bronze Age Mindset", Date: 03122022, Summary: "A journey past the bugmen's cruel world."}

		//Send out the encoded data
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dumbData1)
		return
	} else {
		fmt.Println("For some reason its not a GET Request?")
	}
}

type DummyPost struct {
	Title   string `json:"title"`
	Date    int    `json:"date"`
	Summary string `json:"summary"`
}

type DumbPost struct {
	Title   string
	Date    int
	Summary string
}
