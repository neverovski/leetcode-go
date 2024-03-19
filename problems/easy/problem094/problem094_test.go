package problem094

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  *TreeNode
	out []int
}

func TestInorderTraversalUsingStack(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			out: []int{1, 3, 2},
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
		a.Equal(tc.out, inorderTraversalUsingStack(tc.in), fmt.Sprintf("TestInorderTraversalUsingStack number # %d", key+1))
	}
}

func TestInorderTraversalUsingSlice(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			out: []int{1, 3, 2},
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
		a.Equal(tc.out, inorderTraversalUsingSlice(tc.in), fmt.Sprintf("TestInorderTraversalUsingSlice number # %d", key+1))
	}
}

func TestInorderTraversalUsingRecursive(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			out: []int{1, 3, 2},
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
		a.Equal(tc.out, inorderTraversalUsingRecursive(tc.in), fmt.Sprintf("TestInorderTraversalUsingRecursive number # %d", key+1))
	}
}

func BenchmarkInorderTraversalUsingStack(b *testing.B) {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		inorderTraversalUsingStack(root)
	}
}

func BenchmarkInorderTraversalUsingSlice(b *testing.B) {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		inorderTraversalUsingSlice(root)
	}
}

func BenchmarkInorderTraversalUsingRecursive(b *testing.B) {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		inorderTraversalUsingRecursive(root)
	}
}
