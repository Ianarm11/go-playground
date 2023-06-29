package urlservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var path = "/Users/ian/Projects/go-workspace/go-playground/urlservice/"

func GetUrls() []Url {
	data, err := ioutil.ReadFile(path + "urls.json")
	if err != nil {
		fmt.Print(err)
	}

	var urls []Url
	json.Unmarshal(data, &urls)

	return urls
}
