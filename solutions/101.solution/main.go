package solutions

import (
	"github.com/neverovski/leetcode-go/structures"
)

type TreeNode = structures.TreeNode

func isSymmetricLoop(root *TreeNode) bool {
	if root == nil {
		return true
	}

	arrTree := []*TreeNode{root.Left, root.Right}

	for len(arrTree) != 0 {
		var leftNode *TreeNode
		indexLeft := len(arrTree) - 1

		if indexLeft >= 0 {
			leftNode = arrTree[indexLeft]
			arrTree = arrTree[:indexLeft]
		}

		var rightNode *TreeNode
		indexRight := len(arrTree) - 1

		if indexRight >= 0 {
			rightNode = arrTree[indexRight]
			arrTree = arrTree[:indexRight]
		}

		if leftNode == nil && rightNode == nil {
			continue
		}

		if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
			return false
		}

		arrTree = append(
			arrTree,
			leftNode.Left,
			rightNode.Right,
			leftNode.Right,
			rightNode.Left,
		)
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
