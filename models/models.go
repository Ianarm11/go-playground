package models

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
