package utilities

import (
	"encoding/json"
	Constants "go-playground/go-playground/constants"
	"log"

	"fmt"
	"os"
)

func WritePostToDatabase(url string) {
	data, _ := json.MarshalIndent(url, "", " ")

	file, err := os.OpenFile(Constants.Database, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Open file error")
		fmt.Println(err)
	}

	_, err = file.Write(data)
	if err != nil {
		file.Close()
		fmt.Println("Write error")
		log.Fatal(err)
	}
}

func ReadPostFromDatabase() {

}
