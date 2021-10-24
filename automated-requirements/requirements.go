package automated_requirements

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var path = "/Users/IansIpad/Projects/goworkspace/src/go-playground/go-playground/text-files/"

type TraceObj struct {
	values []string
}

func AutomateRequirements() {
	file := GetTextFile("requirements-3nfr-60fr.txt")
	wordsInFile := Gather(file)
	NFRs, FRs := Sort(wordsInFile)
	NFRs, FRs = Clean(NFRs, FRs)
	tracedResults := Trace(NFRs, FRs)
	OutputResults(tracedResults)
	fmt.Println("The requirements have been automatically traced. Great Success!")
}

func OutputResults(results []TraceObj) {
	fmt.Println("Begin writing..")
	f, err := os.Create("automated-requirements-output.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if len(results) != 0 {
		for _, res := range results {
			f.WriteString(strings.Join(res.values, ","))
			f.WriteString("\n")
		}
	}
	fmt.Println("End writing..")
}

func JackardPercent(sharedMembers int, totalMembers int) float64 {
	percent := (float64(sharedMembers) / float64(totalMembers)) * 100
	//traceFlag := Threshold(jackardPercent)
	return percent
}

func Threshold(percent float64) string {
	var x int = 0
	if percent < 17.00 {
		x = 0
	} else {
		x = 1
	}
	return strconv.Itoa(x)
}

func Trace(NFRs [][]string, FRs [][]string) []TraceObj {
	var traceResults []TraceObj
	score := 0
	var DataSet []float64

	for _, FR := range FRs {
		trace := TraceObj{}
		trace.values = append(trace.values, FR[0])
		for _, NFR := range NFRs {
			for i := 1; i < len(FR); i++ {
				for j := 1; j < len(NFR); j++ {
					if FR[i] == NFR[j] {
						score++
					}
				}
			}
			overTotal := len(FR) + len(NFR)
			trace.values = append(trace.values, Threshold(JackardPercent(score, overTotal)))
			DataSet = append(DataSet, JackardPercent(score, overTotal))
			score = 0
		}
		traceResults = append(traceResults, trace)
	}
	Median(DataSet)
	return traceResults
}

func Clean(NFRs [][]string, FRs [][]string) ([][]string, [][]string) {

	for _, NFR := range NFRs {
		for i := 0; i < len(NFR); i++ {
			if strings.Contains(NFR[i], ".") || strings.Contains(NFR[i], ":") {
				NFR[i] = strings.ReplaceAll(NFR[i], ".", "")
				NFR[i] = strings.ReplaceAll(NFR[i], ":", "")
			}
			if strings.Contains(NFR[i], "the")  ||  strings.Contains(NFR[i], "The") {
				NFR[i] = strings.ReplaceAll(NFR[i], "the", "")
				NFR[i] = strings.ReplaceAll(NFR[i], "The", "")
			}
		}
	}
	for _, FR := range FRs {
		for i := 0; i < len(FR); i++ {
			if strings.Contains(FR[i], ".") || strings.Contains(FR[i], ":") {
				FR[i] = strings.ReplaceAll(FR[i], ".", "")
				FR[i] = strings.ReplaceAll(FR[i], ":", "")
			}
			if strings.Contains(FR[i], "the") ||  strings.Contains(FR[i], "The") {
				FR[i] = strings.ReplaceAll(FR[i], "the", "")
				FR[i] = strings.ReplaceAll(FR[i], "The", "")
			}
		}
	}
	return NFRs, FRs
}

func Sort(words []string) ([][]string, [][]string) {
	var NFR [][]string
	var FR [][]string
	length := len(words)

	//Loop through each word
	for i := 0; i < length; i++ {
		var line []string
		line = append(line, words[i])
		//If the current word is an NFR, loop through the next words until it hits another NFR or FR
		if strings.Contains(words[i], "NFR") {
			for j := i+1; j < length; j++ {
				if strings.Contains(words[j], "NFR") || strings.Contains(words[j], "FR") {
					break
				} else {
					line = append(line, words[j])
				}
			}
			NFR = append(NFR, line)
		//If the current word is an FR, loop through the next words until it hits another FR
		} else if strings.Contains(words[i], "FR") {
			for j := i+1; j < length; j++ {
				if strings.Contains(words[j], "FR") {
					break
				} else {
					line = append(line, words[j])
				}
			}
			FR = append(FR, line)
		}
	}
	return NFR, FR
}

func Gather(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var wordsInDocument []string

	for scanner.Scan() {
		wordsInDocument = append(wordsInDocument, scanner.Text())
	}
	file.Close()

	return wordsInDocument
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

func Median(values []float64) {
	sum := 0.00
	for _, val := range values {
		sum += val
	}
	median := math.Round( sum / float64(len(values)) )
	fmt.Println(median)
}
