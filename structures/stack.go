package structures

type Stack[T any] []*T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(n *T) {
	*s = append(*s, n)
}

func (s *Stack[T]) Pop() (*T, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return node, true
}
