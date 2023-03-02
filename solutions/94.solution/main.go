package solutions

import (
	"github.com/neverovski/leetcode-go/structures"
)

type TreeNode = structures.TreeNode

func inorderTraversalLoop(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := []int{}
	arrTree := []*TreeNode{}

	for len(arrTree) != 0 || root != nil {
		if root != nil {
			arrTree = append(arrTree, root)
			root = root.Left
		} else {
			index := len(arrTree) - 1

			if index >= 0 {
				root = arrTree[index]
				arrTree = arrTree[:index]
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
