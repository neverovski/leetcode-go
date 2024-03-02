package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPushAndPop(t *testing.T) {
	a := assert.New(t)
	s := New()

	a.Equal(0, s.Len(), "Stack should be empty initially")

	s.Push(1)
	s.Push(2)
	s.Push(3)

	a.Equal(3, s.Len(), "Stack should not be empty after pushing elements")

	a.Equal(3, s.Pop(), "Popped value should be 3")
	a.Equal(2, s.Pop(), "Popped value should be 2")
	a.Equal(1, s.Pop(), "Popped value should be 1")

	a.Equal(0, s.Len(), "Stack should be empty after popping all elements")
}

func TestStackPeek(t *testing.T) {
	a := assert.New(t)
	s := New()

	a.Nil(s.Peek(), "Stack is empty")

	s.Push(10)
	a.Equal(10, s.Peek(), "Peeked value should be 10")
}
