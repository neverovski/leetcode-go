package problem145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)
	nums = append(nums, root.Val)

	return nums
}

func postorderTraversalUsingSlice(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	nums := make([]int, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			nums = append([]int{root.Val}, nums...)
			stack = append(stack, root)

			root = root.Right
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		root = root.Left
	}
	return nums
}
