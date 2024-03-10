package problem021

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	nodes := append(listToSlice(list1), listToSlice(list2)...)

	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i] < nodes[j]
	})

	return sliceToList(nodes)
}

func listToSlice(list *ListNode) []int {
	count := countList(list)

	if count == 0 {
		return []int{}
	}

	nodes := make([]int, 0, count)
	current := list

	for current != nil {
		nodes = append(nodes, current.Val)
		current = current.Next
	}

	return nodes
}

func sliceToList(nodes []int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}

	list := &ListNode{Val: nodes[0]}
	current := list

	for i := 1; i < len(nodes); i++ {
		current.Next = &ListNode{Val: nodes[i]}

		current = current.Next
	}

	return list
}

func countList(list *ListNode) int {
	count := 0
	current := list

	for current != nil {
		current = current.Next
		count++
	}

	return count
}
