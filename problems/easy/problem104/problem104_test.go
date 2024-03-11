package problem104

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	a := assert.New(t)

	a.Equal(0, maxDepth(nil), "Should return 0 for an empty tree")
	a.Equal(1, maxDepth(&TreeNode{Val: 1}), "Should return 1 for a tree with one node")
	a.Equal(2, maxDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}), "Should return 2 for a tree with two levels")
	a.Equal(2, maxDepth(&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}), "Should return 2 for a tree with two levels")
	a.Equal(3, maxDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}), "Should return 3 for a tree with three levels")
	a.Equal(3, maxDepth(&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}), "Should return 3 for a tree with three levels")
	a.Equal(4, maxDepth(&TreeNode{Val: 0, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: -1, Right: &TreeNode{Val: 8}}}}), "Should return 4 for a tree with four levels")
}

func BenchmarkMaxDepth(b *testing.B) {
	tree := &TreeNode{Val: 0, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: -1, Right: &TreeNode{Val: 8}}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maxDepth(tree)
	}
}
