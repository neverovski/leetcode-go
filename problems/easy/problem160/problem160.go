package problem160

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashNodeA := make(map[*ListNode]struct{}, 0)

	for headA != nil {
		hashNodeA[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := hashNodeA[headB]; ok {
			return headB
		}

		headB = headB.Next
	}

	return nil
}
