package problem101

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	a := assert.New(t)

	a.Equal(true, isSymmetric(&TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}))
	a.Equal(false, isSymmetric(&TreeNode{Val: 2, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}))
}

func BenchmarkIsSymmetric(b *testing.B) {
	tree := &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isSymmetric(tree)
	}
}
