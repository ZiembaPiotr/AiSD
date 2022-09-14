// dummy

func mergeNodes(head *ListNode) *ListNode {
    var partSum int
    var sumArray []int
    merged := &ListNode{}
    prevNode := &ListNode{}
    
    iterator := head.Next
    
    for ; iterator != nil; iterator = iterator.Next {
        partSum += iterator.Val
        if iterator.Val == 0 {
            sumArray = append(sumArray, partSum)
            partSum = 0
        }
    }
    
    for i, sum := range sumArray {
        newNode := &ListNode{sum, nil}
        
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

// less dummy


func mergeNodes(head *ListNode) *ListNode {
    var partSum int
    
    merged := &ListNode{}
    prevNode := merged
    
    for head.Next != nil {
        partSum += head.Val
        
        if head.Val == 0 {
            prevNode.Next = &ListNode{partSum, nil}
            prevNode = prevNode.Next
            partSum = 0
        }

        head = head.Next
    }
    
    prevNode.Next = &ListNode{partSum, nil}
    
    return merged.Next.Next
}
