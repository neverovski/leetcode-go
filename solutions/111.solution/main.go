package solutions

import "github.com/neverovski/leetcode-go/structures"

type TreeNode = structures.TreeNode
type Queue = structures.Queue

func minDepthQueue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := Queue{root}
	minDepth := 1

	for !queue.IsEmpty() {
		size := len(queue)

		for i := 0; i < size; i++ {
			if node, ok := queue.Dequeue(); ok {
				if node.Left == nil && node.Right == nil {
					return minDepth
				}

				if node.Left != nil {
					queue.Enqueue(node.Left)
				}

				if node.Right != nil {
					queue.Enqueue(node.Right)
				}
			}
		}

		minDepth += 1
	}

	return minDepth
}
