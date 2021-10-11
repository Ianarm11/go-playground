package text_file_reader

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

var path = "/Users/IansIpad/Projects/goworkspace/src/go-playground/go-playground/text-files/"

type book struct {
	value string
	count int
}

func TextReader() {
	file1 := GetTextFile("IF.txt")
	file2 := GetTextFile("Limerick-1-1.txt")
	file1words, file1top3words := ReadTextFile(file1)
	file2words, _ := ReadTextFile(file2)
	grandTotal := len(file1words) + len(file2words)
	WriteToFile(len(file1words), len(file2words), grandTotal, file1top3words)
}

func GetTextFile(fileName string) *os.File {
	fmt.Println("Attempting to open" + fileName + "...")
	file, err := os.Open(path + fileName)
	if err != nil {
		fmt.Println("Opening file has failed.")
		return nil
	} else {
		fmt.Println("File opened with great success!")
		return file
	}
}

func ListAllFilesInDirectory() []os.FileInfo {
	files, _ := ioutil.ReadDir("/home/data/")
	return files
}

func ReadTextFile(file *os.File) ([]string, []string) {
	//Creating scanner to parse
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	file.Close()

	//Total Number of words
	var wordsInFile = 0
	//var Book []book
	var m map[string]int
	m = make(map[string]int)
	var topThreeWords []string

	var newBookCount int = 0
	var inTheBookCount int = 0

	for _, word := range words {
		wordsInFile++
		if InTheMap(word, m) == true {
			inTheBookCount++
			for key, value := range m {
				if word == key{
					value++
				}
			}
		} else {
			newBookCount++
			m[word] = 1
		}
	}
	firstTopValue := GetMaxValue(m)
	delete(m, firstTopValue)
	topThreeWords = append(topThreeWords, firstTopValue)

	secondTopValue := GetMaxValue(m)
	delete(m, secondTopValue)
	topThreeWords = append(topThreeWords, secondTopValue)

	thirdTopValue := GetMaxValue(m)
	delete(m, thirdTopValue)
	topThreeWords = append(topThreeWords, thirdTopValue)

	return words, topThreeWords
}

func InTheMap(word string, Map map[string]int) bool {
	for key := range Map {
		if word == key{
			return true
		}
	}
	return false
}

func GetMaxValue(Map map[string]int) string {
	var wordWithMaxValue string
	var maxValue int
	for key, value := range Map {
		if wordWithMaxValue == "" {
			maxValue = value
			wordWithMaxValue = key
		} else if value > maxValue {
			maxValue = value
			wordWithMaxValue = key
		}
	}
	return wordWithMaxValue
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func WriteToFile(file1WordCount int, file2WordCount int, grandTotalCount int, top3Words []string) {
	ipAddress := GetOutboundIP()
	//fmt.Println(ipAddress)
	files := ListAllFilesInDirectory()

	fmt.Println("Begin writing..")
	f, err := os.Create("results.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("***************************\n")
	f.WriteString("IF.txt word count: \n")
	f.WriteString(fmt.Sprintf("%d", file1WordCount))
	f.WriteString("\n")
	f.WriteString("***************************\n")
	f.WriteString("LIMERICK.txt word count: \n")
	f.WriteString(fmt.Sprintf("%d", file2WordCount))
	f.WriteString("\n")
	f.WriteString("***************************\n")
	f.WriteString("Grand Total word count: \n")
	f.WriteString(fmt.Sprintf("%d", grandTotalCount))
	f.WriteString("\n")
	f.WriteString("***************************\n")
	f.WriteString("Top 3 words from IF.txt \n")
	f.WriteString(top3Words[0] + " " + top3Words[1] + " " + top3Words[2])
	f.WriteString("\n")
	f.WriteString("***************************\n")
	f.WriteString("Outerbound IP Address of this machine: \n")
	f.WriteString(fmt.Sprintf("%s", ipAddress))
	f.WriteString("\n")
	f.WriteString("***************************\n")
	f.WriteString("Files listed at /home/data/: \n")
	if len(files) != 0 {
		for _, file := range files {
			f.WriteString( fmt.Sprintf(file.Name(), file.IsDir()))
		}
	} else {
		f.WriteString("No files were found :(")
	}
	f.WriteString("\n")
	f.WriteString("***************************\n")

	fmt.Println("Done writing..")
}
