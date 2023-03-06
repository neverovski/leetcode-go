package structures

type Queue[T any] []*T

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Add(n *T) {
	*q = append(*q, n)
}

func (q *Queue[T]) Poll() (*T, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	node := (*q)[0]
	*q = (*q)[1:]

	return node, true
}
