package Leetcode

func longestCommonPrefix(strs []string) string {
	var shortestStringIndex int
	var commonPrefix string
	shortestStringLength := len(strs[0])

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs); i++ {
		if len(strs[i]) <= shortestStringLength {
			shortestStringLength = len(strs[i])
			shortestStringIndex = i
		}
	}

	for i := 0; i < shortestStringLength; i++{
		currentlyChecked := strs[shortestStringIndex][i]
		for j, _ := range strs {
			if strs[j][i] != currentlyChecked {
				return strs[0][:i]
			} else {
				commonPrefix = strs[0][:i + 1]
			}
		}
	}

	return commonPrefix
}
