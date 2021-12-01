package Leetcode

func main(){
	
}

func twoSum(nums []int, target int) []int{
	valueIndexMap := make(map[int]int)

	for i, num := range nums {
		if pos, ok := valueIndexMap[target - num]; ok{
			return []int{pos, i}
		} else {
			valueIndexMap[num] = i
		}
	}

	return nil
}