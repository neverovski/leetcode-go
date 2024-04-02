package problem145

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  *TreeNode
	out []int
}

func TestPostorderTraversal(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			out: []int{3, 2, 1},
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
		a.Equal(tc.out, postorderTraversal(tc.in), fmt.Sprintf("TestPostorderTraversal number # %d", key+1))
	}
}

func TestPostorderTraversalUsingSlice(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			out: []int{3, 2, 1},
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
		a.Equal(tc.out, postorderTraversalUsingSlice(tc.in), fmt.Sprintf("TestPostorderTraversalUsingSlice number # %d", key+1))
	}
}

func BenchmarkPostorderTraversal(b *testing.B) {
	node := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	for i := 0; i < b.N; i++ {
		postorderTraversal(node)
	}
}

func BenchmarkPostorderTraversalUsingSlice(b *testing.B) {
	node := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	for i := 0; i < b.N; i++ {
		postorderTraversalUsingSlice(node)
	}
}
