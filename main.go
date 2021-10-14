package main

import (
	"fmt"
	excel_reader "go-playground/go-playground/excel-reader"
	"html/template"
	"net/http"
	"os"
)

const portNumber = ":8080"
var userInfoArray []UserInfo
var tpl *template.Template

func main() {
	/*http.HandleFunc("/", Home)
	http.HandleFunc("/next", Next)
	http.HandleFunc("/info", Info)

	_ = http.ListenAndServe(portNumber, nil)*/
	excel_reader.ExcelReader()
}

func Home(w http.ResponseWriter, r *http.Request) {
	//renderTemplate(w, "home.page.tmpl")
	tpl.ExecuteTemplate(w, "home.page.tmpl", nil)
}
func Next(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "home.page.tmpl", nil)
	}

	if r.Method == "POST" {
		db := UserInfo{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Firstname: r.FormValue("firstname"),
			Lastname: r.FormValue("lastname"),
			Email: r.FormValue("email"),
		}
		userInfoArray = append(userInfoArray, db)
		fmt.Println("Here")
		err := tpl.ExecuteTemplate(w, "next.page.tmpl", db)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func Info(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "home.page.tmpl", nil)
	}

	if r.Method == "POST" {
		db := userInfoArray[0]
		err := tpl.ExecuteTemplate(w, "info.page.tmpl", db)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func renderTemplate(w http.ResponseWriter, tmpl string) {
	homeDirectory, _ := os.Getwd()
	parsedTemplate, parsingFilesErr := template.ParseFiles(homeDirectory + "/templates/" + tmpl)
	if parsingFilesErr == nil {
		executeErr := parsedTemplate.Execute(w, nil)
		if executeErr != nil {
			fmt.Println("Error in executing templates. Error Output: ")
			fmt.Println(executeErr)
			return
		}
	} else {
		fmt.Println("Error in parsing templates. Error Output: ")
		fmt.Println(parsingFilesErr)
	}
}
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}
type UserInfo struct {
	Username string
	Password string
	Firstname string
	Lastname string
	Email string
}
