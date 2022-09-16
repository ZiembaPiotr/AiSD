func countPoints(points [][]int, queries [][]int) []int {
    answer := make([]int, len(queries))
    
    for i, query := range queries {
        for _, point := range points {
            if (square(query[0] - point[0]) + square(query[1] - point[1]) <= square(query[2])) {
                answer[i] += 1
            }
        }
    }
    
    
    return answer
}
                
func square(n int) int {
    return n*n                    
}
