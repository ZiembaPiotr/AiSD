func minimumSum(num int) int {
    digits := splitSort(num)
    
    return (digits[0] + digits[1])*10 + digits[2] + digits[3]    
}

func splitSort(n int) []int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
    }
    
    sort.Ints(digits)
    
    return digits
}
