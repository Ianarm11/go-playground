package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World. Welcome to main.go!")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	//StringForRangeLoop()
	//StringForLoop()
	//BasicArray()
	//BasicSlice()
	//ArraySliceExample()
	//BasicMap()
	//BasicStruct()
	//EmbeddedStruct()
	InterfaceExample()
}

func StringForRangeLoop() {
	x := "racecar"
	fmt.Println(x)

	for i, letter := range x {
		fmt.Printf("Index: %d Value: %s\n", i, string(letter))
	}
}

func StringForLoop() {
	x := "IanIan"
	fmt.Println(x)
	var y = string(x[0])
	var z = string(x[1])
	y1 := x[0]
	z1 := x[1]
	fmt.Printf("String value: %s UNIT8 value: %d\n", y, y1)
	fmt.Printf("String value: %s UNIT8 value: %d\n", z, z1)

	//Check if first char has a matching char in string
	for i := 0; i < len(x); i++ {
		fmt.Println(string(x[i]))
		if y == string(x[i+1]) {
			fmt.Printf("Found a match. %s is equal to %s\n", y, string(x[i+1]))
			break
		} else {
			continue
		}
	}
}

func BasicArray() {
	//Different methods to initialize arrays
	var x [5]int
	x[0] = 100
	x[1] = 200
	x[2] = 300
	x[3] = 400
	x[4] = 500
	//var x1 = [5]int{100, 200, 300, 400, 500}
	//x2 := [5]int{100, 200, 300, 400, 500}
	//Array literal
	//x3 := [...]int{100, 200, 300, 400, 500}
	x4 := [...]int{100, 200, 300, 400, 500}
	if x4 == x {
		fmt.Println("Equal arrays!")
	}

	fmt.Println("len of 5 array: ", len(x))
	var y [0]int
	var z []int
	fmt.Println("len of O array: ", len(y))
	fmt.Println("len of nil array: ", len(z))
}

func BasicSlice() {
	slice := make([]string, 3)
	fmt.Printf("Empty slice content: %s Empty slice length: %d\n\n", slice, len(slice))
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Printf("Slice content: %s Slice length: %d\n\n", slice, len(slice))

	//append() returns a new slice
	slice = append(slice, "d", "e")
	fmt.Printf("Slice appended content: %s Slice appended length: %d\n\n", slice, len(slice))

	//slices can be copied into other slices
	cSlice := make([]string, len(slice))
	copy(cSlice, slice)
	fmt.Printf("Copied slice content: %s Copied slice length: %d\n\n", cSlice, len(cSlice))

	sSlice := slice[:2]
	sSlice1 := slice[:5]
	fmt.Printf("Sliced slice content: %s Sliced slice length: %d\n\n", sSlice, len(sSlice))
	fmt.Printf("Sliced slice1 content: %s Sliced slice1 length: %d\n\n", sSlice1, len(sSlice1))
}

func ArraySliceExample() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Printf("months (array) content: %s months (array) length: %d\n\n", months, len(months))
	//Length is 13 because arrays start at [0]. We initialized starting at [1]

	summer := months[6:9]
	fall := months[9:12]
	spring := months[3:6]
	//Switching the order of these 2 parameters in append() make Spring print out 'December' instead of 'March'
	//Why?
	winter := append(months[12:], months[1:3]...)
	fmt.Printf("Summer: %s Fall: %s Winter: %s Spring: %s\n\n", summer, fall, winter, spring)
	fmt.Printf("Summer len: %d Fall len: %d Winter len: %d Spring len: %d\n\n", len(summer), len(fall), len(winter), len(spring))
	fmt.Printf("Summer cap: %d Fall cap: %d Winter cap: %d Spring cap: %d\n\n", cap(summer), cap(fall), cap(winter), cap(spring))

	//For loop to help understand capacity
	//ie. the 3rd param in make() is the set capacity for underlying array in splice
	s := make([]int, 0)
	for i := 1; i <= 20; i++ {
		s = append(s, i)
		fmt.Printf("cap: %v, len: %v, val: %d, address: %p\n", cap(s), len(s), s, s)
	}
}

func BasicMap() {
	//Key type, value type
	dictionary := make(map[string]int)
	dictionary["Zurf"] = 1
	dictionary["Hinterlands"] = 2
	dictionary["Vassal"] = 3
	dictionary["Idiom"] = 3
	fmt.Println(dictionary)
	value1 := dictionary["Hinterlands"]
	fmt.Println(value1)
	//To check if there is a value. Not nil, "", or 0
	_, check := dictionary["Zurf"]
	fmt.Println(check)
}

func BasicStruct() {
	var kochland Book
	kochland.Pages = 671
	kochland.Author = "Christopher Leonard"
	kochland.Name = "Kochland"

	name := &kochland.Name
	*name = *name + ": The Secret History of Corporate America"
	fmt.Printf("Name of book (pointer): %s\n\n", *name)
}

//------Stucts------//

type Circle struct {
	X int
	Y int
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}
type Fan struct {
	Circle Circle
	Blades int
}
func EmbeddedStruct() {
	var circle Circle
	var wheel Wheel
	var fan Fan

	circle.Y = 10
	circle.X = 5
	circle.Radius = 80

	wheel.Circle = circle
	wheel.Spokes = 20
	wheel.Y = 500
	wheel.X = 100
	wheel.Radius = 200000

	fmt.Printf("wheel.Circle: %d wheel.Spokes: %d \n\n", wheel.Circle, wheel.Spokes)
	fan.Circle = circle
	fan.Blades = 17
	fan.Circle.X = 777
	fan.Circle.Y = 888
	fan.Circle.Radius = 999
	fmt.Printf("fan.Circle: %d fan.Blades: %d \n\n", fan.Circle, fan.Blades)
}

//-----Interfaces-----//

type Book struct {
	Author string
	Name string
	Pages int
}
type Magazine struct {
	Author string
	Title string
	Pages int
}

func (b Book) Read() string {
	output := "Title: " + b.Name + " Author: " + b.Author + " Pages: " + strconv.Itoa(b.Pages)
	return output
}
func (m Magazine) Read() string {
	output := "Title: " + m.Title + " Author: " + m.Author + " Pages: " + strconv.Itoa(m.Pages)
	return output
}

type Reader interface {
	Read() string
}

func PrintContents(r Reader) {
	fmt.Println(r.Read())
}

func InterfaceExample() {
	book := Book{"David J Armstrong", "Kochland", 651}
	magazine := Magazine{"John E Ross", "The Vagabonds", 350}
	PrintContents(book)
	PrintContents(magazine)
}