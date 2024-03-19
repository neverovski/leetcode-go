package problem111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
* The idea is to traverse the given Binary Tree.
* For every node, check if it is a leaf node.
* If yes, then return 1.
* If not leaf node then if the left subtree is NULL, then recur for the right subtree. And if the right subtree is NULL, then recur for the left subtree.
* If both left and right subtrees are not NULL, then take the minimum of two depths.
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
