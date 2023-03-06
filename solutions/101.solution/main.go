package solutions

import (
	"github.com/neverovski/leetcode-go/structures"
)

type TreeNode = structures.TreeNode
type Stack = structures.Stack[TreeNode]

func isSymmetricLoop(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := Stack{root.Left, root.Right}

	for !stack.IsEmpty() {
		var leftNode *TreeNode
		if node, ok := stack.Pop(); ok {
			leftNode = node
		}

		var rightNode *TreeNode
		if node, ok := stack.Pop(); ok {
			rightNode = node
		}

		if leftNode == nil && rightNode == nil {
			continue
		}

		if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
			return false
		}

		stack.Push(leftNode.Left)
		stack.Push(rightNode.Right)
		stack.Push(leftNode.Right)
		stack.Push(rightNode.Left)
	}

	return true
}

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSameTree(root.Left, root.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
}
