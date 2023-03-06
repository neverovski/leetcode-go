package solutions

import (
	"github.com/neverovski/leetcode-go/structures"
)

type TreeNode = structures.TreeNode
type Queue = structures.Queue[TreeNode]

func maxDepthLoop(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := Queue{root}
	maxDepth := 0

	for !queue.IsEmpty() {
		size := len(queue)
		maxDepth += 1

		for i := 0; i < size; i++ {
			if node, ok := queue.Poll(); ok {
				if node.Left != nil {
					queue.Add(node.Left)
				}

				if node.Right != nil {
					queue.Add(node.Right)
				}
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
