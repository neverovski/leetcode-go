package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func TransformIntToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	temp := l

	for _, num := range nums {
		temp.Next = &ListNode{Val: num}
		temp = temp.Next
	}

	return l.Next
}

func TransformListToInt(l *ListNode) []int {
	nums := []int{}

	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}

	return nums
}
