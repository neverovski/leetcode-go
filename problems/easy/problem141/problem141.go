package problem141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head.Next
	fast := slow.Next

	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next

		if fast != nil {
			fast = fast.Next
		}
	}

	return false
}
