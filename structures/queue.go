package structures

type Queue []*TreeNode

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(n *TreeNode) {
	*q = append(*q, n)
}

func (q *Queue) Dequeue() (*TreeNode, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	node := (*q)[0]
	*q = (*q)[1:]

	return node, true
}
