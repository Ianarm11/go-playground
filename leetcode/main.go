package leetcode

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func LeetCode() {
	x := []int{1,1,1,1,1,11,1,1,1,2}
	fmt.Println("Original Array : ", x)
	size := RemoveDuplicatesFromArray(x)
	fmt.Println("Size of modified array: ", size)
}

func CheckIfPalindrome(str string) bool {
	length := len(str)
	var reverseString string

	for i := length-1; i >= 0 ; i-- {
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
		for j := i+1; j < len(nums); j++ {
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

