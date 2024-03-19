package stack

import "container/list"

type Stack struct {
	items *list.List
}

func New() *Stack {
	return &Stack{items: list.New()}
}

func (s *Stack) Push(item any) {
	s.items.PushBack(item)
}

func (s *Stack) Pop() any {
	if s.Len() == 0 {
		return nil
	}

	tail := s.items.Back()
	val := tail.Value

	s.items.Remove(tail)

	return val
}

func (s *Stack) Peek() any {
	if s.Len() == 0 {
		return nil
	}

	return s.items.Back().Value
}

func (s *Stack) Len() int {
	return s.items.Len()
}
