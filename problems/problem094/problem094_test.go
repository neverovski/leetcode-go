package problem094

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversalUsingStack(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{1, 3, 2}, inorderTraversalUsingStack(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	a.Equal([]int{}, inorderTraversalUsingStack(nil))
	a.Equal([]int{1}, inorderTraversalUsingStack(&TreeNode{Val: 1}))
}

func TestInorderTraversalUsingSlice(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{1, 3, 2}, inorderTraversalUsingSlice(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	a.Equal([]int{}, inorderTraversalUsingRecursive(nil))
	a.Equal([]int{1}, inorderTraversalUsingRecursive(&TreeNode{Val: 1}))
}

func TestInorderTraversalUsingRecursive(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{1, 3, 2}, inorderTraversalUsingRecursive(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	a.Equal([]int{}, inorderTraversalUsingRecursive(nil))
	a.Equal([]int{1}, inorderTraversalUsingRecursive(&TreeNode{Val: 1}))
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
