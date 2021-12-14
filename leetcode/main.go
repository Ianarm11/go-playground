package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func LeetCode() {

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
func RemoveElementByIndex(x []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, x[:index]...)
	return append(ret, x[index+1:]...)
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
