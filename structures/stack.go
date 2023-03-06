package structures

type Stack[T any] []*T

func (q *Stack[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Stack[T]) Push(n *T) {
	*q = append(*q, n)
}

func (q *Stack[T]) Pop() (*T, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	node := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]

	return node, true
}
