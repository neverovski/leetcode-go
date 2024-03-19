package problem083

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesUsingOneList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func deleteDuplicatesUsingTwoList(listIn *ListNode) *ListNode {
	if listIn == nil {
		return nil
	}

	listOut := &ListNode{Val: listIn.Val}

	processList(listIn, listOut)

	return listOut
}

func processList(currentIn, currentOut *ListNode) {
	for currentIn != nil {
		if currentIn.Val != currentOut.Val {
			currentOut.Next = &ListNode{Val: currentIn.Val}

			currentOut = currentOut.Next
		}

		currentIn = currentIn.Next
	}
}
