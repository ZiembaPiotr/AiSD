func maximumWealth(accounts [][]int) int {
    var max int
    
    for _, account := range accounts {
        var sum int
        
        for _, amount := range account {
            sum += amount
        }
        
        if max < sum {
            max = sum
        }
    }
    
    return max
}
