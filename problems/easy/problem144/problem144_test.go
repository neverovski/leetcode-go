package problem144

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  *TreeNode
	out []int
}

func TestPreorderTraversalUsingRecursive(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			out: []int{1, 2, 3},
		},
		{
			in:  nil,
			out: []int{},
		},
		{
			in:  &TreeNode{Val: 1},
			out: []int{1},
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, preorderTraversalUsingRecursive(tc.in), fmt.Sprintf("TestPreorderTraversalUsingRecursive number # %d", key+1))
	}
}

func TestPreorderTraversalUsingSlice(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			out: []int{1, 2, 3},
		},
		{
			in:  nil,
			out: []int{},
		},
		{
			in:  &TreeNode{Val: 1},
			out: []int{1},
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, preorderTraversalUsingSlice(tc.in), fmt.Sprintf("TestPreorderTraversalUsingSlice number # %d", key+1))
	}
}

func BenchmarkPreorderTraversalUsingRecursive(b *testing.B) {
	node := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	for i := 0; i < b.N; i++ {
		preorderTraversalUsingRecursive(node)
	}
}

func BenchmarkPreorderTraversalUsingSlice(b *testing.B) {
	node := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	for i := 0; i < b.N; i++ {
		preorderTraversalUsingSlice(node)
	}
}
