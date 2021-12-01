package Leetcode

import "strings"

func WordCount(s string) map[string]int {
	stringTab := strings.Fields(s)
	wordCount := make(map[string]int)
	for _, word := range stringTab {
		wordCount[word] += 1
	}
	return wordCount
}
