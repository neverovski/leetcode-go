package problem094

import "leetcode/go/utils/stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalUsingStack(root *TreeNode) []int {
	stack := stack.New() // for decrease allocation

	nums := make([]int, 0, 5)
	current := root

	for current != nil || stack.Len() > 0 {
		if current != nil {
			stack.Push(current)
			current = current.Left
		} else {
			current = stack.Pop().(*TreeNode)

			nums = append(nums, current.Val)
			current = current.Right
		}
	}

	return nums
}

func inorderTraversalUsingSlice(root *TreeNode) []int {
	stack := make([]*TreeNode, 0, 5) // for decrease allocation

	nums := make([]int, 0, 5)
	current := root

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			nums = append(nums, current.Val)
			current = current.Right
		}
	}

	return nums
}

func inorderTraversalUsingRecursive(root *TreeNode) []int {
	nums := make([]int, 0, 5) // for decrease allocation

	inorderTraversal(root, &nums)

	return nums
}

func inorderTraversal(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorderTraversal(root.Right, nums)
}
