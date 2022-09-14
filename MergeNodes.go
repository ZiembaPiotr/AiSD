func mergeNodes(head *ListNode) *ListNode {
    var partSum int
    var sumArray []int
    merged := new(ListNode)
    prevNode := new(ListNode)
    
    iterator := head.Next
    
    for ; iterator != nil; iterator = iterator.Next {
        partSum += iterator.Val
        if iterator.Val == 0 {
            sumArray = append(sumArray, partSum)
            partSum = 0
        }
    }
    
    for i, sum := range sumArray {
        newNode := new(ListNode)
        newNode.Val = sum
        
        if i == 0 {
            merged = newNode
            prevNode = merged
            continue
        }
        
        prevNode.Next = newNode 
        
        prevNode = newNode
        
    }
    
    return merged
}
