package main

import (
	arraysandhashing "go-playground/go-playground/arrays-and-hashing"
	"go-playground/go-playground/leetcode"
)

func main() {
	leetcode.LeetCode()
	arraysandhashing.Arraysandhashing()
}

/*
-- Given an integer array, return true if any integer in the array occurs more than once, otherwise return false.

Solution: Use a hash map to keep track if we have seen a value before.
*/
func ContainsDuplicate(nums []int) bool {

	hashMap := make(map[int]bool)

	for _, value := range nums {
		if _, exists := hashMap[value]; exists {
			return false
		} else {
			hashMap[value] = true
		}
	}
	return true
}

/*
-- Given two strings return two if the strings are palindromes, return false otherwise

//acbbca
loop length: 3
a[0] == a, b[4] == a
a[1] == c, b[3] == c
a[2] == b, b[2] == b

Solution: Compare the first half of the string with the back half of the string.
*/
func ValidPalindrome(a string) bool {
	j := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[j] {
			return false
		}
		i++
		j--
	}
	return true
}

/*
-- Two Sum: given an array of integers 'nums' and an integer 'target', return the indices 'i' and 'j' such that
nums[i] + nums[j] == target and i != j. You can assume that every input has one pair of indices that satisfies the problem.

Return the answer with the smaller index first.
*/

func TwoSum(nums []int, target int) []int {
	//key: value, value: index
	hashMap := make(map[int]int)

	for index, value := range nums {
		hashMap[value] = index
	}

	for index, value := range nums {
		difference := target - value

		if map_index, exists := hashMap[difference]; exists && index != map_index {
			return []int{index, map_index}
		}
	}
	return []int{}
}
