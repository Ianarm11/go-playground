package leetcode

import (
	"errors"
	"fmt"
	stack "go-playground/go-playground/stack"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

var board = [][]byte{
	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	{'.', '.', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '.', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '.'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
}
var hourglassarray = [][]int32{
	{1, 1, 1, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{1, 1, 1, 0, 0, 0},
	{0, 0, 2, 4, 4, 0},
	{0, 0, 0, 2, 0, 0},
	{0, 0, 1, 2, 4, 0},
}

func LeetCode() {
	res := CarFleet(10, []int{1, 4}, []int{3, 2})
	fmt.Println(res)
}

// ***QUESTIONS****//
type Fleet struct {
	Speed    int
	Position int
}

func CarFleet(target int, positions []int, speed []int) int {
	n := len(positions)
	endFleetCounter := 0
	flag := true
	var carFleets []Fleet

	//Init fleets
	for i := 0; i < n; i++ {
		fleet := Fleet{Position: positions[i], Speed: speed[i]}
		carFleets = append(carFleets, fleet)
	}

	for flag {
		if len(carFleets) == 0 {
			flag = false
			fmt.Println("Length of car fleets is 0.")
			break
		}

		//Update next positions
		for i := 0; i < len(carFleets); i++ {
			carFleets[i].Position = carFleets[i].Position + carFleets[i].Speed
		}

		hash := make(map[int]Fleet) //Key: position Value: Fleet
		//Check if any fleets have been absorbed
		for i := 0; i < len(carFleets); i++ {
			position := carFleets[i].Position
			existingFleet, exists := hash[position]

			if exists {
				fmt.Println("Position in hashmap exists, checking speeds")
				if carFleets[i].Speed > existingFleet.Speed {
					fmt.Println("The current fleet's speed is greater than the exisiting fleets")
					existingFleet.Speed = carFleets[i].Speed
				}
			} else {
				fmt.Println("Adding new fleet to hashmap")
				hash[position] = carFleets[i]
			}
		}

		//Check if any are finished
		for i := 0; i < len(carFleets); i++ {
			if carFleets[i].Position == target {
				fmt.Println("Fleet has reached target, updating counter, removing from the fleet")
				endFleetCounter++
				carFleets = append(carFleets[:i], carFleets[i+1:]...)
			}
		}
	}
	return endFleetCounter
}

func JumpingOnClouds(clouds []int) int {
	jumps := 0
	i := 0

	for i < len(clouds)-1 {
		if i+2 < len(clouds) && clouds[i+2] == 0 {
			fmt.Println("Hoping over twice")
			jumps++
			i += 2
		} else {
			fmt.Println("Hoping over once")
			jumps++
			i += 1
		}
	}
	return jumps
}
func CountVallyes(n int, path string) int {
	seaLevel := 0
	valleyCount := 0

	for _, step := range path {

		if seaLevel == -1 && step == 'U' {
			valleyCount++
		}

		if step == 'U' {
			seaLevel++
		}

		if step == 'D' {
			seaLevel--
		}
	}
	return valleyCount
}

func PairOfSocks(n int, socks []int) int {
	matchingCounter := 0
	basket := make(map[int]int)

	for i := 0; i < len(socks); i++ {
		sock := socks[i]
		_, exists := basket[sock]

		if exists {
			fmt.Println("Socks is in the basket already, removing it")
			delete(basket, sock)
			matchingCounter = matchingCounter + 1
			fmt.Printf("Counter: %d\n", matchingCounter)

		} else {
			fmt.Println("Socks is not in the basket, adding it")
			basket[sock] = 1
		}

		fmt.Println(basket)
	}

	return matchingCounter
}

func DailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	indices := stack.Constructor_Int()

	for i := 0; i <= len(temperatures)-1; i++ {

		previousIndex, _ := indices.Top()

		if previousIndex != -5000 && temperatures[i] > temperatures[previousIndex] {
			fmt.Printf("Current temperature is greater: %d than the what is in the stack: %d \n", temperatures[i], temperatures[previousIndex])
			indices.Pop()
			results[previousIndex] = i - previousIndex
		}
		indices.Push(i)
	}

	return results
}

func GenerateParanthesis(n int) []string {
	var results []string

	GenerateString(n, "", 0, 0, &results)
	return results
}

func GenerateString(count int, currentString string, openingCounter int, closingCounter int, results *[]string) {
	if openingCounter == count && closingCounter == count {
		*results = append(*results, currentString)
		return
	}

	if openingCounter < count {
		GenerateString(count, currentString+"(", openingCounter+1, closingCounter, results)
	}

	if closingCounter < openingCounter {
		GenerateString(count, currentString+")", openingCounter, closingCounter+1, results)
	}
}

func ReversePolishNotation(tokens []string) int {
	operators := []string{"+", "-", "*", "/"}
	operands := stack.Constructor()

	var expression int

	for _, str := range tokens {
		fmt.Println("New string: " + str)
		var operator string = ""

		//If operand, add to stack
		if !contains(operators, str) {
			fmt.Println("Found an operand, adding to stack")
			operands.Push(str)
		} else {
			fmt.Println("Found an operator: " + operator)
			operator = str
		}

		if operator != "" {
			fmt.Println("Since we found an operator, time to compute")

			val1, err := operands.Top()
			if err == stack.ErrorEmptyStack {
				fmt.Println("Empty stack for value 1")
			}

			if val1 != "" {
				fmt.Println("Popping stack")
				operands.Pop()
			}

			val2, err := operands.Top()
			if err == stack.ErrorEmptyStack {
				fmt.Println("Empty stack for value 2")
			}

			if val2 != "" {
				fmt.Println("Popping stack")
				operands.Pop()
			}

			if val1 != "" && val2 != "" {
				fmt.Println("We have two operands, let's compute")

				result, err := performOperation(val1, val2, operator)
				if err != nil {
					fmt.Println("fuccccccckkkkkk")
				}
				fmt.Printf("Result of two operands: %d\n", result)
				expression += result
			} else if val1 != "" && val2 == "" {
				fmt.Println("We just have one operand to compute with")

				result, err := performOperation(val1, strconv.Itoa(expression), operator)
				if err != nil {
					fmt.Println("fucccccccck again")
				}
				fmt.Printf("Result of one operands: %d\n", result)
				expression += result
			}
		}
	}
	return expression
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func validStackString(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var stack []string
	for _, value := range s {

		str := string(value)

		if str == "[" || str == "{" || str == "(" {
			//Add to stack
			stack = append(stack, str)
		}

		if str == "]" {
			//Check stack, if it's not the opening, fail
			var top string
			if len(stack) != 0 {
				top = string(stack[len(stack)-1])
			}
			if top != "[" {
				return false
			}
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else if str == "}" {
			//Check stack, if it's not the opening, fail
			var top string
			if len(stack) != 0 {
				top = string(stack[len(stack)-1])
			}
			if top != "{" {
				return false
			}
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else if str == ")" {
			//Check stack, if it's not the opening, fail
			var top string
			if len(stack) != 0 {
				top = string(stack[len(stack)-1])
			}
			if top != "(" {
				return false
			}
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

var preordertreearray []int = []int{}

func printPreorderTree(root *TreeNode) {
	if root == nil {
		return
	}
	traverseTree(root)
	for i, val := range preordertreearray {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}

func traverseTree(root *TreeNode) {
	if root == nil {
		return
	}
	preordertreearray = append(preordertreearray, root.Data)
	traverseTree(root.Left)
	traverseTree(root.Right)
}

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func printLinkedList(head *SinglyLinkedListNode) {

	flag := true

	for flag {
		fmt.Println(head.data)

		if head.next != nil {
			head = head.next
		} else {
			flag = false
		}
	}
}

func matchingStrings(stringList []string, queries []string) []int32 {
	res := make([]int32, len(queries))
	for i := 0; i <= len(queries)-1; i++ {
		for j := 0; j <= len(stringList)-1; j++ {
			if queries[i] == stringList[j] {
				res[i]++
			}
		}
	}
	return res
}

func rotateLeft(d int32, arr []int32) []int32 {
	if d == 0 {
		return arr
	}

	for i := int32(0); i < d; i++ {
		var _pop = arr[0]
		fmt.Println(_pop)
		arr = append(arr[1:], _pop)
	}
	return arr
}

// Hour Glass Sum question. 6x6 array, there are 16 hourglasses in an hourglass array, calculate the maximu, hourglass
func hourglassSum(arr [][]int32) int32 {
	var max int32 = 0
	for i := 0; i <= 3; i++ { //rows 0 to 3
		for j := 0; j <= 3; j++ { //columns 0 to 3
			//calculate the top row
			toprow := arr[i][j] + arr[i][j+1] + arr[i][j+2]
			//calculate the middle row
			middlerow := arr[i+1][j+1]
			//calculate the bottom row
			bottomrow := arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			result := toprow + middlerow + bottomrow
			if result > max {
				max = result
			}
		}
	}
	return max
}

type square_key struct {
	a int
	b int
}

func isValidSudokuBoard(board [][]byte) bool {
	// Create maps for each row, column, and 3x3 square
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make(map[square_key]map[byte]bool)

	// Initialize all maps
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
	}

	// Initialize squares map
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			squares[square_key{a: i, b: j}] = make(map[byte]bool)
		}
	}

	// Check each cell in the board
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			value := board[row][col]
			if value == '.' {
				continue
			}

			// Check row
			if rows[row][value] {
				return false
			}
			rows[row][value] = true

			// Check column
			if cols[col][value] {
				return false
			}
			cols[col][value] = true

			// Check 3x3 square
			key := square_key{a: row / 3, b: col / 3}
			if squares[key][value] {
				return false
			}
			squares[key][value] = true
		}
	}
	return true
}

func GroupAnagrams(list []string) [][]string {
	hashMap := make(map[string][]string)
	for _, value := range list {

		bytes := []byte(value)
		for x := 0; x < len(bytes); x++ {
			for y := x + 1; y < len(bytes); y++ {
				//if the current index's value is less thant the index's value ahead, swap them
				if bytes[x] < bytes[y] {
					bytes[x], bytes[y] = bytes[y], bytes[x]
				}
			}
		}
		sortedVal := string(bytes)

		hashMap[sortedVal] = append(hashMap[sortedVal], sortedVal)
	}

	result := make([][]string, 0, len(hashMap))
	for _, group := range hashMap {
		result = append(result, group)
	}
	return result
}

func Encode(list []string) string {
	var encodedString string = ""
	for _, val := range list {
		length := len(val)
		fmt.Printf("String: %s Length: %d\n", val, length)
		lengthAsString := strconv.Itoa(length)
		formattedElement := lengthAsString + val
		fmt.Printf("Formatted element: %s\n", formattedElement)
		encodedString = encodedString + formattedElement
	}

	fmt.Printf("Encoded string: %s\n", encodedString)
	return encodedString
}

func Decode(encodedString string) []string {
	//lengthMarker := int(encodedString[0])
	var decodedStrings []string

	i := 0
	for i < len(encodedString) {
		// Get length marker
		lengthStr := string(encodedString[i])
		lengthMarker, _ := strconv.Atoi(lengthStr)
		fmt.Printf("Epoch #%d -- length of string: %s length marker: %d  \n", i, lengthStr, lengthMarker)

		// Move to start of word
		i++

		// Build word
		newWord := ""
		for j := 0; j < lengthMarker; j++ {
			newWord += string(encodedString[i+j])
		}
		decodedStrings = append(decodedStrings, newWord)

		// Jump to next length marker
		i += lengthMarker
		fmt.Printf("Value of i: %d at end of new word: %s\n", i, newWord)
	}
	return decodedStrings
}

func LongestSubarrayWithoutRepeatingCharacters(array []string) {
	windowStart := 0
	max := 0
	hashMap := make(map[string]int)
	for end := 0; end < len(array); end++ {
		fmt.Printf("New index: %d -- Value: %s \n", end, array[end])
		//Check hashmap if we have see this string before
		prevIndex, seen := hashMap[array[end]]

		if seen && prevIndex >= windowStart {
			//We have see this string before and it's recorded index is ahead of the window's start frame
			windowStart = prevIndex + 1
			fmt.Printf("We have seen this letter before, start of window is moving to index: %d. Letter: %s \n", end, array[prevIndex+1])
		}

		//Add/update to hash map. key == string, value == index
		hashMap[array[end]] = end

		localLength := end - windowStart + 1
		fmt.Printf("Local length: %d\n", localLength)
		if localLength > max {
			max = localLength
			fmt.Printf("New global max: %d\n", max)
		}
		fmt.Println("Map contents:")
		for key, value := range hashMap {
			fmt.Printf("  %s: %d\n", key, value)
		}

	}
	fmt.Printf("Global Max: %d\n", max)
}

func MaximumSubarrayOfSizeK(k int, array []int) int {
	//Get first window
	firstSum := 0
	for i := 0; i < k; i++ {
		fmt.Printf("Number: %d\n", i)
		firstSum += array[i]
	}
	fmt.Printf("First window sum: %d\n", firstSum)
	//Slide window
	globalMax := firstSum
	for i := 1; i < len(array)-k; i++ {
		fmt.Printf("Loop #%d \n", i)

		numberEntering := array[i+k-1]
		fmt.Printf("Number entering window: %d\n", numberEntering)
		numberLeaving := array[i-1]
		fmt.Printf("Number leaving window: %d\n", numberLeaving)

		firstSum += numberEntering
		firstSum -= numberLeaving

		fmt.Printf("Loop sum: %d\n", firstSum)
		if firstSum > globalMax {
			globalMax = firstSum
			fmt.Printf("New global max: %d\n", globalMax)
		}
	}

	fmt.Printf("Final result: %d\n", globalMax)
	return globalMax
}

func CheckIfPalindrome(str string) bool {
	length := len(str)
	var reverseString string

	for i := length - 1; i >= 0; i-- {
		reverseString += string(str[i])
	}

	if reverseString == str {
		return true
	} else {
		return false
	}
}
func MergeTwoSortedList(listHead1 *ListNode, listHead2 *ListNode) *ListNode {
	//Empty Node to hold merged list
	mergedList := &ListNode{}
	//Head of merged list
	mergedListHead := mergedList

	for listHead1 != nil && listHead2 != nil {
		if listHead1.Val < listHead2.Val {
			mergedListHead.Next = listHead1
			mergedListHead = mergedListHead.Next
			listHead1 = listHead1.Next
		} else {
			mergedListHead.Next = listHead2
			mergedListHead = mergedListHead.Next
			listHead2 = listHead2.Next
		}
	}
	if listHead1 != nil {
		mergedListHead.Next = listHead1
	} else if listHead2 != nil {
		mergedListHead.Next = listHead2
	}

	fmt.Println("***************")
	return mergedList.Next
}
func RemoveDuplicatesFromArray(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums = RemoveElementByIndex(nums, j)
				j--
				fmt.Println(nums)
			}
		}
	}
	fmt.Println(nums)
	return len(nums)
}
func NextPermutation(nums []int) []int {
	permutation := make([]int, len(nums))
	copy(permutation, nums)

	//Traverse through nums in a reverse manner
	//Start i at second to last index, start j at last index
	for i := len(nums) - 2; i > 0; i-- {
		head := permutation[i]
		tail := permutation[i+1]
		if head < tail {
			Swap(permutation, i, i+1)
			break
		}
	}
	if EqualArrays(permutation, nums) == true {
		ReverseArray(permutation)
	}
	return permutation
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := ListNode{0, nil}
	current := &dummyHead

	p := l1
	q := l2
	carry := 0

	for p != nil && q != nil {
		var x int
		var y int

		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum := x + y + carry
		carry = sum / 10

		current.Next = &ListNode{sum % 10, nil}
		current = current.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry != 0 {
		current.Next = &ListNode{carry, nil}
	}
	return dummyHead.Next
}
func BinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	middle := low + (high-low)/2
	result := -1

	for low <= high {
		if nums[middle] == target {
			result = middle
			break
		} else if nums[middle] < target {
			middle = middle + 1
			low = middle
		} else if nums[middle] > target {
			middle = middle - 1
			high = middle
		}
	}
	return result
}
func UnknownSizeSearch(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}

	low := 0
	high := 1
	result := -1

	for nums[high] < target {
		low = high
		high <<= 1
	}

	for low <= high {
		middle := low + (high-low)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			low = middle + 1
		} else if nums[middle] > target {
			high = middle - 1
		}
	}
	return result
}
func MoveZeros(nums []int) {
	fmt.Println("Original:  ", nums)
	num1 := make([]int, len(nums))
	copy(num1, nums)
	num2 := make([]int, len(nums))
	copy(num2, nums)

	lastNonZeroIndex := 0
	for currentIndex := 0; currentIndex < len(num1); currentIndex++ {
		if num1[currentIndex] != 0 {
			temp := num1[lastNonZeroIndex]
			num1[lastNonZeroIndex] = num1[currentIndex]
			num1[currentIndex] = temp
			lastNonZeroIndex++
		}
	}
	fmt.Println("Incorrect: ", num1)

	lastNonZeroIndex = 0
	for currentIndex := 0; currentIndex < len(num2); currentIndex++ {
		if num2[currentIndex] != 0 {
			//Trick for swapping index values
			num2[lastNonZeroIndex], num2[currentIndex] = num2[currentIndex], num2[lastNonZeroIndex]
			lastNonZeroIndex++
		}
	}
	fmt.Println("Correct:   ", num2)
}
func TwoSum(numbers []int, target int) []int {
	indexes := make([]int, 2)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				indexes[0] = i + 1
				indexes[1] = j + 1
			}
		}
	}
	//fmt.Println(indexes)
	return indexes
}
func ReverseWordsInString(s string) string {
	bytes := len(s)
	//strings.Split(s, " ")
	/*for _, word := range s {
		fmt.Printf("%c \n", word)
	}*/
	for i := 0; i < bytes; i++ {
		if s[i] == ' ' {
			//Reverse the strong in s[i-1]
			//i = i+1
		}
	}
	return s
}
func LengthOfLongestSubstringSlidingWindow(s string) int {
	if len(s) == 1 {
		return 1
	}
	maxLength := 1
	left := 0
	for right := 1; right < len(s); right++ {
		if s[right] != s[left] {
			left = right
		}

		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	fmt.Println(maxLength)
	return maxLength

}
func LengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	var globalSubstring []byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("Looking at new sustring. Base byte: %d\n", s[i])
		var localSubstring []byte = []byte{s[i]}
		for j := i + 1; j < len(s); j++ {
			if s[j] == s[i] {
				fmt.Println("Found a match.")
				localSubstring = append(localSubstring, s[j])
				if len(localSubstring) > len(globalSubstring) {
					fmt.Println("Local substring is greater than global")
					globalSubstring = localSubstring
				}
			} else {
				fmt.Println("Characters don't match. Moving onto next substring.")
				break
			}
		}
	}
	fmt.Println(len(globalSubstring))
	return len(globalSubstring)
}

