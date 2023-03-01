package solutions

import structures "github.com/neverovski/leetcode-go/structures"

const MIN_NUMBER_NODE = 0
const MAX_NUMBER_NODE = 9

type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}

	current := l3
	carry := 0

	for l1 != nil || l2 != nil {
		val1 := 0
		if l1 != nil && l1.Val >= MIN_NUMBER_NODE && l1.Val <= MAX_NUMBER_NODE {
			val1 = l1.Val
		}

		val2 := 0
		if l2 != nil && l2.Val >= MIN_NUMBER_NODE && l2.Val <= MAX_NUMBER_NODE {
			val2 = l2.Val
		}

		sum := val1 + val2 + carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return l3.Next
}
