package problem112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	subSum := targetSum - root.Val

	if subSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, subSum) || hasPathSum(root.Right, subSum)
}

func hasPathSumUsingStack(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{root}
	sums := []int{root.Val}

	for len(stack) > 0 {
		index := len(sums) - 1
		sum := sums[index]
		node := stack[index]

		if node.Left == nil && node.Right == nil && targetSum-sum == 0 {
			return true
		}

		sums = sums[:index]
		stack = stack[:index]

		if node.Left != nil {
			sums = append(sums, sum+node.Left.Val)
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			sums = append(sums, sum+node.Right.Val)
			stack = append(stack, node.Right)
		}
	}

	return false
}
