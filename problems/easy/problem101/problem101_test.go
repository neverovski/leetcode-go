package problem101

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  *TreeNode
	out bool
}

func TestIsSymmetric(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
			out: true,
		},
		{
			in:  &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			out: false,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, isSymmetric(tc.in), fmt.Sprintf("TestIsSymmetric number # %d", key+1))
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	tree := &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isSymmetric(tree)
	}
}
