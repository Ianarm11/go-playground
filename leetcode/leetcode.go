package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func LeetCode() {
	//nums := []int{0,1,2,0,3,4,5,0,6,7,0,8,9}
	LengthOfLongestSubstring("abcacbbcbbdfghjik")
}

//***QUESTIONS****//

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
func LengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	bestSol := 0
	start := 0
	seenCharacters := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		duplicateIndex, flag := seenCharacters[s[i]]

		if flag {
			if bestSol < len(seenCharacters) {
				bestSol = len(seenCharacters)
			}

			fmt.Println("Duplicate Index: ", duplicateIndex, "Start: ", start)
			for j := start; j <= duplicateIndex; j++ {
				delete(seenCharacters, s[j])
				fmt.Printf("Keys that just got deleted: %c \n", s[j])
			}
			start = duplicateIndex + 1
			fmt.Println("New Start: ", start)
		}
		seenCharacters[s[i]] = i
	}

	if bestSol < len(seenCharacters) {
		bestSol = len(seenCharacters)
	}
	fmt.Println("Final Best Solution: ", bestSol)
	return bestSol
}

//***HELPER FUNCTIONS****//

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
