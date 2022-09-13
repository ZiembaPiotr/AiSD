type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	var merged *ListNode
	var partSum int

	iterator := head.Next

	for ; iterator != nil; iterator = iterator.Next {
		partSum += iterator.Val
		if iterator.Val == 0 {
			merged.addNode(partSum)
			fmt.Println(merged)
			partSum = 0
		}
	}

	return merged
}

func (linkedList *ListNode) addNode(value int) {
	fmt.Println(linkedList)

	if linkedList == nil {
		linkedList.Val = value
		return
	}

	newNode := new(ListNode)
	newNode.Val = value

	iterator := linkedList.Next
	for ; iterator.Next != nil; iterator = iterator.Next {

	}

	iterator.Next = newNode
}
