package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPop(t *testing.T) {
	a := assert.New(t)

	stack := New()
	stack.Push("Item1")
	stack.Push("Item2")
	stack.Push("Item3")

	a.Equal("Item3", stack.Peek())
	a.Equal("Item3", stack.Pop())

	a.Equal("Item2", stack.Peek())
	a.Equal("Item2", stack.Pop())

	a.Equal("Item1", stack.Peek())
	a.Equal("Item1", stack.Pop())

	a.Nil(stack.Pop())
}
