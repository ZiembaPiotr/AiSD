func shuffle(nums []int, n int) []int {
    result := make([]int, 2*n)
    var oddCounter int
    var evenCounter int
    
    for i, _ := range result {
        if i%2 == 0 {
            result[i] = nums[i - evenCounter]
            evenCounter++
            continue
        }
        result[i] = nums[oddCounter + n]
        oddCounter++
    }
    
    return result
}
