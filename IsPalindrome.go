package Leetcode

import "strconv"

func isPalindrome(x int) bool {
	stringX := strconv.Itoa(x)

	for i, value := range stringX {
		if int32(stringX[len(stringX) - (i + 1)]) != value {
			return false
		}
	}

	return true
}