package Leetcode

func romanToInt(s string) int {
	romanianValueMap := map[string]int{"I":1, "V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
	var valueArray []int
	var prevValue int
	var nextValue int
	var result int

	for i := 0; i < len(s); i++ {
		valueArray = append(valueArray, romanianValueMap[string(s[i])])
	}

	for i, val := range valueArray {
		if len(valueArray) - i >= 2 {
			prevValue = valueArray[i]
			nextValue = valueArray[i + 1]
			if prevValue < nextValue {
				result -= prevValue
			} else {
				result += prevValue
			}
		} else {
			result += val
		}
	}
	return result
}
