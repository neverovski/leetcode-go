package problem100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stackP, stackQ := make([]*TreeNode, 0, 5), make([]*TreeNode, 0, 5)

	currP, currQ := p, q

	for currP != nil || currQ != nil || (len(stackP) > 0 && len(stackQ) > 0) {
		if currP != nil && currQ != nil {
			stackP, stackQ = append(stackP, currP), append(stackQ, currQ)

			currP, currQ = currP.Left, currQ.Left
		} else if currP == nil && currQ == nil {
			currP, stackP = stackP[len(stackP)-1], stackP[:len(stackP)-1]
			currQ, stackQ = stackQ[len(stackQ)-1], stackQ[:len(stackQ)-1]

			if currP.Val != currQ.Val {
				return false
			}

			currP, currQ = currP.Right, currQ.Right
		} else {
			return false
		}
	}

	return true
}

func isSameTreeUsingRecursive(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
