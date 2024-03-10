package problem101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	stackL, stackR := []*TreeNode{}, []*TreeNode{}

	l, r := root.Left, root.Right

	for l != nil || r != nil || (len(stackL) > 0 && len(stackR) > 0) {
		if l != nil && r != nil {
			stackL, stackR = append(stackL, l), append(stackR, r)

			l, r = l.Left, r.Right
		} else if l == nil && r == nil {
			l, stackL = stackL[len(stackL)-1], stackL[:len(stackL)-1]
			r, stackR = stackR[len(stackR)-1], stackR[:len(stackR)-1]

			if l.Val != r.Val {
				return false
			}

			l, r = l.Right, r.Left
		} else {
			return false
		}
	}

	return true
}