//***HELPER FUNCTIONS****//

func performOperation(leftStr, operator, rightStr string) (int, error) {
	left, err := strconv.Atoi(leftStr)
	if err != nil {
		return 0, fmt.Errorf("invalid left operand: %s", leftStr)
	}

	right, err := strconv.Atoi(rightStr)
	if err != nil {
		return 0, fmt.Errorf("invalid right operand: %s", rightStr)
	}

	switch operator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, errors.New("division by zero")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func RemoveElementByIndex(x []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, x[:index]...)
	return append(ret, x[index+1:]...)
}
func Swap(nums []int, i int, j int) []int {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
	return nums
}
func EqualArrays(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func ReverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func ReverseArrayInt32(arr []int32) []int32 {
	if len(arr) == 1 {
		return arr
	}

	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]

		left++
		right--
	}

	return arr
}

func Max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
func TopKFrequent(nums []int, k int) []int {
	myMap := make(map[int]int)

	for _, val := range nums {
		myMap[val]++
	}

	//Loop through k amount of times
	var numbers []int
	for i := 0; i < k; i++ {
		//Get highest count, then remove that entry from map
		var maxValue int
		var maxKey int
		for key, value := range myMap {
			if value > maxValue {
				maxKey = key
				maxValue = value
			}
		}
		numbers = append(numbers, maxKey)
		delete(myMap, maxKey)
	}
	return numbers
}
