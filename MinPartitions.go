func minPartitions(n string) int {
    var max int
    
    for _, number := range n {
        decodedNumber := int(number) - 48
        
        if decodedNumber == 9 {
            return 9
        }
        if max < decodedNumber {
            max = decodedNumber
        }
    }
    
    return max
}
