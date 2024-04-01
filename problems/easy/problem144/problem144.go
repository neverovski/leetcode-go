package problem144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalUsingSlice(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)

	nums := make([]int, 0)
	current := root

	for current != nil || len(stack) > 0 {
		if current != nil {
			nums = append(nums, current.Val)

			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			current = current.Right
		}
	}

	return nums
}

func preorderTraversalUsingRecursive(root *TreeNode) []int {
	nums := make([]int, 0)

	preorderTraversal(root, &nums)

	return nums
}

func preorderTraversal(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	*nums = append(*nums, root.Val)
	preorderTraversal(root.Left, nums)
	preorderTraversal(root.Right, nums)
}
