package text_file_reader

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var path = "/Users/IansIpad/Projects/goworkspace/src/go-playground/go-playground/text-files/"

func Bigdog() {
	//Input here
	file1 := GetTextFile()
	ReadTextFile(file1)
	//ListAllFilesInDirectory()
}

func GetTextFile() *os.File {
	fmt.Println("Attempting to open file...")
	file, err := os.Open(path + "example.txt")
	if err != nil {
		fmt.Println("Opening file has failed.")
		return nil
	} else {
		return file
	}
}

func ListAllFilesInDirectory() {
	files, err := ioutil.ReadDir("/home/data/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}

func ReadTextFile(file *os.File) []string {
	//Creating scanner to parse
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	file.Close()

	var wordsInFile = 0
	var i = 0

	for _, word := range words {
		wordsInFile++
		i++
		if word == words[i] {

		}
		fmt.Println(word)
	}
	fmt.Println(wordsInFile)
	return words
}

func NumberOfWordsInFile(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}