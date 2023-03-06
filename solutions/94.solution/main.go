package solutions

import (
	"github.com/neverovski/leetcode-go/structures"
)

type TreeNode = structures.TreeNode
type Stack = structures.Stack[TreeNode]

func inorderTraversalLoop(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := []int{}
	stack := Stack{}

	for !stack.IsEmpty() || root != nil {
		if root != nil {
			stack.Push(root)
			root = root.Left
		} else {
			if node, ok := stack.Pop(); ok {
				root = node
			}

			nums = append(nums, root.Val)
			root = root.Right
		}
	}

	return nums
}

func inorderTraversalRecursive(root *TreeNode) []int {
	nums := []int{}
	inorderRecursive(root, &nums)

	return nums
}

func inorderRecursive(root *TreeNode, out *[]int) {
	if root != nil {
		inorderRecursive(root.Left, out)
		*out = append(*out, root.Val)
		inorderRecursive(root.Right, out)
	}
}
