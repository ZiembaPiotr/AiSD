package Leetcode

func twoSum(nums []int, target int) []int{
	valueIndexMap := make(map[int]int)

	for i, num := range nums {
		if pos, exist := valueIndexMap[target - num]; exist{
			return []int{pos, i}
		} else {
			valueIndexMap[num] = i
		}
	}

	return nil
}