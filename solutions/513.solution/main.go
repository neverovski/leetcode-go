package solutions

import "github.com/neverovski/leetcode-go/structures"

type TreeNode = structures.TreeNode
type Queue = structures.Queue[TreeNode]

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := Queue{root}
	bottomLeftVal := 0

	for !queue.IsEmpty() {
		if node, ok := queue.Poll(); ok {
			bottomLeftVal = node.Val

			if node.Right != nil {
				queue.Add(node.Right)
			}

			if node.Left != nil {
				queue.Add(node.Left)
			}
		}
	}

	return bottomLeftVal
}
