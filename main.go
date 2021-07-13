package main

import "fmt"

func main() {
	fmt.Println("Hello, World.")
	testName := "Ian"
	fooValue := foo(testName)
	fmt.Println(fooValue)
}

func foo(test string) string {
	//var name string = "testing"
	name := "testing"
	if name == test {
		return "They are equal"
	}
	return "They are not equal"
}
