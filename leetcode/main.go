package leetcode

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func LeetCode() {

	var list1 ListNode
	var list1two ListNode
	var list1four ListNode

	list1.Val = 1
	list1two.Val = 2
	list1four.Val = 4

	list1.Next = &list1two
	list1two.Next = &list1four
	list1four.Next = nil

	var list2 ListNode
	var list2three ListNode
	var list2four ListNode
	var list2five ListNode

	list2.Val = 1
	list2three.Val = 3
	list2four.Val = 4
	list2five.Val = 5

	list2.Next = &list2three
	list2three.Next = &list2four
	list2four.Next = &list2five
	list2five.Next = nil
	head := MergeTwoSortedList(&list1, &list2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
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

