package solutions

import (
	"github.com/neverovski/leetcode-go/structures"
)

type TreeNode = structures.TreeNode

func maxDepthLoop(root *TreeNode) int {
	if root == nil {
		return 0
	}

	arrTree := []*TreeNode{root}
	maxDepth := 0

	for len(arrTree) != 0 {
		size := len(arrTree)
		maxDepth += 1

		for i := 0; i < size; i++ {
			node := arrTree[0]
			arrTree = arrTree[1:]

			if node.Left != nil {
				arrTree = append(arrTree, node.Left)
			}

			if node.Right != nil {
				arrTree = append(arrTree, node.Right)
			}
		}
	}

	return maxDepth
}

func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepthRecursive(root.Left)
	rightDepth := maxDepthRecursive(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}
