package problem144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := append([]int{root.Val}, preorderTraversal(root.Left)...)
	nums = append(nums, preorderTraversal(root.Right)...)

	return nums
}

func preorderTraversalUsingSlice(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	nums := make([]int, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			nums = append(nums, root.Val)
			stack = append(stack, root)

			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		root = root.Right
	}
	return nums
}
