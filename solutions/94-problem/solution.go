package solutions

import (
	"github.com/neverovski/leetcode-go/structures"
)

type TreeNode = structures.TreeNode

func inorderTraversalUsingLoop(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := []int{}
	stackTree := []*TreeNode{}

	for len(stackTree) != 0 || root != nil {
		if root != nil {
			stackTree = append(stackTree, root)
			root = root.Left
		} else {
			index := len(stackTree) - 1

			if index >= 0 {
				root = stackTree[index]
				stackTree = stackTree[:index]
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
